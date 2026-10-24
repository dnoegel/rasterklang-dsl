package dsl

import (
	"github.com/antlr4-go/antlr/v4"
)

// lexerErrorListener reproduces lexer.js's specific wording for the 3 real
// lexer failure modes (unterminated "..." string,
// unterminated `...` string, and a genuinely unrecognized character) by
// inspecting the source character at the failure position - ANTLR's own
// generic "token recognition error at: ..." doesn't distinguish these.
// Only the "line N: " prefix is part of the actual cross-language
// conformance contract (see test/conformance/README.md), but matching
// the specific wording keeps this port's errors equally legible.
type lexerErrorListener struct {
	*antlr.DefaultErrorListener
	sourceLines []string
	err         error
}

func newLexerErrorListener(sourceLines []string) *lexerErrorListener {
	return &lexerErrorListener{sourceLines: sourceLines}
}

func (l *lexerErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	if l.err != nil {
		return
	}
	var lineText string
	if line-1 >= 0 && line-1 < len(l.sourceLines) {
		lineText = l.sourceLines[line-1]
	}
	var ch byte
	if column >= 0 && column < len(lineText) {
		ch = lineText[column]
	}
	switch ch {
	case '"':
		l.err = &LexError{Line: line, Message: `unterminated string literal (no closing '"' before end of line)`}
	case '`':
		l.err = &LexError{Line: line, Message: "unterminated backtick string literal (no closing '`' before end of file)"}
	default:
		if ch != 0 {
			l.err = &LexError{Line: line, Message: "unexpected character '" + string(ch) + "'"}
		} else {
			l.err = &LexError{Line: line, Message: msg}
		}
	}
}

// parserErrorListener records the first parser syntax error. Unlike the
// JS port's ParserErrorListener, this doesn't tailor the "missing 'end'"
// wording to match parser.go's original messages exactly -- nothing in
// this repo's Go test suite asserts on error MESSAGE text (only the
// "line N: " prefix, via test/conformance's line-number-only check for
// error fixtures), so ANTLR's own message (which already names the
// missing literal, e.g. "missing '=' at ...", or "extraneous input
// '<EOF>' expecting {...}" for a missing 'end') is left as-is.
type parserErrorListener struct {
	*antlr.DefaultErrorListener
	err error
}

func (l *parserErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	if l.err != nil {
		return
	}
	// Tokenize's EOF token already carries lexer.go's "one past the last
	// line" convention, and that's the token replayed for parsing, so
	// `line` arrives here already adjusted -- no further compensation.
	l.err = &ParseError{Line: line, Message: msg}
}
