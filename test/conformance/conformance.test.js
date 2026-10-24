// Cross-language conformance suite: this file and conformance_test.go (the
// Go side) both read the SAME
// fixtures/*.dsl source files and the SAME fixtures/*.expected.json
// golden files, generated once from this JS reference implementation via
// generate-golden.mjs (see that file's header for the mechanism and why
// "both languages compared against one golden" transitively proves
// "Go and JS agree").
//
// This half only proves the JS lexer+parser still produces exactly what
// the golden recorded (regression protection against JS-side drift); the
// Go half is what actually proves the Go port matches. Re-run
// `node test/conformance/generate-golden.mjs` and re-review the diff
// whenever a fixture or the grammar itself intentionally changes.
import fs from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { tokenize } from '../../src/dsl/lexer.js';
import { parseProgram } from '../../src/dsl/parser.js';

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const fixturesDir = path.join(__dirname, 'fixtures');

const fixtures = fs
  .readdirSync(fixturesDir)
  .filter((f) => f.endsWith('.dsl'))
  .sort();

describe('conformance: JS lexer+parser matches the shared golden fixtures', () => {
  it('has at least one fixture of each documented kind', () => {
    // A cheap guard against the fixture directory silently emptying out --
    // the actual per-fixture assertions below are what matter.
    expect(fixtures.length).toBeGreaterThanOrEqual(15);
  });

  for (const file of fixtures) {
    it(`${file} matches its golden`, () => {
      const src = fs.readFileSync(path.join(fixturesDir, file), 'utf8');
      const goldenPath = path.join(fixturesDir, file.replace(/\.dsl$/, '.expected.json'));
      const golden = JSON.parse(fs.readFileSync(goldenPath, 'utf8'));

      if (golden.ok) {
        const ast = parseProgram(tokenize(src));
        // JSON round-trip on both sides so the comparison is over plain
        // JSON values (matching what the Go side compares against),
        // not live object identity/prototype details.
        expect(JSON.parse(JSON.stringify(ast))).toEqual(golden.ast);
      } else {
        let threw = null;
        try {
          parseProgram(tokenize(src));
        } catch (e) {
          threw = e;
        }
        expect(threw).not.toBeNull();
        expect(threw.message).toMatch(new RegExp(`^line ${golden.line}:`));
      }
    });
  }
});
