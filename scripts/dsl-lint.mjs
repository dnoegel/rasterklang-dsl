#!/usr/bin/env node
// Diagnostics for one or more DSL files, without compiling them.
//
// This is where ADR 0001's migration aid lives (decision 10): a song whose
// `beats` value looks like it was used as a melody-stretch lever gets a
// hint pointing at declared pattern lengths instead. It reports everything
// else the semantic pass finds too - declared-length mismatches, the
// arrangement's true period, the four known-defect warnings from ADR 0001
// section 6 - because a checker that only ever says one thing is a checker
// nobody runs.
//
// It deliberately uses src/dsl/index.js's checkSource(), which needs
// nothing but this repo (no pattern engine, no SID emulation), so linting a
// file is exactly as cheap as parsing it.
//
// Usage:
//   node scripts/dsl-lint.mjs song.dsl [more.dsl ...]
//   npm run lint:dsl -- song.dsl
//
// Exit status: 1 if any file produced an error-severity diagnostic, else 0.
// Warnings and info never fail the run - unusual time signatures and long
// polymetric periods are legal music (ADR 0001's D5/decision 10).
import fs from 'node:fs';
import { checkSource } from '../src/dsl/index.js';

const files = process.argv.slice(2);
if (files.length === 0) {
  console.error('usage: node scripts/dsl-lint.mjs FILE.dsl [FILE.dsl ...]');
  process.exit(2);
}

const ORDER = { error: 0, warning: 1, info: 2 };
let failed = false;

for (const file of files) {
  let diagnostics;
  try {
    diagnostics = checkSource(fs.readFileSync(file, 'utf8'));
  } catch (e) {
    console.error(`${file}: ${e.message}`);
    failed = true;
    continue;
  }
  const sorted = [...diagnostics].sort((a, b) => ORDER[a.severity] - ORDER[b.severity] || a.line - b.line);
  for (const d of sorted) {
    console.log(`${file}:${d.line}: ${d.severity}: ${d.message}`);
    if (d.severity === 'error') failed = true;
  }
  if (sorted.length === 0) console.log(`${file}: ok`);
}

process.exit(failed ? 1 : 0);
