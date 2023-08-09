// Code generated from SFGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SFGrammarLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SFGrammarLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sfgrammarlexerLexerInit() {
	staticData := &SFGrammarLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'Int'", "'Float'", "'String'", "'Bool'", "'Character'", "'if'",
		"'else'", "'print'", "'true'", "'false'", "'nil'", "'var'", "'let'",
		"", "", "", "'!'", "'('", "')'", "':'", "','", "';'", "'='", "'+='",
		"'-='", "'?'", "'+'", "'-'", "'*'", "'/'", "'%'", "'=='", "'!='", "'>'",
		"'>='", "'<'", "'<='", "'&&'", "'||'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "STRING", "BOOL", "CHAR", "IF", "ELSE", "PRINT",
		"TRU", "FAL", "NIL", "DECLARATION_VAR", "DECLARATION_LET", "DIGIT_PRIMITIVE",
		"STRING_PRIMITIVE", "ID_PRIMITIVE", "NEGATION_OPERATOR", "LPAREN", "RPAREN",
		"COLON", "COMMA", "SEMICOLON", "IS_", "PLUS_IS", "MINUS_IS", "QUESTION_MARK",
		"PLUS", "MINUS", "MULTIPLY", "DIVIDE", "MODULO", "EQUALS", "NOT_EQUALS",
		"GREATER", "GREATER_EQUALS", "LESS", "LESS_EQUALS", "AND", "OR", "WHITESPACE",
		"MULTI_COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"INT", "FLOAT", "STRING", "BOOL", "CHAR", "IF", "ELSE", "PRINT", "TRU",
		"FAL", "NIL", "DECLARATION_VAR", "DECLARATION_LET", "DIGIT_PRIMITIVE",
		"STRING_PRIMITIVE", "ID_PRIMITIVE", "NEGATION_OPERATOR", "LPAREN", "RPAREN",
		"COLON", "COMMA", "SEMICOLON", "IS_", "PLUS_IS", "MINUS_IS", "QUESTION_MARK",
		"PLUS", "MINUS", "MULTIPLY", "DIVIDE", "MODULO", "EQUALS", "NOT_EQUALS",
		"GREATER", "GREATER_EQUALS", "LESS", "LESS_EQUALS", "AND", "OR", "WHITESPACE",
		"MULTI_COMMENT", "LINE_COMMENT", "ESC_SEQ",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 42, 274, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 13, 4, 13, 158, 8, 13, 11, 13, 12, 13, 159, 1, 13, 1, 13,
		4, 13, 164, 8, 13, 11, 13, 12, 13, 165, 3, 13, 168, 8, 13, 1, 14, 1, 14,
		5, 14, 172, 8, 14, 10, 14, 12, 14, 175, 9, 14, 1, 14, 1, 14, 1, 15, 1,
		15, 5, 15, 181, 8, 15, 10, 15, 12, 15, 184, 9, 15, 1, 16, 1, 16, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26,
		1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1,
		31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35,
		1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 39, 4,
		39, 241, 8, 39, 11, 39, 12, 39, 242, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40,
		1, 40, 5, 40, 251, 8, 40, 10, 40, 12, 40, 254, 9, 40, 1, 40, 1, 40, 1,
		40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 5, 41, 265, 8, 41, 10, 41,
		12, 41, 268, 9, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 252, 0, 43, 1,
		1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29,
		59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38,
		77, 39, 79, 40, 81, 41, 83, 42, 85, 0, 1, 0, 7, 1, 0, 48, 57, 1, 0, 34,
		34, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122,
		4, 0, 9, 10, 13, 13, 32, 32, 92, 92, 2, 0, 10, 10, 13, 13, 7, 0, 32, 33,
		35, 35, 43, 43, 45, 46, 58, 58, 64, 64, 91, 93, 280, 0, 1, 1, 0, 0, 0,
		0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0,
		0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0,
		0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0,
		0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1,
		0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41,
		1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0,
		49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0,
		0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0,
		0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0,
		0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1,
		0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 1, 87, 1, 0, 0, 0, 3, 91,
		1, 0, 0, 0, 5, 97, 1, 0, 0, 0, 7, 104, 1, 0, 0, 0, 9, 109, 1, 0, 0, 0,
		11, 119, 1, 0, 0, 0, 13, 122, 1, 0, 0, 0, 15, 127, 1, 0, 0, 0, 17, 133,
		1, 0, 0, 0, 19, 138, 1, 0, 0, 0, 21, 144, 1, 0, 0, 0, 23, 148, 1, 0, 0,
		0, 25, 152, 1, 0, 0, 0, 27, 157, 1, 0, 0, 0, 29, 169, 1, 0, 0, 0, 31, 178,
		1, 0, 0, 0, 33, 185, 1, 0, 0, 0, 35, 187, 1, 0, 0, 0, 37, 189, 1, 0, 0,
		0, 39, 191, 1, 0, 0, 0, 41, 193, 1, 0, 0, 0, 43, 195, 1, 0, 0, 0, 45, 197,
		1, 0, 0, 0, 47, 199, 1, 0, 0, 0, 49, 202, 1, 0, 0, 0, 51, 205, 1, 0, 0,
		0, 53, 207, 1, 0, 0, 0, 55, 209, 1, 0, 0, 0, 57, 211, 1, 0, 0, 0, 59, 213,
		1, 0, 0, 0, 61, 215, 1, 0, 0, 0, 63, 217, 1, 0, 0, 0, 65, 220, 1, 0, 0,
		0, 67, 223, 1, 0, 0, 0, 69, 225, 1, 0, 0, 0, 71, 228, 1, 0, 0, 0, 73, 230,
		1, 0, 0, 0, 75, 233, 1, 0, 0, 0, 77, 236, 1, 0, 0, 0, 79, 240, 1, 0, 0,
		0, 81, 246, 1, 0, 0, 0, 83, 260, 1, 0, 0, 0, 85, 271, 1, 0, 0, 0, 87, 88,
		5, 73, 0, 0, 88, 89, 5, 110, 0, 0, 89, 90, 5, 116, 0, 0, 90, 2, 1, 0, 0,
		0, 91, 92, 5, 70, 0, 0, 92, 93, 5, 108, 0, 0, 93, 94, 5, 111, 0, 0, 94,
		95, 5, 97, 0, 0, 95, 96, 5, 116, 0, 0, 96, 4, 1, 0, 0, 0, 97, 98, 5, 83,
		0, 0, 98, 99, 5, 116, 0, 0, 99, 100, 5, 114, 0, 0, 100, 101, 5, 105, 0,
		0, 101, 102, 5, 110, 0, 0, 102, 103, 5, 103, 0, 0, 103, 6, 1, 0, 0, 0,
		104, 105, 5, 66, 0, 0, 105, 106, 5, 111, 0, 0, 106, 107, 5, 111, 0, 0,
		107, 108, 5, 108, 0, 0, 108, 8, 1, 0, 0, 0, 109, 110, 5, 67, 0, 0, 110,
		111, 5, 104, 0, 0, 111, 112, 5, 97, 0, 0, 112, 113, 5, 114, 0, 0, 113,
		114, 5, 97, 0, 0, 114, 115, 5, 99, 0, 0, 115, 116, 5, 116, 0, 0, 116, 117,
		5, 101, 0, 0, 117, 118, 5, 114, 0, 0, 118, 10, 1, 0, 0, 0, 119, 120, 5,
		105, 0, 0, 120, 121, 5, 102, 0, 0, 121, 12, 1, 0, 0, 0, 122, 123, 5, 101,
		0, 0, 123, 124, 5, 108, 0, 0, 124, 125, 5, 115, 0, 0, 125, 126, 5, 101,
		0, 0, 126, 14, 1, 0, 0, 0, 127, 128, 5, 112, 0, 0, 128, 129, 5, 114, 0,
		0, 129, 130, 5, 105, 0, 0, 130, 131, 5, 110, 0, 0, 131, 132, 5, 116, 0,
		0, 132, 16, 1, 0, 0, 0, 133, 134, 5, 116, 0, 0, 134, 135, 5, 114, 0, 0,
		135, 136, 5, 117, 0, 0, 136, 137, 5, 101, 0, 0, 137, 18, 1, 0, 0, 0, 138,
		139, 5, 102, 0, 0, 139, 140, 5, 97, 0, 0, 140, 141, 5, 108, 0, 0, 141,
		142, 5, 115, 0, 0, 142, 143, 5, 101, 0, 0, 143, 20, 1, 0, 0, 0, 144, 145,
		5, 110, 0, 0, 145, 146, 5, 105, 0, 0, 146, 147, 5, 108, 0, 0, 147, 22,
		1, 0, 0, 0, 148, 149, 5, 118, 0, 0, 149, 150, 5, 97, 0, 0, 150, 151, 5,
		114, 0, 0, 151, 24, 1, 0, 0, 0, 152, 153, 5, 108, 0, 0, 153, 154, 5, 101,
		0, 0, 154, 155, 5, 116, 0, 0, 155, 26, 1, 0, 0, 0, 156, 158, 7, 0, 0, 0,
		157, 156, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 159,
		160, 1, 0, 0, 0, 160, 167, 1, 0, 0, 0, 161, 163, 5, 46, 0, 0, 162, 164,
		7, 0, 0, 0, 163, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 163, 1, 0,
		0, 0, 165, 166, 1, 0, 0, 0, 166, 168, 1, 0, 0, 0, 167, 161, 1, 0, 0, 0,
		167, 168, 1, 0, 0, 0, 168, 28, 1, 0, 0, 0, 169, 173, 5, 34, 0, 0, 170,
		172, 8, 1, 0, 0, 171, 170, 1, 0, 0, 0, 172, 175, 1, 0, 0, 0, 173, 171,
		1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 176, 1, 0, 0, 0, 175, 173, 1, 0,
		0, 0, 176, 177, 5, 34, 0, 0, 177, 30, 1, 0, 0, 0, 178, 182, 7, 2, 0, 0,
		179, 181, 7, 3, 0, 0, 180, 179, 1, 0, 0, 0, 181, 184, 1, 0, 0, 0, 182,
		180, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 32, 1, 0, 0, 0, 184, 182, 1,
		0, 0, 0, 185, 186, 5, 33, 0, 0, 186, 34, 1, 0, 0, 0, 187, 188, 5, 40, 0,
		0, 188, 36, 1, 0, 0, 0, 189, 190, 5, 41, 0, 0, 190, 38, 1, 0, 0, 0, 191,
		192, 5, 58, 0, 0, 192, 40, 1, 0, 0, 0, 193, 194, 5, 44, 0, 0, 194, 42,
		1, 0, 0, 0, 195, 196, 5, 59, 0, 0, 196, 44, 1, 0, 0, 0, 197, 198, 5, 61,
		0, 0, 198, 46, 1, 0, 0, 0, 199, 200, 5, 43, 0, 0, 200, 201, 5, 61, 0, 0,
		201, 48, 1, 0, 0, 0, 202, 203, 5, 45, 0, 0, 203, 204, 5, 61, 0, 0, 204,
		50, 1, 0, 0, 0, 205, 206, 5, 63, 0, 0, 206, 52, 1, 0, 0, 0, 207, 208, 5,
		43, 0, 0, 208, 54, 1, 0, 0, 0, 209, 210, 5, 45, 0, 0, 210, 56, 1, 0, 0,
		0, 211, 212, 5, 42, 0, 0, 212, 58, 1, 0, 0, 0, 213, 214, 5, 47, 0, 0, 214,
		60, 1, 0, 0, 0, 215, 216, 5, 37, 0, 0, 216, 62, 1, 0, 0, 0, 217, 218, 5,
		61, 0, 0, 218, 219, 5, 61, 0, 0, 219, 64, 1, 0, 0, 0, 220, 221, 5, 33,
		0, 0, 221, 222, 5, 61, 0, 0, 222, 66, 1, 0, 0, 0, 223, 224, 5, 62, 0, 0,
		224, 68, 1, 0, 0, 0, 225, 226, 5, 62, 0, 0, 226, 227, 5, 61, 0, 0, 227,
		70, 1, 0, 0, 0, 228, 229, 5, 60, 0, 0, 229, 72, 1, 0, 0, 0, 230, 231, 5,
		60, 0, 0, 231, 232, 5, 61, 0, 0, 232, 74, 1, 0, 0, 0, 233, 234, 5, 38,
		0, 0, 234, 235, 5, 38, 0, 0, 235, 76, 1, 0, 0, 0, 236, 237, 5, 124, 0,
		0, 237, 238, 5, 124, 0, 0, 238, 78, 1, 0, 0, 0, 239, 241, 7, 4, 0, 0, 240,
		239, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 242, 243,
		1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 245, 6, 39, 0, 0, 245, 80, 1, 0,
		0, 0, 246, 247, 5, 47, 0, 0, 247, 248, 5, 42, 0, 0, 248, 252, 1, 0, 0,
		0, 249, 251, 9, 0, 0, 0, 250, 249, 1, 0, 0, 0, 251, 254, 1, 0, 0, 0, 252,
		253, 1, 0, 0, 0, 252, 250, 1, 0, 0, 0, 253, 255, 1, 0, 0, 0, 254, 252,
		1, 0, 0, 0, 255, 256, 5, 42, 0, 0, 256, 257, 5, 47, 0, 0, 257, 258, 1,
		0, 0, 0, 258, 259, 6, 40, 0, 0, 259, 82, 1, 0, 0, 0, 260, 261, 5, 47, 0,
		0, 261, 262, 5, 47, 0, 0, 262, 266, 1, 0, 0, 0, 263, 265, 8, 5, 0, 0, 264,
		263, 1, 0, 0, 0, 265, 268, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 266, 267,
		1, 0, 0, 0, 267, 269, 1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 269, 270, 6, 41,
		0, 0, 270, 84, 1, 0, 0, 0, 271, 272, 5, 92, 0, 0, 272, 273, 7, 6, 0, 0,
		273, 86, 1, 0, 0, 0, 9, 0, 159, 165, 167, 173, 182, 242, 252, 266, 1, 6,
		0, 0,
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

// SFGrammarLexerInit initializes any static state used to implement SFGrammarLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSFGrammarLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SFGrammarLexerInit() {
	staticData := &SFGrammarLexerLexerStaticData
	staticData.once.Do(sfgrammarlexerLexerInit)
}

// NewSFGrammarLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSFGrammarLexer(input antlr.CharStream) *SFGrammarLexer {
	SFGrammarLexerInit()
	l := new(SFGrammarLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SFGrammarLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "SFGrammar.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SFGrammarLexer tokens.
const (
	SFGrammarLexerINT               = 1
	SFGrammarLexerFLOAT             = 2
	SFGrammarLexerSTRING            = 3
	SFGrammarLexerBOOL              = 4
	SFGrammarLexerCHAR              = 5
	SFGrammarLexerIF                = 6
	SFGrammarLexerELSE              = 7
	SFGrammarLexerPRINT             = 8
	SFGrammarLexerTRU               = 9
	SFGrammarLexerFAL               = 10
	SFGrammarLexerNIL               = 11
	SFGrammarLexerDECLARATION_VAR   = 12
	SFGrammarLexerDECLARATION_LET   = 13
	SFGrammarLexerDIGIT_PRIMITIVE   = 14
	SFGrammarLexerSTRING_PRIMITIVE  = 15
	SFGrammarLexerID_PRIMITIVE      = 16
	SFGrammarLexerNEGATION_OPERATOR = 17
	SFGrammarLexerLPAREN            = 18
	SFGrammarLexerRPAREN            = 19
	SFGrammarLexerCOLON             = 20
	SFGrammarLexerCOMMA             = 21
	SFGrammarLexerSEMICOLON         = 22
	SFGrammarLexerIS_               = 23
	SFGrammarLexerPLUS_IS           = 24
	SFGrammarLexerMINUS_IS          = 25
	SFGrammarLexerQUESTION_MARK     = 26
	SFGrammarLexerPLUS              = 27
	SFGrammarLexerMINUS             = 28
	SFGrammarLexerMULTIPLY          = 29
	SFGrammarLexerDIVIDE            = 30
	SFGrammarLexerMODULO            = 31
	SFGrammarLexerEQUALS            = 32
	SFGrammarLexerNOT_EQUALS        = 33
	SFGrammarLexerGREATER           = 34
	SFGrammarLexerGREATER_EQUALS    = 35
	SFGrammarLexerLESS              = 36
	SFGrammarLexerLESS_EQUALS       = 37
	SFGrammarLexerAND               = 38
	SFGrammarLexerOR                = 39
	SFGrammarLexerWHITESPACE        = 40
	SFGrammarLexerMULTI_COMMENT     = 41
	SFGrammarLexerLINE_COMMENT      = 42
)
