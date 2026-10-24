#!/usr/bin/env node
// Generates docs/dsl-spec.generated.json, the machine-readable, versioned
// DSL v2 spec. Every field/keyword/transform name and its documentation
// (signature/description/example) comes straight from the code:
// src/dsl/keywords.js (name lists, re-exported from the parser and
// registerMap.js, never re-declared), src/dsl/fieldDocs.js (the prose) and
// src/dsl/version.js (DSL_VERSION). If the output is ever wrong, fix one of
// those source files; test/dsl/generatedSpec.test.js fails if the committed
// JSON drifts from what this script produces.
//
// Usage:
//   node scripts/generate-dsl-spec.mjs            # regenerate docs/dsl-spec.generated.json
//   npm run generate:dsl-spec                     # same, via package.json

import { writeFileSync } from 'node:fs';
import { dirname, join } from 'node:path';
import { fileURLToPath } from 'node:url';

import {
  STATEMENT_KEYWORDS,
  INSTRUMENT_FIELDS,
  VOICE_FIELDS,
  TRANSFORMS,
  FILTER_FIELDS,
  FILTER_MODES,
  SIGNAL_NAMES,
  WAVEFORM_NAMES,
  VOICE_REG_TARGETS,
  FILTER_REG_TARGETS,
  STRUCTURAL_KEYWORDS,
  MINI_NOTATION_TOKENS,
} from '../src/dsl/keywords.js';
import { FIELD_DOCS } from '../src/dsl/fieldDocs.js';
import { DSL_VERSION } from '../src/dsl/version.js';

const __dirname = dirname(fileURLToPath(import.meta.url));

// withDocs(names) -> [{ name, signature?, description?, example? }, ...] -
// enriches each bare name with its FIELD_DOCS entry when one exists (every
// STATEMENT_KEYWORDS/INSTRUMENT_FIELDS/VOICE_FIELDS/TRANSFORMS/FILTER_FIELDS
// name is expected to have one; a handful of value-vocabulary names - filter modes, signal names, waveform
// names, register targets, most structural keywords - don't have their own
// FIELD_DOCS entry, since they're not fields/keywords a user types a value
// *into*, and are left as bare names rather than invented documentation).
function withDocs(names) {
  return names.map((name) => ({ name, ...(FIELD_DOCS[name] ?? {}) }));
}

// buildDslSpec() -> the plain JS object serialized to
// docs/dsl-spec.generated.json - exported (not just used by this file's own
// CLI entry point below) so test/dsl/generatedSpec.test.js can call it
// directly and diff against the committed file without shelling out.
export function buildDslSpec() {
  return {
    dslVersion: DSL_VERSION,
    generatedBy: 'scripts/generate-dsl-spec.mjs',
    sources: [
      'src/dsl/parser.js',
      'src/dsl/keywords.js',
      'src/dsl/fieldDocs.js',
      'src/dsl/version.js',
      'rasterklang-score/compiler/registerMap.js',
    ],
    statementKeywords: withDocs(STATEMENT_KEYWORDS),
    instrumentFields: withDocs(INSTRUMENT_FIELDS),
    voiceFields: withDocs(VOICE_FIELDS),
    transforms: withDocs(TRANSFORMS),
    filterFields: withDocs(FILTER_FIELDS),
    filterModes: withDocs(FILTER_MODES),
    signalNames: withDocs(SIGNAL_NAMES),
    waveformNames: withDocs(WAVEFORM_NAMES),
    voiceRegTargets: withDocs(VOICE_REG_TARGETS),
    filterRegTargets: withDocs(FILTER_REG_TARGETS),
    structuralKeywords: withDocs(STRUCTURAL_KEYWORDS),
    // The vocabulary inside a quoted pattern string (rest/hold/weight/...)
    // - not lexer keywords, but the part of the language a user types most,
    // and since 3.0.0 it has a token (`_`, the hold) that no other table
    // mentions. See keywords.js's MINI_NOTATION_TOKENS.
    miniNotationTokens: withDocs(MINI_NOTATION_TOKENS),
  };
}

// serializeDslSpec() -> the exact JSON text that should be on disk -
// factored out (not just JSON.stringify'd inline in two places) so the
// generator and the staleness-check test are guaranteed to compare using
// identical formatting, not just equivalent data.
export function serializeDslSpec() {
  return `${JSON.stringify(buildDslSpec(), null, 2)}\n`;
}

function main() {
  const outPath = join(__dirname, '..', 'docs', 'dsl-spec.generated.json');
  writeFileSync(outPath, serializeDslSpec());
  console.log(`Wrote ${outPath}`);
}

// Only run as a CLI when executed directly (`node scripts/generate-dsl-spec.mjs`),
// not when imported by the staleness-check test.
if (import.meta.url === `file://${process.argv[1]}`) {
  main();
}
