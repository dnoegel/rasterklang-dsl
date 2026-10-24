// Canonical keyword/name tables for DSL v2 (docs/dsl-grammar-spec.md), so
// editor integrations (highlighting, completion) can read the reserved words
// and known names from one place.
//
// Every list below is a RE-EXPORT of the real thing, not a separately
// maintained copy, because hand-maintained copies drift silently.
// STATEMENT_KEYWORDS/INSTRUMENT_FIELDS/VOICE_FIELDS/TRANSFORMS/FILTER_FIELDS
// live in astVisitor.js (re-exported via parser.js), right next to the switch
// statements that validate them; VOICE_REG_TARGETS/FILTER_REG_TARGETS/
// WAVEFORM_NAMES are derived from registerMap.js's objects via Object.keys().
import {
  STATEMENT_KEYWORDS,
  INSTRUMENT_FIELDS,
  VOICE_FIELDS,
  TRANSFORMS,
  FILTER_FIELDS,
  LOOP_TABLE_NAMES,
  GATE_MODES,
  VERB_NAMES,
  VIDEO_STANDARDS,
} from './parser.js';
// registerMap.js (SID register-map knowledge) lives in rasterklang-score,
// not here, but keywords.js still needs its name vocabulary.
import { VOICE_REGISTER_OFFSETS, FILTER_REGISTERS, WAVEFORM_BITS } from 'rasterklang-score/compiler/registerMap.js';
import { FX_OPS } from './fx.js';

// LOOP_TABLE_NAMES doubles as the canonical list of TABLE-VALUED fields (the
// ones parsed as table rows, and so the ones a `$NAME` table
// reference is valid in) - it originates in instrument.js, re-exported here
// rather than retyped, same discipline as the lists above.
export {
  STATEMENT_KEYWORDS,
  INSTRUMENT_FIELDS,
  VOICE_FIELDS,
  TRANSFORMS,
  FILTER_FIELDS,
  LOOP_TABLE_NAMES,
  // ADR 0003 D11's gate legato/retrigger, D13's table-verb names.
  GATE_MODES,
  VERB_NAMES,
};

// Short-mode inline_attr names (interpret.js's VOICE_LEVEL_INLINE_ATTR_NAMES
// plus INSTRUMENT_FIELDS): `wave:pulse adsr:0,4,10,4 sync:1`, sugar for an
// anonymous one-off instrument plus the two voice-level bit fields.
export const SHORT_MODE_INLINE_ATTRS = [...INSTRUMENT_FIELDS, 'sync', 'ring'];

// No source object exists for these (packFilterMode() takes named boolean
// args, not an enum). These are SID hardware facts (the filter mode bits),
// not expected to change, so a plain manually-maintained list is fine.
export const FILTER_MODES = ['lowpass', 'bandpass', 'highpass', 'none'];

// The `fx` voice field's effect-token names (2.15.0), re-exported from
// fx.js's FX_OPS - the same discipline as the lists above: derived from the
// real table the parser dispatches on, never retyped.
export const FX_OP_NAMES = Object.keys(FX_OPS);

// signals.js's exported continuous generators. Also manually maintained (no
// single collection in signals.js itself to derive from without risking
// picking up a future non-signal export via `Object.keys(module)`).
export const SIGNAL_NAMES = ['sine', 'saw', 'square', 'tri', 'rand', 'perlin'];

// song.js's VoiceChain.s()/wave table accepted names - derived directly from
// registerMap.js's WAVEFORM_BITS, the actual control-byte bit mapping.
export const WAVEFORM_NAMES = Object.keys(WAVEFORM_BITS);

// registerMap.js's VOICE_REGISTER_OFFSETS keys (voice-relative reg targets).
export const VOICE_REG_TARGETS = Object.keys(VOICE_REGISTER_OFFSETS);

// registerMap.js's FILTER_REGISTERS keys (shared, absolute reg targets).
export const FILTER_REG_TARGETS = Object.keys(FILTER_REGISTERS);

// Structural keywords that don't belong to one specific field/statement list
// above: block terminator, tempo's optional suffixes ('beats'/'offset'/
// 'meter'), and switch-block syntax.
// 'ramp'/'step'/'over'/'pingpong' are the table-generator words (`pw ramp
// 0..4095 step 127 over 68 pingpong`, ADR 0003 D9) - contextual, only
// meaningful inside a table row, but worth highlighting as keywords like
// the block/structural words above. 'speed'/'delay' after a table row are
// D10's per-table speed/delay pair - 'speed' doubles as a STATEMENT_KEYWORD
// already (the `speed N` clock spelling), so only 'delay' needs adding here.
// 'meter' is ADR 0003 D2's `meter N/D` tempo setting; 'each' is D3's
// `pattern NAME each 1/D = ...` step-sequencer declaration.
// 'bar' is 3.0.0's word for what a `switch` block dispatches on; 'cycle'
// is the pre-3.0.0 spelling of the same thing and still parses (ADR 0001's
// D3), so both stay in the highlighting vocabulary.
// 'video'/'driveCycles'/'multispeed'/VIDEO_STANDARDS are ADR 0002's
// target-block settings: contextual, only meaningful inside
// `target { ... }`, same status as 'ramp'/'step'/'over' inside a table row.
export const STRUCTURAL_KEYWORDS = [
  'end', 'beats', 'offset', 'meter', 'each', 'switch', 'bar', 'cycle', 'else',
  'ramp', 'step', 'over', 'pingpong', 'delay',
  'video', 'driveCycles', 'multispeed', ...VIDEO_STANDARDS,
];

// The mini-notation vocabulary INSIDE a quoted pattern string. These are
// not lexer keywords - the DSL grammar sees a pattern as one opaque STRING
// and the pattern engine scans it separately - but they are the part of the
// language a user types most, so an editor/completion consumer and the
// generated spec both need them named in one place. `_` (the hold) is new
// in 3.0.0; `/N` (durations, D1) was already representable via the `/`
// entry; `$name` (structural references, D5) reuses the SIGIL vocabulary
// elsewhere, not this list; `slideTo:` (D12) is new - a per-note atom
// prefix, unlike every other entry here which is a bare punctuation
// character. The rest are unchanged (ADR 0001's D6).
export const MINI_NOTATION_TOKENS = ['~', '_', '@', '!', '*', '/', '?', '[]', '<>', '{}', ',', '|', 'slideTo:'];

export const ALL_KEYWORDS = Array.from(
  new Set([
    ...STATEMENT_KEYWORDS,
    ...INSTRUMENT_FIELDS,
    ...VOICE_FIELDS,
    ...TRANSFORMS,
    ...FILTER_FIELDS,
    ...FILTER_MODES,
    ...SIGNAL_NAMES,
    ...WAVEFORM_NAMES,
    ...VOICE_REG_TARGETS,
    ...FILTER_REG_TARGETS,
    ...STRUCTURAL_KEYWORDS,
    ...GATE_MODES,
    ...VERB_NAMES,
  ]),
);
