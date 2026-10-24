// compileDslV2Source(sourceText) -> Song: lexer -> parser -> interpret.
//  - success: returns a plain Song builder object, the same shape
//     rasterklang-score's song() produces (voice()/filter()/stack() methods,
//     .voices/.filterChain fields). rasterklang-score's compile.js turns that
//     into register writes.
//  - failure: throws a plain Error whose message is safe to show directly in
//     an editor's error banner: a 1-based source line number plus a short
//     description. Every stage already throws that shape; the wrapper below
//     only guarantees it for an unexpected non-Error throw.
import { tokenize } from './lexer.js';
import { parseProgram } from './parser.js';
import { interpret } from './interpret.js';
import { checkProgram } from './analyze.js';

export { checkProgram };

// options.onDiagnostic, if given, receives each non-fatal diagnostic the
// semantic pass produced (see interpret()'s doc comment). They are on the
// returned Song as `.dslDiagnostics` either way.
export function compileDslV2Source(sourceText, options = {}) {
  try {
    const tokens = tokenize(sourceText);
    const ast = parseProgram(tokens);
    return interpret(ast, options);
  } catch (err) {
    if (err instanceof Error) throw err;
    throw new Error(String(err));
  }
}

// checkSource(sourceText) -> diagnostics for one file, without compiling
// it and so without needing the pattern engine at all - the entry point
// scripts/dsl-lint.mjs (the ADR 0001 decision 10 migration checker) and any
// editor integration should use. A lex/parse failure comes back as a single
// error-severity diagnostic rather than a throw, so a linter can report it
// the same way it reports everything else.
export function checkSource(sourceText) {
  try {
    return checkProgram(parseProgram(tokenize(sourceText)));
  } catch (err) {
    const message = err instanceof Error ? err.message : String(err);
    const m = /^line (\d+): (.*)$/s.exec(message);
    return [{ severity: 'error', line: m ? Number(m[1]) : 1, message: m ? m[2] : message }];
  }
}
