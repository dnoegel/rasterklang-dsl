# ADR 0004: `ramp` saturation is not GoatTracker's filter-table wrap (documentation correction)

Status: documented (corrects `docs/dsl-grammar-spec.md` and `docs/dsl-changelog.md`; no grammar,
parser or interpreter change)
Date: 2026-08-16
Affects: `rasterklang-dsl` docs only (`docs/dsl-grammar-spec.md`, `docs/dsl-changelog.md`).
Current DSL version: 4.0.0. This ADR changes no grammar-visible behavior and does not bump it.

---

## 1. Context

`docs/dsl-grammar-spec.md`'s "Writing long tables: `ramp` and `!`" section, and the `2.12.0`
entry in `docs/dsl-changelog.md`, both claimed that `ramp`'s saturating clamp is not a DSL
convenience but a faithful reproduction of GoatTracker's own filter/pulse modulation-step command,
down to stating that one `ramp` token "reproduces one original tracker row exactly." That claim is
false for the filter table.

The counter-evidence came from reading GoatTracker 2's player source (`gplay.c`) against a
hardware-accurate register capture of a GoatTracker tune from the High Voltage SID Collection
(10_Orbyte): `$D416` changed 592 times over 600 frames in the original, versus 0 times in an
earlier, wrong conversion. `gplay.c`'s filter block shows the cutoff accumulator wraps; it does
not saturate. This ADR records the correction and the reasoning for keeping `ramp` exactly as it
is.

## 2. The two things being conflated

The false claim conflates two mechanisms that live at different layers and have different
contracts:

- **`ramp` is a DSL table generator.** `ramp A..B step D over N` (`docs/dsl-grammar-spec.md`,
  `interpret.js`'s `expandRamp`) is a pure, compile-time expansion that produces a flat array of
  **absolute per-frame values**, clamped to `[min(A,B), max(A,B)]`. The engine never sees the
  generator, only the expanded array. It has no concept of bit width and no concept of "the value
  after the last one."
- **GoatTracker's wrap lives in the player's 8-bit `filtercutoff` accumulator.** `gplay.c` stores
  the running cutoff in a plain `unsigned char`. A modulation-step row ("for N player calls, add D
  per call") does `filtercutoff += D` every call, in 8-bit arithmetic, for as long as the song
  runs or another row changes it. There is no clamp in that addition. What looks like a ceiling in
  the tracker row semantics (`$01-$7F` = duration + signed delta) is a statement about how long
  that one row keeps stepping, not a statement that the accumulator stops moving.

Saturating at the target is a legitimate design choice for a DSL table generator: it is
predictable, it is what most authors reaching for "sweep from A to B" want, and it matches how the
generator's own duration parameter (`over N`) reads. The defect was never the saturation. It was
claiming that saturation is what the original player does, and that a single `ramp` is
bit-for-bit equivalent to a tracker row.

## 3. What is actually true

From `gplay.c`'s filter block:

- `filtercutoff` is 8 bits and wraps. It is written straight to `$D416` (cutoff high), with
  `$D415` (cutoff low) written as `0x00` every frame, so the effective 11-bit cutoff is
  `filtercutoff * 8`, with modulo-256 arithmetic on `filtercutoff`.
- A long filter modulation-step row therefore does not settle onto a value and hold. In
  effective-cutoff terms it is a **continuous sawtooth**: linear descent (or ascent), a vertical
  jump when the 8-bit accumulator overflows or underflows, and the same linear motion resuming
  from the other end, for as long as the row keeps executing. Since a `$FF` row can jump back into
  the same rows, that can be forever.
- A `ramp`'s saturating tail (holding flat at the target) is the opposite shape: a one-shot sweep
  followed by no motion at all in that table.

Two related facts from the same code, which the spec's other filter-table statements depend on:
`filtertype` is masked with `$70`, not `$F0` (the SID's own `$D418` passband bits), and `$FF` in a
filter-table row is a JUMP that continues execution in the same player call, not a loop marker or
terminator. Only `$FF 00` stops, and it freezes the cutoff at its last value.

## 4. Consequence for converted songs, worked example

A filter sweep of the following shape is exactly what a GoatTracker import produces for a long
filter modulation-step row, and the conformance fixture
`test/conformance/fixtures/table_decl_ramp_repeat.dsl` uses the same expression:

```
table sharedFilterTable = ramp 1088..0 step -48 over 128
```

In DSL units this is 1088 down to 0, step -48, i.e. GoatTracker cutoff units (DSL value / 8) of
`136` (`$88`) stepping `-6` per frame, since the DSL's cutoff is `filtercutoff * 8` throughout.
Working the saturating arithmetic through: `136 - 6n <= 0` first holds at `n = 23`
(`136 - 6*23 = -2`, clamped to 0), so the ramp reaches 0 after 23 frames and then holds flat for
the remaining `128 - 23 = 105` frames. **82% of the table's declared length (105 of 128 frames) is
a flat tail.**

A wrapping equivalent, matching what the GoatTracker accumulator does with the same step and no
clamp, does not hold at 0. At frame 23 (value 0), the next frame subtracts 6 more and, in 8-bit
arithmetic, lands at 250 (`0 - 6 mod 256`), so the sweep re-enters near the top of the range and
keeps descending: a repeating sawtooth for as long as the row is live, not a single 23-frame fall
followed by 105 frames without motion.

This is the exact example the original claim cited as its proof.

## 5. Decision

**Keep `ramp`'s saturating clamp exactly as it is.** It is a predictable, well-understood generator
primitive, every existing `.dsl` song that uses `ramp` was authored (or converted) against that
behavior, and changing the default would silently change the sound of existing songs with no
diagnostic. The defect was in the documentation's claim about what the behavior reproduces, not in
the behavior.

Instead:

1. **Correct the documentation** (`docs/dsl-grammar-spec.md`, `docs/dsl-changelog.md`) to state
   the true relationship: `ramp` saturates by design, GoatTracker's filter accumulator wraps by
   hardware width, and a `ramp` reproduces a bounded tracker row's *shape* only when the row's own
   step count already keeps it inside the 8-bit range. That is not a general guarantee.
2. **Record that a wrapping sawtooth is already expressible**, without any clamp change, using
   `ramp`'s own `loop` shape terminator (ADR 0003 D9, `expandRamp` in `interpret.js`). A shorter
   ramp written to land at or near the wrap boundary, e.g. `ramp 1088..0 step -48 over 23 loop`,
   sets `loopIndex 0`, so once the instrument's table sampling advances past the ramp's last
   entry, it restarts at `from` rather than holding. This produces the same qualitative
   repeating-sawtooth shape as the hardware wrap: linear descent, vertical reset, repeat. It is
   not bit-identical to the modulo-256 accumulator: a `loop` restart snaps back to the exact
   `from` value each cycle, while the hardware accumulator carries whatever remainder the last step
   left, so the real sweep drifts by a fraction of one cutoff step per lap. For the point being
   corrected ("saturates and holds" versus "keeps moving"), `loop` gives the right answer.
3. **Note that this idiom was blocked inside a shared `table NAME = ...` declaration** at the time
   of writing: a ramp's implied loop point was not threaded through a `$name` table reference, only
   through a ramp written inline in the instrument field. That gap was closed in 4.1.0 (see
   `docs/dsl-changelog.md`), which makes the wrapping idiom available to shared and
   converter-hoisted tables too.

## 6. Scope of impact

- **Playback engines are not affected.** Emulators and players that consume already-resolved
  `$D415`/`$D416` register writes never evaluate a `ramp` token or any other DSL table generator.
  The false claim, and this correction, have no bearing on emulated playback correctness.
- **The DSL's own documentation is the direct defect.** `src/dsl/interpret.js`'s `expandRamp`
  carries the same claim in its source comment. The code's behavior (the saturating clamp) is
  correct and is not changed by this ADR. This ADR's scope is documentation only; the source
  comment is to be brought in line the next time that function is touched.
- **Downstream reimplementations inherit the same saturating design, not the false claim.** Any
  tool that reimplements `ramp`'s semantics (for example a sweep editor in a UI) inherits the
  same defensible design. Documentation for such a tool that repeats the "reproduces one tracker
  row exactly" claim needs the same correction.

## 7. Documentation changes made

- `docs/dsl-grammar-spec.md`, "Writing long tables: `ramp` and `!`" section: replaced the paragraph
  claiming saturation "isn't a convenience" and reproduces "one original tracker row exactly" with
  a paragraph stating the true relationship (saturates by design; GoatTracker's accumulator wraps;
  a `ramp` matches a tracker row only when the row stays inside range), and cross-referenced this
  ADR.
- `docs/dsl-changelog.md`, `2.12.0` entry: the same claim, made about the same
  `ramp 1088..0 step -48 over 128` example, is corrected in place with a note pointing at this
  ADR, without altering the record of what `2.12.0` shipped (the `ramp`/`V!N` syntax and its
  saturating behavior are unchanged; only the claim about what that behavior represents is
  corrected).
- A changelog entry records this as a documentation-only correction, explicitly not a grammar or
  semantic version bump.
