// Small ANTLR-runtime support layer shared by lexer.js and parser.js:
// the NEWLINE-suppressing lexer subclass, both error listeners, and the
// token-replay TokenSource that lets parseProgram() re-drive the
// generated parser from an already-tokenized array (see parser.js's own
// header for why that two-step shape is preserved).
//
// See grammar/RasterklangDsl.g4's header comment for the overall design
// rationale (soft keywords, and why the statement keywords and `end` are the
// exception promoted to dedicated lexer tokens).
import antlr4 from 'antlr4';
import RasterklangDslLexer from './generated/RasterklangDslLexer.js';

// Suppresses NEWLINE tokens that immediately follow start-of-input or
// another NEWLINE (i.e. blank/comment-only lines), giving the "one NEWLINE
// per real statement line, nothing for blank/comment lines" invariant. Also
// synthesizes a trailing NEWLINE at EOF when the source has no trailing '\n',
// so the last statement is always NEWLINE-terminated; ANTLR's NEWLINE rule
// can only match a real '\n' character.
export class FilteringLexer extends RasterklangDslLexer {
  constructor(input) {
    super(input);
    this._pendingReal = false;
    this._pendingEOF = null;
  }

  nextToken() {
    if (this._pendingEOF) {
      const eof = this._pendingEOF;
      this._pendingEOF = null;
      return eof;
    }
    for (;;) {
      const t = super.nextToken();
      if (t.type === RasterklangDslLexer.NEWLINE) {
        if (!this._pendingReal) continue;
        this._pendingReal = false;
        return t;
      }
      if (t.type === antlr4.Token.EOF) {
        if (this._pendingReal) {
          this._pendingReal = false;
          this._pendingEOF = t;
          const synthetic = new antlr4.CommonToken(undefined, RasterklangDslLexer.NEWLINE, antlr4.Token.DEFAULT_CHANNEL, -1, -1);
          synthetic.line = t.line;
          synthetic.text = '\n';
          return synthetic;
        }
        return t;
      }
      this._pendingReal = true;
      return t;
    }
  }
}

// LexerErrorListener produces specific wording for the 3 real lexer failure
// modes (unterminated "..." string, unterminated `...` string, and an
// unrecognized character) by inspecting the source character at the failure
// position. ANTLR's generic "token recognition error at: ..." doesn't
// distinguish these, and downstream tests assert on the wording, not just
// the line number.
export class LexerErrorListener extends antlr4.error.ErrorListener {
  constructor(sourceLines) {
    super();
    this.sourceLines = sourceLines;
  }

  syntaxError(recognizer, offendingSymbol, line, column, msg) {
    const lineText = this.sourceLines[line - 1] ?? '';
    const ch = lineText[column];
    if (ch === '"') {
      throw new Error(`line ${line}: unterminated string literal (no closing '"' before end of line)`);
    }
    if (ch === '`') {
      throw new Error(`line ${line}: unterminated backtick string literal (no closing '\`' before end of file)`);
    }
    if (ch !== undefined) {
      throw new Error(`line ${line}: unexpected character '${ch}'`);
    }
    throw new Error(`line ${line}: ${msg}`);
  }
}

// ParserErrorListener tailors the 3 "missing 'end'" cases (downstream tests
// assert on the wording) by inspecting the parser's current rule-context chain at the
// point of failure; everything else falls back to ANTLR's own message,
// which already names the missing literal specifically enough (e.g.
// "missing '=' at ...").
export class ParserErrorListener extends antlr4.error.ErrorListener {
  syntaxError(recognizer, offendingSymbol, line, column, msg) {
    // tokenize()'s EOF token sits one line "past" ANTLR's own EOF line, and
    // that's the token replayed here, so `line` arrives pre-adjusted.
    let ctx = recognizer._ctx;
    while (ctx) {
      const kind = ctx.constructor.name;
      // A pattern declaration's `=` became reachable only after an
      // optional declared length (`pattern bass 4 = "..."`, ADR 0001's
      // D1), so ANTLR's own message here degraded from "missing '='" to
      // "mismatched input 'c4' expecting {NUMBER, '='}". Spell out both
      // legal continuations instead of leaking a token-set dump.
      if (kind === 'PatternDeclContext') {
        const name = ctx.anyIdent() ? ctx.anyIdent().getText() : '?';
        throw new Error(
          `line ${line}: expected '=' after pattern name '${name}' (optionally preceded by the pattern's declared length in beats, e.g. \`pattern ${name} 4 = "..."\`)`,
        );
      }
      if (kind === 'InstrumentBlockContext') {
        const name = ctx.anyIdent() ? ctx.anyIdent().getText() : '?';
        throw new Error(`line ${line}: expected 'end' to close instrument block '${name}'`);
      }
      if (kind === 'VoiceBodyContext' || kind === 'VoiceStmtContext') {
        const voiceStmtCtx = kind === 'VoiceStmtContext' ? ctx : ctx.parentCtx;
        const index = voiceStmtCtx.NUMBER() ? voiceStmtCtx.NUMBER().getText() : '?';
        throw new Error(`line ${line}: expected 'end' to close voice ${index}'s block`);
      }
      if (kind === 'FilterBlockContext') {
        throw new Error(`line ${line}: expected 'end' to close the filter block`);
      }
      ctx = ctx.parentCtx;
    }
    throw new Error(`line ${line}: ${msg}`);
  }
}

// identTokenType(text) -- re-derives which of the statement keywords
// (or 'end') a plain IDENT Token actually is, purely from its text. This
// lets tokenize()'s output stay a flat array with no ANTLR token-type
// numbers while parseProgram()'s replay still recovers the exact lexer token
// type, since promotion is a pure function of the text.
export function identTokenType(text) {
  switch (text) {
    case 'tempo':
      return RasterklangDslLexer.TEMPO;
    case 'speed':
      return RasterklangDslLexer.SPEED;
    case 'voice':
      return RasterklangDslLexer.VOICE;
    case 'pattern':
      return RasterklangDslLexer.PATTERN;
    case 'table':
      return RasterklangDslLexer.TABLE;
    case 'instrument':
      return RasterklangDslLexer.INSTRUMENT;
    case 'filter':
      return RasterklangDslLexer.FILTER;
    case 'target':
      return RasterklangDslLexer.TARGET;
    case 'chip':
      return RasterklangDslLexer.CHIP;
    case 'end':
      return RasterklangDslLexer.END;
    default:
      return RasterklangDslLexer.IDENT;
  }
}

const OLD_TYPE_TO_ANTLR = {
  NUMBER: RasterklangDslLexer.NUMBER,
  HEX: RasterklangDslLexer.HEX,
  STRING: RasterklangDslLexer.DQSTRING,
  SIGIL: RasterklangDslLexer.SIGIL,
  NEWLINE: RasterklangDslLexer.NEWLINE,
  PIPE: RasterklangDslLexer.PIPE,
  DOTDOT: RasterklangDslLexer.DOTDOT,
  SHIFT: RasterklangDslLexer.SHIFT,
  GE: RasterklangDslLexer.GE,
  COLON: RasterklangDslLexer.COLON,
  COMMA: RasterklangDslLexer.COMMA,
  STAR: RasterklangDslLexer.STAR,
  SLASH: RasterklangDslLexer.SLASH,
  LT: RasterklangDslLexer.LT,
  LPAREN: RasterklangDslLexer.LPAREN,
  RPAREN: RasterklangDslLexer.RPAREN,
  EQ: RasterklangDslLexer.EQ,
  PLUS: RasterklangDslLexer.PLUS,
  BANG: RasterklangDslLexer.BANG,
  LBRACE: RasterklangDslLexer.LBRACE,
  RBRACE: RasterklangDslLexer.RBRACE,
  LBRACKET: RasterklangDslLexer.LBRACKET,
  RBRACKET: RasterklangDslLexer.RBRACKET,
};

export function antlrTypeFor(oldToken) {
  if (oldToken.type === 'EOF') return antlr4.Token.EOF;
  if (oldToken.type === 'IDENT') return identTokenType(oldToken.value);
  return OLD_TYPE_TO_ANTLR[oldToken.type];
}

// ReplayTokenSource turns an already-tokenized (old-shape) Token[] back
// into something antlr4.CommonTokenStream can drive -- see parser.js's
// header for why parseProgram(tokens) has to work this way.
export class ReplayTokenSource {
  constructor(tokens) {
    this.tokens = tokens;
    this.pos = 0;
  }

  nextToken() {
    const idx = Math.min(this.pos, this.tokens.length - 1);
    const orig = this.tokens[idx];
    this.pos += 1;
    const ct = new antlr4.CommonToken(undefined, antlrTypeFor(orig), antlr4.Token.DEFAULT_CHANNEL, -1, -1);
    ct.line = orig.line;
    ct.text = orig.value !== undefined ? String(orig.value) : '';
    ct.tokenIndex = idx;
    return ct;
  }
}
