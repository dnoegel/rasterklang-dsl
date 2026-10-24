// Tokenizer for DSL v2 (docs/dsl-grammar-spec.md's "Lexical rules"), backed
// by the ANTLR4-generated lexer for grammar/RasterklangDsl.g4. It produces a
// plain Token[] -- { type, value?, line, docStart?, docEnd? }, one NEWLINE
// token per real-content line (never for blank/comment-only lines), always
// ending in a single EOF sentinel -- which parser.js, editor tokenizers and
// the conformance suite consume.
//
// Both string delimiters (`"..."` single-line, `` `...` `` multi-line)
// collapse to one token type (STRING); only STRING carries docStart/docEnd.
// See antlrSupport.js for the NEWLINE-suppression/synthesis lexer subclass
// and the error listener that produces this package's error wording.
import antlr4 from 'antlr4';
import RasterklangDslLexer from './generated/RasterklangDslLexer.js';
import { FilteringLexer, LexerErrorListener } from './antlrSupport.js';

const L = RasterklangDslLexer;

// TYPE_NAMES maps every ANTLR lexer token-type constant to a plain string
// type name. The keyword tokens (TEMPO/SPEED/VOICE/...) all collapse to
// 'IDENT': the promotion only matters to the parser's token-replay step (see
// antlrSupport.js's identTokenType), so tokenize()'s output stays
// keyword-agnostic.
const TYPE_NAMES = {
  [L.TEMPO]: 'IDENT',
  [L.SPEED]: 'IDENT',
  [L.VOICE]: 'IDENT',
  [L.PATTERN]: 'IDENT',
  [L.TABLE]: 'IDENT',
  [L.INSTRUMENT]: 'IDENT',
  [L.FILTER]: 'IDENT',
  [L.TARGET]: 'IDENT',
  [L.CHIP]: 'IDENT',
  [L.END]: 'IDENT',
  [L.IDENT]: 'IDENT',
  [L.NUMBER]: 'NUMBER',
  [L.HEX]: 'HEX',
  [L.DQSTRING]: 'STRING',
  [L.BTSTRING]: 'STRING',
  [L.SIGIL]: 'SIGIL',
  [L.PIPE]: 'PIPE',
  [L.DOTDOT]: 'DOTDOT',
  [L.SHIFT]: 'SHIFT',
  [L.GE]: 'GE',
  [L.COLON]: 'COLON',
  [L.COMMA]: 'COMMA',
  [L.STAR]: 'STAR',
  [L.SLASH]: 'SLASH',
  [L.LT]: 'LT',
  [L.LPAREN]: 'LPAREN',
  [L.RPAREN]: 'RPAREN',
  [L.EQ]: 'EQ',
  [L.PLUS]: 'PLUS',
  [L.BANG]: 'BANG',
  [L.LBRACE]: 'LBRACE',
  [L.RBRACE]: 'RBRACE',
  [L.LBRACKET]: 'LBRACKET',
  [L.RBRACKET]: 'RBRACKET',
};

function unescapeStr(s, quoteChar) {
  let out = '';
  for (let i = 0; i < s.length; i += 1) {
    if (s[i] === '\\' && (s[i + 1] === quoteChar || s[i + 1] === '\\')) {
      out += s[i + 1];
      i += 1;
    } else {
      out += s[i];
    }
  }
  return out;
}

function convertToken(t) {
  const type = TYPE_NAMES[t.type];
  const line = t.line;
  switch (type) {
    case 'IDENT':
      return { type, value: t.text, line };
    case 'NUMBER': {
      // ADR 0003 D11: a trailing '%' (e.g. 'gate 80%') is this same NUMBER
      // token with '%' folded into its lexeme (see the grammar's own
      // comment) - divide by 100 here, the one place every NUMBER's text
      // becomes a value, so '80%' and '0.8' are indistinguishable from
      // here on.
      const isPercent = t.text.endsWith('%');
      const raw = isPercent ? t.text.slice(0, -1) : t.text;
      return { type, value: isPercent ? Number(raw) / 100 : Number(raw), line };
    }
    case 'HEX':
      return { type, value: parseInt(t.text.slice(2), 16), line };
    case 'SIGIL':
      return { type, value: t.text.slice(1), line };
    case 'STRING': {
      const isBacktick = t.type === L.BTSTRING;
      const inner = t.text.slice(1, -1);
      const value = unescapeStr(inner, isBacktick ? '`' : '"');
      return { type, value, line, docStart: t.start + 1, docEnd: t.stop };
    }
    default:
      return { type, line };
  }
}

// tokenize(sourceText) -> Token[], always ending with a single {type:'EOF'}
// sentinel.
export function tokenize(sourceText) {
  const chars = new antlr4.InputStream(sourceText);
  const lexer = new FilteringLexer(chars);
  lexer.removeErrorListeners();
  lexer.addErrorListener(new LexerErrorListener(sourceText.split('\n')));

  const tokens = [];
  for (;;) {
    const t = lexer.nextToken();
    if (t.type === antlr4.Token.EOF) {
      // lexer.js's own EOF line is always "one past" the last physical
      // line it visited (it iterates a trailing blank pseudo-line before
      // pushing EOF) -- ANTLR's raw EOF line needs the same +1 here.
      tokens.push({ type: 'EOF', line: t.line + 1 });
      break;
    }
    if (t.type === RasterklangDslLexer.NEWLINE) {
      tokens.push({ type: 'NEWLINE', line: t.line });
      continue;
    }
    tokens.push(convertToken(t));
  }
  return tokens;
}
