// Walks the DSL v2 AST (parser.js's output) and calls into song.js's
// builder API (song(), note(), .s(), .adsr(), .pw(), .inst(), .test(),
// .reg(), the FilterChain methods) plus the pattern engine's
// signals/combinators. Everything downstream of interpret() (song.js's
// builder object, compile.js) has no knowledge of the DSL.
//
// Two design decisions not spelled out in the grammar spec's EBNF:
//
// - Named patterns ($name) resolve here, at interpret time, against a symbol
//   table that only ever contains names declared strictly earlier in the
//   statement list (see interpret()'s main loop), which makes a forward
//   reference a compile error for free. In `notes` a $name is a structural
//   reference (ADR 0003 D5, resolveNotesValueStructural); the other
//   pattern-valued fields still splice the referenced source text (see
//   spliceSigils()).
// - `jux`/`off`/`superimpose` are the only pipeline transforms that claim a
//   second SID voice (producing song.js's { left, right } split marker
//   instead of a Pattern) - see applyPipelineStep()'s comment for why that
//   reading is unambiguous within this grammar.

// The pattern engine and the register compiler live in rasterklang-score;
// interpret.js builds directly against them rather than an engine-agnostic
// IR.
import { song, note } from 'rasterklang-score/compiler/song.js';
import { continuous, discrete, pure } from 'rasterklang-score/pattern/query.js';
import { parseMini, parseMiniForHighlight } from 'rasterklang-score/pattern/miniNotation.js';
import { timeline } from 'rasterklang-score/pattern/timeline.js';
import { sine, saw, square, tri, rand, perlin } from 'rasterklang-score/pattern/signals.js';
import {
  rev, ply, iter, degradeBy, add, every, sometimesBy,
  struct, mask, chunk, palindrome, linger, rot, someCyclesBy, scale,
} from 'rasterklang-score/pattern/combinators.js';
import { instrument } from './instrument.js';
import { parseFxToken } from './fx.js';
import { analyzeProgram, patternBeatsOf } from './analyze.js';
import { expandHolds, expandEach } from './miniLength.js';
import { resolveTarget } from './driveRate.js';
import { VOICE_REGISTER_OFFSETS, FILTER_REGISTERS, voiceRegister } from 'rasterklang-score/compiler/registerMap.js';

// mapPatternValues(pattern, fn) -> a new discrete Pattern with every Hap's
// `value` passed through fn, everything else (startFrac/durFrac/any extra
// per-Hap metadata like mini-notation's sourceStart/sourceEnd) untouched.
// Only used for `insts` (see buildVoiceChain below) to turn a mini-notation
// pattern of bare instrument-NAME tokens into a pattern of the actual
// resolved instrument objects compile.js samples - Pattern itself (query.js)
// has no generic map/fmap, and this is the only place in the whole DSL that
// needs one.
function mapPatternValues(pattern, fn) {
  return discrete((cycleIndex) => pattern.query(cycleIndex).map((h) => ({ ...h, value: fn(h.value) })));
}

// diag(severity, line, message) - same plain shape analyze.js's diagnostics
// use, so a diagnostic computed here (ADR 0003 D17 needs frame arithmetic
// only this file has - see below) rides along in softDiagnostics exactly
// like one analyze.js produced.
function diag(severity, line, message) {
  return { severity, line, message };
}

// formatBeatsForDiag(n) - shortest exact decimal for a D17 diagnostic
// number (frames/bar, bpm, ...): an integer prints with no decimal point.
function formatBeatsForDiag(n) {
  return Number.isInteger(n) ? String(n) : String(Math.round(n * 1e4) / 1e4);
}

// exactDenominator(beats) -> the smallest positive integer q such that
// beats*q is a whole number (ADR 0003 D17): beats is always D1's
// beatUnit/K[.5|*2/3], so q is bounded and small in every practical case
// (a 64th-note triplet's beats has q=96) - 256 is a generous ceiling.
function exactDenominator(beats) {
  for (let q = 1; q <= 256; q += 1) {
    if (Math.abs(beats * q - Math.round(beats * q)) < 1e-6) return q;
  }
  return 256;
}

// beatFractionLabel(beats, beatUnit) -> "1/16"-style label for a D17
// diagnostic, from a beats value that is beatUnit/K for some K (dotted/
// triplet durations round to their nearest plain fraction for display,
// which is a cosmetic simplification, not a computation the diagnostic's
// own numbers depend on).
function beatFractionLabel(beats, beatUnit) {
  const k = Math.max(1, Math.round(beatUnit / beats));
  return `1/${k}`;
}

// === Named-pattern resolution ($name) =======================================

const SIGIL_RE = /\$([a-zA-Z_][a-zA-Z0-9_]*)/g;

// ABSOLUTE_IDENTITY_SEGMENTS (ADR 0003 D5): a one-segment "translation"
// covering every non-negative position, under which the editor's segment
// arithmetic (`docStart + (sourceStart - textStart)`) reduces to exactly
// `sourceStart`
// - i.e. a no-op, which is correct once every Hap's sourceStart/sourceEnd is
// already an absolute document position by construction (see
// resolveNotesValueStructural's docOffset). Reused (not rebuilt) per voice:
// it carries no per-voice state.
const ABSOLUTE_IDENTITY_SEGMENTS = [{ textStart: 0, textEnd: Number.MAX_SAFE_INTEGER, docStart: 0 }];

// lookupPattern(name, patterns, line) -> the entry interpret()'s main loop
// stored for a `pattern` declaration, or a compile error for an undeclared
// (including forward-referenced) name.
//
// Highlighting anchors: a `segments` list is an ordered list of
// { textStart, textEnd, docStart } spans that together cover a resolved
// source string, each saying "text[textStart..textEnd) is really the text at
// docStart.. in the document". segments is null when ANY piece has no known
// position: a highlight built from partially-wrong data would be worse than
// none, so the editor falls back to a text search in that case.
function lookupPattern(name, patterns, line) {
  if (!patterns.has(name)) {
    throw new Error(`line ${line}: pattern "$${name}" is not defined (declare it with \`pattern ${name} = "..."\` before this use - forward references aren't supported).`);
  }
  return patterns.get(name);
}

// spliceSigils(text, patterns, line, textDocStart) -> { source, segments }
// - text with every `$name` occurrence replaced by that named pattern's own
// (already-resolved) raw mini-notation source (a plain textual splice, not a
// parsed-then-recombined Pattern, per the grammar spec's "for inline
// interpolation inside a string, splice its literal mini-notation source
// before handing the combined string to parseMini()"), plus the segments
// list describing where every piece of the result actually came from.
// `textDocStart` is `text`'s OWN absolute document position (null if
// unknown) - used to anchor the literal stretches between/around splices;
// each spliced-in pattern anchors itself via its own stored segments.
function spliceSigils(text, patterns, line, textDocStart) {
  let source = '';
  let segments = textDocStart != null ? [] : null;
  let lastIndex = 0; // position in the ORIGINAL `text`, not the resolved `source`

  const pushLiteralSegment = (literalStart, literalText) => {
    if (segments) segments.push({ textStart: source.length, textEnd: source.length + literalText.length, docStart: textDocStart + literalStart });
  };

  for (const m of text.matchAll(SIGIL_RE)) {
    const matchStart = m.index;
    if (matchStart > lastIndex) {
      const literal = text.slice(lastIndex, matchStart);
      pushLiteralSegment(lastIndex, literal);
      source += literal;
    }

    const p = lookupPattern(m[1], patterns, line);
    if (p.elements != null) {
      throw new Error(`line ${line}: pattern "$${m[1]}" is an arrangement (ADR 0003 D4) and has no mini-notation text to splice here.`);
    }
    // A named pattern stores no segments list (its Haps carry absolute
    // positions directly, ADR 0003 D5), so a spliced reference makes the
    // whole result's segments unavailable. The callers that still splice
    // (insts/fx/reg/sync/ring/tempo-as-pattern) then fall back to the
    // editor's text-search highlighting.
    segments = null;
    // p.source may itself contain literal `$name` tokens (a declaration's
    // own refs are not pre-spliced, ADR 0003 D5), so recurse to return
    // fully-resolved text with no dangling `$`.
    source += /\$/.test(p.source) ? spliceSigils(p.source, patterns, line, null).source : p.source;
    lastIndex = matchStart + m[0].length;
  }

  if (lastIndex < text.length) {
    const literal = text.slice(lastIndex);
    pushLiteralSegment(lastIndex, literal);
    source += literal;
  }

  return { source, segments };
}

// validateAdsr(values, line) (ADR 0003 D14): each of attack/decay/
// sustain/release is a 4-bit SID envelope register nibble - out of range
// is a silent-garbage bug (packAD/packSR mask to 4 bits with no warning)
// without this, named here once and called from both `adsr` consumer
// sites (a full `instrument ... end` block's field and short-mode's
// `adsr:` inline attribute), which otherwise parse the same 4 numbers
// through two different grammar productions with no shared validation
// point upstream.
const ADSR_NAMES = ['attack', 'decay', 'sustain', 'release'];
function validateAdsr(values, line) {
  for (let i = 0; i < 4; i += 1) {
    const v = values[i];
    if (typeof v !== 'number' || !Number.isInteger(v) || v < 0 || v > 15) {
      throw new Error(`line ${line}: adsr's ${ADSR_NAMES[i]} value must be an integer 0..15, got ${v}.`);
    }
  }
}

// checkNoteNames(text, line) (ADR 0003 D14): a NOTES-field-only check
// (never run over insts/fx/table-row text, which are free-form word/
// number vocabularies mini-notation deliberately stays agnostic to - see
// miniNotation.js's own header comment) for the one shape that's almost
// certainly a note-name typo rather than an intentional word: a single
// letter, optional accidental, then digits (miniNotation.js's NOTE_RE
// shape) whose letter isn't a-g. `h4` is the flagship case: German note
// naming's H (B natural) is a common false friend for what this DSL calls
// "b".
//
// ':' is deliberately NOT in the boundary-exclusion charset below, even
// though it IS one of scanWordToken's word-chars (so that a `slideTo:h4`
// target gets checked too, same as a bare `h4` would): ':' only ever
// appears as the separator right after the `slideTo` prefix (see
// miniNotation.js's SLIDE_TO_PREFIX), so it should act as a boundary for
// THIS check even though the mini-notation scanner itself treats
// "slideTo:h4" as one token.
//
// '$' IS in the exclusion charset, for the opposite reason to ':': a `$name`
// structural pattern reference (ADR 0003 D5) is a pattern NAME, never a note,
// and names like `$p0`, `$i0` or `$h1` match the note-typo shape exactly. Any
// notes string that splices a reference whose name is a letter followed by
// digits would otherwise be rejected outright. A note is never legitimately
// preceded by '$', so excluding it costs no real coverage of the typo case.
const INVALID_NOTE_RE = /(?<![A-Za-z0-9.#$])([A-Za-z])([sf#]?)(-?\d+)(?![A-Za-z0-9.#])/g;
function checkNoteNames(text, line) {
  for (const m of text.matchAll(INVALID_NOTE_RE)) {
    const letter = m[1].toLowerCase();
    if ('abcdefg'.includes(letter)) continue;
    const suggestion = letter === 'h' ? ` (German notation's "h" is this DSL's "b" - did you mean "b${m[3]}"?)` : ' (valid note letters are a-g)';
    throw new Error(`line ${line}: invalid note name "${m[0]}"${suggestion}`);
  }
}

// resolveNotesValueStructural(notesNode, patterns, analysisPatterns,
// barBeats, beatUnit, line, flavor) -> Pattern | null. ADR 0003 D5: $name is
// a structural reference (miniNotation.js's parseRef()), not text
// substitution - built here via parseMini's `refs` option, against the SAME
// `patterns` map PatternDecl processing populates with already-built
// Pattern objects (each already carrying ABSOLUTE sourceStart/sourceEnd by
// construction - see PatternDecl's own docOffset - so nothing downstream
// needs segments-based translation for a referenced pattern's own Haps).
//
// `flavor` is 'pattern' (audio, parseMini, rests are silence()) or
// 'highlightPattern' (parseMiniForHighlight, rests are real Haps - see
// miniNotation.js's parseMiniForHighlight comment) - PatternDecl processing
// builds and caches both for every named pattern, and a $ref here resolves
// to whichever flavor this call is building, so a highlight-flavored voice
// pattern is built entirely from highlight-flavored referenced patterns,
// never a mix.
//
// Returns null (not a Pattern) for a bare sigil naming an ARRANGEMENT
// pattern - that shape is handled entirely by buildVoiceChain's
// isArrangement path instead (bareArrangementRef), which this function
// defers to rather than duplicating.
function resolveNotesValueStructural(notesNode, patterns, analysisPatterns, barBeats, beatUnit, line, flavor) {
  const parse = flavor === 'highlightPattern' ? parseMiniForHighlight : parseMini;
  const patternRefs = (name) => {
    const p = patterns.get(name);
    if (!p) return undefined;
    if (p.elements != null) {
      throw new Error(
        `line ${line}: pattern "$${name}" is an arrangement (ADR 0003 D4) and can only be used as a voice's whole notes or nested inside another arrangement, not spliced into a mini-notation string.`,
      );
    }
    return p[flavor];
  };
  if (notesNode.isSigil) {
    const p = lookupPattern(notesNode.sigilName, patterns, line);
    if (p.elements != null) return null; // arrangement pattern - buildVoiceChain's bareArrangementRef handles this
    return p[flavor];
  }
  checkNoteNames(notesNode.source, line);
  const weightRefs = (name) => patternBeatsOf(name, analysisPatterns, barBeats);
  const expanded = expandHolds(notesNode.source, beatUnit, weightRefs);
  return parse(expanded, { beatUnit, docOffset: notesNode.docStart ?? 0, refs: patternRefs, weightRefs });
}

// === Pipeline / transforms ====================================================

// isSplit(x) -> true for the { isVoiceSplit: true, left, right } marker
// compile.js's resolveVoiceAssignments() specifically checks for (see
// compile.js's own isVoiceSplit() - NOT song.js's structurally-duck-typed
// isVoiceSplitPair(), which only ever runs inside the VoiceChain
// constructor on whatever note() was given; here we're building/mutating
// chain.notePattern directly, after construction, so this module has to
// produce the exact same tagged shape compile.js expects itself).
function isSplit(x) {
  return x != null && typeof x === 'object' && x.isVoiceSplit === true;
}

// argFn(arg, apply) -> (pattern) => Pattern, resolving a transform argument
// that may be a constant OR a per-cycle pattern (parser.js's parseArg()).
//
// In the constant case `apply` is called once and the resulting Pattern is
// returned directly, with no wrapper or per-cycle re-derivation. The pattern case re-resolves the argument for each cycle and
// applies the transform with THAT cycle's value, which is what makes
// `add "<0 0 2>"` mean "cycles 0 and 1 unchanged, cycle 2 up a whole step".
//
// Sampling takes the first Hap of the argument pattern in that cycle: a
// per-cycle control is normally written as an alternation (`<0 0 2>`), which
// yields exactly one Hap per cycle. A finer-grained argument pattern is not
// an error, it just means only its first value in each cycle is consulted -
// a transform like `fast` has to be applied to a whole cycle at once, so
// there is no coherent way to vary it *within* one.
function argFn(arg, apply) {
  if (arg == null) return apply(undefined);
  if (arg.source === undefined) return apply(arg.value);
  const argPattern = parseMini(arg.source);
  return (p) =>
    discrete((c) => {
      const haps = argPattern.query(c);
      if (haps.length === 0) return p.query(c);
      return apply(haps[0].value)(p).query(c);
    });
}

// transformToPatternFn(node) -> Pattern -> Pattern, for every transform
// EXCEPT the voice-claiming ones (jux/off/superimpose; see applyPipelineStep
// below for why those are handled separately, at the top level only).
// Nesting one inside another transform's fn slot (e.g. `every 4 (jux rev)`)
// has no defined meaning, so it's a clear error.
function transformToPatternFn(node) {
  switch (node.type) {
    case 'rev':
      return (p) => rev(p);
    case 'fast':
      return argFn(node.n, (n) => (p) => p.fast(n));
    case 'slow':
      return argFn(node.n, (n) => (p) => p.slow(n));
    case 'ply':
      return argFn(node.n, (n) => (p) => ply(n, p));
    case 'iter':
      return argFn(node.n, (n) => (p) => iter(n, p));
    case 'degradeBy':
      return argFn(node.p, (v) => (p) => degradeBy(v, p));
    case 'add':
      return argFn(node.n, (n) => (p) => add(n, p));
    case 'every': {
      const fn = transformToPatternFn(node.fn);
      return argFn(node.n, (n) => (p) => every(n, fn, p));
    }
    case 'sometimesBy': {
      const fn = transformToPatternFn(node.fn);
      return argFn(node.p, (v) => (p) => sometimesBy(v, fn, p));
    }
    case 'struct': {
      const boolPattern = parseMini(node.source);
      return (p) => struct(boolPattern, p);
    }
    case 'mask': {
      const boolPattern = parseMini(node.source);
      return (p) => mask(boolPattern, p);
    }
    case 'chunk': {
      const fn = transformToPatternFn(node.fn);
      return argFn(node.n, (n) => (p) => chunk(n, fn, p));
    }
    case 'palindrome':
      return (p) => palindrome(p);
    case 'linger':
      return argFn(node.frac, (v) => (p) => linger(v, p));
    case 'rot':
      return argFn(node.n, (n) => (p) => rot(n, p));
    case 'someCyclesBy': {
      const fn = transformToPatternFn(node.fn);
      return argFn(node.p, (v) => (p) => someCyclesBy(v, fn, p));
    }
    case 'scale':
      return (p) => scale(node.source, p);
    case 'jux':
    case 'off':
    case 'superimpose':
      throw new Error(
        `line ${node.line}: '${node.type}' claims a second SID voice and can only be used as a top-level pipeline step, not nested inside another transform.`,
      );
    default:
      throw new Error(`line ${node.line}: unknown transform '${node.type}'.`);
  }
}

// applyPipelineStep(current, node) -> the voice's note-pattern-or-split
// value after applying one pipeline transform.
//
// jux/off are special-cased here (not folded into transformToPatternFn)
// because they're the only two transforms this grammar exposes that claim a
// second hardware voice, producing song.js's { left, right } split instead
// of a Pattern - see the grammar spec's "jux/off voice-claiming" section.
// The pattern engine's combinators.js off() produces a same-voice echo,
// which is useful on a *control* pattern such as a filter cutoff sweep, but
// this grammar never applies the transform pipeline to a control pattern
// (filter fields use their own signal_expr/switch_block grammar), so here
// `off` always means "echo on a second voice."
function applyPipelineStep(current, node) {
  if (node.type === 'jux') {
    if (isSplit(current)) {
      throw new Error(`line ${node.line}: 'jux' can't be applied to a pattern that's already voice-split (chaining 'jux'/'off' after one isn't supported).`);
    }
    const fn = transformToPatternFn(node.fn);
    return { isVoiceSplit: true, left: current, right: fn(current) };
  }
  if (node.type === 'superimpose') {
    // Same voice-claiming shape as jux: the original and its transformed copy
    // sound simultaneously, and one SID oscillator can't do that, so the copy
    // takes a second hardware voice. (jux's name carries a stereo-panning
    // connotation that means nothing on a real SID; superimpose is the same
    // operation without it.)
    if (isSplit(current)) {
      throw new Error(`line ${node.line}: 'superimpose' can't be applied to a pattern that's already voice-split (chaining 'jux'/'off'/'superimpose' after one isn't supported).`);
    }
    const fn = transformToPatternFn(node.fn);
    return { isVoiceSplit: true, left: current, right: fn(current) };
  }
  if (node.type === 'off') {
    if (isSplit(current)) {
      throw new Error(`line ${node.line}: 'off' can't be applied to a pattern that's already voice-split (chaining 'jux'/'off' after one isn't supported).`);
    }
    const fn = transformToPatternFn(node.fn);
    // `off`'s time offset is patternable too, via the same argFn wrapper the
    // other numeric arguments use - the shift happens on the already-
    // transformed copy, so the wrapper goes around shiftedBy, not inside it.
    const shifted = argFn(node.time, (t) => (p) => p.shiftedBy(t))(fn(current));
    return { isVoiceSplit: true, left: current, right: shifted };
  }
  const fn = transformToPatternFn(node);
  if (isSplit(current)) {
    // A non-splitting transform stacked on top of an already-split pattern
    // (e.g. `|> jux rev |> degradeBy 0.1`) spreads across both sides - a
    // reasonable, graceful reading the grammar spec's examples never
    // exercise, since neither side should silently stop being transformed
    // just because a split happened earlier in the pipeline.
    return { isVoiceSplit: true, left: fn(current.left), right: fn(current.right) };
  }
  return fn(current);
}

function applyTransforms(notePattern, transformNodes) {
  let current = notePattern;
  for (const node of transformNodes) current = applyPipelineStep(current, node);
  return current;
}

// === Instruments ===============================================================

const INLINE_ATTR_NAMES = ['wave', 'adsr', 'pw', 'arp', 'filterTable', 'loop', 'vib', 'vibTri', 'slide', 'test', 'gate', 'sync', 'ring'];

// sync/ring are hard sync / ring modulation (registerMap.js's packControl
// bits 1/2, between a voice and its (i+2)%3 neighbor) - a voice-level
// relationship, not a per-instrument "patch" property, so unlike every other
// inline attr they're pulled out BEFORE buildInstrumentFromInlineAttrs runs
// and applied directly via song.js's existing VoiceChain.sync()/.ring()
// (which accepts a constant/string/Pattern via toPattern()).
const VOICE_LEVEL_INLINE_ATTR_NAMES = new Set(['sync', 'ring']);

// expandRamp({from, to, step, over, shape}) -> { values, loopIndex }, the
// per-frame values a `ramp` writes (ADR 0003 D9). Callers thread .loopIndex
// through to the instrument's per-table loop map - see expandTableRefs
// below.
//
// Deliberately clamps at the target rather than overshooting or stopping
// early: a value that reaches the target before `over` frames have passed
// holds there for the remaining frames, so `ramp 1088..0 step -48 over 128`
// is 23 descending values, then 0 for the rest. The saturation is a DSL
// design choice, not a reproduction of GoatTracker's filter modulation
// step, whose 8-bit cutoff accumulator wraps instead (ADR 0004). A
// qualitatively similar wrapping sawtooth can be written with a `loop`
// shape terminator, e.g. `ramp 1088..0 step -48 over 23 loop`.
//
// shape is null (saturate at the target), 'pingpong' or
// 'loop': the base sweep (below) is exactly the same computation either
// way - shape only decides what gets APPENDED and what loopIndex results,
// reusing the existing per-table loop-point mechanism (compile.js's
// sampleTable()/resolveLoopIndex()) rather than a new runtime primitive:
//   'loop':     values unchanged, loopIndex 0 -> restarts from `from` each
//               pass (a sawtooth: a discontinuous jump back at the seam).
//   'pingpong': values + the sweep reversed (excluding both ends, which
//               already appear once each - the reversed run's first entry
//               duplicates the forward run's last, and looping back to
//               index 0 already supplies `from` again) then loopIndex 0 ->
//               a seamless, endless up-down-up-down triangle, because the
//               sequence's last entry is one step short of `from` and its
//               very next (looped) sample is exactly `from` again.
function expandRamp({ from, to, step, over, shape, line }) {
  const lo = Math.min(from, to);
  const hi = Math.max(from, to);
  // Only `over` given: derive the step that lands exactly on the target.
  // (A 1-frame ramp is just its start value, no step to derive.)
  const delta = step ?? (over > 1 ? (to - from) / (over - 1) : 0);
  if (step !== null && Math.sign(step) !== Math.sign(to - from) && to !== from) {
    throw new Error(
      `line ${line}: a 'ramp' from ${from} to ${to} steps by ${step}, which moves away from the target and never arrives - flip the sign.`,
    );
  }
  const length = over ?? Math.floor(Math.abs(to - from) / Math.abs(delta)) + 1;
  const out = [];
  for (let i = 0; i < length; i += 1) {
    out.push(Math.max(lo, Math.min(hi, from + delta * i)));
  }
  if (shape === 'loop') return { values: out, loopIndex: 0 };
  if (shape === 'pingpong') {
    const back = out.slice(1, -1).reverse();
    return { values: [...out, ...back], loopIndex: 0 };
  }
  return { values: out, loopIndex: null };
}

// expandTableRefs(tokens, tables, line) -> the same token list with every
// `$NAME` marker (the parser's `{ ref }` objects) replaced in place by the
// tokens of the `table` declaration it names, and every `ramp`/`!N` repeat
// expanded.
//
// Splices rather than substitutes-wholesale, so a row can mix a shared table
// with literal values (`pw $attack 2048`) or concatenate two shared tables,
// exactly like `$name` inside a mini-notation string. Unknown names are a
// compile error here rather than silently producing a short table, and
// because interpret() walks statements in source order a forward reference
// hits this same error for free.
//
// Returns tokens[] | { tokens, loopIndex } (ADR 0003 D9). Returns a bare array when nothing in this row
// implied a loop point; returns { tokens, loopIndex } when a `ramp ...
// pingpong`/`ramp ... loop` did, OR when a `$NAME` ref resolved to a shared
// table that itself carries an implied loop point (its own declaration's
// `table NAME = ramp ... pingpong`/`... loop` - see the `tables` Map's
// { values, loopIndex } shape at the TableDecl case below).
//
// Both sources feed ONE precedence rule: the LAST entry in the row that
// carries an implied loop index wins, whether it is an inline shaped ramp
// or a $ref to a shaped table. So `filterTable $sweep` (where `table sweep =
// ramp ... loop`) behaves exactly like writing the ramp inline. `pw $decay
// $sweep` takes $sweep's loopIndex because it's the last entry that has
// one; `pw $sweep 2048 2048` keeps it too, because nothing after it
// supplies a replacement.
//
// The loop index is deliberately NOT offset by out.length at the splice
// point (the same choice expandRamp's loopIndex: 0 makes for an inline
// ramp). That is correct when the loop-carrying entry is the row's whole
// content, which covers real instrument tables; a mixed row with leading
// literals is a known limitation.
function expandTableRefs(tokens, tables, line) {
  if (!tokens || !tokens.some((t) => t && typeof t === 'object')) return tokens;
  const out = [];
  let loopIndex = null;
  for (const t of tokens) {
    if (t && typeof t === 'object' && t.gen === 'ramp') {
      const expanded = expandRamp(t);
      out.push(...expanded.values);
      if (expanded.loopIndex != null) loopIndex = expanded.loopIndex;
      continue;
    }
    if (t && typeof t === 'object' && t.repeat !== undefined) {
      for (let i = 0; i < t.count; i += 1) out.push(t.repeat);
      continue;
    }
    if (t && typeof t === 'object' && t.ref !== undefined) {
      const resolved = tables.get(t.ref);
      if (!resolved) {
        const known = [...tables.keys()];
        throw new Error(
          `line ${line}: unknown table '$${t.ref}'${known.length > 0 ? ` (declared tables: ${known.join(', ')})` : ' (no `table` declarations in this song)'}. A table must be declared with \`table ${t.ref} = ...\` BEFORE the instrument that uses it.`,
        );
      }
      out.push(...resolved.values);
      // resolved.loopIndex is the shared table's own implied loop point
      // (set at the TableDecl case from its declaration's shaped ramp),
      // folded in exactly like an inline ramp's above.
      if (resolved.loopIndex != null) loopIndex = resolved.loopIndex;
    } else {
      out.push(t);
    }
  }
  return loopIndex == null ? out : { tokens: out, loopIndex };
}

// expandTableVerbs(verbs, line) -> { tokens, loopIndex } (ADR 0003 D13):
// runs a verb block (astVisitor.js's buildTableVerbLine output) forward,
// exactly the way a reader would say it out loud, into the same flat
// per-frame array the bare-token form produces. `set`/`add` push one entry
// per step; `wait` holds the CURRENT running value for N frames without
// moving it; `loop`/`stop` both just mark "the loop point is HERE" (this
// table's own length so far) - the distinction between "loop back to a
// mid-table position" and "hold the last value forever" needs no separate
// code path, because sampleTable() (compile.js) already clamps an
// out-of-range loopIndex to the table's last entry, which is exactly what
// a `stop`/trailing `loop` (nothing emitted after it) naturally produces.
function expandTableVerbs(verbs, line) {
  let current = 0;
  const out = [];
  let loopIndex = null;
  for (const v of verbs) {
    switch (v.verb) {
      case 'set':
        current = v.value;
        out.push(current);
        break;
      case 'add':
        for (let i = 0; i < v.repeat; i += 1) {
          current += v.value;
          out.push(current);
        }
        break;
      case 'wait':
        for (let i = 0; i < v.count; i += 1) out.push(current);
        break;
      case 'loop':
      case 'stop':
        if (loopIndex != null) throw new Error(`line ${line}: a table block can only have one 'loop'/'stop' verb`);
        loopIndex = out.length;
        break;
      default:
        throw new Error(`line ${line}: unknown table verb '${v.verb}'`);
    }
  }
  if (out.length === 0) {
    throw new Error(`line ${line}: a table block must produce at least one value ('set'/'add'/'wait')`);
  }
  return { tokens: out, loopIndex };
}

// mergeImpliedLoop(opts, impliedLoop) (ADR 0003 D9): folds a `ramp ...
// pingpong`/`ramp ... loop`'s implied per-table loop index (impliedLoop,
// { tableName: index }) into opts.loop, WITHOUT overriding a table the
// author gave an explicit `loop`/`loop TABLE N` for - an explicit choice
// wins over a ramp's implicit one, the same "explicit beats derived"
// precedent D3's `each` follows for a declared pattern length.
function mergeImpliedLoop(opts, impliedLoop) {
  const names = Object.keys(impliedLoop);
  if (names.length === 0) return;
  const explicit = opts.loop && typeof opts.loop === 'object' && !Array.isArray(opts.loop) ? opts.loop : {};
  const merged = { ...impliedLoop, ...explicit };
  // A single already-set global INT loop (opts.loop a plain number) is
  // itself an explicit author choice, applying to every table - leave it
  // untouched rather than trying to merge a scalar with a map.
  if (typeof opts.loop === 'number') return;
  opts.loop = merged;
}

function buildInstrumentFromBlockFields(fields, tables) {
  const opts = {};
  // ADR 0003 D11/D14: `gate` (a static fraction/hardRestart frame count, a
  // mode word, or a per-frame table) and its `hardRestart` alias both
  // compile to this same field. Giving it twice is rejected: the shapes
  // have no defined combination (e.g. `gate retrigger` forces its own
  // gateReleaseFrames, silently discarding an earlier `gate 0.5`), so a
  // repeat is always a mistake.
  let gateFieldLine = null;
  const impliedLoop = {};
  // ADR 0003 D13: a verb block's own loop point, tracked SEPARATELY from
  // ramp-implied loop points (impliedLoop above) - a ramp's implicit loop
  // silently loses to an explicit `loop` field (mergeImpliedLoop's own
  // precedent), but D13's own text is stricter: "using both on one table
  // is an error naming the conflict," checked below once opts.loop is
  // fully resolved.
  const verbImpliedLoop = {};
  // ADR 0003 D10: speed/delay per table, resolved per table the way
  // impliedLoop/opts.loop already are - a { TABLE: { speed, delay } } map,
  // instrument.js's normalizeSpeedDelay() turns this into what
  // compile.js's resolveSpeedDelay() reads. Only wave/arp/pw/filterTable/
  // test/gate carry f.speed/f.delay at all (astVisitor.js's
  // extractSpeedDelay, gated to those field names).
  const speedByTable = {};
  // One expansion pass up front rather than per-field-case below: every
  // table-valued field carries its row the same way (`f.tokens`), so there's
  // no reason for each `case` to know about `$NAME` refs individually.
  fields = fields.map((f) => {
    if (f.speed != null || f.delay != null) speedByTable[f.type] = { speed: f.speed ?? 1, delay: f.delay ?? 0 };
    if (f.verbs) {
      const { tokens, loopIndex } = expandTableVerbs(f.verbs, f.line);
      if (loopIndex != null) verbImpliedLoop[f.type] = loopIndex;
      return { ...f, tokens };
    }
    if (!f.tokens) return f;
    const expanded = expandTableRefs(f.tokens, tables, f.line);
    if (Array.isArray(expanded)) return { ...f, tokens: expanded };
    if (f.type !== 'loop') impliedLoop[f.type] = expanded.loopIndex;
    return { ...f, tokens: expanded.tokens };
  });
  for (const f of fields) {
    switch (f.type) {
      case 'wave':
        opts.wave = f.tokens;
        break;
      case 'arp':
        opts.arp = f.tokens;
        break;
      case 'pw':
        opts.pw = f.tokens;
        break;
      case 'filterTable':
        opts.filterTable = f.tokens;
        break;
      case 'test':
        opts.test = f.tokens;
        break;
      case 'adsr':
        validateAdsr([f.a, f.d, f.s, f.r], f.line);
        opts.adsr = [f.a, f.d, f.s, f.r];
        break;
      case 'loop':
        // Global INT form (f.value) vs. per-table TABLE INT pairs (f.pairs)
        // - see parser.js's parseLoopField() and instrument.js's
        // normalizeLoop() for the two shapes instrument({loop}) accepts.
        opts.loop = f.pairs ? Object.fromEntries(f.pairs.map((p) => [p.table, p.value])) : f.value;
        break;
      case 'vib':
        opts.vib = { depth: f.depth, rate: f.rate };
        break;
      case 'vibTri':
        opts.vibTri = { cmp: f.cmp, speed: f.speed, speedShift: f.speedShift, delay: f.delay };
        break;
      case 'slide':
        opts.slide = { time: f.time };
        break;
      case 'gate':
        if (gateFieldLine != null) {
          throw new Error(
            `line ${f.line}: this instrument already has a 'gate' (or 'hardRestart') field on line ${gateFieldLine} - ` +
              `gate's value/mode/table shapes have no defined combination, so give it once.`,
          );
        }
        gateFieldLine = f.line;
        // Single value (f.value) = a static fraction (or a negative
        // hardRestart-style frame count); tokens (f.tokens, 2+) = the
        // per-frame gate-bit table; f.mode (ADR 0003 D11) = 'legato' or
        // 'retrigger' behaviour - see parser.js's parseGateField().
        if (f.tokens) opts.gateTable = f.tokens;
        else if (f.mode) opts.gateMode = f.mode;
        else opts.gate = f.value;
        break;
      default:
        throw new Error(`line ${f.line}: unknown instrument field '${f.type}'.`);
    }
  }
  // ADR 0003 D13: "using both [a 'loop'/'loop TABLE N' field AND a verb-
  // block 'loop'/'stop' for the same table] is an error naming the
  // conflict" - stricter than a ramp's implied loop, which silently loses
  // to an explicit field (mergeImpliedLoop's own precedent, left alone).
  for (const name of Object.keys(verbImpliedLoop)) {
    const hasExplicit =
      typeof opts.loop === 'number' || (opts.loop && typeof opts.loop === 'object' && Number.isInteger(opts.loop[name]));
    if (hasExplicit) {
      throw new Error(
        `instrument declares '${name}'s loop point twice: once via the 'loop' instrument field and once via a 'loop'/'stop' verb inside its own '${name}' block - use one or the other.`,
      );
    }
  }
  mergeImpliedLoop(opts, { ...impliedLoop, ...verbImpliedLoop });
  if (Object.keys(speedByTable).length > 0) opts.speed = speedByTable;
  return instrument(opts);
}

// inline_attrs (short mode's `wave:pulse adsr:0,4,10,4` sugar) build the
// exact same instrument({...}) shape as a full `instrument NAME ... end`
// block - see the grammar spec's "inline_attrs in short mode are sugar for
// an anonymous one-off instrument" line.
function buildInstrumentFromInlineAttrs(attrs, line, tables) {
  const opts = {};
  let sawGate = false;
  const impliedLoop = {};
  // Same one-pass `$NAME` expansion as the block form above, so short mode's
  // `filterTable:$sweep` works identically to the block field.
  attrs = attrs.map((a) => {
    if (!a.values) return a;
    const expanded = expandTableRefs(a.values, tables, line);
    if (Array.isArray(expanded)) return { ...a, values: expanded };
    if (a.name !== 'loop') impliedLoop[a.name] = expanded.loopIndex;
    return { ...a, values: expanded.tokens };
  });
  for (const a of attrs) {
    if (!INLINE_ATTR_NAMES.includes(a.name)) {
      throw new Error(`line ${line}: unknown inline attribute '${a.name}:' (expected one of ${INLINE_ATTR_NAMES.join('/')}).`);
    }
    switch (a.name) {
      case 'wave':
        opts.wave = a.values;
        break;
      case 'arp':
        opts.arp = a.values;
        break;
      case 'pw':
        opts.pw = a.values;
        break;
      case 'filterTable':
        opts.filterTable = a.values;
        break;
      case 'test':
        opts.test = a.values;
        break;
      case 'adsr':
        if (a.values.length !== 4) {
          throw new Error(`line ${line}: inline attribute 'adsr:' needs exactly 4 values (attack,decay,sustain,release), got ${a.values.length}.`);
        }
        validateAdsr(a.values, line);
        opts.adsr = a.values;
        break;
      case 'loop':
        opts.loop = a.values[0];
        break;
      case 'vib':
        opts.vib = { depth: a.values[0], rate: a.values[1] };
        break;
      case 'vibTri': {
        // Short-mode `vibTri:5,s3,8` - the same three values, comma-
        // separated (a shift-derived speed arrives as the string 's3').
        const [cmp, speed, delay] = a.values;
        opts.vibTri = typeof speed === 'string'
          ? { cmp, speedShift: Number(String(speed).slice(1)), delay }
          : { cmp, speed, delay };
        break;
      }
      case 'slide':
        opts.slide = { time: a.values[0] };
        break;
      case 'gate':
        // ADR 0003 D14: give 'gate:' at most once per voice - see the
        // block-field form's own comment on why a repeat is always a
        // mistake rather than a meaningful combination.
        if (sawGate) throw new Error(`line ${line}: this voice already has a 'gate:' inline attribute - give it once.`);
        sawGate = true;
        // Same single-value-vs-table split as the block-field form above -
        // `gate:1,0,0` (3 comma-separated values) is the per-frame gate-bit
        // table form; `gate:0.85` (one value) stays the static fraction.
        if (a.values.length > 1) opts.gateTable = a.values;
        else opts.gate = a.values[0];
        break;
      default:
        break;
    }
  }
  mergeImpliedLoop(opts, impliedLoop);
  return instrument(opts);
}

// finalizeInstrument(baseInstrument, overrides) -> the instrument object
// actually attached to a voice's chain, after layering any extended-mode
// per-voice vib/slide/gate field on top of the referenced/inline
// instrument's own defaults (grammar spec: both `instrument` blocks AND
// `voice_field` expose vib/slide/gate - "instrument defaults, overridable
// per voice"). Reuses the single `.inst()` attachment point rather than
// adding separate VoiceChain methods for overrides: compile.js reads
// vib/slide/gate uniformly off whatever instrument object is attached.
function finalizeInstrument(baseInstrument, overrides) {
  const hasOverride = overrides.vib != null || overrides.slide != null || overrides.gate != null;
  if (!baseInstrument && !hasOverride) return null;
  const merged = { ...(baseInstrument || instrument({})) };
  if (overrides.vib != null) merged.vib = overrides.vib;
  if (overrides.slide != null) merged.slide = overrides.slide;
  if (overrides.gate != null) merged.gate = overrides.gate;
  return merged;
}

// === Raw register escape hatch (`reg`) =========================================

// resolveRegTarget(target, voiceIndex, line) -> a chip-relative (0x00-0x1f)
// register address. Symbolic names come from registerMap.js's existing
// VOICE_REGISTER_OFFSETS (resolved relative to THIS voice) and
// FILTER_REGISTERS (absolute, for the shared block) - see the grammar
// spec's "Raw register escape hatch" section.
function resolveRegTarget(target, voiceIndex, line) {
  if (target.kind === 'hex') return target.value & 0x1f;
  if (Object.prototype.hasOwnProperty.call(VOICE_REGISTER_OFFSETS, target.name)) {
    return voiceRegister(voiceIndex, target.name);
  }
  if (Object.prototype.hasOwnProperty.call(FILTER_REGISTERS, target.name)) {
    return FILTER_REGISTERS[target.name];
  }
  throw new Error(
    `line ${line}: unknown register name '${target.name}' (expected one of ` +
      `${Object.keys(VOICE_REGISTER_OFFSETS).join(', ')}, ${Object.keys(FILTER_REGISTERS).join(', ')}, or a 0xNN hex address).`,
  );
}

// === Voices =====================================================================

// resolveNamedInstrument(name, instruments, vstmt) -> the named instrument
// object, or a clear compile error naming the voice and the missing
// declaration - shared by both short mode (`voice N INSTRUMENT "..."`/
// `voice N INSTRUMENT $pattern`) and extended mode (`voice N INSTRUMENT
// \n ... \n end`), since both reference an `instrument` block by name the
// same way.
function resolveNamedInstrument(name, instruments, vstmt) {
  const found = instruments.get(name);
  if (!found) {
    throw new Error(
      `line ${vstmt.line}: voice ${vstmt.index} references unknown instrument "${name}" ` +
        `(declare it with an \`instrument ${name}\` block before this voice).`,
    );
  }
  return found;
}

// === Arrangements (ADR 0003 D4) ==============================================

// compressToOneCycle({pattern, cycles}) -> {pattern, cycles}: `pattern`'s
// own natural period is `cycles` cycles, compressed to exactly 1 so it can
// be slotted as a SINGLE timeline() entry - that outer timeline's own
// .slow(cycles) (built from the returned `cycles`) undoes the compression
// at the entry's absolute offset. Shared by both ways an arrangement
// element can nest another multi-cycle span: a reference to a named
// arrangement pattern, and a bracketed [...] group.
function compressToOneCycle({ pattern, cycles }) {
  return { pattern: cycles > 0 ? pattern.fast(cycles) : pattern, cycles };
}

// patternPatternAndCyclesFor(name, patterns, analysisPatterns, barBeats,
// line) -> { pattern, cycles }: resolves a NAMED pattern reference to
// exactly ONE cycle of its own content (ready for a caller to timeline()
// it) plus how many bar-cycles that one cycle actually represents.
//
// `patterns` (this file's own map, PatternDecl-name -> content) supplies
// CONTENT: parseMini(source) for a plain mini-notation pattern, or a
// recursively-built nested timeline for an arrangement pattern (D4:
// "references nest to any depth"). `analysisPatterns` (analyze.js's
// authoritative length computation, already run before interpret() gets
// here) supplies LENGTH, so this does not re-derive the "declared, or one
// bar if proportional" rule analyze.js's resolveRefLength() already
// applied, so the two cannot diverge.
function patternPatternAndCyclesFor(name, patterns, analysisPatterns, barBeats, line) {
  const p = patterns.get(name);
  if (!p) throw new Error(`line ${line}: unknown pattern reference "$${name}".`);
  const meta = analysisPatterns.get(name);
  const beats = meta && meta.elements != null ? meta.arrangementBeats : meta && meta.declaredLength != null ? meta.declaredLength : barBeats;
  const cycles = beats / barBeats;
  if (p.elements != null) {
    const inner = buildArrangementPattern(p.elements, patterns, analysisPatterns, barBeats, line);
    return compressToOneCycle(inner);
  }
  // p.pattern is already built (PatternDecl processing, above) with its
  // own absolute source positions - reuse it rather than re-parsing, which
  // both saves the work and is what makes a leaf's own $refs (ADR 0003 D5)
  // resolve correctly without this function needing to know about them.
  return { pattern: p.pattern, cycles };
}

// buildArrangementPattern(elements, patterns, analysisPatterns, barBeats,
// line) -> { pattern, cycles }. `elements` is astVisitor.js's flat
// arrangement shape ([{ line, ref? , group?, repeat }]); `*n`/`[...]`
// grouping is flattened into the timeline()'s entry list here (5.2: "not
// runtime constructs") rather than at parse time.
function buildArrangementPattern(elements, patterns, analysisPatterns, barBeats, line) {
  const entries = [];
  for (const el of elements) {
    const elLine = el.line ?? line;
    let sub;
    if (el.ref != null) {
      sub = patternPatternAndCyclesFor(el.ref, patterns, analysisPatterns, barBeats, elLine);
    } else {
      sub = compressToOneCycle(buildArrangementPattern(el.group, patterns, analysisPatterns, barBeats, elLine));
    }
    for (let i = 0; i < el.repeat; i += 1) entries.push({ lengthCycles: sub.cycles, pattern: sub.pattern });
  }
  const cycles = entries.reduce((sum, e) => sum + e.lengthCycles, 0);
  return { pattern: timeline(entries), cycles };
}

function buildVoiceChain(vstmt, ctx) {
  const { instruments } = ctx;

  // notes $name resolution already happened in resolveVoiceStmtSigils (see
  // interpret()'s main loop): vstmt.resolvedNotesPattern/
  // resolvedNotesHighlightPattern (short mode) and the `notes` field's
  // f.resolvedPattern/resolvedHighlightPattern (extended mode) are built
  // Pattern objects (ADR 0003 D5, resolveNotesValueStructural()) whose Haps
  // carry ABSOLUTE document positions.
  let notesPattern;
  let notesHighlightPattern = null;
  let baseInstrument = null;
  let transformNodes = [];
  const regFields = [];
  let vibOverride = null;
  let slideOverride = null;
  let gateOverride = null;
  let syncValue = null;
  let ringValue = null;
  // wavePatternSource - set only when `wave` was given as a quoted
  // mini-notation STRING (e.g. `wave:"pulse|triangle"`), never for the bare
  // per-frame table form (`wave:pulse`, or an instrument's own `wave` table
  // field) - routed to song.js's VoiceChain.s() below (cycle-position
  // sampled), a completely separate mechanism from the instrument-table one.
  // Precedence caveat: compile.js lets an instrument's table-based wave win
  // over a voice-level pattern wave if both are present for one voice (the
  // reverse of vib/slide/gate's "voice-level override" model). Within one
  // `voice` statement's inline attrs the two forms can't collide: the
  // pattern-form attr is split out below and never reaches the instrument.
  let wavePatternSource = null;
  // instsPatternSource - set only when `insts:"..."`/`insts "..."` was given
  // (see the grammar spec's "Per-note instrument switching" section): a full
  // mini-notation pattern of instrument NAME tokens, resolved below (after
  // baseInstrument/finalInstrument, once ctx.instruments is known to hold
  // everything) into a Pattern of actual instrument objects attached via
  // song.js's VoiceChain.instPattern(). Like wavePatternSource, this is
  // pulled out before the rest of a short-mode voice's inline attrs go into
  // the anonymous inline instrument - `insts` picks BETWEEN instruments per
  // note, it doesn't build one.
  let instsPatternSource = null;
  // instsSegments - resolveVoiceStmtSigils' highlighting anchor for
  // instsPatternSource (see that function's comment), threaded through to
  // song.js's VoiceChain.instSourceSegments below - null whenever this
  // voice has no `insts` field at all, or (defensively) if that field
  // somehow reached here without segments attached.
  let instsSegments = null;
  // fxPatternSource - set only when `fx:"..."`/`fx "..."` was given (see
  // docs/dsl-grammar-spec.md's "Per-row effect commands" section): a
  // mini-notation pattern of effect tokens (`tp.24`, `vib.65.2`, `ad.0.10`,
  // ... - see src/dsl/fx.js), structurally parallel to `notes` exactly the
  // way `insts` is. Split out from a short-mode voice's inline attrs before
  // the rest build the anonymous inline instrument, same as `insts`: `fx`
  // triggers one-shot commands on a voice, it doesn't describe an
  // instrument.
  let fxPatternSource = null;

  if (vstmt.mode === 'short') {
    // ADR 0003 D4: arrangement mode (vstmt.arrangement set, resolvedNotesPattern
    // null) has no mini-notation Pattern to carry through here -
    // buildArrangementPattern() below builds chain.notePattern directly.
    if (vstmt.resolvedNotesPattern != null) {
      notesPattern = vstmt.resolvedNotesPattern;
      notesHighlightPattern = vstmt.resolvedNotesHighlightPattern;
    }
    const wavePatternAttr = vstmt.inlineAttrs.find((a) => a.name === 'wave' && a.isPattern);
    if (wavePatternAttr) wavePatternSource = wavePatternAttr.source;
    const instsAttr = vstmt.inlineAttrs.find((a) => a.name === 'insts');
    if (instsAttr) {
      instsPatternSource = instsAttr.source;
      instsSegments = instsAttr.segments ?? null;
    }
    const fxAttr = vstmt.inlineAttrs.find((a) => a.name === 'fx');
    if (fxAttr) fxPatternSource = fxAttr.source;
    // sync/ring are voice-level, not instrument-level (see
    // VOICE_LEVEL_INLINE_ATTR_NAMES above); a pattern-form `wave:` attr and
    // `insts:` are ALSO not instrument-level (see wavePatternSource/
    // instsPatternSource above) - all are split out before the rest go into
    // the anonymous inline instrument.
    const instAttrs = vstmt.inlineAttrs.filter(
      (a) => !VOICE_LEVEL_INLINE_ATTR_NAMES.has(a.name) && a !== wavePatternAttr && a !== instsAttr && a !== fxAttr,
    );
    if (vstmt.instrumentName) {
      // `voice N INSTRUMENT "..."` / `voice N INSTRUMENT $pattern` -- short
      // mode WITH a named instrument (see parseVoiceStmt()'s 2-token
      // lookahead). Mutually exclusive with instrument-scoped inline attrs
      // on purpose: `voice 0 bass "c4" wave:pulse` combining a named
      // instrument with ad-hoc inline attrs has no defined meaning (which
      // wins?), so this fails loudly rather than silently picking one -
      // same "fail loudly on genuine ambiguity" precedent as jux/off running
      // out of free voices.
      if (instAttrs.length > 0) {
        throw new Error(
          `line ${vstmt.line}: voice ${vstmt.index} can't combine a named instrument ("${vstmt.instrumentName}") ` +
            `with inline instrument attributes (${instAttrs.map((a) => a.name).join('/')}) - use one or the other.`,
        );
      }
      baseInstrument = resolveNamedInstrument(vstmt.instrumentName, instruments, vstmt);
    } else if (instAttrs.length > 0) {
      baseInstrument = buildInstrumentFromInlineAttrs(instAttrs, vstmt.line, ctx.tables);
    }
    const syncAttr = vstmt.inlineAttrs.find((a) => a.name === 'sync');
    const ringAttr = vstmt.inlineAttrs.find((a) => a.name === 'ring');
    if (syncAttr) syncValue = syncAttr.values[0];
    if (ringAttr) ringValue = ringAttr.values[0];
    transformNodes = vstmt.transforms;
  } else {
    if (vstmt.instrumentName) {
      baseInstrument = resolveNamedInstrument(vstmt.instrumentName, instruments, vstmt);
    }
    for (const f of vstmt.fields) {
      if (f.type === 'notes') {
        notesPattern = f.resolvedPattern;
        notesHighlightPattern = f.resolvedHighlightPattern;
      } else if (f.type === 'reg') {
        regFields.push(f);
      } else if (f.type === 'vib') {
        vibOverride = { depth: f.depth, rate: f.rate };
      } else if (f.type === 'slide') {
        slideOverride = { time: f.time };
      } else if (f.type === 'gate') {
        gateOverride = f.value;
      } else if (f.type === 'sync') {
        syncValue = f.value.kind === 'pattern' ? f.value.resolvedSource : f.value.value;
      } else if (f.type === 'ring') {
        ringValue = f.value.kind === 'pattern' ? f.value.resolvedSource : f.value.value;
      } else if (f.type === 'wave') {
        wavePatternSource = f.source;
      } else if (f.type === 'insts') {
        // resolvedSource: $name splices already substituted by
        // resolveVoiceStmtSigils (falls back to the raw source for a field
        // built by any path that skipped that resolution). segments is that
        // same pass's highlighting anchor (see resolveVoiceStmtSigils).
        instsPatternSource = f.resolvedSource ?? f.source;
        instsSegments = f.segments ?? null;
      } else if (f.type === 'fx') {
        fxPatternSource = f.resolvedSource ?? f.source;
      } else if (f.type === 'pipeline') {
        transformNodes.push(...f.transforms);
      } else {
        throw new Error(`line ${f.line}: unknown voice field '${f.type}'.`);
      }
    }
  }

  // ADR 0003 D4: an arrangement (`voice 0 $a $b`) builds its note pattern
  // structurally, via buildArrangementPattern()/timeline(); there is no
  // single source string an arrangement's elements could be parsed from.
  // A bare single sigil that names an ARRANGEMENT-form pattern (`voice 0
  // $verse` where `pattern verse = $a $b`) is the same case in disguise:
  // grammatically it stayed on the notes/isSigil path (see the .g4 file's
  // ambiguity-resolution comment), but ctx.patterns has no mini-notation
  // .source for it to splice - only .elements - so it is read here as a
  // degenerate one-element arrangement, same as analyze.js's collectVoice.
  const bareArrangementRef = (() => {
    const isSigil = vstmt.mode === 'short' ? vstmt.notes && vstmt.notes.isSigil : vstmt.fields && vstmt.fields.find((f) => f.type === 'notes' && f.isSigil);
    const sigilName = vstmt.mode === 'short' ? vstmt.notes?.sigilName : isSigil?.sigilName;
    return sigilName != null && ctx.patterns.get(sigilName)?.elements != null ? sigilName : null;
  })();
  const arrangementElements = vstmt.mode === 'short' && vstmt.arrangement != null
    ? vstmt.arrangement
    : bareArrangementRef != null
      ? [{ line: vstmt.line, ref: bareArrangementRef, repeat: 1 }]
      : null;
  const isArrangement = arrangementElements != null;
  if (!isArrangement && notesPattern === undefined) {
    throw new Error(`line ${vstmt.line}: voice ${vstmt.index} has no \`notes\` (a quoted note pattern or a $patternName reference is required).`);
  }

  let chain;
  if (isArrangement) {
    const { pattern } = buildArrangementPattern(arrangementElements, ctx.patterns, ctx.analysisPatterns, ctx.barBeats, vstmt.line);
    chain = note(pattern);
    // No highlight-flavored arrangement pattern is built, so rests in an
    // arrangement voice are not highlighted (a known gap). Sounding notes
    // still are: compile.js samples notePattern for onsets, and
    // noteHighlightPattern only adds coverage for RESTS.
  } else {
    chain = note(notesPattern);
    chain.noteHighlightPattern = notesHighlightPattern;
  }
  // ADR 0003 D5: every Hap's sourceStart/sourceEnd is already an ABSOLUTE
  // document position by construction (resolveNotesValueStructural's
  // docOffset) - ABSOLUTE_IDENTITY_SEGMENTS lets the existing segments-
  // based highlighting consume that directly: one segment spanning every
  // non-negative position, under which `docStart + (sourceStart -
  // textStart)` reduces to exactly `sourceStart`.
  chain.noteSourceSegments = ABSOLUTE_IDENTITY_SEGMENTS;
  // The voice's metric length in beats, or null for a proportional voice
  // (ADR 0001's D1). Carried on the chain so the compiler/editor can draw
  // and reason about the real grid rather than re-deriving it; the timing
  // itself is already applied below.
  chain.metricBeats = vstmt.metricBeats ?? null;
  // wave:"..." (a mini-notation waveform pattern) goes to song.js's .s(),
  // which samples by cycle position like a note pattern. See
  // wavePatternSource's comment above for the precedence caveat vs. an
  // attached instrument's own table.
  // Every pattern-valued field gets the same hold-token folding `notes`
  // gets inside resolveNotesValueStructural(): `_` means "whatever this
  // lane was holding, keep holding it" for an instrument/waveform/register
  // lane just as naturally as it does for a note, and a string without one
  // is returned unchanged, so this costs pre-3.0.0 songs nothing.
  if (wavePatternSource != null) chain.s(expandHolds(wavePatternSource));
  // A METRIC voice (its notes are one named pattern - or an orderlist of
  // them - carrying a declared length) occupies its declared length and
  // loops at that length, spanning or subdividing bars as needed (ADR
  // 0001's D1/D3, and decision 3 for the shorter-than-a-bar case). One
  // bar is `beats` beats and a pattern's contents fill one bar by default,
  // so "make the contents last N beats instead" is exactly slowing the
  // pattern by N/beats - no new engine primitive, and a 4-beat pattern in
  // an 8-beat bar comes out repeating twice per bar for free, which is the
  // tracker behaviour decision 3 asks for.
  //
  // Declared length is checked against the contents by analyze.js before
  // this runs, so by here N really is the number of beats the tokens add
  // up to, and an unweighted token really is one beat.
  //
  // An arrangement already has its absolute per-element placement built in
  // by buildArrangementPattern()/timeline() above - applying a SECOND,
  // single-factor slow() on top would double-scale it, which is exactly
  // the problem ADR 0003 section 1.2 describes ("a single slow factor
  // cannot express a timeline").
  if (!isArrangement && vstmt.metricBeats != null && vstmt.metricBeats !== ctx.barBeats) {
    chain.notePattern = chain.notePattern.slow(vstmt.metricBeats / ctx.barBeats);
  }
  chain.notePattern = applyTransforms(chain.notePattern, transformNodes);
  // sync/ring go straight through song.js's VoiceChain methods, whose
  // toPattern() accepts a bare 0/1 constant or a mini-notation string (see
  // docs/dsl-grammar-spec.md's "Sync/ring" section).
  if (syncValue != null) chain.sync(typeof syncValue === 'string' ? expandHolds(syncValue) : syncValue);
  if (ringValue != null) chain.ring(typeof ringValue === 'string' ? expandHolds(ringValue) : ringValue);

  const finalInstrument = finalizeInstrument(baseInstrument, { vib: vibOverride, slide: slideOverride, gate: gateOverride });
  if (finalInstrument) {
    chain.inst(finalInstrument);
    // adsr is a static value, not one of the per-frame table kinds
    // compile.js samples off an instrument, so it goes through the chain's
    // .adsr().
    if (finalInstrument.adsr) chain.adsr(...finalInstrument.adsr);
    // Same for gate: compile.js only reads chain.gatePattern (set by
    // .gate()), never inst.gate directly.
    if (finalInstrument.gate != null) chain.gate(finalInstrument.gate);
  }

  // `insts` - GoatTracker's per-row instrument column (see this function's
  // instsPatternSource comment above and docs/dsl-grammar-spec.md's "Per-note
  // instrument switching" section). The pattern's bare-word tokens are
  // instrument NAMES, resolved against `instruments` (the SAME map/rules
  // resolveNamedInstrument() already uses for `voice N NAME ...` - an
  // instrument referenced here must exist somewhere in the source, same
  // "declare it with an `instrument NAME` block" error otherwise). Sampled by
  // compile.js at each note's own onset (cycle-position clock - this is a
  // property of the note timeline, not a frames-since-onset tracker table),
  // and held for that note's whole duration; when no insts value overlaps a
  // given onset, compile.js falls back to this chain's own static
  // finalInstrument (attached above), so `insts` never needs to cover every
  // note.
  if (instsPatternSource != null) {
    const namePattern = parseMini(expandHolds(instsPatternSource));
    const resolvedInstPattern = mapPatternValues(namePattern, (name) =>
      typeof name === 'string' ? resolveNamedInstrument(name, instruments, vstmt) : null,
    );
    chain.instPattern(resolvedInstPattern);
    // Source-highlighting anchor for `insts`, set directly on the chain
    // exactly like chain.noteSourceSegments above (see song.js's
    // VoiceChain constructor comment) - null (never highlighted) when
    // instsSegments came back null, e.g. this voice's `insts` string was
    // built some other way than a plain STRING/backtick literal.
    chain.instSourceSegments = instsSegments;
    chain.instSourceText = instsPatternSource;
  }

  // `fx` - GoatTracker's per-row command/command-data columns (see this
  // function's fxPatternSource comment above, src/dsl/fx.js for the command
  // set and its gplay.c provenance, and compile.js's fx block for the
  // runtime). Bare-word tokens are effect tokens, parsed here (so a typo is
  // a compile error at the offending token, not a silently ignored note)
  // into the plain { op, a, b } objects compile.js consumes.
  if (fxPatternSource != null) {
    const fxTokens = parseMini(expandHolds(fxPatternSource));
    const resolvedFxPattern = mapPatternValues(fxTokens, (token) => {
      try {
        return parseFxToken(token);
      } catch (e) {
        throw new Error(`line ${vstmt.line}: voice ${vstmt.index}: ${e.message}`);
      }
    });
    chain.fxPattern(resolvedFxPattern);
  }

  for (const f of regFields) {
    const addr = resolveRegTarget(f.target, vstmt.index, f.line);
    chain.reg(addr, expandHolds(f.resolvedSource));
  }

  return chain;
}

// === Filter =====================================================================

const SIGNALS = { sine, saw, square, tri, rand, perlin };

function buildSignalExprPattern(expr) {
  const base = SIGNALS[expr.signal];
  if (!base) {
    throw new Error(`line ${expr.line}: unknown signal '${expr.signal}' (expected one of sine/saw/square/tri/rand/perlin).`);
  }
  let p = base;
  if (expr.range) p = p.range(expr.range.lo, expr.range.hi);
  for (const speed of expr.speeds) {
    if (speed.op === '*') p = p.fast(speed.value);
    else if (speed.op === '/') p = p.slow(speed.value);
    else if (speed.op === '>>') p = p.shiftedBy(speed.value);
  }
  return p;
}

// routeObjectFromValues(values) -> the { voice0, voice1, voice2 } shape
// FilterChain.route()'s single-Pattern-argument path expects when sampled -
// same claimed-set logic as song.js's own route(...voiceIndices) overload.
function routeObjectFromValues(values) {
  const claimed = new Set(values);
  return { voice0: claimed.has(0), voice1: claimed.has(1), voice2: claimed.has(2) };
}

// buildLeafPattern(fieldKind, leaf) -> a Pattern for one NumberListExpr or
// SignalExpr leaf (either a field's flat value, or one switch-case's
// value), dispatching on which of the four filter fields this leaf belongs
// to:
//  - cutoff/resonance: SignalExpr -> a real numeric ramp (sine/saw/...),
//    NumberListExpr -> a constant number (resonance's flat form only,
//    cutoff never produces this leaf shape - see checkResonanceExpr/the
//    cutoff-rejects-NumberList check in buildFilterField).
//  - mode: SignalExpr -> a constant mode-name string (Range/Speeds are
//    always empty here, enforced by checkModeExpr - mode names aren't
//    numeric signals, so this deliberately does NOT go through
//    buildSignalExprPattern/SIGNALS).
//  - route: NumberListExpr -> a constant { voice0, voice1, voice2 } object
//    (mode never reaches NumberListExpr; route never reaches SignalExpr -
//    both enforced by checkRouteExpr/checkModeExpr).
function buildLeafPattern(fieldKind, leaf) {
  if (leaf.type === 'NumberListExpr') {
    return pure(fieldKind === 'route' ? routeObjectFromValues(leaf.values) : leaf.values[0]);
  }
  if (fieldKind === 'mode') return pure(leaf.signal);
  return buildSignalExprPattern(leaf);
}

// buildSwitchBlockPattern(fieldKind, node) -> a single
// continuous((cycleIndex, posFrac) => ...) callback dispatching on
// cycleIndex to whichever case matches, then delegating to that case's own
// leaf pattern's queryAt() (query.js's continuous() accepts exactly this
// kind of per-(cycleIndex, posFrac) callback; see the grammar spec's
// "Filter" section). Works for all four filter fields via buildLeafPattern.
// Cases are tried in source
// order; the first matching case wins (an `else` case, if present, is
// expected last but that's a style convention, not enforced here).
function buildSwitchBlockPattern(fieldKind, node) {
  const cases = node.cases.map((c) => ({ cond: c.cond, pattern: buildLeafPattern(fieldKind, c.value) }));
  return continuous((cycleIndex, posFrac) => {
    for (const c of cases) {
      if (
        (c.cond.op === '<' && cycleIndex < c.cond.value) ||
        (c.cond.op === '>=' && cycleIndex >= c.cond.value) ||
        c.cond.op === 'else'
      ) {
        return c.pattern.queryAt(cycleIndex, posFrac);
      }
    }
    return undefined;
  });
}

// buildFilterFieldPattern(fieldKind, expr) -> the single shared entry point
// used by all four filter fields: expr is either a *SwitchBlock (time-
// varying, dispatched per cycle) or a flat leaf (SignalExpr/NumberListExpr,
// a constant Pattern via buildLeafPattern).
function buildFilterFieldPattern(fieldKind, expr) {
  return expr.type === 'SwitchBlock' ? buildSwitchBlockPattern(fieldKind, expr) : buildLeafPattern(fieldKind, expr);
}

function applyFilterBlockNode(filterChainBuilder, node) {
  for (const f of node.fields) {
    switch (f.type) {
      case 'cutoff':
        filterChainBuilder.cutoff(buildFilterFieldPattern('cutoff', f.expr));
        break;
      case 'resonance':
        filterChainBuilder.resonance(buildFilterFieldPattern('resonance', f.expr));
        break;
      case 'mode':
        filterChainBuilder.mode(buildFilterFieldPattern('mode', f.expr));
        break;
      case 'route':
        // FilterChain.route() accepts a single Pattern argument for a
        // time-varying routing pattern (see song.js) - always used here,
        // even for a flat constant, since buildFilterFieldPattern always
        // returns a Pattern (a pure() one for the constant case).
        filterChainBuilder.route(buildFilterFieldPattern('route', f.expr));
        break;
      default:
        throw new Error(`line ${f.line}: unknown filter field '${f.type}'.`);
    }
  }
}

// resolveVoiceStmtSigils(vstmt, patterns) -> a shallow copy of vstmt with
// every `notes`/`reg` field's $name reference already substituted, using
// whatever `patterns` contains AT THIS EXACT POINT in the program.
//
// This exists because buildVoiceChain() (below) only actually runs in a
// SECOND pass over voiceStmts, after every statement has been scanned (so
// tempo is known before song() is constructed) - if $name substitution also
// happened lazily in that second pass, a `pattern` declared anywhere in the
// whole file (even AFTER this voice) would already be sitting in the
// `patterns` map by the time buildVoiceChain ran, silently defeating the
// spec's "forward references are a compile error" rule. Resolving eagerly
// here, at the point each VoiceStmt is actually encountered during the
// single top-to-bottom scan, is what makes that check real.
function resolveVoiceStmtSigils(vstmt, patterns, analysisPatterns, barBeats, beatUnit) {
  if (vstmt.mode === 'short') {
    return {
      ...vstmt,
      // ADR 0003 D4: arrangement mode (vstmt.notes === null), and a bare
      // sigil naming an ARRANGEMENT pattern (resolveNotesValueStructural
      // returns null for that shape) both have no mini-notation Pattern to
      // build here - buildVoiceChain's isArrangement/bareArrangementRef
      // path resolves those structurally against `patterns` directly.
      resolvedNotesPattern: vstmt.notes != null ? resolveNotesValueStructural(vstmt.notes, patterns, analysisPatterns, barBeats, beatUnit, vstmt.line, 'pattern') : null,
      resolvedNotesHighlightPattern: vstmt.notes != null ? resolveNotesValueStructural(vstmt.notes, patterns, analysisPatterns, barBeats, beatUnit, vstmt.line, 'highlightPattern') : null,
      // `insts:"..."` is a mini-notation pattern like any other, so a $name
      // splice has to resolve inside it too - an orderlist-shaped `insts`
      // replays the same few blocks over and over, which is exactly what
      // named patterns exist for. Unlike `fx` (never highlighted), `insts`
      // gets a segments-based highlighting anchor so an editor can show
      // which instrument a voice is playing (a.docStart is this inline
      // attr's quoted string's position).
      inlineAttrs: vstmt.inlineAttrs.map((a) => {
        if (a.name === 'fx' && a.isPattern) {
          return { ...a, source: spliceSigils(a.source, patterns, vstmt.line, null).source };
        }
        if (a.name === 'insts' && a.isPattern) {
          const resolved = spliceSigils(a.source, patterns, vstmt.line, a.docStart ?? null);
          return { ...a, source: resolved.source, segments: resolved.segments };
        }
        return a;
      }),
    };
  }
  const fields = vstmt.fields.map((f) => {
    if (f.type === 'notes') {
      return {
        ...f,
        resolvedPattern: resolveNotesValueStructural(f, patterns, analysisPatterns, barBeats, beatUnit, f.line, 'pattern'),
        resolvedHighlightPattern: resolveNotesValueStructural(f, patterns, analysisPatterns, barBeats, beatUnit, f.line, 'highlightPattern'),
      };
    }
    if (f.type === 'fx') return { ...f, resolvedSource: spliceSigils(f.source, patterns, f.line, null).source };
    // Segments-based highlighting anchor for `insts` (see the short-mode
    // branch's comment); f.docStart is this field's quoted string's position.
    if (f.type === 'insts') {
      const resolved = spliceSigils(f.source, patterns, f.line, f.docStart ?? null);
      return { ...f, resolvedSource: resolved.source, segments: resolved.segments };
    }
    // reg/sync/ring values are never highlighted, so only their resolved
    // TEXT matters here - spliceSigils(..., null) still resolves $name
    // splices correctly, it just skips building (and immediately discards)
    // segments, which is exactly what's wanted for these two.
    if (f.type === 'reg') return { ...f, resolvedSource: spliceSigils(f.source, patterns, f.line, null).source };
    if ((f.type === 'sync' || f.type === 'ring') && f.value.kind === 'pattern') {
      return { ...f, value: { ...f.value, resolvedSource: spliceSigils(f.value.source, patterns, f.line, null).source } };
    }
    return f;
  });
  return { ...vstmt, fields };
}

// === Program ====================================================================

// interpret(ast, options) -> Song. Walks the top-level statement list
// exactly once, in source order - see this file's header comment for why
// that order is what makes forward-referencing a named pattern a compile
// error for free.
//
// options.onDiagnostic(d) is called once per non-fatal diagnostic
// (severity 'warning' or 'info') the semantic pass produced - the
// declared-length/period report, the migration hints, the four
// known-defect warnings from ADR 0001 section 6. They are also attached to
// the returned builder as `.dslDiagnostics`, so a caller that has no
// callback (the editor, a test) can still read them off the result.
// ERROR-severity diagnostics are thrown, exactly like every other compile
// error in this pipeline, so nothing downstream has to learn a second
// failure convention.
export function interpret(ast, options = {}) {
  const analysis = analyzeProgram(ast);
  const fatal = analysis.diagnostics.find((d) => d.severity === 'error');
  if (fatal) throw new Error(`line ${fatal.line}: ${fatal.message}`);
  const softDiagnostics = analysis.diagnostics.filter((d) => d.severity !== 'error');
  if (options.onDiagnostic) for (const d of softDiagnostics) options.onDiagnostic(d);

  const patterns = new Map();
  const tables = new Map();
  const instruments = new Map();
  const voiceStmts = [];
  let filterBlockNode = null;
  let tempoNode = null;
  let targetNode = null;
  let chipModel = null;
  let tempoResolvedSource = null;
  let beats = 4;
  // `offset N` on the tempo statement: the whole song's pattern clock,
  // shifted by N beats. See the conversion below for the units.
  let offsetBeats = 0;

  for (const stmt of ast.statements) {
    switch (stmt.type) {
      case 'TempoStmt':
        tempoNode = stmt;
        // ADR 0003 D2: `meter N/D` and `beats N` both set the bar length in
        // beats; only `meter` also carries the beat unit (D). analyze.js has
        // already rejected combining the two on one statement.
        if (stmt.meter != null) beats = stmt.meter.numerator;
        else if (stmt.beats != null) beats = stmt.beats;
        if (stmt.offset != null) offsetBeats = stmt.offset;
        // Resolved right here too, for the same reason as VoiceStmt below -
        // in every example this is the first statement anyway, but nothing
        // in the grammar requires that, so this stays correct regardless.
        if (stmt.value.kind === 'pattern') {
          tempoResolvedSource = spliceSigils(stmt.value.source, patterns, stmt.line, null).source;
        }
        break;
      case 'TargetStmt':
        // ADR 0002. At most one per song: two would leave the drive rate
        // ambiguous, and a mid-song drive-rate change is an explicit
        // non-goal of that ADR (its section 2).
        if (targetNode) {
          throw new Error(`line ${stmt.line}: a song can have at most one target block (the first is on line ${targetNode.line}).`);
        }
        targetNode = stmt;
        break;
      case 'ChipStmt':
        // ADR 0003 D15. At most one per song, same reasoning as target:
        // two would leave the intended chip model ambiguous.
        if (chipModel != null) {
          throw new Error(`line ${stmt.line}: a song can have at most one chip statement.`);
        }
        chipModel = stmt.model;
        break;
      case 'PatternDecl': {
        // ADR 0003 D4: an arrangement declaration (`pattern verse = $a $b`,
        // stmt.elements set) has no mini-notation source text at all - it
        // is stored as structural elements, resolved into a Pattern by
        // buildArrangementPattern() below, at each referencing voice's own
        // build time (patterns can nest, so there is no single flat text
        // to splice here the way a string pattern's $refs are).
        if (stmt.elements != null) {
          patterns.set(stmt.name, { elements: stmt.elements });
          break;
        }
        // ADR 0003 D5: $name is a structural reference - stmt.source
        // keeps its literal `$name` tokens, resolved by parseMini's `refs`
        // option against patterns declared EARLIER in this same forward
        // scan (patternObjectRefs/patternObjectHighlightRefs below), not
        // spliced as text. `docOffset: stmt.docStart` is what makes every
        // Hap this pattern's own Pattern object ever produces carry its
        // ABSOLUTE document position from here on - a later $ref to this
        // pattern reuses that Pattern object as-is, no further offset
        // math needed (see resolveNotesValueStructural's own comment).
        //
        // ADR 0003 D3: `each N/D` has no meaning once referenced by name
        // elsewhere, so it has to be baked into literal per-token weights
        // NOW, the same way a hold token is folded before point of use
        // elsewhere in this function. analysis.beatUnit (not the
        // as-yet-unresolved local `beats`/beatUnit below) is used because
        // analyze.js has already scanned the whole statement list for the
        // song's meter.
        // ADR 0003 D14: note names are checked in named declarations too,
        // not just inline voice strings, on the raw text before hold/each
        // expansion touches it.
        checkNoteNames(stmt.source, stmt.line);
        const weightRefs = (name) => patternBeatsOf(name, analysis.patterns, analysis.barBeats);
        const eachExpanded = stmt.each != null
          ? expandEach(stmt.source, stmt.each.denominator, analysis.beatUnit, weightRefs)
          : expandHolds(stmt.source, analysis.beatUnit, weightRefs);
        const patternObjectRefs = (name) => patterns.get(name)?.pattern;
        const patternObjectHighlightRefs = (name) => patterns.get(name)?.highlightPattern;
        const docOffset = stmt.docStart ?? 0;
        patterns.set(stmt.name, {
          source: eachExpanded,
          pattern: parseMini(eachExpanded, { beatUnit: analysis.beatUnit, docOffset, refs: patternObjectRefs, weightRefs }),
          highlightPattern: parseMiniForHighlight(eachExpanded, { beatUnit: analysis.beatUnit, docOffset, refs: patternObjectHighlightRefs, weightRefs }),
        });
        break;
      }
      case 'TableDecl': {
        // Expanded at declaration time (against the tables declared BEFORE
        // it), so a table composed from other tables - `table full = $attack
        // $sustain` - is stored already-flat and every later reference to it
        // is a plain array lookup.
        //
        // A shaped ramp's implied loop point (ADR 0003 D9) survives a shared
        // `table NAME = ...` declaration: `tables` stores { values,
        // loopIndex } per name (loopIndex null when nothing in the row
        // implied one), so a later `$NAME` reference can fold NAME's
        // loopIndex into the referencing row - see expandTableRefs for the
        // precedence rule in a mixed row of $refs and inline terms.
        const expanded = expandTableRefs(stmt.tokens, tables, stmt.line);
        tables.set(
          stmt.name,
          Array.isArray(expanded) ? { values: expanded, loopIndex: null } : { values: expanded.tokens, loopIndex: expanded.loopIndex },
        );
        break;
      }
      case 'InstrumentBlock':
        instruments.set(stmt.name, buildInstrumentFromBlockFields(stmt.fields, tables));
        break;
      case 'VoiceStmt': {
        // metricBeats comes from the semantic pass (analyze.js), which
        // walked the SAME statement list in the same order, so the Nth
        // voice statement here is the Nth entry there. Resolved before
        // splicing, since the metric length belongs to the NAMED pattern a
        // voice references and a splice erases the name.
        const metricBeats = analysis.voices[voiceStmts.length]?.metricBeats ?? null;
        voiceStmts.push({ ...resolveVoiceStmtSigils(stmt, patterns, analysis.patterns, analysis.barBeats, analysis.beatUnit), metricBeats });
        break;
      }
      case 'FilterBlock':
        filterBlockNode = stmt;
        break;
      default:
        throw new Error(`line ${stmt.line}: unrecognized top-level statement '${stmt.type}'.`);
    }
  }

  // "exactly one clock statement" is checked by analyze.js (ADR 0001
  // decision 11), which has already thrown by here if it was violated;
  // this is the belt-and-braces guard for a caller that built an AST by
  // hand rather than by parsing.
  if (!tempoNode) {
    throw new Error('a song needs exactly one clock statement, and this one has none - add either `tempo 120 beats 4` (BPM) or `speed 6 beats 4` (frames per beat).');
  }

  // `speed N` is N FRAMES PER BEAT, the exact, frame-native spelling of the
  // same clock `tempo N` spells in BPM (ADR 0001's D4). One beat of B BPM
  // lasts 50*60/B = 3000/B frames on a PAL SID, so the two are related by
  // bpm = 3000 / framesPerBeat, exactly - and that is the point: at
  // `speed 6 beats 8` a bar is exactly 48 frames, where `tempo 132 beats 8`
  // computes 181.8 and has to round. Converting here (rather than teaching
  // song.js a second clock unit) keeps everything downstream on the one
  // BPM-shaped input it has always taken.
  // ADR 0002: a "frame" is one play call, and how many of those happen per
  // second is whatever the `target` block
  // declared (default PAL 985248/19705 = exactly 50.00, so every song
  // without a target block converts through the identical 3000 below).
  // Resolved before the tempo conversion because BOTH of the conversions
  // under it (frames-per-beat -> BPM, and `offset` beats -> frames) are
  // conversions between musical time and FRAME time, and the exchange rate
  // between frame time and wall-clock time is exactly what this declares.
  const drive = resolveTarget(targetNode);
  const FRAMES_PER_MINUTE_BASE = drive.driveHz * 60;

  const toBpm = (v) => (tempoNode.unit === 'framesPerBeat' ? (v > 0 ? FRAMES_PER_MINUTE_BASE / v : 0) : v);
  const tempoArg = tempoNode.value.kind === 'pattern'
    ? (tempoNode.unit === 'framesPerBeat'
      ? mapPatternValues(parseMini(tempoResolvedSource), (v) => (typeof v === 'number' ? toBpm(v) : v))
      : parseMini(tempoResolvedSource))
    : toBpm(tempoNode.value.value);

  // ADR 0003 D17: the frame grid is reported, and a misfitting subdivision
  // is named rather than left for the author to discover by ear. Only for
  // a CONSTANT tempo (a tempo-as-pattern song's frames-per-bar varies cycle
  // to cycle, which is not covered yet). Uses the same rounding as the
  // compiler's frames-per-bar clock, so this reports exactly what compile()
  // will produce and changes no emitted frame itself.
  if (tempoNode.value.kind !== 'pattern' && tempoArg > 0) {
    const bpm = tempoArg;
    const framesPerBeatExact = FRAMES_PER_MINUTE_BASE / bpm;
    const framesPerBar = Math.round(framesPerBeatExact * beats);
    const actualBpm = (FRAMES_PER_MINUTE_BASE * beats) / framesPerBar;
    softDiagnostics.push(
      diag(
        'info',
        tempoNode.line,
        `meter ${beats}/${analysis.beatUnit}, tempo ${formatBeatsForDiag(bpm)}, ${drive.video} ${drive.driveHz.toFixed(2)} fps -> ` +
          `${formatBeatsForDiag(framesPerBar)} frames/bar, ${formatBeatsForDiag(framesPerBar / beats)} frames/beat.`,
      ),
    );
    if (Math.abs(actualBpm - bpm) > 0.005) {
      const nearestSpeed = Math.round(framesPerBeatExact);
      softDiagnostics.push(
        diag(
          'warning',
          tempoNode.line,
          `tempo ${formatBeatsForDiag(bpm)} rounds to ${framesPerBar} frames/bar, which is ${actualBpm.toFixed(2)} BPM. ` +
            `\`speed ${nearestSpeed}\` is the nearest exact spelling.`,
        ),
      );
    }
    if (analysis.smallestDurationBeats != null) {
      const smallestFrames = framesPerBar * (analysis.smallestDurationBeats / beats);
      if (Math.abs(smallestFrames - Math.round(smallestFrames)) > 1e-6) {
        // The smallest integer frames-per-beat (`speed`) that makes this
        // exact subdivision land on a whole frame: framesPerSubdivision =
        // speed * smallestDurationBeats must be an integer, so speed must
        // be a multiple of smallestDurationBeats' own reduced denominator.
        // Computed, not hand-waved (the ADR's own requirement): scan
        // multiples of that denominator outward from the CURRENT exact
        // frames/beat and take the closest, so the suggestion changes the
        // song's feel as little as possible while actually being exact.
        const denom = exactDenominator(analysis.smallestDurationBeats);
        const center = framesPerBeatExact;
        let nearestExactSpeed = denom;
        let nearestDelta = Math.abs(denom - center);
        for (let k = 2; k * denom < center * 4 + denom; k += 1) {
          const candidate = k * denom;
          const delta = Math.abs(candidate - center);
          if (delta < nearestDelta) {
            nearestExactSpeed = candidate;
            nearestDelta = delta;
          }
        }
        const subdivisionLabel = beatFractionLabel(analysis.smallestDurationBeats, analysis.beatUnit);
        softDiagnostics.push(
          diag(
            'warning',
            tempoNode.line,
            `the smallest duration in this song is ${subdivisionLabel} = ${smallestFrames.toFixed(2)} frames, which is not a whole number ` +
              `of frames; onsets will jitter by up to 1 frame. \`speed ${nearestExactSpeed}\` gives exactly ${(nearestExactSpeed * analysis.smallestDurationBeats).toFixed(0)} frames per ${subdivisionLabel} at this meter.`,
          ),
        );
      }
    }
  }

  // `offset` is declared in BEATS - the song's own unit, the same one
  // `beats N` counts - but song.js/compile.js take the shift in FRAMES
  // (startOffsetFrames), so it is converted here, once, against the song's
  // base tempo. A positive offset DELAYS the song: real frame N is played
  // as compiler frame N - startOffsetFrames, so the pattern clock reaches
  // cycle 0 that many frames late.
  //
  // Why this exists: a converted GoatTracker song has to be able to declare
  // "this tune's first row sounds one row after the play routine starts"
  // (gplay.c seeds its tick counter from the compiled default tempo before
  // the sequencer ever fetches a row) WITHOUT rotating its own note data to
  // fake it: rotating would move every pattern boundary off the cycle grid,
  // so a generated `<...>` alternation would no longer line up with the
  // song's real patterns. It is equally useful by hand for a tune that
  // wants a pickup beat.
  //
  // A fractional offset is allowed and meaningful (a lead-in row is not
  // always a whole beat of the song's final tempo); it is rounded to a whole
  // frame here because frames are what the SID clock actually has.
  // FRAMES_PER_MINUTE_BASE is the drive rate resolved above, in frames per
  // minute (ADR 0002).
  const baseBpm = tempoNode.value.kind === 'pattern'
    ? tempoArg.queryAt(0, 0)
    : toBpm(tempoNode.value.value);
  const startOffsetFrames = offsetBeats === 0 || !(baseBpm > 0)
    ? 0
    : Math.round((offsetBeats * FRAMES_PER_MINUTE_BASE) / baseBpm);

  const builder = song(tempoArg, {
    beatsPerCycle: beats,
    startOffsetFrames,
    // ADR 0002: carried on the song so the compiler's frames-per-bar
    // arithmetic and any renderer read the same declared rate instead of
    // assuming PAL 50.
    driveCycles: drive.driveCycles,
    cpuClockHz: drive.cpuClockHz,
    // ADR 0003 D15: which SID model this song declares itself against, or
    // null for "no opinion" (the consumer's own default applies).
    chipModel,
  });

  for (const vstmt of voiceStmts) {
    const chain = buildVoiceChain(vstmt, { instruments, tables, barBeats: beats, patterns, analysisPatterns: analysis.patterns });
    builder.voice(vstmt.index, chain);
  }

  if (filterBlockNode) {
    builder.filter((f) => applyFilterBlockNode(f, filterBlockNode));
  }

  // See interpret()'s own doc comment: warnings/info ride along on the
  // builder so a caller without an onDiagnostic callback can still find
  // them, and so nothing has to be printed from inside a compiler.
  builder.dslDiagnostics = softDiagnostics;
  builder.dslPeriodBeats = analysis.periodBeats;
  return builder;
}
