# Cross-language conformance suite

Proves the Go side (`pkg/dsl/lexer.go` + `pkg/dsl/parser.go`) produces the
same result as the canonical JS implementation (`src/dsl/lexer.js` +
`src/dsl/parser.js`) on the same real inputs. Both parsers are generated
from one ANTLR grammar, but the token-replay layers and the AST-transform
visitors that turn a parse tree into each language's AST are hand-written
twice, so this suite is the anti-drift guarantee for everything the
grammar file alone cannot enforce.

## Mechanism

- `fixtures/*.dsl` - shared source fixtures, read by both languages.
- `fixtures/*.expected.json` - one golden file per fixture, generated
  **once** from the JS reference implementation via
  `generate-golden.mjs`. Each golden is either `{ ok: true, ast }` (the
  exact JSON-round-tripped AST `parseProgram()` produced) or
  `{ ok: false, line, message }` (a lex/parse error - only the line
  number is asserted cross-language; message text is JS-specific prose).
- `conformance.test.js` (vitest) re-parses every fixture with the JS
  implementation and asserts the result still matches its golden -
  regression protection against JS-side drift.
- `canonical.go` + `conformance_test.go` (Go) re-parses every fixture with
  the Go implementation, converts the resulting `*dsl.Program` to the same
  plain-JSON shape via `canonical.go`, and asserts it matches the same
  golden with `reflect.DeepEqual` - **this is what actually proves the Go
  port matches the JS reference**, since both halves are checked against
  one shared source of truth.

## The second golden set: mini-notation length arithmetic

Since 3.0.0 a pattern can declare its length in beats (ADR 0001's D1), and
that length is computed on both sides of the pipeline: `src/dsl/
miniLength.js` checks the length a song declares, `pkg/dsl/minilength.go`
is what a Go-side generator uses to WRITE one. Two
implementations of the same arithmetic is exactly the situation the fixture
corpus above exists for, so it gets the same treatment:

- `mini-length.cases.json` - the shared input strings (holds, weights,
  groups, alternation, stacks, replication, and cases that must fail).
- `mini-length.expected.json` - the golden, generated from the JS
  implementation by the same `generate-golden.mjs`, recording each case's
  hold-expanded text, content length, and declared-length candidates.
- `miniLength.test.js` (vitest) and `minilength_test.go` (Go) both assert
  against it.

## Regenerating goldens

Only do this when a fixture or the grammar itself intentionally changes,
and review the diff - a golden changing unexpectedly is exactly the drift
this suite exists to catch:

```
node test/conformance/generate-golden.mjs
```
