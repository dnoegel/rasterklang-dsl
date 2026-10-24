// AST-transform visitor: walks the ANTLR parse tree (grammar/
// RasterklangDsl.g4, generated into ./generated/) and builds the
// plain-object AST node shapes (see docs/dsl-grammar-spec.md's "Grammar
// (EBNF-ish)" section for the production names these mirror 1:1).
//
// Two design choices, both explained at more length in
// grammar/RasterklangDsl.g4's header comment:
// 1. Keywords (field/transform/statement names) are NOT baked into the
//    grammar as separate token types (except the statement keywords and
//    `end`) -- most of them are IDENT-shaped text the parser accepts
//    generically, and THIS file dispatches on the actual text and
//    validates it's a known name.
// 2. buildProgram(tree, tokens) takes the ORIGINAL (pre-replay) Token[]
//    array parser.js's parseProgram() was called with, and reads
//    value/docStart/docEnd for STRING/NUMBER/HEX/SIGIL terminals from
//    THAT array (via each ANTLR terminal's tokenIndex) rather than from
//    the replayed ANTLR token's own .text/.start/.stop -- the replay
//    step (antlrSupport.js's ReplayTokenSource) doesn't reconstruct raw
//    source positions or re-escape string content, so this is the one
//    place that information can still come from.

import { LOOP_TABLE_NAMES } from './instrument.js';

// Exported so parser.js/keywords.js can re-export these rather than keep a
// second, drift-prone copy of the same list.
export const STATEMENT_KEYWORDS = ['tempo', 'speed', 'target', 'chip', 'pattern', 'table', 'instrument', 'voice', 'filter'];
export const INSTRUMENT_FIELDS = ['wave', 'adsr', 'pw', 'arp', 'filterTable', 'loop', 'vib', 'vibTri', 'slide', 'test', 'gate', 'hardRestart'];
export const GATE_MODES = ['legato', 'retrigger'];
export const VOICE_FIELDS = ['notes', 'insts', 'fx', 'reg', 'vib', 'slide', 'gate', 'sync', 'ring', 'wave'];
export const TRANSFORMS = [
  'rev', 'fast', 'slow', 'every', 'iter', 'ply', 'off', 'jux', 'degradeBy', 'sometimesBy', 'add',
  'struct', 'mask', 'chunk', 'palindrome', 'linger', 'rot', 'someBarsBy', 'someCyclesBy', 'superimpose', 'scale',
];
export const FILTER_FIELDS = ['cutoff', 'resonance', 'mode', 'route'];
// `switch bar` is 3.0.0's spelling of the older `switch cycle` (ADR
// 0001's D3). Both are accepted: the rename is a vocabulary change, not a
// behaviour change, and silently breaking every existing filter block over
// a word would be a much worse trade than carrying one alias.
const SWITCH_UNIT_WORDS = ['bar', 'cycle'];
// ADR 0002's `target { video PAL|NTSC ... }`.
export const VIDEO_STANDARDS = ['PAL', 'NTSC'];
export const CHIP_MODELS = [6581, 8580];

function fail(line, msg) {
  throw new Error(`line ${line}: ${msg}`);
}

function ident(ctx) {
  return ctx.getText();
}

// tokens/orig -- see this file's header comment. Set once per buildProgram
// call; safe because parsing is synchronous and non-reentrant.
let TOKENS = null;
function orig(terminalNode) {
  return TOKENS[terminalNode.symbol.tokenIndex];
}

// strInfo(ctx: StrContext) -> the original { value, docStart, docEnd }
function strInfo(ctx) {
  const term = ctx.DQSTRING() || ctx.BTSTRING();
  return orig(term);
}

function numOf(term) {
  return orig(term).value;
}

function buildRamp(ctx) {
  const line = ctx.start.line;
  const idents = ctx.anyIdent();
  const nums = ctx.NUMBER();
  const from = numOf(nums[0]);
  const to = numOf(nums[1]);
  let step = null;
  let over = null;
  // ADR 0003 D9: a trailing bare anyIdent (no NUMBER of its own to pair
  // with) is the shape terminator, pingpong/loop - structurally
  // distinguished from step/over here by whether a NUMBER is still
  // available to consume, not by position, since either optional pair may
  // be absent. idents[0] is the leading 'ramp' keyword itself.
  let shape = null;
  let numIdx = 2;
  for (let i = 1; i < idents.length; i += 1) {
    const word = ident(idents[i]);
    if (numIdx < nums.length) {
      const val = numOf(nums[numIdx]);
      numIdx += 1;
      if (word === 'step') step = val;
      else if (word === 'over') over = Math.trunc(val);
      else fail(line, `expected 'step' or 'over', found '${word}'`);
    } else {
      if (word !== 'pingpong' && word !== 'loop') fail(line, `expected 'pingpong' or 'loop', found '${word}'`);
      shape = word;
    }
  }
  if (step === null && over === null) {
    fail(line, "a 'ramp' needs at least one of 'step N' (per-frame delta) or 'over N' (length in frames).");
  }
  if (step === 0) fail(line, "a 'ramp' step of 0 never reaches its target.");
  if (over !== null && over < 1) fail(line, `a 'ramp' must run over at least 1 frame, got ${over}.`);
  return { gen: 'ramp', line, from, to, step, over, shape };
}

// buildAbsNote(ctx: AbsNoteContext) -> { absNote } | { repeat: { absNote }, count }
//
// `=60` - an ABSOLUTE MIDI note for this table frame, the fixed-pitch
// counterpart to a relative `arp` offset (see the grammar's own comment on
// the absNote rule). `=60!4` repeats it, same `!N` shape every other table
// cell already has.
function buildAbsNote(ctx) {
  const nums = ctx.NUMBER();
  const value = { absNote: numOf(nums[0]) };
  if (ctx.BANG()) {
    const count = numOf(nums[1]);
    if (count < 1) fail(ctx.start.line, `a table repeat needs a count of at least 1, got ${count}.`);
    return { repeat: value, count };
  }
  return value;
}

// buildTableToken(ctx: TableTokenContext) -> plain string | number | {ref} | {absNote} | {repeat,count} | ramp object
function buildTableToken(ctx) {
  if (ctx.rampGenerator()) return buildRamp(ctx.rampGenerator());
  if (ctx.SIGIL()) return { ref: orig(ctx.SIGIL()).value };
  if (ctx.absNote()) return buildAbsNote(ctx.absNote());
  const idents = ctx.anyIdent();
  const nums = ctx.NUMBER();
  const bang = ctx.BANG();
  if (idents.length > 0) {
    const word = idents.map(ident).join('+');
    if (bang) {
      const count = numOf(nums[0]);
      if (count < 1) fail(ctx.start.line, `a table repeat needs a count of at least 1, got ${count}.`);
      return { repeat: word, count };
    }
    return word;
  }
  const base = numOf(nums[0]);
  if (bang) {
    const count = numOf(nums[1]);
    if (count < 1) fail(ctx.start.line, `a table repeat needs a count of at least 1, got ${count}.`);
    return { repeat: base, count };
  }
  return base;
}

function buildTableRow(ctx) {
  return ctx.tableToken().map(buildTableToken);
}

// === Program ===

export function buildProgram(ctx, tokens) {
  TOKENS = tokens;
  try {
    const statements = ctx.statement().map(buildStatement);
    return { type: 'Program', statements };
  } finally {
    TOKENS = null;
  }
}

function buildStatement(ctx) {
  if (ctx.tempoStmt()) return buildTempoStmt(ctx.tempoStmt());
  if (ctx.targetStmt()) return buildTargetStmt(ctx.targetStmt());
  if (ctx.chipStmt()) return buildChipStmt(ctx.chipStmt());
  if (ctx.voiceStmt()) return buildVoiceStmt(ctx.voiceStmt());
  if (ctx.patternDecl()) return buildPatternDecl(ctx.patternDecl());
  if (ctx.tableDecl()) return buildTableDecl(ctx.tableDecl());
  if (ctx.instrumentBlock()) return buildInstrumentBlock(ctx.instrumentBlock());
  if (ctx.filterBlock()) return buildFilterBlock(ctx.filterBlock());
  return buildUnknownStmt(ctx.unknownStmt());
}

function buildUnknownStmt(ctx) {
  const name = ctx.IDENT().getText();
  return fail(ctx.start.line, `unknown statement keyword '${name}' (expected ${STATEMENT_KEYWORDS.join('/')})`);
}

function buildTempoStmt(ctx) {
  const line = ctx.start.line;
  const tv = ctx.tempoValue();
  let value;
  if (tv.NUMBER()) value = { kind: 'number', value: numOf(tv.NUMBER()) };
  else value = { kind: 'pattern', source: strInfo(tv.str()).value };
  // Which of the two clock spellings was written (ADR 0001's D4): 'bpm'
  // for `tempo N` (musical), 'framesPerBeat' for `speed N` (exact, one
  // frame = one PAL raster frame, what GoatTracker natively counts). Both
  // produce the same node type on purpose - everything downstream of the
  // unit conversion in interpret.js treats them identically, and the
  // "exactly one clock statement" rule is a semantic check, not a
  // grammar-level XOR (decision 11).
  const unit = ctx.SPEED() ? 'framesPerBeat' : 'bpm';
  // The trailing settings are name/number(/number) triples (`beats N`,
  // `offset N`, `meter N/D` - ADR 0003 D2) - see the grammar's own
  // comment. Validated here rather than in the grammar because they're
  // soft keywords like every other field name.
  let beats = null;
  let offset = null;
  let meter = null;
  for (const sc of ctx.tempoSetting()) {
    const name = ident(sc.anyIdent());
    const numbers = sc.NUMBER();
    const value_ = numOf(numbers[0]);
    if (name === 'beats') {
      if (numbers.length > 1) fail(line, `'beats' takes a single number, not a fraction - did you mean 'meter ${value_}/${numOf(numbers[1])}'?`);
      if (!Number.isInteger(value_)) fail(line, `expected an integer, got ${value_}`);
      beats = value_;
    } else if (name === 'offset') {
      if (numbers.length > 1) fail(line, `'offset' takes a single number, not a fraction`);
      offset = value_;
    } else if (name === 'meter') {
      if (numbers.length < 2) fail(line, `'meter' needs a fraction, e.g. 'meter 4/4'`);
      const denominator = numOf(numbers[1]);
      if (!Number.isInteger(value_) || !Number.isInteger(denominator)) {
        fail(line, `meter's numerator and denominator must be integers, got ${value_}/${denominator}`);
      }
      meter = { numerator: value_, denominator };
    } else {
      fail(line, `expected 'meter', 'beats' or 'offset', found '${name}'`);
    }
  }
  if (beats != null && meter != null) {
    fail(line, `a tempo statement cannot combine 'meter' and 'beats' - 'beats' is deprecated in favor of 'meter N/4', use one or the other`);
  }
  return { type: 'TempoStmt', line, unit, value, beats, offset, meter };
}

// ADR 0002's `target { ... }` block: the song's DRIVE rate, i.e. how many
// CPU cycles pass between two play calls. Every setting is optional and the
// whole block is optional; absent means driveRate.js's DEFAULT_DRIVE_CYCLES
// (PAL, exactly 50.00 Hz), which is DSL v2's only pre-0002 behaviour.
//
// Validated here rather than in the grammar for the same reason `beats`/
// `offset` are: they are soft keywords, so the grammar only knows the shape
// (name value)*, and this is the layer that knows the names.
function buildTargetStmt(ctx) {
  const line = ctx.start.line;
  let video = null;
  let driveCycles = null;
  let multispeed = null;
  for (const setting of ctx.targetSetting()) {
    const idents = setting.anyIdent();
    const name = ident(idents[0]);
    const numTerm = setting.NUMBER();
    const wordCtx = idents.length > 1 ? idents[1] : null;
    switch (name) {
      case 'video': {
        if (!wordCtx) fail(line, `'video' takes a video standard (${VIDEO_STANDARDS.join(' or ')}), not a number`);
        const word = ident(wordCtx);
        if (!VIDEO_STANDARDS.includes(word)) {
          fail(line, `unknown video standard '${word}' (expected ${VIDEO_STANDARDS.join(' or ')})`);
        }
        if (video !== null) fail(line, "'video' given twice in one target block");
        video = word;
        break;
      }
      case 'driveCycles': {
        if (!numTerm) fail(line, "'driveCycles' takes a number of CPU cycles per play call");
        const v = numOf(numTerm);
        if (!Number.isInteger(v) || v < 1) fail(line, `driveCycles must be a positive integer number of CPU cycles, got ${v}`);
        if (driveCycles !== null) fail(line, "'driveCycles' given twice in one target block");
        driveCycles = v;
        break;
      }
      case 'multispeed': {
        if (!numTerm) fail(line, "'multispeed' takes a number of play calls per video frame");
        const v = numOf(numTerm);
        if (!Number.isInteger(v) || v < 1) fail(line, `multispeed must be a positive integer, got ${v}`);
        if (multispeed !== null) fail(line, "'multispeed' given twice in one target block");
        multispeed = v;
        break;
      }
      default:
        fail(line, `expected 'video', 'driveCycles' or 'multispeed' in a target block, found '${name}'`);
    }
  }
  // driveCycles already IS the answer multispeed exists to compute, so
  // accepting both would leave two readings of the same block with no way
  // to tell which the author meant.
  if (driveCycles !== null && multispeed !== null) {
    fail(line, "'driveCycles' and 'multispeed' are two spellings of the same thing; give one, not both");
  }
  return { type: 'TargetStmt', line, video, driveCycles, multispeed };
}

// ADR 0003 D15's `chip 6581` / `chip 8580`: which SID model this song was
// written for/against. A bare NUMBER, no braces (unlike target), since
// there's exactly one setting. Optional; a song with none gets the
// player's default model. This only lets a song RECORD its choice instead
// of that choice living in transient player state, lost on save/export.
function buildChipStmt(ctx) {
  const line = ctx.start.line;
  const model = numOf(ctx.NUMBER());
  if (!CHIP_MODELS.includes(model)) {
    fail(line, `'chip' takes 6581 or 8580, got ${model}`);
  }
  return { type: 'ChipStmt', line, model };
}

// buildArrangement mirrors ast_visitor.go's buildArrangement(): flattens
// the grammar's arrangementElement+ into a []{ line, ref | group, repeat },
// one entry per top-level SIGIL or [...] group, each carrying its own `*n`
// repeat (default 1). Shared by the arrangement forms of patternDecl and
// voiceBody - see the .g4 file's arrangement rule comment for why this is
// a statement-level construct rather than a mini-notation bracket.
function buildArrangementElements(ctxs) {
  return ctxs.map((ec) => {
    const line = ec.start.line;
    const numTerm = ec.NUMBER();
    let repeat = 1;
    if (numTerm) {
      const rv = numOf(numTerm);
      if (!Number.isInteger(rv) || rv <= 0) fail(line, `an arrangement element's '*' repeat must be a positive integer, got ${rv}`);
      repeat = rv;
    }
    const sig = ec.SIGIL();
    if (sig) return { line, ref: orig(sig).value, repeat };
    return { line, group: buildArrangementElements(ec.arrangementElement()), repeat };
  });
}

function buildArrangement(ctx) {
  return buildArrangementElements(ctx.arrangementElement());
}

function buildPatternDecl(ctx) {
  const line = ctx.start.line;
  const name = ident(ctx.anyIdent());
  // Arrangement form (ADR 0003 D4): `pattern verse = $riffA $riffA $riffB`,
  // no quotes, no declared length - its length is always the sum of what
  // it references.
  const arrCtx = ctx.arrangement();
  if (arrCtx) {
    return { type: 'PatternDecl', line, name, length: null, each: null, elements: buildArrangement(arrCtx), source: null, docStart: null, docEnd: null };
  }
  const s = strInfo(ctx.str());
  // `|>` on a pattern declaration parses (see the grammar's patternDecl
  // rule) purely so this check can replace ANTLR's raw "mismatched input"
  // with a sentence that says where the transform belongs instead - ADR
  // 0001 decision 8 / known defect 3.
  if (ctx.pipelineStep().length > 0) {
    fail(line, `pipeline transforms ('|>') are not allowed on pattern declarations; add the transform on the voice that plays this pattern instead (e.g. \`voice 0 $${name} |> slow 4\`).`);
  }
  // The declared length in beats (`pattern bass 4 = "..."`), or null for a
  // proportional pattern - see the grammar's patternDecl comment and ADR
  // 0001's D1. Validated against the pattern's own summed content weight by
  // the semantic pass (analyze.js), not here: this layer only reports what
  // was written.
  const lengthTerm = ctx.NUMBER();
  const length = lengthTerm ? numOf(lengthTerm) : null;
  if (length !== null && !(length > 0)) {
    fail(line, `a pattern's declared length must be greater than 0 beats, got ${length}.`);
  }
  // `each N/D` (ADR 0003 D3): every token is one 1/D note, so the length is
  // derived rather than declared. Always numerator 1 in this ADR's own
  // examples; enforced here rather than in the grammar (a soft keyword like
  // every other field name) so a numerator other than 1 gets a specific
  // message instead of a generic parse failure.
  let each = null;
  const eachCtx = ctx.patternEach();
  if (eachCtx) {
    const eachName = ident(eachCtx.anyIdent());
    if (eachName !== 'each') fail(line, `expected 'each', found '${eachName}'`);
    const nums = eachCtx.NUMBER();
    const numerator = numOf(nums[0]);
    const denominator = numOf(nums[1]);
    if (numerator !== 1) fail(line, `'each' takes a fraction with numerator 1 (e.g. 'each 1/16'), got ${numerator}/${denominator}`);
    if (!Number.isInteger(denominator) || denominator <= 0) fail(line, `'each' needs a positive integer denominator, got ${denominator}`);
    each = { numerator, denominator };
  }
  return { type: 'PatternDecl', line, name, length, each, elements: null, source: s.value, docStart: s.docStart, docEnd: s.docEnd };
}

function buildTableDecl(ctx) {
  const line = ctx.start.line;
  const name = ident(ctx.anyIdent());
  return { type: 'TableDecl', line, name, tokens: buildTableRow(ctx.tableRow()) };
}

// === instrument_block ===

function buildInstrumentBlock(ctx) {
  const line = ctx.start.line;
  const name = ident(ctx.anyIdent());
  const fields = ctx.instrumentField().map(buildInstrumentField);
  return { type: 'InstrumentBlock', line, name, fields };
}

// extractSpeedDelay(tokens, line) -> { tokens, speed, delay } (ADR 0003
// D10): peels trailing `speed N`/`delay N` pairs (either order, either or
// both absent) off the END of a table row's flat token list - the grammar
// doesn't distinguish them from ordinary table entries (both are
// soft-keyword-shaped, like every other field name), so this is the same
// "structurally generic, validated by name here" split `each`/`meter`
// already use. Applies only to wave/arp/pw/filterTable/test/gate (D10's
// own field list); every other field's tokens are untouched.
function extractSpeedDelay(tokens, line) {
  let speed = null;
  let delay = null;
  const out = [...tokens];
  for (;;) {
    if (out.length < 2) break;
    const name = out[out.length - 2];
    const val = out[out.length - 1];
    if (name !== 'speed' && name !== 'delay') break;
    if (typeof val !== 'number') break;
    if (name === 'speed') {
      if (speed !== null) break;
      speed = val;
    } else {
      if (delay !== null) break;
      delay = val;
    }
    out.length -= 2;
  }
  if (speed !== null && !(Number.isInteger(speed) && speed > 0)) fail(line, `'speed' must be a positive integer, got ${speed}`);
  if (delay !== null && !(Number.isInteger(delay) && delay >= 0)) fail(line, `'delay' must be a non-negative integer, got ${delay}`);
  return { tokens: out, speed, delay };
}

// ADR 0003 D13: table verbs, an alternative to a bare token row for tables
// too long to read as N raw numbers - `set V`/`add V [repeat N]`/`wait N`/
// `loop`/`stop`, `end`-terminated, on any of the six table-shaped fields.
// Same six fields LOOP_TABLE_NAMES already names (every table-shaped
// field) - reused rather than retyped, same discipline as the exports
// above.
const TABLE_VERB_FIELDS = LOOP_TABLE_NAMES;
export const VERB_NAMES = ['set', 'add', 'wait', 'loop', 'stop'];

function buildTableVerbLine(ctx) {
  const line = ctx.start.line;
  const verb = ident(ctx.anyIdent());
  if (!VERB_NAMES.includes(verb)) fail(line, `unknown table verb '${verb}' (expected ${VERB_NAMES.join('/')})`);
  const tokens = ctx.tableRow() ? buildTableRow(ctx.tableRow()) : [];
  switch (verb) {
    case 'set': {
      if (tokens.length !== 1 || typeof tokens[0] !== 'number') fail(line, "'set' expects 1 number");
      return { verb, line, value: tokens[0] };
    }
    case 'add': {
      if (tokens.length !== 1 && tokens.length !== 3) {
        fail(line, "'add' expects a number, optionally followed by 'repeat N'");
      }
      if (typeof tokens[0] !== 'number') fail(line, "'add' expects a number");
      let repeat = 1;
      if (tokens.length === 3) {
        if (tokens[1] !== 'repeat' || typeof tokens[2] !== 'number') {
          fail(line, "'add' expects 'repeat N' after its value");
        }
        if (!Number.isInteger(tokens[2]) || tokens[2] < 1) fail(line, `'repeat' must be a positive integer, got ${tokens[2]}`);
        repeat = tokens[2];
      }
      return { verb, line, value: tokens[0], repeat };
    }
    case 'wait': {
      if (tokens.length !== 1 || typeof tokens[0] !== 'number') fail(line, "'wait' expects 1 number");
      if (!Number.isInteger(tokens[0]) || tokens[0] < 1) fail(line, `'wait' must be a positive integer, got ${tokens[0]}`);
      return { verb, line, count: tokens[0] };
    }
    case 'loop':
    case 'stop':
      if (tokens.length !== 0) fail(line, `'${verb}' takes no arguments`);
      return { verb, line };
    default:
      return fail(line, `unknown table verb '${verb}'`);
  }
}

function buildInstrumentField(ctx) {
  const line = ctx.start.line;
  const name = ident(ctx.anyIdent());
  if (!ctx.tableRow()) {
    // ADR 0003 D13: the verb-block form - see the grammar's own comment on
    // why `ctx.tableRow()` being null is exactly how the two alternatives
    // of `instrumentField` are told apart here.
    if (!TABLE_VERB_FIELDS.includes(name)) {
      fail(line, `'${name}' can't use the verb-block form (only ${TABLE_VERB_FIELDS.join('/')} can)`);
    }
    const verbs = ctx.tableVerbLine().map(buildTableVerbLine);
    return { type: name, line, verbs };
  }
  let tokens = buildTableRow(ctx.tableRow());
  let speed = null;
  let delay = null;
  if (['wave', 'arp', 'pw', 'filterTable', 'test'].includes(name) || (name === 'gate' && tokens.length > 1)) {
    ({ tokens, speed, delay } = extractSpeedDelay(tokens, line));
  }
  switch (name) {
    case 'wave':
    case 'arp':
    case 'pw':
    case 'filterTable':
    case 'test':
      return { type: name, line, tokens, speed, delay };
    case 'adsr': {
      if (tokens.length !== 4 || tokens.some((t) => typeof t !== 'number')) fail(line, 'adsr expects 4 numbers');
      const [a, d, s, r] = tokens;
      return { type: 'adsr', line, a, d, s, r };
    }
    case 'vib': {
      if (tokens.length !== 2 || tokens.some((t) => typeof t !== 'number')) fail(line, 'vib expects 2 numbers');
      const [depth, rate] = tokens;
      return { type: 'vib', line, depth, rate };
    }
    case 'vibTri': {
      // `vibTri CMP SPEED DELAY` - GoatTracker's own vibrato (see
      // fieldDocs.js). SPEED may be the shift-derived `sN` spelling the fx
      // lane already uses, which arrives here as an identifier token.
      if (tokens.length !== 3) fail(line, 'vibTri expects 3 values (CMP SPEED DELAY)');
      const [cmp, speed, delay] = tokens;
      if (typeof cmp !== 'number' || typeof delay !== 'number') fail(line, 'vibTri expects numbers for CMP and DELAY');
      const speedValue = parseVibTriSpeed(speed, line);
      return { type: 'vibTri', line, cmp, ...speedValue, delay };
    }
    case 'slide': {
      if (tokens.length !== 1 || typeof tokens[0] !== 'number') fail(line, 'slide expects 1 number');
      return { type: 'slide', line, time: tokens[0] };
    }
    case 'loop':
      return buildLoopField(line, tokens);
    case 'gate':
      if (tokens.length === 1) {
        const t = tokens[0];
        if (typeof t === 'number') return { type: 'gate', line, value: t };
        if (GATE_MODES.includes(t)) return { type: 'gate', line, mode: t };
        // A single `!N` repeat token is an ordinary one-entry gate TABLE, not
        // a malformed scalar: `!N` is the same table-repeat shape every other
        // instrument field accepts (see buildTableToken), so `gate 1!128`
        // must parse just like `gate 1!128 0!4`. Fall through to the table
        // form; anything that is neither a number, a mode, nor a repeat is
        // a real error.
        if (!(t && typeof t === 'object' && 'repeat' in t)) {
          fail(line, `'gate' expects a number, or one of ${GATE_MODES.join('/')}, found '${t}'`);
        }
      }
      return { type: 'gate', line, tokens, speed, delay };
    case 'hardRestart': {
      // ADR 0003 D11: the readable spelling of 'gate -N' (GoatTracker's
      // gatetimer) - compiles to the exact same AST shape a numeric
      // 'gate -N' field would, so it is a spelling, not a third gate
      // representation compile.js/instrument.js need to know about.
      if (tokens.length !== 1 || typeof tokens[0] !== 'number') fail(line, "'hardRestart' expects 1 number of frames");
      const n = tokens[0];
      if (!Number.isInteger(n) || n <= 0) fail(line, `'hardRestart' must be a positive integer number of frames, got ${n}`);
      return { type: 'gate', line, value: -n };
    }
    default:
      return fail(line, `unknown instrument field '${name}' (expected ${INSTRUMENT_FIELDS.join('/')})`);
  }
}

// parseVibTriSpeed(token) -> { speed } for a plain number, or { speedShift }
// for the `sN` spelling ("one semitone's worth of frequency at the current
// note, shifted right by N"), which is note-dependent and so cannot be a
// constant - exactly the same parameter form fx.js's parseParam() accepts,
// because it is the same gplay.c speed-table encoding.
function parseVibTriSpeed(token, line) {
  if (typeof token === 'number') return { speed: token };
  if (typeof token === 'string' && /^s[0-9]+$/.test(token)) return { speedShift: Number(token.slice(1)) };
  return fail(line, "vibTri's SPEED must be a number or a shift-derived 'sN'");
}

function buildLoopField(line, tokens) {
  if (tokens.length === 1 && typeof tokens[0] === 'number') {
    if (!Number.isInteger(tokens[0])) fail(line, `expected an integer, got ${tokens[0]}`);
    return { type: 'loop', line, value: tokens[0] };
  }
  const pairs = [];
  for (let i = 0; i < tokens.length; i += 2) {
    const table = tokens[i];
    const value = tokens[i + 1];
    if (typeof table !== 'string' || typeof value !== 'number') fail(line, 'malformed loop pairs');
    if (!LOOP_TABLE_NAMES.includes(table)) fail(line, `unknown loop table '${table}' (expected ${LOOP_TABLE_NAMES.join('/')})`);
    pairs.push({ table, value });
  }
  if (pairs.length === 0) fail(line, `expected an integer, or one or more TABLE INT pairs (${LOOP_TABLE_NAMES.join('/')})`);
  return { type: 'loop', line, pairs };
}

// === voice_stmt ===

function buildVoiceStmt(ctx) {
  const line = ctx.start.line;
  const index = numOf(ctx.NUMBER());
  const instrumentName = ctx.anyIdent() ? ident(ctx.anyIdent()) : null;
  const body = ctx.voiceBody();
  if (body.str() || body.SIGIL()) {
    const notes = buildNotesValue(body);
    const tail = body.voiceShortTail();
    const inlineAttrs = tail.inlineAttr().map(buildInlineAttr);
    const transforms = tail.pipelineStep().map((p) => buildTransform(p.transform()));
    return { type: 'VoiceStmt', mode: 'short', line, index, instrumentName, notes, arrangement: null, inlineAttrs, transforms };
  }
  if (body.arrangement()) {
    const arrangement = buildArrangement(body.arrangement());
    const tail = body.voiceShortTail();
    const inlineAttrs = tail.inlineAttr().map(buildInlineAttr);
    const transforms = tail.pipelineStep().map((p) => buildTransform(p.transform()));
    return { type: 'VoiceStmt', mode: 'short', line, index, instrumentName, notes: null, arrangement, inlineAttrs, transforms };
  }
  const fields = body.voiceField().map(buildVoiceField);
  return { type: 'VoiceStmt', mode: 'extended', line, index, instrumentName, fields };
}

// buildNotesValue accepts either a StrContext-bearing ctx (has .str()) or
// one with .SIGIL() -- used both for voiceBody (short mode) and the
// 'notes' extended-mode field.
function buildNotesValue(ctx) {
  if (ctx.str && ctx.str()) {
    const s = strInfo(ctx.str());
    return { isSigil: false, source: s.value, docStart: s.docStart, docEnd: s.docEnd };
  }
  return { isSigil: true, sigilName: orig(ctx.SIGIL()).value };
}

function buildInlineAttr(ctx) {
  const name = ident(ctx.anyIdent());
  if (ctx.str()) {
    const s = strInfo(ctx.str());
    return { name, isPattern: true, source: s.value, docStart: s.docStart, docEnd: s.docEnd };
  }
  const values = ctx.valueList().tableToken().map(buildTableToken);
  return { name, values };
}

function buildVoiceField(ctx) {
  const line = ctx.start.line;
  if (ctx.pipelineStep()) {
    return { type: 'pipeline', line, transforms: [buildTransform(ctx.pipelineStep().transform())] };
  }
  const name = ident(ctx.anyIdent());
  const vv = ctx.voiceFieldValue();
  switch (name) {
    case 'notes':
      return { type: 'notes', line, ...buildNotesValue(vv) };
    case 'insts': {
      if (!vv.str()) fail(line, "'insts' expects a quoted mini-notation pattern");
      const s = strInfo(vv.str());
      return { type: 'insts', line, source: s.value, docStart: s.docStart, docEnd: s.docEnd };
    }
    // `fx` - the per-row effect-command lane (see src/dsl/fx.js). Shaped
    // exactly like `insts`: a parallel mini-notation string, kept as raw
    // source text here and only parsed into effect objects at interpret
    // time, same as every other pattern-valued field.
    case 'fx': {
      if (!vv.str()) fail(line, "'fx' expects a quoted mini-notation pattern");
      const s = strInfo(vv.str());
      return { type: 'fx', line, source: s.value, docStart: s.docStart, docEnd: s.docEnd };
    }
    case 'wave': {
      if (!vv.str()) fail(line, "'wave' expects a quoted mini-notation pattern");
      const s = strInfo(vv.str());
      return { type: 'wave', line, source: s.value, docStart: s.docStart, docEnd: s.docEnd };
    }
    case 'reg': {
      const identOrHex = vv.anyIdent ? vv.anyIdent() : null;
      const hex = vv.HEX ? vv.HEX() : null;
      let target;
      if (hex) target = { kind: 'hex', value: orig(hex).value };
      else if (identOrHex) target = { kind: 'ident', name: ident(identOrHex) };
      else return fail(line, 'expected a register name or a 0xNN hex address');
      if (!vv.str()) fail(line, "'reg' expects a quoted register-value mini-notation pattern");
      const source = strInfo(vv.str()).value;
      return { type: 'reg', line, target, source };
    }
    case 'vib': {
      const nums = vv.NUMBER();
      if (!nums || nums.length !== 2) fail(line, "'vib' expects 2 numbers");
      return { type: 'vib', line, depth: numOf(nums[0]), rate: numOf(nums[1]) };
    }
    case 'slide': {
      const nums = vv.NUMBER();
      if (!nums || nums.length !== 1) fail(line, "'slide' expects 1 number");
      return { type: 'slide', line, time: numOf(nums[0]) };
    }
    case 'gate': {
      const nums = vv.NUMBER();
      if (!nums || nums.length !== 1) fail(line, "'gate' expects 1 number");
      return { type: 'gate', line, value: numOf(nums[0]) };
    }
    case 'sync':
    case 'ring':
      return { type: name, line, value: buildBoolFieldValue(vv, line) };
    default:
      return fail(line, `unknown voice field '${name}' (expected ${VOICE_FIELDS.join('/')}, or a '|>' pipeline)`);
  }
}

function buildBoolFieldValue(vv, line) {
  if (vv.str()) return { kind: 'pattern', source: strInfo(vv.str()).value };
  const nums = vv.NUMBER();
  if (nums && nums.length === 1) return { kind: 'number', value: numOf(nums[0]) };
  return fail(line, 'expected a number (0/1) or a quoted mini-notation pattern');
}

// === pipeline / transform ===

function buildTransform(ctx) {
  const line = ctx.start.line;
  const name = ident(ctx.anyIdent());
  const argCtx = ctx.transformArg();
  const refCtx = ctx.transformRef();
  const arg = argCtx ? buildTransformArg(argCtx) : null;
  const fn = refCtx ? buildTransformRef(refCtx) : null;

  const needArg = () => {
    if (!arg) fail(line, `'${name}' expects an argument`);
    return arg;
  };
  const needIntArg = () => {
    const a = needArg();
    if (a.value !== undefined && !Number.isInteger(a.value)) fail(line, `expected an integer, got ${a.value}`);
    return a;
  };
  const needFn = () => {
    if (!fn) fail(line, `'${name}' expects a transform (or parenthesized transform)`);
    return fn;
  };
  const needStringArg = (label) => {
    if (!argCtx || !argCtx.str()) fail(line, `expected ${label}`);
    return strInfo(argCtx.str()).value;
  };

  switch (name) {
    case 'rev':
      return { type: 'rev', line };
    case 'fast':
      return { type: 'fast', line, n: needArg() };
    case 'slow':
      return { type: 'slow', line, n: needArg() };
    case 'every':
      return { type: 'every', line, n: needIntArg(), fn: needFn() };
    case 'iter':
      return { type: 'iter', line, n: needIntArg() };
    case 'ply':
      return { type: 'ply', line, n: needIntArg() };
    case 'off':
      return { type: 'off', line, time: needArg(), fn: needFn() };
    case 'jux':
      return { type: 'jux', line, fn: needFn() };
    case 'degradeBy':
      return { type: 'degradeBy', line, p: needArg() };
    case 'sometimesBy':
      return { type: 'sometimesBy', line, p: needArg(), fn: needFn() };
    case 'add':
      return { type: 'add', line, n: needArg() };
    case 'struct':
      return { type: 'struct', line, source: needStringArg('a quoted boolean rhythm pattern, e.g. "x ~ x x"') };
    case 'mask':
      return { type: 'mask', line, source: needStringArg('a quoted boolean mask pattern, e.g. "1 0 1 1"') };
    case 'chunk':
      return { type: 'chunk', line, n: needIntArg(), fn: needFn() };
    case 'palindrome':
      return { type: 'palindrome', line };
    case 'linger':
      return { type: 'linger', line, frac: needArg() };
    case 'rot':
      return { type: 'rot', line, n: needIntArg() };
    // `someBarsBy` is 3.0.0's name for `someCyclesBy` (ADR 0001's D3:
    // transforms count bars). Both spellings build the SAME node type, so
    // nothing downstream (interpret.js, the Go AST, generators) has to
    // learn a second case for what is one operation under two names.
    case 'someBarsBy':
    case 'someCyclesBy':
      return { type: 'someCyclesBy', line, p: needArg(), fn: needFn() };
    case 'superimpose':
      return { type: 'superimpose', line, fn: needFn() };
    case 'scale':
      return { type: 'scale', line, source: needStringArg('a quoted scale spec, e.g. "c4:minor"') };
    default:
      return fail(line, `unknown transform '${name}' (expected ${TRANSFORMS.join('/')})`);
  }
}

function buildTransformArg(ctx) {
  if (ctx.str()) return { source: strInfo(ctx.str()).value };
  return { value: numOf(ctx.NUMBER()) };
}

function buildTransformRef(ctx) {
  return buildTransform(ctx.transform());
}

// === filter_block ===

function buildFilterBlock(ctx) {
  const line = ctx.start.line;
  const fields = ctx.filterField().map(buildFilterField);
  return { type: 'FilterBlock', line, fields };
}

function buildFilterField(ctx) {
  const line = ctx.start.line;
  const name = ident(ctx.anyIdent());
  const fv = ctx.filterFieldValue();
  switch (name) {
    case 'cutoff':
      return { type: 'cutoff', line, expr: buildFilterExpr(fv, 'cutoff expects a signal expression or \'switch cycle\' block') };
    case 'resonance': {
      const ai = fv.anyIdent();
      if (!fv.numberList() && (!ai || ai.length < 1)) fail(line, 'resonance expects a number, a signal expression, or \'switch cycle\' block');
      return { type: 'resonance', line, expr: buildFilterExpr(fv, 'resonance expects a number, a signal expression, or \'switch cycle\' block') };
    }
    case 'mode': {
      const ai = fv.anyIdent();
      if (!ai || ai.length < 1) fail(line, "mode expects an identifier or 'switch bar' block");
      return { type: 'mode', line, expr: buildFilterExpr(fv, 'mode expects an identifier or \'switch cycle\' block') };
    }
    case 'route': {
      const ai = fv.anyIdent();
      if (!fv.numberList() && (!ai || ai.length < 1)) fail(line, 'route expects one or more integers or \'switch cycle\' block');
      return { type: 'route', line, expr: buildFilterExpr(fv, 'route expects one or more integers or \'switch cycle\' block') };
    }
    default:
      return fail(line, `unknown filter field '${name}' (expected ${FILTER_FIELDS.join('/')})`);
  }
}

// buildFilterExpr(fv: FilterFieldValueContext) -> SignalExpr, SwitchBlock,
// or NumberListExpr node - all three filterFieldValue shapes handled
// uniformly here; which are legal for a given field name is enforced by
// buildFilterField's callers above, not here.
function buildFilterExpr(fv, errMsg) {
  const line = fv.start.line;
  const ai = fv.anyIdent();
  if (ai && ai.length === 2) {
    if (ident(ai[0]) !== 'switch') fail(line, "expected 'switch'");
    // `switch bar` is the 3.0.0 spelling (ADR 0001's D3: a "cycle" is a
    // "bar"); `switch cycle` still parses so no existing song breaks, and
    // the semantic pass is what nudges toward the new word.
    if (!SWITCH_UNIT_WORDS.includes(ident(ai[1]))) fail(line, "expected 'bar' (or the pre-3.0.0 spelling 'cycle')");
    const cases = fv.switchCase().map(buildSwitchCase);
    return { type: 'SwitchBlock', line, cases };
  }
  if (ai && ai.length === 1) {
    return buildSignalExprFromParts(line, ai[0], fv.signalRangeOpt(), fv.signalSpeedMod());
  }
  const nl = fv.numberList();
  if (nl) return buildNumberListExpr(nl);
  return fail(line, errMsg);
}

function buildNumberListExpr(nl) {
  const line = nl.start.line;
  const values = nl.NUMBER().map((n) => numOf(n));
  return { type: 'NumberListExpr', line, values };
}

function buildSignalExprFromParts(line, signalIdentCtx, rangeCtx, speedCtxs) {
  const signal = ident(signalIdentCtx);
  let range = null;
  if (rangeCtx && rangeCtx.NUMBER().length === 2) {
    range = { lo: numOf(rangeCtx.NUMBER(0)), hi: numOf(rangeCtx.NUMBER(1)) };
  }
  const speeds = (speedCtxs || []).map((s) => {
    if (s.STAR()) return { op: '*', value: numOf(s.NUMBER()) };
    if (s.SLASH()) return { op: '/', value: numOf(s.NUMBER()) };
    return { op: '>>', value: numOf(s.NUMBER()) };
  });
  return { type: 'SignalExpr', line, signal, range, speeds };
}

function buildSwitchCase(ctx) {
  const line = ctx.start.line;
  const sc = ctx.switchCond();
  let cond;
  if (sc.LT()) cond = { op: '<', value: numOf(sc.NUMBER()) };
  else if (sc.GE()) cond = { op: '>=', value: numOf(sc.NUMBER()) };
  else if (sc.anyIdent()) {
    if (ident(sc.anyIdent()) !== 'else') fail(line, "expected a switch case ('< N', '>= N', or 'else')");
    cond = { op: 'else' };
  } else {
    return fail(line, "expected a switch case ('< N', '>= N', or 'else')");
  }
  const nl = ctx.numberList();
  const value = nl ? buildNumberListExpr(nl) : buildSignalExprFromParts(line, ctx.anyIdent(), ctx.signalRangeOpt(), ctx.signalSpeedMod());
  return { line, cond, value };
}
