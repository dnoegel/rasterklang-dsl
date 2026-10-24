# rasterklang-dsl

The rasterklang language: a text notation for writing music for the C64 SID
chip. This repository defines what the language *is*: its grammar, lexicon
and static/dynamic semantics, independent of any editor or player. Turning a
parsed song into SID register writes is the job of
[rasterklang-score](https://github.com/dnoegel/rasterklang-score).

> **Status: research project.** This is experimental code from an ongoing
> exploration of how to write SID music as text. The language changes in
> incompatible ways (see the changelog), there are no releases with
> compatibility guarantees, and parts of it exist only to support
> experiments. Use it to read, learn from or play with, not as a stable
> dependency.

## A taste

```
tempo 120 meter 4/4

instrument lead
  wave pulse
  adsr 0 9 6 4
  pw   2048
end

voice 0 lead
  notes "c4/4 e4/4 g4/4 <c5 b4>/4"
end
```

`docs/dsl-grammar-spec.md` describes the full language.

## What's here

- **`src/dsl/`**: the JS reference implementation: lexer, parser, semantic
  analysis (`analyze.js`, `miniLength.js`), interpreter, keyword/field-doc
  tables and the language version constant. This is the canonical
  implementation; the Go port is validated against it.
- **`pkg/dsl/`**: a Go port of the lexer and parser (not the semantic and
  interpreter layer), producing a Go AST, plus the mini-notation length
  arithmetic (`minilength.go`) that tools generating rasterklang source
  need. Validated against the JS implementation by the shared fixtures in
  `test/conformance/` (see that directory's README).
- **`pkg/dslgen/`**: Go helpers for *emitting* rasterklang source text
  (repeat folding, alternation folding, numeric table formatting).
- **`docs/dsl-grammar-spec.md`**: the narrative language spec.
- **`docs/dsl-changelog.md`**: what changed between language versions.
- **`docs/adr/`**: decision records for changes big enough to need one.
- **`docs/dsl-spec.generated.json`**: machine-readable spec, generated from
  `src/dsl/{keywords,fieldDocs,version}.js` by `npm run generate:dsl-spec`
  (never edited by hand).
- **`scripts/dsl-lint.mjs`** (`npm run lint:dsl -- FILE.dsl`): reports a
  file's diagnostics without compiling it.

## Grammar

[`grammar/RasterklangDsl.g4`](grammar/RasterklangDsl.g4) is the lexical and
syntactic grammar in [ANTLR4](https://www.antlr.org/) form. Both the JS
(`src/dsl/generated/`) and the Go (`pkg/dsl/generated/`) parsers are
generated from it. The language has no reserved words (`wave`, `tempo` and
so on are ordinary identifiers, dispatched by context); the comment block at
the top of the `.g4` explains how that is modelled.

After editing the grammar: `bash scripts/generate-parsers.sh` regenerates
both parsers (needs `java` 17+ and an ANTLR 4.13.2 jar, see the script),
`bash scripts/check-parsers-fresh.sh` verifies the committed output matches
the grammar, and both test suites (`npm test`, `go test ./...`) should still
pass. A new construct also needs a case in the hand-written AST visitors
(`src/dsl/astVisitor.js`, `pkg/dsl/ast_visitor.go`).

## Installing

There is no npm release. Install from GitHub together with the engine
package (the two depend on each other):

```sh
npm install github:dnoegel/rasterklang-dsl github:dnoegel/rasterklang-score
```

```js
import { compileDslV2Source, checkSource } from 'rasterklang-dsl/src/dsl/index.js';
```

The Go packages are available as `github.com/dnoegel/rasterklang-dsl/pkg/dsl`
and `github.com/dnoegel/rasterklang-dsl/pkg/dslgen`.

## Development

The JS tests expect a checkout of `rasterklang-score` next to this one:

```sh
git clone https://github.com/dnoegel/rasterklang-dsl.git
git clone https://github.com/dnoegel/rasterklang-score.git
cd rasterklang-dsl && npm install && npm test && go test ./...
```

## License

MIT, see [LICENSE](LICENSE).
