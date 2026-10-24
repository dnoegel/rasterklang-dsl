// Code generated from RasterklangDsl.g4 by ANTLR 4.13.2. DO NOT EDIT.

package generated // RasterklangDsl
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type RasterklangDslParser struct {
	*antlr.BaseParser
}

var RasterklangDslParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rasterklangdslParserInit() {
	staticData := &RasterklangDslParserStaticData
	staticData.LiteralNames = []string{
		"", "'tempo'", "'speed'", "'voice'", "'pattern'", "'table'", "'instrument'",
		"'filter'", "'target'", "'chip'", "'end'", "", "", "", "", "", "", "'|>'",
		"'..'", "'>>'", "'>='", "':'", "','", "'*'", "'/'", "'<'", "'('", "')'",
		"'{'", "'}'", "'['", "']'", "'='", "'+'", "'!'",
	}
	staticData.SymbolicNames = []string{
		"", "TEMPO", "SPEED", "VOICE", "PATTERN", "TABLE", "INSTRUMENT", "FILTER",
		"TARGET", "CHIP", "END", "HEX", "NUMBER", "DQSTRING", "BTSTRING", "SIGIL",
		"IDENT", "PIPE", "DOTDOT", "SHIFT", "GE", "COLON", "COMMA", "STAR",
		"SLASH", "LT", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACKET", "RBRACKET",
		"EQ", "PLUS", "BANG", "NEWLINE", "WS", "COMMENT",
	}
	staticData.RuleNames = []string{
		"program", "statement", "unknownStmt", "unknownTail", "anyIdent", "str",
		"tempoStmt", "tempoValue", "tempoSetting", "targetStmt", "chipStmt",
		"targetSetting", "patternDecl", "patternEach", "arrangement", "arrangementElement",
		"tableDecl", "instrumentBlock", "instrumentField", "tableVerbLine",
		"tableRow", "tableToken", "absNote", "rampGenerator", "voiceStmt", "voiceBody",
		"voiceShortTail", "inlineAttr", "valueList", "voiceField", "voiceFieldValue",
		"pipelineStep", "transform", "transformArg", "transformRef", "filterBlock",
		"filterField", "filterFieldValue", "switchCase", "switchCond", "signalRangeOpt",
		"signalSpeedMod", "numberList",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 37, 469, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 1, 0, 1, 0, 1, 0, 5, 0, 90, 8, 0, 10, 0, 12, 0, 93, 9, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 106,
		8, 1, 1, 2, 1, 2, 5, 2, 110, 8, 2, 10, 2, 12, 2, 113, 9, 2, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 5, 6, 124, 8, 6, 10, 6, 12, 6,
		127, 9, 6, 1, 7, 1, 7, 3, 7, 131, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 137,
		8, 8, 1, 9, 1, 9, 1, 9, 5, 9, 142, 8, 9, 10, 9, 12, 9, 145, 9, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 3, 11, 155, 8, 11, 1, 12,
		1, 12, 1, 12, 3, 12, 160, 8, 12, 1, 12, 3, 12, 163, 8, 12, 1, 12, 1, 12,
		1, 12, 5, 12, 168, 8, 12, 10, 12, 12, 12, 171, 9, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 3, 12, 178, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 14, 4, 14, 186, 8, 14, 11, 14, 12, 14, 187, 1, 15, 1, 15, 1, 15, 4,
		15, 193, 8, 15, 11, 15, 12, 15, 194, 1, 15, 1, 15, 3, 15, 199, 8, 15, 1,
		15, 1, 15, 3, 15, 203, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17,
		1, 17, 1, 17, 1, 17, 5, 17, 214, 8, 17, 10, 17, 12, 17, 217, 9, 17, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 4, 18, 228,
		8, 18, 11, 18, 12, 18, 229, 1, 18, 1, 18, 1, 18, 3, 18, 235, 8, 18, 1,
		19, 1, 19, 3, 19, 239, 8, 19, 1, 19, 1, 19, 1, 20, 4, 20, 244, 8, 20, 11,
		20, 12, 20, 245, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 254,
		8, 21, 10, 21, 12, 21, 257, 9, 21, 1, 21, 1, 21, 3, 21, 261, 8, 21, 1,
		21, 1, 21, 1, 21, 3, 21, 266, 8, 21, 3, 21, 268, 8, 21, 1, 22, 1, 22, 1,
		22, 1, 22, 3, 22, 274, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 3, 23, 283, 8, 23, 1, 23, 1, 23, 1, 23, 3, 23, 288, 8, 23, 1, 23,
		3, 23, 291, 8, 23, 1, 24, 1, 24, 1, 24, 3, 24, 296, 8, 24, 1, 24, 1, 24,
		1, 25, 1, 25, 3, 25, 302, 8, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 5, 25, 310, 8, 25, 10, 25, 12, 25, 313, 9, 25, 1, 25, 3, 25, 316, 8,
		25, 1, 26, 3, 26, 319, 8, 26, 1, 26, 5, 26, 322, 8, 26, 10, 26, 12, 26,
		325, 9, 26, 1, 26, 3, 26, 328, 8, 26, 1, 26, 5, 26, 331, 8, 26, 10, 26,
		12, 26, 334, 9, 26, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 340, 8, 27, 1, 28,
		1, 28, 1, 28, 5, 28, 345, 8, 28, 10, 28, 12, 28, 348, 9, 28, 1, 29, 1,
		29, 1, 29, 1, 29, 3, 29, 354, 8, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30,
		1, 30, 3, 30, 362, 8, 30, 1, 30, 1, 30, 1, 30, 3, 30, 367, 8, 30, 3, 30,
		369, 8, 30, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 3, 32, 376, 8, 32, 1, 32,
		3, 32, 379, 8, 32, 1, 33, 1, 33, 3, 33, 383, 8, 33, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 3, 34, 390, 8, 34, 1, 35, 1, 35, 1, 35, 5, 35, 395, 8, 35,
		10, 35, 12, 35, 398, 9, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 5, 37, 412, 8, 37, 10, 37, 12, 37,
		415, 9, 37, 1, 37, 1, 37, 1, 37, 5, 37, 420, 8, 37, 10, 37, 12, 37, 423,
		9, 37, 1, 37, 3, 37, 426, 8, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 5,
		38, 433, 8, 38, 10, 38, 12, 38, 436, 9, 38, 1, 38, 3, 38, 439, 8, 38, 1,
		39, 1, 39, 1, 39, 1, 39, 1, 39, 3, 39, 446, 8, 39, 1, 40, 1, 40, 1, 40,
		3, 40, 451, 8, 40, 1, 41, 1, 41, 1, 41, 1, 41, 3, 41, 457, 8, 41, 1, 42,
		1, 42, 3, 42, 461, 8, 42, 1, 42, 5, 42, 464, 8, 42, 10, 42, 12, 42, 467,
		9, 42, 1, 42, 0, 0, 43, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
		62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 0, 5, 1, 0, 35, 35, 2,
		0, 1, 9, 16, 16, 1, 0, 13, 14, 1, 0, 1, 2, 1, 0, 23, 24, 498, 0, 91, 1,
		0, 0, 0, 2, 105, 1, 0, 0, 0, 4, 107, 1, 0, 0, 0, 6, 114, 1, 0, 0, 0, 8,
		116, 1, 0, 0, 0, 10, 118, 1, 0, 0, 0, 12, 120, 1, 0, 0, 0, 14, 130, 1,
		0, 0, 0, 16, 132, 1, 0, 0, 0, 18, 138, 1, 0, 0, 0, 20, 148, 1, 0, 0, 0,
		22, 151, 1, 0, 0, 0, 24, 177, 1, 0, 0, 0, 26, 179, 1, 0, 0, 0, 28, 185,
		1, 0, 0, 0, 30, 198, 1, 0, 0, 0, 32, 204, 1, 0, 0, 0, 34, 209, 1, 0, 0,
		0, 36, 234, 1, 0, 0, 0, 38, 236, 1, 0, 0, 0, 40, 243, 1, 0, 0, 0, 42, 267,
		1, 0, 0, 0, 44, 269, 1, 0, 0, 0, 46, 275, 1, 0, 0, 0, 48, 292, 1, 0, 0,
		0, 50, 315, 1, 0, 0, 0, 52, 323, 1, 0, 0, 0, 54, 335, 1, 0, 0, 0, 56, 341,
		1, 0, 0, 0, 58, 353, 1, 0, 0, 0, 60, 368, 1, 0, 0, 0, 62, 370, 1, 0, 0,
		0, 64, 373, 1, 0, 0, 0, 66, 382, 1, 0, 0, 0, 68, 389, 1, 0, 0, 0, 70, 391,
		1, 0, 0, 0, 72, 401, 1, 0, 0, 0, 74, 425, 1, 0, 0, 0, 76, 427, 1, 0, 0,
		0, 78, 445, 1, 0, 0, 0, 80, 450, 1, 0, 0, 0, 82, 456, 1, 0, 0, 0, 84, 458,
		1, 0, 0, 0, 86, 87, 3, 2, 1, 0, 87, 88, 5, 35, 0, 0, 88, 90, 1, 0, 0, 0,
		89, 86, 1, 0, 0, 0, 90, 93, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92, 1,
		0, 0, 0, 92, 94, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 94, 95, 5, 0, 0, 1, 95,
		1, 1, 0, 0, 0, 96, 106, 3, 12, 6, 0, 97, 106, 3, 18, 9, 0, 98, 106, 3,
		20, 10, 0, 99, 106, 3, 48, 24, 0, 100, 106, 3, 24, 12, 0, 101, 106, 3,
		32, 16, 0, 102, 106, 3, 34, 17, 0, 103, 106, 3, 70, 35, 0, 104, 106, 3,
		4, 2, 0, 105, 96, 1, 0, 0, 0, 105, 97, 1, 0, 0, 0, 105, 98, 1, 0, 0, 0,
		105, 99, 1, 0, 0, 0, 105, 100, 1, 0, 0, 0, 105, 101, 1, 0, 0, 0, 105, 102,
		1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 105, 104, 1, 0, 0, 0, 106, 3, 1, 0, 0,
		0, 107, 111, 5, 16, 0, 0, 108, 110, 3, 6, 3, 0, 109, 108, 1, 0, 0, 0, 110,
		113, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 5, 1,
		0, 0, 0, 113, 111, 1, 0, 0, 0, 114, 115, 8, 0, 0, 0, 115, 7, 1, 0, 0, 0,
		116, 117, 7, 1, 0, 0, 117, 9, 1, 0, 0, 0, 118, 119, 7, 2, 0, 0, 119, 11,
		1, 0, 0, 0, 120, 121, 7, 3, 0, 0, 121, 125, 3, 14, 7, 0, 122, 124, 3, 16,
		8, 0, 123, 122, 1, 0, 0, 0, 124, 127, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0,
		125, 126, 1, 0, 0, 0, 126, 13, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 128, 131,
		5, 12, 0, 0, 129, 131, 3, 10, 5, 0, 130, 128, 1, 0, 0, 0, 130, 129, 1,
		0, 0, 0, 131, 15, 1, 0, 0, 0, 132, 133, 3, 8, 4, 0, 133, 136, 5, 12, 0,
		0, 134, 135, 5, 24, 0, 0, 135, 137, 5, 12, 0, 0, 136, 134, 1, 0, 0, 0,
		136, 137, 1, 0, 0, 0, 137, 17, 1, 0, 0, 0, 138, 139, 5, 8, 0, 0, 139, 143,
		5, 28, 0, 0, 140, 142, 3, 22, 11, 0, 141, 140, 1, 0, 0, 0, 142, 145, 1,
		0, 0, 0, 143, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 146, 1, 0, 0,
		0, 145, 143, 1, 0, 0, 0, 146, 147, 5, 29, 0, 0, 147, 19, 1, 0, 0, 0, 148,
		149, 5, 9, 0, 0, 149, 150, 5, 12, 0, 0, 150, 21, 1, 0, 0, 0, 151, 154,
		3, 8, 4, 0, 152, 155, 5, 12, 0, 0, 153, 155, 3, 8, 4, 0, 154, 152, 1, 0,
		0, 0, 154, 153, 1, 0, 0, 0, 155, 23, 1, 0, 0, 0, 156, 157, 5, 4, 0, 0,
		157, 159, 3, 8, 4, 0, 158, 160, 5, 12, 0, 0, 159, 158, 1, 0, 0, 0, 159,
		160, 1, 0, 0, 0, 160, 162, 1, 0, 0, 0, 161, 163, 3, 26, 13, 0, 162, 161,
		1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 165, 5, 32,
		0, 0, 165, 169, 3, 10, 5, 0, 166, 168, 3, 62, 31, 0, 167, 166, 1, 0, 0,
		0, 168, 171, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170,
		178, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 172, 173, 5, 4, 0, 0, 173, 174,
		3, 8, 4, 0, 174, 175, 5, 32, 0, 0, 175, 176, 3, 28, 14, 0, 176, 178, 1,
		0, 0, 0, 177, 156, 1, 0, 0, 0, 177, 172, 1, 0, 0, 0, 178, 25, 1, 0, 0,
		0, 179, 180, 3, 8, 4, 0, 180, 181, 5, 12, 0, 0, 181, 182, 5, 24, 0, 0,
		182, 183, 5, 12, 0, 0, 183, 27, 1, 0, 0, 0, 184, 186, 3, 30, 15, 0, 185,
		184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 187, 188,
		1, 0, 0, 0, 188, 29, 1, 0, 0, 0, 189, 199, 5, 15, 0, 0, 190, 192, 5, 30,
		0, 0, 191, 193, 3, 30, 15, 0, 192, 191, 1, 0, 0, 0, 193, 194, 1, 0, 0,
		0, 194, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196,
		197, 5, 31, 0, 0, 197, 199, 1, 0, 0, 0, 198, 189, 1, 0, 0, 0, 198, 190,
		1, 0, 0, 0, 199, 202, 1, 0, 0, 0, 200, 201, 5, 23, 0, 0, 201, 203, 5, 12,
		0, 0, 202, 200, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 31, 1, 0, 0, 0,
		204, 205, 5, 5, 0, 0, 205, 206, 3, 8, 4, 0, 206, 207, 5, 32, 0, 0, 207,
		208, 3, 40, 20, 0, 208, 33, 1, 0, 0, 0, 209, 210, 5, 6, 0, 0, 210, 211,
		3, 8, 4, 0, 211, 215, 5, 35, 0, 0, 212, 214, 3, 36, 18, 0, 213, 212, 1,
		0, 0, 0, 214, 217, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 215, 216, 1, 0, 0,
		0, 216, 218, 1, 0, 0, 0, 217, 215, 1, 0, 0, 0, 218, 219, 5, 10, 0, 0, 219,
		35, 1, 0, 0, 0, 220, 221, 3, 8, 4, 0, 221, 222, 3, 40, 20, 0, 222, 223,
		5, 35, 0, 0, 223, 235, 1, 0, 0, 0, 224, 225, 3, 8, 4, 0, 225, 227, 5, 35,
		0, 0, 226, 228, 3, 38, 19, 0, 227, 226, 1, 0, 0, 0, 228, 229, 1, 0, 0,
		0, 229, 227, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231,
		232, 5, 10, 0, 0, 232, 233, 5, 35, 0, 0, 233, 235, 1, 0, 0, 0, 234, 220,
		1, 0, 0, 0, 234, 224, 1, 0, 0, 0, 235, 37, 1, 0, 0, 0, 236, 238, 3, 8,
		4, 0, 237, 239, 3, 40, 20, 0, 238, 237, 1, 0, 0, 0, 238, 239, 1, 0, 0,
		0, 239, 240, 1, 0, 0, 0, 240, 241, 5, 35, 0, 0, 241, 39, 1, 0, 0, 0, 242,
		244, 3, 42, 21, 0, 243, 242, 1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245, 243,
		1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 41, 1, 0, 0, 0, 247, 268, 3, 46,
		23, 0, 248, 268, 5, 15, 0, 0, 249, 268, 3, 44, 22, 0, 250, 255, 3, 8, 4,
		0, 251, 252, 5, 33, 0, 0, 252, 254, 3, 8, 4, 0, 253, 251, 1, 0, 0, 0, 254,
		257, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256, 260,
		1, 0, 0, 0, 257, 255, 1, 0, 0, 0, 258, 259, 5, 34, 0, 0, 259, 261, 5, 12,
		0, 0, 260, 258, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261, 268, 1, 0, 0, 0,
		262, 265, 5, 12, 0, 0, 263, 264, 5, 34, 0, 0, 264, 266, 5, 12, 0, 0, 265,
		263, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 268, 1, 0, 0, 0, 267, 247,
		1, 0, 0, 0, 267, 248, 1, 0, 0, 0, 267, 249, 1, 0, 0, 0, 267, 250, 1, 0,
		0, 0, 267, 262, 1, 0, 0, 0, 268, 43, 1, 0, 0, 0, 269, 270, 5, 32, 0, 0,
		270, 273, 5, 12, 0, 0, 271, 272, 5, 34, 0, 0, 272, 274, 5, 12, 0, 0, 273,
		271, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 45, 1, 0, 0, 0, 275, 276, 3,
		8, 4, 0, 276, 277, 5, 12, 0, 0, 277, 278, 5, 18, 0, 0, 278, 282, 5, 12,
		0, 0, 279, 280, 3, 8, 4, 0, 280, 281, 5, 12, 0, 0, 281, 283, 1, 0, 0, 0,
		282, 279, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 287, 1, 0, 0, 0, 284,
		285, 3, 8, 4, 0, 285, 286, 5, 12, 0, 0, 286, 288, 1, 0, 0, 0, 287, 284,
		1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288, 290, 1, 0, 0, 0, 289, 291, 3, 8,
		4, 0, 290, 289, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 47, 1, 0, 0, 0,
		292, 293, 5, 3, 0, 0, 293, 295, 5, 12, 0, 0, 294, 296, 3, 8, 4, 0, 295,
		294, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 297, 1, 0, 0, 0, 297, 298,
		3, 50, 25, 0, 298, 49, 1, 0, 0, 0, 299, 302, 3, 10, 5, 0, 300, 302, 5,
		15, 0, 0, 301, 299, 1, 0, 0, 0, 301, 300, 1, 0, 0, 0, 302, 303, 1, 0, 0,
		0, 303, 316, 3, 52, 26, 0, 304, 305, 3, 28, 14, 0, 305, 306, 3, 52, 26,
		0, 306, 316, 1, 0, 0, 0, 307, 311, 5, 35, 0, 0, 308, 310, 3, 58, 29, 0,
		309, 308, 1, 0, 0, 0, 310, 313, 1, 0, 0, 0, 311, 309, 1, 0, 0, 0, 311,
		312, 1, 0, 0, 0, 312, 314, 1, 0, 0, 0, 313, 311, 1, 0, 0, 0, 314, 316,
		5, 10, 0, 0, 315, 301, 1, 0, 0, 0, 315, 304, 1, 0, 0, 0, 315, 307, 1, 0,
		0, 0, 316, 51, 1, 0, 0, 0, 317, 319, 5, 35, 0, 0, 318, 317, 1, 0, 0, 0,
		318, 319, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320, 322, 3, 54, 27, 0, 321,
		318, 1, 0, 0, 0, 322, 325, 1, 0, 0, 0, 323, 321, 1, 0, 0, 0, 323, 324,
		1, 0, 0, 0, 324, 332, 1, 0, 0, 0, 325, 323, 1, 0, 0, 0, 326, 328, 5, 35,
		0, 0, 327, 326, 1, 0, 0, 0, 327, 328, 1, 0, 0, 0, 328, 329, 1, 0, 0, 0,
		329, 331, 3, 62, 31, 0, 330, 327, 1, 0, 0, 0, 331, 334, 1, 0, 0, 0, 332,
		330, 1, 0, 0, 0, 332, 333, 1, 0, 0, 0, 333, 53, 1, 0, 0, 0, 334, 332, 1,
		0, 0, 0, 335, 336, 3, 8, 4, 0, 336, 339, 5, 21, 0, 0, 337, 340, 3, 10,
		5, 0, 338, 340, 3, 56, 28, 0, 339, 337, 1, 0, 0, 0, 339, 338, 1, 0, 0,
		0, 340, 55, 1, 0, 0, 0, 341, 346, 3, 42, 21, 0, 342, 343, 5, 22, 0, 0,
		343, 345, 3, 42, 21, 0, 344, 342, 1, 0, 0, 0, 345, 348, 1, 0, 0, 0, 346,
		344, 1, 0, 0, 0, 346, 347, 1, 0, 0, 0, 347, 57, 1, 0, 0, 0, 348, 346, 1,
		0, 0, 0, 349, 354, 3, 62, 31, 0, 350, 351, 3, 8, 4, 0, 351, 352, 3, 60,
		30, 0, 352, 354, 1, 0, 0, 0, 353, 349, 1, 0, 0, 0, 353, 350, 1, 0, 0, 0,
		354, 355, 1, 0, 0, 0, 355, 356, 5, 35, 0, 0, 356, 59, 1, 0, 0, 0, 357,
		369, 3, 10, 5, 0, 358, 369, 5, 15, 0, 0, 359, 362, 3, 8, 4, 0, 360, 362,
		5, 11, 0, 0, 361, 359, 1, 0, 0, 0, 361, 360, 1, 0, 0, 0, 362, 363, 1, 0,
		0, 0, 363, 369, 3, 10, 5, 0, 364, 366, 5, 12, 0, 0, 365, 367, 5, 12, 0,
		0, 366, 365, 1, 0, 0, 0, 366, 367, 1, 0, 0, 0, 367, 369, 1, 0, 0, 0, 368,
		357, 1, 0, 0, 0, 368, 358, 1, 0, 0, 0, 368, 361, 1, 0, 0, 0, 368, 364,
		1, 0, 0, 0, 369, 61, 1, 0, 0, 0, 370, 371, 5, 17, 0, 0, 371, 372, 3, 64,
		32, 0, 372, 63, 1, 0, 0, 0, 373, 375, 3, 8, 4, 0, 374, 376, 3, 66, 33,
		0, 375, 374, 1, 0, 0, 0, 375, 376, 1, 0, 0, 0, 376, 378, 1, 0, 0, 0, 377,
		379, 3, 68, 34, 0, 378, 377, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 65,
		1, 0, 0, 0, 380, 383, 5, 12, 0, 0, 381, 383, 3, 10, 5, 0, 382, 380, 1,
		0, 0, 0, 382, 381, 1, 0, 0, 0, 383, 67, 1, 0, 0, 0, 384, 390, 3, 64, 32,
		0, 385, 386, 5, 26, 0, 0, 386, 387, 3, 64, 32, 0, 387, 388, 5, 27, 0, 0,
		388, 390, 1, 0, 0, 0, 389, 384, 1, 0, 0, 0, 389, 385, 1, 0, 0, 0, 390,
		69, 1, 0, 0, 0, 391, 392, 5, 7, 0, 0, 392, 396, 5, 35, 0, 0, 393, 395,
		3, 72, 36, 0, 394, 393, 1, 0, 0, 0, 395, 398, 1, 0, 0, 0, 396, 394, 1,
		0, 0, 0, 396, 397, 1, 0, 0, 0, 397, 399, 1, 0, 0, 0, 398, 396, 1, 0, 0,
		0, 399, 400, 5, 10, 0, 0, 400, 71, 1, 0, 0, 0, 401, 402, 3, 8, 4, 0, 402,
		403, 3, 74, 37, 0, 403, 404, 5, 35, 0, 0, 404, 73, 1, 0, 0, 0, 405, 406,
		3, 8, 4, 0, 406, 407, 3, 8, 4, 0, 407, 408, 5, 35, 0, 0, 408, 413, 3, 76,
		38, 0, 409, 410, 5, 35, 0, 0, 410, 412, 3, 76, 38, 0, 411, 409, 1, 0, 0,
		0, 412, 415, 1, 0, 0, 0, 413, 411, 1, 0, 0, 0, 413, 414, 1, 0, 0, 0, 414,
		426, 1, 0, 0, 0, 415, 413, 1, 0, 0, 0, 416, 417, 3, 8, 4, 0, 417, 421,
		3, 80, 40, 0, 418, 420, 3, 82, 41, 0, 419, 418, 1, 0, 0, 0, 420, 423, 1,
		0, 0, 0, 421, 419, 1, 0, 0, 0, 421, 422, 1, 0, 0, 0, 422, 426, 1, 0, 0,
		0, 423, 421, 1, 0, 0, 0, 424, 426, 3, 84, 42, 0, 425, 405, 1, 0, 0, 0,
		425, 416, 1, 0, 0, 0, 425, 424, 1, 0, 0, 0, 426, 75, 1, 0, 0, 0, 427, 428,
		3, 78, 39, 0, 428, 438, 5, 21, 0, 0, 429, 430, 3, 8, 4, 0, 430, 434, 3,
		80, 40, 0, 431, 433, 3, 82, 41, 0, 432, 431, 1, 0, 0, 0, 433, 436, 1, 0,
		0, 0, 434, 432, 1, 0, 0, 0, 434, 435, 1, 0, 0, 0, 435, 439, 1, 0, 0, 0,
		436, 434, 1, 0, 0, 0, 437, 439, 3, 84, 42, 0, 438, 429, 1, 0, 0, 0, 438,
		437, 1, 0, 0, 0, 439, 77, 1, 0, 0, 0, 440, 441, 5, 25, 0, 0, 441, 446,
		5, 12, 0, 0, 442, 443, 5, 20, 0, 0, 443, 446, 5, 12, 0, 0, 444, 446, 3,
		8, 4, 0, 445, 440, 1, 0, 0, 0, 445, 442, 1, 0, 0, 0, 445, 444, 1, 0, 0,
		0, 446, 79, 1, 0, 0, 0, 447, 448, 5, 12, 0, 0, 448, 449, 5, 18, 0, 0, 449,
		451, 5, 12, 0, 0, 450, 447, 1, 0, 0, 0, 450, 451, 1, 0, 0, 0, 451, 81,
		1, 0, 0, 0, 452, 453, 7, 4, 0, 0, 453, 457, 5, 12, 0, 0, 454, 455, 5, 19,
		0, 0, 455, 457, 5, 12, 0, 0, 456, 452, 1, 0, 0, 0, 456, 454, 1, 0, 0, 0,
		457, 83, 1, 0, 0, 0, 458, 465, 5, 12, 0, 0, 459, 461, 5, 22, 0, 0, 460,
		459, 1, 0, 0, 0, 460, 461, 1, 0, 0, 0, 461, 462, 1, 0, 0, 0, 462, 464,
		5, 12, 0, 0, 463, 460, 1, 0, 0, 0, 464, 467, 1, 0, 0, 0, 465, 463, 1, 0,
		0, 0, 465, 466, 1, 0, 0, 0, 466, 85, 1, 0, 0, 0, 467, 465, 1, 0, 0, 0,
		58, 91, 105, 111, 125, 130, 136, 143, 154, 159, 162, 169, 177, 187, 194,
		198, 202, 215, 229, 234, 238, 245, 255, 260, 265, 267, 273, 282, 287, 290,
		295, 301, 311, 315, 318, 323, 327, 332, 339, 346, 353, 361, 366, 368, 375,
		378, 382, 389, 396, 413, 421, 425, 434, 438, 445, 450, 456, 460, 465,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// RasterklangDslParserInit initializes any static state used to implement RasterklangDslParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRasterklangDslParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RasterklangDslParserInit() {
	staticData := &RasterklangDslParserStaticData
	staticData.once.Do(rasterklangdslParserInit)
}

// NewRasterklangDslParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRasterklangDslParser(input antlr.TokenStream) *RasterklangDslParser {
	RasterklangDslParserInit()
	this := new(RasterklangDslParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &RasterklangDslParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "RasterklangDsl.g4"

	return this
}

// RasterklangDslParser tokens.
const (
	RasterklangDslParserEOF        = antlr.TokenEOF
	RasterklangDslParserTEMPO      = 1
	RasterklangDslParserSPEED      = 2
	RasterklangDslParserVOICE      = 3
	RasterklangDslParserPATTERN    = 4
	RasterklangDslParserTABLE      = 5
	RasterklangDslParserINSTRUMENT = 6
	RasterklangDslParserFILTER     = 7
	RasterklangDslParserTARGET     = 8
	RasterklangDslParserCHIP       = 9
	RasterklangDslParserEND        = 10
	RasterklangDslParserHEX        = 11
	RasterklangDslParserNUMBER     = 12
	RasterklangDslParserDQSTRING   = 13
	RasterklangDslParserBTSTRING   = 14
	RasterklangDslParserSIGIL      = 15
	RasterklangDslParserIDENT      = 16
	RasterklangDslParserPIPE       = 17
	RasterklangDslParserDOTDOT     = 18
	RasterklangDslParserSHIFT      = 19
	RasterklangDslParserGE         = 20
	RasterklangDslParserCOLON      = 21
	RasterklangDslParserCOMMA      = 22
	RasterklangDslParserSTAR       = 23
	RasterklangDslParserSLASH      = 24
	RasterklangDslParserLT         = 25
	RasterklangDslParserLPAREN     = 26
	RasterklangDslParserRPAREN     = 27
	RasterklangDslParserLBRACE     = 28
	RasterklangDslParserRBRACE     = 29
	RasterklangDslParserLBRACKET   = 30
	RasterklangDslParserRBRACKET   = 31
	RasterklangDslParserEQ         = 32
	RasterklangDslParserPLUS       = 33
	RasterklangDslParserBANG       = 34
	RasterklangDslParserNEWLINE    = 35
	RasterklangDslParserWS         = 36
	RasterklangDslParserCOMMENT    = 37
)

// RasterklangDslParser rules.
const (
	RasterklangDslParserRULE_program            = 0
	RasterklangDslParserRULE_statement          = 1
	RasterklangDslParserRULE_unknownStmt        = 2
	RasterklangDslParserRULE_unknownTail        = 3
	RasterklangDslParserRULE_anyIdent           = 4
	RasterklangDslParserRULE_str                = 5
	RasterklangDslParserRULE_tempoStmt          = 6
	RasterklangDslParserRULE_tempoValue         = 7
	RasterklangDslParserRULE_tempoSetting       = 8
	RasterklangDslParserRULE_targetStmt         = 9
	RasterklangDslParserRULE_chipStmt           = 10
	RasterklangDslParserRULE_targetSetting      = 11
	RasterklangDslParserRULE_patternDecl        = 12
	RasterklangDslParserRULE_patternEach        = 13
	RasterklangDslParserRULE_arrangement        = 14
	RasterklangDslParserRULE_arrangementElement = 15
	RasterklangDslParserRULE_tableDecl          = 16
	RasterklangDslParserRULE_instrumentBlock    = 17
	RasterklangDslParserRULE_instrumentField    = 18
	RasterklangDslParserRULE_tableVerbLine      = 19
	RasterklangDslParserRULE_tableRow           = 20
	RasterklangDslParserRULE_tableToken         = 21
	RasterklangDslParserRULE_absNote            = 22
	RasterklangDslParserRULE_rampGenerator      = 23
	RasterklangDslParserRULE_voiceStmt          = 24
	RasterklangDslParserRULE_voiceBody          = 25
	RasterklangDslParserRULE_voiceShortTail     = 26
	RasterklangDslParserRULE_inlineAttr         = 27
	RasterklangDslParserRULE_valueList          = 28
	RasterklangDslParserRULE_voiceField         = 29
	RasterklangDslParserRULE_voiceFieldValue    = 30
	RasterklangDslParserRULE_pipelineStep       = 31
	RasterklangDslParserRULE_transform          = 32
	RasterklangDslParserRULE_transformArg       = 33
	RasterklangDslParserRULE_transformRef       = 34
	RasterklangDslParserRULE_filterBlock        = 35
	RasterklangDslParserRULE_filterField        = 36
	RasterklangDslParserRULE_filterFieldValue   = 37
	RasterklangDslParserRULE_switchCase         = 38
	RasterklangDslParserRULE_switchCond         = 39
	RasterklangDslParserRULE_signalRangeOpt     = 40
	RasterklangDslParserRULE_signalSpeedMod     = 41
	RasterklangDslParserRULE_numberList         = 42
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserEOF, 0)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(RasterklangDslParserNEWLINE)
}

func (s *ProgramContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNEWLINE, i)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RasterklangDslParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&66558) != 0 {
		{
			p.SetState(86)
			p.Statement()
		}
		{
			p.SetState(87)
			p.Match(RasterklangDslParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(94)
		p.Match(RasterklangDslParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TempoStmt() ITempoStmtContext
	TargetStmt() ITargetStmtContext
	ChipStmt() IChipStmtContext
	VoiceStmt() IVoiceStmtContext
	PatternDecl() IPatternDeclContext
	TableDecl() ITableDeclContext
	InstrumentBlock() IInstrumentBlockContext
	FilterBlock() IFilterBlockContext
	UnknownStmt() IUnknownStmtContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) TempoStmt() ITempoStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITempoStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITempoStmtContext)
}

func (s *StatementContext) TargetStmt() ITargetStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetStmtContext)
}

func (s *StatementContext) ChipStmt() IChipStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChipStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChipStmtContext)
}

func (s *StatementContext) VoiceStmt() IVoiceStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVoiceStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVoiceStmtContext)
}

func (s *StatementContext) PatternDecl() IPatternDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatternDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatternDeclContext)
}

func (s *StatementContext) TableDecl() ITableDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableDeclContext)
}

func (s *StatementContext) InstrumentBlock() IInstrumentBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstrumentBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstrumentBlockContext)
}

func (s *StatementContext) FilterBlock() IFilterBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterBlockContext)
}

func (s *StatementContext) UnknownStmt() IUnknownStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnknownStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnknownStmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RasterklangDslParserRULE_statement)
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RasterklangDslParserTEMPO, RasterklangDslParserSPEED:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(96)
			p.TempoStmt()
		}

	case RasterklangDslParserTARGET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(97)
			p.TargetStmt()
		}

	case RasterklangDslParserCHIP:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(98)
			p.ChipStmt()
		}

	case RasterklangDslParserVOICE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(99)
			p.VoiceStmt()
		}

	case RasterklangDslParserPATTERN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(100)
			p.PatternDecl()
		}

	case RasterklangDslParserTABLE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(101)
			p.TableDecl()
		}

	case RasterklangDslParserINSTRUMENT:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(102)
			p.InstrumentBlock()
		}

	case RasterklangDslParserFILTER:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(103)
			p.FilterBlock()
		}

	case RasterklangDslParserIDENT:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(104)
			p.UnknownStmt()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnknownStmtContext is an interface to support dynamic dispatch.
type IUnknownStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	AllUnknownTail() []IUnknownTailContext
	UnknownTail(i int) IUnknownTailContext

	// IsUnknownStmtContext differentiates from other interfaces.
	IsUnknownStmtContext()
}

type UnknownStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnknownStmtContext() *UnknownStmtContext {
	var p = new(UnknownStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_unknownStmt
	return p
}

func InitEmptyUnknownStmtContext(p *UnknownStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_unknownStmt
}

func (*UnknownStmtContext) IsUnknownStmtContext() {}

func NewUnknownStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnknownStmtContext {
	var p = new(UnknownStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_unknownStmt

	return p
}

func (s *UnknownStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *UnknownStmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserIDENT, 0)
}

func (s *UnknownStmtContext) AllUnknownTail() []IUnknownTailContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnknownTailContext); ok {
			len++
		}
	}

	tst := make([]IUnknownTailContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnknownTailContext); ok {
			tst[i] = t.(IUnknownTailContext)
			i++
		}
	}

	return tst
}

func (s *UnknownStmtContext) UnknownTail(i int) IUnknownTailContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnknownTailContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnknownTailContext)
}

func (s *UnknownStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnknownStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnknownStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterUnknownStmt(s)
	}
}

func (s *UnknownStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitUnknownStmt(s)
	}
}

func (s *UnknownStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitUnknownStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) UnknownStmt() (localctx IUnknownStmtContext) {
	localctx = NewUnknownStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RasterklangDslParserRULE_unknownStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Match(RasterklangDslParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&240518168574) != 0 {
		{
			p.SetState(108)
			p.UnknownTail()
		}

		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnknownTailContext is an interface to support dynamic dispatch.
type IUnknownTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NEWLINE() antlr.TerminalNode

	// IsUnknownTailContext differentiates from other interfaces.
	IsUnknownTailContext()
}

type UnknownTailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnknownTailContext() *UnknownTailContext {
	var p = new(UnknownTailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_unknownTail
	return p
}

func InitEmptyUnknownTailContext(p *UnknownTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_unknownTail
}

func (*UnknownTailContext) IsUnknownTailContext() {}

func NewUnknownTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnknownTailContext {
	var p = new(UnknownTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_unknownTail

	return p
}

func (s *UnknownTailContext) GetParser() antlr.Parser { return s.parser }

func (s *UnknownTailContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNEWLINE, 0)
}

func (s *UnknownTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnknownTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnknownTailContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterUnknownTail(s)
	}
}

func (s *UnknownTailContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitUnknownTail(s)
	}
}

func (s *UnknownTailContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitUnknownTail(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) UnknownTail() (localctx IUnknownTailContext) {
	localctx = NewUnknownTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RasterklangDslParserRULE_unknownTail)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		_la = p.GetTokenStream().LA(1)

		if _la <= 0 || _la == RasterklangDslParserNEWLINE {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAnyIdentContext is an interface to support dynamic dispatch.
type IAnyIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	TEMPO() antlr.TerminalNode
	SPEED() antlr.TerminalNode
	VOICE() antlr.TerminalNode
	PATTERN() antlr.TerminalNode
	TABLE() antlr.TerminalNode
	INSTRUMENT() antlr.TerminalNode
	FILTER() antlr.TerminalNode
	TARGET() antlr.TerminalNode
	CHIP() antlr.TerminalNode

	// IsAnyIdentContext differentiates from other interfaces.
	IsAnyIdentContext()
}

type AnyIdentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnyIdentContext() *AnyIdentContext {
	var p = new(AnyIdentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_anyIdent
	return p
}

func InitEmptyAnyIdentContext(p *AnyIdentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_anyIdent
}

func (*AnyIdentContext) IsAnyIdentContext() {}

func NewAnyIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnyIdentContext {
	var p = new(AnyIdentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_anyIdent

	return p
}

func (s *AnyIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *AnyIdentContext) IDENT() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserIDENT, 0)
}

func (s *AnyIdentContext) TEMPO() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserTEMPO, 0)
}

func (s *AnyIdentContext) SPEED() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserSPEED, 0)
}

func (s *AnyIdentContext) VOICE() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserVOICE, 0)
}

func (s *AnyIdentContext) PATTERN() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserPATTERN, 0)
}

func (s *AnyIdentContext) TABLE() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserTABLE, 0)
}

func (s *AnyIdentContext) INSTRUMENT() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserINSTRUMENT, 0)
}

func (s *AnyIdentContext) FILTER() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserFILTER, 0)
}

func (s *AnyIdentContext) TARGET() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserTARGET, 0)
}

func (s *AnyIdentContext) CHIP() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserCHIP, 0)
}

func (s *AnyIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnyIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnyIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterAnyIdent(s)
	}
}

func (s *AnyIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitAnyIdent(s)
	}
}

func (s *AnyIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitAnyIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) AnyIdent() (localctx IAnyIdentContext) {
	localctx = NewAnyIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RasterklangDslParserRULE_anyIdent)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&66558) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStrContext is an interface to support dynamic dispatch.
type IStrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DQSTRING() antlr.TerminalNode
	BTSTRING() antlr.TerminalNode

	// IsStrContext differentiates from other interfaces.
	IsStrContext()
}

type StrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrContext() *StrContext {
	var p = new(StrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_str
	return p
}

func InitEmptyStrContext(p *StrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_str
}

func (*StrContext) IsStrContext() {}

func NewStrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrContext {
	var p = new(StrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_str

	return p
}

func (s *StrContext) GetParser() antlr.Parser { return s.parser }

func (s *StrContext) DQSTRING() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserDQSTRING, 0)
}

func (s *StrContext) BTSTRING() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserBTSTRING, 0)
}

func (s *StrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterStr(s)
	}
}

func (s *StrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitStr(s)
	}
}

func (s *StrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitStr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) Str() (localctx IStrContext) {
	localctx = NewStrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RasterklangDslParserRULE_str)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RasterklangDslParserDQSTRING || _la == RasterklangDslParserBTSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITempoStmtContext is an interface to support dynamic dispatch.
type ITempoStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TempoValue() ITempoValueContext
	TEMPO() antlr.TerminalNode
	SPEED() antlr.TerminalNode
	AllTempoSetting() []ITempoSettingContext
	TempoSetting(i int) ITempoSettingContext

	// IsTempoStmtContext differentiates from other interfaces.
	IsTempoStmtContext()
}

type TempoStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTempoStmtContext() *TempoStmtContext {
	var p = new(TempoStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_tempoStmt
	return p
}

func InitEmptyTempoStmtContext(p *TempoStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_tempoStmt
}

func (*TempoStmtContext) IsTempoStmtContext() {}

func NewTempoStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TempoStmtContext {
	var p = new(TempoStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_tempoStmt

	return p
}

func (s *TempoStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *TempoStmtContext) TempoValue() ITempoValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITempoValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITempoValueContext)
}

func (s *TempoStmtContext) TEMPO() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserTEMPO, 0)
}

func (s *TempoStmtContext) SPEED() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserSPEED, 0)
}

func (s *TempoStmtContext) AllTempoSetting() []ITempoSettingContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITempoSettingContext); ok {
			len++
		}
	}

	tst := make([]ITempoSettingContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITempoSettingContext); ok {
			tst[i] = t.(ITempoSettingContext)
			i++
		}
	}

	return tst
}

func (s *TempoStmtContext) TempoSetting(i int) ITempoSettingContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITempoSettingContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITempoSettingContext)
}

func (s *TempoStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TempoStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TempoStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterTempoStmt(s)
	}
}

func (s *TempoStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitTempoStmt(s)
	}
}

func (s *TempoStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitTempoStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) TempoStmt() (localctx ITempoStmtContext) {
	localctx = NewTempoStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RasterklangDslParserRULE_tempoStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RasterklangDslParserTEMPO || _la == RasterklangDslParserSPEED) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(121)
		p.TempoValue()
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&66558) != 0 {
		{
			p.SetState(122)
			p.TempoSetting()
		}

		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITempoValueContext is an interface to support dynamic dispatch.
type ITempoValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	Str() IStrContext

	// IsTempoValueContext differentiates from other interfaces.
	IsTempoValueContext()
}

type TempoValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTempoValueContext() *TempoValueContext {
	var p = new(TempoValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_tempoValue
	return p
}

func InitEmptyTempoValueContext(p *TempoValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_tempoValue
}

func (*TempoValueContext) IsTempoValueContext() {}

func NewTempoValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TempoValueContext {
	var p = new(TempoValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_tempoValue

	return p
}

func (s *TempoValueContext) GetParser() antlr.Parser { return s.parser }

func (s *TempoValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNUMBER, 0)
}

func (s *TempoValueContext) Str() IStrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrContext)
}

func (s *TempoValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TempoValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TempoValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterTempoValue(s)
	}
}

func (s *TempoValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitTempoValue(s)
	}
}

func (s *TempoValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitTempoValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) TempoValue() (localctx ITempoValueContext) {
	localctx = NewTempoValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RasterklangDslParserRULE_tempoValue)
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RasterklangDslParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.Match(RasterklangDslParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RasterklangDslParserDQSTRING, RasterklangDslParserBTSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.Str()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITempoSettingContext is an interface to support dynamic dispatch.
type ITempoSettingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AnyIdent() IAnyIdentContext
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode
	SLASH() antlr.TerminalNode

	// IsTempoSettingContext differentiates from other interfaces.
	IsTempoSettingContext()
}

type TempoSettingContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTempoSettingContext() *TempoSettingContext {
	var p = new(TempoSettingContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_tempoSetting
	return p
}

func InitEmptyTempoSettingContext(p *TempoSettingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_tempoSetting
}

func (*TempoSettingContext) IsTempoSettingContext() {}

func NewTempoSettingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TempoSettingContext {
	var p = new(TempoSettingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_tempoSetting

	return p
}

func (s *TempoSettingContext) GetParser() antlr.Parser { return s.parser }

func (s *TempoSettingContext) AnyIdent() IAnyIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyIdentContext)
}

func (s *TempoSettingContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(RasterklangDslParserNUMBER)
}

func (s *TempoSettingContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNUMBER, i)
}

func (s *TempoSettingContext) SLASH() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserSLASH, 0)
}

func (s *TempoSettingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TempoSettingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TempoSettingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterTempoSetting(s)
	}
}

func (s *TempoSettingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitTempoSetting(s)
	}
}

func (s *TempoSettingContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitTempoSetting(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) TempoSetting() (localctx ITempoSettingContext) {
	localctx = NewTempoSettingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RasterklangDslParserRULE_tempoSetting)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.AnyIdent()
	}
	{
		p.SetState(133)
		p.Match(RasterklangDslParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == RasterklangDslParserSLASH {
		{
			p.SetState(134)
			p.Match(RasterklangDslParserSLASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.Match(RasterklangDslParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITargetStmtContext is an interface to support dynamic dispatch.
type ITargetStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TARGET() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllTargetSetting() []ITargetSettingContext
	TargetSetting(i int) ITargetSettingContext

	// IsTargetStmtContext differentiates from other interfaces.
	IsTargetStmtContext()
}

type TargetStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetStmtContext() *TargetStmtContext {
	var p = new(TargetStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_targetStmt
	return p
}

func InitEmptyTargetStmtContext(p *TargetStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_targetStmt
}

func (*TargetStmtContext) IsTargetStmtContext() {}

func NewTargetStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetStmtContext {
	var p = new(TargetStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_targetStmt

	return p
}

func (s *TargetStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetStmtContext) TARGET() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserTARGET, 0)
}

func (s *TargetStmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserLBRACE, 0)
}

func (s *TargetStmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserRBRACE, 0)
}

func (s *TargetStmtContext) AllTargetSetting() []ITargetSettingContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITargetSettingContext); ok {
			len++
		}
	}

	tst := make([]ITargetSettingContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITargetSettingContext); ok {
			tst[i] = t.(ITargetSettingContext)
			i++
		}
	}

	return tst
}

func (s *TargetStmtContext) TargetSetting(i int) ITargetSettingContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetSettingContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetSettingContext)
}

func (s *TargetStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterTargetStmt(s)
	}
}

func (s *TargetStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitTargetStmt(s)
	}
}

func (s *TargetStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitTargetStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) TargetStmt() (localctx ITargetStmtContext) {
	localctx = NewTargetStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RasterklangDslParserRULE_targetStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(RasterklangDslParserTARGET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Match(RasterklangDslParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&66558) != 0 {
		{
			p.SetState(140)
			p.TargetSetting()
		}

		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(146)
		p.Match(RasterklangDslParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IChipStmtContext is an interface to support dynamic dispatch.
type IChipStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CHIP() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsChipStmtContext differentiates from other interfaces.
	IsChipStmtContext()
}

type ChipStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChipStmtContext() *ChipStmtContext {
	var p = new(ChipStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_chipStmt
	return p
}

func InitEmptyChipStmtContext(p *ChipStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_chipStmt
}

func (*ChipStmtContext) IsChipStmtContext() {}

func NewChipStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChipStmtContext {
	var p = new(ChipStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_chipStmt

	return p
}

func (s *ChipStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ChipStmtContext) CHIP() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserCHIP, 0)
}

func (s *ChipStmtContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNUMBER, 0)
}

func (s *ChipStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChipStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChipStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterChipStmt(s)
	}
}

func (s *ChipStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitChipStmt(s)
	}
}

func (s *ChipStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitChipStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) ChipStmt() (localctx IChipStmtContext) {
	localctx = NewChipStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, RasterklangDslParserRULE_chipStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(RasterklangDslParserCHIP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.Match(RasterklangDslParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITargetSettingContext is an interface to support dynamic dispatch.
type ITargetSettingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAnyIdent() []IAnyIdentContext
	AnyIdent(i int) IAnyIdentContext
	NUMBER() antlr.TerminalNode

	// IsTargetSettingContext differentiates from other interfaces.
	IsTargetSettingContext()
}

type TargetSettingContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetSettingContext() *TargetSettingContext {
	var p = new(TargetSettingContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_targetSetting
	return p
}

func InitEmptyTargetSettingContext(p *TargetSettingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_targetSetting
}

func (*TargetSettingContext) IsTargetSettingContext() {}

func NewTargetSettingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetSettingContext {
	var p = new(TargetSettingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_targetSetting

	return p
}

func (s *TargetSettingContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetSettingContext) AllAnyIdent() []IAnyIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAnyIdentContext); ok {
			len++
		}
	}

	tst := make([]IAnyIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAnyIdentContext); ok {
			tst[i] = t.(IAnyIdentContext)
			i++
		}
	}

	return tst
}

func (s *TargetSettingContext) AnyIdent(i int) IAnyIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyIdentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyIdentContext)
}

func (s *TargetSettingContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNUMBER, 0)
}

func (s *TargetSettingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetSettingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetSettingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterTargetSetting(s)
	}
}

func (s *TargetSettingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitTargetSetting(s)
	}
}

func (s *TargetSettingContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitTargetSetting(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) TargetSetting() (localctx ITargetSettingContext) {
	localctx = NewTargetSettingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RasterklangDslParserRULE_targetSetting)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.AnyIdent()
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RasterklangDslParserNUMBER:
		{
			p.SetState(152)
			p.Match(RasterklangDslParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RasterklangDslParserTEMPO, RasterklangDslParserSPEED, RasterklangDslParserVOICE, RasterklangDslParserPATTERN, RasterklangDslParserTABLE, RasterklangDslParserINSTRUMENT, RasterklangDslParserFILTER, RasterklangDslParserTARGET, RasterklangDslParserCHIP, RasterklangDslParserIDENT:
		{
			p.SetState(153)
			p.AnyIdent()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPatternDeclContext is an interface to support dynamic dispatch.
type IPatternDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PATTERN() antlr.TerminalNode
	AnyIdent() IAnyIdentContext
	EQ() antlr.TerminalNode
	Str() IStrContext
	NUMBER() antlr.TerminalNode
	PatternEach() IPatternEachContext
	AllPipelineStep() []IPipelineStepContext
	PipelineStep(i int) IPipelineStepContext
	Arrangement() IArrangementContext

	// IsPatternDeclContext differentiates from other interfaces.
	IsPatternDeclContext()
}

type PatternDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatternDeclContext() *PatternDeclContext {
	var p = new(PatternDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_patternDecl
	return p
}

func InitEmptyPatternDeclContext(p *PatternDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_patternDecl
}

func (*PatternDeclContext) IsPatternDeclContext() {}

func NewPatternDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatternDeclContext {
	var p = new(PatternDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_patternDecl

	return p
}

func (s *PatternDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *PatternDeclContext) PATTERN() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserPATTERN, 0)
}

func (s *PatternDeclContext) AnyIdent() IAnyIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyIdentContext)
}

func (s *PatternDeclContext) EQ() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserEQ, 0)
}

func (s *PatternDeclContext) Str() IStrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrContext)
}

func (s *PatternDeclContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNUMBER, 0)
}

func (s *PatternDeclContext) PatternEach() IPatternEachContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatternEachContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatternEachContext)
}

func (s *PatternDeclContext) AllPipelineStep() []IPipelineStepContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPipelineStepContext); ok {
			len++
		}
	}

	tst := make([]IPipelineStepContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPipelineStepContext); ok {
			tst[i] = t.(IPipelineStepContext)
			i++
		}
	}

	return tst
}

func (s *PatternDeclContext) PipelineStep(i int) IPipelineStepContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipelineStepContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipelineStepContext)
}

func (s *PatternDeclContext) Arrangement() IArrangementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrangementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrangementContext)
}

func (s *PatternDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PatternDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterPatternDecl(s)
	}
}

func (s *PatternDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitPatternDecl(s)
	}
}

func (s *PatternDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitPatternDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) PatternDecl() (localctx IPatternDeclContext) {
	localctx = NewPatternDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, RasterklangDslParserRULE_patternDecl)
	var _la int

	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(156)
			p.Match(RasterklangDslParserPATTERN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)
			p.AnyIdent()
		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == RasterklangDslParserNUMBER {
			{
				p.SetState(158)
				p.Match(RasterklangDslParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&66558) != 0 {
			{
				p.SetState(161)
				p.PatternEach()
			}

		}
		{
			p.SetState(164)
			p.Match(RasterklangDslParserEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)
			p.Str()
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RasterklangDslParserPIPE {
			{
				p.SetState(166)
				p.PipelineStep()
			}

			p.SetState(171)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(172)
			p.Match(RasterklangDslParserPATTERN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(173)
			p.AnyIdent()
		}
		{
			p.SetState(174)
			p.Match(RasterklangDslParserEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(175)
			p.Arrangement()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPatternEachContext is an interface to support dynamic dispatch.
type IPatternEachContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AnyIdent() IAnyIdentContext
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode
	SLASH() antlr.TerminalNode

	// IsPatternEachContext differentiates from other interfaces.
	IsPatternEachContext()
}

type PatternEachContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatternEachContext() *PatternEachContext {
	var p = new(PatternEachContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_patternEach
	return p
}

func InitEmptyPatternEachContext(p *PatternEachContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_patternEach
}

func (*PatternEachContext) IsPatternEachContext() {}

func NewPatternEachContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatternEachContext {
	var p = new(PatternEachContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_patternEach

	return p
}

func (s *PatternEachContext) GetParser() antlr.Parser { return s.parser }

func (s *PatternEachContext) AnyIdent() IAnyIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyIdentContext)
}

func (s *PatternEachContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(RasterklangDslParserNUMBER)
}

func (s *PatternEachContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNUMBER, i)
}

func (s *PatternEachContext) SLASH() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserSLASH, 0)
}

func (s *PatternEachContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternEachContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PatternEachContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterPatternEach(s)
	}
}

func (s *PatternEachContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitPatternEach(s)
	}
}

func (s *PatternEachContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitPatternEach(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) PatternEach() (localctx IPatternEachContext) {
	localctx = NewPatternEachContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, RasterklangDslParserRULE_patternEach)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.AnyIdent()
	}
	{
		p.SetState(180)
		p.Match(RasterklangDslParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
		p.Match(RasterklangDslParserSLASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.Match(RasterklangDslParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrangementContext is an interface to support dynamic dispatch.
type IArrangementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllArrangementElement() []IArrangementElementContext
	ArrangementElement(i int) IArrangementElementContext

	// IsArrangementContext differentiates from other interfaces.
	IsArrangementContext()
}

type ArrangementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrangementContext() *ArrangementContext {
	var p = new(ArrangementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_arrangement
	return p
}

func InitEmptyArrangementContext(p *ArrangementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_arrangement
}

func (*ArrangementContext) IsArrangementContext() {}

func NewArrangementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrangementContext {
	var p = new(ArrangementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_arrangement

	return p
}

func (s *ArrangementContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrangementContext) AllArrangementElement() []IArrangementElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArrangementElementContext); ok {
			len++
		}
	}

	tst := make([]IArrangementElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArrangementElementContext); ok {
			tst[i] = t.(IArrangementElementContext)
			i++
		}
	}

	return tst
}

func (s *ArrangementContext) ArrangementElement(i int) IArrangementElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrangementElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrangementElementContext)
}

func (s *ArrangementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrangementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrangementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterArrangement(s)
	}
}

func (s *ArrangementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitArrangement(s)
	}
}

func (s *ArrangementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitArrangement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) Arrangement() (localctx IArrangementContext) {
	localctx = NewArrangementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, RasterklangDslParserRULE_arrangement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == RasterklangDslParserSIGIL || _la == RasterklangDslParserLBRACKET {
		{
			p.SetState(184)
			p.ArrangementElement()
		}

		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrangementElementContext is an interface to support dynamic dispatch.
type IArrangementElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SIGIL() antlr.TerminalNode
	LBRACKET() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode
	STAR() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	AllArrangementElement() []IArrangementElementContext
	ArrangementElement(i int) IArrangementElementContext

	// IsArrangementElementContext differentiates from other interfaces.
	IsArrangementElementContext()
}

type ArrangementElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrangementElementContext() *ArrangementElementContext {
	var p = new(ArrangementElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_arrangementElement
	return p
}

func InitEmptyArrangementElementContext(p *ArrangementElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_arrangementElement
}

func (*ArrangementElementContext) IsArrangementElementContext() {}

func NewArrangementElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrangementElementContext {
	var p = new(ArrangementElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_arrangementElement

	return p
}

func (s *ArrangementElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrangementElementContext) SIGIL() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserSIGIL, 0)
}

func (s *ArrangementElementContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserLBRACKET, 0)
}

func (s *ArrangementElementContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserRBRACKET, 0)
}

func (s *ArrangementElementContext) STAR() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserSTAR, 0)
}

func (s *ArrangementElementContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNUMBER, 0)
}

func (s *ArrangementElementContext) AllArrangementElement() []IArrangementElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArrangementElementContext); ok {
			len++
		}
	}

	tst := make([]IArrangementElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArrangementElementContext); ok {
			tst[i] = t.(IArrangementElementContext)
			i++
		}
	}

	return tst
}

func (s *ArrangementElementContext) ArrangementElement(i int) IArrangementElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrangementElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrangementElementContext)
}

func (s *ArrangementElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrangementElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrangementElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterArrangementElement(s)
	}
}

func (s *ArrangementElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitArrangementElement(s)
	}
}

func (s *ArrangementElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitArrangementElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) ArrangementElement() (localctx IArrangementElementContext) {
	localctx = NewArrangementElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, RasterklangDslParserRULE_arrangementElement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RasterklangDslParserSIGIL:
		{
			p.SetState(189)
			p.Match(RasterklangDslParserSIGIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RasterklangDslParserLBRACKET:
		{
			p.SetState(190)
			p.Match(RasterklangDslParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == RasterklangDslParserSIGIL || _la == RasterklangDslParserLBRACKET {
			{
				p.SetState(191)
				p.ArrangementElement()
			}

			p.SetState(194)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(196)
			p.Match(RasterklangDslParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == RasterklangDslParserSTAR {
		{
			p.SetState(200)
			p.Match(RasterklangDslParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
			p.Match(RasterklangDslParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITableDeclContext is an interface to support dynamic dispatch.
type ITableDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TABLE() antlr.TerminalNode
	AnyIdent() IAnyIdentContext
	EQ() antlr.TerminalNode
	TableRow() ITableRowContext

	// IsTableDeclContext differentiates from other interfaces.
	IsTableDeclContext()
}

type TableDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableDeclContext() *TableDeclContext {
	var p = new(TableDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_tableDecl
	return p
}

func InitEmptyTableDeclContext(p *TableDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_tableDecl
}

func (*TableDeclContext) IsTableDeclContext() {}

func NewTableDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableDeclContext {
	var p = new(TableDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_tableDecl

	return p
}

func (s *TableDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TableDeclContext) TABLE() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserTABLE, 0)
}

func (s *TableDeclContext) AnyIdent() IAnyIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyIdentContext)
}

func (s *TableDeclContext) EQ() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserEQ, 0)
}

func (s *TableDeclContext) TableRow() ITableRowContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableRowContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableRowContext)
}

func (s *TableDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterTableDecl(s)
	}
}

func (s *TableDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitTableDecl(s)
	}
}

func (s *TableDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitTableDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) TableDecl() (localctx ITableDeclContext) {
	localctx = NewTableDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, RasterklangDslParserRULE_tableDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Match(RasterklangDslParserTABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.AnyIdent()
	}
	{
		p.SetState(206)
		p.Match(RasterklangDslParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(207)
		p.TableRow()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstrumentBlockContext is an interface to support dynamic dispatch.
type IInstrumentBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INSTRUMENT() antlr.TerminalNode
	AnyIdent() IAnyIdentContext
	NEWLINE() antlr.TerminalNode
	END() antlr.TerminalNode
	AllInstrumentField() []IInstrumentFieldContext
	InstrumentField(i int) IInstrumentFieldContext

	// IsInstrumentBlockContext differentiates from other interfaces.
	IsInstrumentBlockContext()
}

type InstrumentBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstrumentBlockContext() *InstrumentBlockContext {
	var p = new(InstrumentBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_instrumentBlock
	return p
}

func InitEmptyInstrumentBlockContext(p *InstrumentBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_instrumentBlock
}

func (*InstrumentBlockContext) IsInstrumentBlockContext() {}

func NewInstrumentBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstrumentBlockContext {
	var p = new(InstrumentBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_instrumentBlock

	return p
}

func (s *InstrumentBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *InstrumentBlockContext) INSTRUMENT() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserINSTRUMENT, 0)
}

func (s *InstrumentBlockContext) AnyIdent() IAnyIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyIdentContext)
}

func (s *InstrumentBlockContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNEWLINE, 0)
}

func (s *InstrumentBlockContext) END() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserEND, 0)
}

func (s *InstrumentBlockContext) AllInstrumentField() []IInstrumentFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstrumentFieldContext); ok {
			len++
		}
	}

	tst := make([]IInstrumentFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstrumentFieldContext); ok {
			tst[i] = t.(IInstrumentFieldContext)
			i++
		}
	}

	return tst
}

func (s *InstrumentBlockContext) InstrumentField(i int) IInstrumentFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstrumentFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstrumentFieldContext)
}

func (s *InstrumentBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstrumentBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstrumentBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterInstrumentBlock(s)
	}
}

func (s *InstrumentBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitInstrumentBlock(s)
	}
}

func (s *InstrumentBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitInstrumentBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) InstrumentBlock() (localctx IInstrumentBlockContext) {
	localctx = NewInstrumentBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, RasterklangDslParserRULE_instrumentBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		p.Match(RasterklangDslParserINSTRUMENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(210)
		p.AnyIdent()
	}
	{
		p.SetState(211)
		p.Match(RasterklangDslParserNEWLINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&66558) != 0 {
		{
			p.SetState(212)
			p.InstrumentField()
		}

		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(218)
		p.Match(RasterklangDslParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstrumentFieldContext is an interface to support dynamic dispatch.
type IInstrumentFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AnyIdent() IAnyIdentContext
	TableRow() ITableRowContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	END() antlr.TerminalNode
	AllTableVerbLine() []ITableVerbLineContext
	TableVerbLine(i int) ITableVerbLineContext

	// IsInstrumentFieldContext differentiates from other interfaces.
	IsInstrumentFieldContext()
}

type InstrumentFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstrumentFieldContext() *InstrumentFieldContext {
	var p = new(InstrumentFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_instrumentField
	return p
}

func InitEmptyInstrumentFieldContext(p *InstrumentFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_instrumentField
}

func (*InstrumentFieldContext) IsInstrumentFieldContext() {}

func NewInstrumentFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstrumentFieldContext {
	var p = new(InstrumentFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_instrumentField

	return p
}

func (s *InstrumentFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *InstrumentFieldContext) AnyIdent() IAnyIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyIdentContext)
}

func (s *InstrumentFieldContext) TableRow() ITableRowContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableRowContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableRowContext)
}

func (s *InstrumentFieldContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(RasterklangDslParserNEWLINE)
}

func (s *InstrumentFieldContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNEWLINE, i)
}

func (s *InstrumentFieldContext) END() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserEND, 0)
}

func (s *InstrumentFieldContext) AllTableVerbLine() []ITableVerbLineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITableVerbLineContext); ok {
			len++
		}
	}

	tst := make([]ITableVerbLineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITableVerbLineContext); ok {
			tst[i] = t.(ITableVerbLineContext)
			i++
		}
	}

	return tst
}

func (s *InstrumentFieldContext) TableVerbLine(i int) ITableVerbLineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableVerbLineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableVerbLineContext)
}

func (s *InstrumentFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstrumentFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstrumentFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterInstrumentField(s)
	}
}

func (s *InstrumentFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitInstrumentField(s)
	}
}

func (s *InstrumentFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitInstrumentField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) InstrumentField() (localctx IInstrumentFieldContext) {
	localctx = NewInstrumentFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, RasterklangDslParserRULE_instrumentField)
	var _la int

	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(220)
			p.AnyIdent()
		}
		{
			p.SetState(221)
			p.TableRow()
		}
		{
			p.SetState(222)
			p.Match(RasterklangDslParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(224)
			p.AnyIdent()
		}
		{
			p.SetState(225)
			p.Match(RasterklangDslParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&66558) != 0) {
			{
				p.SetState(226)
				p.TableVerbLine()
			}

			p.SetState(229)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(231)
			p.Match(RasterklangDslParserEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(232)
			p.Match(RasterklangDslParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITableVerbLineContext is an interface to support dynamic dispatch.
type ITableVerbLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AnyIdent() IAnyIdentContext
	NEWLINE() antlr.TerminalNode
	TableRow() ITableRowContext

	// IsTableVerbLineContext differentiates from other interfaces.
	IsTableVerbLineContext()
}

type TableVerbLineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableVerbLineContext() *TableVerbLineContext {
	var p = new(TableVerbLineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_tableVerbLine
	return p
}

func InitEmptyTableVerbLineContext(p *TableVerbLineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_tableVerbLine
}

func (*TableVerbLineContext) IsTableVerbLineContext() {}

func NewTableVerbLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableVerbLineContext {
	var p = new(TableVerbLineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_tableVerbLine

	return p
}

func (s *TableVerbLineContext) GetParser() antlr.Parser { return s.parser }

func (s *TableVerbLineContext) AnyIdent() IAnyIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyIdentContext)
}

func (s *TableVerbLineContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNEWLINE, 0)
}

func (s *TableVerbLineContext) TableRow() ITableRowContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableRowContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableRowContext)
}

func (s *TableVerbLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableVerbLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableVerbLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterTableVerbLine(s)
	}
}

func (s *TableVerbLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitTableVerbLine(s)
	}
}

func (s *TableVerbLineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitTableVerbLine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) TableVerbLine() (localctx ITableVerbLineContext) {
	localctx = NewTableVerbLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, RasterklangDslParserRULE_tableVerbLine)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.AnyIdent()
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4295070718) != 0 {
		{
			p.SetState(237)
			p.TableRow()
		}

	}
	{
		p.SetState(240)
		p.Match(RasterklangDslParserNEWLINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITableRowContext is an interface to support dynamic dispatch.
type ITableRowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTableToken() []ITableTokenContext
	TableToken(i int) ITableTokenContext

	// IsTableRowContext differentiates from other interfaces.
	IsTableRowContext()
}

type TableRowContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableRowContext() *TableRowContext {
	var p = new(TableRowContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_tableRow
	return p
}

func InitEmptyTableRowContext(p *TableRowContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_tableRow
}

func (*TableRowContext) IsTableRowContext() {}

func NewTableRowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableRowContext {
	var p = new(TableRowContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_tableRow

	return p
}

func (s *TableRowContext) GetParser() antlr.Parser { return s.parser }

func (s *TableRowContext) AllTableToken() []ITableTokenContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITableTokenContext); ok {
			len++
		}
	}

	tst := make([]ITableTokenContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITableTokenContext); ok {
			tst[i] = t.(ITableTokenContext)
			i++
		}
	}

	return tst
}

func (s *TableRowContext) TableToken(i int) ITableTokenContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableTokenContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableTokenContext)
}

func (s *TableRowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableRowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableRowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterTableRow(s)
	}
}

func (s *TableRowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitTableRow(s)
	}
}

func (s *TableRowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitTableRow(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) TableRow() (localctx ITableRowContext) {
	localctx = NewTableRowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, RasterklangDslParserRULE_tableRow)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4295070718) != 0) {
		{
			p.SetState(242)
			p.TableToken()
		}

		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITableTokenContext is an interface to support dynamic dispatch.
type ITableTokenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RampGenerator() IRampGeneratorContext
	SIGIL() antlr.TerminalNode
	AbsNote() IAbsNoteContext
	AllAnyIdent() []IAnyIdentContext
	AnyIdent(i int) IAnyIdentContext
	AllPLUS() []antlr.TerminalNode
	PLUS(i int) antlr.TerminalNode
	BANG() antlr.TerminalNode
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode

	// IsTableTokenContext differentiates from other interfaces.
	IsTableTokenContext()
}

type TableTokenContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableTokenContext() *TableTokenContext {
	var p = new(TableTokenContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_tableToken
	return p
}

func InitEmptyTableTokenContext(p *TableTokenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_tableToken
}

func (*TableTokenContext) IsTableTokenContext() {}

func NewTableTokenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableTokenContext {
	var p = new(TableTokenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_tableToken

	return p
}

func (s *TableTokenContext) GetParser() antlr.Parser { return s.parser }

func (s *TableTokenContext) RampGenerator() IRampGeneratorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRampGeneratorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRampGeneratorContext)
}

func (s *TableTokenContext) SIGIL() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserSIGIL, 0)
}

func (s *TableTokenContext) AbsNote() IAbsNoteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAbsNoteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAbsNoteContext)
}

func (s *TableTokenContext) AllAnyIdent() []IAnyIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAnyIdentContext); ok {
			len++
		}
	}

	tst := make([]IAnyIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAnyIdentContext); ok {
			tst[i] = t.(IAnyIdentContext)
			i++
		}
	}

	return tst
}

func (s *TableTokenContext) AnyIdent(i int) IAnyIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyIdentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyIdentContext)
}

func (s *TableTokenContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(RasterklangDslParserPLUS)
}

func (s *TableTokenContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserPLUS, i)
}

func (s *TableTokenContext) BANG() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserBANG, 0)
}

func (s *TableTokenContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(RasterklangDslParserNUMBER)
}

func (s *TableTokenContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNUMBER, i)
}

func (s *TableTokenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableTokenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableTokenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterTableToken(s)
	}
}

func (s *TableTokenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitTableToken(s)
	}
}

func (s *TableTokenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitTableToken(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) TableToken() (localctx ITableTokenContext) {
	localctx = NewTableTokenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, RasterklangDslParserRULE_tableToken)
	var _la int

	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(247)
			p.RampGenerator()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(248)
			p.Match(RasterklangDslParserSIGIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(249)
			p.AbsNote()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(250)
			p.AnyIdent()
		}
		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RasterklangDslParserPLUS {
			{
				p.SetState(251)
				p.Match(RasterklangDslParserPLUS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(252)
				p.AnyIdent()
			}

			p.SetState(257)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(260)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == RasterklangDslParserBANG {
			{
				p.SetState(258)
				p.Match(RasterklangDslParserBANG)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(259)
				p.Match(RasterklangDslParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(262)
			p.Match(RasterklangDslParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(265)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == RasterklangDslParserBANG {
			{
				p.SetState(263)
				p.Match(RasterklangDslParserBANG)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(264)
				p.Match(RasterklangDslParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAbsNoteContext is an interface to support dynamic dispatch.
type IAbsNoteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ() antlr.TerminalNode
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode
	BANG() antlr.TerminalNode

	// IsAbsNoteContext differentiates from other interfaces.
	IsAbsNoteContext()
}

type AbsNoteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAbsNoteContext() *AbsNoteContext {
	var p = new(AbsNoteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_absNote
	return p
}

func InitEmptyAbsNoteContext(p *AbsNoteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_absNote
}

func (*AbsNoteContext) IsAbsNoteContext() {}

func NewAbsNoteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AbsNoteContext {
	var p = new(AbsNoteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_absNote

	return p
}

func (s *AbsNoteContext) GetParser() antlr.Parser { return s.parser }

func (s *AbsNoteContext) EQ() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserEQ, 0)
}

func (s *AbsNoteContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(RasterklangDslParserNUMBER)
}

func (s *AbsNoteContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNUMBER, i)
}

func (s *AbsNoteContext) BANG() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserBANG, 0)
}

func (s *AbsNoteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AbsNoteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AbsNoteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterAbsNote(s)
	}
}

func (s *AbsNoteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitAbsNote(s)
	}
}

func (s *AbsNoteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitAbsNote(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) AbsNote() (localctx IAbsNoteContext) {
	localctx = NewAbsNoteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, RasterklangDslParserRULE_absNote)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(269)
		p.Match(RasterklangDslParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(270)
		p.Match(RasterklangDslParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == RasterklangDslParserBANG {
		{
			p.SetState(271)
			p.Match(RasterklangDslParserBANG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(272)
			p.Match(RasterklangDslParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRampGeneratorContext is an interface to support dynamic dispatch.
type IRampGeneratorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAnyIdent() []IAnyIdentContext
	AnyIdent(i int) IAnyIdentContext
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode
	DOTDOT() antlr.TerminalNode

	// IsRampGeneratorContext differentiates from other interfaces.
	IsRampGeneratorContext()
}

type RampGeneratorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRampGeneratorContext() *RampGeneratorContext {
	var p = new(RampGeneratorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_rampGenerator
	return p
}

func InitEmptyRampGeneratorContext(p *RampGeneratorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_rampGenerator
}

func (*RampGeneratorContext) IsRampGeneratorContext() {}

func NewRampGeneratorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RampGeneratorContext {
	var p = new(RampGeneratorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_rampGenerator

	return p
}

func (s *RampGeneratorContext) GetParser() antlr.Parser { return s.parser }

func (s *RampGeneratorContext) AllAnyIdent() []IAnyIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAnyIdentContext); ok {
			len++
		}
	}

	tst := make([]IAnyIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAnyIdentContext); ok {
			tst[i] = t.(IAnyIdentContext)
			i++
		}
	}

	return tst
}

func (s *RampGeneratorContext) AnyIdent(i int) IAnyIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyIdentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyIdentContext)
}

func (s *RampGeneratorContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(RasterklangDslParserNUMBER)
}

func (s *RampGeneratorContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNUMBER, i)
}

func (s *RampGeneratorContext) DOTDOT() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserDOTDOT, 0)
}

func (s *RampGeneratorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RampGeneratorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RampGeneratorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterRampGenerator(s)
	}
}

func (s *RampGeneratorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitRampGenerator(s)
	}
}

func (s *RampGeneratorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitRampGenerator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) RampGenerator() (localctx IRampGeneratorContext) {
	localctx = NewRampGeneratorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, RasterklangDslParserRULE_rampGenerator)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(275)
		p.AnyIdent()
	}
	{
		p.SetState(276)
		p.Match(RasterklangDslParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(277)
		p.Match(RasterklangDslParserDOTDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(278)
		p.Match(RasterklangDslParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(279)
			p.AnyIdent()
		}
		{
			p.SetState(280)
			p.Match(RasterklangDslParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(287)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(284)
			p.AnyIdent()
		}
		{
			p.SetState(285)
			p.Match(RasterklangDslParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(290)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(289)
			p.AnyIdent()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVoiceStmtContext is an interface to support dynamic dispatch.
type IVoiceStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VOICE() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	VoiceBody() IVoiceBodyContext
	AnyIdent() IAnyIdentContext

	// IsVoiceStmtContext differentiates from other interfaces.
	IsVoiceStmtContext()
}

type VoiceStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVoiceStmtContext() *VoiceStmtContext {
	var p = new(VoiceStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_voiceStmt
	return p
}

func InitEmptyVoiceStmtContext(p *VoiceStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_voiceStmt
}

func (*VoiceStmtContext) IsVoiceStmtContext() {}

func NewVoiceStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VoiceStmtContext {
	var p = new(VoiceStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_voiceStmt

	return p
}

func (s *VoiceStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *VoiceStmtContext) VOICE() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserVOICE, 0)
}

func (s *VoiceStmtContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNUMBER, 0)
}

func (s *VoiceStmtContext) VoiceBody() IVoiceBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVoiceBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVoiceBodyContext)
}

func (s *VoiceStmtContext) AnyIdent() IAnyIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyIdentContext)
}

func (s *VoiceStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VoiceStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VoiceStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterVoiceStmt(s)
	}
}

func (s *VoiceStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitVoiceStmt(s)
	}
}

func (s *VoiceStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitVoiceStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) VoiceStmt() (localctx IVoiceStmtContext) {
	localctx = NewVoiceStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, RasterklangDslParserRULE_voiceStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(292)
		p.Match(RasterklangDslParserVOICE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(293)
		p.Match(RasterklangDslParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&66558) != 0 {
		{
			p.SetState(294)
			p.AnyIdent()
		}

	}
	{
		p.SetState(297)
		p.VoiceBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVoiceBodyContext is an interface to support dynamic dispatch.
type IVoiceBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VoiceShortTail() IVoiceShortTailContext
	Str() IStrContext
	SIGIL() antlr.TerminalNode
	Arrangement() IArrangementContext
	NEWLINE() antlr.TerminalNode
	END() antlr.TerminalNode
	AllVoiceField() []IVoiceFieldContext
	VoiceField(i int) IVoiceFieldContext

	// IsVoiceBodyContext differentiates from other interfaces.
	IsVoiceBodyContext()
}

type VoiceBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVoiceBodyContext() *VoiceBodyContext {
	var p = new(VoiceBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_voiceBody
	return p
}

func InitEmptyVoiceBodyContext(p *VoiceBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_voiceBody
}

func (*VoiceBodyContext) IsVoiceBodyContext() {}

func NewVoiceBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VoiceBodyContext {
	var p = new(VoiceBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_voiceBody

	return p
}

func (s *VoiceBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *VoiceBodyContext) VoiceShortTail() IVoiceShortTailContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVoiceShortTailContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVoiceShortTailContext)
}

func (s *VoiceBodyContext) Str() IStrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrContext)
}

func (s *VoiceBodyContext) SIGIL() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserSIGIL, 0)
}

func (s *VoiceBodyContext) Arrangement() IArrangementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrangementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrangementContext)
}

func (s *VoiceBodyContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNEWLINE, 0)
}

func (s *VoiceBodyContext) END() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserEND, 0)
}

func (s *VoiceBodyContext) AllVoiceField() []IVoiceFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVoiceFieldContext); ok {
			len++
		}
	}

	tst := make([]IVoiceFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVoiceFieldContext); ok {
			tst[i] = t.(IVoiceFieldContext)
			i++
		}
	}

	return tst
}

func (s *VoiceBodyContext) VoiceField(i int) IVoiceFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVoiceFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVoiceFieldContext)
}

func (s *VoiceBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VoiceBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VoiceBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterVoiceBody(s)
	}
}

func (s *VoiceBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitVoiceBody(s)
	}
}

func (s *VoiceBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitVoiceBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) VoiceBody() (localctx IVoiceBodyContext) {
	localctx = NewVoiceBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, RasterklangDslParserRULE_voiceBody)
	var _la int

	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(301)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case RasterklangDslParserDQSTRING, RasterklangDslParserBTSTRING:
			{
				p.SetState(299)
				p.Str()
			}

		case RasterklangDslParserSIGIL:
			{
				p.SetState(300)
				p.Match(RasterklangDslParserSIGIL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(303)
			p.VoiceShortTail()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(304)
			p.Arrangement()
		}
		{
			p.SetState(305)
			p.VoiceShortTail()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(307)
			p.Match(RasterklangDslParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&197630) != 0 {
			{
				p.SetState(308)
				p.VoiceField()
			}

			p.SetState(313)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(314)
			p.Match(RasterklangDslParserEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVoiceShortTailContext is an interface to support dynamic dispatch.
type IVoiceShortTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllInlineAttr() []IInlineAttrContext
	InlineAttr(i int) IInlineAttrContext
	AllPipelineStep() []IPipelineStepContext
	PipelineStep(i int) IPipelineStepContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsVoiceShortTailContext differentiates from other interfaces.
	IsVoiceShortTailContext()
}

type VoiceShortTailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVoiceShortTailContext() *VoiceShortTailContext {
	var p = new(VoiceShortTailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_voiceShortTail
	return p
}

func InitEmptyVoiceShortTailContext(p *VoiceShortTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_voiceShortTail
}

func (*VoiceShortTailContext) IsVoiceShortTailContext() {}

func NewVoiceShortTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VoiceShortTailContext {
	var p = new(VoiceShortTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_voiceShortTail

	return p
}

func (s *VoiceShortTailContext) GetParser() antlr.Parser { return s.parser }

func (s *VoiceShortTailContext) AllInlineAttr() []IInlineAttrContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInlineAttrContext); ok {
			len++
		}
	}

	tst := make([]IInlineAttrContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInlineAttrContext); ok {
			tst[i] = t.(IInlineAttrContext)
			i++
		}
	}

	return tst
}

func (s *VoiceShortTailContext) InlineAttr(i int) IInlineAttrContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlineAttrContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInlineAttrContext)
}

func (s *VoiceShortTailContext) AllPipelineStep() []IPipelineStepContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPipelineStepContext); ok {
			len++
		}
	}

	tst := make([]IPipelineStepContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPipelineStepContext); ok {
			tst[i] = t.(IPipelineStepContext)
			i++
		}
	}

	return tst
}

func (s *VoiceShortTailContext) PipelineStep(i int) IPipelineStepContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipelineStepContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipelineStepContext)
}

func (s *VoiceShortTailContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(RasterklangDslParserNEWLINE)
}

func (s *VoiceShortTailContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNEWLINE, i)
}

func (s *VoiceShortTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VoiceShortTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VoiceShortTailContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterVoiceShortTail(s)
	}
}

func (s *VoiceShortTailContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitVoiceShortTail(s)
	}
}

func (s *VoiceShortTailContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitVoiceShortTail(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) VoiceShortTail() (localctx IVoiceShortTailContext) {
	localctx = NewVoiceShortTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, RasterklangDslParserRULE_voiceShortTail)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(318)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == RasterklangDslParserNEWLINE {
				{
					p.SetState(317)
					p.Match(RasterklangDslParserNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(320)
				p.InlineAttr()
			}

		}
		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(327)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == RasterklangDslParserNEWLINE {
				{
					p.SetState(326)
					p.Match(RasterklangDslParserNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(329)
				p.PipelineStep()
			}

		}
		p.SetState(334)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInlineAttrContext is an interface to support dynamic dispatch.
type IInlineAttrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AnyIdent() IAnyIdentContext
	COLON() antlr.TerminalNode
	Str() IStrContext
	ValueList() IValueListContext

	// IsInlineAttrContext differentiates from other interfaces.
	IsInlineAttrContext()
}

type InlineAttrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineAttrContext() *InlineAttrContext {
	var p = new(InlineAttrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_inlineAttr
	return p
}

func InitEmptyInlineAttrContext(p *InlineAttrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_inlineAttr
}

func (*InlineAttrContext) IsInlineAttrContext() {}

func NewInlineAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineAttrContext {
	var p = new(InlineAttrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_inlineAttr

	return p
}

func (s *InlineAttrContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineAttrContext) AnyIdent() IAnyIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyIdentContext)
}

func (s *InlineAttrContext) COLON() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserCOLON, 0)
}

func (s *InlineAttrContext) Str() IStrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrContext)
}

func (s *InlineAttrContext) ValueList() IValueListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueListContext)
}

func (s *InlineAttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineAttrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InlineAttrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterInlineAttr(s)
	}
}

func (s *InlineAttrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitInlineAttr(s)
	}
}

func (s *InlineAttrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitInlineAttr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) InlineAttr() (localctx IInlineAttrContext) {
	localctx = NewInlineAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, RasterklangDslParserRULE_inlineAttr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(335)
		p.AnyIdent()
	}
	{
		p.SetState(336)
		p.Match(RasterklangDslParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(339)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RasterklangDslParserDQSTRING, RasterklangDslParserBTSTRING:
		{
			p.SetState(337)
			p.Str()
		}

	case RasterklangDslParserTEMPO, RasterklangDslParserSPEED, RasterklangDslParserVOICE, RasterklangDslParserPATTERN, RasterklangDslParserTABLE, RasterklangDslParserINSTRUMENT, RasterklangDslParserFILTER, RasterklangDslParserTARGET, RasterklangDslParserCHIP, RasterklangDslParserNUMBER, RasterklangDslParserSIGIL, RasterklangDslParserIDENT, RasterklangDslParserEQ:
		{
			p.SetState(338)
			p.ValueList()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueListContext is an interface to support dynamic dispatch.
type IValueListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTableToken() []ITableTokenContext
	TableToken(i int) ITableTokenContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsValueListContext differentiates from other interfaces.
	IsValueListContext()
}

type ValueListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueListContext() *ValueListContext {
	var p = new(ValueListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_valueList
	return p
}

func InitEmptyValueListContext(p *ValueListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_valueList
}

func (*ValueListContext) IsValueListContext() {}

func NewValueListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueListContext {
	var p = new(ValueListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_valueList

	return p
}

func (s *ValueListContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueListContext) AllTableToken() []ITableTokenContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITableTokenContext); ok {
			len++
		}
	}

	tst := make([]ITableTokenContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITableTokenContext); ok {
			tst[i] = t.(ITableTokenContext)
			i++
		}
	}

	return tst
}

func (s *ValueListContext) TableToken(i int) ITableTokenContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableTokenContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableTokenContext)
}

func (s *ValueListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(RasterklangDslParserCOMMA)
}

func (s *ValueListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserCOMMA, i)
}

func (s *ValueListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterValueList(s)
	}
}

func (s *ValueListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitValueList(s)
	}
}

func (s *ValueListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitValueList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) ValueList() (localctx IValueListContext) {
	localctx = NewValueListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, RasterklangDslParserRULE_valueList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(341)
		p.TableToken()
	}
	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == RasterklangDslParserCOMMA {
		{
			p.SetState(342)
			p.Match(RasterklangDslParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)
			p.TableToken()
		}

		p.SetState(348)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVoiceFieldContext is an interface to support dynamic dispatch.
type IVoiceFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NEWLINE() antlr.TerminalNode
	PipelineStep() IPipelineStepContext
	AnyIdent() IAnyIdentContext
	VoiceFieldValue() IVoiceFieldValueContext

	// IsVoiceFieldContext differentiates from other interfaces.
	IsVoiceFieldContext()
}

type VoiceFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVoiceFieldContext() *VoiceFieldContext {
	var p = new(VoiceFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_voiceField
	return p
}

func InitEmptyVoiceFieldContext(p *VoiceFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_voiceField
}

func (*VoiceFieldContext) IsVoiceFieldContext() {}

func NewVoiceFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VoiceFieldContext {
	var p = new(VoiceFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_voiceField

	return p
}

func (s *VoiceFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *VoiceFieldContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNEWLINE, 0)
}

func (s *VoiceFieldContext) PipelineStep() IPipelineStepContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipelineStepContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipelineStepContext)
}

func (s *VoiceFieldContext) AnyIdent() IAnyIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyIdentContext)
}

func (s *VoiceFieldContext) VoiceFieldValue() IVoiceFieldValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVoiceFieldValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVoiceFieldValueContext)
}

func (s *VoiceFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VoiceFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VoiceFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterVoiceField(s)
	}
}

func (s *VoiceFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitVoiceField(s)
	}
}

func (s *VoiceFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitVoiceField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) VoiceField() (localctx IVoiceFieldContext) {
	localctx = NewVoiceFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, RasterklangDslParserRULE_voiceField)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RasterklangDslParserPIPE:
		{
			p.SetState(349)
			p.PipelineStep()
		}

	case RasterklangDslParserTEMPO, RasterklangDslParserSPEED, RasterklangDslParserVOICE, RasterklangDslParserPATTERN, RasterklangDslParserTABLE, RasterklangDslParserINSTRUMENT, RasterklangDslParserFILTER, RasterklangDslParserTARGET, RasterklangDslParserCHIP, RasterklangDslParserIDENT:
		{
			p.SetState(350)
			p.AnyIdent()
		}
		{
			p.SetState(351)
			p.VoiceFieldValue()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(355)
		p.Match(RasterklangDslParserNEWLINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVoiceFieldValueContext is an interface to support dynamic dispatch.
type IVoiceFieldValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Str() IStrContext
	SIGIL() antlr.TerminalNode
	AnyIdent() IAnyIdentContext
	HEX() antlr.TerminalNode
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode

	// IsVoiceFieldValueContext differentiates from other interfaces.
	IsVoiceFieldValueContext()
}

type VoiceFieldValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVoiceFieldValueContext() *VoiceFieldValueContext {
	var p = new(VoiceFieldValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_voiceFieldValue
	return p
}

func InitEmptyVoiceFieldValueContext(p *VoiceFieldValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_voiceFieldValue
}

func (*VoiceFieldValueContext) IsVoiceFieldValueContext() {}

func NewVoiceFieldValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VoiceFieldValueContext {
	var p = new(VoiceFieldValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_voiceFieldValue

	return p
}

func (s *VoiceFieldValueContext) GetParser() antlr.Parser { return s.parser }

func (s *VoiceFieldValueContext) Str() IStrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrContext)
}

func (s *VoiceFieldValueContext) SIGIL() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserSIGIL, 0)
}

func (s *VoiceFieldValueContext) AnyIdent() IAnyIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyIdentContext)
}

func (s *VoiceFieldValueContext) HEX() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserHEX, 0)
}

func (s *VoiceFieldValueContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(RasterklangDslParserNUMBER)
}

func (s *VoiceFieldValueContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNUMBER, i)
}

func (s *VoiceFieldValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VoiceFieldValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VoiceFieldValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterVoiceFieldValue(s)
	}
}

func (s *VoiceFieldValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitVoiceFieldValue(s)
	}
}

func (s *VoiceFieldValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitVoiceFieldValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) VoiceFieldValue() (localctx IVoiceFieldValueContext) {
	localctx = NewVoiceFieldValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, RasterklangDslParserRULE_voiceFieldValue)
	var _la int

	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RasterklangDslParserDQSTRING, RasterklangDslParserBTSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(357)
			p.Str()
		}

	case RasterklangDslParserSIGIL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(358)
			p.Match(RasterklangDslParserSIGIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RasterklangDslParserTEMPO, RasterklangDslParserSPEED, RasterklangDslParserVOICE, RasterklangDslParserPATTERN, RasterklangDslParserTABLE, RasterklangDslParserINSTRUMENT, RasterklangDslParserFILTER, RasterklangDslParserTARGET, RasterklangDslParserCHIP, RasterklangDslParserHEX, RasterklangDslParserIDENT:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(361)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case RasterklangDslParserTEMPO, RasterklangDslParserSPEED, RasterklangDslParserVOICE, RasterklangDslParserPATTERN, RasterklangDslParserTABLE, RasterklangDslParserINSTRUMENT, RasterklangDslParserFILTER, RasterklangDslParserTARGET, RasterklangDslParserCHIP, RasterklangDslParserIDENT:
			{
				p.SetState(359)
				p.AnyIdent()
			}

		case RasterklangDslParserHEX:
			{
				p.SetState(360)
				p.Match(RasterklangDslParserHEX)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(363)
			p.Str()
		}

	case RasterklangDslParserNUMBER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(364)
			p.Match(RasterklangDslParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(366)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == RasterklangDslParserNUMBER {
			{
				p.SetState(365)
				p.Match(RasterklangDslParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPipelineStepContext is an interface to support dynamic dispatch.
type IPipelineStepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PIPE() antlr.TerminalNode
	Transform() ITransformContext

	// IsPipelineStepContext differentiates from other interfaces.
	IsPipelineStepContext()
}

type PipelineStepContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipelineStepContext() *PipelineStepContext {
	var p = new(PipelineStepContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_pipelineStep
	return p
}

func InitEmptyPipelineStepContext(p *PipelineStepContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_pipelineStep
}

func (*PipelineStepContext) IsPipelineStepContext() {}

func NewPipelineStepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipelineStepContext {
	var p = new(PipelineStepContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_pipelineStep

	return p
}

func (s *PipelineStepContext) GetParser() antlr.Parser { return s.parser }

func (s *PipelineStepContext) PIPE() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserPIPE, 0)
}

func (s *PipelineStepContext) Transform() ITransformContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITransformContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITransformContext)
}

func (s *PipelineStepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipelineStepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipelineStepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterPipelineStep(s)
	}
}

func (s *PipelineStepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitPipelineStep(s)
	}
}

func (s *PipelineStepContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitPipelineStep(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) PipelineStep() (localctx IPipelineStepContext) {
	localctx = NewPipelineStepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, RasterklangDslParserRULE_pipelineStep)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		p.Match(RasterklangDslParserPIPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(371)
		p.Transform()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITransformContext is an interface to support dynamic dispatch.
type ITransformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AnyIdent() IAnyIdentContext
	TransformArg() ITransformArgContext
	TransformRef() ITransformRefContext

	// IsTransformContext differentiates from other interfaces.
	IsTransformContext()
}

type TransformContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransformContext() *TransformContext {
	var p = new(TransformContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_transform
	return p
}

func InitEmptyTransformContext(p *TransformContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_transform
}

func (*TransformContext) IsTransformContext() {}

func NewTransformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransformContext {
	var p = new(TransformContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_transform

	return p
}

func (s *TransformContext) GetParser() antlr.Parser { return s.parser }

func (s *TransformContext) AnyIdent() IAnyIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyIdentContext)
}

func (s *TransformContext) TransformArg() ITransformArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITransformArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITransformArgContext)
}

func (s *TransformContext) TransformRef() ITransformRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITransformRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITransformRefContext)
}

func (s *TransformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TransformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterTransform(s)
	}
}

func (s *TransformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitTransform(s)
	}
}

func (s *TransformContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitTransform(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) Transform() (localctx ITransformContext) {
	localctx = NewTransformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, RasterklangDslParserRULE_transform)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(373)
		p.AnyIdent()
	}
	p.SetState(375)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&28672) != 0 {
		{
			p.SetState(374)
			p.TransformArg()
		}

	}
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67175422) != 0 {
		{
			p.SetState(377)
			p.TransformRef()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITransformArgContext is an interface to support dynamic dispatch.
type ITransformArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	Str() IStrContext

	// IsTransformArgContext differentiates from other interfaces.
	IsTransformArgContext()
}

type TransformArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransformArgContext() *TransformArgContext {
	var p = new(TransformArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_transformArg
	return p
}

func InitEmptyTransformArgContext(p *TransformArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_transformArg
}

func (*TransformArgContext) IsTransformArgContext() {}

func NewTransformArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransformArgContext {
	var p = new(TransformArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_transformArg

	return p
}

func (s *TransformArgContext) GetParser() antlr.Parser { return s.parser }

func (s *TransformArgContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNUMBER, 0)
}

func (s *TransformArgContext) Str() IStrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrContext)
}

func (s *TransformArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransformArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TransformArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterTransformArg(s)
	}
}

func (s *TransformArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitTransformArg(s)
	}
}

func (s *TransformArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitTransformArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) TransformArg() (localctx ITransformArgContext) {
	localctx = NewTransformArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, RasterklangDslParserRULE_transformArg)
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RasterklangDslParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(380)
			p.Match(RasterklangDslParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RasterklangDslParserDQSTRING, RasterklangDslParserBTSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(381)
			p.Str()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITransformRefContext is an interface to support dynamic dispatch.
type ITransformRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Transform() ITransformContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsTransformRefContext differentiates from other interfaces.
	IsTransformRefContext()
}

type TransformRefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransformRefContext() *TransformRefContext {
	var p = new(TransformRefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_transformRef
	return p
}

func InitEmptyTransformRefContext(p *TransformRefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_transformRef
}

func (*TransformRefContext) IsTransformRefContext() {}

func NewTransformRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransformRefContext {
	var p = new(TransformRefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_transformRef

	return p
}

func (s *TransformRefContext) GetParser() antlr.Parser { return s.parser }

func (s *TransformRefContext) Transform() ITransformContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITransformContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITransformContext)
}

func (s *TransformRefContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserLPAREN, 0)
}

func (s *TransformRefContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserRPAREN, 0)
}

func (s *TransformRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransformRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TransformRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterTransformRef(s)
	}
}

func (s *TransformRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitTransformRef(s)
	}
}

func (s *TransformRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitTransformRef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) TransformRef() (localctx ITransformRefContext) {
	localctx = NewTransformRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, RasterklangDslParserRULE_transformRef)
	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RasterklangDslParserTEMPO, RasterklangDslParserSPEED, RasterklangDslParserVOICE, RasterklangDslParserPATTERN, RasterklangDslParserTABLE, RasterklangDslParserINSTRUMENT, RasterklangDslParserFILTER, RasterklangDslParserTARGET, RasterklangDslParserCHIP, RasterklangDslParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(384)
			p.Transform()
		}

	case RasterklangDslParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(385)
			p.Match(RasterklangDslParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(386)
			p.Transform()
		}
		{
			p.SetState(387)
			p.Match(RasterklangDslParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFilterBlockContext is an interface to support dynamic dispatch.
type IFilterBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FILTER() antlr.TerminalNode
	NEWLINE() antlr.TerminalNode
	END() antlr.TerminalNode
	AllFilterField() []IFilterFieldContext
	FilterField(i int) IFilterFieldContext

	// IsFilterBlockContext differentiates from other interfaces.
	IsFilterBlockContext()
}

type FilterBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterBlockContext() *FilterBlockContext {
	var p = new(FilterBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_filterBlock
	return p
}

func InitEmptyFilterBlockContext(p *FilterBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_filterBlock
}

func (*FilterBlockContext) IsFilterBlockContext() {}

func NewFilterBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterBlockContext {
	var p = new(FilterBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_filterBlock

	return p
}

func (s *FilterBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterBlockContext) FILTER() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserFILTER, 0)
}

func (s *FilterBlockContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNEWLINE, 0)
}

func (s *FilterBlockContext) END() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserEND, 0)
}

func (s *FilterBlockContext) AllFilterField() []IFilterFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFilterFieldContext); ok {
			len++
		}
	}

	tst := make([]IFilterFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFilterFieldContext); ok {
			tst[i] = t.(IFilterFieldContext)
			i++
		}
	}

	return tst
}

func (s *FilterBlockContext) FilterField(i int) IFilterFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterFieldContext)
}

func (s *FilterBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterFilterBlock(s)
	}
}

func (s *FilterBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitFilterBlock(s)
	}
}

func (s *FilterBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitFilterBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) FilterBlock() (localctx IFilterBlockContext) {
	localctx = NewFilterBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, RasterklangDslParserRULE_filterBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(391)
		p.Match(RasterklangDslParserFILTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(392)
		p.Match(RasterklangDslParserNEWLINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&66558) != 0 {
		{
			p.SetState(393)
			p.FilterField()
		}

		p.SetState(398)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(399)
		p.Match(RasterklangDslParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFilterFieldContext is an interface to support dynamic dispatch.
type IFilterFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AnyIdent() IAnyIdentContext
	FilterFieldValue() IFilterFieldValueContext
	NEWLINE() antlr.TerminalNode

	// IsFilterFieldContext differentiates from other interfaces.
	IsFilterFieldContext()
}

type FilterFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterFieldContext() *FilterFieldContext {
	var p = new(FilterFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_filterField
	return p
}

func InitEmptyFilterFieldContext(p *FilterFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_filterField
}

func (*FilterFieldContext) IsFilterFieldContext() {}

func NewFilterFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterFieldContext {
	var p = new(FilterFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_filterField

	return p
}

func (s *FilterFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterFieldContext) AnyIdent() IAnyIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyIdentContext)
}

func (s *FilterFieldContext) FilterFieldValue() IFilterFieldValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterFieldValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterFieldValueContext)
}

func (s *FilterFieldContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNEWLINE, 0)
}

func (s *FilterFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterFilterField(s)
	}
}

func (s *FilterFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitFilterField(s)
	}
}

func (s *FilterFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitFilterField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) FilterField() (localctx IFilterFieldContext) {
	localctx = NewFilterFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, RasterklangDslParserRULE_filterField)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(401)
		p.AnyIdent()
	}
	{
		p.SetState(402)
		p.FilterFieldValue()
	}
	{
		p.SetState(403)
		p.Match(RasterklangDslParserNEWLINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFilterFieldValueContext is an interface to support dynamic dispatch.
type IFilterFieldValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAnyIdent() []IAnyIdentContext
	AnyIdent(i int) IAnyIdentContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	AllSwitchCase() []ISwitchCaseContext
	SwitchCase(i int) ISwitchCaseContext
	SignalRangeOpt() ISignalRangeOptContext
	AllSignalSpeedMod() []ISignalSpeedModContext
	SignalSpeedMod(i int) ISignalSpeedModContext
	NumberList() INumberListContext

	// IsFilterFieldValueContext differentiates from other interfaces.
	IsFilterFieldValueContext()
}

type FilterFieldValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterFieldValueContext() *FilterFieldValueContext {
	var p = new(FilterFieldValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_filterFieldValue
	return p
}

func InitEmptyFilterFieldValueContext(p *FilterFieldValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_filterFieldValue
}

func (*FilterFieldValueContext) IsFilterFieldValueContext() {}

func NewFilterFieldValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterFieldValueContext {
	var p = new(FilterFieldValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_filterFieldValue

	return p
}

func (s *FilterFieldValueContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterFieldValueContext) AllAnyIdent() []IAnyIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAnyIdentContext); ok {
			len++
		}
	}

	tst := make([]IAnyIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAnyIdentContext); ok {
			tst[i] = t.(IAnyIdentContext)
			i++
		}
	}

	return tst
}

func (s *FilterFieldValueContext) AnyIdent(i int) IAnyIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyIdentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyIdentContext)
}

func (s *FilterFieldValueContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(RasterklangDslParserNEWLINE)
}

func (s *FilterFieldValueContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNEWLINE, i)
}

func (s *FilterFieldValueContext) AllSwitchCase() []ISwitchCaseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISwitchCaseContext); ok {
			len++
		}
	}

	tst := make([]ISwitchCaseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISwitchCaseContext); ok {
			tst[i] = t.(ISwitchCaseContext)
			i++
		}
	}

	return tst
}

func (s *FilterFieldValueContext) SwitchCase(i int) ISwitchCaseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchCaseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchCaseContext)
}

func (s *FilterFieldValueContext) SignalRangeOpt() ISignalRangeOptContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISignalRangeOptContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISignalRangeOptContext)
}

func (s *FilterFieldValueContext) AllSignalSpeedMod() []ISignalSpeedModContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISignalSpeedModContext); ok {
			len++
		}
	}

	tst := make([]ISignalSpeedModContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISignalSpeedModContext); ok {
			tst[i] = t.(ISignalSpeedModContext)
			i++
		}
	}

	return tst
}

func (s *FilterFieldValueContext) SignalSpeedMod(i int) ISignalSpeedModContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISignalSpeedModContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISignalSpeedModContext)
}

func (s *FilterFieldValueContext) NumberList() INumberListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberListContext)
}

func (s *FilterFieldValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterFieldValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterFieldValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterFilterFieldValue(s)
	}
}

func (s *FilterFieldValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitFilterFieldValue(s)
	}
}

func (s *FilterFieldValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitFilterFieldValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) FilterFieldValue() (localctx IFilterFieldValueContext) {
	localctx = NewFilterFieldValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, RasterklangDslParserRULE_filterFieldValue)
	var _la int

	var _alt int

	p.SetState(425)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(405)
			p.AnyIdent()
		}
		{
			p.SetState(406)
			p.AnyIdent()
		}
		{
			p.SetState(407)
			p.Match(RasterklangDslParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(408)
			p.SwitchCase()
		}
		p.SetState(413)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(409)
					p.Match(RasterklangDslParserNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(410)
					p.SwitchCase()
				}

			}
			p.SetState(415)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(416)
			p.AnyIdent()
		}
		{
			p.SetState(417)
			p.SignalRangeOpt()
		}
		p.SetState(421)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&25690112) != 0 {
			{
				p.SetState(418)
				p.SignalSpeedMod()
			}

			p.SetState(423)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(424)
			p.NumberList()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchCaseContext is an interface to support dynamic dispatch.
type ISwitchCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SwitchCond() ISwitchCondContext
	COLON() antlr.TerminalNode
	AnyIdent() IAnyIdentContext
	SignalRangeOpt() ISignalRangeOptContext
	NumberList() INumberListContext
	AllSignalSpeedMod() []ISignalSpeedModContext
	SignalSpeedMod(i int) ISignalSpeedModContext

	// IsSwitchCaseContext differentiates from other interfaces.
	IsSwitchCaseContext()
}

type SwitchCaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchCaseContext() *SwitchCaseContext {
	var p = new(SwitchCaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_switchCase
	return p
}

func InitEmptySwitchCaseContext(p *SwitchCaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_switchCase
}

func (*SwitchCaseContext) IsSwitchCaseContext() {}

func NewSwitchCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchCaseContext {
	var p = new(SwitchCaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_switchCase

	return p
}

func (s *SwitchCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchCaseContext) SwitchCond() ISwitchCondContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchCondContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchCondContext)
}

func (s *SwitchCaseContext) COLON() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserCOLON, 0)
}

func (s *SwitchCaseContext) AnyIdent() IAnyIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyIdentContext)
}

func (s *SwitchCaseContext) SignalRangeOpt() ISignalRangeOptContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISignalRangeOptContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISignalRangeOptContext)
}

func (s *SwitchCaseContext) NumberList() INumberListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberListContext)
}

func (s *SwitchCaseContext) AllSignalSpeedMod() []ISignalSpeedModContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISignalSpeedModContext); ok {
			len++
		}
	}

	tst := make([]ISignalSpeedModContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISignalSpeedModContext); ok {
			tst[i] = t.(ISignalSpeedModContext)
			i++
		}
	}

	return tst
}

func (s *SwitchCaseContext) SignalSpeedMod(i int) ISignalSpeedModContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISignalSpeedModContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISignalSpeedModContext)
}

func (s *SwitchCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterSwitchCase(s)
	}
}

func (s *SwitchCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitSwitchCase(s)
	}
}

func (s *SwitchCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitSwitchCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) SwitchCase() (localctx ISwitchCaseContext) {
	localctx = NewSwitchCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, RasterklangDslParserRULE_switchCase)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(427)
		p.SwitchCond()
	}
	{
		p.SetState(428)
		p.Match(RasterklangDslParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RasterklangDslParserTEMPO, RasterklangDslParserSPEED, RasterklangDslParserVOICE, RasterklangDslParserPATTERN, RasterklangDslParserTABLE, RasterklangDslParserINSTRUMENT, RasterklangDslParserFILTER, RasterklangDslParserTARGET, RasterklangDslParserCHIP, RasterklangDslParserIDENT:
		{
			p.SetState(429)
			p.AnyIdent()
		}
		{
			p.SetState(430)
			p.SignalRangeOpt()
		}
		p.SetState(434)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&25690112) != 0 {
			{
				p.SetState(431)
				p.SignalSpeedMod()
			}

			p.SetState(436)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case RasterklangDslParserNUMBER:
		{
			p.SetState(437)
			p.NumberList()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchCondContext is an interface to support dynamic dispatch.
type ISwitchCondContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LT() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	GE() antlr.TerminalNode
	AnyIdent() IAnyIdentContext

	// IsSwitchCondContext differentiates from other interfaces.
	IsSwitchCondContext()
}

type SwitchCondContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchCondContext() *SwitchCondContext {
	var p = new(SwitchCondContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_switchCond
	return p
}

func InitEmptySwitchCondContext(p *SwitchCondContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_switchCond
}

func (*SwitchCondContext) IsSwitchCondContext() {}

func NewSwitchCondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchCondContext {
	var p = new(SwitchCondContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_switchCond

	return p
}

func (s *SwitchCondContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchCondContext) LT() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserLT, 0)
}

func (s *SwitchCondContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNUMBER, 0)
}

func (s *SwitchCondContext) GE() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserGE, 0)
}

func (s *SwitchCondContext) AnyIdent() IAnyIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyIdentContext)
}

func (s *SwitchCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchCondContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchCondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterSwitchCond(s)
	}
}

func (s *SwitchCondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitSwitchCond(s)
	}
}

func (s *SwitchCondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitSwitchCond(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) SwitchCond() (localctx ISwitchCondContext) {
	localctx = NewSwitchCondContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, RasterklangDslParserRULE_switchCond)
	p.SetState(445)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RasterklangDslParserLT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(440)
			p.Match(RasterklangDslParserLT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(441)
			p.Match(RasterklangDslParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RasterklangDslParserGE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(442)
			p.Match(RasterklangDslParserGE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(443)
			p.Match(RasterklangDslParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RasterklangDslParserTEMPO, RasterklangDslParserSPEED, RasterklangDslParserVOICE, RasterklangDslParserPATTERN, RasterklangDslParserTABLE, RasterklangDslParserINSTRUMENT, RasterklangDslParserFILTER, RasterklangDslParserTARGET, RasterklangDslParserCHIP, RasterklangDslParserIDENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(444)
			p.AnyIdent()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISignalRangeOptContext is an interface to support dynamic dispatch.
type ISignalRangeOptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode
	DOTDOT() antlr.TerminalNode

	// IsSignalRangeOptContext differentiates from other interfaces.
	IsSignalRangeOptContext()
}

type SignalRangeOptContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySignalRangeOptContext() *SignalRangeOptContext {
	var p = new(SignalRangeOptContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_signalRangeOpt
	return p
}

func InitEmptySignalRangeOptContext(p *SignalRangeOptContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_signalRangeOpt
}

func (*SignalRangeOptContext) IsSignalRangeOptContext() {}

func NewSignalRangeOptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SignalRangeOptContext {
	var p = new(SignalRangeOptContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_signalRangeOpt

	return p
}

func (s *SignalRangeOptContext) GetParser() antlr.Parser { return s.parser }

func (s *SignalRangeOptContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(RasterklangDslParserNUMBER)
}

func (s *SignalRangeOptContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNUMBER, i)
}

func (s *SignalRangeOptContext) DOTDOT() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserDOTDOT, 0)
}

func (s *SignalRangeOptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SignalRangeOptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SignalRangeOptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterSignalRangeOpt(s)
	}
}

func (s *SignalRangeOptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitSignalRangeOpt(s)
	}
}

func (s *SignalRangeOptContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitSignalRangeOpt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) SignalRangeOpt() (localctx ISignalRangeOptContext) {
	localctx = NewSignalRangeOptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, RasterklangDslParserRULE_signalRangeOpt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(450)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == RasterklangDslParserNUMBER {
		{
			p.SetState(447)
			p.Match(RasterklangDslParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(448)
			p.Match(RasterklangDslParserDOTDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(449)
			p.Match(RasterklangDslParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISignalSpeedModContext is an interface to support dynamic dispatch.
type ISignalSpeedModContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	STAR() antlr.TerminalNode
	SLASH() antlr.TerminalNode
	SHIFT() antlr.TerminalNode

	// IsSignalSpeedModContext differentiates from other interfaces.
	IsSignalSpeedModContext()
}

type SignalSpeedModContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySignalSpeedModContext() *SignalSpeedModContext {
	var p = new(SignalSpeedModContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_signalSpeedMod
	return p
}

func InitEmptySignalSpeedModContext(p *SignalSpeedModContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_signalSpeedMod
}

func (*SignalSpeedModContext) IsSignalSpeedModContext() {}

func NewSignalSpeedModContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SignalSpeedModContext {
	var p = new(SignalSpeedModContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_signalSpeedMod

	return p
}

func (s *SignalSpeedModContext) GetParser() antlr.Parser { return s.parser }

func (s *SignalSpeedModContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNUMBER, 0)
}

func (s *SignalSpeedModContext) STAR() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserSTAR, 0)
}

func (s *SignalSpeedModContext) SLASH() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserSLASH, 0)
}

func (s *SignalSpeedModContext) SHIFT() antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserSHIFT, 0)
}

func (s *SignalSpeedModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SignalSpeedModContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SignalSpeedModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterSignalSpeedMod(s)
	}
}

func (s *SignalSpeedModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitSignalSpeedMod(s)
	}
}

func (s *SignalSpeedModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitSignalSpeedMod(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) SignalSpeedMod() (localctx ISignalSpeedModContext) {
	localctx = NewSignalSpeedModContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, RasterklangDslParserRULE_signalSpeedMod)
	var _la int

	p.SetState(456)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RasterklangDslParserSTAR, RasterklangDslParserSLASH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(452)
			_la = p.GetTokenStream().LA(1)

			if !(_la == RasterklangDslParserSTAR || _la == RasterklangDslParserSLASH) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(453)
			p.Match(RasterklangDslParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RasterklangDslParserSHIFT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(454)
			p.Match(RasterklangDslParserSHIFT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(455)
			p.Match(RasterklangDslParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumberListContext is an interface to support dynamic dispatch.
type INumberListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsNumberListContext differentiates from other interfaces.
	IsNumberListContext()
}

type NumberListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberListContext() *NumberListContext {
	var p = new(NumberListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_numberList
	return p
}

func InitEmptyNumberListContext(p *NumberListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RasterklangDslParserRULE_numberList
}

func (*NumberListContext) IsNumberListContext() {}

func NewNumberListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberListContext {
	var p = new(NumberListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RasterklangDslParserRULE_numberList

	return p
}

func (s *NumberListContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberListContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(RasterklangDslParserNUMBER)
}

func (s *NumberListContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserNUMBER, i)
}

func (s *NumberListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(RasterklangDslParserCOMMA)
}

func (s *NumberListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(RasterklangDslParserCOMMA, i)
}

func (s *NumberListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.EnterNumberList(s)
	}
}

func (s *NumberListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RasterklangDslListener); ok {
		listenerT.ExitNumberList(s)
	}
}

func (s *NumberListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RasterklangDslVisitor:
		return t.VisitNumberList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RasterklangDslParser) NumberList() (localctx INumberListContext) {
	localctx = NewNumberListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, RasterklangDslParserRULE_numberList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(458)
		p.Match(RasterklangDslParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(465)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == RasterklangDslParserNUMBER || _la == RasterklangDslParserCOMMA {
		p.SetState(460)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == RasterklangDslParserCOMMA {
			{
				p.SetState(459)
				p.Match(RasterklangDslParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(462)
			p.Match(RasterklangDslParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(467)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
