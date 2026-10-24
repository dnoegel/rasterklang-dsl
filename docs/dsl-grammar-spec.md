# rasterklang DSL v2 - grammar specification

Status: **live**. The current version is `src/dsl/version.js`'s `DSL_VERSION`; see
`docs/dsl-changelog.md` for what changed at each version and `docs/adr/` for the larger design
decisions. This is the narrative syntax and semantics reference for the language.

**Where the facts live.** Several places encode grammar facts, kept honest in different ways:

1. **`grammar/RasterklangDsl.g4`** is the lexical/syntactic grammar both the JS and Go parsers are
   generated from, and **`src/dsl/astVisitor.js`** (with its Go counterpart
   `pkg/dsl/ast_visitor.go`) turns the parse tree into the AST and performs the semantic checks.
   Together they are the executable ground truth: if this doc and the implementation ever
   disagree, the implementation is right and this doc is stale. The five canonical name lists
   (`STATEMENT_KEYWORDS`/`INSTRUMENT_FIELDS`/`VOICE_FIELDS`/`TRANSFORMS`/`FILTER_FIELDS`) live in
   `astVisitor.js`, next to the switch statements that validate them, so a change to one is hard
   to miss making to the other. Keeping a single copy is a lesson learned: an earlier, separately
   hardcoded copy went stale (missing `wave`) without any test noticing (see `2.7.0`'s changelog
   entry).
2. **`src/dsl/keywords.js`** is the machine-readable *name* index for tooling (syntax highlighting,
   autocomplete). Every list in it is a re-export of those canonical arrays (or, for register and
   waveform names, derived from the register map's real objects), not a second hand-typed copy.
3. **`src/dsl/fieldDocs.js`** is the hand-written *prose*: a signature, a one-line description, and
   a parseable example for every name `keywords.js` lists.
4. **`docs/dsl-spec.generated.json`** is the generated, versioned, machine-readable spec (run
   `npm run generate:dsl-spec` to regenerate it via `scripts/generate-dsl-spec.mjs`): every
   name/signature/description/example from (2) and (3), tagged with the current `DSL_VERSION`.
   `test/dsl/generatedSpec.test.js` fails if the committed file doesn't match what the generator
   would produce. It is the file to point a script at for "the complete current DSL vocabulary",
   but it is generated *from* (2) and (3), so a wrong description has to be fixed in
   `fieldDocs.js`.
5. **This document** is the human-facing *narrative*: the EBNF, the semantics, the worked examples,
   and the *why* behind each design choice. It is manually maintained and not guarded by any test
   against the others, so treat it as the best available narrative spec, not an infallible one.

**The headline change from v1**: the source text is no longer evaluated as JavaScript (see "Why
not JS" below). It's a real, small, line-oriented grammar with its own lexer/parser, compiling to
the *same* internal `Song` builder object the compiler (`rasterklang-score`'s `compile.js`)
already consumes, so the pattern engine (`query.js`, `combinators.js`, `signals.js`,
`miniNotation.js`) and the SID compiler (`compile.js`, `registerMap.js`) needed **no changes** for
anything already expressible in v1's JS API. Only genuinely new capabilities (instrument tables,
vibrato/slide, test-bit, raw register pokes, tempo-as-pattern) needed engine-level additions.

## Why not JS

v1 ran the editor buffer through `new Function(...PARAM_NAMES, sourceText)`: arbitrary JS eval
with some names injected. Two concrete reasons this DSL replaces that instead of keeping it as an
option:

1. **Safety**: a real grammar can't reach `fetch`/`window`/infinite loops; a JS eval path always
   can, no matter how "trusted."
2. **It doesn't read like a music language.** `song(125).voice(0, note('c4').s('saw'))` reads as
   "calling a JS API," not as domain notation the way GoatTracker/Alda/MML text does. A bounded
   grammar with a small vocabulary is also far easier to learn, generate and validate than "valid
   arbitrary JS that happens to call our functions correctly."

This is a **full replacement**, not a dual-mode system - see "What happens to the old JS builder
API" at the end of this doc.

## Lexical rules

- **Encoding**: UTF-8 source text.
- **Comments**: `;` to end of line (GoatTracker/tracker-culture convention).
- **Whitespace**: spaces/tabs separate tokens within a line; otherwise insignificant. Newlines
  terminate a statement. **No indentation sensitivity anywhere in this grammar** - deliberate,
  to keep the parser simple and unambiguous; blocks are delimited by explicit `end` keywords,
  never by indent level (see "Blocks" below).
- **Identifiers**: `[a-zA-Z_][a-zA-Z0-9_]*` - used for instrument names, pattern names, keywords.
- **Numbers**: `-?[0-9]+` (int) or `-?[0-9]+\.[0-9]+` (float). Hex literals (register addresses
  only): `0x` followed by 1-2 hex digits, e.g. `0x15`.
- **Strings**: two delimiters, one token type. `"..."` must close on the same physical line.
  `` `...` `` (backtick) is the same string, except it may **span multiple physical lines** -
  useful for a long pattern (a 12-bar melody, say) that's more readable laid out one bar per
  line than crammed onto one very long line. A backtick string's embedded newlines (and any
  indentation) are preserved *literally*, not collapsed to a space - `miniNotation.js`'s own
  tokenizer already treats any whitespace, including a literal `\n`, as an inter-token
  separator, so no normalization is needed for parsing, and the stored source text stays an
  exact substring of the document, which is what keeps the live editor's note-highlighting
  working with no special-casing anywhere downstream. Escapes for both forms: `\"`/`` \` ``
  and `\\` only. A string's contents (either delimiter) are parsed by `miniNotation.js`'s
  **existing, unmodified grammar** (`~`/`-` rest, `[]` subdivide, `[a,b]` stack, `<>`
  alternation, `*n`/`/n` speed, `@n` weight, `!n` replicate, `?p` drop, `a|b` choice, `(p,s,r)`
  euclidean, note names, numbers, bare words), plus `_`, the hold token added in 3.0.0 (see
  "Song time").
- **Named-pattern reference**: `$name` - a sigil-prefixed identifier, resolved at compile time
  (see "Named patterns").

### Quoted vs. bare fields - the one rule to hold onto

- Any field whose value can be a **full mini-notation pattern** (nesting, alternation, euclidean,
  etc.) is a **quoted string**: voice `notes`, `pattern NAME = "..."` declarations, `reg`'s value
  pattern, `test`'s per-frame pattern.
- Any field that's a **flat per-frame table** (no nesting ever needed) is **bare, space-separated
  tokens**: `wave`, `arp`, `pw`, `filterTable` inside an `instrument` block. These are tracker
  tables (see "Instruments" below), not mini-notation patterns - different sampling semantics,
  see the callout in that section.
- Fixed positional scalars (`adsr A D S R`, `vib DEPTH RATE`, `slide TIME`, `gate FRACTION`,
  `pw` as a plain instrument default) are bare numbers.

## Blocks

A block starts with a header line ending in a keyword-appropriate way and is closed by a line
containing only `end`. No block nests more than what's listed below (the grammar is
deliberately flat - at most one level of block, `switch bar` inside `filter`'s `cutoff` field
being the sole two-deep exception).

```
instrument NAME
  ... fields ...
end

voice N [INSTRUMENT_NAME]
  ... fields ...
end

filter
  ... fields ...
end
```

**Short-mode exception**: `voice` also has a single-line form with no `end` - see "Voices"
below. This is what makes "short mode" and "extended mode" the same grammar, not two grammars:
short mode is just the single-line production of `voice_stmt`; extended mode is the
multi-line/`end`-terminated production. Everything else in the language (instruments, filter,
tempo, named patterns) already exists in exactly one form.

## Grammar (EBNF-ish)

```
program        := statement*
statement       := comment | tempo_stmt | target_stmt | pattern_decl | table_decl
                  | instrument_block | voice_stmt | filter_block

comment         := ';' <rest of line>

tempo_stmt      := ('tempo' | 'speed') tempo_value (tempo_setting)*
                                            ; 'tempo' = BPM, 'speed' = FRAMES PER BEAT.
                                            ; Exactly one clock statement per song; having
                                            ; neither or both is a semantic error, not a
                                            ; grammar one, so the message can say which.
tempo_setting   := 'beats' INT | 'offset' NUMBER
                                            ; 'beats' is the TIME SIGNATURE (how long a bar
                                            ; is). It does not set any pattern's length.
tempo_value     := NUMBER | STRING          ; STRING form allows "<120 96>" etc.

target_stmt     := 'target' '{' target_setting* '}'
                                            ; the DRIVE rate: CPU cycles per play call.
                                            ; Optional, at most one per song; absent means
                                            ; PAL 19705 cycles = exactly 50.00 fps.
target_setting  := 'video' ('PAL' | 'NTSC')
                  | 'driveCycles' INT       ; exact, the primitive
                  | 'multispeed' INT        ; play calls per video frame
                                            ; 'driveCycles' and 'multispeed' together is a
                                            ; semantic error (two spellings of one number).

pattern_decl    := 'pattern' IDENT NUMBER? '=' STRING
                                            ; the optional NUMBER is the pattern's declared
                                            ; length IN BEATS, which is what makes it metric
                                            ; (see "Song time"). Decimals allowed. No number
                                            ; = proportional, i.e. it fills one bar.

table_decl      := 'table' IDENT '=' table_row   ; reusable named per-frame table, used as $NAME

instrument_block:= 'instrument' IDENT NEWLINE instrument_field* 'end'
instrument_field:= 'wave' table_row
                  | 'adsr' NUMBER NUMBER NUMBER NUMBER
                  | 'pw' (NUMBER | table_row)
                  | 'arp' table_row
                  | 'filterTable' table_row
                  | 'loop' (INT | loop_pair+)   ; INT = one global default; TABLE INT pairs = per-table
                  | 'vib' NUMBER NUMBER        ; depth (semitones), rate (Hz)
                  | 'vibTri' NUMBER (NUMBER | SHIFT) NUMBER
                                                ; GoatTracker's own triangle vibrato:
                                                ;   compare value, speed (frequency-register
                                                ;   units per frame, or 'sN' = one semitone
                                                ;   at the current note >> N), delay frames
                  | 'slide' NUMBER              ; portamento time, fraction of a beat
                  | 'test' table_row            ; 0/1 per-frame table (test-bit)
                  | 'gate' (NUMBER | table_row) ; NUMBER >= 0 = static fraction of the slot; NUMBER < 0 = gate-off
                                                ;   frames at the note's END (GoatTracker's gatetimer/hard restart);
                                                ;   table_row = per-frame gate-bit program

loop_pair       := ('wave'|'pw'|'arp'|'filterTable'|'test'|'gate') INT

table_row       := table_token+                 ; space-separated bare tokens
table_token     := (NUMBER | WORD ('+' WORD)*) ('!' INT)?
                                                  ; a number, or a waveform name / several
                                                  ; '+'-joined names sounding SIMULTANEOUSLY
                                                  ; in one frame slot (noise+pulse+saw).
                                                  ; '!N' repeats the value N times.
                 | '$' IDENT                    ; a table_decl, spliced in place
                 | ramp_generator

ramp_generator  := 'ramp' NUMBER '..' NUMBER ('step' NUMBER)? ('over' INT)?
                                                  ; linear per-frame sweep, saturating at
                                                  ; the target. At least one of step/over.

voice_stmt      := 'voice' INT IDENT? voice_body
voice_body      := voice_inline_tail             ; short mode, single line, no 'end'
                  | NEWLINE voice_field* 'end'    ; extended mode

; IDENT? (an instrument name) is shared by BOTH productions above, not just
; extended mode - 'voice 0 bass "..."'/'voice 0 bass $pattern' (short mode
; WITH a named instrument) is real syntax, disambiguated from extended mode
; by one extra token of lookahead (IDENT followed by STRING/SIGIL = short
; mode; IDENT followed by NEWLINE = extended mode) - see parser.js's
; parseVoiceStmt(). A named instrument and inline_attrs are mutually
; exclusive on one voice (a compile error, not silently picking one - see
; "Voices" below).
voice_inline_tail := (STRING | SIGIL) inline_attr* pipeline?
inline_attr     := IDENT ':' value_list           ; e.g. wave:pulse  adsr:0,4,10,4
                  | 'wave' ':' STRING              ; e.g. wave:"pulse|triangle" -- see "Waveform" below
value_list      := token (',' token)*

voice_field     := 'notes' STRING
                  | 'insts' STRING                 ; per-note instrument-switching pattern, see "Instruments"
                  | 'reg' reg_target STRING
                  | 'vib' NUMBER NUMBER
                  | 'slide' NUMBER
                  | 'gate' NUMBER                   ; voice-level override stays the single-number form only
                                                    ;   (fraction, or negative = gate-off frames)
                  | 'sync' bool_value
                  | 'ring' bool_value
                  | 'wave' STRING                  ; extended-mode counterpart to wave:"..." above
                  | pipeline

bool_value      := NUMBER | STRING     ; a 0/1 constant, or a quoted mini-notation
                                        ; pattern if the bit should vary over time

pipeline        := '|>' transform (WS '|>' transform)*
transform       := 'rev'
                  | 'fast' arg | 'slow' arg
                  | 'every' arg transform_ref
                  | 'iter' arg
                  | 'ply' arg
                  | 'off' arg transform_ref
                  | 'jux' transform_ref
                  | 'degradeBy' arg
                  | 'sometimesBy' arg transform_ref
                  | 'add' arg
                  | 'struct' STRING              ; boolean rhythm pattern
                  | 'mask' STRING                ; boolean filter pattern
                  | 'chunk' arg transform_ref
                  | 'palindrome'
                  | 'linger' arg
                  | 'rot' arg
                  | 'someCyclesBy' arg transform_ref
                  | 'superimpose' transform_ref  ; claims a second voice, like jux
                  | 'scale' STRING               ; e.g. "c4:minor"

arg             := NUMBER | STRING               ; STRING = a mini-notation pattern sampled
                                                  ; once per bar, e.g. add "<0 0 2>"
transform_ref   := transform | '(' transform ')'   ; parens only needed when a transform
                                                    ; itself takes an argument, e.g. (add 12)

reg_target      := IDENT | HEX                   ; symbolic register name or raw 0xNN address

filter_block    := 'filter' NEWLINE filter_field* 'end'
filter_field    := 'cutoff' filter_expr           ; signal_expr | switch_block only - a bare
                                                    ; number never means anything for cutoff
                  | 'resonance' filter_expr        ; NUMBER | signal_expr | switch_block
                  | 'mode' filter_expr             ; IDENT | switch_block of IDENT cases -
                                                    ; lowpass | bandpass | highpass | none;
                                                    ; signal_expr's range/speed modifiers are
                                                    ; rejected (mode is categorical, not numeric)
                  | 'route' filter_expr            ; INT (',' INT)* | switch_block of INT lists -
                                                    ; signal_expr is rejected (routing is a set
                                                    ; of voices, not a continuous quantity)

filter_expr     := signal_expr | number_list | switch_block
                                                    ; all four filter fields share this one
                                                    ; grammar shape; which of the three is
                                                    ; actually meaningful for a given field is a
                                                    ; semantic check (ast_visitor.go/astVisitor.js),
                                                    ; not a parse-time restriction
signal_expr     := SIGNAL_NAME range? speed*      ; e.g. sine 300..2200 / 4 >> 3 -
                                                    ; speed steps chain, applied in order;
                                                    ; also doubles as a bare-identifier leaf
                                                    ; (no range/speed) for mode's IDENT values
number_list     := NUMBER (',' NUMBER)*            ; resonance's single constant, or route's
                                                    ; list of voice indices
range           := NUMBER '..' NUMBER
speed           := ('*' | '/') NUMBER
                  | '>>' NUMBER                    ; rotate/phase-shift

switch_block    := 'switch' ('bar' | 'cycle') NEWLINE switch_case+
                                            ; 'bar' since 3.0.0; 'cycle' is the old
                                            ; spelling of the same thing, still accepted
                                                    ; no closing 'end' - the case list simply
                                                    ; ends at the first line that isn't
                                                    ; '<'/'>='/'else'-shaped (still token-based,
                                                    ; not indentation-based: an ordinary line
                                                    ; like the next 'resonance' field is what
                                                    ; actually ends it, same as any other
                                                    ; filter_field)
switch_case     := ('<' NUMBER | '>=' NUMBER | 'else') ':' (signal_expr | number_list)
                                                    ; number_list here is what lets route's
                                                    ; per-case values (a voice-index list) be
                                                    ; spelled at all - they can't fit signal_expr's
                                                    ; identifier-led shape
```

(`route`'s INT list accepts either comma- or space-separated values in the implementation -
both `route 0, 1` and `route 0 1` parse identically - same for `number_list` generally.)

`SIGNAL_NAME` is one of `sine saw square tri rand perlin` (unchanged from `signals.js`).

## Semantics

### Song time: bars, pattern length, metric vs proportional

This section is the one to read first. Everything below it assumes the model described here,
and its absence was the direct cause of most of the confusion recorded in
`docs/adr/0001-metric-patterns-and-song-time.md` section 1.

**Two clocks, not one.** "How long is a beat" (musical time: `tempo`/`speed`, `meter`, a
pattern's own declared length) and "how long is a frame" (SID tick time: `target`'s drive rate,
ADR 0002) are independent questions. Musical time is what this whole section is about - beats,
bars, pattern length - and it is expressed entirely in BEATS, never frames, so a song's
melodic content doesn't change shape when its drive rate does. Frame time is what the compiler
converts musical time INTO at the very last step, once per song, using whatever `target` block
(or its PAL default) declared. A `speed N` clock statement is the one place these two clocks
meet directly - "N frames per beat" - which is why it exists at all: it is GoatTracker data's
native unit, and every other frame-shaped question in this spec (table sampling, gate timing)
is answered in frames-since-onset, never in beats.

**A bar is `meter N/D` beats long** (`meter 4/4`, `meter 6/8`) **and that is all `meter` does.**
It is a time signature: N beats per bar, each beat a 1/D note. It says nothing about how long
any particular pattern is - that is what a pattern's own declared length (below) is for. The
older `beats INT` spelling still works (it only ever gave the numerator, always assuming
quarter-note beats) but is mutually exclusive with `meter` on one clock statement. Before
3.0.0, `beats` doubled as every pattern's length, which is why people reached for `beats 28` to
stretch a melody; it no longer has that effect. (The word for a bar used to be "cycle", a
Strudel implementation term that leaked into the user model. Same thing, better name.)

**A pattern is proportional by default and metric if it declares a length.**

```
pattern hats    = "bd*4"          ; proportional: fills exactly one bar, whatever a bar is
pattern bass 4  = "c2 ~ c2 ~"     ; metric: 4 beats, so one token = one beat
pattern lead 32 = "..."           ; metric: 32 beats
```

- **Proportional** (no number) is the pre-3.0.0 reading, unchanged: the contents subdivide one
  bar between them. `"c4 d4 e4"` is three thirds of a bar; 32 tokens are 32 thirty-seconds of a
  bar. Nothing about existing songs changes.
- **Metric** (a number before `=`) means the pattern is that many beats long. An unweighted
  token is exactly **one beat**, and `@n` reads as **n beats** rather than n proportional
  shares. Decimals are legal in both places (`pattern odd 7.5 = "c4@0.5 d4@3 e4@4"`).

The number does three jobs at once, which is why it is a number on the declaration rather than
a keyword somewhere else:

1. **It is the opt-in.** No number, no change.
2. **It is checked.** If the contents add up to something else, that is a compile error naming
   both lengths: `pattern "lead" declares 32 beats but its contents add up to 31 beats`. A
   31-beat melody where 32 was meant used to be silently squeezed to fit.
3. **It declares polymetre deliberately.** Writing `3` means you meant 3 (see below).

**A metric pattern loops at its own length**, spanning or subdividing bars as needed. A 4-beat
pattern in an 8-beat bar repeats twice per bar; a 32-beat pattern in an 8-beat bar takes four
bars to come round. This is what a tracker orderlist entry does, and it is what makes the
declared length a musical statement rather than a formatting detail.

**Bar counting for transforms stays global.** `every 4 rev` means "every four bars", counted on
the song clock, not per pattern. Two voices with `every 4` fire together even when their
pattern lengths differ - that global coordination is a real strength of this engine and is not
traded away for per-pattern repetition counting.

**An inline string on a voice is an anonymous pattern and stays proportional.** `voice 0 "c4 ~
c4 ~"` behaves exactly as it always has. To make a melody metric, name it - which is also the
form you want for anything long enough to care about.

#### Durations: `/N`, and `each 1/D`

A declared length says how long the WHOLE pattern is; `/N` says how long ONE STEP is, directly,
in the pattern's own beat unit:

```
pattern lead = "c4/4 e4/8 g4/8 c5/4."   ; quarter, eighth, eighth, dotted quarter
```

`/N` is `beatUnit/N` beats - a real musical duration, not a speed multiplier (that reading was
retired in 4.0.0; `/N` used to mean "slow this step down by N", which nothing else in the
notation calls a duration). A trailing `.` dots it (x1.5); a trailing `t` makes it a triplet
(x2/3). `/N` and `@n` are two authoring models for the same underlying weight - a duration and
a proportional share - and mixing them in one pattern is an error naming which is which.

`each 1/D` is the step-sequencer's own spelling, on the pattern DECLARATION rather than per
step: every token is exactly one `1/D` note, so the pattern's total length falls out of its own
token count instead of being declared up front.

```
pattern hats each 1/16 = "x ~ x ~ x ~ x ~ x ~ x ~ x ~ x ~"   ; 16 sixteenth notes, 1 bar at 4/4
```

#### Note, slot, and gate are three different things

A **note** is material - a pitch, a duration, an intent to sound something. A **slot** is its
span in the grid - the beats it occupies, whether or not anything actually plays there (a rest
is a slot with no note in it). A **gate** is the SID control-register bit - a hardware fact
about whether the envelope generator is currently attacking/sustaining, that outlives the slot
that triggered it.

The three diverge constantly, and most of what `gate`/`slide`/`slideTo` (see the instrument
fields section) exist to control is exactly that divergence:

- A slot can hold a note whose GATE stays open past the slot's own boundary - `gate legato`
  (D11) and `slideTo` (D12) both mean "the next slot's note changes what's SOUNDING without
  closing the gate that's already open."
- A slot can close the gate before its own boundary ends - `gate 0.8`/`gate -N`/`hardRestart N`
  all mean "release before the slot is over," at three different granularities (a fraction of
  the slot, an absolute frame count from the end, the same thing under a readable name).
- The hold token `_` (below) extends a SLOT, not a note by itself - it is what lets one note's
  material span several slots' worth of grid with the gate never dropping in between.

Keeping these separate is what makes "how long is this note" and "is the envelope retriggering"
answerable independently instead of one fact standing in for both, the way a bar's worth of
slots alone cannot say whether two adjacent same-pitch notes are one sustained sound or two
articulated ones.

#### Arrangements vs. `<...>` alternation

Two different things both look like "play these patterns in sequence," and they answer
different questions:

- **`<a b>` alternation** (below) is a per-VOICE, per-PASS orderlist: one element sounds per
  cycle, cycling. It answers "what does this voice play on pass N."
- **`$a $b` (an arrangement, D4)** is structural sequencing: `$a` plays in full, THEN `$b`
  plays in full, in ONE pass - a song section list, not a rotation.

```
pattern verse  = $a $b          ; an arrangement: play $a then $b, once through
voice 0 "<[$a] [$b]>"           ; alternation: play $a on pass 1, $b on pass 2, repeat
```

An arrangement can repeat (`[$a $b]*2`) and nest (`$intro [$verse $chorus]*2 $outro`), and a
named arrangement is itself usable as one voice's whole `notes` value or nested inside another
arrangement - it maps directly onto a GoatTracker orderlist, which `<...>` alternation never
could (an orderlist plays its whole list once through per song, it does not rotate one entry
per pass).

#### Polymetre and the arrangement period

Different voices having different lengths is **not** an error. It is polymetre, and trackers
support it by design:

```
pattern a 4 = "c2 ~ c2 ~"
pattern b 3 = "e3 f3 g3"
voice 0 $a
voice 1 $b
```

Instead of warning, the compiler reports the arrangement's true period - the LCM of the voice
lengths, the number a composer actually wants:

```
Arrangement: voice 0 4 beats, voice 1 3 beats. Full period: 12 beats (3 bars).
```

A soft hint (not an error) appears when that period is implausibly long, since a 500-bar period
is usually a mistyped declared length rather than a phase piece. What IS an error is a
declared length disagreeing with its own pattern's contents.

#### Orderlists via `<...>`

`<...>` alternation is the per-voice orderlist, and it always was: one element per pass, each
voice cycling independently. With metric elements each element occupies **its own declared
length** (orderlist semantics), which is exactly how a GoatTracker orderlist replays a column
of patterns:

```
pattern verse  8 = "..."
pattern chorus 8 = "..."
voice 0 "<[$verse] [$chorus]>"
```

Note the brackets. `$name` splices raw text, so `<$verse!3 $chorus>` applies the `!3` to the
last token of `verse`, not to `verse` as a whole; `[$verse]!3` is the correct spelling and the
compiler warns when it sees the unbracketed form. Mixing elements of DIFFERENT declared lengths
inside one `<...>` is rejected: an orderlist plays one length per voice, so give the patterns
the same length or split them across voices.

#### The hold token `_`

A tracker row has three states: trigger, hold, release. Mini-notation had two - a note and a
rest - and a rest is not a hold: it CUTS the note. `_` is the third:

```
notes "c4 _ _ _ _ _ _ _"     ; one note, ringing for eight beats
notes "c4 _ ~ ~"             ; ring for two beats, then release
```

Each `_` extends the step before it by one step's worth of time; `_@2` and `_!2` add two at
once. A hold with nothing before it is an error. Holds work after a rest too (`"~ _"` is a
two-beat rest), and a **weighted rest closes the gate for its whole span** exactly like the
equivalent run of plain rests - `"c4@4 ~@4"` releases at the halfway point, it does not ring
through.

Being able to write a held note is what makes it practical to split a long melody into
per-bar patterns: a note tied across a bar line has a spelling now.

#### `speed` vs `tempo`

```
tempo 132 beats 8      ; musical: BPM, for composing
speed 6   beats 8      ; exact: 6 frames per beat, for imports
```

Both set the same clock; a song has exactly one of them (neither or both is an error naming
which case applies). BPM is lossy at frame resolution - `tempo 132 beats 8` computes 181.8
frames per bar and rounds to 182 - while GoatTracker natively counts frames per row and not
every frame count maps to a round BPM. `speed 6 beats 8` is exactly 48 frames per bar. The two
are related exactly by `bpm = 3000 / framesPerBeat` on a 50Hz PAL machine, and `speed` accepts
the same `beats`/`offset` suffixes and the same `"<6 8>"` pattern form as `tempo`.

#### What is deliberately NOT part of this

The metric grid is a rendering, not syntax. Alignment is never encoded in source whitespace,
column layout, or absolute step indices: column layout in text is fragile and hostile to diffs,
and absolute indices trade a duration domino for an insertion domino. The editor draws bar
rulers and a playhead over clean source instead, using the provenance the compiler already
carries.

#### Design rules this section follows

Adapted as a closing summary, not a separate ruleset - every construct above chose its spelling
by these:

1. **A number means what a musician means by it.** A duration is beats (`/N`), a bar is a time
   signature (`meter N/D`), a fraction of a slot is a fraction (`gate 0.8`). Never frames, never
   an index, unless the field is explicitly about frames (`speed`, `hardRestart`, a table's own
   per-frame contents).
2. **The common case stays the short case.** A proportional pattern, an inline note string, a
   bare `gate -N` are all still one token or one line; the declarative, checked, multi-line
   forms (`each`, verb blocks, `chip`) are opt-in for when a song actually needs them.
3. **A new spelling is sugar over an existing mechanism until proven otherwise.** `hardRestart
   N` compiles to the same AST a `gate -N` field would; `gate legato`/`slideTo` share one
   onset-suppression primitive with the pre-existing `fx:"tp"` legato path. Reaching for a new
   runtime mechanism is the last resort, not the first move.
4. **Silence is never the answer to ambiguity.** A declared length that doesn't match its
   contents, a bar line off the grid, a subdivision that doesn't land on a whole frame, an `@n`
   pattern mixed with `/N` durations - every one of these is a named, specific compile
   diagnostic (error or, where the song is still playable, a warning/info naming exactly what
   was auto-fit and why), never a silent best guess.
5. **Note, slot, and gate stay three different things, on purpose** (above) - the distinction
   that makes "how long is this note" and "does the envelope retrigger" two independently
   answerable questions instead of one fact doing both jobs badly.

### Tempo (`tempo_stmt`)

- `tempo 120 beats 4` - constant, exactly today's `song(120, {beatsPerCycle: 4})`.
- `speed 6 beats 4` - the same clock in frames per beat instead of BPM, exactly (`bpm = 3000 /
  framesPerBeat`). Mutually exclusive with `tempo`; see "Song time" above.
- `tempo 120 beats 4 offset 1` - the same, shifted: the whole song's pattern clock moves by that
  many **beats**, so a positive value delays the song (real frame N is played as compiler frame
  `N - offset*framesPerBeat`, i.e. `song()`'s `startOffsetFrames`, which before 2.17.0 had no DSL
  spelling). Fractional and negative are both legal. By hand this is a pickup beat; for the
  GoatTracker importer it is how a tune whose player takes one row to get going declares that
  once, at song level, instead of rotating its own note data and knocking every pattern boundary
  off the bar grid. `beats` and `offset` may appear in either order and either may be omitted.
- `tempo "<120 96>" beats 4` - tempo varies **per bar**, using the *same* `<>` alternation
  mini-notation already means everywhere else (deliberate reuse, not a new mechanism to learn).
  Maps to GoatTracker's `EXY`/`FXY` funktempo - a real, precedented technique, not a novelty.
- With a constant tempo the compiler derives frames per bar once. With a tempo pattern it
  re-derives frames per bar at each bar boundary, sampling the pattern once per bar.

### Target: the drive rate (`target_stmt`)

ADR 0002. `tempo`/`speed`/`beats` answer "how long is a beat, counted in frames". This answers
"how long is a frame, counted in CPU cycles" - i.e. how often the player is invoked. DSL v2 had
one hard-coded answer to that (PAL, 50.00 Hz) and real tunes have many: NTSC tunes run at ~60 Hz,
GoatTracker multispeed tunes at 2x/3x/4x the video rate, and CIA-timer-driven players at rates
with no video relationship at all (`Welle_Erdball_Alles_Luege.sid`: 12,801 cycles = 76.97 Hz).

- `target { video PAL }` - 985248/19705 = 49.99990 Hz. Identical to declaring nothing.
- `target { video NTSC }` - 1022727/17045 = 60.00158 Hz.
- `target { video PAL multispeed 3 }` - 3 play calls per video frame, 19705/3 = 6568 cycles.
- `target { driveCycles 12801 }` - the general form: this many CPU cycles per play call, exact.

`driveCycles` is the primitive; `video` and `multispeed` are macros over it, so `multispeed`
without `video` means "multispeed against PAL". Giving both `driveCycles` and `multispeed` is an
error, since they compute the same number two ways.

Cycles rather than Hz because a float Hz accumulates rounding over a multi-minute song, because
SID playback is cycle-exact, and because `driveCycles` is exactly what a play-call clock
measurement produces, so a SID-to-DSL converter can populate it mechanically instead of rounding
through Hz.

One note on the PAL number: 19705 is `machineClock / 50`, the frame grid the rest of the toolchain
uses, not the hardware-true 312 * 63 = 19656-cycle raster period, which would be 0.25% faster and
would silently change the meaning of every existing song. See `src/dsl/driveRate.js`.

Out of scope, same as ADR 0001's mid-song tempo changes: changing the drive rate mid-song, and
sub-frame (digi) timing.

### Named patterns (`pattern_decl`, `$name`)

- `pattern hook = "c4!3 e4!2 g4"` then `voice 1 lead $hook` (as the whole `notes` value) or
  inline inside another string, e.g. `"$hook $hook c4"`.
- `pattern hook 4 = "..."` declares the pattern's length in beats and makes it metric - see
  "Song time" above for what that means and what it is checked against.
- **No engine change needed.** This resolves entirely at the AST-interpretation stage: before
  building a voice's note pattern, substitute every `$name` occurrence with the *already-parsed*
  named pattern (or, for inline interpolation inside a string, splice its literal mini-notation
  source before handing the combined string to `parseMini()`). Forward references are a compile
  error (`pattern` declarations must appear before first use - no need to support recursion).

### Patterned transform arguments

Strudel's defining idea is that a control is never a constant - it is itself a pattern. This DSL
already applied that to `insts`, `wave:"..."`, `sync`/`ring` and tempo; since 2.14.0 it applies to
**transform arguments** too. Anywhere a transform takes a number, a quoted mini-notation string
works instead, sampled once per bar:

```
voice 0 $riff |> add "<0 0 0 2>"      ; bars 0-2 as written, bar 3 up a whole step
voice 1 $hook |> fast "<1 1 2>"
```

This is the only way to scope a transform to *part* of a song. A bare `add 2` applies to the
voice's whole pattern for all time, which is why the GoatTracker converter originally had to
duplicate a pattern per transposition instead of using `add` at all.

Sampling takes the first event of the argument pattern in each bar, so the natural spelling is
an alternation (`<...>`), which yields exactly one value per bar. A finer-grained argument isn't
an error - only its first value in each bar is consulted, since a transform like `fast` applies
to a whole bar at once and has no coherent meaning *within* one.

### Scale degrees (`scale`)

`scale "c4:minor"` reads the pattern's values as **scale degrees** and maps each to a real pitch.
Degrees past the end of the scale wrap into the next octave; negative degrees wrap downward.

```
voice 0 "0 2 4 7" |> scale "c4:minor"        ; C Eb G C
voice 0 "0 2 4" |> add 2 |> scale "c4:minor" ; Eb G Bb -- a diatonic third, still in key
```

That second line is the point: on degrees, `add 2` moves two **scale** steps and stays in key,
where the same `add 2` on note numbers shifts two semitones and leaves the key behind.

**Write the pattern as degrees, not note names.** `miniNotation.js` resolves `c4` to the number
60 while parsing, so by the time `scale` sees it there is nothing left to distinguish the pitch C4
from degree 60 - it will be read as the latter. This is a real constraint, not a rough edge to be
smoothed over later; Strudel draws the same line by giving degrees their own control (`n`) separate
from `note`. Available scales are listed in `src/pattern/scales.js` (`SCALE_INTERVALS`).

### Named tables (`table_decl`, `$name`)

The same idea as `pattern_decl` above, one level down: `pattern` shares a note sequence between
voices, `table` shares a **per-frame table** between instruments.

```
table filterSweep = 1088 1040 992 944 896 848 800

instrument bass
  adsr 0 8 8 8
  filterTable $filterSweep
end

instrument lead
  adsr 4 7 7 7
  filterTable $filterSweep
end
```

- Valid in any table-valued field (`wave`/`pw`/`arp`/`filterTable`/`test`/`gate`), in an
  `instrument` block or as a short-mode inline attr (`filterTable:$filterSweep`).
- **Spliced in place**, not substituted wholesale - so a row can mix a shared table with literal
  values (`pw $attack 2048`), concatenate two shared tables, or build one table from others
  (`table full = $attack $sustain`). A composed table is flattened at declaration time.
- Forward references are a compile error, same as `pattern` and for the same reason: `interpret()`
  walks the statement list once, in source order, so a name used before it's declared simply
  isn't in the map yet.
- **No engine change needed** - this resolves entirely during AST interpretation, producing the
  exact same flat array the literal form would.

Why it matters in practice: a converted GoatTracker song routinely gives *every* instrument a
byte-identical filter sweep (a 128-entry table repeated across nine instruments is over a thousand
numbers of pure duplication). `table` collapses that to one declaration and one reference per
instrument.

### Writing long tables: `ramp` and `!`

Per-frame tables get long fast - a filter sweep is routinely 100+ values. Two spellings keep
that readable, both **pure expansions**: they produce the exact same flat array the literal
values would, and the engine never sees them.

```
pw ramp 2175..4095 step 127 over 68     ; sweep up, then hold at the ceiling
filterTable ramp 1088..0 step -48 over 128
pw 2048!104 ramp 4..36 step 8 over 5     ; hold 2048 for 104 frames, then sweep
```

- `ramp A..B step D over N` - start at `A`, add `D` each frame, for `N` frames, **saturating at
  `B`**. At least one of `step`/`over` is required (with neither, the ramp has no length). With
  only `over`, the step is derived to land exactly on the target; with only `step`, it runs until
  it reaches the target.
- The saturation is a deliberate DSL choice, not a reproduction of GoatTracker's own hardware
  behavior - and the two are genuinely different things, not two descriptions of the same one
  (ADR 0004). A GoatTracker pulse/filter **modulation-step** row ("for N frames, add D per frame")
  writes into an 8-bit accumulator (`gplay.c`'s `filtercutoff`/pulse equivalent) that **wraps**,
  not saturates: there is no clamp anywhere in that addition, so a row that keeps stepping past the
  accumulator's range re-enters from the other side and keeps moving (confirmed by a
  hardware-accurate register capture, not inferred). `ramp 1088..0 step -48 over 128` (a typical
  converted 128-value filter sweep) reaches its saturating floor of 0 after 23 frames and then
  holds flat for the remaining 105 - 82% of the table - where the real
  8-bit accumulator would instead wrap and keep sweeping. `ramp` only reproduces a tracker row
  exactly for the common case where the row's own step count already keeps it inside range; it is
  not a general guarantee, and this table is not that case. A wrapping sawtooth is still
  expressible without changing `ramp`'s default: a shorter ramp carrying the `loop` shape
  terminator (`ramp A..B step D over N loop`) restarts from `A` once its own table position is
  exhausted, producing the same repeating-sawtooth shape as the hardware wrap. See ADR 0004 for the
  full correction and the worked example.
- `V!N` - `N` copies of `V`, the same repetition spelling mini-notation already uses inside a
  string (`"c4!3"`). Long constant stretches are common in real tables (a pulse width held for
  104 frames) and read far better than a degenerate `ramp V..V over N`.

Both compose with everything else a table row can hold, including `$name` table references.

### Instruments (`instrument_block`)

The actual headline addition - table-driven, per-note-relative programs, modeled directly on
GoatTracker's wavetable/pulsetable/filtertable/`7XY`-`CXY` commands.

**Critical semantic distinction, worth restating because it's easy to get backwards**: a `voice`
's `notes` string samples by **bar position** (`cycleIndex`, `posFrac`) - Strudel semantics,
unchanged. An instrument's `wave`/`arp`/`pw`/`filterTable`/`test` tables sample by
**frames-since-this-note's-own-onset** - tracker semantics. These are genuinely different clocks
sampled against genuinely different counters. Get this wrong and arpeggios/PWM sweeps will
appear to phase-drift against the song instead of restarting fresh on every note.

- `wave pulse tri pulse noise` - per-frame waveform sequence. A single bare word (`wave pulse`)
  is today's static case, unchanged. Join names with `+` (`wave noise+pulse pulse`) for waveforms
  sounding **simultaneously** in one frame slot: a table slot holds a single value, so unlike the
  bar-sampled pattern form below there's no `[a,b]` stack available and the combination has to
  live inside the slot itself. A bare number is the same thing spelled as a raw control-byte
  bitmask (`224` == `saw+pulse+noise`), which is what real tracker data carries.
- `arp 0 4 7 12` - per-frame **relative semitone offsets** added to the voice's currently
  sounding note. This is the table-driven alternative to `[c4,ef4,g4]` inline chord syntax -
  both exist and both produce the same class of effect (fake polyphony via fast retrigger); an
  instrument's `arp` table is the reusable, named version.

  A cell may instead be written `=MIDI` (e.g. `arp 0 =36 0 12`), an **absolute** MIDI note that
  pins that one frame to a fixed pitch regardless of which note is playing. This is GoatTracker's
  wave-table note column $80-$DF - the classic fixed-pitch drum/percussion technique - which a
  relative-only `arp` cannot express at all (`gplay.c`'s `WAVEEXEC` adds the sounding note only
  for the sub-$80, relative half). Relative and absolute cells mix freely in one row, and `=60!4`
  repeats like any other cell. It is spelled as a number rather than a note name on purpose: this
  grammar's lexer has no `#` character, so a table cell could never spell `c#4` the way
  mini-notation does, and a second, subtly different note-name convention just for table rows
  would be a worse wart than one explicit number.
- `pw` - same idea for pulse width.
- `filterTable` - same idea for the shared filter's cutoff, except that it **writes** that one
  cutoff rather than offsetting it: while the table runs, its value IS the cutoff, replacing the
  song-level `filter` block's own `cutoff` (and any `fx fcut.N`). Since `2.19.0` this is
  independent of `route` - routing (`$D417`'s low bits) only picks which voices pass THROUGH the
  one filter, it never decides what the cutoff VALUE is - and two voices whose tables run at the
  same time do not sum: the most recently started table wins, ties going to the higher voice
  index. This is what the SID does (one filter, one cutoff register) and what GoatTracker does
  (`gplay.c` has ONE global `filterptr`/`filtercutoff` pair, re-pointed by whichever note most
  recently started an instrument carrying a filter table, and it writes `sidreg[0x16]` outright).
  Before `2.19.0` it was a per-voice OFFSET, summed across the routed voices and added to the
  `filter` block's `cutoff` - under which a filter table on an unrouted voice did nothing at all,
  silently.
- `loop N` - once a table runs past its last index, wrap back to index `N` (not necessarily 0)
  and continue - GoatTracker/SID-Wizard's sustain-loop idiom (short attack transient, then loop
  a sustain segment for as long as the note's gate stays open). A bare `N` is one global default
  applied to every table this instrument has (unchanged since this field's introduction). Since
  `2.9.0`, `loop` also accepts one or more `TABLE INT` pairs instead - `loop wave 1 pw 3
  filterTable 0` - giving each named table (`wave`/`pw`/`arp`/`filterTable`/`test`/`gate`) its OWN
  independent loop point, matching GoatTracker/SID-Wizard's real model (a wave table's loop marker
  and a pulse table's loop marker are genuinely independent data in the original tracker file - a
  single shared index can't preserve both when they differ, which real ported songs routinely do).
  A table not mentioned in a `TABLE INT` form falls back to 0 (plain wraparound), same default as
  the global form.
- `vib DEPTH RATE` / `slide TIME` - continuous per-note pitch modulation. `vib` is a sine
  wobble of `DEPTH` semitones at `RATE` Hz, re-triggered (phase-zeroed) on every note onset -
  same "frames since onset" clock as the tables, not the song's absolute bar position, since a
  real vibrato LFO restarts its phase each new note. `slide` glides from the previous note's
  pitch to this note's pitch over `TIME` (fraction of a beat) - GoatTracker's toneportamento.
- `vibTri CMP SPEED DELAY` (since `2.18.0`) - the SAME lane, but GoatTracker's own vibrato rather
  than a sine: `gplay.c`'s `vibtime` counter steps by 2 and inverts whenever it passes `CMP`,
  and its low bit decides whether this frame adds or subtracts `SPEED` from the running frequency
  register. That integrates into a TRIANGLE in raw register units (not semitones), whose first
  quarter period is only half the swing of every later one, and it starts `DELAY` frames after
  the onset (GoatTracker's per-instrument `vibdelay`). One half period lasts `CMP+2` frames and
  the swing is `+/-(CMP/2+1)*SPEED`. `SPEED` may be `sN` instead of a number, meaning "one
  semitone at the current note, shifted right by `N`" - GoatTracker's own note-relative speed
  encoding, also accepted by the `fx` lane's parameters. An instrument setting both `vib` and
  `vibTri` gets `vibTri` (the exact one); a converted GoatTracker song only ever emits `vibTri`.
- `test 1 0 0 0` - per-frame test-bit table. The cheapest addition in the spec: `packControl`
  in `registerMap.js` already accepted a `test` field, so only the table plumbing was new.
- `gate FRACTION` - identical to today's `.gate()` on `VoiceChain` (`song.js`), just also
  settable per-instrument as a default. Since `2.16.0`, a NEGATIVE number is accepted as a second
  spelling of the same idea in **absolute frames** rather than a proportion: `gate -2` holds the
  gate high until the note's last 2 frames, whatever that note's length is. That is GoatTracker's
  per-instrument `gatetimer` (`gplay.c` fetches the next note `gatetimer & 0x3f` ticks early and
  keys the gate off there, leaving it low until the new note's own tick 0), i.e. the classic hard
  restart, and it is a fixed TICK COUNT - not a proportion of a row - so no single `FRACTION` can
  express it for a song whose notes are not all the same length. Converters emit this form;
  see "Articulation defaults" below for why the *default* is still legato. Since `2.9.0`, `gate`
  also accepts 2+ space-separated
  tokens instead of one `FRACTION` - a per-frame 0/1 gate-bit **program**, frames-since-onset
  sampled and loopable exactly like `wave`/`arp`/`pw`/`test` (`gate 1 1 1 0 1 1 1 0`). This is a
  genuinely different mechanism from the single fraction: a fraction can only ever turn the gate
  off once per note, while a gate-bit program can toggle it off then back ON mid-note -
  GoatTracker/SID-Wizard's real technique of retriggering a percussive ADSR envelope from a
  wave-table row without starting a brand-new note. When present, the table wins over the static
  fraction entirely (same "instrument table wins" precedence as `wave`/`pw`/`test` above). A
  voice-level `gate` field/override (see "Voices" below) stays the static-fraction-only form -
  genuinely two different clocks/mechanisms sharing one field name depending on where they're
  written, the same split `wave` already has between an instrument's bare table and a voice's
  quoted pattern.

**Articulation defaults (why a `gate` is needed at all).** With no `gate` field, a note holds the
gate high for its whole slot, so a run of consecutive notes is one continuous gate-high run: the
envelope never retriggers, and anything that samples the registers once per frame (a `.sid`
register dump, a conversion scorer, a player's own frame loop) sees ONE note. `compile.js` does
force a gate-off/gate-on pair of writes on a new onset, which restarts a cycle-accurate emulator's
envelope, but both writes land inside the same frame and are therefore invisible at frame
granularity. This legato default is deliberate and stays: it is the right feel for live coding,
and hand-written songs depend on it. Anything that wants per-note articulation - in particular
a SID-to-DSL converter, whose input is a real tracker song where every note ends with real
gate-low frames - must say so explicitly with `gate -N` (or a `FRACTION`, or a gate table).

### Per-row effect commands (`fx`)

`vib`/`slide` are *instrument-level, continuous* modulation - a property of a patch, running for
as long as a note sounds. Tracker songs also carry a completely different thing: **one-shot,
per-row commands**, fired from a pattern's command/command-data columns on a specific row, with
no relation to which instrument is playing. GoatTracker's set of these was, until 2.15.0, the
single largest structural gap between that model and this one, and the GoatTracker importer said
so in its own warnings: those rows were simply dropped.

`fx` is a parallel mini-notation string, structurally the same idea as `insts`:

```
voice 1 lead "c4 e4 g4 c5" fx:"~ tp.24 vib.16.100 ad.0.10"
```

```
voice 2
  notes "c3 ~ g3 ~"
  fx "pup.128 ~ pdn.s3 sr.4.4"
end
```

**Why a parallel pattern rather than inline annotations on `notes`.** `insts` already solved
exactly this shape of problem - a second tracker column that lines up with the note column but is
sampled independently - and it is the proven idiom here. Reusing it means `fx` gets `<...>`
alternation, `!n`, `@n`, `$name` splices and every transform for free; the importer reuses its
whole orderlist-folding machinery unchanged; and a voice's note text stays readable instead of
growing a second, effect-specific micro-syntax inside every note token.

**Sampling differs from `insts`, deliberately.** An `insts` value is read only at a note ONSET (a
row's instrument column belongs to that row's note). An `fx` value is read on the fx pattern's
OWN step boundaries, because a tracker command executes on its row whether or not that row
carries a note - and the rows carrying portamento and vibrato are overwhelmingly the note-less
ones.

Token syntax is `name.param` (or `name.param.param`). A param is a plain integer, or `sN`
meaning "one semitone's worth of frequency at the current note, shifted right by N" -
GoatTracker's own note-relative speed form, which cannot be resolved to a constant ahead of time.

| Token | Effect |
| --- | --- |
| `pup.N` / `pdn.N` | Portamento up/down: add/subtract N frequency-register units every frame. |
| `tp.N` | Tone portamento: glide toward this note's pitch at N per frame, clamping on arrival (`tp.0` snaps). Also **legato** - it suppresses the note's gate retrigger and instrument re-init. |
| `vib.CMP.SPEED` | Vibrato: GoatTracker's own triangle LFO (compare value, then speed). |
| `ad.A.D` / `sr.S.R` | Overwrite this voice's attack/decay or sustain/release nibbles until the next note onset. |
| `wave.N` | Overwrite the waveform bits (ignored when the instrument has a `wave` table, which rewrites them every frame anyway). |
| `pw.N` | Overwrite the pulse width, `0-4095` (the same 12-bit domain as an instrument's own `pw` table column; a value outside it is an error, not a clamp). Ignored when the instrument has a `pw` table, for the same reason `wave.N` yields to a `wave` table. |
| `fcut.N` | Set the shared filter's cutoff (GoatTracker's 8-bit value, scaled to the SID's 11-bit register). |
| `vol.N` | Set the master volume nibble. |

**`pw.N` is the one token here with no GoatTracker command behind it** (since `4.2.0`). GoatTracker
reaches pulse width only indirectly, through `CMD_SETPULSEPTR` (`0x09`), which repoints the running
pulse TABLE and is deliberately unsupported here (see the triage table below). `pw.N` is instead a
DSL-native one-shot register write, in the shape `wave.N` already established, added because pulse
width was the only SID register class with no `fx` spelling at all while being the largest class of
register writes a SID-to-DSL converter could not express, measured over a sample of tunes from the
High Voltage SID Collection.

#### GoatTracker command triage

Every command is taken from GoatTracker 2's actual player (`src/gplay.c`'s `CMD_*` switch and its
per-tick `TICKNEFFECTS` block), not from reverse-engineering byte patterns.

| GT cmd | Name | Disposition |
| --- | --- | --- |
| `0x01` | portamento up | new syntax: `pup.N` |
| `0x02` | portamento down | new syntax: `pdn.N` |
| `0x03` | tone portamento | new syntax: `tp.N` (incl. the legato retrigger suppression) |
| `0x04` | vibrato | new syntax: `vib.CMP.SPEED`. NOT foldable into instrument-level `vib`: that is a continuous per-patch LFO, this is a per-row trigger with GoatTracker's own triangle shape and its own reset semantics. |
| `0x05` | set AD | new syntax: `ad.A.D`. Overlaps `insts` in principle (switching to an instrument with a different `adsr`), but only in principle: the command changes AD *without* touching the wave/pulse/filter tables or restarting anything, which an instrument switch cannot express. |
| `0x06` | set SR | new syntax: `sr.S.R`, same reasoning. |
| `0x07` | set waveform | new syntax: `wave.N`. Overlaps an instrument `wave` table, and where a table exists the table genuinely wins (gplay.c runs `WAVEEXEC` after the command switch) - so this only matters for table-less instruments, and is implemented that way. |
| `0x08` | set wavetable pointer | **deliberately unsupported.** Repoints the running table to an absolute index in GoatTracker's single GLOBAL wave table. DSL v2 instruments own private, already-sliced copies of their tables; there is no global index space to jump within, and synthesizing one instrument per reachable jump target would change instrument identity (and therefore envelope/onset behavior) rather than just the table position. |
| `0x09` | set pulse pointer | **deliberately unsupported**, same reason. Note that the `pw.N` token is NOT this command: it writes the pulse-width register directly and does not touch any table pointer. |
| `0x0a` | set filter pointer | **deliberately unsupported**, same reason. |
| `0x0b` | set filter control | **deliberately unsupported.** Packs filter resonance and per-voice routing into one byte; DSL v2's `filter` block owns both as song-level fields with no per-row lane to write them from. |
| `0x0c` | set filter cutoff | new syntax: `fcut.N` |
| `0x0d` | set master volume | new syntax: `vol.N` |
| `0x0e` | funktempo | **deliberately unsupported.** Alternates between two tempos row by row; DSL v2's tempo varies per BAR at finest, and one bar is a whole pattern play. |
| `0x0f` | set tempo | **maps onto the existing `tempo` feature**, not onto `fx`. The GoatTracker importer folds the song's commanded tempo into the `tempo`/`beats` it emits (previously it hardcoded GoatTracker's power-on default of 6 frames per row, which ran every self-tempo-setting song at the wrong speed). A song that changes tempo mid-tune still cannot be represented exactly; the most-commanded value wins and the importer warns. |

### Per-note instrument switching (`insts`)

A single `voice` statement binds exactly one static instrument for its whole duration - but real
tracker songs constantly switch instruments *within one channel*, row by row (GoatTracker's own
per-row instrument column: e.g. a bass line and kick/snare hits sharing one physical SID voice
when only 3 are available). `insts` closes this gap without changing how a static instrument is
attached at all:

```
voice 0 bass "c2 c2 c2 c2 c2 c2 c2 c2" insts:"bass bass drum bass bass bass drum bass"
```

```
voice 0 bass
  notes "c2 c2 c2 c2 c2 c2 c2 c2"
  insts "bass bass drum bass bass bass drum bass"
end
```

`insts` is a full mini-notation pattern of already-declared `instrument` NAMEs (bare words) -
always quoted, never a bare per-frame table, since which instrument a note uses is a property of
the note pattern's own timeline (bar-position sampled, the SAME clock as `notes` itself), not
something that varies mid-note the way a tracker table does. Sampled once at each note's own
onset and held fixed for that note's whole duration (re-sampling every frame would let the
instrument change out from under a still-sounding note if `insts`'s own subdivisions are finer
than the note pattern's - no tracker analogue for that). A note whose onset has no active `insts`
value falls back to the voice's own static instrument (the named instrument, or inline attrs) -
so `insts` never needs an entry for every note, only the ones that actually switch.

**Engine notes**: `song.js`'s `VoiceChain` carries the per-note instrument pattern
(`.instPattern()`), and `compile.js` samples it at each onset. Instrument tables are sampled by
`voiceRuntime[v].framesSinceOnset` (looping at `loop`'s index once past the end), and the
continuous vibrato/slide pitch-offset lane is keyed by the same per-onset frame counter, which
already existed for chord-arpeggio/`.gate()` purposes: additive state usage, not new
architecture.

### Voices (`voice_stmt`)

Both forms compile to the same internal `VoiceChain` - short mode is exactly the single-line
grammar production, nothing else differs:

```
voice 0 "c2 ~ c2 g1" wave:pulse adsr:0,4,10,4 pw:1536      ; short mode
```

```
voice 0 bass                          ; extended mode - bass is a named instrument
  notes "c2 ~ c2 g1"
  vib   0.15 5
end
```

`inline_attr`s in short mode (`wave:pulse`, `adsr:0,4,10,4`) are sugar for an anonymous
one-off instrument built inline - they don't require a separate `instrument` block. `pipeline`
(`|>`) attaches after everything else, whichever form is used.

Short mode can also reference a named `instrument` block directly, instead of building an
anonymous one from inline attrs - put the instrument name right after the voice index, before
the note value (a string or a `$name` reference):

```
voice 0 bass "c2 ~ c2 g1"      ; short mode, named instrument + literal string
voice 0 bass $bassLine         ; short mode, named instrument + a $name reference
```

A named instrument and inline attrs are mutually exclusive on one voice - `voice 0 bass "c4"
wave:pulse` is a compile error (`can't combine a named instrument... with inline instrument
attributes`), not silently one winning over the other, since which one should apply is
genuinely ambiguous.

`jux`/`off` transforms that operate on a **note** pattern auto-claim a second free hardware
voice, exactly like today's `song.js`/`compile.js` `{left, right}`-pair mechanism (unchanged -
`resolveVoiceAssignments` in `compile.js` already implements this).

A long `notes` pattern can use a backtick string instead of `"..."` to lay it out one bar per
physical line rather than as one very long line - no other change, `notes` still just takes one
STRING value either way:

```
voice 0
  notes `c4 e4 g4 c5
         e4 g4 b4 e5
         g4 b4 d5 g5`
end
```

Short mode's `inline_attr`s and `pipeline` may also continue on their own lines, not just the
line the notes string closes on - a real gap once `notes` can be a multi-line backtick string:
without this, the line break right after a multi-line string's closing backtick would end the
voice statement, and whatever followed (`wave:...`, `|> ...`) would fail to parse as a bogus new
top-level statement. The parser tells "the next line continues this voice statement" apart from
"the next line starts a new statement" purely by token shape (an `IDENT ':'` or a `PIPE`
immediately after the line break - the same style already used for `switch_case`'s own
termination, not indentation-based):

```
voice 0 `c3 ~ g2 ~
  a2 ~ f2 ~`
  wave:pulse
  adsr:0,6,12,6
  pw:1024
  |> every 4 rev
```

### Sync/ring (hard sync / ring modulation)

```
voice 0 "c2 c2 c2 c2" sync:1                 ; short mode - bare 0/1 constant
```

```
voice 2                                      ; extended mode - a pattern that varies over time
  notes "e2 e2 e2 e2"
  ring "1 0"
end
```

Hard sync (bit 1) and ring modulation (bit 2) of the voice control byte, against this voice's
`(i+2)%3` neighbor. **No engine change was needed for this**: `song.js`'s
`VoiceChain.sync()`/`.ring()` already existed before this grammar did and already accept a
constant, a mini-notation string, or a `Pattern` via `toPattern()`; a `sync`/`ring` voice field
(or short-mode inline attr) is purely grammar/interpreter wiring that calls straight into those
existing methods. (Added in 2.1.0; sync and ring are load-bearing SID lead/bell techniques, not an
optional nicety.)

### Waveform: table vs. pattern

`wave` is the one field with two genuinely different forms, disambiguated by quoting (the same
"quoted vs. bare" rule the rest of the grammar already follows):

```
voice 0 "c4 c4 c4 c4" wave:pulse                 ; bare -- a per-frame TABLE (tracker semantics,
                                                  ; frames-since-onset, see "Instruments" above)
voice 0 "c4" wave:"pulse|triangle"               ; quoted -- a full mini-notation PATTERN
                                                  ; (bar-position sampled, same clock as notes)
```

The quoted form routes to `song.js`'s `VoiceChain.s()` - a method that already existed before
DSL v2 and already samples by bar position exactly like a note pattern; DSL v2's interpreter
previously just never called it. This means the quoted form gets the *entire* mini-notation
grammar for free, including two things worth calling out explicitly:

- `wave:"pulse|triangle"` - picks randomly between waveforms once per bar (`a|b` choice, same
  mechanism a note pattern already uses), instead of a fixed value.
- `wave:"[pulse,triangle]"` - combines two waveform bits at once (a stack, same as a chord's
  `,`). The bare table form spells the same thing as `pulse+triangle` inside one slot.

Extended mode has the equivalent as its own voice field: `wave "pulse|triangle"`.

**One accepted precedence caveat**: if a voice has *both* an attached instrument with its own
`wave` table *and* a `wave:"..."` pattern, the instrument's table still wins (`compile.js`'s
existing, unchanged rule - an attached instrument's table always takes priority over the chain's
own static waveform). This doesn't match `vib`/`slide`/`gate`'s "voice-level value overrides the
instrument's default" model; making it consistent would mean changing that longstanding
precedence rule, which is out of scope here - noted as a real, documented asymmetry rather than
a silent trap.

### Raw register escape hatch (`reg`)

```
voice 0 bass
  notes "c2 ~ c2 ~"
  reg control "9 8 9 8"    ; symbolic register name
  reg 0x15     "3 3 3 3"   ; raw hex fallback for anything not named
end
```

Symbolic names come straight from `registerMap.js`'s existing `VOICE_REGISTER_OFFSETS` keys
(`freqLo freqHi pwLo pwHi control attackDecay sustainRelease`, resolved relative to *this*
voice's base) plus `FILTER_REGISTERS`'s keys (`cutoffLo cutoffHi resonanceRouting modeVolume`,
absolute, not voice-relative) for the shared block. The string's mini-notation pattern is
sampled per-frame and, if a Hap is active, **overwrites** whatever byte the rest of the compiler
computed for that address on that frame, applied last, after every other field. This is
deliberately the lowest-level, least-guarded construct in the language - the explicit
"deep-hack" trapdoor, matching GoatTracker's `5XY`/`6XY`/`7XY`/`CXY` direct-register commands.

### Filter (`filter_block`)

`cutoff` is the shared filter's cutoff for as long as nothing else is writing it. Two things
outrank it, in this order: an instrument's `filterTable` (see the instrument section - since
`2.19.0` it WRITES this one cutoff absolutely, whatever `route` says), then the `fx` lane's
`fcut.N`. `route` chooses which voices pass through the filter and nothing else; it has never
had, and since `2.19.0` no longer appears to have, any say in what the cutoff value is.

Otherwise unchanged semantics from today's `FilterChain` (`cutoff`/`resonance`/`mode`/`route`), extended so
that **all four fields**, not just `cutoff`, accept either a `signal_expr` (a direct
grammar-level spelling of what's today written as JS `sine.range(300,2200).slow(4)`) or a
`switch_block` for genuinely piecewise automation - not only a flat constant:

```
filter
  cutoff switch bar
    < 3   : saw 160..0 / 3
    else  : saw 2047..1440 / 9 >> 3
  resonance 15
  mode      bandpass
  route     1 2
end
```

This is a direct, non-JS spelling of a v1-style `continuous((cycleIndex, posFrac) =>
cycleIndex < 3 ? ... : ...)` callback - no engine change needed, `query.js`'s `continuous()`
already accepts exactly this kind of custom per-`(cycleIndex, posFrac)` callback; the AST
interpreter generates one instead of requiring hand-written JS. `song.js`'s `FilterChain` already
stored `resonancePattern`/`modePattern`/`routingPattern` as real `Pattern`s (via `toPattern()`),
same as `cutoffPattern`, so the grammar only had to reach that path for every field.

`resonance` and `mode` use exactly the same `signal_expr`/`switch_block` shapes as `cutoff` (a
bare `switch bar` case for `mode` is just an identifier with no range/speed, e.g. `bandpass`,
since mode is a category name, not a numeric sweep). `route`'s case values can't be spelled as an
identifier-led `signal_expr` at all - a routing value is a list of voice indices, e.g. `0 1` - so
`switch_case` also accepts a bare `number_list` leaf; `route` is the only field that ever uses it
inside a switch case. Worked example - a filter chain where every field genuinely varies over
time, not just cutoff:

```
filter
  cutoff switch bar
    < 4   : sine 300..2200 / 4
    else  : saw 2047..1440 / 9 >> 3
  resonance switch bar
    < 4   : sine 4..15
    else  : 12
  mode switch bar
    < 4   : bandpass
    else  : lowpass
  route switch bar
    < 4   : 0
    else  : 0 1
end
```

Before bar 4: cutoff sweeps 300-2200Hz, resonance sweeps 4-15, mode is bandpass, and only voice
0 is routed through the filter. From bar 4 on: cutoff switches to the descending saw sweep,
resonance locks to a constant 12, mode switches to lowpass, and both voice 0 and voice 1 are
routed through the filter.

## What happens to the old JS builder API

`song()`/`.voice()`/`.filter()`/`note()` in `song.js` (now in `rasterklang-score`) **do not go
away** - they stop being something an author types and become the compile target this grammar's
AST interpreter calls into (same relationship the mini-notation parser already has to `Pattern`
objects: source text in, calls into an existing builder API out).

## Example: a complete song

```
; --- tempo & meter -----------------------------------------------------
tempo 120 beats 4

; --- reusable named pattern ---------------------------------------------
pattern hook = "c4!3 e4!2 g4"

; --- instruments ----------------------------------------------------------
instrument bass
  wave pulse
  adsr 0 4 10 4
  pw   1536
  vib  0.15 5
  gate 0.85
end

instrument lead
  wave triangle pulse triangle noise
  arp  0 4 7 12
  pw   2048 2400 2800 3200 2800 2400
  filterTable 0 20 40 60 40 20
  adsr 0 9 0 0
  loop 2
  slide 0.08
end

instrument perc
  wave noise
  adsr 0 2 0 2
  test 1 0 0 0
end

; --- voices ----------------------------------------------------------------
voice 0 bass
  notes "c2!3 e2!2 g1 <c2 f2>"
  |> every 4 rev
end

voice 1 lead
  notes $hook
  |> every 4 rev
  |> degradeBy 0.1
  |> sometimesBy 0.3 (ply 2)
end

voice 2 perc
  notes "x(5,8,2)"     ; euclidean modifies an atom, same as mini-notation everywhere else -
                        ; it can't stand alone; "x" is a placeholder atom, irrelevant under a
                        ; noise waveform where pitch isn't perceptible
end

; --- shared filter -----------------------------------------------------------
filter
  cutoff switch bar
    < 3  : saw 160..0 / 3
    else : saw 2047..1440 / 9 >> 3
  resonance 15
  mode      bandpass
  route     1 2
end
```
