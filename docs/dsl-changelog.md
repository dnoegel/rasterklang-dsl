# DSL v2 grammar changelog

Tracks changes to the DSL's actual grammar/semantics (`src/dsl/{lexer,parser,interpret}.js`,
and the `VoiceChain`/`FilterChain` surface of `song.js`, now in `rasterklang-score`) - not any
package version (`package.json`), and not editor-tooling-only changes (syntax highlighting,
autocomplete) unless they reflect a real grammar change. See `src/dsl/version.js`'s
`DSL_VERSION` for the current value, and `docs/dsl-grammar-spec.md` for the full current spec
this history led to.

Entries below `2.7.0` were reconstructed after the fact (this changelog and `DSL_VERSION` did not
exist yet when those changes landed). Everything from `2.7.0` on is added in the same change that
bumps the version.

Several entries mention a GoatTracker-to-DSL converter. The converter is a separate tool; it is
mentioned here only where its needs motivated a grammar change.

## 4.2.0

**New `fx` token `pw.N`: a per-row, one-shot pulse-width overwrite.** The `fx` lane already had a
command for every other SID register class (`pup.N`/`pdn.N`/`tp.N` for frequency, `wave.N` for the
control nibble, `ad.A.D`/`sr.S.R` for the envelope, `fcut.N`/`vol.N` for the global lanes). Pulse
width had none, so the only way to move it was an instrument `pw` TABLE, which is a per-patch,
per-frame program rather than a per-row event.

Why now: measuring which register writes a SID-to-DSL converter could not express, over a sample
of tunes from the High Voltage SID Collection, pulse width was the largest single class of
unexpressible events, and it was the only register class with no `fx` spelling.

Semantics mirror `wave.N` exactly: parsed into the same `{ op, a }` object, installed as a sticky
per-voice override on the frame the `fx` hap becomes active, cleared at the next note ONSET, and
ignored when the instrument has a `pw` table of its own (that table rewrites the register every
frame anyway, the same reason `wave.N` yields to a `wave` table). Unlike GoatTracker's other
commands, `pw.N` has no `CMD_*` counterpart: `CMD_SETPULSEPTR` (`0x09`) repoints a table and stays
deliberately unsupported.

The parameter is the SID's 12-bit pulse width, `0-4095` - deliberately the same value domain as the
instrument `pw` column, so the two spellings cannot disagree. A value outside that range is a parse
ERROR rather than a clamp, because the caller most likely to produce one is an importer confusing an
8-bit tracker pulse value with the 12-bit register one, and a silent clamp would hide that.

Compatibility: purely additive. No existing token, and no song not containing `pw.N`, changes
behavior.

## 4.1.0

**A shaped `ramp`'s implied loop point now survives a shared `table NAME = ...` declaration.**
`ramp A..B step D pingpong` and `ramp ... loop` (ADR 0003 D9) set an implied per-table loop index
of 0. Written inline in an instrument field that already worked; written inside a `table`
declaration and referenced as `$NAME`, the values were spliced in but the loop point was silently
dropped, so the table fell back to whatever the instrument's own `loop TABLE N` field said.
`interpret.js`'s `expandTableRefs` documented this as "a known, narrow gap" with the workaround
"write the ramp inline in the instrument field itself", which defeats the point of sharing.

`tables` now stores `{ values, loopIndex }` per name rather than a bare array, and a `$NAME`
reference folds the referenced table's own loop index into the referencing row through the SAME
last-one-wins rule an inline shaped ramp already used. No second precedence rule was introduced.

Why now: ADR 0004 establishes that a wrapping GoatTracker-style filter sawtooth is expressible in
this DSL precisely as a short ramp carrying `loop`, and that idiom is exactly what this gap blocked
inside a shared table. Tables a converter hoists into a shared declaration are shared by
construction, so the gap applied to the very songs most likely to need it.

Compatibility: this changes compiled output only for a song that puts a SHAPED ramp inside a shared
table, which is the case that was broken. A shared `ramp 1088..0 step -48 over 128` with no shape
terminator produces a byte-identical register stream (SHA256 over 2000 frames) before and after.

## Documentation correction (2026-08-16, no version bump)

`docs/adr/0004-ramp-saturation-vs-goattracker-filter-wrap.md` corrects a false claim in
`docs/dsl-grammar-spec.md`'s "Writing long tables: `ramp` and `!`" section and in this file's own
`2.12.0` entry below: both stated that `ramp`'s saturating clamp is exactly what GoatTracker's
pulse/filter modulation-step command does, and that one `ramp` reproduces one original tracker row
exactly. It is not and it does not - GoatTracker's own `filtercutoff` accumulator is an unclamped
8-bit value that wraps, as GoatTracker's player source (`gplay.c`) and a register capture of a
real tune show. `ramp`'s actual behavior (saturating at the target) is
**unchanged** by this entry; only the documentation describing what that behavior represents is
corrected, in place, at both sites. No grammar, parser, or interpreter change; `DSL_VERSION` stays
`4.0.0`.

## 4.0.0

Musical durations and song structure. Implements
`docs/adr/0003-musical-durations-and-song-structure.md`; that document carries the full
rationale and the measured evidence, decision by decision. This entry is grouped by decision.

**BREAKING: `/N` in mini-notation now means a DURATION in beats (D1), not "slow down by N".**
`c4/4` is a quarter note (`beatUnit/4` beats), `c4/8.` a dotted eighth, `c4/8t` a triplet
eighth. A pattern mixing `/N` durations with `@n` weights is now an error - they are two
authoring models for the same underlying number. Any existing song using `/N` for its old
"slow this step down" meaning will now parse differently; that old meaning had no real users
this DSL is aware of and a duration is what a musician actually means by the syntax.

**New: `meter N/D` (D2) alongside `tempo`/`speed`, a real time signature.** Replaces `beats
INT` (which only ever gave the numerator, always assuming quarter-note beats); `beats` stays
legal but is mutually exclusive with `meter` on one clock statement.

**New: `pattern NAME each 1/D = ...` (D3), a step-sequencer declaration.** Every token is
exactly one `1/D` note; the pattern's length is DERIVED from its own token count instead of
declared - the natural spelling for a fixed-grid drum row.

**New: arrangements (D4) - `pattern verse = $a $b`, `[$a $b]*2`.** A structural sequencing
primitive distinct from `<...>` alternation: `$a $b` plays `$a` then `$b` in ONE pass (an
orderlist), where `<$a $b>` plays one OR the other, alternating by pass. Backed by a new
`timeline` primitive in the pattern engine (`src/pattern/timeline.js`, now in
`rasterklang-score`).

**New: `$name` is a first-class structural pattern reference (D5), not text substitution.**
Every Hap a `$name` reference contributes now carries its ABSOLUTE document position by
construction, fixing a whole class of source-highlighting bugs a text-splice implementation
could never close (including the removal of a known `$name!n` splice-hazard warning, which is
now structurally impossible rather than merely detected).

**New: checked bar lines `|` in mini-notation (D6)**, verifying a pattern's declared beats
land exactly on the boundary the author drew, named with the offending beat position.

**New: proportional-pattern auto-fit is reported (D8)**, an info-level diagnostic rather than
silent behaviour.

**New: `pw ramp 400..800 step 20 pingpong` / `... loop` (D9)**, sweep SHAPE terminators. Without
one, a ramp still saturates at its target exactly as before.

**New: `speed N` / `delay N` on every instrument table (D10)** - wave/arp/pw/filterTable/test/
gate - alongside the existing `loop`, resolved independently per table.

**New: `gate legato` / `gate retrigger` / `gate N%` / `hardRestart N` (D11).** `legato`
suppresses retriggering between consecutive notes entirely (the gate never drops,
frames-since-onset keeps counting); `retrigger` guarantees a real, frame-observable
low-to-high edge on every note including a repeated pitch (today's forced pair alone is
invisible to anything sampling registers once per frame); `N%` is percentage sugar for the
existing fraction; `hardRestart N` is the readable spelling of `gate -N`.

**New: `slideTo:NOTE` (D12)**, portamento to an explicit target pitch, legato with the
previous note, using the instrument's own `slide TIME`.

**New: table verbs (D13)** - `set V` / `add V [repeat N]` / `wait N` / `loop` / `stop`, an
`end`-terminated block form on any table-shaped field, for tables too long to read as a row of
bare numbers. Compiles to exactly the same per-frame array the bare-row form produces.

**New: `chip 6581` / `chip 8580` (D15)**, song-level SID model declaration, carried on the
compiled song as `chipModel` and written to the `.sid` exporter's header flags, taking
precedence over an editor's current model selection.

**Deferred to a future version, recorded not dropped (D16):** `vib ... delay N` (sine
vibrato's own delay parameter, unlike `vibTri`'s); instrument sequence `continue` mode (a
free-running table that ignores note boundaries); combined wave+pitch table rows; relative
pitch slide `slide +12 speed 2`; `jump` as a table verb. See the ADR's own D16 section for the
reasoning behind each deferral, including two explicitly revisited after D10 landed.

Every cross-language-lockstep change above (grammar/AST/beat-arithmetic) landed in JS
(`src/dsl/`) and Go (`pkg/dsl/`) in the same commit, with regenerated parsers and conformance
goldens - the two implementations stay verified-identical, not just believed identical.

## 3.1.0

Native play rate. Implements `docs/adr/0002-native-play-rate.md`; that document carries the
measured evidence (NTSC tunes re-synthesised at about 0.83x speed, `Danger_Zone.sid`'s 3x
multispeed played at 0.33x, `Welle_Erdball_Alles_Luege.sid`'s 12,801-cycle CIA clock) and the
rationale.

**New: an optional song-level `target { ... }` block declaring the DRIVE rate.** Where
`tempo`/`speed`/`beats` say how long a beat is in frames, this says how long a frame is in CPU
cycles - the question DSL v2 previously had exactly one hard-coded answer to (PAL, 50.00 play
calls per second).

```
target { video PAL }               ; 985248/19705  = 49.99990 Hz - the default, unchanged
target { video NTSC }              ; 1022727/17045 = 60.00158 Hz
target { video PAL multispeed 3 }  ; 3 play calls per video frame (GoatTracker funktempo)
target { driveCycles 12801 }       ; exact: CPU cycles per play call, for a CIA-driven rate
```

`driveCycles` (a positive integer count of CPU cycles) is the one new primitive; `video
PAL|NTSC` and `multispeed N` are named macros over it. `driveCycles` and `multispeed` together
are an error (two spellings of one number). At most one `target` block per song; a mid-song
drive-rate change is out of scope (ADR 0002 section 2).

Purely additive: with no `target` block a song means `driveCycles 19705` at the PAL clock,
i.e. 50.00 fps, which is DSL v2's only pre-3.1.0 behaviour, so no existing `.dsl` file
changes meaning. The declared rate is what `speed N`'s frames-per-beat and `offset N`'s
beats-to-frames conversions divide by (both were hard-coded `50 * 60` before), and it rides on
the compiled song as `driveCycles`/`cpuClockHz` for the compiler and its renderers.

Cycle counts, not Hz, because SID playback is cycle-exact and because `driveCycles` is the number
a play-call clock measurement produces directly, so a converter can populate it mechanically. Note
PAL is 19705 (`machineClock / 50`), not the hardware-true 19656-cycle raster period; see
`src/dsl/driveRate.js`'s header for why.

## 3.0.0

Metric patterns, bars instead of cycles, and a frame-native clock. Implements
`docs/adr/0001-metric-patterns-and-song-time.md`; that document carries the measured evidence
and the rationale, this entry carries what actually changed.

**BREAKING: `beats N` is now a pure time signature.** It declares how long a BAR is and nothing
else. It no longer determines how long a pattern is. Any song that used a large `beats` value
as a melody-stretch lever (`tempo 132 beats 28` was a real case) will now sound different, and
`npm run lint:dsl -- song.dsl` flags exactly that shape with a hint pointing at declared
pattern lengths instead. Songs that used `beats` as an honest time signature are unaffected.

**BREAKING (vocabulary): "cycle" is now "bar".** A bar is `beats` beats. `every 4 rev` means
"every four bars". Bar counting stays global rather than per-pattern, so two voices with
`every 4` still fire together. Both user-visible spellings kept an alias so no existing song
stops compiling: `switch bar` is the new spelling of `switch cycle`, and `someBarsBy` of
`someCyclesBy`; the old words still parse and mean exactly what they always did.

**New: a pattern may declare its length in beats, which makes it metric.**

```
pattern hats    = "bd*4"          ; proportional (unchanged): fills exactly one bar
pattern bass 4  = "c2 ~ c2 ~"     ; metric: 4 beats, one token = one beat
pattern odd 7.5 = "c4@0.5 d4@3 e4@4"
```

In a metric pattern an unweighted token is one beat and `@n` is n beats; the contents are
checked against the declared length and a mismatch is a compile error naming both numbers; the
pattern loops at its own length rather than at the bar length. Patterns with no declared length
keep today's proportional semantics exactly, so this is opt-in and no existing song changes
behaviour because of it. Inline strings on a voice stay proportional (they are anonymous
patterns).

**New: `speed N` as an exact, frame-native clock.** `speed 6 beats 8` is 6 frames per beat -
48 frames per bar, no rounding - where `tempo 132 beats 8` computes 181.8 and has to round.
`bpm = 3000 / framesPerBeat` exactly. `tempo` and `speed` are mutually exclusive and a song
needs exactly one of them: having neither or both is now an error that says which case applies.

**New: `_`, the hold token.** `"c4 _ _ _"` is one note ringing for four beats. Unlike `~`, a
hold does not cut the note - trackers have trigger/hold/release and mini-notation only had
trigger and release. `_@2`/`_!2` add that many steps at once; a hold with nothing before it is
an error; a hold after a rest extends the rest. This is additive (`_` used to be a parse
error), and it is what makes it practical to split a long melody into per-bar patterns, since a
note tied across a bar line finally has a spelling.

**Polymetre across voices is explicitly legal** and now reports the arrangement's true period
(`Full period: 12 beats (3 bars)`) instead of being silently confusing, with a soft hint when
that period is implausibly long. What is an error is a declared length disagreeing with its own
pattern's contents, or one `<...>` orderlist mixing elements of different declared lengths.

**Diagnostics.** Compilation now produces warnings and info alongside errors: they are passed
to `compileDslV2Source(source, { onDiagnostic })` and attached to the returned song as
`.dslDiagnostics`. `checkSource(text)` (and the `npm run lint:dsl` CLI over it) returns the
same list without compiling, which is what the migration checker runs on.

Fixed alongside (all four defects listed in the ADR's section 6):

1. **Weighted rests close the gate.** A rest closes the gate for its own span whether or not it
   carries a weight, so `"c4@4 ~@4"` releases halfway instead of ringing through. The rest-span
   gate behaviour itself lives in the pattern engine (now `rasterklang-score`); on this side the
   rule is specified, and the hold token is expressed as weights so that "ring" and "release"
   are two different spellings rather than two readings of one.
2. **`$name!n` warns.** `<$bassA!3 $bassB>` applies the `!3` to the pattern's last token,
   because `$name` splices raw text. The behaviour is unchanged (changing it would silently
   re-time existing songs); the compiler now says so and suggests `[$bassA]!3`.
3. **`|>` on a pattern declaration** produces "pipeline transforms ('|>') are not allowed on
   pattern declarations; add the transform on the voice that plays this pattern instead"
   instead of ANTLR's `mismatched input '' expecting NEWLINE`.
4. **Cutoff above 2047 warns** ("cutoff 2600 clipped to 2047 (11-bit filter cutoff)") instead
   of clipping in silence. The clipping itself is unchanged.

## 2.19.0

**BREAKING: an instrument's `filterTable` now WRITES the shared filter's cutoff, absolutely and
regardless of `route`.** It used to be a per-voice cutoff OFFSET, summed across whichever voices
were routed into the filter that frame and added on top of the `filter` block's own `cutoff`.
Both halves of that are gone:

- **Absolute, not an offset.** While a table runs, its value IS the cutoff - it replaces the
  `filter` block's `cutoff` and any `fx fcut.N`.
- **Independent of `route`.** A SID has exactly one filter with one cutoff register. `$D417`'s
  low bits (`route`) choose which voices pass THROUGH that filter; they never decide what its
  cutoff VALUE is. Gating the table on routing meant a filter sweep on an unrouted voice was
  silently discarded.
- **Last writer wins, not summed.** When two voices run a table at once, the most recently
  started one owns the cutoff (ties go to the higher voice index), the way GoatTracker's single
  global `filterptr` behaves - `gplay.c` has ONE `filterptr`/`filtercutoff` pair, re-pointed by
  whichever note most recently started an instrument that carries a filter table, and it writes
  `sidreg[0x16]` outright.

Why it changed: converting the GoatTracker tune 10_Orbyte (from the High Voltage SID Collection)
produced a completely dead filter. The original sweeps its cutoff on 592 of 600 frames; the
conversion changed it 0 times. The song routes voice 0, and the instrument carrying the filter
sweep plays on voice 1, so the whole sweep was thrown away by the routing gate. (The converter had
a separate bug of its own in decoding the filter table.) The new rule also removes a silent
failure mode for hand-written songs, where a filter sweep does nothing because the author did not
think to route a voice.

No grammar change: `filterTable` is still `instrumentField := anyIdent tableRow`, and the
`filter` block is untouched. Only the meaning of the numbers changed, which is exactly why this
is a version bump rather than a patch. To restore roughly the old sound in an existing song,
fold the `filter` block's `cutoff` into the instrument table's own values.

## 2.18.0

**`vibTri CMP SPEED DELAY` on an instrument - GoatTracker's own vibrato, exactly.** A second,
separately-named vibrato field rather than a mode on `vib`, because the two are different shapes
in different units and neither can express the other. `vib DEPTH RATE` is a sine of DEPTH
semitones at RATE Hz. `vibTri` is `gplay.c`'s vibrato: a `vibtime` counter stepping by 2 and
inverting whenever it passes CMP, whose low bit decides whether this frame ADDS or SUBTRACTS
SPEED from the running frequency register - an integrating triangle in raw register units, whose
first quarter period is only half the swing of every later one, started DELAY frames after the
note's onset. One half period is `CMP+2` frames; the swing is `+/-(CMP/2+1)*SPEED` register
units. SPEED may instead be `sN` ("one semitone at the current note, shifted right by N"), the
same note-relative encoding the `fx` lane's parameters already accept, since GoatTracker's speed
table stores it that way and it cannot be resolved to a constant at conversion time.

No grammar change was needed: `instrumentField` is `anyIdent tableRow`, so this is one
AST-transform-visitor case in each language.

Why it exists: the `fx` lane's `vib.CMP.SPEED` (2.15.0) already had the exact per-ROW command,
but GoatTracker also carries vibrato on the INSTRUMENT (a speed-table pointer plus a delay byte),
where it runs continuously for every note that instrument plays and only while the row carries no
command of its own. That is an instrument-level continuous modulator, which is what this field
is. Converted GoatTracker songs used to emit no vibrato at all for those instruments.

## 2.17.0

**`offset N` on the `tempo` statement - a song-level pattern-clock shift, in beats.** `tempo 125
beats 16 offset 1` plays the whole song one beat later: real frame N is played as compiler frame
`N - offset*framesPerBeat` (`song.js`'s already-existing `startOffsetFrames`, which until now had
no DSL spelling at all and could only be set from JS). Fractional and negative values are both
legal and meaningful. The grammar change is `tempoStmt: TEMPO tempoValue (anyIdent NUMBER)*` -
the trailing settings became a repeated name/number pair instead of the single optional `beats
N`, with the AST-transform visitor validating the names, so a third setting later costs a visitor
case rather than a grammar change. `beats`/`offset` may appear in either order, and omitting
`offset` is exactly the old behaviour.

Why it exists: the GoatTracker importer has to say "this tune's first row sounds one row after
the play routine starts" (`gplay.c` seeds its tick counter from the compiled default tempo before
the sequencer ever fetches a row). It used to say it by ROTATING each voice's row stream by one
row, which put every GoatTracker pattern boundary off the DSL cycle grid - so the generated
`<...>` alternation no longer corresponded to the song's real patterns, and the declared
`pattern` names stopped meaning anything. Declaring the shift once, at song level, gets
pattern-aligned cycles back without reintroducing the phase error. By hand it is simply a pickup
beat.

## 2.16.0

**`gate -N` - gate-off frames at a note's end (per-note retrigger).** A negative `gate` value is
a second spelling of the existing static fraction, in absolute frames instead of a proportion:
`gate -2` holds the gate high until the note's last 2 frames, whatever that note's length is.
This is GoatTracker's per-instrument `gatetimer`, its hard restart (`gplay.c` fetches the next
note `gatetimer & 0x3f` ticks early and keys the gate off there), which is a fixed tick count -
so a `FRACTION` can only reproduce it for one specific note length, and cuts a quarter off a
whole-bar held note if used anyway. Purely additive: no grammar change was needed (`NUMBER`
already lexes a leading `-`), and a non-negative `gate` behaves exactly as before.

The DSL's default articulation is unchanged and stays legato - see the grammar spec's
"Articulation defaults" note. What this fixes is that SID-to-DSL converters had no way to
say "retrigger per note" at all: `compile.js` forces a gate-off/gate-on pair on a note onset, but
both writes land in the SAME frame, so anything sampling the registers once per frame (a `.sid`
dump, a conversion scorer, a player's frame loop) saw a whole converted tune as one sustained
note.

Also in this change (compiler, not grammar): a voice whose instrument switches per note via
`insts` now takes its `gate` from the instrument actually sounding, instead of only from the
voice's single static instrument - the same "the instrument owns it" precedence `wave`/`pw`/
`test` already had.

## 2.15.0

Closes the two biggest remaining expressiveness gaps against real GoatTracker songs - both of
which the GoatTracker importer had been reporting in its own conversion warnings, as content that
was being silently dropped.

**`fx` - per-row effect commands (new voice field).** A parallel mini-notation string of one-shot
effect tokens, structurally the same idea as `insts`: `voice 1 lead "c4 e4 g4" fx:"~ tp.24
vib.16.100"`. Tokens are `pup.N`/`pdn.N` (portamento), `tp.N` (tone portamento, including its
legato gate-retrigger suppression), `vib.CMP.SPEED`, `ad.A.D`/`sr.S.R`, `wave.N`, `fcut.N`,
`vol.N`; a speed param may be `sN` for GoatTracker's note-relative "one semitone shifted right by
N" form. Unlike `insts`, `fx` is sampled on its own step boundaries rather than at note onsets,
because a tracker command runs on its row whether or not that row carries a note. Semantics are
transcribed from GoatTracker 2's actual player (`gplay.c`), and the full command-by-command triage
- including which commands are deliberately left unsupported and why - is in
`docs/dsl-grammar-spec.md`'s "Per-row effect commands" section.

Note that this is a genuinely different thing from the existing instrument-level `vib`/`slide`,
which are continuous per-patch modulation; per-row commands are one-shot triggers with their own
reset semantics, so they share the frequency-register machinery but not the configuration model.

**`=MIDI` - absolute-pitch table cells.** An `arp` row cell may now be written `=36`, pinning that
frame to a fixed MIDI note instead of an offset from the sounding note: `arp 0 =36 0 12`. This is
GoatTracker's wave-table note column $80-$DF, the fixed-pitch drum technique, which a
relative-only `arp` could not express. Grammar-level addition (`absNote` in
`grammar/RasterklangDsl.g4`), so both the JS and Go parsers carry it.

Both features are inert for any song that does not use them: the whole `fx` runtime is gated on
the chain having an effect pattern, and a table with no `=` cell parses and samples exactly as
before.

## 2.14.0

A pass over what Strudel has that this DSL didn't. Two of the motivating examples already worked,
because the "every control is a pattern" idea was already in place for `insts`/`wave`/`sync`/
`ring`/tempo: "play it twice then swap instrument" is just `insts:"<bass bass lead>"`, and "repeat
this" is `!n` or `<...>`. What was missing was that principle applied to *transform arguments*,
plus a set of combinators that were never ported.

- **Patternable transform arguments.** Every numeric transform argument now also accepts a quoted
  mini-notation string, sampled once per cycle: `add "<0 0 2>"`, `fast "<1 2>"`, `off "<0.125 0.25>"
  rev`. This closes a gap a GoatTracker converter could not work around: a bare `add` applies to a
  voice's whole pattern with no way to scope it to particular orderlist slots. A patterned `add`
  aligned to the same alternation as the notes says exactly that. Constant arguments are
  untouched, taking the same single code path they always did (`argFn()` in `interpret.js`).
- **New transforms**: `struct "x ~ x x"` (take the rhythm from a boolean pattern, the values from
  the pattern), `mask "1 0 1 1"` (its complement - keep the original rhythm, drop the masked-out
  events), `chunk N f` (transform a different 1/N of the cycle each cycle, where `every` transforms
  a whole cycle at a time), `palindrome`, `linger FRACTION`, `rot N` (rotate values, keep the
  rhythm - distinct from `iter`, which rotates whole events), `someCyclesBy P f` (per-CYCLE random
  choice, where `sometimesBy` is per-event), `superimpose f`, and `scale "c4:minor"`.
- **`superimpose` is voice-claiming**, handled exactly like `jux`/`off`: the original and its
  transformed copy sound simultaneously and one SID oscillator can't do that. It is `jux` without
  the stereo-panning connotation, which means nothing on real hardware.
- **`scale` + scale degrees** (`src/pattern/scales.js`, 16 scales) is the one that changes how you
  write: on degrees, `add 2` moves a phrase two SCALE steps and stays in key, where the same `add 2`
  on note numbers shifts two semitones and leaves the key behind. One sharp edge, documented and
  asserted in tests rather than papered over: `miniNotation.js` resolves note names to MIDI at parse
  time, so `scale` genuinely cannot tell degree 60 from the pitch `c4` - a pattern intended for a
  scale must be written as degrees. Strudel draws the same line with separate `n`/`note` controls.
- **Polymeter `{a b c}%4`** in mini-notation: lay a sequence over a fixed step grid, running across
  cycle boundaries instead of restarting, so a 3-step phrase drifts against a 4-step bar.

**Converter follow-up.** The GoatTracker converter now emits orderlist transpose as a patterned
`add` (e.g. `|> add "<0 0 2 0>"`) lined up slot-for-slot with the notes alternation, instead of
declaring a transposed copy of each affected pattern. Its output was also made considerably more
compact (merged `insts` runs, repeated figures folded into `[...]!n`, each voice's patterns emitted
directly above it), with a bit-identical register stream verified at every step. Two points from
that work are about the DSL itself, not the converter:

- **Folding into groups can change timing.** At the top level, mini-notation divides the bar in
  proportion to each step's `@weight`, but a bracketed GROUP is one step of weight 1 no matter
  what it contains. So folding `c3@2 c3 c3@2 c3 c3 c3` into `[c3@2 c3]!2 c3!2` keeps the same
  tokens in the same order and **does not sound the same**: in the original each `c3@2 c3` figure
  spans 3 units and each trailing `c3` 1 unit, but after folding all four steps get equal width. A
  bare `X@w!n` is always safe; group folding is only safe when every resulting group has equal
  total weight.
- **Declarations must precede use.** `interpret()` resolves `$name` in a single source-order walk,
  so a declaration after its use is a compile error.

**Considered and deliberately not added:** `echo`/`stut`. An N-tap echo needs N oscillators
sounding at once, but the voice-split model is binary (`{ left, right }`), so anything past two taps
can't be represented - and two taps is already exactly what `off` does. Adding it would have meant
either an N-way split through `song.js`/`compile.js` or a feature that quietly does less than its
name says. `run N` was also dropped before shipping: mini-notation's `"0 1 2 3"` already says it.

## 2.13.0

- **`$name` splicing now resolves inside an `insts` pattern**, in both short and extended mode. A
  real gap, not an enhancement: `insts` is a mini-notation pattern like any other, but
  `resolveVoiceStmtSigils()` only ever substituted `notes`/`reg`/`sync`/`ring`, so a `$name` in an
  `insts` string reached `miniNotation.js` literally and threw. Text-only resolution (no
  highlighting segments), same as `reg`/`sync`/`ring` - `insts` isn't source-highlighted.

The GoatTracker converter was reviewed alongside: it now hoists identical `insts` slots into
named patterns (which needed the fix above), folds runs with `!n`, and collapses repeated
per-instrument warnings into one. Still open at this version: GoatTracker wave-table rows that set
an **absolute** note (`$80-$DF`, a fixed drum pitch) have no DSL equivalent, since `arp` is
relative-semitone-only (closed in 2.15.0 by `=MIDI` table cells).

## 2.12.0

- **`ramp A..B step D over N`** - a linear per-frame sweep inside any table-valued field, instead
  of spelling out every value. **Saturates at the target**, a deliberate DSL design choice: at
  least one of `step`/`over` is required; with only `over` the step is derived, with only `step`
  the length is.
  [**Corrected by ADR 0004 (2026-08-16)**: the sentence originally here claimed this saturation
  was "exactly what GoatTracker's pulse/filter modulation-step command does" and that one `ramp`
  token "reproduces one original tracker row rather than approximating it," citing
  `ramp 1088..0 step -48 over 128` (a converted GoatTracker filter sweep) as proof. That is
  false: GoatTracker's own accumulator (`gplay.c`'s `filtercutoff`) is an unclamped 8-bit value
  that wraps, not saturates, as a capture of real register writes confirms. That exact example
  reaches its saturating floor after 23 of its 128 frames and holds flat for the remaining 105
  (82%), where the real hardware would instead wrap and keep sweeping. See ADR 0004 for the full correction; `ramp`'s behavior itself is unchanged by this
  note.]
- **`V!N`** - N copies of a table value, deliberately the same repetition spelling mini-notation
  already uses inside a string (`"c4!3"`). Long constant stretches are ordinary in real tables (a
  pulse width held for 104 frames) and read far better than a degenerate `ramp V..V over N`.
- Both are pure expansions performed during interpretation, alongside `$name` table splicing -
  they produce the identical flat array the literal values would, so **no engine change** and no
  new sampling path.

The GoatTracker converter emits both idioms, and hoists any table value two or more instruments
share into a single `table sharedX = ...` declaration, so a converted instrument goes from a
128-number `filterTable` line to `filterTable $sharedFilterTable`. This is presentation-only and
was verified by round-tripping table shapes through the real parser and by comparing register
streams of the hoisted and fully inlined source.

## 2.11.0

- **Shared named tables**: `table NAME = TOKEN+` plus `$NAME` references from any table-valued
  instrument field (`filterTable $sweep`, `pw:$attack`). Deliberately the same shape as
  `pattern NAME = "..."` because it's the same idea one level down - `pattern` shares a note
  sequence between voices, `table` shares a per-frame table between instruments. Motivated by real
  converted tracker data, where every instrument of a song can carry the same byte-identical
  128-entry filter sweep: over a thousand numbers of duplication that this collapses to one
  declaration and one reference per instrument.
- Spliced in place rather than substituted wholesale, so a row can mix a shared table with literal
  values (`pw $attack 2048`), concatenate several, or compose one table from others
  (`table full = $attack $sustain`, flattened at declaration time). Unknown/forward references are
  a compile error naming the known tables, for free from `interpret()`'s single source-order walk -
  the same mechanism that already makes forward `pattern` references an error.
- Implemented entirely in the lexer/parser/interpreter: `$NAME` handling moved into
  `parseTableToken()` (so block fields and short mode's comma-separated inline attrs get it from
  one place), and `expandTableRefs()` runs one pass per instrument rather than per field-`case`.
  **No engine change** - it produces the exact same flat array the literal form would.
- `keywords.js` now re-exports `LOOP_TABLE_NAMES`, which doubles as the canonical list of
  table-valued fields; the editor's completion uses it to decide whether a `$` should offer
  `table` names or `pattern` names, since both can appear on the same short-mode voice line.

## 2.10.0

- **Combined waveform names in a table**: `wave noise+pulse+saw` - several waveforms sounding
  simultaneously in ONE frame slot. A per-frame table slot holds a single value, so the
  mini-notation stack spelling a cycle-sampled wave PATTERN can use (`wave:"[pulse,triangle]"`,
  2.6.0) has nowhere to go here; the combination has to live inside the slot itself. Lexed as a new
  `PLUS` token, joined by `parser.js`'s `parseTableToken()` into one token, and resolved by
  `registerMap.js`'s `waveformBitsFromValue()` (which now splits on `+` and OR-s the bits) - so the
  parser stays free of waveform semantics, exactly as it already was for a plain single name.
  `waveformBitsToName()` is the new inverse, kept beside it.
- `parseTableRow()` (block mode) now goes through `parseTableToken()` rather than collecting tokens
  itself, so the `+` spelling works identically in an `instrument` block and in short mode's
  comma-separated inline attrs - one spelling, one code path.

**Converter fixes shipped alongside** (no grammar change of their own): the GoatTracker
converter had been discarding the wave table's note column, so instruments that differ only in
their arpeggio offsets (a minor and a major chord instrument, say) collapsed into the same sound.
It now decodes that column into a real `arp` field sharing the wave table's loop point, and emits
readable waveform names (`noise pulse saw+pulse+noise`) instead of raw bitmasks, which is what
motivated the `+` spelling above. Absolute-note rows (`$80-$DF`, fixed drum pitches) still had no
relative-offset equivalent and were reported as a warning.

## 2.9.0

Prompted by porting a real GoatTracker `.sng` song to the DSL, whose biggest gaps were all
structural, not one-off bugs. Three of those gaps are feasible engine additions; the rest (see
"Explicitly deferred" below) aren't, at least not yet.

- **Per-note instrument switching**: a new `insts` voice field (extended: `insts "bass bass drum
  bass"`; short-mode inline attr: `insts:"..."`) - a quoted mini-notation pattern of already-
  declared `instrument` NAMEs, sampled at each note's own onset and held for that note's whole
  duration, falling back to the voice's own static instrument when no `insts` value is active.
  This was the single biggest reported gap: GoatTracker constantly switches instruments *within
  one channel* per row (e.g. bass/drum/snare alternating on one voice); DSL v2 previously bound
  exactly one instrument per `voice` statement, for its entire duration. `song.js`'s
  `VoiceChain.instPattern()` / `compile.js`'s onset-gated sampling of it are the new engine
  surface; `dsl/interpret.js`'s `mapPatternValues()` helper resolves the pattern's name tokens
  into actual instrument objects.
- **Per-table independent loop points**: `loop` (an `instrument` block field) now also accepts one
  or more `TABLE INT` pairs (`loop wave 1 pw 3 filterTable 0`), giving wave/pw/arp/filterTable/
  test/gate tables each their OWN loop index, alongside the existing bare-`INT` global-default
  form (unchanged, fully backward compatible). GoatTracker's wave/pulse/filter tables each carry
  an independent loop marker in the original tracker data; DSL v2's old single global `loop N`
  could only ever describe one of them faithfully, forcing a lossy "pick one table's loop point to
  keep" choice when porting a real song. `instrument.js`'s `normalizeLoop()` / `compile.js`'s new
  `resolveLoopIndex(inst, tableName)` implement this; every existing `sampleTable()` call site now
  resolves its own table's loop index instead of reading `inst.loop` directly.
- **Per-frame gate-bit table** (percussive envelope retriggering): `gate` (an `instrument` block
  field or inline attr) now also accepts 2+ space/comma-separated tokens - a per-frame 0/1
  gate-bit PROGRAM, frames-since-onset sampled and loopable exactly like wave/arp/pw/test - instead
  of only ever a single static 0..1 fraction-of-slot (unchanged, still the exact old behavior for
  one value). GoatTracker/SID-Wizard songs routinely toggle the gate bit OFF then back ON *within*
  one note's sustain (a wave-table row program) to retrigger the ADSR envelope for a percussive
  pluck without starting a new note onset - the old single fraction can only ever turn the gate off
  once per note, never back on, so it structurally couldn't express this. New `gateTable` field on
  `instrument.js`'s runtime object; wins over the static fraction in `compile.js` when present,
  same "instrument table wins" precedent already established for wave/pw/test.
- Fixed a latent bug found while implementing the gate-table change: the short-mode inline attr
  `gate:1,0,0` (comma-separated values) silently kept only the FIRST value
  (`buildInstrumentFromInlineAttrs`'s old `opts.gate = a.values[0]`) - now a 2+-value `gate:` inline
  attr routes to the new gate table instead of silently truncating.
- The GoatTracker converter uses all three new constructs instead of the lossy simplifications it
  previously fell back to: every instrument a voice touches gets its own `instrument` block, a
  generated `insts:"..."` pattern reproduces the per-row instrument switching, wave/pulse/filter
  tables keep their own loop points, and a wave table's gate-bit-drop rows become a `gate` table.
  One naming constraint surfaced along the way: `_` is a valid DSL `instrument IDENT` character
  but not a valid mini-notation bare-word character (`[A-Za-z0-9.#]`), so an instrument name used
  as an `insts` token must avoid it.

**Explicitly deferred, not silently dropped** - real gaps the same port raised, kept
out of this change because the DSL genuinely isn't the right layer to fix them (or the source
semantics themselves are unknown, not just unimplemented):
- Pulse-/filter-table **right column**: GoatTracker's own pulsetable/filtertable format has a
  second column whose meaning lives in the player's code, not the `.sng` file format. Nothing to
  build until that meaning is actually known; the raw `reg` escape hatch remains the way to
  hand-drive this once it is.
- **Vibrato-table parameter conversion formula**: same "meaning lives in player code, not
  guessable from the file alone" situation as above - `vib DEPTH RATE` already exists as the DSL's
  own (differently-parameterized, semitones+Hz) vibrato mechanism.
- **Pattern effect commands** (GoatTracker's one-shot-per-row `0xy`-`fxy` command bytes):
  structurally incompatible with this DSL's continuous per-note instrument-program model, not
  merely unimplemented - a real, permanent architectural mismatch, not a gap this DSL should try to
  close.
- **Orderlist restart to a non-zero index**: GoatTracker can loop playback back to any orderlist
  position, not just the start; DSL v2's `<...>` alternation (the orderlist analogue) always wraps
  to element 0. Not encountered in the song that prompted this round, but noted as a known gap -
  would need a real `miniNotation.js` grammar change (a new `<...>@N` restart-index
  spelling), deliberately not undertaken lightly given how carefully `<>`'s existing semantics are
  relied on elsewhere.
- **2SID/stereo songs**: out of scope by design - this compiler targets exactly 3 SID voices (see
  `docs/dsl-grammar-spec.md`'s own framing), not two independent SID chips.

## 2.8.0

- **Short mode can now reference a named `instrument` block directly** - `voice N NAME "..."` /
  `voice N NAME $pattern`, e.g. `voice 0 bass $bassLine`. A real, structural gap found in
  practice: short mode's grammar had no slot for an instrument name at all before this; only
  extended mode (`voice N NAME \n notes ... \n end`) could reference a named instrument.
  Disambiguated from extended mode by one extra token of lookahead in `parser.js`'s
  `parseVoiceStmt()` (`IDENT` then `STRING`/`SIGIL` = short mode with a named instrument; `IDENT`
  then `NEWLINE` = extended mode, unchanged).
- A named instrument and inline attrs together on one voice (`voice 0 bass "c4" wave:pulse`) is a
  compile error, not silent precedence - which one should apply is genuinely ambiguous.
- The editor's completion-context classifier, which had the old shape hardcoded, was updated in
  the same change.

## 2.7.0

- **Fixed a real drift bug**: `keywords.js` (the CodeMirror editor's vocabulary source) hardcoded
  its own copy of `STATEMENT_KEYWORDS`/`INSTRUMENT_FIELDS`/`VOICE_FIELDS`/`TRANSFORMS`/
  `FILTER_FIELDS`, separate from `parser.js`'s real switch statements - `VOICE_FIELDS` was
  missing `wave` (added in 2.6.0) because nothing forced the two to stay in sync in that
  direction (`test/dsl/keywords.test.js` was explicitly one-directional).
- These five lists now live in `parser.js` itself, immediately above the switch statement each
  one describes, and `keywords.js` re-exports them instead of re-declaring them.
  `VOICE_REG_TARGETS`/`FILTER_REG_TARGETS`/`WAVEFORM_NAMES` now derive directly from
  `registerMap.js`'s real objects (`Object.keys(...)`) instead of being retyped.
- Added `DSL_VERSION` (`src/dsl/version.js`) and this changelog.

## 2.6.0

- `wave:"..."` (quoted) - a full mini-notation waveform pattern, e.g. `wave:"pulse|triangle"`
  (random per-cycle choice) or `wave:"[pulse,triangle]"` (combined bits). Routes to `song.js`'s
  pre-existing `VoiceChain.s()` (cycle-position sampled), distinct from `wave:pulse` (bare - a
  per-frame instrument table, frame-since-onset sampled). Added the extended-mode counterpart,
  the `wave "..."` voice field.
- Incidentally closed a previously-documented gap: combined/simultaneous waveforms, since `.s()`
  never had that limitation - it just wasn't wired up by the interpreter yet.
- One documented, accepted precedence caveat: an attached instrument's own `wave` table still
  wins over a voice-level `wave:"..."` pattern if both are present on one voice.

## 2.5.0

- Source-highlighting anchor upgraded from a single `(docStart, docEnd)` span
  (`noteSourceDocStart`/`noteSourceDocEnd`) to `noteSourceSegments`, an ordered list of
  `{textStart, textEnd, docStart}` spans (a small source map) - the single-span design broke
  (no highlight at all) the moment a voice combined more than one named pattern (e.g.
  `"[$counterMelody]|[$counterMelody2]"`), since the resolved text is assembled from pieces
  living in genuinely different places in the document. Fully recursive: a pattern referenced
  inside another pattern inside a voice all resolve correctly, including the same name
  referenced more than once.

## 2.4.0

- Source-highlighting anchor introduced: `noteSourceDocStart`/`noteSourceDocEnd`, tracked from
  `lexer.js`'s STRING tokens (which gained `docStart`/`docEnd`, the token's absolute document
  position) through `parser.js`/`interpret.js`/`song.js`/`compile.js` to the editor, replacing a
  runtime `String.indexOf(noteSourceText)` search that could highlight the wrong occurrence
  (identical text existing elsewhere in the document for an unrelated reason) or never
  correctly anchor a `$name`-referenced voice at all.

## 2.3.0

- Short mode's inline attrs (`wave:`/`adsr:`/`pw:`/...) and `|>` pipeline steps may continue on
  their own lines after a multi-line backtick `notes` string closes, not just on the line the
  closing backtick is on - a real gap once `notes` could be multi-line (2.2.0): without this,
  the line break right after a multi-line string's close would end the voice statement, and
  whatever followed would fail to parse as a bogus new top-level statement.

## 2.2.0

- Backtick strings (`` `...` ``) added as a second string delimiter alongside `"..."`, may span
  multiple physical lines (`"..."` must still close on the line it starts). Embedded newlines
  are preserved literally, not collapsed to a space - `miniNotation.js`'s own tokenizer already
  treats any whitespace, including `\n`, as an inter-token separator, so no normalization is
  needed, and highlighting keeps working with zero downstream changes.

## 2.1.0

- `sync`/`ring` voice fields added (hard sync / ring modulation, control-byte bits 1/2) - a
  real gap in the grammar's first cut, found while migrating the editor's cheatsheet: `song.js`'s
  `VoiceChain.sync()`/`.ring()` already existed and already worked, the grammar just never
  exposed a construct for them. Bare `0`/`1` constant in short mode, or a full quoted
  mini-notation pattern in extended mode if the bit should vary over time.

## 2.0.0 - initial release

Replaced the v1 JS-hosted builder (`song()/.voice()/.filter()` evaluated via `new Function(...)`)
with a real, standalone, non-JS grammar: `tempo`/`pattern`/`instrument`/`voice`/`filter`
statements, short mode (single-line `voice N "..." attr:value ... |> transform`) and extended
mode (`voice N NAME \n field... \n end`), named patterns (`pattern NAME = "..."`, `$name`
reference/splicing), table-driven instruments (`wave`/`arp`/`pw`/`filterTable`/`test`, sampled
by frames-since-note-onset - tracker semantics) alongside cycle-position-sampled note/filter
patterns (Strudel semantics), the `|>` transform pipeline, filter `switch cycle` piecewise
automation, a raw register escape hatch (`reg`), and tempo-as-pattern. See
`docs/dsl-grammar-spec.md` for the full current spec.
