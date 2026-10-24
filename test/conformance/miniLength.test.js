// JS half of the mini-notation LENGTH conformance suite (the Go half is
// minilength_test.go). Both read the same mini-length.expected.json golden,
// generated from this implementation by generate-golden.mjs - see that
// file's second half for why length arithmetic needs the same
// two-implementations-one-golden treatment the parser already gets.
import fs from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { contentBeats, declaredLengthCandidates, expandHolds } from '../../src/dsl/miniLength.js';

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const golden = JSON.parse(fs.readFileSync(path.join(__dirname, 'mini-length.expected.json'), 'utf8'));

describe('conformance: mini-notation length arithmetic matches the shared golden', () => {
  it('has cases covering holds, weights, groups, alternation and stacks', () => {
    expect(golden.length).toBeGreaterThanOrEqual(15);
    expect(golden.some((c) => c.source.includes('_'))).toBe(true);
    expect(golden.some((c) => c.ok === false)).toBe(true);
  });

  for (const c of golden) {
    it(`${JSON.stringify(c.source)} matches its golden`, () => {
      if (!c.ok) {
        // expandHolds() alone doesn't scan (and so can't throw) a hold-free
        // string - see its own "return source unchanged" fast path - so a
        // case that is only invalid for a reason other than hold placement
        // (e.g. ADR 0003 D1's mixed /N-and-@n check) only surfaces once
        // something that always scans, like contentBeats(), runs.
        expect(() => contentBeats(expandHolds(c.source))).toThrow();
        return;
      }
      const expanded = expandHolds(c.source);
      expect(expanded).toBe(c.expanded);
      expect(contentBeats(expanded)).toBeCloseTo(c.contentBeats, 9);
      expect(JSON.parse(JSON.stringify(declaredLengthCandidates(expanded)))).toEqual(c.candidates);
    });
  }
});
