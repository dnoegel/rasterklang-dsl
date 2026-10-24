// Structural scanner for mini-notation strings: how many BEATS a pattern's
// contents add up to, where its hold tokens are, and how to rewrite those
// holds into weights the pattern engine understands.
//
// Why this lives in rasterklang-dsl and not in the pattern engine: since
// 3.0.0 a pattern may DECLARE its length (`pattern bass 4 = "c4 ~ c4 ~"`,
// ADR 0001's D1), and "does the content actually add up to the declared
// length" is a compile-time question about the SOURCE, answerable without
// building a single Hap. Making it a property of the grammar layer means
// the check, its error message, and the length arithmetic itself have one
// home that both the JS interpreter and Go consumers (via the port in
// pkg/dsl/minilength.go), such as DSL generators, can rely on.
//
// This is deliberately NOT a second mini-notation parser in the sense of
// producing patterns: it computes structure only (steps, weights,
// nesting), and it mirrors rasterklang-score's pattern/miniNotation.js
// scanning rules step for step - same atoms (`~`/`-` rest, word, `[
// ]` group, `< >` alternation, `{ }` polymetre, `a|b` choice), same
// modifiers (`*n` `/n` `@n` `!n` `?p` `(p,s,r)`), same "a `,` stacks"
// reading. Any divergence between the two is a bug in this file.
//
// `/N[.t]` (ADR 0003 D1) is a duration in beats, not the pre-4.0.0 slow
// operator: `/K` is `beatUnit/K` beats, `.` multiplies by 1.5 (dotted),
// `t` by 2/3 (triplet). It sets a step's weight exactly the way `@n` does
// - `/8` and `@0.5` mean the same thing in a 4/4-beat context - so `/N`
// and `@n` in one pattern is rejected (one is a duration, the other a
// proportional weight, and mixing them has no single sensible reading).
//
// The one token this file knows that miniNotation.js's 2.x grammar did not
// is `_`, the hold (ADR 0001's D6): a tracker's "empty row", i.e. the note
// keeps ringing for one more step rather than being cut by a rest.

// WEIGHT_EPSILON - the tolerance for "these two beat counts are the same".
// Weights are decimal source text (`@0.25`), so exact float equality would
// report a 32 x 0.25 melody as 7.999999999999999 beats against a declared
// 8. One part in a thousand of a beat is far finer than anything the 50Hz
// frame clock can express and far coarser than float noise.
const WEIGHT_EPSILON = 1e-6;

export const HOLD_TOKEN = '_';

// Default beat unit (the meter's denominator D, `meter N/D`) a `/N` duration
// token converts against when the caller has not threaded the song's actual
// meter through - see miniNotation.js's MiniParser constructor comment for
// the same gap on the pattern-engine side.
const DEFAULT_BEAT_UNIT = 4;

export function beatsEqual(a, b) {
  return Math.abs(a - b) < WEIGHT_EPSILON;
}

// === Scanner =================================================================

// A parsed step: `text` is the atom exactly as written (word/rest/hold, or
// a rendered group), `weight` its `@n` (default 1), `replicate` its `!n`
// (default 1), `suffix` every other modifier verbatim (`*2`, `?0.3`,
// `(3,8)`) so re-rendering is lossless.
class Step {
  constructor({ kind, text, weight, replicate, suffix, seqs, entries }) {
    this.kind = kind; // 'atom' | 'hold' | 'rest' | 'group' | 'alt' | 'poly' | 'choice'
    this.text = text;
    this.weight = weight;
    this.replicate = replicate;
    this.suffix = suffix;
    this.seqs = seqs ?? null; // group/poly: one entry per ','-stacked sequence
    this.entries = entries ?? null; // alt: its elements, one per bar/orderlist slot
  }
}

class MiniScanner {
  constructor(str, beatUnit = DEFAULT_BEAT_UNIT, defaultWeight = 1, refs = null, barBeats = null) {
    this.str = str;
    this.pos = 0;
    this.beatUnit = beatUnit;
    // barBeats (ADR 0003 D6): a bar is this many beats, needed only to
    // check a standalone `|`'s claimed position. null means "don't check
    // bar lines" (every caller except the dedicated checker below) -
    // scanMini()'s other callers (contentBeats, expandHolds, ...) don't
    // need this and a `|` there is scanned and ignored, same as always.
    this.barBeats = barBeats;
    // ADR 0003 D3 (`each 1/D`): the weight a step gets when it carries no
    // explicit `@n`/`/N` of its own. 1 unless a caller
    // is expanding an `each`-declared pattern, in which case this is
    // beatUnit/D - see expandEach() below.
    this.defaultWeight = defaultWeight;
    // refs (ADR 0003 D5): (name) -> beats | null/undefined, the same
    // resolver shape analyze.js's resolveRefLength() uses - a $name
    // reference's WEIGHT is the referenced pattern's own beats, not the
    // default 1. Mirrors miniNotation.js's MiniParser.refs (which resolves
    // to a Pattern, not a number) one for one.
    this.refs = refs;
    // Position of the first `/N` duration and first `@n` weight seen so
    // far - the "mixed /N and @n in one pattern" check (D1, D14). Mirrors
    // miniNotation.js's MiniParser one for one; any divergence is a bug in
    // one of the two files (this file's own header comment).
    this.sawDurationAt = null;
    this.sawWeightAt = null;
    // Every /N[.t] duration's beats value seen so far (ADR 0003 D17): the
    // smallest one is what the compiler checks against the frame grid for
    // subdivision jitter. Empty for a pattern using no explicit durations.
    this.durationsSeen = [];
  }

  eof() {
    return this.pos >= this.str.length;
  }

  peek() {
    return this.pos < this.str.length ? this.str[this.pos] : undefined;
  }

  skipWs() {
    while (!this.eof() && /\s/.test(this.peek())) this.pos += 1;
  }

  fail(msg) {
    throw new Error(`${msg} in mini-notation "${this.str}"`);
  }

  expect(ch) {
    if (this.peek() !== ch) this.fail(`expected '${ch}' at position ${this.pos}`);
    this.pos += 1;
  }

  parseNumberLiteral() {
    const start = this.pos;
    if (this.peek() === '-') this.pos += 1;
    while (!this.eof() && /[0-9.]/.test(this.peek())) this.pos += 1;
    const raw = this.str.slice(start, this.pos);
    if (raw === '' || raw === '-') this.fail(`expected a number at position ${start}`);
    return { value: Number(raw), text: raw };
  }

  tryParseNumberLiteral() {
    if (this.eof()) return null;
    const ch = this.peek();
    const next = this.str[this.pos + 1];
    if (/[0-9.]/.test(ch) || (ch === '-' && next && /[0-9]/.test(next))) return this.parseNumberLiteral();
    return null;
  }

  // parseDuration() -> beats, for a `/` modifier already consumed. Shape is
  // `INT ('.' | 't')?` - a dedicated digit scan, not parseNumberLiteral,
  // because that would swallow the dot of `/4.` as a decimal point. Mirrors
  // miniNotation.js's MiniParser.parseDuration() exactly.
  parseDuration() {
    const start = this.pos;
    let digits = '';
    while (!this.eof() && /[0-9]/.test(this.peek())) {
      digits += this.peek();
      this.pos += 1;
    }
    if (digits === '') this.fail(`expected a duration (e.g. /4, /8., /4t) at position ${start}`);
    const denom = Number(digits);
    let beats = this.beatUnit / denom;
    let text = `/${digits}`;
    if (this.peek() === '.') {
      this.pos += 1;
      beats *= 1.5;
      text += '.';
    } else if (this.peek() === 't') {
      this.pos += 1;
      beats *= 2 / 3;
      text += 't';
    }
    return { value: beats, text };
  }

  failMixedDurationWeight() {
    this.fail(
      `mixes /N durations (first at position ${this.sawDurationAt}) with @N weights ` +
        `(first at position ${this.sawWeightAt}) - one is a duration and the other a proportional weight, use one or the other`,
    );
  }

  scanWordToken() {
    const start = this.pos;
    if (this.peek() === '-' && /[0-9]/.test(this.str[this.pos + 1] || '')) this.pos += 1;
    // ':' only ever appears as part of 'slideTo:' (ADR 0003 D12) - mirrors
    // miniNotation.js's scanWordToken exactly (this file's own invariant:
    // any divergence from the engine's own scanning rules is a bug here).
    while (!this.eof() && /[A-Za-z0-9.#:]/.test(this.peek())) this.pos += 1;
    const raw = this.str.slice(start, this.pos);
    if (raw === '') this.fail(`expected a step at position ${start}`);
    return raw;
  }

  // atom := rest | hold | word | group | alternation | polymetre
  parseAtomUnit() {
    this.skipWs();
    const ch = this.peek();
    if (ch === '~') {
      this.pos += 1;
      return new Step({ kind: 'rest', text: '~', weight: 1, replicate: 1, suffix: '' });
    }
    if (ch === '-' && !/[0-9]/.test(this.str[this.pos + 1] || '')) {
      this.pos += 1;
      return new Step({ kind: 'rest', text: '-', weight: 1, replicate: 1, suffix: '' });
    }
    if (ch === HOLD_TOKEN) {
      this.pos += 1;
      return new Step({ kind: 'hold', text: HOLD_TOKEN, weight: 1, replicate: 1, suffix: '' });
    }
    if (ch === '[' || ch === '{') {
      const close = ch === '[' ? ']' : '}';
      this.pos += 1;
      const seqs = [this.parseSequence([',', close])];
      this.skipWs();
      while (this.peek() === ',') {
        this.pos += 1;
        seqs.push(this.parseSequence([',', close]));
        this.skipWs();
      }
      this.expect(close);
      return new Step({ kind: ch === '[' ? 'group' : 'poly', text: null, weight: 1, replicate: 1, suffix: '', seqs });
    }
    if (ch === '<') {
      this.pos += 1;
      const entries = [];
      this.skipWs();
      while (!this.eof() && this.peek() !== '>') {
        entries.push(this.parseStep());
        this.skipWs();
      }
      this.expect('>');
      return new Step({ kind: 'alt', text: null, weight: 1, replicate: 1, suffix: '', entries });
    }
    // `$name` (ADR 0003 D5): a structural reference, weight = the
    // referenced pattern's own beats (declared, derived, or one bar if
    // proportional - resolved by the caller's `refs` function, the same
    // rule resolveRefLength() in analyze.js already applies for
    // arrangements). Only recognized when a resolver was given; otherwise
    // `$` falls through to scanWordToken(), which rejects it.
    if (ch === '$' && this.refs != null) {
      const start = this.pos;
      this.pos += 1;
      const nameStart = this.pos;
      while (!this.eof() && /[A-Za-z0-9_]/.test(this.peek())) this.pos += 1;
      const name = this.str.slice(nameStart, this.pos);
      if (name === '') this.fail(`expected a pattern name after '$' at position ${start}`);
      const beats = this.refs(name);
      if (beats == null) {
        this.fail(`pattern "$${name}" is not defined (declare it with \`pattern ${name} = "..."\` before this use - forward references aren't supported)`);
      }
      return new Step({ kind: 'ref', text: `$${name}`, weight: beats, replicate: 1, suffix: '' });
    }
    if (ch === undefined) this.fail('unexpected end of input while parsing a step');
    return new Step({ kind: 'atom', text: this.scanWordToken(), weight: 1, replicate: 1, suffix: '' });
  }

  // choice := atom ('|' atom)+ - kept as one step, its branches scanned
  // recursively so a hold inside a branch group is still seen.
  parseChoice() {
    const first = this.parseAtomUnit();
    if (this.peek() !== '|') return first;
    const branches = [first];
    while (this.peek() === '|') {
      this.pos += 1;
      branches.push(this.parseAtomUnit());
    }
    return new Step({ kind: 'choice', text: null, weight: 1, replicate: 1, suffix: '', entries: branches });
  }

  // step := atom modifier*
  parseStep() {
    const step = this.parseChoice();
    // each's default (see the constructor comment); '@'/'/' below overwrite
    // it with the step's own explicit weight/duration when present. Not
    // applied to a $name reference (ADR 0003 D5): its weight is already the
    // REFERENCED pattern's own beats (set in parseAtomUnit), which `each`
    // has no business overriding - `@n`/`/N` still can, explicitly.
    if (step.kind !== 'ref') step.weight = this.defaultWeight;
    for (;;) {
      const ch = this.peek();
      if (ch === '@') {
        const modStart = this.pos;
        this.pos += 1;
        step.weight = this.parseNumberLiteral().value;
        if (this.sawWeightAt === null) this.sawWeightAt = modStart;
        if (this.sawDurationAt !== null) this.failMixedDurationWeight();
      } else if (ch === '!') {
        this.pos += 1;
        const n = this.tryParseNumberLiteral();
        step.replicate = Math.trunc(n === null ? 2 : n.value);
      } else if (ch === '/') {
        const modStart = this.pos;
        this.pos += 1;
        step.weight = this.parseDuration().value;
        this.durationsSeen.push(step.weight);
        if (this.sawDurationAt === null) this.sawDurationAt = modStart;
        if (this.sawWeightAt !== null) this.failMixedDurationWeight();
      } else if (ch === '*') {
        this.pos += 1;
        step.suffix += '*' + this.parseNumberLiteral().text;
      } else if (ch === '?') {
        this.pos += 1;
        const n = this.tryParseNumberLiteral();
        step.suffix += `?${n === null ? '' : n.text}`;
      } else if (ch === '(') {
        const start = this.pos;
        const close = this.str.indexOf(')', this.pos);
        if (close === -1) this.fail(`unclosed '(' at position ${start}`);
        this.pos = close + 1;
        step.suffix += this.str.slice(start, this.pos);
      } else {
        break;
      }
    }
    // each's default weight (ADR 0003 D3) is itself a duration - every
    // unweighted token in an each-pattern is really "1/D" - so it counts
    // toward D17's smallest-subdivision scan the same as an explicit /N.
    if (this.defaultWeight !== 1 && step.weight === this.defaultWeight) this.durationsSeen.push(step.weight);
    return step;
  }

  // sequence := (step | barLine) (WS (step | barLine))*
  //
  // A `|` reached here (as opposed to inside parseStep()'s choice handling,
  // `a|b`, which only fires immediately after an atom with no intervening
  // whitespace) is always the STANDALONE form (ADR 0003 D6): a checked bar
  // line, purely an assertion with no effect on timing. Checked only when
  // this.barBeats is set (checkBarLines() below); every other caller of
  // scanMini() just skips over it, same as always.
  parseSequence(stopChars) {
    this.skipWs();
    const steps = [];
    let runningBeats = 0;
    while (!this.eof() && !stopChars.includes(this.peek())) {
      if (this.peek() === '|') {
        const barPos = this.pos;
        this.pos += 1;
        if (this.barBeats != null && !beatsEqual(runningBeats % this.barBeats, 0) && !beatsEqual(runningBeats % this.barBeats, this.barBeats)) {
          this.fail(`bar line '|' at position ${barPos} is off the grid - it lands at beat ${formatNumber(runningBeats)}, not a multiple of the ${formatNumber(this.barBeats)}-beat bar`);
        }
        this.skipWs();
        continue;
      }
      const step = this.parseStep();
      steps.push(step);
      runningBeats += step.weight * step.replicate;
      this.skipWs();
    }
    return steps;
  }

  // top-level := sequence (',' sequence)*   (',' stacks, same as a group)
  parseTop() {
    const seqs = [this.parseSequence([','])];
    while (this.peek() === ',') {
      this.pos += 1;
      seqs.push(this.parseSequence([',']));
    }
    if (!this.eof()) this.fail(`unexpected '${this.peek()}' at position ${this.pos}`);
    return seqs;
  }
}

// scanMini(source) -> the ','-stacked top-level sequences, each a Step[].
// Throws a plain Error (no line prefix - callers add their own) when the
// string is not scannable at all.
export function scanMini(source, beatUnit = DEFAULT_BEAT_UNIT, defaultWeight = 1, refs = null, barBeats = null) {
  return new MiniScanner(source, beatUnit, defaultWeight, refs, barBeats).parseTop();
}

// checkBarLines(source, barBeats, beatUnit, refs) (ADR 0003 D6): throws if
// any standalone `|` in `source` does not land on a multiple of `barBeats`
// beats, naming the beat offset it actually reached. Purely a check - it
// does not change how the pattern plays, so it is called for its
// side-effect (throwing), not its return value.
export function checkBarLines(source, barBeats, beatUnit = DEFAULT_BEAT_UNIT, refs = null) {
  scanMini(source, beatUnit, 1, refs, barBeats);
}

// smallestDurationBeats(source, beatUnit, defaultWeight, refs) -> beats |
// null (ADR 0003 D17): the smallest explicit /N[.t] duration used anywhere
// in source, in beats - including an each-pattern's own implied 1/D for
// every unweighted token (defaultWeight != 1). null when no duration is
// used at all (a plain proportional or @n-weighted pattern), which is not
// a misfit case D17 cares about.
export function smallestDurationBeats(source, beatUnit = DEFAULT_BEAT_UNIT, defaultWeight = 1, refs = null) {
  const scanner = new MiniScanner(source, beatUnit, defaultWeight, refs);
  scanner.parseTop();
  if (scanner.durationsSeen.length === 0) return null;
  return Math.min(...scanner.durationsSeen);
}

// proportionalFitInfo(source, beatUnit, refs) -> { tokenCount } | null (ADR
// 0003 D8): non-null only for a PURELY proportional pattern - no /N
// duration, no @n weight anywhere in it (each's own default doesn't count
// as "explicit", so an each-pattern is never "proportional" in this sense
// - it already declares its grid) - giving the token count proportional
// auto-fit divides the bar into, so the compiler can say so.
export function proportionalFitInfo(source, beatUnit = DEFAULT_BEAT_UNIT, refs = null) {
  const scanner = new MiniScanner(source, beatUnit, 1, refs);
  const seqs = scanner.parseTop();
  if (scanner.sawDurationAt !== null || scanner.sawWeightAt !== null) return null;
  const tokenCount = Math.max(...seqs.map((s) => s.reduce((n, step) => n + step.replicate, 0)), 0);
  return { tokenCount };
}

// === Lengths =================================================================

function sequenceBeats(steps) {
  let total = 0;
  for (const s of steps) total += s.weight * s.replicate;
  return total;
}

// contentBeats(source) -> how many beats one pass of this pattern's
// contents occupies, reading every unweighted token as one beat (ADR
// 0001's D1). Stacked sequences (`,`) sound simultaneously, so the stack's
// length is the longest of them, not their sum.
export function contentBeats(source, beatUnit = DEFAULT_BEAT_UNIT, refs = null) {
  const seqs = scanMini(source, beatUnit, 1, refs);
  return Math.max(...seqs.map(sequenceBeats), 0);
}

// stepBeats(step) -> one step's own content length, used for alternation
// elements (orderlist entries). A bare atom is one beat; a group is its
// contents; an alternation nested inside an element is that element's own
// business and counts as one beat here.
function stepBeats(step) {
  if (step.kind === 'group' || step.kind === 'poly') {
    return Math.max(...step.seqs.map(sequenceBeats), 0) * step.weight;
  }
  return step.weight;
}

// declaredLengthCandidates(source) -> every content length that a declared
// length has to match, so `pattern p 4 = "..."` can be checked with one
// comparison per candidate.
//
// Normally that is a single number: the top-level sequence's total (one
// per ','-stacked sequence, since each stacked layer plays the same span).
//
// The exception is ORDERLIST SHAPE - a pattern whose whole content is one
// `<...>` alternation. ADR 0001 decision 1 resolves alternation of metric
// elements as orderlist semantics: each element occupies its OWN declared
// length rather than a fixed bar each, which is exactly how a GoatTracker
// orderlist replays a column of patterns. So for `pattern song 4 = "<$a
// $b>"` the number 4 describes each element, not their sum, and every
// element is checked against it individually.
export function declaredLengthCandidates(source, beatUnit = DEFAULT_BEAT_UNIT, refs = null) {
  const seqs = scanMini(source, beatUnit, 1, refs);
  if (seqs.length === 1 && seqs[0].length === 1 && seqs[0][0].kind === 'alt' && seqs[0][0].weight === 1) {
    const entries = seqs[0][0].entries;
    return entries.map((e, i) => ({ beats: stepBeats(e), what: `orderlist element ${i + 1}` }));
  }
  return seqs.map((s, i) => ({
    beats: sequenceBeats(s),
    what: seqs.length === 1 ? 'contents' : `stacked layer ${i + 1}`,
  }));
}

// === Hold-token rewriting =====================================================

// A hold (`_`) means "the previous step keeps sounding for one more step's
// worth of time" - a tracker's empty row, and the exact opposite of a rest,
// which cuts the note (ADR 0001's D6). Folding each hold into the weight of
// the step it extends (`"c4 _ _ _"` -> `"c4@4"`) expresses precisely that
// with a construct the pattern engine has always had, so the hold token
// needs no engine change to sound right: one Hap, four beats long, gate
// open throughout.
//
// A hold after a REST extends the rest the same way (`"c4 _ ~ _"` ->
// `"c4@2 ~@2"`), which is how "ring two beats then release for two" is
// spelled.
function renderSteps(steps, path) {
  const folded = [];
  for (const step of steps) {
    if (step.kind === 'hold') {
      const prev = folded[folded.length - 1];
      if (!prev) {
        throw new Error(
          `the hold token '_' extends the step before it, but there is no step before it${path ? ` (inside ${path})` : ''} - start the sequence with a note or a rest.`,
        );
      }
      // A hold's own `@n`/`!n` say how MANY steps' worth of ring it adds.
      prev.extra += step.weight * step.replicate;
      continue;
    }
    folded.push({ step, extra: 0 });
  }

  const out = [];
  for (const { step, extra } of folded) {
    const body = renderStepBody(step);
    if (extra === 0) {
      const w = step.weight === 1 ? '' : `@${formatNumber(step.weight)}`;
      const r = step.replicate === 1 ? '' : `!${step.replicate}`;
      out.push(`${body}${step.suffix}${w}${r}`);
      continue;
    }
    // Only the LAST of a replicated step is extended by the hold, so the
    // replication is written out as separate steps first. `!n` is exactly
    // n copies of the same step, so this is a spelling change, not a
    // semantic one.
    for (let i = 0; i < step.replicate - 1; i += 1) {
      const w = step.weight === 1 ? '' : `@${formatNumber(step.weight)}`;
      out.push(`${body}${step.suffix}${w}`);
    }
    out.push(`${body}${step.suffix}@${formatNumber(step.weight + extra)}`);
  }
  return out.join(' ');
}

function renderStepBody(step) {
  switch (step.kind) {
    case 'group':
      return `[${step.seqs.map((s) => renderSteps(s, 'a [...] group')).join(', ')}]`;
    case 'poly':
      return `{${step.seqs.map((s) => renderSteps(s, 'a {...} group')).join(', ')}}`;
    case 'alt':
      return `<${step.entries.map((e) => renderSteps([e], 'a <...> alternation')).join(' ')}>`;
    case 'choice':
      return step.entries.map((e) => renderSteps([e], 'a choice')).join('|');
    default:
      return step.text;
  }
}

// formatNumber - shortest exact decimal spelling, so a folded weight reads
// as `@4` rather than `@4.000000000000001` when it came from summing
// quarter-beat holds.
function formatNumber(n) {
  const rounded = Math.round(n * 1e6) / 1e6;
  return String(rounded);
}

export function hasHoldToken(source) {
  // Cheap pre-check so a string with no hold at all is returned byte-for-
  // byte unchanged by expandHolds (every pre-3.0.0 song takes this path).
  return source.includes(HOLD_TOKEN);
}

// expandHolds(source) -> the same pattern with every `_` folded into the
// weight of the step it extends. Returns `source` itself, untouched, when
// there is no hold token in it.
export function expandHolds(source, beatUnit = DEFAULT_BEAT_UNIT, refs = null) {
  if (!hasHoldToken(source)) return source;
  const seqs = scanMini(source, beatUnit, 1, refs);
  return seqs.map((s) => renderSteps(s, null)).join(', ');
}

// expandEach(source, denominator, beatUnit) -> the same pattern with every
// step's default weight set to beatUnit/denominator instead of 1 (ADR 0003
// D3's `each 1/D`), rendered out explicitly (`@n`) the same way
// expandHolds() renders a folded hold - a step that already carries its own
// `@n`/`/N` is untouched (`each` composes: it is the default, an explicit
// duration overrides it for that token). Also folds any hold tokens present,
// in the same pass, since renderSteps() always does.
export function expandEach(source, denominator, beatUnit = DEFAULT_BEAT_UNIT, refs = null) {
  const seqs = scanMini(source, beatUnit, beatUnit / denominator, refs);
  return seqs.map((s) => renderSteps(s, null)).join(', ');
}

// findSpliceRepeatHazards (ADR 0001 decision 7 / known defect 2) is
// deleted as of ADR 0003 D5: `$name` is a structural reference now, not
// text substitution, so the hazard it warned about (`$bassA!3` repeating
// only the last SPLICED token, because `!` landed on already-substituted
// text) cannot occur - `!` on a $ref applies to the WHOLE resolved
// pattern (see miniNotation.js's parseRef()/parseStep()).
