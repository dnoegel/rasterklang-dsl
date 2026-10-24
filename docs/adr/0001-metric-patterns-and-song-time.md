# ADR 0001: Metric patterns, bars instead of cycles, and frame-native tempo

Status: implemented (3.0.0)
Date: 2026-08-01
Affects: `rasterklang-dsl`, `rasterklang-score` (compiler and pattern engine), and tools that emit DSL
Current DSL version: 2.19.0. This ADR targets 3.0.0 (see Compatibility).

---

## 1. Context

DSL v2's note notation is inherited from Strudel/TidalCycles: a pattern is a
*proportional* object that subdivides its own cycle by its contents.
`"c4 d4 e4"` divides its cycle into thirds; `"c3 ~ g2 ~"` divides it into
quarters. Two patterns line up only because both happen to be one cycle long.

That model is correct for grooves, which is what Strudel is for. It is the
wrong primitive for melodies, which is most of what C64 music is. The
following problems were all observed while writing real songs, and all trace
back to the same root cause: **there is no shared metric.**

### 1.1 Measured evidence

Everything below was measured on the real implementation, not inferred.

**Auto-fit silently rescales.** At `tempo 132 beats 8` (182 frames per cycle):

```
"c4 ~ c4 ~"   -> each note 1.93 beats   (4 tokens stretched over 8 beats)
32 notes      -> each note 0.18 beats   (32 tokens squeezed into 8 beats)
```

Neither is an error. Both patterns are "one cycle long". An author writing a
4-note bass riff gets 2-beat notes; an author writing a 32-note melody gets a
melody playing 4x too fast.

**`beats` becomes a composition lever by accident.** Because pattern duration
is `beats`, the only way to stretch a melody was to inflate the time
signature. Real songs ended up at `tempo 132 beats 28` this way.

**Length errors are invisible.** One melody summed to 31 beats where 32
was intended. Nothing reported it; auto-fit simply squeezed 31 beats into one
cycle. In a tracker a 31-row column next to a 32-row column is visible
immediately.

**Duration edits cascade (the domino effect).** Because position is the sum of
predecessors, changing one note's length moves every later note:

```
change c4@1.25 -> c4@1.5 in a 12-token melody:
  note  before  after
  c4      0.00   0.00
  d4      1.50   1.75   <- moved
  e4      2.00   2.25   <- moved
  f4      3.00   3.25   <- moved
  ...
  total   7.50   7.75   <- pattern no longer fits the bass
```

This makes per-note tuning ("is @1 or @2 better here?") structurally
expensive, which is the single biggest obstacle to writing tight melodies.

**The material is already metric.** The same melody had every weight a
multiple of 0.25, i.e. a clean 1/16 grid, 32 steps per bar. The fractional
weights were not free-form timing; they were grid positions expressed as
proportions. The syncopated note starting on step 31 and tying across the bar
line is ordinary music, not noise.

### 1.2 What already works and must be preserved

`<...>` alternation already **is** a per-voice orderlist with independent
lengths per voice, which is the GoatTracker model. Verified:

```
voice 0 i "<$bassA-two-elements>"     -> period 2 bars
voice 1 i "<...eight elements...>"    -> period 8 bars, phases independently
```

This is what lets a GoatTracker converter represent songs at all. Nothing in
this ADR may break it.

---

## 2. Decisions

### D1. Patterns may declare a length in beats; that declaration is the opt-in to metric behaviour

```
pattern bass 4  = "c4 ~ c4 ~"     ; metric: 4 beats, one token per beat
pattern lead 32 = "..."           ; metric: 32 beats
pattern hats    = "bd*4"          ; no number -> proportional, exactly as today
```

The number does three jobs at once, which is why it is preferred over a
keyword:

1. **Opt-in.** No number means today's behaviour, so every existing song keeps
   working unchanged.
2. **Length check.** Declared 32, contents sum to 31 -> compile error naming
   the difference.
3. **Polymetre declaration.** Writing `3` means the author meant 3 (see D5).

In a metric pattern, an unweighted token is **one beat**. `@n` remains
available and is then read as "n beats" rather than "n proportional units".

Rationale for putting length on the pattern rather than the voice: a pattern
is the thing that has a length; a voice is a place where one is played. This
also makes inline strings in a voice (`voice 0 i "c4 ~ c4 ~"`) anonymous
patterns that stay proportional, which keeps the simple case simple and
nudges tight melodies toward being named.

### D2. `beats` becomes a pure time signature

`beats N` declares the bar length and nothing else. It no longer determines
how long a pattern is.

Measured today for reference (`frames_per_cycle = round(3000 / tempo * beats)`):

```
tempo 120 beats  8  -> 200 frames, beat = 0.50 s
tempo 120 beats  4  -> 100 frames, beat = 0.50 s
tempo 120 beats 16  -> 400 frames, beat = 0.50 s
```

`beats` never changed beat duration; it only changed cycle length. After this
ADR it only changes bar length, which is what its name implies.

### D3. "Cycle" is renamed to "bar", with a meaning shift

`cycle` is a Strudel implementation concept that leaked into the user model.
Composers think in bars. After this ADR:

- A **bar** is `beats` beats long.
- A proportional pattern fills exactly one bar (identical to today's
  behaviour, new name).
- A metric pattern occupies its declared length and loops at that length,
  spanning or subdividing bars as needed.
- Transforms that counted cycles now count bars. `every 4 rev` means "every
  four bars".

Bar counting stays **global**, not per-pattern. Two voices with `every 4` must
fire together even when their pattern lengths differ; that global coordination
is a genuine strength of the Strudel model and is not to be traded away for
per-pattern repetition counting.

For proportional patterns this rename is purely terminological: cycle length
already equals `beats` beats. Semantics do not change for existing songs.

### D4. A frame-native tempo spelling is added

```
tempo 132 beats 8      ; musical, for composing
speed 6   beats 8      ; 6 frames per beat, exact, for imports
```

BPM is lossy for frame-exact music. `tempo 132 beats 8` computes to 181.8
frames per bar and is rounded to 182. GoatTracker natively counts frames per
row, and not every frame count maps to a round BPM. `speed` removes an entire
class of rounding-induced timing error from the import path.

`tempo` and `speed` are mutually exclusive; exactly one must be present.

### D5. Polymetre is legal; only declaration mismatches are errors

- Contents disagreeing with a pattern's **declared** length: **error**.
- Different **voices** having different lengths: **not an error**. This is
  polymetre and is a legitimate technique that trackers support by design.

Instead of warning, the tooling reports the arrangement's true period:

```
Arrangement: bass 4 beats, lead 32 beats, hats 3 beats
Full period: 96 beats (12 bars)
```

That is the number a composer wants. A soft hint (not an error) is
appropriate when the LCM is implausibly large, since that usually indicates a
typo rather than intent.

### D6. A hold token is added to mini-notation

Trackers have three per-row states: trigger, hold (empty row), release. Mini
notation has only two: note and rest (`~`). Verified today, one 96-frame bar
of 8 beats:

```
"c4@8"                 -> gate on 96/96 frames   (note rings the whole bar)
"c4 ~!7"               -> gate on 24/96 frames   (rests cut the note)
"c4 _ _ _ _ _ _ _"     -> parse error            (no such token)
```

A rest is not a hold. Padding with `~` silences the note, which is why the
importer must express sustain as weights. Add `_` (the mini-notation family's
elongation token; `.` is a reasonable alternative if `_` conflicts) so a voice
can read like a tracker column:

```
notes "c4 _ _ _ _ _ _ _"     ; note, then 7 beats of ring
notes "c4 _ ~ ~"             ; ring 2 beats, then release
```

This is also a prerequisite for splitting a melody into per-bar patterns: the
example melody's syncopation across the bar line cannot be cut without it.
Splitting into short patterns is what bounds the domino effect in D1, so this
token is load-bearing, not cosmetic.

### D7. The metric grid is a rendering, not syntax

Do **not** encode alignment in source whitespace, column layout, or absolute
step indices. All three were considered and rejected:

- Column layout in text is fragile (one space shifts everything) and hostile
  to diffs.
- Absolute indices trade the duration domino for an insertion domino: adding
  a note means renumbering.

An editor already has the second dimension: it can draw bar rulers, column
guides and a playhead over clean source. The provenance needed for this
already exists: `miniNotation.js` attaches `sourceStart`/`sourceEnd` to every
Hap, the lexer records `docStart`/`docEnd` on strings, and `song.js` maintains
`noteSourceSegments`, so an editor can answer "which source span produced this
sounding note".

A tracker **view** therefore falls out of existing infrastructure. Editing
through that view is out of scope here (see Deferred).

---

## 3. Explicitly deferred

- **Editable tracker view.** Read-only is cheap and unambiguous. Editing is a
  view-update problem: transforms that permute or annotate events (`rev`,
  `add`, `slow`, `iter`, `rot`) preserve provenance, but transforms that
  create or delete events (`ply`, `degradeBy`, `jux`, `superimpose`) do not.
  A sane later step is "editable only where the mapping is unambiguous, with
  a visible reason when it is not".
- **Bar lines inside strings** (`"c3 ~ g2 ~ | a2 ~ f2 ~"`). Cheap and useful
  for non-syncopated material, but they cannot be placed where a note ties
  across the bar, which is exactly the interesting case. D1's declared length
  catches the same class of error without that limitation.
- **Transforms on `pattern` declarations.** Currently `pattern p = "..." |> slow 4`
  is a parse error, and authors reach for it because the transform feels like a
  property of the melody. Keeping transforms on voices preserves `$name`
  splicing as pure text substitution. Revisit only with a concrete design for
  how a transformed pattern splices.

---

## 4. Implementation scope

- **`rasterklang-dsl`**: the grammar (`grammar/RasterklangDsl.g4`) gains the
  pattern length in `patternDecl`, `speed` alongside `tempo`, and the hold
  token in mini-notation. The Go and JS AST-transform visitors
  (`pkg/dsl/ast_visitor.go`, `src/dsl/astVisitor.js`) stay in lockstep, and
  the conformance fixtures (`test/conformance/`) cover metric patterns, the
  hold token, `speed`, declared-length mismatches and polymetre. Length
  validation and its error message live in the interpreter.
  `docs/dsl-grammar-spec.md` gains a prose section on song time (bars,
  pattern length, metric vs proportional, orderlists via `<...>`); its
  absence was the root cause of most of the confusion documented in
  section 1.
- **Compiler and pattern engine** (now `rasterklang-score`): metric pattern
  timing, bar-based transform counting (`every`/`slow`/`fast`), and
  hold-token gate semantics. The compiler already has non-trivial gate
  handling (`gate FRACTION` and `gate -N`); the hold token interacts with it
  and must not regress either. Tests prove (a) a metric pattern's notes land
  on declared beats, (b) two voices of different metric lengths phase
  correctly and realign at the LCM, (c) the hold token rings rather than
  silences.
- **Converters**: GoatTracker data is already metric and was being flattened
  into proportions, so a converter benefits most from emitting declared
  pattern lengths, `speed` instead of a computed-and-rounded BPM, and hold
  tokens instead of weight chains.

---

## 5. Compatibility and migration

- Patterns without a declared length keep today's proportional semantics
  exactly. Existing songs are unaffected by D1.
- D3 is terminological for existing songs: a proportional pattern already
  fills one cycle, and cycle length already equals `beats` beats. Renaming to
  "bar" changes documentation, error text and the mental model, not output.
- D4 is additive.
- D6 is additive (`_` is currently a parse error, so no existing program uses
  it).
- The version bump is nevertheless **3.0.0**, because `beats` losing its
  effect on pattern length is a real semantic change for anyone who used it as
  a stretch lever, and because "cycle" disappearing from the vocabulary is a
  breaking documentation change.

Migration aid worth building: a checker that flags songs whose `beats` value
looks like a stretch lever (unusual values such as 28) and suggests the
metric spelling.

---

## 6. Known defects to fix alongside

Found while investigating this ADR; each is small, real, and in scope because
each is a symptom of the same confusion.

1. **Weighted rests do not close the gate.** `"c4@4 ~@4"` measured gate-on
   96/96 frames, while `"c4 ~!7"` measured 24/96. A weighted rest should
   behave like a plain rest.
2. **`$name!n` silently misapplies.** `<$bassA!3 $bassB>` applies `!3` to the
   pattern's last token, because `$name` splices raw text; measured output was
   `36 36 43 43 43 43 41 41` instead of the intended `36 36 36 41 ...`.
   `[$bassA]!3` is correct. Emit a warning suggesting the bracketed form.
3. **Unhelpful error for `|>` on a pattern declaration.** Produces
   `line 9: mismatched input '' expecting NEWLINE`, ANTLR's default. Should
   say that pipeline transforms belong on a voice.
4. **Filter cutoff above 2047 is silently clipped.** `cutoff sine 900..2600`
   clips to 2047 with no diagnostic; cutoff is 11-bit.

---

## 7. Implementation decisions

Design questions this ADR left open, and implementation choices not spelled
out above, with the resolution that shipped in 3.0.0. Numbering is stable and
referenced from code comments (e.g. "decision 10").

1. **Alternation of metric elements.** In `<[$a] [$b]>`, each element used to
   take one cycle. Resolved as orderlist semantics: each element of `<...>`
   occupies its own declared length rather than one bar each. This matches
   trackers and the way GoatTracker conversions already rely on alternation
   for per-voice orderlists (section 1.2), with no change to that
   representation.
2. **Hold token spelling.** `_`. No lexer conflict: it is not used elsewhere
   in mini-notation. `.` was rejected because it is visually close to `~` at
   small sizes and has no family precedent as strong as Tidal's `_`
   (elongate the previous event).
3. **Metric pattern shorter than a bar.** It loops at its own declared
   length, exactly like a GoatTracker orderlist entry shorter than the
   arrangement's bar count. A 4-beat pattern repeats twice in an 8-beat bar.
4. **Fractional `@n` in metric patterns.** Allowing `@0.25` keeps
   expressiveness but reintroduces the arithmetic the model is meant to
   remove. Decimal fractions remain legal and are read as fractional beats,
   same arithmetic as before. A musical spelling (`1/16`, dotted, triplets)
   is the better long-term answer and was deferred as additive sugar over the
   same beat count (later delivered by ADR 0003 D1).
5. **Declared length type.** The number in `pattern name N = "..."` accepts
   both integers and decimals (`pattern lead 7.5 = "..."` is legal), since
   the length check compares against the summed content weight, which is
   already fractional.
6. **Weighted-rest gate defect.** Fixed by making rest gate behaviour
   independent of weight: a rest (weighted or not) always closes the gate
   for its own span. `"c4@4 ~@4"` now releases at the same point as the
   equivalent run of plain rests.
7. **`$name!n` splice defect.** Behaviour is unchanged (still applies `!n` to
   the pattern's last token when unbracketed), but the compiler emits a
   warning when it detects `$name!n` outside brackets, suggesting the
   `[$name]!n` form. Diagnostic-only; no silent behaviour change for existing
   songs. (ADR 0003 D5 later made `$name` a resolved reference, which removes
   the hazard entirely.)
8. **`|>` on a pattern declaration.** The raw ANTLR mismatch error is replaced
   by a semantic check that reports: "pipeline transforms ('|>') are not
   allowed on pattern declarations; add the transform on the voice that
   plays this pattern instead."
9. **Cutoff-over-2047 diagnostic.** Compiling a `cutoff` literal or
   `sine`/`saw`/etc. range endpoint above 2047 emits a warning
   ("cutoff N clipped to 2047 (11-bit filter cutoff)") instead of silently
   clipping. The clipped value is still used at runtime; diagnostic-only.
10. **Migration checker heuristic.** `beats` values outside the common set
    `{1, 2, 3, 4, 6, 8, 9, 12, 16}` are flagged with a hint: "beats N is
    unusual for a time signature; if you meant to stretch a melody, give
    the pattern a declared length instead (`pattern name N = ...`)."
    Warning-level, not an error, since unusual time signatures are legal
    music (e.g. 5, 7).
11. **`speed`/`tempo` exclusivity.** Enforced as a semantic check (both
    absent or both present is an error), not a grammar-level XOR, so the
    error message can name which case applies.
12. **"Cycle" to "bar" rename scope.** Applied to grammar-adjacent
    identifiers, error messages, docs and comments in code touched by this
    ADR. Internal variable names deep in the pattern engine that never
    surface to users (e.g. Strudel-derived internals) were left alone where
    renaming them would be churn with no user-facing effect.

---

## 8. Verification

- Cross-language conformance suite green (JS and Go against shared goldens).
- Compiler and pattern-engine test suites green.
- Converter output does not regress against its existing structural sample.
- New behaviour is proven at the register level, not by aggregate score.
  Aggregate metrics are easy to fool: pitch RMS computed only over matched
  onsets can score a flattened 6-octave sweep as perfect; byte-match rewards
  register stasis; waveform match can sample exactly the frame where a hard
  restart deliberately has no waveform bits. Timing claims are proven by
  comparing gate onsets and frequency registers frame by frame.
