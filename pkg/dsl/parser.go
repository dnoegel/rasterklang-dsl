// Parser for DSL v2, backed by the ANTLR4-generated parser for
// grammar/RasterklangDsl.g4 (see its header comment for the design
// rationale and scripts/generate-parsers.sh for regeneration).
//
// ParseProgram(tokens) takes the []Token slice Tokenize produced, not
// source text: the two-step Tokenize-then-ParseProgram contract is public
// API. Internally, replayLexer (below) turns []Token back into something
// antlr.NewCommonTokenStream can drive, re-deriving each keyword's promoted
// token type (TEMPO/VOICE/... vs plain IDENT) from its text exactly as
// Tokenize did, so the round-trip is lossless for parsing purposes.
//
// ast_visitor.go turns ANTLR's parse tree into this package's AST types
// (ast.go); the ParserRuleContext tree is never handed to a consumer
// directly.
package dsl

import (
	"strconv"

	"github.com/antlr4-go/antlr/v4"

	gen "github.com/dnoegel/rasterklang-dsl/pkg/dsl/generated"
)

// ParseError is a line-numbered parser failure. Error() always starts
// with "line N: " -- see LexError's comment; Found is populated where
// available (kept for API compatibility) but callers should not depend on
// it being present.
type ParseError struct {
	Line    int
	Message string
	Found   string
}

func (e *ParseError) Error() string {
	return "line " + strconv.Itoa(e.Line) + ": " + e.Message
}

// replayLexer turns an already-lexed []Token (produced by Tokenize) back
// into something antlr.NewCommonTokenStream can drive. It embeds
// *antlr.BaseLexer purely to satisfy antlr.Lexer's unexported methods
// (only satisfiable via a type from the antlr package itself) and
// overrides NextToken() to replay Tokens instead of scanning characters.
type replayLexer struct {
	*antlr.BaseLexer
	tokens []Token
	pos    int
}

func newReplayLexer(tokens []Token) *replayLexer {
	base := antlr.NewBaseLexer(antlr.NewInputStream(""))
	return &replayLexer{BaseLexer: base, tokens: tokens}
}

// identTokenType re-derives which of the 8 statement keywords (or 'end')
// a TokenIdent actually is, purely from its text -- exactly how
// Tokenize's convertToken decided it in the first place (see
// lexer.go's convertToken).
func identTokenType(text string) int {
	switch text {
	case "tempo":
		return gen.RasterklangDslLexerTEMPO
	case "speed":
		return gen.RasterklangDslLexerSPEED
	case "voice":
		return gen.RasterklangDslLexerVOICE
	case "pattern":
		return gen.RasterklangDslLexerPATTERN
	case "table":
		return gen.RasterklangDslLexerTABLE
	case "instrument":
		return gen.RasterklangDslLexerINSTRUMENT
	case "filter":
		return gen.RasterklangDslLexerFILTER
	case "target":
		return gen.RasterklangDslLexerTARGET
	case "chip":
		return gen.RasterklangDslLexerCHIP
	case "end":
		return gen.RasterklangDslLexerEND
	default:
		return gen.RasterklangDslLexerIDENT
	}
}

func antlrTypeFor(t Token) int {
	switch t.Type {
	case TokenIdent:
		return identTokenType(t.Str)
	case TokenNumber:
		return gen.RasterklangDslLexerNUMBER
	case TokenHex:
		return gen.RasterklangDslLexerHEX
	case TokenString:
		return gen.RasterklangDslLexerDQSTRING
	case TokenSigil:
		return gen.RasterklangDslLexerSIGIL
	case TokenNewline:
		return gen.RasterklangDslLexerNEWLINE
	case TokenEOF:
		return antlr.TokenEOF
	case TokenPipe:
		return gen.RasterklangDslLexerPIPE
	case TokenColon:
		return gen.RasterklangDslLexerCOLON
	case TokenComma:
		return gen.RasterklangDslLexerCOMMA
	case TokenStar:
		return gen.RasterklangDslLexerSTAR
	case TokenSlash:
		return gen.RasterklangDslLexerSLASH
	case TokenLT:
		return gen.RasterklangDslLexerLT
	case TokenGE:
		return gen.RasterklangDslLexerGE
	case TokenShift:
		return gen.RasterklangDslLexerSHIFT
	case TokenLParen:
		return gen.RasterklangDslLexerLPAREN
	case TokenRParen:
		return gen.RasterklangDslLexerRPAREN
	case TokenEq:
		return gen.RasterklangDslLexerEQ
	case TokenPlus:
		return gen.RasterklangDslLexerPLUS
	case TokenBang:
		return gen.RasterklangDslLexerBANG
	case TokenDotDot:
		return gen.RasterklangDslLexerDOTDOT
	case TokenLBrace:
		return gen.RasterklangDslLexerLBRACE
	case TokenRBrace:
		return gen.RasterklangDslLexerRBRACE
	case TokenLBracket:
		return gen.RasterklangDslLexerLBRACKET
	case TokenRBracket:
		return gen.RasterklangDslLexerRBRACKET
	}
	return antlr.TokenInvalidType
}

func (r *replayLexer) NextToken() antlr.Token {
	factory := r.GetTokenFactory()
	if r.pos >= len(r.tokens) {
		return factory.Create(&antlr.TokenSourceCharStreamPair{}, antlr.TokenEOF, "<EOF>", antlr.TokenDefaultChannel, -1, -1, -1, 0)
	}
	t := r.tokens[r.pos]
	idx := r.pos
	r.pos++
	tt := antlrTypeFor(t)
	ct := factory.Create(&antlr.TokenSourceCharStreamPair{}, tt, t.Str, antlr.TokenDefaultChannel, -1, -1, t.Line, 0)
	ct.SetTokenIndex(idx)
	return ct
}

// ParseProgram drives the ANTLR-generated parser + AST-transform visitor
// over an already-tokenized []Token (from Tokenize), preserving
// ParseProgram's exact existing signature.
func ParseProgram(tokens []Token) (*Program, error) {
	lexer := newReplayLexer(tokens)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := gen.NewRasterklangDslParser(tokenStream)
	listener := &parserErrorListener{}
	parser.RemoveErrorListeners()
	parser.AddErrorListener(listener)
	tree := parser.Program()
	if listener.err != nil {
		return nil, listener.err
	}
	return buildProgram(tree, tokens)
}
