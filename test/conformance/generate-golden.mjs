// One-time (re-run manually whenever a fixture changes) generator for the
// conformance suite's golden files: runs the JS reference lexer+parser
// (src/dsl/{lexer,parser}.js, the canonical implementation) over every fixtures/*.dsl file and writes a sibling
// fixtures/<name>.expected.json golden. Both the JS conformance test
// (conformance.test.js) and the Go conformance test
// (conformance_test.go) read these SAME golden files and assert their own
// parse output matches - that's the actual cross-language equivalence
// check, not a promise that re-running this script is itself a test.
//
// Usage: node test/conformance/generate-golden.mjs
import fs from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { tokenize } from '../../src/dsl/lexer.js';
import { parseProgram } from '../../src/dsl/parser.js';
import { contentBeats, declaredLengthCandidates, expandHolds } from '../../src/dsl/miniLength.js';

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const fixturesDir = path.join(__dirname, 'fixtures');

// A golden is either { ok: true, ast } (successful parse -- the exact AST
// parser.js produced, JSON-round-tripped) or { ok: false, line, message }
// (a lex/parse error -- only the 1-based line number is asserted for
// cross-language equivalence; the message text itself is JS-specific
// prose that Go doesn't need to match byte-for-byte).
for (const file of fs.readdirSync(fixturesDir).sort()) {
  if (!file.endsWith('.dsl')) continue;
  const src = fs.readFileSync(path.join(fixturesDir, file), 'utf8');
  let golden;
  try {
    const tokens = tokenize(src);
    const ast = parseProgram(tokens);
    golden = { ok: true, ast };
  } catch (e) {
    const m = /^line (\d+):/.exec(e.message);
    golden = { ok: false, line: m ? Number(m[1]) : null, message: e.message };
  }
  const outPath = path.join(fixturesDir, file.replace(/\.dsl$/, '.expected.json'));
  fs.writeFileSync(outPath, `${JSON.stringify(golden, null, 2)}\n`);
  console.log('wrote', path.relative(process.cwd(), outPath));
}

// The second golden set: mini-notation LENGTH arithmetic (src/dsl/
// miniLength.js and its Go port pkg/dsl/minilength.go). Same mechanism and
// same reason as the parser goldens above - a pattern's declared length
// (ADR 0001's D1) is computed on both sides of the pipeline, by the
// compiler that checks it and by Go generators that write it, so the two
// must agree token for token. A case that legitimately
// fails to scan records its error instead of its lengths; only the fact
// that BOTH ports fail is asserted, not the message text (same convention
// as the parse-error fixtures).
const miniCases = JSON.parse(fs.readFileSync(path.join(__dirname, 'mini-length.cases.json'), 'utf8'));
const miniGolden = miniCases.map((source) => {
  try {
    return {
      source,
      ok: true,
      contentBeats: contentBeats(expandHolds(source)),
      candidates: declaredLengthCandidates(expandHolds(source)),
      expanded: expandHolds(source),
    };
  } catch (e) {
    return { source, ok: false };
  }
});
const miniOut = path.join(__dirname, 'mini-length.expected.json');
fs.writeFileSync(miniOut, `${JSON.stringify(miniGolden, null, 2)}\n`);
console.log('wrote', path.relative(process.cwd(), miniOut));
