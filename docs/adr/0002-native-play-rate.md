# ADR 0002: Native play rate (NTSC, multispeed, and custom-tempo drivers)

Status: implemented (2026-08-02)
Date: 2026-08-02
Affects: `rasterklang-dsl` (grammar, `pkg/dsl`, `src/dsl`), `rasterklang-score`
(`src/compiler/compile.js`, `src/engine/liveEngine.js`), and any tool that emits DSL from a
SID file.
Current DSL version: 2.19.0. Shipped in 3.1.0.

---

## 1. Context

`compile.js` derives `framesPerCycle` from a single hard-coded constant: DSL v2
assumes every song's player is invoked exactly 50.00 times per second (PAL
video frame rate). ADR 0001 already names this as a known one-time constant
that a future change will need to touch.

That assumption is false for a measurable, non-trivial slice of real SID
tunes, in three distinct ways:

1. **NTSC.** Machines clocked at 1,022,727 Hz with a 60 Hz frame rate exist
   throughout the High Voltage SID Collection (HVSC); the video standard is a
   per-file property, not a global constant. Re-synthesising an NTSC tune
   through DSL v2 plays it at roughly 50/60 = 0.83x real speed, because the
   compiler's `framesPerCycle` never asks which video standard the source
   declared.
2. **Integer multispeed.** GoatTracker and other trackers commonly invoke the
   player 2, 3 or 4 times per video frame ("funktempo"/multispeed). Owen
   Crowley's `Danger_Zone.sid` (3x) and `Interest.sid` (4x) are two measured
   examples. DSL v2 cannot represent this.
3. **Non-video-locked custom tempo.** Some players are driven by a CIA timer
   set to an arbitrary reload value with no relationship to the video frame
   rate at all. A concrete example: `Welle_Erdball_Alles_Luege.sid`
   (Soundmonitor), whose play routine's CIA timer A latch is 12,800/12,801
   cycles, i.e. **76.97 Hz** (about 1.54 calls per 19,705-cycle PAL frame):
   not 50, not 60, not an integer multiple of either. This is what a
   CIA-driven player looks like, and DSL v2 has no representation for it: a
   custom-tempo file is force-fit to 50 Hz with no error, exactly like the NTSC
   case, just with a worse ratio.

All three are the same underlying gap: **DSL v2 has one play rate, hard-coded,
and real tunes have many.** Fixing NTSC alone would be a point fix; fixing all
three needs one mechanism, because integer multispeed and NTSC are both
special cases of "driven at N Hz, declared per song."

## 2. Non-goals

- Mid-song changes to the *drive rate* itself (as opposed to musical
  tempo within a fixed drive rate, which `tempo`/`beats`/`speed` already
  handle). Real players that change their own interrupt rate mid-tune exist
  but are rare; out of scope for this ADR, flagged as a known limitation the
  same way ADR 0001 flags mid-song tempo changes.
- Sub-frame timing (digi playback, `$D418` sample streams faster than the
  drive rate). Already out of scope for DSL v2's frame-sampling compiler
  model; unaffected by this ADR.
- Automatically detecting the right value. Detecting a tune's play-call clock
  is a job for the tool that analyses the SID file; this ADR gives DSL v2
  somewhere to put the measured number, not a way to infer it inside the DSL
  toolchain.

## 3. Proposed mechanism

One song-level declaration:

```
target { video PAL }              // sugar for driveHz 50.0000238... (985248/19656)
target { video NTSC }              // sugar for driveHz 59.8260... (1022727/17095)
target { video PAL  multispeed 3 } // sugar for driveHz = 3 * PAL video rate
target { driveCycles 12801 }       // exact: this many CPU cycles per play call,
                                    // for a CIA-driven rate with no clean video
                                    // relationship (the Soundmonitor case above)
```

`driveCycles` (integer, CPU cycles per invocation, exact) is the one new
primitive; `video PAL|NTSC` and `multispeed N` are named shorthands that
expand to it at parse time (`video PAL` = `driveCycles 19656` at 985248 Hz;
`multispeed 3` multiplies the currently-declared video's per-frame cycle count
by 3). `driveCycles` is the generic mechanism, `video`/`multispeed` are macros
over it for the two common named cases. Omitting the whole `target` block
keeps today's default (PAL, `driveCycles 19656`): purely additive, no change
to any existing `.dsl` file's meaning. (The PAL cycle count was corrected to
19705 during implementation; see 7.2.)

**Why cycles, not Hz, as the primitive.** A floating-point Hz value
accumulates rounding error over a multi-minute song, while SID playback and
the traces it is compared against are cycle-exact. `driveCycles` keeps the
new primitive in the same units as the thing that has to reproduce it
exactly, and it is the number a play-call clock measurement produces
directly, so a SID-to-DSL converter can populate it mechanically rather than
needing a Hz-rounding step.

## 4. Compiler changes

- **`rasterklang-score/src/compiler/compile.js`**: `framesPerCycle`'s
  derivation (currently one constant, per ADR 0001) reads the song's
  `driveCycles` instead. This is the one load-bearing change; everything
  downstream that already consumes `framesPerCycle` (row/beat timing,
  frames-since-onset tables) needs no further change, since it is already
  expressed in frame units, not wall-clock seconds.
- **`rasterklang-score/src/engine/liveEngine.js`**: computes "how many
  emulator cycles between one register frame and the next" for real-time and
  offline rendering, and needs to read the same `driveCycles` rather than
  assume PAL 50 Hz. The output sample rate (44,100 Hz) is unaffected; only
  the SID-write cadence changes.
- **Offline renderers** that turn compiled register frames into audio need
  the same drive rate passed through.
- **SID-to-DSL converters** have no way to express a non-50 Hz source until
  the grammar supports the declaration above, and every emitted `.dsl`
  silently inherits the 50 Hz assumption. They need to read the source's
  declared clock and video standard plus the measured play-call clock, and
  emit the matching `target { ... }` line. The grammar change alone does
  nothing until an emitter uses it.
- **Grammar (`grammar/RasterklangDsl.g4`) and both parsers** (`pkg/dsl`,
  `src/dsl`): add `driveCycles INT`, `multispeed INT` and `video PAL|NTSC`
  as statements inside `target { }`.

## 5. Validation

1. **`video PAL|NTSC` first**: the cheapest change (one constant swap, no new
   numeric primitive), and it alone addresses the 0.83x-speed failure on NTSC
   tunes. Verify by rendering NTSC tunes before and after and comparing
   against reference renders of the original SID files.
2. **`driveCycles` general form second**, verified against
   `Welle_Erdball_Alles_Luege.sid` (12,801 cycles/call), since it exercises
   the general arithmetic path a PAL/NTSC toggle does not.
3. **`multispeed N` sugar third**, verified against `Danger_Zone.sid` (3x)
   and `Interest.sid` (4x).
4. At every step, a claim of improvement is accompanied by a before/after
   measurement on named files, not a general assertion that the fix "should"
   work.

## 6. Compatibility

Purely additive. No existing `.dsl` file changes behavior: the default
(PAL 50 Hz) is today's only behavior. This follows the same low-risk shape as
earlier DSL v2 additions logged in `dsl-changelog.md` (e.g. 2.18.0's `vibTri`,
2.17.0's tempo `offset`): a targeted grammar addition driven by one measured,
named gap, not a redesign.

---

## 7. Implementation notes and measured outcome (2026-08-02)

The mechanism exists end to end, and all three validation steps of section 5
were run on the named files.

### 7.1 Measured, before and after

"Before" is a converted `.dsl` with its `target` line deleted (the pre-ADR
behaviour); "after" is the same file unmodified. Both were rendered and
compared against a reference render of the original SID, over 600 frames.

- **`multispeed` (step 3), the clean result.** `Danger_Zone.sid` (3x) and
  `Interest.sid` (4x) went from playing at 0.33x and 0.25x speed to 1.00x,
  and from clearly divergent to spectrally equivalent to the reference.
- **`driveCycles` (step 2).** `Welle_Erdball_Alles_Luege.sid` (measured
  12,787 cycles/call against the CIA latch's nominal 12,801) went from 0.65x
  speed to 1.00x, with rendered length matching the reference within 2%.
- **`video PAL|NTSC` (step 1), the partial result.** All tested NTSC tunes
  now render within 1.3% of their reference's length and at 1.00x speed.
  Spectral similarity improved on half of them and stayed flat on the rest.
  On those, the dominant remaining loss is intra-frame register collision
  (two or more values written to one register inside one play call), not
  tempo. That is this ADR's own declared non-goal (section 2).
- **Compatibility, checked rather than asserted.** PAL tunes emitting
  `target { video PAL }` render byte-for-byte identical audio with and without
  the line.

### 7.2 Deviations from this document, and why

1. **PAL is 19705 cycles, not 19656.** Section 3 says `video PAL` =
   `driveCycles 19656` while also calling it 50.0000 Hz; those two are
   inconsistent (985248/19656 = 50.1245). The rest of the toolchain defines a
   video frame as `machineClock/50` for PAL and `machineClock/60` for NTSC,
   and section 6 requires the no-target default to stay today's behaviour. So
   PAL is 19705 (49.99990 Hz, 2e-5 from an exact 50, and still exactly 882
   samples per frame at 44.1 kHz) and NTSC is 17045 at 1022727 Hz
   (60.00158 Hz). Documented in `src/dsl/driveRate.js`.
2. **Converters may prefer `driveCycles` over `video NTSC` / `multispeed N`.**
   Both macros exist, parse, and are covered by tests, but a converter that
   has a measured play-call period can write that value directly; the
   nominal macro values can leave a small residual against a measured
   reference. `video PAL` remains the natural spelling for single-speed PAL
   tunes, where it is numerically free.
3. **`wavExport.js` needed no change.** It is a pure RIFF/PCM header encoder
   with no frame-rate concept. The frame-cadence assumptions were in
   `liveEngine.js` (`samplesPerFrameFor`, `renderOffline`, `createLiveEngine`,
   `setFrameSource`) and in the editor's call sites, which were changed
   instead.
4. **Syntax.** `target { ... }` is spelled with braces exactly as this
   document writes it, which required two new lexer tokens (`{`, `}`) and a
   new soft keyword (`target`, added to `anyIdent` like the other statement
   keywords, so it still works as an ordinary identifier). `driveCycles` and
   `multispeed` together is a semantic error rather than a silently-resolved
   combination, and at most one `target` block per song is enforced (a
   mid-song drive-rate change is section 2's non-goal).
