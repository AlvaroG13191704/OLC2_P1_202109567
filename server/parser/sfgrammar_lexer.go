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
		"", "", "", "'!'", "'('", "')'", "'{'", "'}'", "':'", "','", "';'",
		"'='", "'+='", "'-='", "'?'", "'+'", "'-'", "'*'", "'/'", "'%'", "'=='",
		"'!='", "'>'", "'>='", "'<'", "'<='", "'&&'", "'||'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "STRING", "BOOL", "CHAR", "IF", "ELSE", "PRINT",
		"TRU", "FAL", "NIL", "DECLARATION_VAR", "DECLARATION_LET", "DIGIT_PRIMITIVE",
		"STRING_PRIMITIVE", "ID_PRIMITIVE", "NEGATION_OPERATOR", "LPAREN", "RPAREN",
		"LBRACE", "RBRACE", "COLON", "COMMA", "SEMICOLON", "IS_", "PLUS_IS",
		"MINUS_IS", "QUESTION_MARK", "PLUS", "MINUS", "MULTIPLY", "DIVIDE",
		"MODULO", "EQUALS", "NOT_EQUALS", "GREATER", "GREATER_EQUALS", "LESS",
		"LESS_EQUALS", "AND", "OR", "WHITESPACE", "MULTI_COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"INT", "FLOAT", "STRING", "BOOL", "CHAR", "IF", "ELSE", "PRINT", "TRU",
		"FAL", "NIL", "DECLARATION_VAR", "DECLARATION_LET", "DIGIT_PRIMITIVE",
		"STRING_PRIMITIVE", "ID_PRIMITIVE", "NEGATION_OPERATOR", "LPAREN", "RPAREN",
		"LBRACE", "RBRACE", "COLON", "COMMA", "SEMICOLON", "IS_", "PLUS_IS",
		"MINUS_IS", "QUESTION_MARK", "PLUS", "MINUS", "MULTIPLY", "DIVIDE",
		"MODULO", "EQUALS", "NOT_EQUALS", "GREATER", "GREATER_EQUALS", "LESS",
		"LESS_EQUALS", "AND", "OR", "WHITESPACE", "MULTI_COMMENT", "LINE_COMMENT",
		"ESC_SEQ",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 44, 282, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 4, 13, 162, 8, 13, 11, 13,
		12, 13, 163, 1, 13, 1, 13, 4, 13, 168, 8, 13, 11, 13, 12, 13, 169, 3, 13,
		172, 8, 13, 1, 14, 1, 14, 5, 14, 176, 8, 14, 10, 14, 12, 14, 179, 9, 14,
		1, 14, 1, 14, 1, 15, 1, 15, 5, 15, 185, 8, 15, 10, 15, 12, 15, 188, 9,
		15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1,
		25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30,
		1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1,
		34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38,
		1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 41, 4, 41, 249, 8, 41, 11,
		41, 12, 41, 250, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 5, 42, 259,
		8, 42, 10, 42, 12, 42, 262, 9, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1,
		43, 1, 43, 1, 43, 1, 43, 5, 43, 273, 8, 43, 10, 43, 12, 43, 276, 9, 43,
		1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 260, 0, 45, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27,
		14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45,
		23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63,
		32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81,
		41, 83, 42, 85, 43, 87, 44, 89, 0, 1, 0, 7, 1, 0, 48, 57, 1, 0, 34, 34,
		3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 4,
		0, 9, 10, 13, 13, 32, 32, 92, 92, 2, 0, 10, 10, 13, 13, 7, 0, 32, 33, 35,
		35, 43, 43, 45, 46, 58, 58, 64, 64, 91, 93, 288, 0, 1, 1, 0, 0, 0, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0,
		19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0,
		0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0,
		0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0,
		0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1,
		0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57,
		1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0,
		65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0,
		0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0,
		0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0,
		0, 0, 1, 91, 1, 0, 0, 0, 3, 95, 1, 0, 0, 0, 5, 101, 1, 0, 0, 0, 7, 108,
		1, 0, 0, 0, 9, 113, 1, 0, 0, 0, 11, 123, 1, 0, 0, 0, 13, 126, 1, 0, 0,
		0, 15, 131, 1, 0, 0, 0, 17, 137, 1, 0, 0, 0, 19, 142, 1, 0, 0, 0, 21, 148,
		1, 0, 0, 0, 23, 152, 1, 0, 0, 0, 25, 156, 1, 0, 0, 0, 27, 161, 1, 0, 0,
		0, 29, 173, 1, 0, 0, 0, 31, 182, 1, 0, 0, 0, 33, 189, 1, 0, 0, 0, 35, 191,
		1, 0, 0, 0, 37, 193, 1, 0, 0, 0, 39, 195, 1, 0, 0, 0, 41, 197, 1, 0, 0,
		0, 43, 199, 1, 0, 0, 0, 45, 201, 1, 0, 0, 0, 47, 203, 1, 0, 0, 0, 49, 205,
		1, 0, 0, 0, 51, 207, 1, 0, 0, 0, 53, 210, 1, 0, 0, 0, 55, 213, 1, 0, 0,
		0, 57, 215, 1, 0, 0, 0, 59, 217, 1, 0, 0, 0, 61, 219, 1, 0, 0, 0, 63, 221,
		1, 0, 0, 0, 65, 223, 1, 0, 0, 0, 67, 225, 1, 0, 0, 0, 69, 228, 1, 0, 0,
		0, 71, 231, 1, 0, 0, 0, 73, 233, 1, 0, 0, 0, 75, 236, 1, 0, 0, 0, 77, 238,
		1, 0, 0, 0, 79, 241, 1, 0, 0, 0, 81, 244, 1, 0, 0, 0, 83, 248, 1, 0, 0,
		0, 85, 254, 1, 0, 0, 0, 87, 268, 1, 0, 0, 0, 89, 279, 1, 0, 0, 0, 91, 92,
		5, 73, 0, 0, 92, 93, 5, 110, 0, 0, 93, 94, 5, 116, 0, 0, 94, 2, 1, 0, 0,
		0, 95, 96, 5, 70, 0, 0, 96, 97, 5, 108, 0, 0, 97, 98, 5, 111, 0, 0, 98,
		99, 5, 97, 0, 0, 99, 100, 5, 116, 0, 0, 100, 4, 1, 0, 0, 0, 101, 102, 5,
		83, 0, 0, 102, 103, 5, 116, 0, 0, 103, 104, 5, 114, 0, 0, 104, 105, 5,
		105, 0, 0, 105, 106, 5, 110, 0, 0, 106, 107, 5, 103, 0, 0, 107, 6, 1, 0,
		0, 0, 108, 109, 5, 66, 0, 0, 109, 110, 5, 111, 0, 0, 110, 111, 5, 111,
		0, 0, 111, 112, 5, 108, 0, 0, 112, 8, 1, 0, 0, 0, 113, 114, 5, 67, 0, 0,
		114, 115, 5, 104, 0, 0, 115, 116, 5, 97, 0, 0, 116, 117, 5, 114, 0, 0,
		117, 118, 5, 97, 0, 0, 118, 119, 5, 99, 0, 0, 119, 120, 5, 116, 0, 0, 120,
		121, 5, 101, 0, 0, 121, 122, 5, 114, 0, 0, 122, 10, 1, 0, 0, 0, 123, 124,
		5, 105, 0, 0, 124, 125, 5, 102, 0, 0, 125, 12, 1, 0, 0, 0, 126, 127, 5,
		101, 0, 0, 127, 128, 5, 108, 0, 0, 128, 129, 5, 115, 0, 0, 129, 130, 5,
		101, 0, 0, 130, 14, 1, 0, 0, 0, 131, 132, 5, 112, 0, 0, 132, 133, 5, 114,
		0, 0, 133, 134, 5, 105, 0, 0, 134, 135, 5, 110, 0, 0, 135, 136, 5, 116,
		0, 0, 136, 16, 1, 0, 0, 0, 137, 138, 5, 116, 0, 0, 138, 139, 5, 114, 0,
		0, 139, 140, 5, 117, 0, 0, 140, 141, 5, 101, 0, 0, 141, 18, 1, 0, 0, 0,
		142, 143, 5, 102, 0, 0, 143, 144, 5, 97, 0, 0, 144, 145, 5, 108, 0, 0,
		145, 146, 5, 115, 0, 0, 146, 147, 5, 101, 0, 0, 147, 20, 1, 0, 0, 0, 148,
		149, 5, 110, 0, 0, 149, 150, 5, 105, 0, 0, 150, 151, 5, 108, 0, 0, 151,
		22, 1, 0, 0, 0, 152, 153, 5, 118, 0, 0, 153, 154, 5, 97, 0, 0, 154, 155,
		5, 114, 0, 0, 155, 24, 1, 0, 0, 0, 156, 157, 5, 108, 0, 0, 157, 158, 5,
		101, 0, 0, 158, 159, 5, 116, 0, 0, 159, 26, 1, 0, 0, 0, 160, 162, 7, 0,
		0, 0, 161, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0,
		163, 164, 1, 0, 0, 0, 164, 171, 1, 0, 0, 0, 165, 167, 5, 46, 0, 0, 166,
		168, 7, 0, 0, 0, 167, 166, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 167,
		1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 172, 1, 0, 0, 0, 171, 165, 1, 0,
		0, 0, 171, 172, 1, 0, 0, 0, 172, 28, 1, 0, 0, 0, 173, 177, 5, 34, 0, 0,
		174, 176, 8, 1, 0, 0, 175, 174, 1, 0, 0, 0, 176, 179, 1, 0, 0, 0, 177,
		175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 180, 1, 0, 0, 0, 179, 177,
		1, 0, 0, 0, 180, 181, 5, 34, 0, 0, 181, 30, 1, 0, 0, 0, 182, 186, 7, 2,
		0, 0, 183, 185, 7, 3, 0, 0, 184, 183, 1, 0, 0, 0, 185, 188, 1, 0, 0, 0,
		186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 32, 1, 0, 0, 0, 188, 186,
		1, 0, 0, 0, 189, 190, 5, 33, 0, 0, 190, 34, 1, 0, 0, 0, 191, 192, 5, 40,
		0, 0, 192, 36, 1, 0, 0, 0, 193, 194, 5, 41, 0, 0, 194, 38, 1, 0, 0, 0,
		195, 196, 5, 123, 0, 0, 196, 40, 1, 0, 0, 0, 197, 198, 5, 125, 0, 0, 198,
		42, 1, 0, 0, 0, 199, 200, 5, 58, 0, 0, 200, 44, 1, 0, 0, 0, 201, 202, 5,
		44, 0, 0, 202, 46, 1, 0, 0, 0, 203, 204, 5, 59, 0, 0, 204, 48, 1, 0, 0,
		0, 205, 206, 5, 61, 0, 0, 206, 50, 1, 0, 0, 0, 207, 208, 5, 43, 0, 0, 208,
		209, 5, 61, 0, 0, 209, 52, 1, 0, 0, 0, 210, 211, 5, 45, 0, 0, 211, 212,
		5, 61, 0, 0, 212, 54, 1, 0, 0, 0, 213, 214, 5, 63, 0, 0, 214, 56, 1, 0,
		0, 0, 215, 216, 5, 43, 0, 0, 216, 58, 1, 0, 0, 0, 217, 218, 5, 45, 0, 0,
		218, 60, 1, 0, 0, 0, 219, 220, 5, 42, 0, 0, 220, 62, 1, 0, 0, 0, 221, 222,
		5, 47, 0, 0, 222, 64, 1, 0, 0, 0, 223, 224, 5, 37, 0, 0, 224, 66, 1, 0,
		0, 0, 225, 226, 5, 61, 0, 0, 226, 227, 5, 61, 0, 0, 227, 68, 1, 0, 0, 0,
		228, 229, 5, 33, 0, 0, 229, 230, 5, 61, 0, 0, 230, 70, 1, 0, 0, 0, 231,
		232, 5, 62, 0, 0, 232, 72, 1, 0, 0, 0, 233, 234, 5, 62, 0, 0, 234, 235,
		5, 61, 0, 0, 235, 74, 1, 0, 0, 0, 236, 237, 5, 60, 0, 0, 237, 76, 1, 0,
		0, 0, 238, 239, 5, 60, 0, 0, 239, 240, 5, 61, 0, 0, 240, 78, 1, 0, 0, 0,
		241, 242, 5, 38, 0, 0, 242, 243, 5, 38, 0, 0, 243, 80, 1, 0, 0, 0, 244,
		245, 5, 124, 0, 0, 245, 246, 5, 124, 0, 0, 246, 82, 1, 0, 0, 0, 247, 249,
		7, 4, 0, 0, 248, 247, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 248, 1, 0,
		0, 0, 250, 251, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252, 253, 6, 41, 0, 0,
		253, 84, 1, 0, 0, 0, 254, 255, 5, 47, 0, 0, 255, 256, 5, 42, 0, 0, 256,
		260, 1, 0, 0, 0, 257, 259, 9, 0, 0, 0, 258, 257, 1, 0, 0, 0, 259, 262,
		1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 260, 258, 1, 0, 0, 0, 261, 263, 1, 0,
		0, 0, 262, 260, 1, 0, 0, 0, 263, 264, 5, 42, 0, 0, 264, 265, 5, 47, 0,
		0, 265, 266, 1, 0, 0, 0, 266, 267, 6, 42, 0, 0, 267, 86, 1, 0, 0, 0, 268,
		269, 5, 47, 0, 0, 269, 270, 5, 47, 0, 0, 270, 274, 1, 0, 0, 0, 271, 273,
		8, 5, 0, 0, 272, 271, 1, 0, 0, 0, 273, 276, 1, 0, 0, 0, 274, 272, 1, 0,
		0, 0, 274, 275, 1, 0, 0, 0, 275, 277, 1, 0, 0, 0, 276, 274, 1, 0, 0, 0,
		277, 278, 6, 43, 0, 0, 278, 88, 1, 0, 0, 0, 279, 280, 5, 92, 0, 0, 280,
		281, 7, 6, 0, 0, 281, 90, 1, 0, 0, 0, 9, 0, 163, 169, 171, 177, 186, 250,
		260, 274, 1, 6, 0, 0,
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
	SFGrammarLexerLBRACE            = 20
	SFGrammarLexerRBRACE            = 21
	SFGrammarLexerCOLON             = 22
	SFGrammarLexerCOMMA             = 23
	SFGrammarLexerSEMICOLON         = 24
	SFGrammarLexerIS_               = 25
	SFGrammarLexerPLUS_IS           = 26
	SFGrammarLexerMINUS_IS          = 27
	SFGrammarLexerQUESTION_MARK     = 28
	SFGrammarLexerPLUS              = 29
	SFGrammarLexerMINUS             = 30
	SFGrammarLexerMULTIPLY          = 31
	SFGrammarLexerDIVIDE            = 32
	SFGrammarLexerMODULO            = 33
	SFGrammarLexerEQUALS            = 34
	SFGrammarLexerNOT_EQUALS        = 35
	SFGrammarLexerGREATER           = 36
	SFGrammarLexerGREATER_EQUALS    = 37
	SFGrammarLexerLESS              = 38
	SFGrammarLexerLESS_EQUALS       = 39
	SFGrammarLexerAND               = 40
	SFGrammarLexerOR                = 41
	SFGrammarLexerWHITESPACE        = 42
	SFGrammarLexerMULTI_COMMENT     = 43
	SFGrammarLexerLINE_COMMENT      = 44
)
