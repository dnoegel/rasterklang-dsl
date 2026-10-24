// Semantic analysis for DSL v2/3: everything that is decidable from the AST
// alone, before a single Hap or register write exists.
//
// This layer (since 3.0.0) exists because ADR 0001 added rules that are
// neither grammatical (the parser can't see them) nor
// engine-level (the pattern engine shouldn't have to care): a pattern's
// declared length has to match its contents, a song has to have exactly one
// clock statement, an arrangement of different-length voices has a period
// worth telling the composer about. Keeping them here rather than sprinkled
// through interpret.js means they are also usable WITHOUT the pattern
// engine - `checkProgram()` needs nothing but this repo, which is what lets
// `scripts/dsl-lint.mjs` (and any editor integration) run the migration
// checker over a file without compiling it.
//
// Nothing in this file throws for a diagnostic; it returns them. interpret()
// is what decides that an error-severity diagnostic ends compilation, and
// that is deliberately its decision rather than this module's, so a linter
// can report ALL of a file's problems instead of only the first.
import {
  contentBeats,
  declaredLengthCandidates,
  beatsEqual,
  expandHolds,
  expandEach,
  checkBarLines,
  proportionalFitInfo,
  smallestDurationBeats,
} from './miniLength.js';

// The 11-bit filter cutoff register's maximum. A larger value has always
// been clipped to this silently; since 3.0.0 it also says so (ADR 0001
// decision 9 / known defect 4).
export const MAX_CUTOFF = 2047;

// Time signatures a real piece of music plausibly uses. Anything outside
// this set is legal (5/8 and 7/8 exist) but is also exactly what `beats`
// looked like when it was being abused as a melody-stretch lever before
// declared pattern lengths existed - the migration hint in ADR 0001
// decision 10.
const COMMON_BEATS = [1, 2, 3, 4, 6, 8, 9, 12, 16];

// An arrangement whose voices only realign after this many bars is far more
// likely to be a typo'd declared length than an intentional phase piece, so
// the period report gains a "did you mean this?" note past it. Not an
// error: intentional long-period polymetre is legitimate music (ADR 0001's
// D5).
const IMPLAUSIBLE_PERIOD_BARS = 64;

const DEFAULT_BAR_BEATS = 4;
const DEFAULT_BEAT_UNIT = 4;

function diag(severity, line, message) {
  return { severity, line, message };
}

// === Rational helpers ========================================================

// Declared lengths may be decimals (`pattern lead 7.5 = ...`, ADR 0001
// decision 5), so the arrangement period is an LCM over RATIONALS, not
// integers. Scaling by a fixed 1/1000-beat denominator is exact for every
// length the DSL can spell (source text is decimal) and avoids float GCD.
const LENGTH_SCALE = 1000;

function gcdInt(a, b) {
  let x = Math.abs(Math.round(a));
  let y = Math.abs(Math.round(b));
  while (y) {
    const t = x % y;
    x = y;
    y = t;
  }
  return x;
}

function lcmBeats(a, b) {
  const ai = Math.round(a * LENGTH_SCALE);
  const bi = Math.round(b * LENGTH_SCALE);
  if (ai === 0 || bi === 0) return 0;
  return (Math.abs(ai * bi) / gcdInt(ai, bi)) / LENGTH_SCALE;
}

// === The pass ================================================================

// analyzeProgram(ast) -> { diagnostics, barBeats, clock, patterns, voices,
// periodBeats }.
//
// `patterns` is a Map name -> { declaredLength, source, contentBeats }, the
// same symbol table interpret.js needs; `voices` is one entry per voice
// statement with the metric length that voice actually plays at.
export function analyzeProgram(ast) {
  const diagnostics = [];
  const patterns = new Map();
  const clocks = [];
  const voices = [];
  let barBeats = DEFAULT_BAR_BEATS;
  let beatUnit = DEFAULT_BEAT_UNIT;
  let barBeatsLine = null;
  // ADR 0003 D17: the smallest explicit duration used ANYWHERE in the song
  // (a named pattern or a voice's own inline notes), in beats - null if the
  // song uses no explicit durations at all. interpret.js uses this against
  // the song's actual frame grid to warn about subdivision jitter; computed
  // here (not there) because this is the layer that already scans every
  // pattern/voice's content for weights.
  let songSmallestDurationBeats = null;
  const reportDuration = (beats) => {
    if (beats == null) return;
    songSmallestDurationBeats = songSmallestDurationBeats == null ? beats : Math.min(songSmallestDurationBeats, beats);
  };

  // Pre-scanned so a `pattern ... each N/D = ...` or an arrangement
  // declared BEFORE the song's tempo statement (unusual, but not
  // grammatically forbidden) still resolves against the song's real
  // meter rather than the default - the main loop below only updates
  // barBeats/beatUnit as it reaches the TempoStmt, which is too late for
  // an earlier PatternDecl's own resolution.
  const earlyMeter = (() => {
    for (const s of ast.statements) {
      if (s.type !== 'TempoStmt') continue;
      if (s.meter != null) return { barBeats: s.meter.numerator, beatUnit: s.meter.denominator };
      if (s.beats != null) return { barBeats: s.beats, beatUnit: DEFAULT_BEAT_UNIT };
      return { barBeats: DEFAULT_BAR_BEATS, beatUnit: DEFAULT_BEAT_UNIT };
    }
    return { barBeats: DEFAULT_BAR_BEATS, beatUnit: DEFAULT_BEAT_UNIT };
  })();

  for (const stmt of ast.statements) {
    switch (stmt.type) {
      case 'TempoStmt':
        clocks.push(stmt);
        if (stmt.meter != null) {
          barBeats = stmt.meter.numerator;
          beatUnit = stmt.meter.denominator;
          barBeatsLine = stmt.line;
        } else if (stmt.beats != null) {
          // ADR 0003 D2: `beats N` stays legal as an alias for `meter N/4`,
          // documented as deprecated (docs/dsl-grammar-spec.md) - it never
          // said what a beat unit was, `meter` gives it the denominator it
          // always implied. No new diagnostic per usage: the existing
          // unusual-value migration hint (checkBeatsHeuristic, ADR 0001
          // decision 10) already flags the case worth flagging.
          barBeats = stmt.beats;
          beatUnit = DEFAULT_BEAT_UNIT;
          barBeatsLine = stmt.line;
        }
        break;
      case 'PatternDecl':
        if (stmt.elements != null) checkArrangementPatternDecl(stmt, patterns, diagnostics, earlyMeter.barBeats);
        else checkPatternDecl(stmt, patterns, diagnostics, earlyMeter.beatUnit, earlyMeter.barBeats, reportDuration);
        break;
      case 'VoiceStmt':
        voices.push(collectVoice(stmt, patterns, diagnostics, earlyMeter.barBeats, earlyMeter.beatUnit, reportDuration));
        break;
      case 'FilterBlock':
        checkFilterBlock(stmt, diagnostics);
        break;
      default:
        break;
    }
  }

  checkClock(clocks, diagnostics);
  checkBeatsHeuristic(barBeats, barBeatsLine, diagnostics);

  // Resolve every voice's length now that `beats` is known: a voice with no
  // declared-length pattern is proportional and so fills exactly one bar,
  // which is the pre-3.0.0 reading, restated in the new vocabulary (ADR
  // 0001's D3).
  for (const v of voices) {
    if (v.metricBeats == null) v.effectiveBeats = barBeats;
    else v.effectiveBeats = v.metricBeats;
  }

  const periodBeats = voices.reduce((acc, v) => (acc === null ? v.effectiveBeats : lcmBeats(acc, v.effectiveBeats)), null);
  if (periodBeats !== null && voices.length > 0) {
    diagnostics.push(diag('info', clocks.length > 0 ? clocks[0].line : 1, arrangementReport(voices, periodBeats, barBeats)));
    if (periodBeats / barBeats > IMPLAUSIBLE_PERIOD_BARS) {
      diagnostics.push(
        diag(
          'warning',
          clocks.length > 0 ? clocks[0].line : 1,
          `this arrangement's voices only realign after ${plural(periodBeats / barBeats, 'bar')}, which is usually a mistyped declared length rather than intended polymetre - check the pattern lengths above.`,
        ),
      );
    }
  }

  return { diagnostics, barBeats, beatUnit, clock: clocks[0] ?? null, patterns, voices, periodBeats, smallestDurationBeats: songSmallestDurationBeats };
}

// checkProgram(ast) -> diagnostics only, the entry point for a lint pass
// (scripts/dsl-lint.mjs) that wants to report everything without
// compiling.
export function checkProgram(ast) {
  return analyzeProgram(ast).diagnostics;
}

function arrangementReport(voices, periodBeats, barBeats) {
  const parts = voices.map((v) => `voice ${v.index} ${formatBeats(v.effectiveBeats)} beats${v.metricBeats == null ? ' (proportional, one bar)' : ''}`);
  return `Arrangement: ${parts.join(', ')}. Full period: ${formatBeats(periodBeats)} beats (${plural(periodBeats / barBeats, 'bar')}).`;
}

function plural(n, word) {
  return `${formatBeats(n)} ${word}${beatsEqual(n, 1) ? '' : 's'}`;
}

function formatBeats(n) {
  const rounded = Math.round(n * 1e6) / 1e6;
  return String(rounded);
}

// --- clock (ADR 0001's D4 / decision 11) ------------------------------------

function checkClock(clocks, diagnostics) {
  if (clocks.length === 0) {
    diagnostics.push(
      diag('error', 1, 'a song needs exactly one clock statement, and this one has none - add either `tempo 120 beats 4` (BPM) or `speed 6 beats 4` (frames per beat).'),
    );
    return;
  }
  if (clocks.length === 1) return;
  const units = new Set(clocks.map((c) => c.unit));
  const lines = clocks.map((c) => c.line).join(', ');
  if (units.size > 1) {
    diagnostics.push(
      diag(
        'error',
        clocks[1].line,
        `\`tempo\` and \`speed\` are two spellings of the same clock and are mutually exclusive, but this song has both (lines ${lines}) - keep the one you mean and delete the other.`,
      ),
    );
    return;
  }
  const word = clocks[0].unit === 'framesPerBeat' ? 'speed' : 'tempo';
  diagnostics.push(diag('error', clocks[1].line, `a song needs exactly one clock statement, but this one has ${clocks.length} \`${word}\` statements (lines ${lines}).`));
}

// --- beats heuristic (ADR 0001 decision 10) ---------------------------------

function checkBeatsHeuristic(barBeats, line, diagnostics) {
  if (line === null || COMMON_BEATS.includes(barBeats)) return;
  diagnostics.push(
    diag(
      'warning',
      line,
      `beats ${barBeats} is unusual for a time signature; if you meant to stretch a melody, give the pattern a declared length instead (\`pattern name ${barBeats} = ...\`).`,
    ),
  );
}

// --- patterns (ADR 0001's D1) ------------------------------------------------

// patternBeatsOf(name, patterns, barBeats) -> a named pattern's own beats,
// or null if undeclared (a forward/unknown reference - reported by whatever
// throws when it actually tries to resolve one, e.g. miniLength.js's
// scanner or interpret.js's lookupPattern; this function only looks up
// already-declared entries, never guesses). An arrangement-form pattern's
// beats were already computed at ITS OWN declaration (see
// checkArrangementPatternDecl); a plain (string) pattern's beats are its
// declared length, or one bar if proportional (D4's "an element with no
// metric length occupies one bar" - the same rule for a $ref used inside a
// mini-notation string as for an arrangement element, since both are the
// SAME resolveRefLength()-shaped question, "how many beats is $name").
export function patternBeatsOf(name, patterns, barBeats) {
  const p = patterns.get(name);
  if (!p) return null;
  if (p.elements != null) return p.arrangementBeats;
  return p.declaredLength != null ? p.declaredLength : barBeats;
}

function checkPatternDecl(stmt, patterns, diagnostics, beatUnit = DEFAULT_BEAT_UNIT, barBeats = DEFAULT_BAR_BEATS, reportDuration = () => {}) {
  const refs = (name) => patternBeatsOf(name, patterns, barBeats);
  // ADR 0003 D17: scanned on the RAW source (each's own implied 1/D counts
  // as a duration too, via smallestDurationBeats' defaultWeight parameter -
  // scanning the ALREADY each-expanded text would find only @n weights,
  // losing that).
  try {
    reportDuration(smallestDurationBeats(stmt.source, beatUnit, stmt.each != null ? beatUnit / stmt.each.denominator : 1, refs));
  } catch {
    // Reported again (and more specifically) by the try block below.
  }
  let candidates = null;
  // ADR 0003 D5: $name is a structural reference now, resolved by
  // miniLength.js's scanner itself (via `refs`) rather than spliced into
  // the text beforehand - stmt.source keeps its literal `$name` tokens
  // all the way through.
  let expandedSource = stmt.source;
  try {
    // ADR 0003 D3: `each N/D` sets every unweighted token's default
    // weight to beatUnit/D instead of the usual 1, folding holds in the
    // same pass; otherwise holds alone are folded so `"c4 _ _ _"`
    // measures as the 4 beats it sounds for, not as one note plus three
    // mystery tokens. Either way, a $ref token's own weight (from `refs`)
    // is untouched - `each`'s default only fills in for a token with no
    // explicit weight of its own, and a reference's weight is never that.
    expandedSource = stmt.each != null
      ? expandEach(stmt.source, stmt.each.denominator, beatUnit, refs)
      : expandHolds(stmt.source, beatUnit, refs);
    candidates = declaredLengthCandidates(expandedSource, beatUnit, refs);
  } catch (e) {
    diagnostics.push(diag('error', stmt.line, `pattern "${stmt.name}": ${e.message}`));
  }

  // ADR 0003 D6: a standalone `|` asserts a bar boundary; check it against
  // the same content the length check just used (barBeats already known).
  try {
    checkBarLines(expandedSource, barBeats, beatUnit, refs);
  } catch (e) {
    diagnostics.push(diag('error', stmt.line, `pattern "${stmt.name}": ${e.message}`));
  }

  // ADR 0003 D8: a pattern with no /N, no each and no declared number is
  // auto-fit proportionally to one bar, and says so instead of reaching the
  // author only as sound.
  if (stmt.length == null && stmt.each == null) {
    const info = proportionalFitInfo(expandedSource, beatUnit, refs);
    if (info) {
      diagnostics.push(
        diag(
          'info',
          stmt.line,
          `pattern "${stmt.name}" is proportional: ${info.tokenCount} token${info.tokenCount === 1 ? '' : 's'} over 1 bar = 1/${info.tokenCount} note each. ` +
            `Write durations (c4/${info.tokenCount}) or \`each 1/${info.tokenCount}\` to fix the grid.`,
        ),
      );
    }
  }

  // ADR 0003 D3: `each N/D` alone (no explicit declared number) still makes
  // the pattern metric - "length is derived" - so a voice referencing it
  // bare (`$hats`, no brackets) gets the same declared-length stretching a
  // written number would trigger, using the summed duration as that number.
  const derivedLength = stmt.length ?? (stmt.each != null && candidates ? candidates[0].beats : null);

  const entry = {
    name: stmt.name,
    line: stmt.line,
    source: stmt.source,
    declaredLength: derivedLength,
    contentBeats: candidates ? candidates[0].beats : null,
  };
  patterns.set(stmt.name, entry);

  if (stmt.length == null || candidates === null) return;
  for (const c of candidates) {
    if (beatsEqual(c.beats, stmt.length)) continue;
    diagnostics.push(
      diag(
        'error',
        stmt.line,
        `pattern "${stmt.name}" declares ${formatBeats(stmt.length)} beats but its ${c.what} add up to ${formatBeats(c.beats)} beats ` +
          `(in a pattern with a declared length, an unweighted token is one beat and \`@n\` is n beats) - ` +
          `either fix the contents or declare \`pattern ${stmt.name} ${formatBeats(c.beats)} = ...\`.`,
      ),
    );
  }
}

// --- voices ------------------------------------------------------------------

// notesSourceOf(vstmt) -> the raw (pre-splice) notes text of a voice
// statement, in either grammar mode, plus whether it was a bare `$name`.
function notesSourceOf(vstmt) {
  if (vstmt.mode === 'short') {
    if (vstmt.notes == null) return null; // arrangement mode (ADR 0003 D4) - no old-shape source text
    return vstmt.notes.isSigil ? `$${vstmt.notes.sigilName}` : vstmt.notes.source;
  }
  const f = vstmt.fields.find((x) => x.type === 'notes');
  if (!f) return null;
  return f.isSigil ? `$${f.sigilName}` : f.source;
}

// bareSigilNameOf(vstmt) -> the referenced pattern's name if this voice's
// notes are a single unquoted `$name` (either grammar mode), else null.
// A QUOTED string containing only "$name"/"[$name]" deliberately does not
// count: it is an anonymous inline pattern and stays proportional.
function bareSigilNameOf(vstmt) {
  if (vstmt.mode === 'short') {
    return vstmt.notes && vstmt.notes.isSigil ? vstmt.notes.sigilName : null;
  }
  const f = vstmt.fields.find((x) => x.type === 'notes');
  return f && f.isSigil ? f.sigilName : null;
}

function collectVoice(vstmt, patterns, diagnostics, barBeats, beatUnit, reportDuration = () => {}) {
  const entry = { index: vstmt.index, line: vstmt.line, metricBeats: null, effectiveBeats: null, elements: null };
  const source = notesSourceOf(vstmt);

  // ADR 0003 D4/D5: an arrangement (`voice 0 $a $b`) or a bare single
  // sigil (`voice 0 $lead`, read as a degenerate one-element arrangement)
  // both resolve bottom-up over the reference tree; sequencing different
  // lengths is an arrangement's whole point, so it is not an error. A
  // quoted string (including one that happens to read "$lead" or
  // "<$a $b>") stays proportional at the voice level.
  const bareSigil = bareSigilNameOf(vstmt);
  const elements = vstmt.arrangement ?? (bareSigil != null ? [{ line: vstmt.line, ref: bareSigil, repeat: 1 }] : null);

  if (source != null && elements == null) {
    const refs = (name) => patternBeatsOf(name, patterns, barBeats);
    try {
      reportDuration(smallestDurationBeats(source, beatUnit, 1, refs));
    } catch {
      // Reported again (and more specifically) below.
    }
    // A voice's own inline notes text is an anonymous pattern (ADR 0001's
    // D1: it stays proportional), but its hold tokens (and, since ADR 0003
    // D5, any $name references in it) still have to make sense, and a hold
    // with nothing before it - or a reference to an undeclared/forward
    // pattern - is worth catching here rather than as a mid-compile throw.
    // ADR 0001 decision 7's `$name!n` warning is not needed: `$name!n`
    // replicates the WHOLE referenced pattern (miniNotation.js's
    // parseRef()/parseStep()).
    let expanded = source;
    try {
      expanded = expandHolds(source, beatUnit, refs);
    } catch (e) {
      diagnostics.push(diag('error', vstmt.line, `voice ${vstmt.index}: ${e.message}`));
    }
    // ADR 0003 D6: a standalone `|` asserts a bar boundary in this voice's
    // own inline pattern too, not just a named declaration.
    try {
      checkBarLines(expanded, barBeats, beatUnit, refs);
    } catch (e) {
      diagnostics.push(diag('error', vstmt.line, `voice ${vstmt.index}: ${e.message}`));
    }
    // ADR 0003 D8: same proportional-auto-fit report a named pattern gets.
    const info = proportionalFitInfo(expanded, beatUnit, refs);
    if (info) {
      diagnostics.push(
        diag(
          'info',
          vstmt.line,
          `voice ${vstmt.index} is proportional: ${info.tokenCount} token${info.tokenCount === 1 ? '' : 's'} over 1 bar = 1/${info.tokenCount} note each. ` +
            `Write durations (c4/${info.tokenCount}) or \`each 1/${info.tokenCount}\` on a named pattern to fix the grid.`,
        ),
      );
    }
  }

  if (elements == null) return entry;

  const result = resolveArrangementLength(elements, patterns, new Set(), diagnostics, barBeats, vstmt.line);
  if (result) {
    entry.metricBeats = result.beats;
    entry.elements = result.flat;
  }
  return entry;
}

// === Arrangements (ADR 0003 D4) =============================================

// resolveArrangementLength(elements, patterns, visiting, diagnostics,
// barBeats, contextLine) -> { beats, flat } | null (null when a leaf could
// not be resolved at all - either a forward/unknown reference, reported
// elsewhere, or a recursive one, reported here).
//
// `elements` is the flat AST shape astVisitor.js's buildArrangement()
// produces: [{ line, ref? , group?, repeat }]. `*n`/`[...]` grouping is
// flattened into the entry list HERE, at analysis time, matching 5.2's
// "D4's *n and [...] grouping are flattened into the entry list before
// the timeline is built. They are not runtime constructs" - interpret.js
// gets the already-flat list and never re-derives it.
function resolveArrangementLength(elements, patterns, visiting, diagnostics, barBeats, contextLine) {
  const flat = [];
  for (const el of elements) {
    const line = el.line ?? contextLine;
    let subFlat;
    if (el.ref != null) {
      const r = resolveRefLength(el.ref, patterns, visiting, diagnostics, barBeats, line);
      if (r == null) return null;
      subFlat = r.flat;
    } else {
      const r = resolveArrangementLength(el.group, patterns, visiting, diagnostics, barBeats, line);
      if (r == null) return null;
      subFlat = r.flat;
    }
    for (let i = 0; i < el.repeat; i += 1) flat.push(...subFlat);
  }
  const beats = flat.reduce((sum, e) => sum + e.beats, 0);
  return { beats, flat };
}

function resolveRefLength(name, patterns, visiting, diagnostics, barBeats, line) {
  if (visiting.has(name)) {
    diagnostics.push(diag('error', line, `pattern "${name}" is a recursive reference (it references itself, directly or through another pattern) - an arrangement cannot have a cycle.`));
    return null;
  }
  // patternBeatsOf returns null for both "not declared" (an unresolved
  // forward reference - interpret.js reports this) and "declared, but its
  // OWN resolution already failed" (already reported at that declaration).
  const beats = patternBeatsOf(name, patterns, barBeats);
  if (beats == null) return null;
  return { beats, flat: [{ name, beats }] };
}

// checkArrangementPatternDecl mirrors checkPatternDecl but for the
// arrangement form of patternDecl (`pattern verse = $riffA $riffA $riffB`,
// ADR 0003 D4) - no declared length, no mini-notation content, just a
// reference tree resolved bottom-up. `visiting` starts as {stmt.name} so a
// pattern referencing itself (directly - true multi-hop cycles are
// impossible under this language's forward-only declaration order) is
// caught as recursion rather than silently read as an unresolved forward
// reference.
function checkArrangementPatternDecl(stmt, patterns, diagnostics, barBeats) {
  const visiting = new Set([stmt.name]);
  const result = resolveArrangementLength(stmt.elements, patterns, visiting, diagnostics, barBeats, stmt.line);
  patterns.set(stmt.name, {
    name: stmt.name,
    line: stmt.line,
    elements: stmt.elements,
    arrangementBeats: result ? result.beats : null,
    flat: result ? result.flat : null,
  });
}

// --- filter (ADR 0001 decision 9 / known defect 4) ---------------------------

function checkFilterBlock(stmt, diagnostics) {
  for (const f of stmt.fields) {
    if (f.type !== 'cutoff') continue;
    for (const value of cutoffValuesOf(f.expr)) {
      if (value > MAX_CUTOFF) {
        diagnostics.push(diag('warning', f.line, `cutoff ${formatBeats(value)} clipped to ${MAX_CUTOFF} (11-bit filter cutoff).`));
      }
    }
  }
}

function cutoffValuesOf(expr) {
  if (!expr) return [];
  if (expr.type === 'SwitchBlock') return expr.cases.flatMap((c) => cutoffValuesOf(c.value));
  if (expr.type === 'NumberListExpr') return expr.values;
  if (expr.type === 'SignalExpr') return expr.range ? [expr.range.lo, expr.range.hi] : [];
  return [];
}

export { contentBeats };
