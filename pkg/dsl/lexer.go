// Package dsl is the Go implementation of the DSL v2 grammar.
// Tokenize/ParseProgram are backed by the ANTLR4-generated lexer/parser in
// pkg/dsl/generated/ (from grammar/RasterklangDsl.g4; see its header
// comment) plus the hand-written AST-transform visitor in ast_visitor.go.
// This file produces the public []Token shape and the
// Tokenize(string) ([]Token, error) entry point that ParseProgram and other
// consumers use.
package dsl

import (
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"

	gen "github.com/dnoegel/rasterklang-dsl/pkg/dsl/generated"
)

// TokenType identifies what kind of lexeme a Token is. Punctuation types
// carry no payload (see Token.Value) - the type alone is the payload,
// same as lexer.js's token shape.
type TokenType int

const (
	TokenIdent TokenType = iota
	TokenNumber
	TokenHex
	TokenString
	TokenSigil
	TokenNewline
	TokenEOF
	TokenPipe   // |>
	TokenColon  // :
	TokenComma  // ,
	TokenStar   // *
	TokenSlash  // /
	TokenLT     // <
	TokenGE     // >=
	TokenShift  // >>
	TokenLParen // (
	TokenRParen // )
	TokenEq     // =
	TokenPlus   // +
	TokenBang   // !
	TokenDotDot // ..
	TokenLBrace   // {
	TokenRBrace   // }
	TokenLBracket // [ (ADR 0003 D4 arrangements)
	TokenRBracket // ]
)

func (t TokenType) String() string {
	switch t {
	case TokenIdent:
		return "IDENT"
	case TokenNumber:
		return "NUMBER"
	case TokenHex:
		return "HEX"
	case TokenString:
		return "STRING"
	case TokenSigil:
		return "SIGIL"
	case TokenNewline:
		return "NEWLINE"
	case TokenEOF:
		return "EOF"
	case TokenPipe:
		return "PIPE"
	case TokenColon:
		return "COLON"
	case TokenComma:
		return "COMMA"
	case TokenStar:
		return "STAR"
	case TokenSlash:
		return "SLASH"
	case TokenLT:
		return "LT"
	case TokenGE:
		return "GE"
	case TokenShift:
		return "SHIFT"
	case TokenLParen:
		return "LPAREN"
	case TokenRParen:
		return "RPAREN"
	case TokenEq:
		return "EQ"
	case TokenPlus:
		return "PLUS"
	case TokenBang:
		return "BANG"
	case TokenDotDot:
		return "DOTDOT"
	case TokenLBrace:
		return "LBRACE"
	case TokenRBrace:
		return "RBRACE"
	case TokenLBracket:
		return "LBRACKET"
	case TokenRBracket:
		return "RBRACKET"
	default:
		return "UNKNOWN"
	}
}

// Token mirrors lexer.js's { type, value?, line, docStart?, docEnd? }
// shape. Only one of Str/Num/HexVal is meaningful, chosen by Type
// (IDENT/SIGIL/STRING -> Str, NUMBER -> Num, HEX -> HexVal); DocStart/
// DocEnd are only meaningful for STRING.
type Token struct {
	Type     TokenType
	Line     int
	Str      string
	Num      float64
	HexVal   int
	DocStart int
	DocEnd   int
}

// LexError is a line-numbered lexer failure. Error() always starts with
// "line N: " - that prefix, not the rest of the text, is the actual
// cross-language conformance contract (see test/conformance's README).
type LexError struct {
	Line    int
	Message string
}

func (e *LexError) Error() string {
	return "line " + strconv.Itoa(e.Line) + ": " + e.Message
}

// Tokenize(sourceText) -> Token slice, always ending with a single
// {Type: TokenEOF} sentinel, so the parser never needs to special-case
// "ran off the end of the slice" separately from "hit a real EOF token."
func Tokenize(sourceText string) ([]Token, error) {
	input := antlr.NewInputStream(sourceText)
	lexer := gen.NewRasterklangDslLexer(input)
	sourceLines := strings.Split(sourceText, "\n")
	listener := newLexerErrorListener(sourceLines)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(listener)

	var tokens []Token
	pendingReal := false
	for {
		t := lexer.NextToken()
		if listener.err != nil {
			return nil, listener.err
		}
		tt := t.GetTokenType()
		if tt == gen.RasterklangDslLexerNEWLINE {
			if !pendingReal {
				continue // blank/comment-only line: drop it, matching lexer.js
			}
			pendingReal = false
			tokens = append(tokens, Token{Type: TokenNewline, Line: t.GetLine()})
			continue
		}
		if tt == antlr.TokenEOF {
			// lexer.js always emits a NEWLINE after the last real-content
			// line even with no actual trailing '\n' character in the
			// source (it iterates strings.Split(sourceText, "\n"), which
			// always yields a final "line" for whatever follows the last
			// newline, including nothing) -- synthesize that same virtual
			// NEWLINE here, and its own EOF line is always "one past" the
			// last physical line it visited.
			if pendingReal {
				tokens = append(tokens, Token{Type: TokenNewline, Line: t.GetLine()})
			}
			tokens = append(tokens, Token{Type: TokenEOF, Line: t.GetLine() + 1})
			break
		}
		pendingReal = true
		tokens = append(tokens, convertToken(t))
	}
	return tokens, nil
}

func convertToken(t antlr.Token) Token {
	line := t.GetLine()
	text := t.GetText()
	switch t.GetTokenType() {
	case gen.RasterklangDslLexerTEMPO, gen.RasterklangDslLexerSPEED, gen.RasterklangDslLexerVOICE, gen.RasterklangDslLexerPATTERN,
		gen.RasterklangDslLexerTABLE, gen.RasterklangDslLexerINSTRUMENT, gen.RasterklangDslLexerFILTER,
		gen.RasterklangDslLexerTARGET, gen.RasterklangDslLexerCHIP, gen.RasterklangDslLexerEND, gen.RasterklangDslLexerIDENT:
		return Token{Type: TokenIdent, Line: line, Str: text}
	case gen.RasterklangDslLexerNUMBER:
		// ADR 0003 D11: a trailing '%' is this same NUMBER token with '%'
		// folded into its lexeme (see the grammar's own comment) - divide by
		// 100 here, mirroring lexer.js's convertToken exactly.
		raw := text
		isPercent := strings.HasSuffix(raw, "%")
		if isPercent {
			raw = raw[:len(raw)-1]
		}
		v, _ := strconv.ParseFloat(raw, 64)
		if isPercent {
			v /= 100
		}
		return Token{Type: TokenNumber, Line: line, Num: v}
	case gen.RasterklangDslLexerHEX:
		v, _ := strconv.ParseInt(text[2:], 16, 64)
		return Token{Type: TokenHex, Line: line, HexVal: int(v)}
	case gen.RasterklangDslLexerDQSTRING:
		val, dStart, dEnd := unquoteString(t, '"')
		return Token{Type: TokenString, Line: line, Str: val, DocStart: dStart, DocEnd: dEnd}
	case gen.RasterklangDslLexerBTSTRING:
		val, dStart, dEnd := unquoteString(t, '`')
		return Token{Type: TokenString, Line: line, Str: val, DocStart: dStart, DocEnd: dEnd}
	case gen.RasterklangDslLexerSIGIL:
		return Token{Type: TokenSigil, Line: line, Str: text[1:]}
	case gen.RasterklangDslLexerPIPE:
		return Token{Type: TokenPipe, Line: line}
	case gen.RasterklangDslLexerDOTDOT:
		return Token{Type: TokenDotDot, Line: line}
	case gen.RasterklangDslLexerSHIFT:
		return Token{Type: TokenShift, Line: line}
	case gen.RasterklangDslLexerGE:
		return Token{Type: TokenGE, Line: line}
	case gen.RasterklangDslLexerCOLON:
		return Token{Type: TokenColon, Line: line}
	case gen.RasterklangDslLexerCOMMA:
		return Token{Type: TokenComma, Line: line}
	case gen.RasterklangDslLexerSTAR:
		return Token{Type: TokenStar, Line: line}
	case gen.RasterklangDslLexerSLASH:
		return Token{Type: TokenSlash, Line: line}
	case gen.RasterklangDslLexerLT:
		return Token{Type: TokenLT, Line: line}
	case gen.RasterklangDslLexerLPAREN:
		return Token{Type: TokenLParen, Line: line}
	case gen.RasterklangDslLexerRPAREN:
		return Token{Type: TokenRParen, Line: line}
	case gen.RasterklangDslLexerEQ:
		return Token{Type: TokenEq, Line: line}
	case gen.RasterklangDslLexerPLUS:
		return Token{Type: TokenPlus, Line: line}
	case gen.RasterklangDslLexerBANG:
		return Token{Type: TokenBang, Line: line}
	case gen.RasterklangDslLexerLBRACE:
		return Token{Type: TokenLBrace, Line: line}
	case gen.RasterklangDslLexerRBRACE:
		return Token{Type: TokenRBrace, Line: line}
	case gen.RasterklangDslLexerLBRACKET:
		return Token{Type: TokenLBracket, Line: line}
	case gen.RasterklangDslLexerRBRACKET:
		return Token{Type: TokenRBracket, Line: line}
	}
	// Unreachable for well-formed grammar output; fall back to IDENT
	// rather than panicking on a future grammar change that adds a token
	// this switch hasn't been taught about yet.
	return Token{Type: TokenIdent, Line: line, Str: text}
}

func unquoteString(t antlr.Token, quote byte) (value string, docStart, docEnd int) {
	text := t.GetText()
	inner := text[1 : len(text)-1]
	out := make([]byte, 0, len(inner))
	for i := 0; i < len(inner); i++ {
		if inner[i] == '\\' && i+1 < len(inner) && (inner[i+1] == quote || inner[i+1] == '\\') {
			out = append(out, inner[i+1])
			i++
			continue
		}
		out = append(out, inner[i])
	}
	start := t.GetStart() + 1
	stop := t.GetStop()
	return string(out), start, stop
}
