// Parser for DSL v2, backed by the ANTLR4-generated parser for
// grammar/RasterklangDsl.g4 (see its header comment for the design rationale
// and scripts/generate-parsers.sh for regeneration).
//
// parseProgram(tokens) takes the Token[] array lexer.js's tokenize()
// produced, not source text: the two-step tokenize()-then-parseProgram()
// contract is public API. Internally, antlrSupport.js's ReplayTokenSource
// turns Token[] back into an ANTLR token source, re-deriving each keyword's
// promoted token type (TEMPO/VOICE/... vs plain IDENT) from its text exactly
// as tokenize() did, so the round-trip is lossless for parsing purposes.
//
// astVisitor.js turns ANTLR's parse tree into this package's AST node shapes
// (plain objects whose `type` fields mirror docs/dsl-grammar-spec.md's
// production names 1:1). The ParserRuleContext tree is never handed to a
// consumer directly.
import antlr4 from 'antlr4';
import RasterklangDslParser from './generated/RasterklangDslParser.js';
import { ReplayTokenSource, ParserErrorListener } from './antlrSupport.js';
import {
  buildProgram,
  STATEMENT_KEYWORDS,
  INSTRUMENT_FIELDS,
  VOICE_FIELDS,
  TRANSFORMS,
  FILTER_FIELDS,
  GATE_MODES,
  VERB_NAMES,
  VIDEO_STANDARDS,
  CHIP_MODELS,
} from './astVisitor.js';
import { LOOP_TABLE_NAMES } from './instrument.js';

// Re-exported, not retyped: these keyword lists live in astVisitor.js,
// right next to the switch statements that actually validate them, so a
// diff adding a case there can't silently leave a second, drifting copy
// behind.
export {
  LOOP_TABLE_NAMES,
  STATEMENT_KEYWORDS,
  INSTRUMENT_FIELDS,
  VOICE_FIELDS,
  TRANSFORMS,
  FILTER_FIELDS,
  GATE_MODES,
  VERB_NAMES,
  VIDEO_STANDARDS,
  CHIP_MODELS,
};

// parseProgram(tokens) -> the Program AST node.
export function parseProgram(tokens) {
  const source = new ReplayTokenSource(tokens);
  const tokenStream = new antlr4.CommonTokenStream(source);
  const parser = new RasterklangDslParser(tokenStream);
  parser.removeErrorListeners();
  parser.addErrorListener(new ParserErrorListener());
  const tree = parser.program();
  return buildProgram(tree, tokens);
}
