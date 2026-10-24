# ADR 0003: A notation for writing good SID songs

Status: implemented (4.0.0)
Date: 2026-08-07
Affects: `rasterklang-dsl`, `rasterklang-score` (compiler and pattern engine), and tools that emit DSL
Current DSL version: 3.0.0. This ADR targets 4.0.0.

**Priority statement.** The goal is a notation in which good SID songs are
easy to write. That goal outranks backward compatibility, including
compatibility with tools that generate DSL source; those are updated after
the notation, not allowed to narrow it.

Throughout, "the notation draft" refers to an earlier, unpublished
composer-side proposal for a SID live coding language. Several decisions
below adopt or reject its ideas, and the reasons are recorded here so the
draft itself is not needed to follow them. File and line references such as
`compile.js:216` describe the code as it was at 3.0.0; `compile.js`,
`song.js` and `miniNotation.js` now live in `rasterklang-score`.

---

## 1. Context

Two independent reviews reached the same conclusion. ADR 0001 diagnosed the
timing model from the implementation side; the notation draft reached it from
the composer's side. Both identified that a proportional pattern model leaves
a composer with no shared metric. ADR 0001 shipped `pattern name N = "..."`
and 3.0.0 does apply that number.

The number works. Everything around it does not, in two distinct ways: the
metric model collapses as soon as a pattern is not the whole voice (1.1), and
the instrument layer makes classic SID technique harder to write than it needs
to be (1.4).

### 1.1 Measured evidence: the timing model

Measured on 3.0.0 at `tempo 120 beats 4`, so one bar is 100 frames and one
beat is 25 frames. Onsets are frequency-register changes on voice 0, read
frame by frame out of `compile()`'s frame source by a small probe script.
(Gate is useless as an onset signal for legato material, since it stays high
across notes; the frequency registers are the reliable signal.)

```
pattern lead 8 = "c4 d4 e4 f4 g4 a4 b4 c5"   voice 0 $lead
  -> onsets 0 25 50 75 100 125 150 175 (loop 200)     one note per beat, correct

pattern lead   = "c4 d4 e4 f4 g4 a4 b4 c5"   voice 0 $lead
  -> onsets 0 13 25 38 50 63 75 88 (loop 100)         8 notes squeezed into one bar
```

The declared length is real. Three failures surround it.

**F1. A metric pattern spliced into a larger string silently loses its length,
with no diagnostic at all.**

```
pattern a 4 = "c4 d4 e4 f4"
voice 0 "[$a] g4"
  -> onsets 0 13 25 38 | 50 | 100 113 125 138 | 150
```

`$a` declared 4 beats and plays in 2. Nothing is reported.
`analyze.js`'s `referencedPatternLength()` returns null for anything that is
not a whole-string reference, and the voice falls back to proportional
auto-fit. Its comment argues that guessing would be worse than falling back,
which is right; falling back *silently* is the worst of the three options.

**F2. `<$a $b>` is not an orderlist, because `$name` is raw text
substitution.**

```
pattern a 4 = "c4 d4 e4 f4"
voice 0 "<$a $a>"
  -> onsets 0 100 200 300 400          ONE note per bar
```

The splice produces `<c4 d4 e4 f4 c4 d4 e4 f4>`, an eight-element alternation,
not a two-element orderlist. `<[$a] [$a]>` is correct (onsets every 25 frames)
and only the bracketed form is. ADR 0001 section 1.2 asserted that `<...>`
"already is a per-voice orderlist"; that holds for the single-token elements
a GoatTracker converter emits and breaks for any hand-written multi-note
pattern.

**F3. Sequencing patterns of different lengths in one voice is a hard error.**

```
pattern a 4 = "c4 d4 e4 f4"
pattern b 8 = "c5 d5 e5 f5 g5 a5 b5 c6"
voice 0 "<[$a] [$b]>"
  -> error: voice 0's orderlist mixes patterns of different declared lengths
```

The error suggests "split them across voices", which is not an answer: a voice
is a SID voice and there are three. **Song structure is therefore not
expressible.** An intro/verse/chorus arrangement of differing lengths in one
voice cannot be written at all, which is the most basic thing a tracker
orderlist does, and a central concern of the notation draft.

### 1.2 Root cause

Declared length is a **voice-level** property wearing a **pattern-level**
costume. `interpret.js:805` applies it as
`chain.notePattern.slow(N / barBeats)`: one multiplier over the voice's whole
note pattern. That is why it only works when the pattern is the whole voice,
and why elements of different lengths cannot coexist. A single `slow` factor
cannot express a timeline.

### 1.3 Length is declared, not written

Even where it works, the length is a number the author must compute and keep
in sync. ADR 0001's own error message says it out loud:

```
pattern "a" declares 4 beats but its contents add up to 3 beats
```

The author writes music and then separately asserts its length. The notation
draft makes the better case: if a note carries its own duration, the
pattern's length is the sum and nothing needs declaring. `@0.25` is
arithmetic; `/16` is music.

### 1.4 The instrument layer: expressible, but not writable

Everything classic SID technique needs is *expressible* in 3.0.0, because
`wave`/`arp`/`pw`/`filterTable` are per-frame tables and a table can encode
anything. That is not the same as writable. Checked against
`parser.js`'s `INSTRUMENT_FIELDS`
(`wave, adsr, pw, arp, filterTable, loop, vib, vibTri, slide, test, gate`)
and `compile.js`, these are real authoring gaps:

| technique | notation draft | 3.0.0 | gap |
| --- | --- | --- | --- |
| PWM sweep | `pwm 400 -> 800 speed 2 pingpong` | `ramp 400..800 step 20` | `ramp` saturates; no pingpong, no loop, so a sustained PWM needs a hand-unrolled table |
| arpeggio speed | `arp 0 4 7 speed 1` | `arp 0 4 7` | speed is **song-global** (`song.js:352`, `arpeggioFrameRate` default 3), not per instrument. Two instruments cannot arpeggiate at different rates |
| arpeggio delay | `arp 0 4 7 delay 3` | none | must be padded into the table by hand |
| hard restart | `restart hard 2` | `gate -2` | works, but nothing about `gate -2` says "hard restart" to a reader |
| legato | `gate legato` | none | no way to say "do not retrigger between notes" |
| gate percent | `gate 80%` | `gate 0.8` | cosmetic, but `%` is what a musician writes |
| slide to a note | `slideTo g4 speed 3` | `slide TIME` | portamento time exists; a target pitch does not |
| readable tables | `set 400; add 20 repeat 8; loop` | bare token rows | a 32-entry pulse table is 32 hand-written numbers |

The pattern under all of these: 3.0.0 exposes the *mechanism* (a per-frame
table) and not the *intent*. For a live coding editor that is backwards. The
mechanism must stay, since it is the escape hatch, but the intent needs a
spelling.

### 1.5 What must be preserved

Mini-notation and the transform pipeline (`[]`, `<>`, `{}`, `*n`, `!n`, `?p`,
euclidean, `rev`, `every`, `iter`, `jux`, `degradeBy` and the rest of
`TRANSFORMS`) are the reason this is a live coding language and not a
tracker with a text file. The notation draft dropped all of it while still
listing live coding as a goal. Nothing here removes any of it except one
redundant operator (D1).

---

## 2. Decisions, part A: time and structure

### D1. Durations are written on the note, and that is the metric opt-in

```
c4/1   whole      c4/8    eighth       c4/4.  dotted quarter
c4/2   half       c4/16   sixteenth    c4/4t  quarter triplet
c4/4   quarter    c4/32   thirty-second
~/8    rest       _/4     hold
```

A pattern containing **any** `/N` token is metric. Its length is the sum of
its durations. No declared number, no arithmetic, and no domino effect: each
duration is absolute and local, so changing one note's length does not move
the others relative to the grid.

`@n` stays legal and unchanged. `/N` and `@n` in the same pattern is an error:
one is a duration and the other a proportional weight, and mixing them has no
single sensible reading.

**Duration in beats.** In `meter N/D` the beat unit is a `1/D` note and a bar
is `N` beats. So `/K` is `D/K` beats, `/K.` is `1.5 * D/K`, and `/Kt` is
`(2/3) * D/K`. In 4/4 an eighth is 0.5 beats, which is what a composer expects.
An unannotated token is one beat, so in 4/4 the default is a quarter, matching
the notation draft.

**Rest and hold keep the 3.0.0 spelling, against the notation draft.** The
draft uses `.` for rest and `~` for hold. `~` is already the rest token in
mini-notation. Adopting the draft's spelling would silently invert the meaning
of `~` in every existing song and every converter's output, turning rests into
holds with no diagnostic possible. Rest stays
`~`/`-`, hold stays `_` (ADR 0001's D6). `.` is not added as a rest synonym:
`scanWordToken` (`miniNotation.js:266`) accepts `.` inside a word, so it would
be ambiguous.

**The `/` collision, and why it is acceptable.** Mini-notation already uses
`/n` as the slow operator (`miniNotation.js:330`, `pattern.slow(...)`), so
`c4/2` today means "c4 slowed 2x". This ADR takes `/` for duration and removes
slow from mini-notation. Measured justification: `/n`-as-slow appears in no
converter output, no fixture, no example song, and no test asserts it. It is
also fully redundant, since `*0.5` is the identical
operation inside a string and `|> slow 2` is the identical operation on a
voice. Rejected alternatives: `c4:8` (`:` already reads as a label in inline
attributes), `c4=8` (no precedent), `@1/8` (keeps the weight concept and its
arithmetic in the author's head, which is what D1 exists to remove).

**Implementation note.** A duration is a weight (see 5.1). The parse shape is
`'/' INT ('.' | 't')?`, matched explicitly so the number parser does not
swallow the dot of `/4.` as a decimal point; it sets the step's weight to
`beatUnit/N` (dotted x1.5, triplet x2/3) through the same `timeCatWeighted`
path `@n` uses. At `meter 4/4 tempo 120`, `"c4/4 d4/8 e4/8 f4/2"` gives onsets
0, 25, 38, 50 and loops at 100, frame-identical to `"c4@1 d4@.5 e4@.5 f4@2"`.

### D2. `meter N/D` replaces `beats N`

`/8` has no meaning without a beat unit, and `beats 4` never said what a beat
is.

```
meter 4/4      ; bar = 4 quarter notes
meter 7/8      ; bar = 7 eighth notes
```

`beats N` stays legal as an alias for `meter N/4` so 3.0.0 songs keep working,
documented as deprecated. ADR 0001's D2 reduced `beats` to a pure time
signature; this gives it the denominator it always implied. ADR 0001's
migration hint (its decision 10, flagging unusual `beats` values as probable
stretch levers) stays and gains a better suggestion: write the durations.

`meter` is a soft keyword on the clock statement carrying `NUMBER '/' NUMBER`;
combining `meter` and `beats` on one statement is an error. `tempo 120 beats 4`
and `tempo 120 meter 4/4` produce byte-identical onset frames.

### D3. `each 1/D` for step-sequencer patterns

```
pattern hats each 1/16 = "c6 ~ c6 c6 ~ c6 ~ c6"
```

Every token is one sixteenth; length is derived. This is the tracker step grid
and the shortest possible spelling for a drum pattern. The notation draft's
distinction is worth keeping verbatim: `each` sets the
duration of *each event*, a declared length sets the duration of the *whole
pattern*.

`each` and `/N` compose: `each` is the default, a `/N` token overrides it for
that token.

**Implementation note.** `each` is a source-level expansion (`expandEach` in
`miniLength.js`/`minilength.go`, the same shape as `expandHolds`): every
unweighted token's default weight becomes `beatUnit/D`, and the result is
baked into the pattern's stored text at declaration time, since `each` has
no meaning once the pattern is referenced from a voice. An `each`-only pattern
(no declared number) derives an effective declared length from its summed
durations, so a bare `$name` voice reference is stretched exactly as a
written number would stretch it. `each` plus a declared length is legal; the
declared length is checked against the derived one.

### D4. Arrangements: pattern references sequence, repeat, group and nest

```
pattern verse  = $riffA $riffA $riffB $riffA
pattern chorus = $riffB*2 $fill $riffA
pattern song   = $intro [$verse $chorus]*2 $verse

voice 0 bass $song
voice 1 lead $leadA $leadB |> every 4 rev
```

Rules:

- A sequence of `$refs` concatenates in beat time. Each element occupies **its
  own** metric length. Elements of different lengths are legal and are the
  entire point (this is F3).
- `$name*n` repeats an element; `[$a $b]*n` repeats a group.
- References nest to any depth, so `pattern` doubles as the song-part
  construct and no separate section system is needed: voice, song part,
  musical pattern, notes.
- A voice's loop length is the sum of its elements. Voices with different sums
  loop independently against one shared clock, which is already how 3.0.0
  behaves across voices.
- An element with no metric length (a proportional pattern) occupies one bar,
  so the construct is usable before every pattern is converted.
- Recursive references are a compile error naming the cycle.

**Deliberately not a mini-notation bracket.** `{...}` was the obvious spelling
and is already taken: `miniNotation.js:296` dispatches `{` to
`parsePolymeter()`. The arrangement lives at statement level in the ANTLR
grammar instead, where `$` is already unambiguous and where lengths are
already computed. Mini-notation therefore does not change for D4 at all.

`<...>` alternation is untouched: one element per bar, phasing against other
voices, still what GoatTracker conversions rely on. The two are now clearly
distinct. `<a b>` alternates by bar; `$a $b` concatenates by length.

**Implementation note.** The grammar gains `arrangement: arrangementElement+`
(a sigil or a `[...]` group, each with an optional `*n`), used by the
short-mode `voice` statement and by a second `patternDecl` alternative
(no quotes, no declared length). A bare single sigil (`voice 0 $lead`)
stays byte-identical to the old behaviour. Extended-mode voice syntax's
`notes` field was not extended, so an arrangement is not writable there.
Lengths are computed bottom-up over the reference tree
(`resolveArrangementLength()`), replacing the old whole-string
`referencedPatternLength()` match and the different-lengths-in-an-orderlist
error. Because declarations are forward-only, the reachable recursion case is
direct self-reference, which is reported as a recursive reference.

Measured at `meter 4/4 tempo 120` with patterns of 4 and 8 declared beats,
`voice 0 $a $b` gives onsets every 25 frames from 0 to 275 and loops at 300
(F3, previously a compile error). `$a*2 $b` is onset-identical to `$a $a $b`.
`[$a $b]*2` and a nested arrangement are onset-identical to their unrolled
forms within one frame: composing `.fast()`/`.slow()` through a nested
`timeline()` introduces floating-point rounding of the same class as D17's
subdivision jitter. Note values and counts match exactly.

### D5. Metric length never degrades silently

Once D4's timeline exists, F1 has a correct answer rather than a fallback: a
metric pattern referenced anywhere occupies its metric length. Where a reading
is genuinely ambiguous it is an **error naming the ambiguity**, never a silent
rescale. The rule is absolute: no construct may change a metric pattern's
length without saying so.

`$name` also stops being text substitution and becomes a resolved reference.
That fixes F2 structurally rather than by warning: `<$a $b>` becomes a genuine
two-element alternation and `$a!3` repeats the pattern rather than its last
token. ADR 0001's decision 7 chose diagnostic-only for backward compatibility;
F2's measurement shows the silent behaviour is severe enough to be worth the
break, and this ADR does not trade correctness for compatibility.

**Implementation note.** `$name` is a first-class mini-notation atom, resolved
through a `refs` option on `parseMini`, not by text substitution. A bare
`$ref` atom gets a default sequence-slot weight equal to the referenced
pattern's own beats (overridable by an explicit `@n` or `/N`), which is F1's
fix. `!n` and `<...>` needed no special-casing: since a reference resolves to
the referenced pattern object, the existing replicate and alternation
machinery does the right thing, which fixes F2. Every Hap's
`sourceStart`/`sourceEnd` is now an absolute document position by
construction, and the old `$name!n` splice warning is removed because the
hazard can no longer occur.

Measured at `meter 4/4 tempo 120`: `"<$a $a>"` (pattern `a` of 4 beats) gives
onsets every 25 frames; `"[$a] g4"` gives onsets `0 20 40 61 80 100`, i.e.
`$a`'s 4 beats against `g4`'s 1; `$a!3` replicates all four notes.

Scope limits: `insts`, `fx`, `reg`, `sync`/`ring` and tempo-as-pattern still
resolve `$name` by text splicing (correct text, coarser editor highlighting).
Only a bracket wrapping exactly one bare `$ref` inherits that reference's
weight; a multi-element `[...]` still defaults to weight 1.

### D6. `|` is a checked bar line when it stands alone

```
pattern lead = "c4/4 d4/4 e4/4 f4/4 | g4/4 a4/4 b4/2"
```

In a metric pattern a whitespace-delimited `|` asserts a bar boundary. A `|`
that does not land on one is an error reporting the beat offset it reached.
Purely a check, no effect on timing. It also makes a one-bar-per-line layout
(already supported by backtick strings) self-verifying.

**No collision with the choice operator**, verified: `peekChar()`
(`miniNotation.js:219`) does not skip whitespace and `parseChoice()` (`:306`)
only treats `|` as choice when it is immediately adjacent to the preceding
atom. `a|b` is a choice; `a | b` currently scans `|` as a bare word atom,
which is a meaningless value no song relies on. Only the standalone form is
reclaimed.

ADR 0001 deferred bar lines because they cannot be placed where a note ties
across the bar. D1 removes that objection: a tie is `c4/4 _/4` and the `|`
goes between the tokens.

### D7. The declared number is demoted to an assertion, honestly documented

`pattern lead 32 = "..."` stays and keeps checking that the contents sum to
32. It is no longer presented as the mechanism that makes a pattern metric,
because after D1 the durations do that. This matches what users already
perceive it to be.

### D8. Proportional auto-fit becomes loud

A pattern with no `/N`, no `each` and no declared number keeps today's
behaviour, and the compiler reports the duration it derived:

```
info: pattern "lead" is proportional: 8 tokens over 1 bar = 1/8 note each.
      Write durations (c4/8) or `each 1/8` to fix the grid.
```

The 32-notes-squeezed-into-a-bar case (ADR 0001 section 1.1) never again
reaches the user as sound-only feedback. The editor already has the bar grid
from ADR 0001's D7 to show it.

---

## 3. Decisions, part B: instrument expressiveness

The governing principle, shared with the notation draft: every high-level
construct compiles down to the same per-frame table engine that exists today.
Nothing here is a new runtime mechanism. They are spellings for intent, and
the raw table stays available as the escape hatch for everything they do not
cover.

### D9. Sweeps get shapes: `pingpong` and `loop`

```
pw   ramp 400..800 step 20 pingpong
filterTable ramp 300..1800 step 12 pingpong
```

`ramp` exists (`ramp NUMBER '..' NUMBER ('step' N)? ('over' N)?`) and
saturates at its target, so a sustained PWM has to be hand-unrolled today. Add
two terminators:

- `pingpong`: on reaching the target, reverse and sweep back, forever.
- `loop`: on reaching the target, jump back to the start, forever.
- absent: today's saturating behaviour, unchanged.

This covers the notation draft's `pwm A -> B speed S pingpong` and
`cutoff A -> B speed S pingpong` without inventing a second sweep syntax. The
draft's `speed S` is `step` here: `step` is frequency units per frame, which
is the primitive the engine has.

The whole change is one function, `expandRamp` (`interpret.js`), which
produces the per-frame array. Its current clamp is deliberate and remains the
default. (Its source comment also claimed that this clamp is what a
GoatTracker modulation-step row does; ADR 0004 corrects that claim, and shows
that `loop` is the way to write GoatTracker's wrapping filter sweep.)

A dedicated `pwm` field is deliberately **not** added: `pw ramp ... pingpong`
says the same thing with one concept instead of two.

**Implementation note.** `pw ramp 400..800 step 20 pingpong` produces a
symmetric triangle at register level. A shape terminator sets the table's
implied loop index to 0. Without any terminator, a ramp that fills its whole
table still wraps to entry 0 at the end through the ordinary default
`loopIndex 0`; true saturation (holding at the target within an
`over`-padded ramp) is unchanged.

### D10. Every instrument table gains `speed`, `delay` and keeps `loop`

```
instrument lead
  arp 0 4 7 speed 2 delay 3
  wave pulse noise speed 4
end
```

`arpeggioFrameRate`
(`song.js:352`, used at `compile.js:577`) is *not* the instrument arp table's
speed. It is the **chord round-robin rate**: when more than one Hap is
simultaneously active (a stacked `[c4,e4,g4]`), the compiler cycles through
those tones at that rate. It never touches `inst.arp`.

The instrument's own tables are sampled by `sampleTable(table,
framesSinceOnset, loopIndex)` (`compile.js:216`), which advances **exactly one
entry per frame** with no rate control whatsoever. Today the only way to slow
an arpeggio is `arp 0!2 4!2 7!2`, hand-replicating every entry, which is
precisely the arithmetic the notation draft's `speed` exists to remove, and it becomes
unmanageable the moment the table is long or the speed changes.

So rather than an `arp`-only fix, this generalizes to the sampling function
all tables already share:

- `speed N`: hold each entry for N frames. Default 1, today's behaviour.
- `delay N`: wait N frames after note-on before the table starts advancing,
  so an attack transient stays clean. Default 0.
- `loop` already exists per table (`resolveLoopIndex`, `compile.js:242`, with
  a real per-table map so a wave table's loop marker and a filter table's are
  independent). `speed` and `delay` follow the same per-table shape.

One mechanism, available on `wave`, `arp`, `pw`, `filterTable`, `test` and
`gate`. This is strictly more than the notation draft asked for (it wanted
speed and delay on `arp` and delay on `vibrato`) and costs less, because it is
three parameters on one existing function instead of a special case per field.

`arpeggioFrameRate` is untouched and keeps its real job.

**Implementation note.** The sampled index is
`Math.floor(max(0, f - delay) / speed)`, so entry 0 is held for `delay + 1`
frames (frames 0 to `delay` inclusive). Verified at register level:
`arp 0 4 7 speed 2` is frame-identical to `arp 0!2 4!2 7!2`.

### D11. Gate modes say what they mean

```
gate 80%           ; percentage of the slot, same as gate 0.8
gate legato        ; do not retrigger between consecutive notes
gate retrigger     ; always retrigger, even on a repeated pitch
hardRestart 2      ; gate off 2 frames before the next note
```

- `80%` is a spelling of the existing fraction. `gate 0.8` stays legal.
- `hardRestart N` is the readable spelling of the existing `gate -N`
  (GoatTracker's gatetimer). `gate -N` stays legal. Nothing about `gate -2`
  tells a reader it is a hard restart, and hard restart is not an advanced
  technique in SID music, it is table stakes.
- `legato` is a declarative surface over an existing, verified path.
  `compile.js:530-540` computes `toneportaLegato` for the `fx` lane's `tp`
  op: it suppresses the onset, keeps `framesSinceOnset` counting so the
  instrument's tables keep running, and emits no forced gate pair. That is
  exactly legato. Today it is reachable only through a per-row `fx` command;
  `gate legato` makes it an instrument property.
- `retrigger` is a **spelling over `gate -N`, not over the onset forcing.**
  `compile.js:941`'s own
  comment is explicit: the forced gate-off/gate-on pair emitted on every
  onset "restarts the envelope but is invisible to anything that samples the
  registers once per frame", whereas `gate -N` produces "a genuine SID hard
  restart" with a real, frame-observable low-to-high edge. So forced control
  is not a retrigger mode that can be toggled; it already happens on every
  onset and cannot be observed per frame. `gate retrigger` therefore means
  "guarantee a frame-observable re-attack on every note, including a repeated
  pitch", and must be implemented through the `gateReleaseFrames` path with a
  small default, not through `forcedControl`.

Legato matters because it is the difference between a melody line and a
sequence of separate notes, and it is a prerequisite for `slideTo` (D12) to
sound like anything.

**Implementation note.** `gate N%` is folded into the `NUMBER` lexeme
(`'-'? [0-9]+ ('.' [0-9]+)? '%'?`), so `80%` and `0.8` are the same value from
the moment the token is read. `hardRestart 2` produces the same AST as
`gate -2` and is frame-identical to it. `gate retrigger` defaults
`gateReleaseFrames` to 1 when no more specific `gate`/`hardRestart` value sets
one, giving four real low-to-high edges for four same-pitch notes. `gate
legato` and `slideTo` share one onset-suppression primitive in `compile.js`
(`suppressOnset`, generalising the `toneportaLegato` check). Gate modes are
read from an instrument field only; the voice-level `gate` override was not
extended. For a voice whose instrument switches per note (`insts`), legato
suppression reads the ending note's instrument, to avoid a circular
dependency on the onset decision itself.

### D12. `slideTo NOTE` for portamento with a target

```
instrument lead
  slide 0.5             ; portamento time, unchanged
end

pattern riff = "c4/4 g4/4 slideTo:c5/2"
```

`slide TIME` (portamento time as a fraction of a beat) exists; a target pitch
does not, so a slide can only be to whatever note comes next. The notation
draft wants both `slide +12` (relative) and `slideTo g4` (absolute). The
absolute form is the useful one and maps onto `CMD_TONEPORTA`, which the
compiler already implements.

Spelled as a per-note attribute rather than an instrument field, because it is
a property of a specific note, not of the sound.

**Implementation note: a per-note tag on the Hap.** Mini-notation atoms were
plain values with no per-note attribute channel, so this needed a decision.
`slideTo:NOTE` marks its Hap (`{ value, slideTo: true, ... }`, through the
same extra-field mechanism that carries `sourceStart`/`sourceEnd`) so the
compiler treats the transition as legato (no retrigger, `framesSinceOnset`
keeps counting) even when the instrument's gate mode is not `legato`. The
existing `slide TIME` glide does the rest, because it already interpolates
toward whatever note is currently active. Two alternatives were rejected:
routing through the `fx` lane's `tp` would mean authoring two parallel
patterns in row-for-row sync for one musical idea, which is the ergonomic gap
D12 exists to close; and wrapping the note value in an object would touch
every arithmetic site in `compile.js` that reads a note as a bare number,
while tagging the Hap touches only the onset decision. Verified at register
level: a monotonic frequency ramp reaching exactly the target register, with
no control-register write at the transition.

### D13. Table verbs as an alternative to bare token rows

```
instrument bass
  pw
    set 400
    add 20 repeat 8
    add -20 repeat 8
    loop
  end
end
```

Compiles to exactly the same per-frame table as the bare-token form, which
stays legal and is still the right spelling for short tables. The verb form
exists because a 32-entry pulse table written as 32 numbers is unreadable and
unmaintainable, and because `add ... repeat N` says what the author meant.

Verbs: `set V`, `add V [repeat N]`, `wait N`, `loop`, `stop` (hold last
value). Available on any table field: `wave`, `pw`, `arp`, `filterTable`,
`test`, `gate`.

**`loop` collides with the existing instrument field and is unified with it,
not renamed.** `loop` is already an `INSTRUMENT_FIELD`
(`loop INT`, or `loop wave 4 pw 2` pairs) setting a per-table loop point by
index. The verb `loop` inside a table block marks the loop point **at its own
position**, which is the same concept expressed positionally instead of by
index, and is strictly easier to keep correct when the table above it is
edited. The two forms mean the same thing and must compile to the same loop
index. Using both on one table is an error naming the conflict.

Rejected: making this the only spelling. `wave noise noise triangle` is
clearer than three `set` lines, and converters emit bare rows.

**Implementation note.** `instrumentField` gains a block alternative
(`anyIdent NEWLINE tableVerbLine+ END NEWLINE`), distinguished from the
bare-row form by the very next token, so no extra lookahead is needed; every
non-table field rejects the block form with a named error. Verbs expand at
interpret time (`expandTableVerbs`): `set` and each `add` step push one entry,
and `wait N` holds the current value for N more entries. `loop` and `stop`
both just record the loop point at their own position; the table's
out-of-range clamp to its last entry is what makes a trailing `loop`/`stop`
mean "hold the last value". Using an instrument-level `loop` field and a verb
`loop`/`stop` for the same table is an error, whereas a ramp's implied loop
point (D9) silently yields to an explicit `loop` field.

### D14. A named error catalogue

The notation draft is right that the compiler's errors are part of the
notation. Every one of these is a specific message with the offending text and
a suggestion, not a generic parse failure:

```
unknown instrument, unknown pattern, unknown table
more than 3 voices
invalid note name              (e.g. "h4", suggest "b4")
invalid duration               (e.g. "/7", list the legal subdivisions)
ADSR value out of 0..15
cutoff above 2047              (already ADR 0001 decision 9)
recursive pattern reference    (name the cycle)
bar line off the grid          (D6, name the beat offset)
mixed /N and @n in one pattern (D1)
proportional auto-fit          (D8, info level)
```

`|>` on a pattern declaration already has a good message (ADR 0001 decision 8)
and is the template for tone.

**Implementation note.** Two entries were missing before 4.0.0 and were added:
an invalid note name check (`checkNoteNames`, scoped to `notes` fields only so
mini-notation stays field-agnostic; `h4` gets the German-notation suggestion
"did you mean b4") and an ADSR range check (`validateAdsr`, shared by the
instrument block and the short-mode `adsr:` attribute, which parse through
different grammar productions). Both are interpreter-only semantic checks
with nothing for the Go side to mirror.

### D15. `chip 6581` / `chip 8580`

```
chip 6581
```

The engine already honours the SID model: `liveEngine.js` has a
`DEFAULT_MODEL`, the editor exposes a 6581/8580 selector, and `sidExport.js`
writes the model bits into the `.sid` header flags. The only thing missing is
that a **song cannot record which model it was written for**, so the choice
lives in UI state and is lost on save, re-open and export.

That is a real defect for a publishable tool, since 6581 and 8580 differ
audibly in filter behaviour and in how combined waveforms sound, and a song
authored against one can be plainly wrong on the other. `chip` joins `target`
as song-level configuration. Absent means today's default.

**Implementation note.** `chip` is its own top-level statement
(`chipStmt: CHIP NUMBER`), mirroring `target` but with a bare number since
there is exactly one setting. It is carried on the compiled song as
`chipModel` and written to the PSID flags (model bits 4-5: `10` for 8580,
`01` for 6581); a `.sid` export prefers the song's declared model over the
UI selection.

### D17. The frame grid is reported, and misfitting subdivisions are named

The two-clock model (musical time versus SID tick time) is real in the
implementation, but it is worth being precise about *how*, because the answer
determines whether a written `/16` actually sounds even.

**How it works today, verified.** Musical time is rational and is never stored
in frames. The only conversion point is `makeCycleClock`
(`compile.js:265-275`):

```
framesPerCycle = round(framesPerMinute / bpm * beatsPerCycle)
resolve(frame)  -> { cycleIndex, posFrac }     posFrac = (frame % framesPerCycle) / framesPerCycle
```

Frames are the SID tick, at whatever rate `target` declares (default PAL
985248/19705 = exactly 50.00 fps, ADR 0002). Every frame *samples* the
pattern at its own `posFrac`. There is no snapping of note starts to frames:
a note begins on the first frame whose `posFrac` has passed its boundary.
Quantization is therefore a consequence of sampling, and it is per frame, not
per note.

**What that means for the metric, measured.** Sixteen sixteenth notes, onset
frames:

```
tempo 120 beats 4   bar 100 frames, 1/16 = 6.25 frames
  0 7 13 19 25 32 38 44 50 57 ...        spacing 7,6,6,6,7,6,6,6

tempo 132 beats 4   bar  91 frames, 1/16 = 5.69 frames
  0 6 12 18 23 29 35 40 46 52 ...        spacing 6,6,6,5,6,6,5,6

speed 8 beats 4     bar  32 frames, 1/16 = 2 frames exactly
  0 2 4 6 8 10 12 14 16 18 ...           spacing 2,2,2,2,2,2,2,2
```

Two separate effects, both invisible today:

1. **Subdivision jitter.** At `tempo 120` a sixteenth is 6.25 frames, so
   onsets alternate 7,6,6,6. The written rhythm is even; the played rhythm is
   not. This is inaudible at a sixteenth and very audible at a thirty-second
   or a triplet.
2. **Tempo rounding.** `tempo 132 beats 4` computes 90.909 frames per bar and
   rounds to 91, so the song actually plays at 131.87 BPM. ADR 0001's D4
   introduced `speed` precisely for this, but nothing tells an author it is
   happening.

**Decision.** The compiler reports the grid, and names a misfit rather than
letting the author discover it by ear:

```
info:  meter 4/4, tempo 120, PAL 50.00 fps -> 100 frames/bar, 25 frames/beat.
warn:  the smallest duration in pattern "hats" is 1/16 = 6.25 frames, which is
       not a whole number of frames; onsets will jitter by up to 1 frame.
       `speed 8` gives exactly 2 frames per 1/16 at this meter.
warn:  tempo 132 rounds to 91 frames/bar, which is 131.87 BPM.
       `speed 6` is the nearest exact spelling.
```

Warning level, never an error: uneven subdivisions are legal and often
inaudible, and a composer may not care. But "I wrote an even rhythm and it
does not sound even" must never again be something the author has to diagnose
from the audio. This is the most direct answer in the whole ADR to the
original complaint about lacking control over timing and rhythm.

The suggestion must be computed, not hand-waved: given the meter and the
smallest subdivision used, the nearest `speed N` making that subdivision
integral is arithmetic the compiler can do exactly.

**Implementation note.** The `speed 8` and `speed 6` in the example above are
illustrative and not consistent with a single rule. The implementation
suggests the exact multiple of the subdivision's denominator nearest to the
current frames per beat, which gives `speed 24` for the 1/16-at-tempo-120 case
and `speed 23` for the tempo-132 case. Substituting the suggested `speed`
produces exactly even onset spacing. The grid report's frame arithmetic
mirrors the compiler's own `makeCycleClock` formula and drive rate, and the
smallest-subdivision scan lives next to the beat summing in `miniLength`.
None of these diagnostics changes an emitted frame.

### D16. Deferred but recorded, not dropped

Everything below is a real idea from the notation draft that this ADR does
not implement in 4.0.0. Each is recorded with the reason and what it would take,
so it is a decision rather than an omission. None should be quietly lost.

- **`vib ... delay N`**. `vibTri` already takes a delay
  (GoatTracker's own `vibdelay`); the sine-based `vib` does not
  (revisited after D10 landed). D10's per-table `delay` did
  make the right shape obvious: gate `vibOffsetSemitones` on
  `runtime.framesSinceOnset >= vibCfg.delay`, exactly the `vibTriCfg.delay`
  check three lines below it already does. Still deferred to 4.1, though,
  because it is not a one-parameter gap in the SPELLING - `vib` is
  currently `vib RATE DEPTH`, and the grammar/AST/instrument.js changes to
  add a third optional field need to land in JS and Go lockstep with their
  own conformance fixtures, the same cost as any other new field, not the
  free byproduct D10 landing might have made it. Recorded so it stays a
  decision, not a rediscovered gap.
- **Instrument sequence restart mode**. Tables restart on
  every note-on, because they are sampled by `framesSinceOnset`. The notation
  draft asks for `restart` / `continue` / `start at position` per sequence.
  `continue` (a free-running table that ignores note boundaries) is the
  genuinely useful one: it is how continuous PWM shimmer and slow filter
  sweeps that span a phrase are written. Not hard, but it needs a second
  frame counter per voice that does not reset on onset, which is an engine
  change rather than a spelling. Strong candidate for 4.1.
- **Combined wave and pitch table rows**
  (`noise +24 / noise +12 / triangle -12`). Expressible today as parallel
  `wave` and `arp` tables that advance together, so this is readability only.
  It becomes more attractive once D10's per-table `speed` exists, because two
  tables at different speeds no longer line up and a combined row would then
  be ambiguous. Decide after D10 lands, not before.
  **Revisited after D10: still deferred.** D10 landing confirms the ambiguity
  concern is real (independent per-table `speed`/`delay` on `wave` and `arp`
  now genuinely can drift out of phase with each other), which is an
  argument for eventually adding a combined-row spelling that keeps two
  tracks locked together by construction - but it is a new syntax and a new
  AST shape, not a fix to something D10 broke (today's parallel-table
  spelling still works exactly as before), so it stays out of 4.0.0 and is
  a candidate for the same 4.1 pass as instrument sequence restart mode.
- **Relative pitch slide `slide +12 speed 2`**. Distinct
  from portamento-to-a-target (D12): a bend by a fixed interval, which is a
  fall or a rise effect rather than a glide between notes. Cheap once D12's
  mechanism question is answered.
- **`jump` as a table verb**. `loop` and `stop` cover the
  cases that occur in real tracker data. `jump` is deferred until something
  needs it, to keep D13's verb set small enough to memorize.
- **Naming note, not a gap**: the notation draft's `pitch` table is this
  DSL's `arp`. Both are per-frame semitone offsets from the note; the draft
  splits them only by intent (musical intervals versus drum drops and fine
  detune). `arp` serves both, and the worked example's kick uses it that way.
  The docs must say this explicitly, or a reader expecting a `pitch` table
  will hunt for a field that does not exist.
- **Migration aid.** Mechanical `@n` to `/N` rewriting is exact whenever every
  weight in a pattern maps onto a representable duration, which ADR 0001's
  evidence (all weights multiples of 0.25) suggests is the common case. If
  built, it belongs in a `scripts/` tool, not in compiler behaviour.

---

## 4. Non-goals and deferrals

Evaluated on the merits, not dismissed.

- **Indentation-sensitive blocks.** The notation draft is indentation-based
  throughout. The current grammar rejects indentation significance
  deliberately (`dsl-grammar-spec.md`, "Lexical rules"). For a live editor
  where lines are constantly reordered, and for generated source, whitespace
  significance adds a failure mode for no expressive gain. D13's table blocks
  use the existing `end` terminator.
- **Positional voice assignment** (`$ bass:`, declaration order picks the SID
  voice). Voice index is not neutral on a SID: ring and sync couple to the
  neighbouring oscillator and voice 3 has its own filter and mute role.
  Silently reassigning on reorder is a real hazard. `voice 0` stays explicit.
- **`ring from N` / `sync from N`**. SID topology is
  fixed: voice N is modulated by voice N-1, always. A `from` clause would
  imply a choice the hardware does not offer. Rejected as actively misleading.
- **Removing mini-notation and the transform pipeline.** See 1.5.
- **Digi, samples, multi-SID.** Out of scope for this ADR.

---

## 5. Implementation notes

Findings from reading the code that shaped the work.

### 5.1 D1 needs almost no engine work

A duration is a weight. `miniNotation.js` already routes any sequence
containing `@n` through `timeCatWeighted`, which handles unequal per-slot
shares. So `"c4/4 d4/8 e4/8"` in 4/4 is exactly `"c4@1 d4@0.5 e4@0.5"` plus a
computed metric length of 2 beats. The work is: parse `/N[.t]` into a weight
and surface the summed beat count upward. The pattern engine's timing code is
untouched.

### 5.2 The only new engine primitive is D4's timeline

Everything else composes from existing pieces. The timeline is needed because
concatenating differently-sized blocks cannot be expressed as a transform over
one global cycle clock (1.2).

Full definition in cycle units (bar = 1 cycle), for entries
`[{ lengthCycles, pattern }]` where each `pattern` is content filling one
cycle on its own:

```
total    = sum(lengthCycles)
offset_i = sum(lengthCycles[0..i-1])
element i occupies absolute cycles [offset_i, offset_i + lengthCycles_i)
   and plays its content once across that span, i.e. pattern_i.slow(lengthCycles_i)
the whole repeats with period `total`
```

`query(c)` returns the haps falling in the window `[c, c+1)` mapped into that
timeline modulo `total`, with `startFrac`/`durFrac` relative to cycle `c`.
Constraints that are not optional:

- **A Hap never straddles a cycle boundary** (`query.js` header). Clip.
- **`sourceStart`/`sourceEnd` provenance must survive** the shift and slow
  untouched. Editor note highlighting and ADR 0001's bar grid both depend
  on it.
- `total` may be fractional, so the window can wrap the period and can span
  several elements. Split `[c, c+1)` at element boundaries and at the period
  boundary, then query each sub-interval.
- Negative `c` must work (`startOffsetFrames` lead-in walks backwards, see
  `makeCycleClock`'s backward branch).
- `queryAt` comes free from the `Pattern` base class.

D4's `*n` and `[...]` grouping are flattened into the entry list before the
timeline is built. They are not runtime constructs.

The primitive lives in the pattern engine as `timeline(entries)`
(`src/pattern/timeline.js` in `rasterklang-score`). It splits `[c, c+1)` at
period boundaries first (looping when `total < 1`, so a window can span
several periods), then at entry boundaries, and clips each entry through the
existing `queryCycleRange` primitive, so provenance survives the same way it
survives `fast()`/`slow()`. Its tests cover: a single entry of length 1 is
identical to the pattern queried directly; two length-1 entries concatenate
with period 2; entries of length 1 and 2 land at the right offsets; a
fractional total (0.5 + 1) wraps correctly; negative cycle indices resolve to
the right element; provenance survives on every hap; no hap straddles a cycle
boundary.

### 5.3 Beat arithmetic already has two implementations and a shared golden

`src/dsl/miniLength.js` and `pkg/dsl/minilength.go` are a JS/Go pair
computing how many beats a mini-notation string's contents add up to, checked
against each other by `test/conformance/mini-length.cases.json`. Duration
parsing belongs there, in both, in the same change.

`miniNotation.js` must not become a third, independent implementation: it
either calls the same rules or keeps its scanner in lockstep and says so in a
comment, as `miniLength.js`'s header already demands ("Any divergence between
the two is a bug in this file").

Any change to the grammar, the AST, or beat arithmetic lands in the JS and Go
sides together, with regenerated parsers and conformance goldens. New lexer
tokens (`[`/`]` for D4, `chip` for D15) also need the hand-written
token-replay layers (`lexer.js`/`antlrSupport.js`, `pkg/dsl/lexer.go`/
`parser.go`) updated; ANTLR generation gives no warning when that is missing,
and the symptom is a "no viable alternative" parse error on real input.

### 5.4 The part-B fields are wiring, not new mechanisms

Each of D9 to D13 has an existing path to build on:

- `expandRamp` (`interpret.js`) already generates a per-frame array;
  `pingpong`/`loop` extend that one function, not the runtime. The existing
  clamp stays the default.
- **`arpeggioFrameRate` is not touched by D10.** It is the chord round-robin
  rate (`compile.js:577`, used only when more than one Hap is simultaneously
  active), not the instrument arp table's speed. D10's speed and delay belong
  in `sampleTable` (`compile.js:216`), next to the `loopIndex` parameter it
  already takes, and are resolved per table exactly the way
  `resolveLoopIndex` (`:242`) already resolves loop points.
- `toneportaLegato` (`compile.js:530-540`) is the legato mechanism and is
  verified: it suppresses the onset, keeps `framesSinceOnset` counting, emits
  no forced gate pair. `gate legato` is wired to it.
- `forcedControl`'s gate-off-then-on on an onset (`:941`) is **not** a
  retrigger mode. Its own comment states it is invisible to anything sampling
  registers once per frame. Frame-observable re-attack comes from
  `gateReleaseFrames` (the `gate -N` path), which is what both `hardRestart`
  and `gate retrigger` must use.
- D13's verbs expand to the same arrays the bare rows produce, at interpret
  time. Nothing reaches the compiler that does not reach it today.

---

## 6. Compatibility

4.0.0, and it is a real major version.

Breaks:

1. `/n` as mini-notation slow is removed (D1). No known users; `*0.5` and
   `|> slow` are exact replacements.
2. `$name` becomes a resolved reference rather than text substitution (D5).
   The only known reliance on the old behaviour is the `$name!n`
   mis-application, which was a defect.
3. Patterns that silently auto-fit now report it (D8), and constructs that
   silently dropped a declared length now honour it (D5).
4. Tools that generate DSL source need updating to emit durations and
   arrangements. Accepted by decision (see the priority statement).

Additive: `meter`, `each`, arrangements, `|`, durations, dotted, triplets,
`pingpong`/`loop`, `arp speed`/`delay`, `gate N%`, `gate legato`,
`gate retrigger`, `hardRestart`, `slideTo`, table verbs.

Unchanged: `beats N`, `@n`, `<...>`, `{...}` polymetre, `~`/`-` rest, `_`
hold, `speed`, `target`, `reg`, and every transform.

---

## 7. Verification

- Cross-language conformance suite green, JS and Go against shared goldens.
- Compiler and pattern-engine test suites green.
- F1, F2 and F3 all produce correct frames under the onset probe described in
  1.1. A timing or register claim proven by anything other than
  frame-by-frame comparison does not count (ADR 0001 section 8 lists how
  aggregate metrics mislead).
- Every part-B decision proven at register level, not by listening.
- The editor's example song is written in the new notation and a reader who
  has not seen the DSL before can tell what it plays.

---

## 8. Worked example

The target, end to end.

```
meter 4/4
tempo 125

instrument bass
  wave pulse
  adsr 0 6 12 3
  pw
    set 400
    add 20 repeat 8
    add -20 repeat 8
    loop
  end
  gate 80%
  hardRestart 2
end

instrument lead
  wave pulse
  adsr 2 8 10 4
  pw ramp 600..900 step 8 pingpong
  arp 0 4 7 speed 1 delay 3
  vib 0.3 6
  gate legato
end

instrument kick
  adsr 0 4 0 2
  wave triangle
  arp
    set 36
    add -12 repeat 3
    stop
  end
end

pattern riffA = "c2/4 c2/4 g1/4 c2/4"
pattern riffB = "c2/8 c2/8 g1/4 c2/8 e2/8 f2/4 | c2/4 c2/8 g1/8 c2/2"
pattern fill  = "c3/16 d3/16 e3/16 f3/16 g3/8 f3/8 e3/4"

pattern verse  = $riffA*2 $riffB $riffA
pattern chorus = $riffB*2 $fill $riffA

pattern leadA = "a4/8 e5/8 f4/4 g4/8 a4/8 e4/4"
pattern leadB = "c5/2 _/4 g4/8 a4/8 bb4/4"

pattern kickP each 1/8 = "c2 ~ ~ ~ c2 ~ c2 ~"

voice 0 bass  $verse $chorus $verse
voice 1 lead  $leadA $leadB $leadA
voice 2 kick  $kickP*8

filter
  mode lowpass
  cutoff ramp 300..1800 step 12 pingpong
  resonance 10
  route 0
end
```

What this gives a composer that 3.0.0 does not:

- `voice 0` sequences song parts of different lengths. In 3.0.0 that is a
  hard error, and it is the most basic thing a song needs.
- Every pattern's length is the sum of its written durations. Nothing to
  declare, nothing to keep in sync.
- The `|` in `$riffB` is checked. A dropped note is a compile error at the bar
  it broke, not a rhythm that sounds slightly wrong.
- `$verse` and `$chorus` are song parts built from riffs, using the same
  `pattern` construct, with no section system.
- The lead arpeggiates at its own speed with its own delay while the bass does
  not arpeggiate at all.
- `gate legato` makes the lead a melody line rather than a row of separate
  notes; `hardRestart 2` gives the bass its attack.
- The pulse table says `add 20 repeat 8` instead of eight numbers, and the
  filter sweeps forever instead of saturating.
- Every transform still works: `voice 1 lead $leadA |> every 4 rev` is
  unchanged from today.
