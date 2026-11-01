# ADR 0005: Player idioms for converting real SID tunes (lessons from unitracker)

Status: implemented (4.3.0)
Date: 2026-11-01
Affects: `rasterklang-dsl` (fx lane, `hardRestart`, the tempo statement, JS and Go AST visitors),
`rasterklang-score` (compiler), and tools that emit DSL (the GoatTracker importer and the
behavioral decompiler in `rasterklang-conv`)
Current DSL version: 4.2.0. This ADR targets 4.3.0.

**Compatibility statement.** Every decision here is additive. Each new spelling was a compile error
before 4.3.0 (an unknown `fx` op, extra `hardRestart` tokens, an unknown tempo setting), so no song
that compiled under 4.2.0 can change meaning. This was verified by rendering 785 existing songs
(the conformance fixtures, the example songs of `rasterklang-compose`/`rasterklang-daw`, and a
one-in-six sample of the converted tunes in `rasterklang-conv/out`) for 1500 frames each before and
after the change and comparing the register poke streams byte for byte: all identical, and the 30
intentional error fixtures fail with identical messages.

---

## 1. Context

[unitracker-grammar](https://github.com/anarkiwi/unitracker-grammar) is a canonical, player-agnostic
document format for C64 SID tracker music, and
[unitracker](https://github.com/anarkiwi/unitracker) lifts real tunes into it from three native
players (GoatTracker 2, defMON, Music Assembler), verifying every lift against the reference
player's register grid over HVSC. Its constructs are therefore a list of what real players
actually need, found the hard way.

It is not a model for this language. It is deliberately machine-dense (one canonical text, no
line wraps), it embeds 6502 code templates, and it models player internals (`regs`, `lag`, tick
phase `order`, sub-frame `calls`) that a readable composition notation should not expose. But
comparing its construct list with the open gaps of our own converter
(`rasterklang-conv/docs/dsl-feedback.md`) and with the "deliberately unsupported" rows of the
GoatTracker command triage in `docs/dsl-grammar-spec.md` shows five places where a real-tune idiom
has no spelling here and a small additive one would do.

## 2. What unitracker has that bears on converting tunes

| unitracker construct | what it expresses | our gap |
| --- | --- | --- |
| named programs, started from a row (`run:c7`, also inline) | jump a voice's wave/pulse/filter program to another one mid-note | GoatTracker `8xx`/`9xx`/`Axx` are unsupported: instrument tables are private, there is no global index |
| `res`, `route`, `route_on`/`route_off`, `mode` as row commands | resonance, routing and filter mode change per row | GoatTracker `Bxx` unsupported; feedback F-4 (filter mode has no per-row surface) |
| modulator bindings (`pw:'v + 257*s'`, prelude `pwm`, `sweep`, `slide`) | a register ramped per tick from a row, with wrap or bounce | feedback F-8: register ramps (`m5-lane-lift`, 53% pulse) have no construct |
| `restart:hr(lead,ad,sr,test,stop)` | hard restart: lead frames before a note with their own AD/SR and test bit | feedback F-7: our `hardRestart N` only drops the gate, the AD/SR blip is dropped |
| tempo set on a row (Music Assembler `$FC`), funktempo via alternating speeds | non-uniform row lengths | GoatTracker `Exx` (funktempo) unsupported |

Also noted, not adopted now (section 9): layers (several programs of one kind per voice, needed by
defMON), per-voice orderlists with independent loops (our `<...>` alternation and polymetre
already cover the musical cases), `calls` (multispeed where only some calls read rows, possibly
relevant to feedback F-6), the note-table generator survey (`tools/notetables.py`), and the
canonical-print/round-trip discipline.

## 3. Decision D1: start a named table from a row (`wtab`/`atab`/`ptab`/`ftab`)

New `fx` tokens, one per table kind that GoatTracker can repoint:

| token | replaces this voice's | GoatTracker |
| --- | --- | --- |
| `wtab.NAME` | `wave` table | `8xx` (wave table pointer; the waveform column) |
| `atab.NAME` | `arp` table | `8xx` (the note column of the same table) |
| `ptab.NAME` | `pw` table | `9xx` |
| `ftab.NAME` | `filterTable` | `Axx` |

`NAME` is a `table NAME = ...` declaration. On the frame the fx hap becomes active the named
table starts at index 0 and runs on its own clock (frames since that fx fired), with the same
loop, ramp and `!N` semantics as an instrument table. An optional loop index overrides the
table's own: `ptab.sweep.4` wraps to index 4. Without one, a ramp's implied loop point applies,
else 0, exactly as for an instrument table that references the same declaration.

The override lasts until the next note ONSET, which re-initialises the instrument's own tables,
as GoatTracker's note init resets the pointers. A legato transition (`tp`, `gate legato`,
`slideTo`) is not an onset and keeps the override. It applies whether or not the instrument has a
table of that kind. An `ftab` takes part in the "most recent filter table wins" rule with the
frame it fired as its start.

Why per-kind tokens rather than one `run.NAME`: a `table` declaration is untyped (a row of
numbers or waveform words), so the kind has to come from somewhere, and the four tokens map
one-to-one onto the commands an importer meets. Why not `$NAME`: inside a mini-notation string
`$NAME` already means a pattern splice. The name is checked when the voice is built: an unknown
table, or a word cell in a `ptab`/`ftab` table, is an error.

This removes the reason the triage gave for leaving `8xx`/`9xx`/`Axx` unsupported. The importer
can emit each reachable jump target as a named table instead of needing a global index space,
and instrument identity is untouched.

## 4. Decision D2: filter control per row (`res`, `route`, `fmode`, `fctl`)

| token | sets | range |
| --- | --- | --- |
| `res.N` | resonance (`$D417` high nibble) | 0-15 |
| `route.M` | routing mask (`$D417` low nibble; bit 0 = voice 0, bit 3 = external input) | 0-15 |
| `fmode.N` | mode bits (`$D418` high nibble: 1 lowpass, 2 bandpass, 4 highpass, 8 voice 3 off) | 0-15 |
| `fctl.R.M` | resonance and routing at once, GoatTracker `Bxx` exactly | 0-15 each |

They behave like the existing `fcut.N`/`vol.N`: song-global (the chip has one filter), set from
whichever voice's row carries them, and from that frame on they override the `filter` block's
own value for that field. Out-of-range values are an error rather than a clamp, for the same
reason as `pw.N`: an importer mixing up a nibble with a byte should fail loudly. `fmode` takes the
raw nibble (not names) because that is what tracker data carries and because a single row must
be able to combine modes.

This closes feedback F-4 and the `Bxx` triage row.

## 5. Decision D3: register sweeps from a row (`pws`, `pwm`, `fsw`)

| token | per frame, while the fx step is active |
| --- | --- |
| `pws.D` | pulse width += D (signed), wrapping modulo 4096 like the 12-bit register |
| `pwm.LO.HI.S` | pulse width moves by S toward HI, bounces between LO and HI (ping-pong) |
| `fsw.D` | filter cutoff += D in GoatTracker's 8-bit cutoff units (x8, as `fcut`), wrapping modulo 256 |

The pitch commands `pup`/`pdn` already work this way for frequency; these are the same shape for
the two other registers real players sweep most. Each starts from the value the register holds
this frame (an earlier `pw.N`/`fcut.N`, else the instrument or `filter` block value) and leaves its
result in place when the step ends, as a sticky `pw.N`/`fcut.N` would.

The wraps are deliberate and follow the players: GoatTracker's pulse arithmetic is 12-bit, and
its filter accumulator is 8-bit and wraps (ADR 0004 documents that sawtooth and why `ramp`
saturates instead). `fsw` is the faithful spelling of the wrapping filter sweep that ADR 0004 said
`ramp` cannot be. `pwm` is the bounded alternative most hand-written PWM wants.

Precedence matches `pw.N`: a running `pw` table (the instrument's or a `ptab`) outranks a pulse
sweep, and a running filter table outranks `fsw`. The pulse result is cleared at the next note
onset like `pw.N`, except that a `pws`/`pwm` step that began before the onset and is still active
keeps sweeping through it, so a sweep written across several notes (`fx "pws.8@4"`) is one
continuous ramp.

This is the narrow, readable part of unitracker's modulator idea. Arbitrary expressions and code
templates (section 9) stay out.

## 6. Decision D4: `hardRestart N` with its own envelope and test bit

```
hardRestart 2 ad 0 0 sr 0 0 test
```

The existing form `hardRestart N` is unchanged (exactly `gate -N`). The optional trailing clauses
apply during those N gate-off frames, and only when another note follows directly (a rest after
the note releases normally, as GoatTracker only hard-restarts before a new note):

- `ad A D` / `sr S R`: write this attack/decay and sustain/release (nibbles 0-15) instead of the
  instrument's. This is the classic AD/SR reset that makes the next attack start clean.
- `test`: set the test bit for those frames (unitracker's `hr(...,test=1)`).

Clauses may come in any order, each at most once. This gives feedback F-7 a spelling: the
behavioral decompiler can emit the reset instead of suppressing it. The clauses are carried on the
same gate AST node as an optional `hardRestart` object, present only when a clause is given, so
every existing AST (and conformance golden) is unchanged.

## 7. Decision D5: funktempo (`groove`)

```
speed 6 groove 5 beats 8
```

With `speed N`, `groove M` makes beats alternate between N and M frames, counted from the start of
the song: beat 0 lasts N, beat 1 lasts M, beat 2 N, and so on. GoatTracker's funktempo alternates
row lengths the same way, and an importer emits one beat per row. A bar of an odd number of beats
therefore alternates between two lengths, which is also what the original does. Notes keep their
position in the bar; only the frame mapping inside the bar becomes piecewise linear.

`groove` is only valid with a constant `speed` (frame counts), not with `tempo` (BPM) and not with
a `speed "<...>"` pattern, each a named error. The frame-grid diagnostics of ADR 0003 D17 are
skipped for a grooved song (they assume uniform beats); an info line reports both beat lengths.

Mid-bar tempo changes on arbitrary rows (Music Assembler `$FC`, GoatTracker `Fxx` inside a
pattern) remain out of scope: the song clock is derived from bar position and an `fx` lane is
sampled by bar position, so a tempo command inside the lane would make the clock depend on
itself. Tempo changes at bar boundaries are already expressible as `speed "<6 6 5>"`.

## 8. Implementation notes

- No grammar change. `fx` tokens live inside a string, `hardRestart`'s clauses fit the existing
  `tableRow`, and `groove N` fits the existing `tempoSetting` shape. No parser regeneration.
- `src/dsl/fx.js` gains the ops, a per-op arity range, range checks and a `{ table }` parameter
  for the table ops; `interpret.js` resolves table names against the song's `table` declarations.
- `astVisitor.js` and `pkg/dsl/ast_visitor.go` parse the `hardRestart` clauses and `groove`; the
  conformance canonicaliser emits the new keys only when set.
- `rasterklang-score`'s `compile.js` implements the runtime: per-voice table overrides, the global
  filter overrides, the sweeps, the hard-restart window and the groove clock.

## 9. Not adopted, and why

- **Expression and 6502 code modulators.** Powerful and exact, but they turn a composition
  language into a player-construction kit. D3 covers the ramps the converter actually measured.
- **Layers** (several programs of one kind per voice). Needed for defMON, not for GoatTracker; no
  measured demand here yet.
- **`regs`, `lag`, tick-phase `order`, sub-frame `calls`.** Exact player emulation. Feedback F-2
  already declares sub-frame writes a boundary. `calls` is worth revisiting for F-6.
- **A canonical printed form with a round-trip test.** Useful for corpora (stable diffs, normalised
  training data) and worth its own ADR; independent of the five idioms here.

## 10. Follow-ups

- The GoatTracker importer (`rasterklang-conv`'s `sngcodegen`) can now emit `8xx`/`9xx`/`Axx` as
  table ops, `Bxx` as `fctl`, and funktempo as `groove`; until it does, the triage rows describe
  what the importer drops, not what the language lacks.
- The behavioral decompiler can emit `pws`/`pwm`/`fsw` for register ramps (F-8) and the
  `hardRestart` clauses (F-7); `dsl-feedback.md` should record the new status when it does.
- unitracker's `tools/hardrestart.py` classifies hard-restart idioms over HVSC and would show
  whether `ad`/`sr`/`test` cover the common variants.
