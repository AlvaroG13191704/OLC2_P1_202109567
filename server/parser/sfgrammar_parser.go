// Code generated from SFGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SFGrammar
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

type SFGrammarParser struct {
	*antlr.BaseParser
}

var SFGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sfgrammarParserInit() {
	staticData := &SFGrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "'Int'", "'Float'", "'String'", "'Bool'", "'Character'", "'if'",
		"'struct'", "'mutating'", "'append'", "'count'", "'removeLast'", "'remove'",
		"'isEmpty'", "'self'", "'at'", "'else'", "'switch'", "'case'", "'default'",
		"'break'", "'continue'", "'return'", "'while'", "'for'", "'func'", "'->'",
		"'in'", "'.'", "'guard'", "'print'", "'true'", "'false'", "'nil'", "'var'",
		"'let'", "'&'", "'inout'", "'_'", "", "", "", "'!'", "'('", "')'", "'{'",
		"'}'", "'['", "']'", "':'", "','", "';'", "'='", "'+='", "'-='", "'?'",
		"'+'", "'-'", "'*'", "'/'", "'%'", "'=='", "'!='", "'>'", "'>='", "'<'",
		"'<='", "'&&'", "'||'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "STRING", "BOOL", "CHAR", "IF", "STRUCT", "MUTATING",
		"APPEND", "COUNT", "REMOVELAST", "REMOVE", "ISEMPTY", "SELF", "AT",
		"ELSE", "SWITCH", "CASE", "DEFAULT", "BREAK", "CONTINUE", "RETURN",
		"WHILE", "FOR", "FUNC", "ARROW_FUNCTION", "IN", "DOT", "GUARD", "PRINT",
		"TRU", "FAL", "NIL", "DECLARATION_VAR", "DECLARATION_LET", "REFERENCE",
		"INOUT", "NOT_PARAM", "DIGIT_PRIMITIVE", "STRING_PRIMITIVE", "ID_PRIMITIVE",
		"NEGATION_OPERATOR", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACKET",
		"RBRACKET", "COLON", "COMMA", "SEMICOLON", "IS_", "PLUS_IS", "MINUS_IS",
		"QUESTION_MARK", "PLUS", "MINUS", "MULTIPLY", "DIVIDE", "MODULO", "EQUALS",
		"NOT_EQUALS", "GREATER", "GREATER_EQUALS", "LESS", "LESS_EQUALS", "AND",
		"OR", "WHITESPACE", "MULTI_COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"start", "block", "stmts", "transferStmt", "structStmt", "structBlock",
		"structStmts", "declarationStructs", "functionStructs", "structCallList",
		"declaration", "type_declaration", "assignment", "ifstmt", "switchStmt",
		"caseBlock", "defaultBlock", "whileStmt", "forStmt", "forRange", "guardStmt",
		"functionStmt", "listFunctionParams", "callFunctionStmt", "listCallFunctionStmt",
		"callBack", "embbededFunc", "printstmt", "exprList", "intstmt", "floatstmt",
		"stringstmt", "expr", "type",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 71, 758, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 73, 8, 1,
		10, 1, 12, 1, 76, 9, 1, 1, 2, 1, 2, 3, 2, 80, 8, 2, 1, 2, 1, 2, 3, 2, 84,
		8, 2, 1, 2, 1, 2, 3, 2, 88, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 3, 2, 99, 8, 2, 1, 2, 1, 2, 3, 2, 103, 8, 2, 1, 2, 1, 2,
		3, 2, 107, 8, 2, 1, 2, 3, 2, 110, 8, 2, 1, 3, 1, 3, 3, 3, 114, 8, 3, 1,
		3, 1, 3, 3, 3, 118, 8, 3, 1, 3, 1, 3, 3, 3, 122, 8, 3, 1, 3, 3, 3, 125,
		8, 3, 3, 3, 127, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 5, 5,
		136, 8, 5, 10, 5, 12, 5, 139, 9, 5, 1, 6, 1, 6, 3, 6, 143, 8, 6, 1, 6,
		3, 6, 146, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 3, 7, 160, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7,
		179, 8, 7, 3, 7, 181, 8, 7, 1, 8, 3, 8, 184, 8, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 3, 8, 192, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 199,
		8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 208, 8, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 3, 8, 214, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 5, 9, 223, 8, 9, 10, 9, 12, 9, 226, 9, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 3, 10, 266, 8, 10, 3, 10, 268, 8, 10, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 302,
		8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 3, 13, 314, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 3, 13, 324, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 330, 8, 14,
		10, 14, 12, 14, 333, 9, 14, 1, 14, 3, 14, 336, 8, 14, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 359, 8,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 375, 8, 18, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 396, 8, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 409,
		8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 415, 8, 21, 1, 22, 1, 22, 1,
		22, 1, 22, 3, 22, 421, 8, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		5, 22, 429, 8, 22, 10, 22, 12, 22, 432, 9, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 3, 22, 438, 8, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22,
		446, 8, 22, 10, 22, 12, 22, 449, 9, 22, 1, 22, 1, 22, 1, 22, 3, 22, 454,
		8, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 461, 8, 22, 1, 22, 5,
		22, 464, 8, 22, 10, 22, 12, 22, 467, 9, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		3, 22, 473, 8, 22, 1, 22, 3, 22, 476, 8, 22, 1, 22, 1, 22, 3, 22, 480,
		8, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 487, 8, 22, 1, 22, 3,
		22, 490, 8, 22, 1, 22, 1, 22, 3, 22, 494, 8, 22, 5, 22, 496, 8, 22, 10,
		22, 12, 22, 499, 9, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 505, 8, 22,
		1, 22, 3, 22, 508, 8, 22, 1, 22, 1, 22, 3, 22, 512, 8, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 3, 22, 519, 8, 22, 1, 22, 3, 22, 522, 8, 22, 1, 22,
		1, 22, 3, 22, 526, 8, 22, 5, 22, 528, 8, 22, 10, 22, 12, 22, 531, 9, 22,
		1, 22, 1, 22, 1, 22, 3, 22, 536, 8, 22, 1, 22, 3, 22, 539, 8, 22, 1, 22,
		1, 22, 3, 22, 543, 8, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 549, 8, 22,
		1, 22, 3, 22, 552, 8, 22, 1, 22, 1, 22, 3, 22, 556, 8, 22, 5, 22, 558,
		8, 22, 10, 22, 12, 22, 561, 9, 22, 3, 22, 563, 8, 22, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 573, 8, 23, 1, 24, 3, 24,
		576, 8, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 583, 8, 24, 1, 24,
		1, 24, 1, 24, 5, 24, 588, 8, 24, 10, 24, 12, 24, 591, 9, 24, 1, 24, 3,
		24, 594, 8, 24, 1, 24, 1, 24, 1, 24, 3, 24, 599, 8, 24, 1, 24, 5, 24, 602,
		8, 24, 10, 24, 12, 24, 605, 9, 24, 3, 24, 607, 8, 24, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 4, 25, 644, 8, 25, 11, 25, 12, 25, 645, 1, 25, 1, 25, 3, 25, 650,
		8, 25, 1, 25, 1, 25, 1, 25, 1, 25, 4, 25, 656, 8, 25, 11, 25, 12, 25, 657,
		1, 25, 1, 25, 3, 25, 662, 8, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3,
		25, 669, 8, 25, 3, 25, 671, 8, 25, 1, 26, 1, 26, 1, 26, 3, 26, 676, 8,
		26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 5, 28, 686,
		8, 28, 10, 28, 12, 28, 689, 9, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 722, 8, 32, 1, 32, 1, 32, 3, 32,
		726, 8, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 734, 8, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 5, 32, 751, 8, 32, 10, 32, 12, 32, 754,
		9, 32, 1, 33, 1, 33, 1, 33, 0, 1, 64, 34, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
		54, 56, 58, 60, 62, 64, 66, 0, 8, 1, 0, 34, 35, 1, 0, 31, 32, 1, 0, 58,
		60, 1, 0, 56, 57, 1, 0, 61, 62, 1, 0, 63, 66, 1, 0, 67, 68, 2, 0, 1, 5,
		41, 41, 858, 0, 68, 1, 0, 0, 0, 2, 74, 1, 0, 0, 0, 4, 109, 1, 0, 0, 0,
		6, 126, 1, 0, 0, 0, 8, 128, 1, 0, 0, 0, 10, 137, 1, 0, 0, 0, 12, 145, 1,
		0, 0, 0, 14, 180, 1, 0, 0, 0, 16, 213, 1, 0, 0, 0, 18, 215, 1, 0, 0, 0,
		20, 267, 1, 0, 0, 0, 22, 269, 1, 0, 0, 0, 24, 301, 1, 0, 0, 0, 26, 323,
		1, 0, 0, 0, 28, 325, 1, 0, 0, 0, 30, 339, 1, 0, 0, 0, 32, 344, 1, 0, 0,
		0, 34, 348, 1, 0, 0, 0, 36, 374, 1, 0, 0, 0, 38, 376, 1, 0, 0, 0, 40, 382,
		1, 0, 0, 0, 42, 414, 1, 0, 0, 0, 44, 562, 1, 0, 0, 0, 46, 572, 1, 0, 0,
		0, 48, 606, 1, 0, 0, 0, 50, 670, 1, 0, 0, 0, 52, 675, 1, 0, 0, 0, 54, 677,
		1, 0, 0, 0, 56, 682, 1, 0, 0, 0, 58, 690, 1, 0, 0, 0, 60, 695, 1, 0, 0,
		0, 62, 700, 1, 0, 0, 0, 64, 733, 1, 0, 0, 0, 66, 755, 1, 0, 0, 0, 68, 69,
		3, 2, 1, 0, 69, 70, 5, 0, 0, 1, 70, 1, 1, 0, 0, 0, 71, 73, 3, 4, 2, 0,
		72, 71, 1, 0, 0, 0, 73, 76, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 74, 75, 1,
		0, 0, 0, 75, 3, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 77, 79, 3, 20, 10, 0, 78,
		80, 5, 51, 0, 0, 79, 78, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 110, 1, 0,
		0, 0, 81, 83, 3, 24, 12, 0, 82, 84, 5, 51, 0, 0, 83, 82, 1, 0, 0, 0, 83,
		84, 1, 0, 0, 0, 84, 110, 1, 0, 0, 0, 85, 87, 3, 52, 26, 0, 86, 88, 5, 51,
		0, 0, 87, 86, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 110, 1, 0, 0, 0, 89,
		110, 3, 26, 13, 0, 90, 110, 3, 28, 14, 0, 91, 110, 3, 34, 17, 0, 92, 110,
		3, 36, 18, 0, 93, 110, 3, 40, 20, 0, 94, 110, 3, 6, 3, 0, 95, 110, 3, 42,
		21, 0, 96, 98, 3, 54, 27, 0, 97, 99, 5, 51, 0, 0, 98, 97, 1, 0, 0, 0, 98,
		99, 1, 0, 0, 0, 99, 110, 1, 0, 0, 0, 100, 102, 3, 46, 23, 0, 101, 103,
		5, 51, 0, 0, 102, 101, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 110, 1, 0,
		0, 0, 104, 106, 3, 50, 25, 0, 105, 107, 5, 51, 0, 0, 106, 105, 1, 0, 0,
		0, 106, 107, 1, 0, 0, 0, 107, 110, 1, 0, 0, 0, 108, 110, 3, 8, 4, 0, 109,
		77, 1, 0, 0, 0, 109, 81, 1, 0, 0, 0, 109, 85, 1, 0, 0, 0, 109, 89, 1, 0,
		0, 0, 109, 90, 1, 0, 0, 0, 109, 91, 1, 0, 0, 0, 109, 92, 1, 0, 0, 0, 109,
		93, 1, 0, 0, 0, 109, 94, 1, 0, 0, 0, 109, 95, 1, 0, 0, 0, 109, 96, 1, 0,
		0, 0, 109, 100, 1, 0, 0, 0, 109, 104, 1, 0, 0, 0, 109, 108, 1, 0, 0, 0,
		110, 5, 1, 0, 0, 0, 111, 113, 5, 20, 0, 0, 112, 114, 5, 51, 0, 0, 113,
		112, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 127, 1, 0, 0, 0, 115, 117,
		5, 21, 0, 0, 116, 118, 5, 51, 0, 0, 117, 116, 1, 0, 0, 0, 117, 118, 1,
		0, 0, 0, 118, 127, 1, 0, 0, 0, 119, 121, 5, 22, 0, 0, 120, 122, 3, 64,
		32, 0, 121, 120, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 124, 1, 0, 0, 0,
		123, 125, 5, 51, 0, 0, 124, 123, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125,
		127, 1, 0, 0, 0, 126, 111, 1, 0, 0, 0, 126, 115, 1, 0, 0, 0, 126, 119,
		1, 0, 0, 0, 127, 7, 1, 0, 0, 0, 128, 129, 5, 7, 0, 0, 129, 130, 5, 41,
		0, 0, 130, 131, 5, 45, 0, 0, 131, 132, 3, 10, 5, 0, 132, 133, 5, 46, 0,
		0, 133, 9, 1, 0, 0, 0, 134, 136, 3, 12, 6, 0, 135, 134, 1, 0, 0, 0, 136,
		139, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 11, 1,
		0, 0, 0, 139, 137, 1, 0, 0, 0, 140, 142, 3, 14, 7, 0, 141, 143, 5, 51,
		0, 0, 142, 141, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 146, 1, 0, 0, 0,
		144, 146, 3, 16, 8, 0, 145, 140, 1, 0, 0, 0, 145, 144, 1, 0, 0, 0, 146,
		13, 1, 0, 0, 0, 147, 148, 3, 22, 11, 0, 148, 149, 5, 41, 0, 0, 149, 150,
		5, 49, 0, 0, 150, 151, 3, 66, 33, 0, 151, 152, 5, 52, 0, 0, 152, 153, 3,
		64, 32, 0, 153, 181, 1, 0, 0, 0, 154, 155, 3, 22, 11, 0, 155, 156, 5, 41,
		0, 0, 156, 157, 5, 49, 0, 0, 157, 159, 3, 66, 33, 0, 158, 160, 5, 55, 0,
		0, 159, 158, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 181, 1, 0, 0, 0, 161,
		162, 3, 22, 11, 0, 162, 163, 5, 41, 0, 0, 163, 164, 5, 52, 0, 0, 164, 165,
		3, 64, 32, 0, 165, 181, 1, 0, 0, 0, 166, 167, 3, 22, 11, 0, 167, 168, 5,
		41, 0, 0, 168, 169, 5, 49, 0, 0, 169, 170, 5, 47, 0, 0, 170, 171, 3, 66,
		33, 0, 171, 172, 5, 48, 0, 0, 172, 178, 5, 52, 0, 0, 173, 174, 5, 47, 0,
		0, 174, 175, 3, 56, 28, 0, 175, 176, 5, 48, 0, 0, 176, 179, 1, 0, 0, 0,
		177, 179, 5, 41, 0, 0, 178, 173, 1, 0, 0, 0, 178, 177, 1, 0, 0, 0, 179,
		181, 1, 0, 0, 0, 180, 147, 1, 0, 0, 0, 180, 154, 1, 0, 0, 0, 180, 161,
		1, 0, 0, 0, 180, 166, 1, 0, 0, 0, 181, 15, 1, 0, 0, 0, 182, 184, 5, 8,
		0, 0, 183, 182, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0,
		185, 186, 5, 25, 0, 0, 186, 187, 5, 41, 0, 0, 187, 188, 5, 43, 0, 0, 188,
		191, 5, 44, 0, 0, 189, 190, 5, 26, 0, 0, 190, 192, 3, 66, 33, 0, 191, 189,
		1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 194, 5, 45,
		0, 0, 194, 195, 3, 2, 1, 0, 195, 196, 5, 46, 0, 0, 196, 214, 1, 0, 0, 0,
		197, 199, 5, 8, 0, 0, 198, 197, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199,
		200, 1, 0, 0, 0, 200, 201, 5, 25, 0, 0, 201, 202, 5, 41, 0, 0, 202, 203,
		5, 43, 0, 0, 203, 204, 3, 44, 22, 0, 204, 207, 5, 44, 0, 0, 205, 206, 5,
		26, 0, 0, 206, 208, 3, 66, 33, 0, 207, 205, 1, 0, 0, 0, 207, 208, 1, 0,
		0, 0, 208, 209, 1, 0, 0, 0, 209, 210, 5, 45, 0, 0, 210, 211, 3, 2, 1, 0,
		211, 212, 5, 46, 0, 0, 212, 214, 1, 0, 0, 0, 213, 183, 1, 0, 0, 0, 213,
		198, 1, 0, 0, 0, 214, 17, 1, 0, 0, 0, 215, 216, 5, 41, 0, 0, 216, 217,
		5, 49, 0, 0, 217, 224, 3, 64, 32, 0, 218, 219, 5, 50, 0, 0, 219, 220, 5,
		41, 0, 0, 220, 221, 5, 49, 0, 0, 221, 223, 3, 64, 32, 0, 222, 218, 1, 0,
		0, 0, 223, 226, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0,
		225, 19, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 227, 228, 3, 22, 11, 0, 228,
		229, 5, 41, 0, 0, 229, 230, 5, 52, 0, 0, 230, 231, 5, 41, 0, 0, 231, 232,
		5, 43, 0, 0, 232, 233, 3, 18, 9, 0, 233, 234, 5, 44, 0, 0, 234, 268, 1,
		0, 0, 0, 235, 236, 3, 22, 11, 0, 236, 237, 5, 41, 0, 0, 237, 238, 5, 49,
		0, 0, 238, 239, 3, 66, 33, 0, 239, 240, 5, 52, 0, 0, 240, 241, 3, 64, 32,
		0, 241, 268, 1, 0, 0, 0, 242, 243, 3, 22, 11, 0, 243, 244, 5, 41, 0, 0,
		244, 245, 5, 49, 0, 0, 245, 246, 3, 66, 33, 0, 246, 247, 5, 55, 0, 0, 247,
		268, 1, 0, 0, 0, 248, 249, 3, 22, 11, 0, 249, 250, 5, 41, 0, 0, 250, 251,
		5, 52, 0, 0, 251, 252, 3, 64, 32, 0, 252, 268, 1, 0, 0, 0, 253, 254, 3,
		22, 11, 0, 254, 255, 5, 41, 0, 0, 255, 256, 5, 49, 0, 0, 256, 257, 5, 47,
		0, 0, 257, 258, 3, 66, 33, 0, 258, 259, 5, 48, 0, 0, 259, 265, 5, 52, 0,
		0, 260, 261, 5, 47, 0, 0, 261, 262, 3, 56, 28, 0, 262, 263, 5, 48, 0, 0,
		263, 266, 1, 0, 0, 0, 264, 266, 5, 41, 0, 0, 265, 260, 1, 0, 0, 0, 265,
		264, 1, 0, 0, 0, 266, 268, 1, 0, 0, 0, 267, 227, 1, 0, 0, 0, 267, 235,
		1, 0, 0, 0, 267, 242, 1, 0, 0, 0, 267, 248, 1, 0, 0, 0, 267, 253, 1, 0,
		0, 0, 268, 21, 1, 0, 0, 0, 269, 270, 7, 0, 0, 0, 270, 23, 1, 0, 0, 0, 271,
		272, 5, 41, 0, 0, 272, 273, 5, 52, 0, 0, 273, 302, 3, 64, 32, 0, 274, 275,
		5, 41, 0, 0, 275, 276, 5, 53, 0, 0, 276, 302, 3, 64, 32, 0, 277, 278, 5,
		41, 0, 0, 278, 279, 5, 54, 0, 0, 279, 302, 3, 64, 32, 0, 280, 281, 5, 41,
		0, 0, 281, 282, 5, 47, 0, 0, 282, 283, 3, 64, 32, 0, 283, 284, 5, 48, 0,
		0, 284, 285, 5, 52, 0, 0, 285, 286, 3, 64, 32, 0, 286, 302, 1, 0, 0, 0,
		287, 288, 5, 41, 0, 0, 288, 289, 5, 47, 0, 0, 289, 290, 3, 64, 32, 0, 290,
		291, 5, 48, 0, 0, 291, 292, 5, 54, 0, 0, 292, 293, 3, 64, 32, 0, 293, 302,
		1, 0, 0, 0, 294, 295, 5, 41, 0, 0, 295, 296, 5, 47, 0, 0, 296, 297, 3,
		64, 32, 0, 297, 298, 5, 48, 0, 0, 298, 299, 5, 53, 0, 0, 299, 300, 3, 64,
		32, 0, 300, 302, 1, 0, 0, 0, 301, 271, 1, 0, 0, 0, 301, 274, 1, 0, 0, 0,
		301, 277, 1, 0, 0, 0, 301, 280, 1, 0, 0, 0, 301, 287, 1, 0, 0, 0, 301,
		294, 1, 0, 0, 0, 302, 25, 1, 0, 0, 0, 303, 304, 5, 6, 0, 0, 304, 305, 3,
		64, 32, 0, 305, 306, 5, 45, 0, 0, 306, 307, 3, 2, 1, 0, 307, 313, 5, 46,
		0, 0, 308, 309, 5, 16, 0, 0, 309, 310, 5, 45, 0, 0, 310, 311, 3, 2, 1,
		0, 311, 312, 5, 46, 0, 0, 312, 314, 1, 0, 0, 0, 313, 308, 1, 0, 0, 0, 313,
		314, 1, 0, 0, 0, 314, 324, 1, 0, 0, 0, 315, 316, 5, 6, 0, 0, 316, 317,
		3, 64, 32, 0, 317, 318, 5, 45, 0, 0, 318, 319, 3, 2, 1, 0, 319, 320, 5,
		46, 0, 0, 320, 321, 5, 16, 0, 0, 321, 322, 3, 26, 13, 0, 322, 324, 1, 0,
		0, 0, 323, 303, 1, 0, 0, 0, 323, 315, 1, 0, 0, 0, 324, 27, 1, 0, 0, 0,
		325, 326, 5, 17, 0, 0, 326, 327, 3, 64, 32, 0, 327, 331, 5, 45, 0, 0, 328,
		330, 3, 30, 15, 0, 329, 328, 1, 0, 0, 0, 330, 333, 1, 0, 0, 0, 331, 329,
		1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332, 335, 1, 0, 0, 0, 333, 331, 1, 0,
		0, 0, 334, 336, 3, 32, 16, 0, 335, 334, 1, 0, 0, 0, 335, 336, 1, 0, 0,
		0, 336, 337, 1, 0, 0, 0, 337, 338, 5, 46, 0, 0, 338, 29, 1, 0, 0, 0, 339,
		340, 5, 18, 0, 0, 340, 341, 3, 64, 32, 0, 341, 342, 5, 49, 0, 0, 342, 343,
		3, 2, 1, 0, 343, 31, 1, 0, 0, 0, 344, 345, 5, 19, 0, 0, 345, 346, 5, 49,
		0, 0, 346, 347, 3, 2, 1, 0, 347, 33, 1, 0, 0, 0, 348, 349, 5, 23, 0, 0,
		349, 350, 3, 64, 32, 0, 350, 351, 5, 45, 0, 0, 351, 352, 3, 2, 1, 0, 352,
		353, 5, 46, 0, 0, 353, 35, 1, 0, 0, 0, 354, 358, 5, 24, 0, 0, 355, 359,
		5, 41, 0, 0, 356, 359, 1, 0, 0, 0, 357, 359, 5, 38, 0, 0, 358, 355, 1,
		0, 0, 0, 358, 356, 1, 0, 0, 0, 358, 357, 1, 0, 0, 0, 359, 360, 1, 0, 0,
		0, 360, 361, 5, 27, 0, 0, 361, 362, 3, 38, 19, 0, 362, 363, 5, 45, 0, 0,
		363, 364, 3, 2, 1, 0, 364, 365, 5, 46, 0, 0, 365, 375, 1, 0, 0, 0, 366,
		367, 5, 24, 0, 0, 367, 368, 5, 41, 0, 0, 368, 369, 5, 27, 0, 0, 369, 370,
		3, 64, 32, 0, 370, 371, 5, 45, 0, 0, 371, 372, 3, 2, 1, 0, 372, 373, 5,
		46, 0, 0, 373, 375, 1, 0, 0, 0, 374, 354, 1, 0, 0, 0, 374, 366, 1, 0, 0,
		0, 375, 37, 1, 0, 0, 0, 376, 377, 3, 64, 32, 0, 377, 378, 5, 28, 0, 0,
		378, 379, 5, 28, 0, 0, 379, 380, 5, 28, 0, 0, 380, 381, 3, 64, 32, 0, 381,
		39, 1, 0, 0, 0, 382, 383, 5, 29, 0, 0, 383, 384, 3, 64, 32, 0, 384, 385,
		5, 16, 0, 0, 385, 386, 5, 45, 0, 0, 386, 387, 3, 2, 1, 0, 387, 388, 5,
		46, 0, 0, 388, 41, 1, 0, 0, 0, 389, 390, 5, 25, 0, 0, 390, 391, 5, 41,
		0, 0, 391, 392, 5, 43, 0, 0, 392, 395, 5, 44, 0, 0, 393, 394, 5, 26, 0,
		0, 394, 396, 3, 66, 33, 0, 395, 393, 1, 0, 0, 0, 395, 396, 1, 0, 0, 0,
		396, 397, 1, 0, 0, 0, 397, 398, 5, 45, 0, 0, 398, 399, 3, 2, 1, 0, 399,
		400, 5, 46, 0, 0, 400, 415, 1, 0, 0, 0, 401, 402, 5, 25, 0, 0, 402, 403,
		5, 41, 0, 0, 403, 404, 5, 43, 0, 0, 404, 405, 3, 44, 22, 0, 405, 408, 5,
		44, 0, 0, 406, 407, 5, 26, 0, 0, 407, 409, 3, 66, 33, 0, 408, 406, 1, 0,
		0, 0, 408, 409, 1, 0, 0, 0, 409, 410, 1, 0, 0, 0, 410, 411, 5, 45, 0, 0,
		411, 412, 3, 2, 1, 0, 412, 413, 5, 46, 0, 0, 413, 415, 1, 0, 0, 0, 414,
		389, 1, 0, 0, 0, 414, 401, 1, 0, 0, 0, 415, 43, 1, 0, 0, 0, 416, 417, 5,
		41, 0, 0, 417, 418, 5, 41, 0, 0, 418, 420, 5, 49, 0, 0, 419, 421, 5, 37,
		0, 0, 420, 419, 1, 0, 0, 0, 420, 421, 1, 0, 0, 0, 421, 422, 1, 0, 0, 0,
		422, 430, 3, 66, 33, 0, 423, 424, 5, 50, 0, 0, 424, 425, 5, 41, 0, 0, 425,
		426, 5, 41, 0, 0, 426, 427, 5, 49, 0, 0, 427, 429, 3, 66, 33, 0, 428, 423,
		1, 0, 0, 0, 429, 432, 1, 0, 0, 0, 430, 428, 1, 0, 0, 0, 430, 431, 1, 0,
		0, 0, 431, 563, 1, 0, 0, 0, 432, 430, 1, 0, 0, 0, 433, 434, 5, 38, 0, 0,
		434, 435, 5, 41, 0, 0, 435, 437, 5, 49, 0, 0, 436, 438, 5, 37, 0, 0, 437,
		436, 1, 0, 0, 0, 437, 438, 1, 0, 0, 0, 438, 439, 1, 0, 0, 0, 439, 447,
		3, 66, 33, 0, 440, 441, 5, 50, 0, 0, 441, 442, 5, 38, 0, 0, 442, 443, 5,
		41, 0, 0, 443, 444, 5, 49, 0, 0, 444, 446, 3, 66, 33, 0, 445, 440, 1, 0,
		0, 0, 446, 449, 1, 0, 0, 0, 447, 445, 1, 0, 0, 0, 447, 448, 1, 0, 0, 0,
		448, 563, 1, 0, 0, 0, 449, 447, 1, 0, 0, 0, 450, 451, 5, 41, 0, 0, 451,
		453, 5, 49, 0, 0, 452, 454, 5, 37, 0, 0, 453, 452, 1, 0, 0, 0, 453, 454,
		1, 0, 0, 0, 454, 455, 1, 0, 0, 0, 455, 465, 3, 66, 33, 0, 456, 457, 5,
		50, 0, 0, 457, 458, 5, 41, 0, 0, 458, 460, 5, 49, 0, 0, 459, 461, 5, 37,
		0, 0, 460, 459, 1, 0, 0, 0, 460, 461, 1, 0, 0, 0, 461, 462, 1, 0, 0, 0,
		462, 464, 3, 66, 33, 0, 463, 456, 1, 0, 0, 0, 464, 467, 1, 0, 0, 0, 465,
		463, 1, 0, 0, 0, 465, 466, 1, 0, 0, 0, 466, 563, 1, 0, 0, 0, 467, 465,
		1, 0, 0, 0, 468, 469, 5, 41, 0, 0, 469, 470, 5, 41, 0, 0, 470, 472, 5,
		49, 0, 0, 471, 473, 5, 37, 0, 0, 472, 471, 1, 0, 0, 0, 472, 473, 1, 0,
		0, 0, 473, 475, 1, 0, 0, 0, 474, 476, 5, 47, 0, 0, 475, 474, 1, 0, 0, 0,
		475, 476, 1, 0, 0, 0, 476, 477, 1, 0, 0, 0, 477, 479, 3, 66, 33, 0, 478,
		480, 5, 48, 0, 0, 479, 478, 1, 0, 0, 0, 479, 480, 1, 0, 0, 0, 480, 497,
		1, 0, 0, 0, 481, 482, 5, 50, 0, 0, 482, 483, 5, 41, 0, 0, 483, 484, 5,
		41, 0, 0, 484, 486, 5, 49, 0, 0, 485, 487, 5, 37, 0, 0, 486, 485, 1, 0,
		0, 0, 486, 487, 1, 0, 0, 0, 487, 489, 1, 0, 0, 0, 488, 490, 5, 47, 0, 0,
		489, 488, 1, 0, 0, 0, 489, 490, 1, 0, 0, 0, 490, 491, 1, 0, 0, 0, 491,
		493, 3, 66, 33, 0, 492, 494, 5, 48, 0, 0, 493, 492, 1, 0, 0, 0, 493, 494,
		1, 0, 0, 0, 494, 496, 1, 0, 0, 0, 495, 481, 1, 0, 0, 0, 496, 499, 1, 0,
		0, 0, 497, 495, 1, 0, 0, 0, 497, 498, 1, 0, 0, 0, 498, 563, 1, 0, 0, 0,
		499, 497, 1, 0, 0, 0, 500, 501, 5, 38, 0, 0, 501, 502, 5, 41, 0, 0, 502,
		504, 5, 49, 0, 0, 503, 505, 5, 37, 0, 0, 504, 503, 1, 0, 0, 0, 504, 505,
		1, 0, 0, 0, 505, 507, 1, 0, 0, 0, 506, 508, 5, 47, 0, 0, 507, 506, 1, 0,
		0, 0, 507, 508, 1, 0, 0, 0, 508, 509, 1, 0, 0, 0, 509, 511, 3, 66, 33,
		0, 510, 512, 5, 48, 0, 0, 511, 510, 1, 0, 0, 0, 511, 512, 1, 0, 0, 0, 512,
		529, 1, 0, 0, 0, 513, 514, 5, 50, 0, 0, 514, 515, 5, 38, 0, 0, 515, 516,
		5, 41, 0, 0, 516, 518, 5, 49, 0, 0, 517, 519, 5, 37, 0, 0, 518, 517, 1,
		0, 0, 0, 518, 519, 1, 0, 0, 0, 519, 521, 1, 0, 0, 0, 520, 522, 5, 47, 0,
		0, 521, 520, 1, 0, 0, 0, 521, 522, 1, 0, 0, 0, 522, 523, 1, 0, 0, 0, 523,
		525, 3, 66, 33, 0, 524, 526, 5, 48, 0, 0, 525, 524, 1, 0, 0, 0, 525, 526,
		1, 0, 0, 0, 526, 528, 1, 0, 0, 0, 527, 513, 1, 0, 0, 0, 528, 531, 1, 0,
		0, 0, 529, 527, 1, 0, 0, 0, 529, 530, 1, 0, 0, 0, 530, 563, 1, 0, 0, 0,
		531, 529, 1, 0, 0, 0, 532, 533, 5, 41, 0, 0, 533, 535, 5, 49, 0, 0, 534,
		536, 5, 37, 0, 0, 535, 534, 1, 0, 0, 0, 535, 536, 1, 0, 0, 0, 536, 538,
		1, 0, 0, 0, 537, 539, 5, 47, 0, 0, 538, 537, 1, 0, 0, 0, 538, 539, 1, 0,
		0, 0, 539, 540, 1, 0, 0, 0, 540, 542, 3, 66, 33, 0, 541, 543, 5, 48, 0,
		0, 542, 541, 1, 0, 0, 0, 542, 543, 1, 0, 0, 0, 543, 559, 1, 0, 0, 0, 544,
		545, 5, 50, 0, 0, 545, 546, 5, 41, 0, 0, 546, 548, 5, 49, 0, 0, 547, 549,
		5, 37, 0, 0, 548, 547, 1, 0, 0, 0, 548, 549, 1, 0, 0, 0, 549, 551, 1, 0,
		0, 0, 550, 552, 5, 47, 0, 0, 551, 550, 1, 0, 0, 0, 551, 552, 1, 0, 0, 0,
		552, 553, 1, 0, 0, 0, 553, 555, 3, 66, 33, 0, 554, 556, 5, 48, 0, 0, 555,
		554, 1, 0, 0, 0, 555, 556, 1, 0, 0, 0, 556, 558, 1, 0, 0, 0, 557, 544,
		1, 0, 0, 0, 558, 561, 1, 0, 0, 0, 559, 557, 1, 0, 0, 0, 559, 560, 1, 0,
		0, 0, 560, 563, 1, 0, 0, 0, 561, 559, 1, 0, 0, 0, 562, 416, 1, 0, 0, 0,
		562, 433, 1, 0, 0, 0, 562, 450, 1, 0, 0, 0, 562, 468, 1, 0, 0, 0, 562,
		500, 1, 0, 0, 0, 562, 532, 1, 0, 0, 0, 563, 45, 1, 0, 0, 0, 564, 565, 5,
		41, 0, 0, 565, 566, 5, 43, 0, 0, 566, 573, 5, 44, 0, 0, 567, 568, 5, 41,
		0, 0, 568, 569, 5, 43, 0, 0, 569, 570, 3, 48, 24, 0, 570, 571, 5, 44, 0,
		0, 571, 573, 1, 0, 0, 0, 572, 564, 1, 0, 0, 0, 572, 567, 1, 0, 0, 0, 573,
		47, 1, 0, 0, 0, 574, 576, 5, 36, 0, 0, 575, 574, 1, 0, 0, 0, 575, 576,
		1, 0, 0, 0, 576, 577, 1, 0, 0, 0, 577, 578, 5, 41, 0, 0, 578, 579, 5, 49,
		0, 0, 579, 589, 3, 64, 32, 0, 580, 582, 5, 50, 0, 0, 581, 583, 5, 36, 0,
		0, 582, 581, 1, 0, 0, 0, 582, 583, 1, 0, 0, 0, 583, 584, 1, 0, 0, 0, 584,
		585, 5, 41, 0, 0, 585, 586, 5, 49, 0, 0, 586, 588, 3, 64, 32, 0, 587, 580,
		1, 0, 0, 0, 588, 591, 1, 0, 0, 0, 589, 587, 1, 0, 0, 0, 589, 590, 1, 0,
		0, 0, 590, 607, 1, 0, 0, 0, 591, 589, 1, 0, 0, 0, 592, 594, 5, 36, 0, 0,
		593, 592, 1, 0, 0, 0, 593, 594, 1, 0, 0, 0, 594, 595, 1, 0, 0, 0, 595,
		603, 3, 64, 32, 0, 596, 598, 5, 50, 0, 0, 597, 599, 5, 36, 0, 0, 598, 597,
		1, 0, 0, 0, 598, 599, 1, 0, 0, 0, 599, 600, 1, 0, 0, 0, 600, 602, 3, 64,
		32, 0, 601, 596, 1, 0, 0, 0, 602, 605, 1, 0, 0, 0, 603, 601, 1, 0, 0, 0,
		603, 604, 1, 0, 0, 0, 604, 607, 1, 0, 0, 0, 605, 603, 1, 0, 0, 0, 606,
		575, 1, 0, 0, 0, 606, 593, 1, 0, 0, 0, 607, 49, 1, 0, 0, 0, 608, 609, 5,
		41, 0, 0, 609, 610, 5, 28, 0, 0, 610, 611, 5, 9, 0, 0, 611, 612, 5, 43,
		0, 0, 612, 613, 3, 64, 32, 0, 613, 614, 5, 44, 0, 0, 614, 671, 1, 0, 0,
		0, 615, 616, 5, 41, 0, 0, 616, 617, 5, 28, 0, 0, 617, 618, 5, 11, 0, 0,
		618, 619, 5, 43, 0, 0, 619, 671, 5, 44, 0, 0, 620, 621, 5, 41, 0, 0, 621,
		622, 5, 28, 0, 0, 622, 623, 5, 12, 0, 0, 623, 624, 5, 43, 0, 0, 624, 625,
		5, 15, 0, 0, 625, 626, 5, 49, 0, 0, 626, 627, 3, 64, 32, 0, 627, 628, 5,
		44, 0, 0, 628, 671, 1, 0, 0, 0, 629, 630, 5, 41, 0, 0, 630, 631, 5, 28,
		0, 0, 631, 671, 5, 13, 0, 0, 632, 633, 5, 41, 0, 0, 633, 634, 5, 28, 0,
		0, 634, 671, 5, 10, 0, 0, 635, 636, 5, 41, 0, 0, 636, 637, 5, 47, 0, 0,
		637, 638, 3, 64, 32, 0, 638, 639, 5, 48, 0, 0, 639, 671, 1, 0, 0, 0, 640,
		643, 5, 41, 0, 0, 641, 642, 5, 28, 0, 0, 642, 644, 5, 41, 0, 0, 643, 641,
		1, 0, 0, 0, 644, 645, 1, 0, 0, 0, 645, 643, 1, 0, 0, 0, 645, 646, 1, 0,
		0, 0, 646, 647, 1, 0, 0, 0, 647, 649, 5, 43, 0, 0, 648, 650, 3, 44, 22,
		0, 649, 648, 1, 0, 0, 0, 649, 650, 1, 0, 0, 0, 650, 651, 1, 0, 0, 0, 651,
		671, 5, 44, 0, 0, 652, 655, 5, 41, 0, 0, 653, 654, 5, 28, 0, 0, 654, 656,
		5, 41, 0, 0, 655, 653, 1, 0, 0, 0, 656, 657, 1, 0, 0, 0, 657, 655, 1, 0,
		0, 0, 657, 658, 1, 0, 0, 0, 658, 661, 1, 0, 0, 0, 659, 660, 5, 52, 0, 0,
		660, 662, 3, 64, 32, 0, 661, 659, 1, 0, 0, 0, 661, 662, 1, 0, 0, 0, 662,
		671, 1, 0, 0, 0, 663, 664, 5, 14, 0, 0, 664, 665, 5, 28, 0, 0, 665, 668,
		5, 41, 0, 0, 666, 667, 5, 52, 0, 0, 667, 669, 3, 64, 32, 0, 668, 666, 1,
		0, 0, 0, 668, 669, 1, 0, 0, 0, 669, 671, 1, 0, 0, 0, 670, 608, 1, 0, 0,
		0, 670, 615, 1, 0, 0, 0, 670, 620, 1, 0, 0, 0, 670, 629, 1, 0, 0, 0, 670,
		632, 1, 0, 0, 0, 670, 635, 1, 0, 0, 0, 670, 640, 1, 0, 0, 0, 670, 652,
		1, 0, 0, 0, 670, 663, 1, 0, 0, 0, 671, 51, 1, 0, 0, 0, 672, 676, 3, 58,
		29, 0, 673, 676, 3, 60, 30, 0, 674, 676, 3, 62, 31, 0, 675, 672, 1, 0,
		0, 0, 675, 673, 1, 0, 0, 0, 675, 674, 1, 0, 0, 0, 676, 53, 1, 0, 0, 0,
		677, 678, 5, 30, 0, 0, 678, 679, 5, 43, 0, 0, 679, 680, 3, 56, 28, 0, 680,
		681, 5, 44, 0, 0, 681, 55, 1, 0, 0, 0, 682, 687, 3, 64, 32, 0, 683, 684,
		5, 50, 0, 0, 684, 686, 3, 64, 32, 0, 685, 683, 1, 0, 0, 0, 686, 689, 1,
		0, 0, 0, 687, 685, 1, 0, 0, 0, 687, 688, 1, 0, 0, 0, 688, 57, 1, 0, 0,
		0, 689, 687, 1, 0, 0, 0, 690, 691, 5, 1, 0, 0, 691, 692, 5, 43, 0, 0, 692,
		693, 3, 64, 32, 0, 693, 694, 5, 44, 0, 0, 694, 59, 1, 0, 0, 0, 695, 696,
		5, 2, 0, 0, 696, 697, 5, 43, 0, 0, 697, 698, 3, 64, 32, 0, 698, 699, 5,
		44, 0, 0, 699, 61, 1, 0, 0, 0, 700, 701, 5, 3, 0, 0, 701, 702, 5, 43, 0,
		0, 702, 703, 3, 64, 32, 0, 703, 704, 5, 44, 0, 0, 704, 63, 1, 0, 0, 0,
		705, 706, 6, 32, -1, 0, 706, 707, 5, 42, 0, 0, 707, 734, 3, 64, 32, 17,
		708, 709, 5, 57, 0, 0, 709, 734, 3, 64, 32, 16, 710, 711, 5, 43, 0, 0,
		711, 712, 3, 64, 32, 0, 712, 713, 5, 44, 0, 0, 713, 734, 1, 0, 0, 0, 714,
		715, 5, 41, 0, 0, 715, 716, 5, 43, 0, 0, 716, 717, 3, 18, 9, 0, 717, 718,
		5, 44, 0, 0, 718, 734, 1, 0, 0, 0, 719, 721, 3, 46, 23, 0, 720, 722, 5,
		51, 0, 0, 721, 720, 1, 0, 0, 0, 721, 722, 1, 0, 0, 0, 722, 734, 1, 0, 0,
		0, 723, 725, 3, 50, 25, 0, 724, 726, 5, 51, 0, 0, 725, 724, 1, 0, 0, 0,
		725, 726, 1, 0, 0, 0, 726, 734, 1, 0, 0, 0, 727, 734, 3, 52, 26, 0, 728,
		734, 5, 39, 0, 0, 729, 734, 5, 40, 0, 0, 730, 734, 5, 41, 0, 0, 731, 734,
		5, 33, 0, 0, 732, 734, 7, 1, 0, 0, 733, 705, 1, 0, 0, 0, 733, 708, 1, 0,
		0, 0, 733, 710, 1, 0, 0, 0, 733, 714, 1, 0, 0, 0, 733, 719, 1, 0, 0, 0,
		733, 723, 1, 0, 0, 0, 733, 727, 1, 0, 0, 0, 733, 728, 1, 0, 0, 0, 733,
		729, 1, 0, 0, 0, 733, 730, 1, 0, 0, 0, 733, 731, 1, 0, 0, 0, 733, 732,
		1, 0, 0, 0, 734, 752, 1, 0, 0, 0, 735, 736, 10, 15, 0, 0, 736, 737, 7,
		2, 0, 0, 737, 751, 3, 64, 32, 16, 738, 739, 10, 14, 0, 0, 739, 740, 7,
		3, 0, 0, 740, 751, 3, 64, 32, 15, 741, 742, 10, 13, 0, 0, 742, 743, 7,
		4, 0, 0, 743, 751, 3, 64, 32, 14, 744, 745, 10, 12, 0, 0, 745, 746, 7,
		5, 0, 0, 746, 751, 3, 64, 32, 13, 747, 748, 10, 11, 0, 0, 748, 749, 7,
		6, 0, 0, 749, 751, 3, 64, 32, 12, 750, 735, 1, 0, 0, 0, 750, 738, 1, 0,
		0, 0, 750, 741, 1, 0, 0, 0, 750, 744, 1, 0, 0, 0, 750, 747, 1, 0, 0, 0,
		751, 754, 1, 0, 0, 0, 752, 750, 1, 0, 0, 0, 752, 753, 1, 0, 0, 0, 753,
		65, 1, 0, 0, 0, 754, 752, 1, 0, 0, 0, 755, 756, 7, 7, 0, 0, 756, 67, 1,
		0, 0, 0, 87, 74, 79, 83, 87, 98, 102, 106, 109, 113, 117, 121, 124, 126,
		137, 142, 145, 159, 178, 180, 183, 191, 198, 207, 213, 224, 265, 267, 301,
		313, 323, 331, 335, 358, 374, 395, 408, 414, 420, 430, 437, 447, 453, 460,
		465, 472, 475, 479, 486, 489, 493, 497, 504, 507, 511, 518, 521, 525, 529,
		535, 538, 542, 548, 551, 555, 559, 562, 572, 575, 582, 589, 593, 598, 603,
		606, 645, 649, 657, 661, 668, 670, 675, 687, 721, 725, 733, 750, 752,
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

// SFGrammarParserInit initializes any static state used to implement SFGrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSFGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SFGrammarParserInit() {
	staticData := &SFGrammarParserStaticData
	staticData.once.Do(sfgrammarParserInit)
}

// NewSFGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSFGrammarParser(input antlr.TokenStream) *SFGrammarParser {
	SFGrammarParserInit()
	this := new(SFGrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SFGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SFGrammar.g4"

	return this
}

// SFGrammarParser tokens.
const (
	SFGrammarParserEOF               = antlr.TokenEOF
	SFGrammarParserINT               = 1
	SFGrammarParserFLOAT             = 2
	SFGrammarParserSTRING            = 3
	SFGrammarParserBOOL              = 4
	SFGrammarParserCHAR              = 5
	SFGrammarParserIF                = 6
	SFGrammarParserSTRUCT            = 7
	SFGrammarParserMUTATING          = 8
	SFGrammarParserAPPEND            = 9
	SFGrammarParserCOUNT             = 10
	SFGrammarParserREMOVELAST        = 11
	SFGrammarParserREMOVE            = 12
	SFGrammarParserISEMPTY           = 13
	SFGrammarParserSELF              = 14
	SFGrammarParserAT                = 15
	SFGrammarParserELSE              = 16
	SFGrammarParserSWITCH            = 17
	SFGrammarParserCASE              = 18
	SFGrammarParserDEFAULT           = 19
	SFGrammarParserBREAK             = 20
	SFGrammarParserCONTINUE          = 21
	SFGrammarParserRETURN            = 22
	SFGrammarParserWHILE             = 23
	SFGrammarParserFOR               = 24
	SFGrammarParserFUNC              = 25
	SFGrammarParserARROW_FUNCTION    = 26
	SFGrammarParserIN                = 27
	SFGrammarParserDOT               = 28
	SFGrammarParserGUARD             = 29
	SFGrammarParserPRINT             = 30
	SFGrammarParserTRU               = 31
	SFGrammarParserFAL               = 32
	SFGrammarParserNIL               = 33
	SFGrammarParserDECLARATION_VAR   = 34
	SFGrammarParserDECLARATION_LET   = 35
	SFGrammarParserREFERENCE         = 36
	SFGrammarParserINOUT             = 37
	SFGrammarParserNOT_PARAM         = 38
	SFGrammarParserDIGIT_PRIMITIVE   = 39
	SFGrammarParserSTRING_PRIMITIVE  = 40
	SFGrammarParserID_PRIMITIVE      = 41
	SFGrammarParserNEGATION_OPERATOR = 42
	SFGrammarParserLPAREN            = 43
	SFGrammarParserRPAREN            = 44
	SFGrammarParserLBRACE            = 45
	SFGrammarParserRBRACE            = 46
	SFGrammarParserLBRACKET          = 47
	SFGrammarParserRBRACKET          = 48
	SFGrammarParserCOLON             = 49
	SFGrammarParserCOMMA             = 50
	SFGrammarParserSEMICOLON         = 51
	SFGrammarParserIS_               = 52
	SFGrammarParserPLUS_IS           = 53
	SFGrammarParserMINUS_IS          = 54
	SFGrammarParserQUESTION_MARK     = 55
	SFGrammarParserPLUS              = 56
	SFGrammarParserMINUS             = 57
	SFGrammarParserMULTIPLY          = 58
	SFGrammarParserDIVIDE            = 59
	SFGrammarParserMODULO            = 60
	SFGrammarParserEQUALS            = 61
	SFGrammarParserNOT_EQUALS        = 62
	SFGrammarParserGREATER           = 63
	SFGrammarParserGREATER_EQUALS    = 64
	SFGrammarParserLESS              = 65
	SFGrammarParserLESS_EQUALS       = 66
	SFGrammarParserAND               = 67
	SFGrammarParserOR                = 68
	SFGrammarParserWHITESPACE        = 69
	SFGrammarParserMULTI_COMMENT     = 70
	SFGrammarParserLINE_COMMENT      = 71
)

// SFGrammarParser rules.
const (
	SFGrammarParserRULE_start                = 0
	SFGrammarParserRULE_block                = 1
	SFGrammarParserRULE_stmts                = 2
	SFGrammarParserRULE_transferStmt         = 3
	SFGrammarParserRULE_structStmt           = 4
	SFGrammarParserRULE_structBlock          = 5
	SFGrammarParserRULE_structStmts          = 6
	SFGrammarParserRULE_declarationStructs   = 7
	SFGrammarParserRULE_functionStructs      = 8
	SFGrammarParserRULE_structCallList       = 9
	SFGrammarParserRULE_declaration          = 10
	SFGrammarParserRULE_type_declaration     = 11
	SFGrammarParserRULE_assignment           = 12
	SFGrammarParserRULE_ifstmt               = 13
	SFGrammarParserRULE_switchStmt           = 14
	SFGrammarParserRULE_caseBlock            = 15
	SFGrammarParserRULE_defaultBlock         = 16
	SFGrammarParserRULE_whileStmt            = 17
	SFGrammarParserRULE_forStmt              = 18
	SFGrammarParserRULE_forRange             = 19
	SFGrammarParserRULE_guardStmt            = 20
	SFGrammarParserRULE_functionStmt         = 21
	SFGrammarParserRULE_listFunctionParams   = 22
	SFGrammarParserRULE_callFunctionStmt     = 23
	SFGrammarParserRULE_listCallFunctionStmt = 24
	SFGrammarParserRULE_callBack             = 25
	SFGrammarParserRULE_embbededFunc         = 26
	SFGrammarParserRULE_printstmt            = 27
	SFGrammarParserRULE_exprList             = 28
	SFGrammarParserRULE_intstmt              = 29
	SFGrammarParserRULE_floatstmt            = 30
	SFGrammarParserRULE_stringstmt           = 31
	SFGrammarParserRULE_expr                 = 32
	SFGrammarParserRULE_type                 = 33
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_start
	return p
}

func InitEmptyStartContext(p *StartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_start
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitStart(s)
	}
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) Start_() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SFGrammarParserRULE_start)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Block()
	}
	{
		p.SetState(69)
		p.Match(SFGrammarParserEOF)
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmts() []IStmtsContext
	Stmts(i int) IStmtsContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStmts() []IStmtsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtsContext); ok {
			len++
		}
	}

	tst := make([]IStmtsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtsContext); ok {
			tst[i] = t.(IStmtsContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Stmts(i int) IStmtsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtsContext); ok {
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

	return t.(IStmtsContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SFGrammarParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2252239683790) != 0 {
		{
			p.SetState(71)
			p.Stmts()
		}

		p.SetState(76)
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

// IStmtsContext is an interface to support dynamic dispatch.
type IStmtsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Declaration() IDeclarationContext
	SEMICOLON() antlr.TerminalNode
	Assignment() IAssignmentContext
	EmbbededFunc() IEmbbededFuncContext
	Ifstmt() IIfstmtContext
	SwitchStmt() ISwitchStmtContext
	WhileStmt() IWhileStmtContext
	ForStmt() IForStmtContext
	GuardStmt() IGuardStmtContext
	TransferStmt() ITransferStmtContext
	FunctionStmt() IFunctionStmtContext
	Printstmt() IPrintstmtContext
	CallFunctionStmt() ICallFunctionStmtContext
	CallBack() ICallBackContext
	StructStmt() IStructStmtContext

	// IsStmtsContext differentiates from other interfaces.
	IsStmtsContext()
}

type StmtsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtsContext() *StmtsContext {
	var p = new(StmtsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_stmts
	return p
}

func InitEmptyStmtsContext(p *StmtsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_stmts
}

func (*StmtsContext) IsStmtsContext() {}

func NewStmtsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtsContext {
	var p = new(StmtsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_stmts

	return p
}

func (s *StmtsContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtsContext) Declaration() IDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *StmtsContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserSEMICOLON, 0)
}

func (s *StmtsContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StmtsContext) EmbbededFunc() IEmbbededFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmbbededFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmbbededFuncContext)
}

func (s *StmtsContext) Ifstmt() IIfstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfstmtContext)
}

func (s *StmtsContext) SwitchStmt() ISwitchStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchStmtContext)
}

func (s *StmtsContext) WhileStmt() IWhileStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStmtContext)
}

func (s *StmtsContext) ForStmt() IForStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStmtContext)
}

func (s *StmtsContext) GuardStmt() IGuardStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuardStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuardStmtContext)
}

func (s *StmtsContext) TransferStmt() ITransferStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITransferStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITransferStmtContext)
}

func (s *StmtsContext) FunctionStmt() IFunctionStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionStmtContext)
}

func (s *StmtsContext) Printstmt() IPrintstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintstmtContext)
}

func (s *StmtsContext) CallFunctionStmt() ICallFunctionStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallFunctionStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallFunctionStmtContext)
}

func (s *StmtsContext) CallBack() ICallBackContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallBackContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallBackContext)
}

func (s *StmtsContext) StructStmt() IStructStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructStmtContext)
}

func (s *StmtsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterStmts(s)
	}
}

func (s *StmtsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitStmts(s)
	}
}

func (s *StmtsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitStmts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) Stmts() (localctx IStmtsContext) {
	localctx = NewStmtsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SFGrammarParserRULE_stmts)
	var _la int

	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(77)
			p.Declaration()
		}
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserSEMICOLON {
			{
				p.SetState(78)
				p.Match(SFGrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.Assignment()
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserSEMICOLON {
			{
				p.SetState(82)
				p.Match(SFGrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(85)
			p.EmbbededFunc()
		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserSEMICOLON {
			{
				p.SetState(86)
				p.Match(SFGrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(89)
			p.Ifstmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(90)
			p.SwitchStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(91)
			p.WhileStmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(92)
			p.ForStmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(93)
			p.GuardStmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(94)
			p.TransferStmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(95)
			p.FunctionStmt()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(96)
			p.Printstmt()
		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserSEMICOLON {
			{
				p.SetState(97)
				p.Match(SFGrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(100)
			p.CallFunctionStmt()
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserSEMICOLON {
			{
				p.SetState(101)
				p.Match(SFGrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(104)
			p.CallBack()
		}
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserSEMICOLON {
			{
				p.SetState(105)
				p.Match(SFGrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(108)
			p.StructStmt()
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

// ITransferStmtContext is an interface to support dynamic dispatch.
type ITransferStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTransferStmtContext differentiates from other interfaces.
	IsTransferStmtContext()
}

type TransferStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransferStmtContext() *TransferStmtContext {
	var p = new(TransferStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_transferStmt
	return p
}

func InitEmptyTransferStmtContext(p *TransferStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_transferStmt
}

func (*TransferStmtContext) IsTransferStmtContext() {}

func NewTransferStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransferStmtContext {
	var p = new(TransferStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_transferStmt

	return p
}

func (s *TransferStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *TransferStmtContext) CopyAll(ctx *TransferStmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TransferStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransferStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ContinueStmtContext struct {
	TransferStmtContext
}

func NewContinueStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	InitEmptyTransferStmtContext(&p.TransferStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*TransferStmtContext))

	return p
}

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCONTINUE, 0)
}

func (s *ContinueStmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserSEMICOLON, 0)
}

func (s *ContinueStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterContinueStmt(s)
	}
}

func (s *ContinueStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitContinueStmt(s)
	}
}

func (s *ContinueStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitContinueStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type BreakStmtContext struct {
	TransferStmtContext
}

func NewBreakStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStmtContext {
	var p = new(BreakStmtContext)

	InitEmptyTransferStmtContext(&p.TransferStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*TransferStmtContext))

	return p
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserBREAK, 0)
}

func (s *BreakStmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserSEMICOLON, 0)
}

func (s *BreakStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterBreakStmt(s)
	}
}

func (s *BreakStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitBreakStmt(s)
	}
}

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStmtContext struct {
	TransferStmtContext
}

func NewReturnStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	InitEmptyTransferStmtContext(&p.TransferStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*TransferStmtContext))

	return p
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRETURN, 0)
}

func (s *ReturnStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturnStmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserSEMICOLON, 0)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) TransferStmt() (localctx ITransferStmtContext) {
	localctx = NewTransferStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SFGrammarParserRULE_transferStmt)
	var _la int

	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SFGrammarParserBREAK:
		localctx = NewBreakStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.Match(SFGrammarParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserSEMICOLON {
			{
				p.SetState(112)
				p.Match(SFGrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case SFGrammarParserCONTINUE:
		localctx = NewContinueStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(115)
			p.Match(SFGrammarParserCONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserSEMICOLON {
			{
				p.SetState(116)
				p.Match(SFGrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case SFGrammarParserRETURN:
		localctx = NewReturnStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(119)
			p.Match(SFGrammarParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(120)
				p.expr(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserSEMICOLON {
			{
				p.SetState(123)
				p.Match(SFGrammarParserSEMICOLON)
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

// IStructStmtContext is an interface to support dynamic dispatch.
type IStructStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRUCT() antlr.TerminalNode
	ID_PRIMITIVE() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	StructBlock() IStructBlockContext
	RBRACE() antlr.TerminalNode

	// IsStructStmtContext differentiates from other interfaces.
	IsStructStmtContext()
}

type StructStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructStmtContext() *StructStmtContext {
	var p = new(StructStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_structStmt
	return p
}

func InitEmptyStructStmtContext(p *StructStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_structStmt
}

func (*StructStmtContext) IsStructStmtContext() {}

func NewStructStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructStmtContext {
	var p = new(StructStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_structStmt

	return p
}

func (s *StructStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StructStmtContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserSTRUCT, 0)
}

func (s *StructStmtContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *StructStmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLBRACE, 0)
}

func (s *StructStmtContext) StructBlock() IStructBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructBlockContext)
}

func (s *StructStmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRBRACE, 0)
}

func (s *StructStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterStructStmt(s)
	}
}

func (s *StructStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitStructStmt(s)
	}
}

func (s *StructStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitStructStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) StructStmt() (localctx IStructStmtContext) {
	localctx = NewStructStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SFGrammarParserRULE_structStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(SFGrammarParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Match(SFGrammarParserID_PRIMITIVE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(130)
		p.Match(SFGrammarParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
		p.StructBlock()
	}
	{
		p.SetState(132)
		p.Match(SFGrammarParserRBRACE)
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

// IStructBlockContext is an interface to support dynamic dispatch.
type IStructBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStructStmts() []IStructStmtsContext
	StructStmts(i int) IStructStmtsContext

	// IsStructBlockContext differentiates from other interfaces.
	IsStructBlockContext()
}

type StructBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructBlockContext() *StructBlockContext {
	var p = new(StructBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_structBlock
	return p
}

func InitEmptyStructBlockContext(p *StructBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_structBlock
}

func (*StructBlockContext) IsStructBlockContext() {}

func NewStructBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructBlockContext {
	var p = new(StructBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_structBlock

	return p
}

func (s *StructBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *StructBlockContext) AllStructStmts() []IStructStmtsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructStmtsContext); ok {
			len++
		}
	}

	tst := make([]IStructStmtsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructStmtsContext); ok {
			tst[i] = t.(IStructStmtsContext)
			i++
		}
	}

	return tst
}

func (s *StructBlockContext) StructStmts(i int) IStructStmtsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructStmtsContext); ok {
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

	return t.(IStructStmtsContext)
}

func (s *StructBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterStructBlock(s)
	}
}

func (s *StructBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitStructBlock(s)
	}
}

func (s *StructBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitStructBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) StructBlock() (localctx IStructBlockContext) {
	localctx = NewStructBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SFGrammarParserRULE_structBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&51573162240) != 0 {
		{
			p.SetState(134)
			p.StructStmts()
		}

		p.SetState(139)
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

// IStructStmtsContext is an interface to support dynamic dispatch.
type IStructStmtsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DeclarationStructs() IDeclarationStructsContext
	SEMICOLON() antlr.TerminalNode
	FunctionStructs() IFunctionStructsContext

	// IsStructStmtsContext differentiates from other interfaces.
	IsStructStmtsContext()
}

type StructStmtsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructStmtsContext() *StructStmtsContext {
	var p = new(StructStmtsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_structStmts
	return p
}

func InitEmptyStructStmtsContext(p *StructStmtsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_structStmts
}

func (*StructStmtsContext) IsStructStmtsContext() {}

func NewStructStmtsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructStmtsContext {
	var p = new(StructStmtsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_structStmts

	return p
}

func (s *StructStmtsContext) GetParser() antlr.Parser { return s.parser }

func (s *StructStmtsContext) DeclarationStructs() IDeclarationStructsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationStructsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationStructsContext)
}

func (s *StructStmtsContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserSEMICOLON, 0)
}

func (s *StructStmtsContext) FunctionStructs() IFunctionStructsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionStructsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionStructsContext)
}

func (s *StructStmtsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructStmtsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructStmtsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterStructStmts(s)
	}
}

func (s *StructStmtsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitStructStmts(s)
	}
}

func (s *StructStmtsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitStructStmts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) StructStmts() (localctx IStructStmtsContext) {
	localctx = NewStructStmtsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SFGrammarParserRULE_structStmts)
	var _la int

	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SFGrammarParserDECLARATION_VAR, SFGrammarParserDECLARATION_LET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(140)
			p.DeclarationStructs()
		}
		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserSEMICOLON {
			{
				p.SetState(141)
				p.Match(SFGrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case SFGrammarParserMUTATING, SFGrammarParserFUNC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(144)
			p.FunctionStructs()
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

// IDeclarationStructsContext is an interface to support dynamic dispatch.
type IDeclarationStructsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclarationStructsContext differentiates from other interfaces.
	IsDeclarationStructsContext()
}

type DeclarationStructsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationStructsContext() *DeclarationStructsContext {
	var p = new(DeclarationStructsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_declarationStructs
	return p
}

func InitEmptyDeclarationStructsContext(p *DeclarationStructsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_declarationStructs
}

func (*DeclarationStructsContext) IsDeclarationStructsContext() {}

func NewDeclarationStructsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationStructsContext {
	var p = new(DeclarationStructsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_declarationStructs

	return p
}

func (s *DeclarationStructsContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationStructsContext) CopyAll(ctx *DeclarationStructsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DeclarationStructsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationStructsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructDeclarationWithValueAndTypeContext struct {
	DeclarationStructsContext
}

func NewStructDeclarationWithValueAndTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructDeclarationWithValueAndTypeContext {
	var p = new(StructDeclarationWithValueAndTypeContext)

	InitEmptyDeclarationStructsContext(&p.DeclarationStructsContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarationStructsContext))

	return p
}

func (s *StructDeclarationWithValueAndTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclarationWithValueAndTypeContext) Type_declaration() IType_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_declarationContext)
}

func (s *StructDeclarationWithValueAndTypeContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *StructDeclarationWithValueAndTypeContext) COLON() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOLON, 0)
}

func (s *StructDeclarationWithValueAndTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *StructDeclarationWithValueAndTypeContext) IS_() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserIS_, 0)
}

func (s *StructDeclarationWithValueAndTypeContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StructDeclarationWithValueAndTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterStructDeclarationWithValueAndType(s)
	}
}

func (s *StructDeclarationWithValueAndTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitStructDeclarationWithValueAndType(s)
	}
}

func (s *StructDeclarationWithValueAndTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitStructDeclarationWithValueAndType(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructDeclarationWithoutValueContext struct {
	DeclarationStructsContext
}

func NewStructDeclarationWithoutValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructDeclarationWithoutValueContext {
	var p = new(StructDeclarationWithoutValueContext)

	InitEmptyDeclarationStructsContext(&p.DeclarationStructsContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarationStructsContext))

	return p
}

func (s *StructDeclarationWithoutValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclarationWithoutValueContext) Type_declaration() IType_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_declarationContext)
}

func (s *StructDeclarationWithoutValueContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *StructDeclarationWithoutValueContext) COLON() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOLON, 0)
}

func (s *StructDeclarationWithoutValueContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *StructDeclarationWithoutValueContext) QUESTION_MARK() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserQUESTION_MARK, 0)
}

func (s *StructDeclarationWithoutValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterStructDeclarationWithoutValue(s)
	}
}

func (s *StructDeclarationWithoutValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitStructDeclarationWithoutValue(s)
	}
}

func (s *StructDeclarationWithoutValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitStructDeclarationWithoutValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructDeclarationVectorContext struct {
	DeclarationStructsContext
}

func NewStructDeclarationVectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructDeclarationVectorContext {
	var p = new(StructDeclarationVectorContext)

	InitEmptyDeclarationStructsContext(&p.DeclarationStructsContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarationStructsContext))

	return p
}

func (s *StructDeclarationVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclarationVectorContext) Type_declaration() IType_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_declarationContext)
}

func (s *StructDeclarationVectorContext) AllID_PRIMITIVE() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserID_PRIMITIVE)
}

func (s *StructDeclarationVectorContext) ID_PRIMITIVE(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, i)
}

func (s *StructDeclarationVectorContext) COLON() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOLON, 0)
}

func (s *StructDeclarationVectorContext) AllLBRACKET() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserLBRACKET)
}

func (s *StructDeclarationVectorContext) LBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLBRACKET, i)
}

func (s *StructDeclarationVectorContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *StructDeclarationVectorContext) AllRBRACKET() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserRBRACKET)
}

func (s *StructDeclarationVectorContext) RBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRBRACKET, i)
}

func (s *StructDeclarationVectorContext) IS_() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserIS_, 0)
}

func (s *StructDeclarationVectorContext) ExprList() IExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *StructDeclarationVectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterStructDeclarationVector(s)
	}
}

func (s *StructDeclarationVectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitStructDeclarationVector(s)
	}
}

func (s *StructDeclarationVectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitStructDeclarationVector(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructDeclarationImplicitValueContext struct {
	DeclarationStructsContext
}

func NewStructDeclarationImplicitValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructDeclarationImplicitValueContext {
	var p = new(StructDeclarationImplicitValueContext)

	InitEmptyDeclarationStructsContext(&p.DeclarationStructsContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarationStructsContext))

	return p
}

func (s *StructDeclarationImplicitValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclarationImplicitValueContext) Type_declaration() IType_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_declarationContext)
}

func (s *StructDeclarationImplicitValueContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *StructDeclarationImplicitValueContext) IS_() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserIS_, 0)
}

func (s *StructDeclarationImplicitValueContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StructDeclarationImplicitValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterStructDeclarationImplicitValue(s)
	}
}

func (s *StructDeclarationImplicitValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitStructDeclarationImplicitValue(s)
	}
}

func (s *StructDeclarationImplicitValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitStructDeclarationImplicitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) DeclarationStructs() (localctx IDeclarationStructsContext) {
	localctx = NewDeclarationStructsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SFGrammarParserRULE_declarationStructs)
	var _la int

	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStructDeclarationWithValueAndTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(147)
			p.Type_declaration()
		}
		{
			p.SetState(148)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.Type_()
		}
		{
			p.SetState(151)
			p.Match(SFGrammarParserIS_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)
			p.expr(0)
		}

	case 2:
		localctx = NewStructDeclarationWithoutValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)
			p.Type_declaration()
		}
		{
			p.SetState(155)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)
			p.Type_()
		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserQUESTION_MARK {
			{
				p.SetState(158)
				p.Match(SFGrammarParserQUESTION_MARK)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 3:
		localctx = NewStructDeclarationImplicitValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(161)
			p.Type_declaration()
		}
		{
			p.SetState(162)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)
			p.Match(SFGrammarParserIS_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)
			p.expr(0)
		}

	case 4:
		localctx = NewStructDeclarationVectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(166)
			p.Type_declaration()
		}
		{
			p.SetState(167)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.Match(SFGrammarParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(170)
			p.Type_()
		}
		{
			p.SetState(171)
			p.Match(SFGrammarParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)
			p.Match(SFGrammarParserIS_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case SFGrammarParserLBRACKET:
			{
				p.SetState(173)
				p.Match(SFGrammarParserLBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(174)
				p.ExprList()
			}
			{
				p.SetState(175)
				p.Match(SFGrammarParserRBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case SFGrammarParserID_PRIMITIVE:
			{
				p.SetState(177)
				p.Match(SFGrammarParserID_PRIMITIVE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
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

// IFunctionStructsContext is an interface to support dynamic dispatch.
type IFunctionStructsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunctionStructsContext differentiates from other interfaces.
	IsFunctionStructsContext()
}

type FunctionStructsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionStructsContext() *FunctionStructsContext {
	var p = new(FunctionStructsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_functionStructs
	return p
}

func InitEmptyFunctionStructsContext(p *FunctionStructsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_functionStructs
}

func (*FunctionStructsContext) IsFunctionStructsContext() {}

func NewFunctionStructsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionStructsContext {
	var p = new(FunctionStructsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_functionStructs

	return p
}

func (s *FunctionStructsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionStructsContext) CopyAll(ctx *FunctionStructsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FunctionStructsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionStructsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructFunctionWithoutParamsContext struct {
	FunctionStructsContext
}

func NewStructFunctionWithoutParamsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructFunctionWithoutParamsContext {
	var p = new(StructFunctionWithoutParamsContext)

	InitEmptyFunctionStructsContext(&p.FunctionStructsContext)
	p.parser = parser
	p.CopyAll(ctx.(*FunctionStructsContext))

	return p
}

func (s *StructFunctionWithoutParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructFunctionWithoutParamsContext) FUNC() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserFUNC, 0)
}

func (s *StructFunctionWithoutParamsContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *StructFunctionWithoutParamsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLPAREN, 0)
}

func (s *StructFunctionWithoutParamsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRPAREN, 0)
}

func (s *StructFunctionWithoutParamsContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLBRACE, 0)
}

func (s *StructFunctionWithoutParamsContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StructFunctionWithoutParamsContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRBRACE, 0)
}

func (s *StructFunctionWithoutParamsContext) MUTATING() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserMUTATING, 0)
}

func (s *StructFunctionWithoutParamsContext) ARROW_FUNCTION() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserARROW_FUNCTION, 0)
}

func (s *StructFunctionWithoutParamsContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *StructFunctionWithoutParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterStructFunctionWithoutParams(s)
	}
}

func (s *StructFunctionWithoutParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitStructFunctionWithoutParams(s)
	}
}

func (s *StructFunctionWithoutParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitStructFunctionWithoutParams(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructFunctionWithParamsContext struct {
	FunctionStructsContext
}

func NewStructFunctionWithParamsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructFunctionWithParamsContext {
	var p = new(StructFunctionWithParamsContext)

	InitEmptyFunctionStructsContext(&p.FunctionStructsContext)
	p.parser = parser
	p.CopyAll(ctx.(*FunctionStructsContext))

	return p
}

func (s *StructFunctionWithParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructFunctionWithParamsContext) FUNC() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserFUNC, 0)
}

func (s *StructFunctionWithParamsContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *StructFunctionWithParamsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLPAREN, 0)
}

func (s *StructFunctionWithParamsContext) ListFunctionParams() IListFunctionParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListFunctionParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListFunctionParamsContext)
}

func (s *StructFunctionWithParamsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRPAREN, 0)
}

func (s *StructFunctionWithParamsContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLBRACE, 0)
}

func (s *StructFunctionWithParamsContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StructFunctionWithParamsContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRBRACE, 0)
}

func (s *StructFunctionWithParamsContext) MUTATING() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserMUTATING, 0)
}

func (s *StructFunctionWithParamsContext) ARROW_FUNCTION() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserARROW_FUNCTION, 0)
}

func (s *StructFunctionWithParamsContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *StructFunctionWithParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterStructFunctionWithParams(s)
	}
}

func (s *StructFunctionWithParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitStructFunctionWithParams(s)
	}
}

func (s *StructFunctionWithParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitStructFunctionWithParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) FunctionStructs() (localctx IFunctionStructsContext) {
	localctx = NewFunctionStructsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SFGrammarParserRULE_functionStructs)
	var _la int

	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStructFunctionWithoutParamsContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserMUTATING {
			{
				p.SetState(182)
				p.Match(SFGrammarParserMUTATING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(185)
			p.Match(SFGrammarParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.Match(SFGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserARROW_FUNCTION {
			{
				p.SetState(189)
				p.Match(SFGrammarParserARROW_FUNCTION)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(190)
				p.Type_()
			}

		}
		{
			p.SetState(193)
			p.Match(SFGrammarParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)
			p.Block()
		}
		{
			p.SetState(195)
			p.Match(SFGrammarParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewStructFunctionWithParamsContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserMUTATING {
			{
				p.SetState(197)
				p.Match(SFGrammarParserMUTATING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(200)
			p.Match(SFGrammarParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(202)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)
			p.ListFunctionParams()
		}
		{
			p.SetState(204)
			p.Match(SFGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserARROW_FUNCTION {
			{
				p.SetState(205)
				p.Match(SFGrammarParserARROW_FUNCTION)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(206)
				p.Type_()
			}

		}
		{
			p.SetState(209)
			p.Match(SFGrammarParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)
			p.Block()
		}
		{
			p.SetState(211)
			p.Match(SFGrammarParserRBRACE)
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

// IStructCallListContext is an interface to support dynamic dispatch.
type IStructCallListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID_PRIMITIVE() []antlr.TerminalNode
	ID_PRIMITIVE(i int) antlr.TerminalNode
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsStructCallListContext differentiates from other interfaces.
	IsStructCallListContext()
}

type StructCallListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructCallListContext() *StructCallListContext {
	var p = new(StructCallListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_structCallList
	return p
}

func InitEmptyStructCallListContext(p *StructCallListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_structCallList
}

func (*StructCallListContext) IsStructCallListContext() {}

func NewStructCallListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructCallListContext {
	var p = new(StructCallListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_structCallList

	return p
}

func (s *StructCallListContext) GetParser() antlr.Parser { return s.parser }

func (s *StructCallListContext) AllID_PRIMITIVE() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserID_PRIMITIVE)
}

func (s *StructCallListContext) ID_PRIMITIVE(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, i)
}

func (s *StructCallListContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserCOLON)
}

func (s *StructCallListContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOLON, i)
}

func (s *StructCallListContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *StructCallListContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *StructCallListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserCOMMA)
}

func (s *StructCallListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOMMA, i)
}

func (s *StructCallListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructCallListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructCallListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterStructCallList(s)
	}
}

func (s *StructCallListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitStructCallList(s)
	}
}

func (s *StructCallListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitStructCallList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) StructCallList() (localctx IStructCallListContext) {
	localctx = NewStructCallListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SFGrammarParserRULE_structCallList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.Match(SFGrammarParserID_PRIMITIVE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(216)
		p.Match(SFGrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(217)
		p.expr(0)
	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SFGrammarParserCOMMA {
		{
			p.SetState(218)
			p.Match(SFGrammarParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(219)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(220)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(221)
			p.expr(0)
		}

		p.SetState(226)
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

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_declaration
	return p
}

func InitEmptyDeclarationContext(p *DeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_declaration
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) CopyAll(ctx *DeclarationContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VectorDeclarationContext struct {
	DeclarationContext
}

func NewVectorDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorDeclarationContext {
	var p = new(VectorDeclarationContext)

	InitEmptyDeclarationContext(&p.DeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarationContext))

	return p
}

func (s *VectorDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorDeclarationContext) Type_declaration() IType_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_declarationContext)
}

func (s *VectorDeclarationContext) AllID_PRIMITIVE() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserID_PRIMITIVE)
}

func (s *VectorDeclarationContext) ID_PRIMITIVE(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, i)
}

func (s *VectorDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOLON, 0)
}

func (s *VectorDeclarationContext) AllLBRACKET() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserLBRACKET)
}

func (s *VectorDeclarationContext) LBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLBRACKET, i)
}

func (s *VectorDeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *VectorDeclarationContext) AllRBRACKET() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserRBRACKET)
}

func (s *VectorDeclarationContext) RBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRBRACKET, i)
}

func (s *VectorDeclarationContext) IS_() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserIS_, 0)
}

func (s *VectorDeclarationContext) ExprList() IExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *VectorDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterVectorDeclaration(s)
	}
}

func (s *VectorDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitVectorDeclaration(s)
	}
}

func (s *VectorDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitVectorDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValueDeclarationContext struct {
	DeclarationContext
}

func NewValueDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueDeclarationContext {
	var p = new(ValueDeclarationContext)

	InitEmptyDeclarationContext(&p.DeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarationContext))

	return p
}

func (s *ValueDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueDeclarationContext) Type_declaration() IType_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_declarationContext)
}

func (s *ValueDeclarationContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *ValueDeclarationContext) IS_() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserIS_, 0)
}

func (s *ValueDeclarationContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ValueDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterValueDeclaration(s)
	}
}

func (s *ValueDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitValueDeclaration(s)
	}
}

func (s *ValueDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitValueDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructCreationContext struct {
	DeclarationContext
}

func NewStructCreationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructCreationContext {
	var p = new(StructCreationContext)

	InitEmptyDeclarationContext(&p.DeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarationContext))

	return p
}

func (s *StructCreationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructCreationContext) Type_declaration() IType_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_declarationContext)
}

func (s *StructCreationContext) AllID_PRIMITIVE() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserID_PRIMITIVE)
}

func (s *StructCreationContext) ID_PRIMITIVE(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, i)
}

func (s *StructCreationContext) IS_() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserIS_, 0)
}

func (s *StructCreationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLPAREN, 0)
}

func (s *StructCreationContext) StructCallList() IStructCallListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructCallListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructCallListContext)
}

func (s *StructCreationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRPAREN, 0)
}

func (s *StructCreationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterStructCreation(s)
	}
}

func (s *StructCreationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitStructCreation(s)
	}
}

func (s *StructCreationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitStructCreation(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeValueDeclarationContext struct {
	DeclarationContext
}

func NewTypeValueDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeValueDeclarationContext {
	var p = new(TypeValueDeclarationContext)

	InitEmptyDeclarationContext(&p.DeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarationContext))

	return p
}

func (s *TypeValueDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeValueDeclarationContext) Type_declaration() IType_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_declarationContext)
}

func (s *TypeValueDeclarationContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *TypeValueDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOLON, 0)
}

func (s *TypeValueDeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TypeValueDeclarationContext) IS_() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserIS_, 0)
}

func (s *TypeValueDeclarationContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TypeValueDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterTypeValueDeclaration(s)
	}
}

func (s *TypeValueDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitTypeValueDeclaration(s)
	}
}

func (s *TypeValueDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitTypeValueDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeOptionalValueDeclarationContext struct {
	DeclarationContext
}

func NewTypeOptionalValueDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeOptionalValueDeclarationContext {
	var p = new(TypeOptionalValueDeclarationContext)

	InitEmptyDeclarationContext(&p.DeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarationContext))

	return p
}

func (s *TypeOptionalValueDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeOptionalValueDeclarationContext) Type_declaration() IType_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_declarationContext)
}

func (s *TypeOptionalValueDeclarationContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *TypeOptionalValueDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOLON, 0)
}

func (s *TypeOptionalValueDeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TypeOptionalValueDeclarationContext) QUESTION_MARK() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserQUESTION_MARK, 0)
}

func (s *TypeOptionalValueDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterTypeOptionalValueDeclaration(s)
	}
}

func (s *TypeOptionalValueDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitTypeOptionalValueDeclaration(s)
	}
}

func (s *TypeOptionalValueDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitTypeOptionalValueDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SFGrammarParserRULE_declaration)
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStructCreationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(227)
			p.Type_declaration()
		}
		{
			p.SetState(228)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.Match(SFGrammarParserIS_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(230)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(232)
			p.StructCallList()
		}
		{
			p.SetState(233)
			p.Match(SFGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewTypeValueDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(235)
			p.Type_declaration()
		}
		{
			p.SetState(236)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(237)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.Type_()
		}
		{
			p.SetState(239)
			p.Match(SFGrammarParserIS_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(240)
			p.expr(0)
		}

	case 3:
		localctx = NewTypeOptionalValueDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(242)
			p.Type_declaration()
		}
		{
			p.SetState(243)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(244)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(245)
			p.Type_()
		}
		{
			p.SetState(246)
			p.Match(SFGrammarParserQUESTION_MARK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewValueDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(248)
			p.Type_declaration()
		}
		{
			p.SetState(249)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(250)
			p.Match(SFGrammarParserIS_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(251)
			p.expr(0)
		}

	case 5:
		localctx = NewVectorDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(253)
			p.Type_declaration()
		}
		{
			p.SetState(254)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(256)
			p.Match(SFGrammarParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(257)
			p.Type_()
		}
		{
			p.SetState(258)
			p.Match(SFGrammarParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(259)
			p.Match(SFGrammarParserIS_)
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

		switch p.GetTokenStream().LA(1) {
		case SFGrammarParserLBRACKET:
			{
				p.SetState(260)
				p.Match(SFGrammarParserLBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(261)
				p.ExprList()
			}
			{
				p.SetState(262)
				p.Match(SFGrammarParserRBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case SFGrammarParserID_PRIMITIVE:
			{
				p.SetState(264)
				p.Match(SFGrammarParserID_PRIMITIVE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
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

// IType_declarationContext is an interface to support dynamic dispatch.
type IType_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DECLARATION_VAR() antlr.TerminalNode
	DECLARATION_LET() antlr.TerminalNode

	// IsType_declarationContext differentiates from other interfaces.
	IsType_declarationContext()
}

type Type_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_declarationContext() *Type_declarationContext {
	var p = new(Type_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_type_declaration
	return p
}

func InitEmptyType_declarationContext(p *Type_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_type_declaration
}

func (*Type_declarationContext) IsType_declarationContext() {}

func NewType_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_declarationContext {
	var p = new(Type_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_type_declaration

	return p
}

func (s *Type_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_declarationContext) DECLARATION_VAR() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserDECLARATION_VAR, 0)
}

func (s *Type_declarationContext) DECLARATION_LET() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserDECLARATION_LET, 0)
}

func (s *Type_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterType_declaration(s)
	}
}

func (s *Type_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitType_declaration(s)
	}
}

func (s *Type_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitType_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) Type_declaration() (localctx IType_declarationContext) {
	localctx = NewType_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SFGrammarParserRULE_type_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(269)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SFGrammarParserDECLARATION_VAR || _la == SFGrammarParserDECLARATION_LET) {
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

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) CopyAll(ctx *AssignmentContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ValueAssignmentContext struct {
	AssignmentContext
}

func NewValueAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueAssignmentContext {
	var p = new(ValueAssignmentContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *ValueAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueAssignmentContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *ValueAssignmentContext) IS_() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserIS_, 0)
}

func (s *ValueAssignmentContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ValueAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterValueAssignment(s)
	}
}

func (s *ValueAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitValueAssignment(s)
	}
}

func (s *ValueAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitValueAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

type MinusAssignmentContext struct {
	AssignmentContext
}

func NewMinusAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MinusAssignmentContext {
	var p = new(MinusAssignmentContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *MinusAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinusAssignmentContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *MinusAssignmentContext) MINUS_IS() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserMINUS_IS, 0)
}

func (s *MinusAssignmentContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MinusAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterMinusAssignment(s)
	}
}

func (s *MinusAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitMinusAssignment(s)
	}
}

func (s *MinusAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitMinusAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorAssignmentContext struct {
	AssignmentContext
}

func NewVectorAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorAssignmentContext {
	var p = new(VectorAssignmentContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *VectorAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorAssignmentContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *VectorAssignmentContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLBRACKET, 0)
}

func (s *VectorAssignmentContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *VectorAssignmentContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *VectorAssignmentContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRBRACKET, 0)
}

func (s *VectorAssignmentContext) IS_() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserIS_, 0)
}

func (s *VectorAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterVectorAssignment(s)
	}
}

func (s *VectorAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitVectorAssignment(s)
	}
}

func (s *VectorAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitVectorAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorPlusAssignmentContext struct {
	AssignmentContext
}

func NewVectorPlusAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorPlusAssignmentContext {
	var p = new(VectorPlusAssignmentContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *VectorPlusAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorPlusAssignmentContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *VectorPlusAssignmentContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLBRACKET, 0)
}

func (s *VectorPlusAssignmentContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *VectorPlusAssignmentContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *VectorPlusAssignmentContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRBRACKET, 0)
}

func (s *VectorPlusAssignmentContext) PLUS_IS() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserPLUS_IS, 0)
}

func (s *VectorPlusAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterVectorPlusAssignment(s)
	}
}

func (s *VectorPlusAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitVectorPlusAssignment(s)
	}
}

func (s *VectorPlusAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitVectorPlusAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

type PlusAssignmentContext struct {
	AssignmentContext
}

func NewPlusAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlusAssignmentContext {
	var p = new(PlusAssignmentContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *PlusAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlusAssignmentContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *PlusAssignmentContext) PLUS_IS() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserPLUS_IS, 0)
}

func (s *PlusAssignmentContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PlusAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterPlusAssignment(s)
	}
}

func (s *PlusAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitPlusAssignment(s)
	}
}

func (s *PlusAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitPlusAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorMinusAssignmentContext struct {
	AssignmentContext
}

func NewVectorMinusAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorMinusAssignmentContext {
	var p = new(VectorMinusAssignmentContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *VectorMinusAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorMinusAssignmentContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *VectorMinusAssignmentContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLBRACKET, 0)
}

func (s *VectorMinusAssignmentContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *VectorMinusAssignmentContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *VectorMinusAssignmentContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRBRACKET, 0)
}

func (s *VectorMinusAssignmentContext) MINUS_IS() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserMINUS_IS, 0)
}

func (s *VectorMinusAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterVectorMinusAssignment(s)
	}
}

func (s *VectorMinusAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitVectorMinusAssignment(s)
	}
}

func (s *VectorMinusAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitVectorMinusAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SFGrammarParserRULE_assignment)
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		localctx = NewValueAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(271)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(272)
			p.Match(SFGrammarParserIS_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(273)
			p.expr(0)
		}

	case 2:
		localctx = NewPlusAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(274)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(275)
			p.Match(SFGrammarParserPLUS_IS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(276)
			p.expr(0)
		}

	case 3:
		localctx = NewMinusAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(277)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(278)
			p.Match(SFGrammarParserMINUS_IS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)
			p.expr(0)
		}

	case 4:
		localctx = NewVectorAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(280)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(281)
			p.Match(SFGrammarParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(282)
			p.expr(0)
		}
		{
			p.SetState(283)
			p.Match(SFGrammarParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)
			p.Match(SFGrammarParserIS_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(285)
			p.expr(0)
		}

	case 5:
		localctx = NewVectorMinusAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(287)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(288)
			p.Match(SFGrammarParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(289)
			p.expr(0)
		}
		{
			p.SetState(290)
			p.Match(SFGrammarParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(291)
			p.Match(SFGrammarParserMINUS_IS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(292)
			p.expr(0)
		}

	case 6:
		localctx = NewVectorPlusAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(294)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(295)
			p.Match(SFGrammarParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(296)
			p.expr(0)
		}
		{
			p.SetState(297)
			p.Match(SFGrammarParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(298)
			p.Match(SFGrammarParserPLUS_IS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)
			p.expr(0)
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

// IIfstmtContext is an interface to support dynamic dispatch.
type IIfstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIfstmtContext differentiates from other interfaces.
	IsIfstmtContext()
}

type IfstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfstmtContext() *IfstmtContext {
	var p = new(IfstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_ifstmt
	return p
}

func InitEmptyIfstmtContext(p *IfstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_ifstmt
}

func (*IfstmtContext) IsIfstmtContext() {}

func NewIfstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstmtContext {
	var p = new(IfstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_ifstmt

	return p
}

func (s *IfstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstmtContext) CopyAll(ctx *IfstmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IfstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ElseIfStmtContext struct {
	IfstmtContext
}

func NewElseIfStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ElseIfStmtContext {
	var p = new(ElseIfStmtContext)

	InitEmptyIfstmtContext(&p.IfstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfstmtContext))

	return p
}

func (s *ElseIfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserIF, 0)
}

func (s *ElseIfStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ElseIfStmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLBRACE, 0)
}

func (s *ElseIfStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseIfStmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRBRACE, 0)
}

func (s *ElseIfStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserELSE, 0)
}

func (s *ElseIfStmtContext) Ifstmt() IIfstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfstmtContext)
}

func (s *ElseIfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterElseIfStmt(s)
	}
}

func (s *ElseIfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitElseIfStmt(s)
	}
}

func (s *ElseIfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitElseIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfElseStmtContext struct {
	IfstmtContext
}

func NewIfElseStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfElseStmtContext {
	var p = new(IfElseStmtContext)

	InitEmptyIfstmtContext(&p.IfstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfstmtContext))

	return p
}

func (s *IfElseStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserIF, 0)
}

func (s *IfElseStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfElseStmtContext) AllLBRACE() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserLBRACE)
}

func (s *IfElseStmtContext) LBRACE(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLBRACE, i)
}

func (s *IfElseStmtContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfElseStmtContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
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

	return t.(IBlockContext)
}

func (s *IfElseStmtContext) AllRBRACE() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserRBRACE)
}

func (s *IfElseStmtContext) RBRACE(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRBRACE, i)
}

func (s *IfElseStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserELSE, 0)
}

func (s *IfElseStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterIfElseStmt(s)
	}
}

func (s *IfElseStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitIfElseStmt(s)
	}
}

func (s *IfElseStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitIfElseStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) Ifstmt() (localctx IIfstmtContext) {
	localctx = NewIfstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SFGrammarParserRULE_ifstmt)
	var _la int

	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIfElseStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(303)
			p.Match(SFGrammarParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(304)
			p.expr(0)
		}
		{
			p.SetState(305)
			p.Match(SFGrammarParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(306)
			p.Block()
		}
		{
			p.SetState(307)
			p.Match(SFGrammarParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserELSE {
			{
				p.SetState(308)
				p.Match(SFGrammarParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(309)
				p.Match(SFGrammarParserLBRACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(310)
				p.Block()
			}
			{
				p.SetState(311)
				p.Match(SFGrammarParserRBRACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		localctx = NewElseIfStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(315)
			p.Match(SFGrammarParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(316)
			p.expr(0)
		}
		{
			p.SetState(317)
			p.Match(SFGrammarParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(318)
			p.Block()
		}
		{
			p.SetState(319)
			p.Match(SFGrammarParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(320)
			p.Match(SFGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(321)
			p.Ifstmt()
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

// ISwitchStmtContext is an interface to support dynamic dispatch.
type ISwitchStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SWITCH() antlr.TerminalNode
	Expr() IExprContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllCaseBlock() []ICaseBlockContext
	CaseBlock(i int) ICaseBlockContext
	DefaultBlock() IDefaultBlockContext

	// IsSwitchStmtContext differentiates from other interfaces.
	IsSwitchStmtContext()
}

type SwitchStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchStmtContext() *SwitchStmtContext {
	var p = new(SwitchStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_switchStmt
	return p
}

func InitEmptySwitchStmtContext(p *SwitchStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_switchStmt
}

func (*SwitchStmtContext) IsSwitchStmtContext() {}

func NewSwitchStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStmtContext {
	var p = new(SwitchStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_switchStmt

	return p
}

func (s *SwitchStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStmtContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserSWITCH, 0)
}

func (s *SwitchStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchStmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLBRACE, 0)
}

func (s *SwitchStmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRBRACE, 0)
}

func (s *SwitchStmtContext) AllCaseBlock() []ICaseBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICaseBlockContext); ok {
			len++
		}
	}

	tst := make([]ICaseBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICaseBlockContext); ok {
			tst[i] = t.(ICaseBlockContext)
			i++
		}
	}

	return tst
}

func (s *SwitchStmtContext) CaseBlock(i int) ICaseBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseBlockContext); ok {
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

	return t.(ICaseBlockContext)
}

func (s *SwitchStmtContext) DefaultBlock() IDefaultBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultBlockContext)
}

func (s *SwitchStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterSwitchStmt(s)
	}
}

func (s *SwitchStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitSwitchStmt(s)
	}
}

func (s *SwitchStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitSwitchStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) SwitchStmt() (localctx ISwitchStmtContext) {
	localctx = NewSwitchStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SFGrammarParserRULE_switchStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(325)
		p.Match(SFGrammarParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(326)
		p.expr(0)
	}
	{
		p.SetState(327)
		p.Match(SFGrammarParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SFGrammarParserCASE {
		{
			p.SetState(328)
			p.CaseBlock()
		}

		p.SetState(333)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SFGrammarParserDEFAULT {
		{
			p.SetState(334)
			p.DefaultBlock()
		}

	}
	{
		p.SetState(337)
		p.Match(SFGrammarParserRBRACE)
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

// ICaseBlockContext is an interface to support dynamic dispatch.
type ICaseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	Expr() IExprContext
	COLON() antlr.TerminalNode
	Block() IBlockContext

	// IsCaseBlockContext differentiates from other interfaces.
	IsCaseBlockContext()
}

type CaseBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseBlockContext() *CaseBlockContext {
	var p = new(CaseBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_caseBlock
	return p
}

func InitEmptyCaseBlockContext(p *CaseBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_caseBlock
}

func (*CaseBlockContext) IsCaseBlockContext() {}

func NewCaseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseBlockContext {
	var p = new(CaseBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_caseBlock

	return p
}

func (s *CaseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseBlockContext) CASE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCASE, 0)
}

func (s *CaseBlockContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CaseBlockContext) COLON() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOLON, 0)
}

func (s *CaseBlockContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *CaseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterCaseBlock(s)
	}
}

func (s *CaseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitCaseBlock(s)
	}
}

func (s *CaseBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitCaseBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) CaseBlock() (localctx ICaseBlockContext) {
	localctx = NewCaseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SFGrammarParserRULE_caseBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.Match(SFGrammarParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(340)
		p.expr(0)
	}
	{
		p.SetState(341)
		p.Match(SFGrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(342)
		p.Block()
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

// IDefaultBlockContext is an interface to support dynamic dispatch.
type IDefaultBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFAULT() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Block() IBlockContext

	// IsDefaultBlockContext differentiates from other interfaces.
	IsDefaultBlockContext()
}

type DefaultBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultBlockContext() *DefaultBlockContext {
	var p = new(DefaultBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_defaultBlock
	return p
}

func InitEmptyDefaultBlockContext(p *DefaultBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_defaultBlock
}

func (*DefaultBlockContext) IsDefaultBlockContext() {}

func NewDefaultBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultBlockContext {
	var p = new(DefaultBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_defaultBlock

	return p
}

func (s *DefaultBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultBlockContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserDEFAULT, 0)
}

func (s *DefaultBlockContext) COLON() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOLON, 0)
}

func (s *DefaultBlockContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *DefaultBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterDefaultBlock(s)
	}
}

func (s *DefaultBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitDefaultBlock(s)
	}
}

func (s *DefaultBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitDefaultBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) DefaultBlock() (localctx IDefaultBlockContext) {
	localctx = NewDefaultBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SFGrammarParserRULE_defaultBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(344)
		p.Match(SFGrammarParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(345)
		p.Match(SFGrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(346)
		p.Block()
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

// IWhileStmtContext is an interface to support dynamic dispatch.
type IWhileStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expr() IExprContext
	LBRACE() antlr.TerminalNode
	Block() IBlockContext
	RBRACE() antlr.TerminalNode

	// IsWhileStmtContext differentiates from other interfaces.
	IsWhileStmtContext()
}

type WhileStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStmtContext() *WhileStmtContext {
	var p = new(WhileStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_whileStmt
	return p
}

func InitEmptyWhileStmtContext(p *WhileStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_whileStmt
}

func (*WhileStmtContext) IsWhileStmtContext() {}

func NewWhileStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStmtContext {
	var p = new(WhileStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_whileStmt

	return p
}

func (s *WhileStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStmtContext) WHILE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserWHILE, 0)
}

func (s *WhileStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhileStmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLBRACE, 0)
}

func (s *WhileStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileStmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRBRACE, 0)
}

func (s *WhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterWhileStmt(s)
	}
}

func (s *WhileStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitWhileStmt(s)
	}
}

func (s *WhileStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitWhileStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) WhileStmt() (localctx IWhileStmtContext) {
	localctx = NewWhileStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SFGrammarParserRULE_whileStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(348)
		p.Match(SFGrammarParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(349)
		p.expr(0)
	}
	{
		p.SetState(350)
		p.Match(SFGrammarParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(351)
		p.Block()
	}
	{
		p.SetState(352)
		p.Match(SFGrammarParserRBRACE)
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

// IForStmtContext is an interface to support dynamic dispatch.
type IForStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsForStmtContext differentiates from other interfaces.
	IsForStmtContext()
}

type ForStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStmtContext() *ForStmtContext {
	var p = new(ForStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_forStmt
	return p
}

func InitEmptyForStmtContext(p *ForStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_forStmt
}

func (*ForStmtContext) IsForStmtContext() {}

func NewForStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStmtContext {
	var p = new(ForStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_forStmt

	return p
}

func (s *ForStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStmtContext) CopyAll(ctx *ForStmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForRangeExprContext struct {
	ForStmtContext
}

func NewForRangeExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForRangeExprContext {
	var p = new(ForRangeExprContext)

	InitEmptyForStmtContext(&p.ForStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*ForStmtContext))

	return p
}

func (s *ForRangeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForRangeExprContext) FOR() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserFOR, 0)
}

func (s *ForRangeExprContext) IN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserIN, 0)
}

func (s *ForRangeExprContext) ForRange() IForRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForRangeContext)
}

func (s *ForRangeExprContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLBRACE, 0)
}

func (s *ForRangeExprContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForRangeExprContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRBRACE, 0)
}

func (s *ForRangeExprContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *ForRangeExprContext) NOT_PARAM() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserNOT_PARAM, 0)
}

func (s *ForRangeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterForRangeExpr(s)
	}
}

func (s *ForRangeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitForRangeExpr(s)
	}
}

func (s *ForRangeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitForRangeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForExprContext struct {
	ForStmtContext
}

func NewForExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForExprContext {
	var p = new(ForExprContext)

	InitEmptyForStmtContext(&p.ForStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*ForStmtContext))

	return p
}

func (s *ForExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForExprContext) FOR() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserFOR, 0)
}

func (s *ForExprContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *ForExprContext) IN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserIN, 0)
}

func (s *ForExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForExprContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLBRACE, 0)
}

func (s *ForExprContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForExprContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRBRACE, 0)
}

func (s *ForExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterForExpr(s)
	}
}

func (s *ForExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitForExpr(s)
	}
}

func (s *ForExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitForExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) ForStmt() (localctx IForStmtContext) {
	localctx = NewForStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SFGrammarParserRULE_forStmt)
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForRangeExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(354)
			p.Match(SFGrammarParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(358)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case SFGrammarParserID_PRIMITIVE:
			{
				p.SetState(355)
				p.Match(SFGrammarParserID_PRIMITIVE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case SFGrammarParserIN:

		case SFGrammarParserNOT_PARAM:
			{
				p.SetState(357)
				p.Match(SFGrammarParserNOT_PARAM)
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
			p.SetState(360)
			p.Match(SFGrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(361)
			p.ForRange()
		}
		{
			p.SetState(362)
			p.Match(SFGrammarParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(363)
			p.Block()
		}
		{
			p.SetState(364)
			p.Match(SFGrammarParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewForExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(366)
			p.Match(SFGrammarParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(367)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(368)
			p.Match(SFGrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(369)
			p.expr(0)
		}
		{
			p.SetState(370)
			p.Match(SFGrammarParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(371)
			p.Block()
		}
		{
			p.SetState(372)
			p.Match(SFGrammarParserRBRACE)
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

// IForRangeContext is an interface to support dynamic dispatch.
type IForRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the left rule contexts.
	GetLeft() IExprContext

	// GetRight returns the right rule contexts.
	GetRight() IExprContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExprContext)

	// SetRight sets the right rule contexts.
	SetRight(IExprContext)

	// Getter signatures
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsForRangeContext differentiates from other interfaces.
	IsForRangeContext()
}

type ForRangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	left   IExprContext
	right  IExprContext
}

func NewEmptyForRangeContext() *ForRangeContext {
	var p = new(ForRangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_forRange
	return p
}

func InitEmptyForRangeContext(p *ForRangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_forRange
}

func (*ForRangeContext) IsForRangeContext() {}

func NewForRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForRangeContext {
	var p = new(ForRangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_forRange

	return p
}

func (s *ForRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *ForRangeContext) GetLeft() IExprContext { return s.left }

func (s *ForRangeContext) GetRight() IExprContext { return s.right }

func (s *ForRangeContext) SetLeft(v IExprContext) { s.left = v }

func (s *ForRangeContext) SetRight(v IExprContext) { s.right = v }

func (s *ForRangeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserDOT)
}

func (s *ForRangeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserDOT, i)
}

func (s *ForRangeContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ForRangeContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ForRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterForRange(s)
	}
}

func (s *ForRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitForRange(s)
	}
}

func (s *ForRangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitForRange(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) ForRange() (localctx IForRangeContext) {
	localctx = NewForRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SFGrammarParserRULE_forRange)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)

		var _x = p.expr(0)

		localctx.(*ForRangeContext).left = _x
	}
	{
		p.SetState(377)
		p.Match(SFGrammarParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(378)
		p.Match(SFGrammarParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(379)
		p.Match(SFGrammarParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(380)

		var _x = p.expr(0)

		localctx.(*ForRangeContext).right = _x
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

// IGuardStmtContext is an interface to support dynamic dispatch.
type IGuardStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GUARD() antlr.TerminalNode
	Expr() IExprContext
	ELSE() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	Block() IBlockContext
	RBRACE() antlr.TerminalNode

	// IsGuardStmtContext differentiates from other interfaces.
	IsGuardStmtContext()
}

type GuardStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGuardStmtContext() *GuardStmtContext {
	var p = new(GuardStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_guardStmt
	return p
}

func InitEmptyGuardStmtContext(p *GuardStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_guardStmt
}

func (*GuardStmtContext) IsGuardStmtContext() {}

func NewGuardStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GuardStmtContext {
	var p = new(GuardStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_guardStmt

	return p
}

func (s *GuardStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *GuardStmtContext) GUARD() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserGUARD, 0)
}

func (s *GuardStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GuardStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserELSE, 0)
}

func (s *GuardStmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLBRACE, 0)
}

func (s *GuardStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *GuardStmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRBRACE, 0)
}

func (s *GuardStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GuardStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GuardStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterGuardStmt(s)
	}
}

func (s *GuardStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitGuardStmt(s)
	}
}

func (s *GuardStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitGuardStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) GuardStmt() (localctx IGuardStmtContext) {
	localctx = NewGuardStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SFGrammarParserRULE_guardStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(382)
		p.Match(SFGrammarParserGUARD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(383)
		p.expr(0)
	}
	{
		p.SetState(384)
		p.Match(SFGrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(385)
		p.Match(SFGrammarParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(386)
		p.Block()
	}
	{
		p.SetState(387)
		p.Match(SFGrammarParserRBRACE)
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

// IFunctionStmtContext is an interface to support dynamic dispatch.
type IFunctionStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunctionStmtContext differentiates from other interfaces.
	IsFunctionStmtContext()
}

type FunctionStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionStmtContext() *FunctionStmtContext {
	var p = new(FunctionStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_functionStmt
	return p
}

func InitEmptyFunctionStmtContext(p *FunctionStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_functionStmt
}

func (*FunctionStmtContext) IsFunctionStmtContext() {}

func NewFunctionStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionStmtContext {
	var p = new(FunctionStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_functionStmt

	return p
}

func (s *FunctionStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionStmtContext) CopyAll(ctx *FunctionStmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FunctionStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionWithParamsContext struct {
	FunctionStmtContext
}

func NewFunctionWithParamsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionWithParamsContext {
	var p = new(FunctionWithParamsContext)

	InitEmptyFunctionStmtContext(&p.FunctionStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*FunctionStmtContext))

	return p
}

func (s *FunctionWithParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionWithParamsContext) FUNC() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserFUNC, 0)
}

func (s *FunctionWithParamsContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *FunctionWithParamsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLPAREN, 0)
}

func (s *FunctionWithParamsContext) ListFunctionParams() IListFunctionParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListFunctionParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListFunctionParamsContext)
}

func (s *FunctionWithParamsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRPAREN, 0)
}

func (s *FunctionWithParamsContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLBRACE, 0)
}

func (s *FunctionWithParamsContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionWithParamsContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRBRACE, 0)
}

func (s *FunctionWithParamsContext) ARROW_FUNCTION() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserARROW_FUNCTION, 0)
}

func (s *FunctionWithParamsContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FunctionWithParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterFunctionWithParams(s)
	}
}

func (s *FunctionWithParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitFunctionWithParams(s)
	}
}

func (s *FunctionWithParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitFunctionWithParams(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionWithoutParamsContext struct {
	FunctionStmtContext
}

func NewFunctionWithoutParamsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionWithoutParamsContext {
	var p = new(FunctionWithoutParamsContext)

	InitEmptyFunctionStmtContext(&p.FunctionStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*FunctionStmtContext))

	return p
}

func (s *FunctionWithoutParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionWithoutParamsContext) FUNC() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserFUNC, 0)
}

func (s *FunctionWithoutParamsContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *FunctionWithoutParamsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLPAREN, 0)
}

func (s *FunctionWithoutParamsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRPAREN, 0)
}

func (s *FunctionWithoutParamsContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLBRACE, 0)
}

func (s *FunctionWithoutParamsContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionWithoutParamsContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRBRACE, 0)
}

func (s *FunctionWithoutParamsContext) ARROW_FUNCTION() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserARROW_FUNCTION, 0)
}

func (s *FunctionWithoutParamsContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FunctionWithoutParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterFunctionWithoutParams(s)
	}
}

func (s *FunctionWithoutParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitFunctionWithoutParams(s)
	}
}

func (s *FunctionWithoutParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitFunctionWithoutParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) FunctionStmt() (localctx IFunctionStmtContext) {
	localctx = NewFunctionStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SFGrammarParserRULE_functionStmt)
	var _la int

	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFunctionWithoutParamsContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(389)
			p.Match(SFGrammarParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(390)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(391)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(392)
			p.Match(SFGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(395)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserARROW_FUNCTION {
			{
				p.SetState(393)
				p.Match(SFGrammarParserARROW_FUNCTION)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(394)
				p.Type_()
			}

		}
		{
			p.SetState(397)
			p.Match(SFGrammarParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(398)
			p.Block()
		}
		{
			p.SetState(399)
			p.Match(SFGrammarParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewFunctionWithParamsContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(401)
			p.Match(SFGrammarParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(402)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(403)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(404)
			p.ListFunctionParams()
		}
		{
			p.SetState(405)
			p.Match(SFGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(408)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserARROW_FUNCTION {
			{
				p.SetState(406)
				p.Match(SFGrammarParserARROW_FUNCTION)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(407)
				p.Type_()
			}

		}
		{
			p.SetState(410)
			p.Match(SFGrammarParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(411)
			p.Block()
		}
		{
			p.SetState(412)
			p.Match(SFGrammarParserRBRACE)
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

// IListFunctionParamsContext is an interface to support dynamic dispatch.
type IListFunctionParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsListFunctionParamsContext differentiates from other interfaces.
	IsListFunctionParamsContext()
}

type ListFunctionParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListFunctionParamsContext() *ListFunctionParamsContext {
	var p = new(ListFunctionParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_listFunctionParams
	return p
}

func InitEmptyListFunctionParamsContext(p *ListFunctionParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_listFunctionParams
}

func (*ListFunctionParamsContext) IsListFunctionParamsContext() {}

func NewListFunctionParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListFunctionParamsContext {
	var p = new(ListFunctionParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_listFunctionParams

	return p
}

func (s *ListFunctionParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListFunctionParamsContext) CopyAll(ctx *ListFunctionParamsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ListFunctionParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListFunctionParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListFunctionParamsEIVectorContext struct {
	ListFunctionParamsContext
}

func NewListFunctionParamsEIVectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListFunctionParamsEIVectorContext {
	var p = new(ListFunctionParamsEIVectorContext)

	InitEmptyListFunctionParamsContext(&p.ListFunctionParamsContext)
	p.parser = parser
	p.CopyAll(ctx.(*ListFunctionParamsContext))

	return p
}

func (s *ListFunctionParamsEIVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListFunctionParamsEIVectorContext) AllID_PRIMITIVE() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserID_PRIMITIVE)
}

func (s *ListFunctionParamsEIVectorContext) ID_PRIMITIVE(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, i)
}

func (s *ListFunctionParamsEIVectorContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserCOLON)
}

func (s *ListFunctionParamsEIVectorContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOLON, i)
}

func (s *ListFunctionParamsEIVectorContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *ListFunctionParamsEIVectorContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
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

	return t.(ITypeContext)
}

func (s *ListFunctionParamsEIVectorContext) AllINOUT() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserINOUT)
}

func (s *ListFunctionParamsEIVectorContext) INOUT(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserINOUT, i)
}

func (s *ListFunctionParamsEIVectorContext) AllLBRACKET() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserLBRACKET)
}

func (s *ListFunctionParamsEIVectorContext) LBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLBRACKET, i)
}

func (s *ListFunctionParamsEIVectorContext) AllRBRACKET() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserRBRACKET)
}

func (s *ListFunctionParamsEIVectorContext) RBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRBRACKET, i)
}

func (s *ListFunctionParamsEIVectorContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserCOMMA)
}

func (s *ListFunctionParamsEIVectorContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOMMA, i)
}

func (s *ListFunctionParamsEIVectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterListFunctionParamsEIVector(s)
	}
}

func (s *ListFunctionParamsEIVectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitListFunctionParamsEIVector(s)
	}
}

func (s *ListFunctionParamsEIVectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitListFunctionParamsEIVector(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListFunctionParamsNEIVectorContext struct {
	ListFunctionParamsContext
}

func NewListFunctionParamsNEIVectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListFunctionParamsNEIVectorContext {
	var p = new(ListFunctionParamsNEIVectorContext)

	InitEmptyListFunctionParamsContext(&p.ListFunctionParamsContext)
	p.parser = parser
	p.CopyAll(ctx.(*ListFunctionParamsContext))

	return p
}

func (s *ListFunctionParamsNEIVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListFunctionParamsNEIVectorContext) AllNOT_PARAM() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserNOT_PARAM)
}

func (s *ListFunctionParamsNEIVectorContext) NOT_PARAM(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserNOT_PARAM, i)
}

func (s *ListFunctionParamsNEIVectorContext) AllID_PRIMITIVE() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserID_PRIMITIVE)
}

func (s *ListFunctionParamsNEIVectorContext) ID_PRIMITIVE(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, i)
}

func (s *ListFunctionParamsNEIVectorContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserCOLON)
}

func (s *ListFunctionParamsNEIVectorContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOLON, i)
}

func (s *ListFunctionParamsNEIVectorContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *ListFunctionParamsNEIVectorContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
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

	return t.(ITypeContext)
}

func (s *ListFunctionParamsNEIVectorContext) AllINOUT() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserINOUT)
}

func (s *ListFunctionParamsNEIVectorContext) INOUT(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserINOUT, i)
}

func (s *ListFunctionParamsNEIVectorContext) AllLBRACKET() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserLBRACKET)
}

func (s *ListFunctionParamsNEIVectorContext) LBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLBRACKET, i)
}

func (s *ListFunctionParamsNEIVectorContext) AllRBRACKET() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserRBRACKET)
}

func (s *ListFunctionParamsNEIVectorContext) RBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRBRACKET, i)
}

func (s *ListFunctionParamsNEIVectorContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserCOMMA)
}

func (s *ListFunctionParamsNEIVectorContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOMMA, i)
}

func (s *ListFunctionParamsNEIVectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterListFunctionParamsNEIVector(s)
	}
}

func (s *ListFunctionParamsNEIVectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitListFunctionParamsNEIVector(s)
	}
}

func (s *ListFunctionParamsNEIVectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitListFunctionParamsNEIVector(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListFunctionParamsEIContext struct {
	ListFunctionParamsContext
}

func NewListFunctionParamsEIContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListFunctionParamsEIContext {
	var p = new(ListFunctionParamsEIContext)

	InitEmptyListFunctionParamsContext(&p.ListFunctionParamsContext)
	p.parser = parser
	p.CopyAll(ctx.(*ListFunctionParamsContext))

	return p
}

func (s *ListFunctionParamsEIContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListFunctionParamsEIContext) AllID_PRIMITIVE() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserID_PRIMITIVE)
}

func (s *ListFunctionParamsEIContext) ID_PRIMITIVE(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, i)
}

func (s *ListFunctionParamsEIContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserCOLON)
}

func (s *ListFunctionParamsEIContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOLON, i)
}

func (s *ListFunctionParamsEIContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *ListFunctionParamsEIContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
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

	return t.(ITypeContext)
}

func (s *ListFunctionParamsEIContext) INOUT() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserINOUT, 0)
}

func (s *ListFunctionParamsEIContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserCOMMA)
}

func (s *ListFunctionParamsEIContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOMMA, i)
}

func (s *ListFunctionParamsEIContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterListFunctionParamsEI(s)
	}
}

func (s *ListFunctionParamsEIContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitListFunctionParamsEI(s)
	}
}

func (s *ListFunctionParamsEIContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitListFunctionParamsEI(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListFunctionParamsNEIContext struct {
	ListFunctionParamsContext
}

func NewListFunctionParamsNEIContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListFunctionParamsNEIContext {
	var p = new(ListFunctionParamsNEIContext)

	InitEmptyListFunctionParamsContext(&p.ListFunctionParamsContext)
	p.parser = parser
	p.CopyAll(ctx.(*ListFunctionParamsContext))

	return p
}

func (s *ListFunctionParamsNEIContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListFunctionParamsNEIContext) AllNOT_PARAM() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserNOT_PARAM)
}

func (s *ListFunctionParamsNEIContext) NOT_PARAM(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserNOT_PARAM, i)
}

func (s *ListFunctionParamsNEIContext) AllID_PRIMITIVE() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserID_PRIMITIVE)
}

func (s *ListFunctionParamsNEIContext) ID_PRIMITIVE(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, i)
}

func (s *ListFunctionParamsNEIContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserCOLON)
}

func (s *ListFunctionParamsNEIContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOLON, i)
}

func (s *ListFunctionParamsNEIContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *ListFunctionParamsNEIContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
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

	return t.(ITypeContext)
}

func (s *ListFunctionParamsNEIContext) INOUT() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserINOUT, 0)
}

func (s *ListFunctionParamsNEIContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserCOMMA)
}

func (s *ListFunctionParamsNEIContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOMMA, i)
}

func (s *ListFunctionParamsNEIContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterListFunctionParamsNEI(s)
	}
}

func (s *ListFunctionParamsNEIContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitListFunctionParamsNEI(s)
	}
}

func (s *ListFunctionParamsNEIContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitListFunctionParamsNEI(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListFunctionParamsBEIContext struct {
	ListFunctionParamsContext
}

func NewListFunctionParamsBEIContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListFunctionParamsBEIContext {
	var p = new(ListFunctionParamsBEIContext)

	InitEmptyListFunctionParamsContext(&p.ListFunctionParamsContext)
	p.parser = parser
	p.CopyAll(ctx.(*ListFunctionParamsContext))

	return p
}

func (s *ListFunctionParamsBEIContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListFunctionParamsBEIContext) AllID_PRIMITIVE() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserID_PRIMITIVE)
}

func (s *ListFunctionParamsBEIContext) ID_PRIMITIVE(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, i)
}

func (s *ListFunctionParamsBEIContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserCOLON)
}

func (s *ListFunctionParamsBEIContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOLON, i)
}

func (s *ListFunctionParamsBEIContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *ListFunctionParamsBEIContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
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

	return t.(ITypeContext)
}

func (s *ListFunctionParamsBEIContext) AllINOUT() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserINOUT)
}

func (s *ListFunctionParamsBEIContext) INOUT(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserINOUT, i)
}

func (s *ListFunctionParamsBEIContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserCOMMA)
}

func (s *ListFunctionParamsBEIContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOMMA, i)
}

func (s *ListFunctionParamsBEIContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterListFunctionParamsBEI(s)
	}
}

func (s *ListFunctionParamsBEIContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitListFunctionParamsBEI(s)
	}
}

func (s *ListFunctionParamsBEIContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitListFunctionParamsBEI(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListFunctionParamsBEIVectorContext struct {
	ListFunctionParamsContext
}

func NewListFunctionParamsBEIVectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListFunctionParamsBEIVectorContext {
	var p = new(ListFunctionParamsBEIVectorContext)

	InitEmptyListFunctionParamsContext(&p.ListFunctionParamsContext)
	p.parser = parser
	p.CopyAll(ctx.(*ListFunctionParamsContext))

	return p
}

func (s *ListFunctionParamsBEIVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListFunctionParamsBEIVectorContext) AllID_PRIMITIVE() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserID_PRIMITIVE)
}

func (s *ListFunctionParamsBEIVectorContext) ID_PRIMITIVE(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, i)
}

func (s *ListFunctionParamsBEIVectorContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserCOLON)
}

func (s *ListFunctionParamsBEIVectorContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOLON, i)
}

func (s *ListFunctionParamsBEIVectorContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *ListFunctionParamsBEIVectorContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
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

	return t.(ITypeContext)
}

func (s *ListFunctionParamsBEIVectorContext) AllINOUT() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserINOUT)
}

func (s *ListFunctionParamsBEIVectorContext) INOUT(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserINOUT, i)
}

func (s *ListFunctionParamsBEIVectorContext) AllLBRACKET() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserLBRACKET)
}

func (s *ListFunctionParamsBEIVectorContext) LBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLBRACKET, i)
}

func (s *ListFunctionParamsBEIVectorContext) AllRBRACKET() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserRBRACKET)
}

func (s *ListFunctionParamsBEIVectorContext) RBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRBRACKET, i)
}

func (s *ListFunctionParamsBEIVectorContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserCOMMA)
}

func (s *ListFunctionParamsBEIVectorContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOMMA, i)
}

func (s *ListFunctionParamsBEIVectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterListFunctionParamsBEIVector(s)
	}
}

func (s *ListFunctionParamsBEIVectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitListFunctionParamsBEIVector(s)
	}
}

func (s *ListFunctionParamsBEIVectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitListFunctionParamsBEIVector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) ListFunctionParams() (localctx IListFunctionParamsContext) {
	localctx = NewListFunctionParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SFGrammarParserRULE_listFunctionParams)
	var _la int

	p.SetState(562)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 65, p.GetParserRuleContext()) {
	case 1:
		localctx = NewListFunctionParamsEIContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(416)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(417)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(418)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(420)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserINOUT {
			{
				p.SetState(419)
				p.Match(SFGrammarParserINOUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(422)
			p.Type_()
		}
		p.SetState(430)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SFGrammarParserCOMMA {
			{
				p.SetState(423)
				p.Match(SFGrammarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(424)
				p.Match(SFGrammarParserID_PRIMITIVE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(425)
				p.Match(SFGrammarParserID_PRIMITIVE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(426)
				p.Match(SFGrammarParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(427)
				p.Type_()
			}

			p.SetState(432)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		localctx = NewListFunctionParamsNEIContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(433)
			p.Match(SFGrammarParserNOT_PARAM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(434)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(435)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(437)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserINOUT {
			{
				p.SetState(436)
				p.Match(SFGrammarParserINOUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(439)
			p.Type_()
		}
		p.SetState(447)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SFGrammarParserCOMMA {
			{
				p.SetState(440)
				p.Match(SFGrammarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(441)
				p.Match(SFGrammarParserNOT_PARAM)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(442)
				p.Match(SFGrammarParserID_PRIMITIVE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(443)
				p.Match(SFGrammarParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(444)
				p.Type_()
			}

			p.SetState(449)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 3:
		localctx = NewListFunctionParamsBEIContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(450)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(451)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(453)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserINOUT {
			{
				p.SetState(452)
				p.Match(SFGrammarParserINOUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(455)
			p.Type_()
		}
		p.SetState(465)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SFGrammarParserCOMMA {
			{
				p.SetState(456)
				p.Match(SFGrammarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(457)
				p.Match(SFGrammarParserID_PRIMITIVE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(458)
				p.Match(SFGrammarParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(460)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SFGrammarParserINOUT {
				{
					p.SetState(459)
					p.Match(SFGrammarParserINOUT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(462)
				p.Type_()
			}

			p.SetState(467)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 4:
		localctx = NewListFunctionParamsEIVectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(468)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(469)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(470)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(472)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserINOUT {
			{
				p.SetState(471)
				p.Match(SFGrammarParserINOUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(475)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserLBRACKET {
			{
				p.SetState(474)
				p.Match(SFGrammarParserLBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(477)
			p.Type_()
		}
		p.SetState(479)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserRBRACKET {
			{
				p.SetState(478)
				p.Match(SFGrammarParserRBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(497)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SFGrammarParserCOMMA {
			{
				p.SetState(481)
				p.Match(SFGrammarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(482)
				p.Match(SFGrammarParserID_PRIMITIVE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(483)
				p.Match(SFGrammarParserID_PRIMITIVE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(484)
				p.Match(SFGrammarParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(486)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SFGrammarParserINOUT {
				{
					p.SetState(485)
					p.Match(SFGrammarParserINOUT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(489)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SFGrammarParserLBRACKET {
				{
					p.SetState(488)
					p.Match(SFGrammarParserLBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(491)
				p.Type_()
			}
			p.SetState(493)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SFGrammarParserRBRACKET {
				{
					p.SetState(492)
					p.Match(SFGrammarParserRBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}

			p.SetState(499)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 5:
		localctx = NewListFunctionParamsNEIVectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(500)
			p.Match(SFGrammarParserNOT_PARAM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(501)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(502)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(504)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserINOUT {
			{
				p.SetState(503)
				p.Match(SFGrammarParserINOUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(507)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserLBRACKET {
			{
				p.SetState(506)
				p.Match(SFGrammarParserLBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(509)
			p.Type_()
		}
		p.SetState(511)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserRBRACKET {
			{
				p.SetState(510)
				p.Match(SFGrammarParserRBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(529)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SFGrammarParserCOMMA {
			{
				p.SetState(513)
				p.Match(SFGrammarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(514)
				p.Match(SFGrammarParserNOT_PARAM)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(515)
				p.Match(SFGrammarParserID_PRIMITIVE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(516)
				p.Match(SFGrammarParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(518)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SFGrammarParserINOUT {
				{
					p.SetState(517)
					p.Match(SFGrammarParserINOUT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(521)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SFGrammarParserLBRACKET {
				{
					p.SetState(520)
					p.Match(SFGrammarParserLBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(523)
				p.Type_()
			}
			p.SetState(525)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SFGrammarParserRBRACKET {
				{
					p.SetState(524)
					p.Match(SFGrammarParserRBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}

			p.SetState(531)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 6:
		localctx = NewListFunctionParamsBEIVectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(532)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(533)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(535)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserINOUT {
			{
				p.SetState(534)
				p.Match(SFGrammarParserINOUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(538)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserLBRACKET {
			{
				p.SetState(537)
				p.Match(SFGrammarParserLBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(540)
			p.Type_()
		}
		p.SetState(542)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserRBRACKET {
			{
				p.SetState(541)
				p.Match(SFGrammarParserRBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(559)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SFGrammarParserCOMMA {
			{
				p.SetState(544)
				p.Match(SFGrammarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(545)
				p.Match(SFGrammarParserID_PRIMITIVE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(546)
				p.Match(SFGrammarParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(548)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SFGrammarParserINOUT {
				{
					p.SetState(547)
					p.Match(SFGrammarParserINOUT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(551)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SFGrammarParserLBRACKET {
				{
					p.SetState(550)
					p.Match(SFGrammarParserLBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(553)
				p.Type_()
			}
			p.SetState(555)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SFGrammarParserRBRACKET {
				{
					p.SetState(554)
					p.Match(SFGrammarParserRBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}

			p.SetState(561)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
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

// ICallFunctionStmtContext is an interface to support dynamic dispatch.
type ICallFunctionStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCallFunctionStmtContext differentiates from other interfaces.
	IsCallFunctionStmtContext()
}

type CallFunctionStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallFunctionStmtContext() *CallFunctionStmtContext {
	var p = new(CallFunctionStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_callFunctionStmt
	return p
}

func InitEmptyCallFunctionStmtContext(p *CallFunctionStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_callFunctionStmt
}

func (*CallFunctionStmtContext) IsCallFunctionStmtContext() {}

func NewCallFunctionStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallFunctionStmtContext {
	var p = new(CallFunctionStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_callFunctionStmt

	return p
}

func (s *CallFunctionStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *CallFunctionStmtContext) CopyAll(ctx *CallFunctionStmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *CallFunctionStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallFunctionStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CallFunctionWithoutParamsContext struct {
	CallFunctionStmtContext
}

func NewCallFunctionWithoutParamsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallFunctionWithoutParamsContext {
	var p = new(CallFunctionWithoutParamsContext)

	InitEmptyCallFunctionStmtContext(&p.CallFunctionStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*CallFunctionStmtContext))

	return p
}

func (s *CallFunctionWithoutParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallFunctionWithoutParamsContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *CallFunctionWithoutParamsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLPAREN, 0)
}

func (s *CallFunctionWithoutParamsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRPAREN, 0)
}

func (s *CallFunctionWithoutParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterCallFunctionWithoutParams(s)
	}
}

func (s *CallFunctionWithoutParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitCallFunctionWithoutParams(s)
	}
}

func (s *CallFunctionWithoutParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitCallFunctionWithoutParams(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallFunctionWithParamsContext struct {
	CallFunctionStmtContext
}

func NewCallFunctionWithParamsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallFunctionWithParamsContext {
	var p = new(CallFunctionWithParamsContext)

	InitEmptyCallFunctionStmtContext(&p.CallFunctionStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*CallFunctionStmtContext))

	return p
}

func (s *CallFunctionWithParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallFunctionWithParamsContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *CallFunctionWithParamsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLPAREN, 0)
}

func (s *CallFunctionWithParamsContext) ListCallFunctionStmt() IListCallFunctionStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListCallFunctionStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListCallFunctionStmtContext)
}

func (s *CallFunctionWithParamsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRPAREN, 0)
}

func (s *CallFunctionWithParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterCallFunctionWithParams(s)
	}
}

func (s *CallFunctionWithParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitCallFunctionWithParams(s)
	}
}

func (s *CallFunctionWithParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitCallFunctionWithParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) CallFunctionStmt() (localctx ICallFunctionStmtContext) {
	localctx = NewCallFunctionStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SFGrammarParserRULE_callFunctionStmt)
	p.SetState(572)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 66, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCallFunctionWithoutParamsContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(564)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(565)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(566)
			p.Match(SFGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewCallFunctionWithParamsContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(567)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(568)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(569)
			p.ListCallFunctionStmt()
		}
		{
			p.SetState(570)
			p.Match(SFGrammarParserRPAREN)
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

// IListCallFunctionStmtContext is an interface to support dynamic dispatch.
type IListCallFunctionStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsListCallFunctionStmtContext differentiates from other interfaces.
	IsListCallFunctionStmtContext()
}

type ListCallFunctionStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListCallFunctionStmtContext() *ListCallFunctionStmtContext {
	var p = new(ListCallFunctionStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_listCallFunctionStmt
	return p
}

func InitEmptyListCallFunctionStmtContext(p *ListCallFunctionStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_listCallFunctionStmt
}

func (*ListCallFunctionStmtContext) IsListCallFunctionStmtContext() {}

func NewListCallFunctionStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListCallFunctionStmtContext {
	var p = new(ListCallFunctionStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_listCallFunctionStmt

	return p
}

func (s *ListCallFunctionStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ListCallFunctionStmtContext) CopyAll(ctx *ListCallFunctionStmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ListCallFunctionStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListCallFunctionStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListCallFunctionStmtEIContext struct {
	ListCallFunctionStmtContext
}

func NewListCallFunctionStmtEIContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListCallFunctionStmtEIContext {
	var p = new(ListCallFunctionStmtEIContext)

	InitEmptyListCallFunctionStmtContext(&p.ListCallFunctionStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*ListCallFunctionStmtContext))

	return p
}

func (s *ListCallFunctionStmtEIContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListCallFunctionStmtEIContext) AllID_PRIMITIVE() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserID_PRIMITIVE)
}

func (s *ListCallFunctionStmtEIContext) ID_PRIMITIVE(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, i)
}

func (s *ListCallFunctionStmtEIContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserCOLON)
}

func (s *ListCallFunctionStmtEIContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOLON, i)
}

func (s *ListCallFunctionStmtEIContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ListCallFunctionStmtEIContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ListCallFunctionStmtEIContext) AllREFERENCE() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserREFERENCE)
}

func (s *ListCallFunctionStmtEIContext) REFERENCE(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserREFERENCE, i)
}

func (s *ListCallFunctionStmtEIContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserCOMMA)
}

func (s *ListCallFunctionStmtEIContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOMMA, i)
}

func (s *ListCallFunctionStmtEIContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterListCallFunctionStmtEI(s)
	}
}

func (s *ListCallFunctionStmtEIContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitListCallFunctionStmtEI(s)
	}
}

func (s *ListCallFunctionStmtEIContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitListCallFunctionStmtEI(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListCallFunctionStmtNEIContext struct {
	ListCallFunctionStmtContext
}

func NewListCallFunctionStmtNEIContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListCallFunctionStmtNEIContext {
	var p = new(ListCallFunctionStmtNEIContext)

	InitEmptyListCallFunctionStmtContext(&p.ListCallFunctionStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*ListCallFunctionStmtContext))

	return p
}

func (s *ListCallFunctionStmtNEIContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListCallFunctionStmtNEIContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ListCallFunctionStmtNEIContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ListCallFunctionStmtNEIContext) AllREFERENCE() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserREFERENCE)
}

func (s *ListCallFunctionStmtNEIContext) REFERENCE(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserREFERENCE, i)
}

func (s *ListCallFunctionStmtNEIContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserCOMMA)
}

func (s *ListCallFunctionStmtNEIContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOMMA, i)
}

func (s *ListCallFunctionStmtNEIContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterListCallFunctionStmtNEI(s)
	}
}

func (s *ListCallFunctionStmtNEIContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitListCallFunctionStmtNEI(s)
	}
}

func (s *ListCallFunctionStmtNEIContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitListCallFunctionStmtNEI(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) ListCallFunctionStmt() (localctx IListCallFunctionStmtContext) {
	localctx = NewListCallFunctionStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SFGrammarParserRULE_listCallFunctionStmt)
	var _la int

	p.SetState(606)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 73, p.GetParserRuleContext()) {
	case 1:
		localctx = NewListCallFunctionStmtEIContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(575)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserREFERENCE {
			{
				p.SetState(574)
				p.Match(SFGrammarParserREFERENCE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(577)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(578)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(579)
			p.expr(0)
		}
		p.SetState(589)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SFGrammarParserCOMMA {
			{
				p.SetState(580)
				p.Match(SFGrammarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(582)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SFGrammarParserREFERENCE {
				{
					p.SetState(581)
					p.Match(SFGrammarParserREFERENCE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(584)
				p.Match(SFGrammarParserID_PRIMITIVE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(585)
				p.Match(SFGrammarParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(586)
				p.expr(0)
			}

			p.SetState(591)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		localctx = NewListCallFunctionStmtNEIContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(593)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserREFERENCE {
			{
				p.SetState(592)
				p.Match(SFGrammarParserREFERENCE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(595)
			p.expr(0)
		}
		p.SetState(603)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SFGrammarParserCOMMA {
			{
				p.SetState(596)
				p.Match(SFGrammarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(598)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SFGrammarParserREFERENCE {
				{
					p.SetState(597)
					p.Match(SFGrammarParserREFERENCE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(600)
				p.expr(0)
			}

			p.SetState(605)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
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

// ICallBackContext is an interface to support dynamic dispatch.
type ICallBackContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCallBackContext differentiates from other interfaces.
	IsCallBackContext()
}

type CallBackContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallBackContext() *CallBackContext {
	var p = new(CallBackContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_callBack
	return p
}

func InitEmptyCallBackContext(p *CallBackContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_callBack
}

func (*CallBackContext) IsCallBackContext() {}

func NewCallBackContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallBackContext {
	var p = new(CallBackContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_callBack

	return p
}

func (s *CallBackContext) GetParser() antlr.Parser { return s.parser }

func (s *CallBackContext) CopyAll(ctx *CallBackContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *CallBackContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallBackContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AccessVectorContext struct {
	CallBackContext
}

func NewAccessVectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AccessVectorContext {
	var p = new(AccessVectorContext)

	InitEmptyCallBackContext(&p.CallBackContext)
	p.parser = parser
	p.CopyAll(ctx.(*CallBackContext))

	return p
}

func (s *AccessVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessVectorContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *AccessVectorContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLBRACKET, 0)
}

func (s *AccessVectorContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AccessVectorContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRBRACKET, 0)
}

func (s *AccessVectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterAccessVector(s)
	}
}

func (s *AccessVectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitAccessVector(s)
	}
}

func (s *AccessVectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitAccessVector(s)

	default:
		return t.VisitChildren(s)
	}
}

type AppendVectorContext struct {
	CallBackContext
}

func NewAppendVectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AppendVectorContext {
	var p = new(AppendVectorContext)

	InitEmptyCallBackContext(&p.CallBackContext)
	p.parser = parser
	p.CopyAll(ctx.(*CallBackContext))

	return p
}

func (s *AppendVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendVectorContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *AppendVectorContext) DOT() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserDOT, 0)
}

func (s *AppendVectorContext) APPEND() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserAPPEND, 0)
}

func (s *AppendVectorContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLPAREN, 0)
}

func (s *AppendVectorContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AppendVectorContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRPAREN, 0)
}

func (s *AppendVectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterAppendVector(s)
	}
}

func (s *AppendVectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitAppendVector(s)
	}
}

func (s *AppendVectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitAppendVector(s)

	default:
		return t.VisitChildren(s)
	}
}

type RemoveLastVectorContext struct {
	CallBackContext
}

func NewRemoveLastVectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RemoveLastVectorContext {
	var p = new(RemoveLastVectorContext)

	InitEmptyCallBackContext(&p.CallBackContext)
	p.parser = parser
	p.CopyAll(ctx.(*CallBackContext))

	return p
}

func (s *RemoveLastVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RemoveLastVectorContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *RemoveLastVectorContext) DOT() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserDOT, 0)
}

func (s *RemoveLastVectorContext) REMOVELAST() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserREMOVELAST, 0)
}

func (s *RemoveLastVectorContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLPAREN, 0)
}

func (s *RemoveLastVectorContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRPAREN, 0)
}

func (s *RemoveLastVectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterRemoveLastVector(s)
	}
}

func (s *RemoveLastVectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitRemoveLastVector(s)
	}
}

func (s *RemoveLastVectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitRemoveLastVector(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructCallFunctionContext struct {
	CallBackContext
}

func NewStructCallFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructCallFunctionContext {
	var p = new(StructCallFunctionContext)

	InitEmptyCallBackContext(&p.CallBackContext)
	p.parser = parser
	p.CopyAll(ctx.(*CallBackContext))

	return p
}

func (s *StructCallFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructCallFunctionContext) AllID_PRIMITIVE() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserID_PRIMITIVE)
}

func (s *StructCallFunctionContext) ID_PRIMITIVE(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, i)
}

func (s *StructCallFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLPAREN, 0)
}

func (s *StructCallFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRPAREN, 0)
}

func (s *StructCallFunctionContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserDOT)
}

func (s *StructCallFunctionContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserDOT, i)
}

func (s *StructCallFunctionContext) ListFunctionParams() IListFunctionParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListFunctionParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListFunctionParamsContext)
}

func (s *StructCallFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterStructCallFunction(s)
	}
}

func (s *StructCallFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitStructCallFunction(s)
	}
}

func (s *StructCallFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitStructCallFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

type SelfFunctionContext struct {
	CallBackContext
}

func NewSelfFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelfFunctionContext {
	var p = new(SelfFunctionContext)

	InitEmptyCallBackContext(&p.CallBackContext)
	p.parser = parser
	p.CopyAll(ctx.(*CallBackContext))

	return p
}

func (s *SelfFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelfFunctionContext) SELF() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserSELF, 0)
}

func (s *SelfFunctionContext) DOT() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserDOT, 0)
}

func (s *SelfFunctionContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *SelfFunctionContext) IS_() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserIS_, 0)
}

func (s *SelfFunctionContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SelfFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterSelfFunction(s)
	}
}

func (s *SelfFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitSelfFunction(s)
	}
}

func (s *SelfFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitSelfFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructAttributeContext struct {
	CallBackContext
}

func NewStructAttributeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructAttributeContext {
	var p = new(StructAttributeContext)

	InitEmptyCallBackContext(&p.CallBackContext)
	p.parser = parser
	p.CopyAll(ctx.(*CallBackContext))

	return p
}

func (s *StructAttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructAttributeContext) AllID_PRIMITIVE() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserID_PRIMITIVE)
}

func (s *StructAttributeContext) ID_PRIMITIVE(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, i)
}

func (s *StructAttributeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserDOT)
}

func (s *StructAttributeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserDOT, i)
}

func (s *StructAttributeContext) IS_() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserIS_, 0)
}

func (s *StructAttributeContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StructAttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterStructAttribute(s)
	}
}

func (s *StructAttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitStructAttribute(s)
	}
}

func (s *StructAttributeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitStructAttribute(s)

	default:
		return t.VisitChildren(s)
	}
}

type RemoveAtVectorContext struct {
	CallBackContext
}

func NewRemoveAtVectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RemoveAtVectorContext {
	var p = new(RemoveAtVectorContext)

	InitEmptyCallBackContext(&p.CallBackContext)
	p.parser = parser
	p.CopyAll(ctx.(*CallBackContext))

	return p
}

func (s *RemoveAtVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RemoveAtVectorContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *RemoveAtVectorContext) DOT() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserDOT, 0)
}

func (s *RemoveAtVectorContext) REMOVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserREMOVE, 0)
}

func (s *RemoveAtVectorContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLPAREN, 0)
}

func (s *RemoveAtVectorContext) AT() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserAT, 0)
}

func (s *RemoveAtVectorContext) COLON() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOLON, 0)
}

func (s *RemoveAtVectorContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RemoveAtVectorContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRPAREN, 0)
}

func (s *RemoveAtVectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterRemoveAtVector(s)
	}
}

func (s *RemoveAtVectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitRemoveAtVector(s)
	}
}

func (s *RemoveAtVectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitRemoveAtVector(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsEmptyVectorContext struct {
	CallBackContext
}

func NewIsEmptyVectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsEmptyVectorContext {
	var p = new(IsEmptyVectorContext)

	InitEmptyCallBackContext(&p.CallBackContext)
	p.parser = parser
	p.CopyAll(ctx.(*CallBackContext))

	return p
}

func (s *IsEmptyVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsEmptyVectorContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *IsEmptyVectorContext) DOT() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserDOT, 0)
}

func (s *IsEmptyVectorContext) ISEMPTY() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserISEMPTY, 0)
}

func (s *IsEmptyVectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterIsEmptyVector(s)
	}
}

func (s *IsEmptyVectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitIsEmptyVector(s)
	}
}

func (s *IsEmptyVectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitIsEmptyVector(s)

	default:
		return t.VisitChildren(s)
	}
}

type CountVectorContext struct {
	CallBackContext
}

func NewCountVectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CountVectorContext {
	var p = new(CountVectorContext)

	InitEmptyCallBackContext(&p.CallBackContext)
	p.parser = parser
	p.CopyAll(ctx.(*CallBackContext))

	return p
}

func (s *CountVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CountVectorContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *CountVectorContext) DOT() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserDOT, 0)
}

func (s *CountVectorContext) COUNT() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOUNT, 0)
}

func (s *CountVectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterCountVector(s)
	}
}

func (s *CountVectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitCountVector(s)
	}
}

func (s *CountVectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitCountVector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) CallBack() (localctx ICallBackContext) {
	localctx = NewCallBackContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SFGrammarParserRULE_callBack)
	var _la int

	var _alt int

	p.SetState(670)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 79, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAppendVectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(608)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(609)
			p.Match(SFGrammarParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(610)
			p.Match(SFGrammarParserAPPEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(611)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(612)
			p.expr(0)
		}
		{
			p.SetState(613)
			p.Match(SFGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewRemoveLastVectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(615)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(616)
			p.Match(SFGrammarParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(617)
			p.Match(SFGrammarParserREMOVELAST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(618)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(619)
			p.Match(SFGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewRemoveAtVectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(620)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(621)
			p.Match(SFGrammarParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(622)
			p.Match(SFGrammarParserREMOVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(623)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(624)
			p.Match(SFGrammarParserAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(625)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(626)
			p.expr(0)
		}
		{
			p.SetState(627)
			p.Match(SFGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewIsEmptyVectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(629)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(630)
			p.Match(SFGrammarParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(631)
			p.Match(SFGrammarParserISEMPTY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewCountVectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(632)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(633)
			p.Match(SFGrammarParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(634)
			p.Match(SFGrammarParserCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewAccessVectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(635)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(636)
			p.Match(SFGrammarParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(637)
			p.expr(0)
		}
		{
			p.SetState(638)
			p.Match(SFGrammarParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewStructCallFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(640)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(643)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SFGrammarParserDOT {
			{
				p.SetState(641)
				p.Match(SFGrammarParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(642)
				p.Match(SFGrammarParserID_PRIMITIVE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(645)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(647)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(649)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserNOT_PARAM || _la == SFGrammarParserID_PRIMITIVE {
			{
				p.SetState(648)
				p.ListFunctionParams()
			}

		}
		{
			p.SetState(651)
			p.Match(SFGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewStructAttributeContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(652)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(655)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(653)
					p.Match(SFGrammarParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(654)
					p.Match(SFGrammarParserID_PRIMITIVE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(657)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 76, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(661)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 77, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(659)
				p.Match(SFGrammarParserIS_)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(660)
				p.expr(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 9:
		localctx = NewSelfFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(663)
			p.Match(SFGrammarParserSELF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(664)
			p.Match(SFGrammarParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(665)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(668)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 78, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(666)
				p.Match(SFGrammarParserIS_)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(667)
				p.expr(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
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

// IEmbbededFuncContext is an interface to support dynamic dispatch.
type IEmbbededFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Intstmt() IIntstmtContext
	Floatstmt() IFloatstmtContext
	Stringstmt() IStringstmtContext

	// IsEmbbededFuncContext differentiates from other interfaces.
	IsEmbbededFuncContext()
}

type EmbbededFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmbbededFuncContext() *EmbbededFuncContext {
	var p = new(EmbbededFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_embbededFunc
	return p
}

func InitEmptyEmbbededFuncContext(p *EmbbededFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_embbededFunc
}

func (*EmbbededFuncContext) IsEmbbededFuncContext() {}

func NewEmbbededFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmbbededFuncContext {
	var p = new(EmbbededFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_embbededFunc

	return p
}

func (s *EmbbededFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *EmbbededFuncContext) Intstmt() IIntstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntstmtContext)
}

func (s *EmbbededFuncContext) Floatstmt() IFloatstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloatstmtContext)
}

func (s *EmbbededFuncContext) Stringstmt() IStringstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringstmtContext)
}

func (s *EmbbededFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmbbededFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmbbededFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterEmbbededFunc(s)
	}
}

func (s *EmbbededFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitEmbbededFunc(s)
	}
}

func (s *EmbbededFuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitEmbbededFunc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) EmbbededFunc() (localctx IEmbbededFuncContext) {
	localctx = NewEmbbededFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SFGrammarParserRULE_embbededFunc)
	p.SetState(675)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SFGrammarParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(672)
			p.Intstmt()
		}

	case SFGrammarParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(673)
			p.Floatstmt()
		}

	case SFGrammarParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(674)
			p.Stringstmt()
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

// IPrintstmtContext is an interface to support dynamic dispatch.
type IPrintstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	ExprList() IExprListContext
	RPAREN() antlr.TerminalNode

	// IsPrintstmtContext differentiates from other interfaces.
	IsPrintstmtContext()
}

type PrintstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintstmtContext() *PrintstmtContext {
	var p = new(PrintstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_printstmt
	return p
}

func InitEmptyPrintstmtContext(p *PrintstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_printstmt
}

func (*PrintstmtContext) IsPrintstmtContext() {}

func NewPrintstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintstmtContext {
	var p = new(PrintstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_printstmt

	return p
}

func (s *PrintstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintstmtContext) PRINT() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserPRINT, 0)
}

func (s *PrintstmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLPAREN, 0)
}

func (s *PrintstmtContext) ExprList() IExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *PrintstmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRPAREN, 0)
}

func (s *PrintstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterPrintstmt(s)
	}
}

func (s *PrintstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitPrintstmt(s)
	}
}

func (s *PrintstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitPrintstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) Printstmt() (localctx IPrintstmtContext) {
	localctx = NewPrintstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SFGrammarParserRULE_printstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(677)
		p.Match(SFGrammarParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(678)
		p.Match(SFGrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(679)
		p.ExprList()
	}
	{
		p.SetState(680)
		p.Match(SFGrammarParserRPAREN)
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

// IExprListContext is an interface to support dynamic dispatch.
type IExprListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsExprListContext differentiates from other interfaces.
	IsExprListContext()
}

type ExprListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprListContext() *ExprListContext {
	var p = new(ExprListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_exprList
	return p
}

func InitEmptyExprListContext(p *ExprListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_exprList
}

func (*ExprListContext) IsExprListContext() {}

func NewExprListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprListContext {
	var p = new(ExprListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_exprList

	return p
}

func (s *ExprListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprListContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprListContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ExprListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserCOMMA)
}

func (s *ExprListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOMMA, i)
}

func (s *ExprListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterExprList(s)
	}
}

func (s *ExprListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitExprList(s)
	}
}

func (s *ExprListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitExprList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) ExprList() (localctx IExprListContext) {
	localctx = NewExprListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SFGrammarParserRULE_exprList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(682)
		p.expr(0)
	}
	p.SetState(687)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SFGrammarParserCOMMA {
		{
			p.SetState(683)
			p.Match(SFGrammarParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(684)
			p.expr(0)
		}

		p.SetState(689)
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

// IIntstmtContext is an interface to support dynamic dispatch.
type IIntstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode

	// IsIntstmtContext differentiates from other interfaces.
	IsIntstmtContext()
}

type IntstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntstmtContext() *IntstmtContext {
	var p = new(IntstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_intstmt
	return p
}

func InitEmptyIntstmtContext(p *IntstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_intstmt
}

func (*IntstmtContext) IsIntstmtContext() {}

func NewIntstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntstmtContext {
	var p = new(IntstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_intstmt

	return p
}

func (s *IntstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IntstmtContext) INT() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserINT, 0)
}

func (s *IntstmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLPAREN, 0)
}

func (s *IntstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IntstmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRPAREN, 0)
}

func (s *IntstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterIntstmt(s)
	}
}

func (s *IntstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitIntstmt(s)
	}
}

func (s *IntstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitIntstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) Intstmt() (localctx IIntstmtContext) {
	localctx = NewIntstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SFGrammarParserRULE_intstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(690)
		p.Match(SFGrammarParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(691)
		p.Match(SFGrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(692)
		p.expr(0)
	}
	{
		p.SetState(693)
		p.Match(SFGrammarParserRPAREN)
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

// IFloatstmtContext is an interface to support dynamic dispatch.
type IFloatstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FLOAT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode

	// IsFloatstmtContext differentiates from other interfaces.
	IsFloatstmtContext()
}

type FloatstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatstmtContext() *FloatstmtContext {
	var p = new(FloatstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_floatstmt
	return p
}

func InitEmptyFloatstmtContext(p *FloatstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_floatstmt
}

func (*FloatstmtContext) IsFloatstmtContext() {}

func NewFloatstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatstmtContext {
	var p = new(FloatstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_floatstmt

	return p
}

func (s *FloatstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatstmtContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserFLOAT, 0)
}

func (s *FloatstmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLPAREN, 0)
}

func (s *FloatstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FloatstmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRPAREN, 0)
}

func (s *FloatstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterFloatstmt(s)
	}
}

func (s *FloatstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitFloatstmt(s)
	}
}

func (s *FloatstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitFloatstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) Floatstmt() (localctx IFloatstmtContext) {
	localctx = NewFloatstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SFGrammarParserRULE_floatstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(695)
		p.Match(SFGrammarParserFLOAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(696)
		p.Match(SFGrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(697)
		p.expr(0)
	}
	{
		p.SetState(698)
		p.Match(SFGrammarParserRPAREN)
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

// IStringstmtContext is an interface to support dynamic dispatch.
type IStringstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode

	// IsStringstmtContext differentiates from other interfaces.
	IsStringstmtContext()
}

type StringstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringstmtContext() *StringstmtContext {
	var p = new(StringstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_stringstmt
	return p
}

func InitEmptyStringstmtContext(p *StringstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_stringstmt
}

func (*StringstmtContext) IsStringstmtContext() {}

func NewStringstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringstmtContext {
	var p = new(StringstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_stringstmt

	return p
}

func (s *StringstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StringstmtContext) STRING() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserSTRING, 0)
}

func (s *StringstmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLPAREN, 0)
}

func (s *StringstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StringstmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRPAREN, 0)
}

func (s *StringstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterStringstmt(s)
	}
}

func (s *StringstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitStringstmt(s)
	}
}

func (s *StringstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitStringstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) Stringstmt() (localctx IStringstmtContext) {
	localctx = NewStringstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SFGrammarParserRULE_stringstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(700)
		p.Match(SFGrammarParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(701)
		p.Match(SFGrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(702)
		p.expr(0)
	}
	{
		p.SetState(703)
		p.Match(SFGrammarParserRPAREN)
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringExprContext struct {
	ExprContext
}

func NewStringExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringExprContext {
	var p = new(StringExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StringExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringExprContext) STRING_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserSTRING_PRIMITIVE, 0)
}

func (s *StringExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterStringExpr(s)
	}
}

func (s *StringExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitStringExpr(s)
	}
}

func (s *StringExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitStringExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type EmbeddedFunctionExprContext struct {
	ExprContext
}

func NewEmbeddedFunctionExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmbeddedFunctionExprContext {
	var p = new(EmbeddedFunctionExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EmbeddedFunctionExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmbeddedFunctionExprContext) EmbbededFunc() IEmbbededFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmbbededFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmbbededFuncContext)
}

func (s *EmbeddedFunctionExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterEmbeddedFunctionExpr(s)
	}
}

func (s *EmbeddedFunctionExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitEmbeddedFunctionExpr(s)
	}
}

func (s *EmbeddedFunctionExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitEmbeddedFunctionExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NilExprContext struct {
	ExprContext
}

func NewNilExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilExprContext {
	var p = new(NilExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NilExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilExprContext) NIL() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserNIL, 0)
}

func (s *NilExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterNilExpr(s)
	}
}

func (s *NilExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitNilExpr(s)
	}
}

func (s *NilExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitNilExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdExprContext struct {
	ExprContext
}

func NewIdExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdExprContext {
	var p = new(IdExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdExprContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *IdExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterIdExpr(s)
	}
}

func (s *IdExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitIdExpr(s)
	}
}

func (s *IdExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitIdExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallBackExprContext struct {
	ExprContext
}

func NewCallBackExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallBackExprContext {
	var p = new(CallBackExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CallBackExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallBackExprContext) CallBack() ICallBackContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallBackContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallBackContext)
}

func (s *CallBackExprContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserSEMICOLON, 0)
}

func (s *CallBackExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterCallBackExpr(s)
	}
}

func (s *CallBackExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitCallBackExpr(s)
	}
}

func (s *CallBackExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitCallBackExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogicalOperationExprContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewLogicalOperationExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalOperationExprContext {
	var p = new(LogicalOperationExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LogicalOperationExprContext) GetOp() antlr.Token { return s.op }

func (s *LogicalOperationExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *LogicalOperationExprContext) GetLeft() IExprContext { return s.left }

func (s *LogicalOperationExprContext) GetRight() IExprContext { return s.right }

func (s *LogicalOperationExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *LogicalOperationExprContext) SetRight(v IExprContext) { s.right = v }

func (s *LogicalOperationExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOperationExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *LogicalOperationExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *LogicalOperationExprContext) AND() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserAND, 0)
}

func (s *LogicalOperationExprContext) OR() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserOR, 0)
}

func (s *LogicalOperationExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterLogicalOperationExpr(s)
	}
}

func (s *LogicalOperationExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitLogicalOperationExpr(s)
	}
}

func (s *LogicalOperationExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitLogicalOperationExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NegExprContext struct {
	ExprContext
	right IExprContext
}

func NewNegExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegExprContext {
	var p = new(NegExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NegExprContext) GetRight() IExprContext { return s.right }

func (s *NegExprContext) SetRight(v IExprContext) { s.right = v }

func (s *NegExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserMINUS, 0)
}

func (s *NegExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NegExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterNegExpr(s)
	}
}

func (s *NegExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitNegExpr(s)
	}
}

func (s *NegExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitNegExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ComparationOperationExprContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewComparationOperationExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparationOperationExprContext {
	var p = new(ComparationOperationExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ComparationOperationExprContext) GetOp() antlr.Token { return s.op }

func (s *ComparationOperationExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ComparationOperationExprContext) GetLeft() IExprContext { return s.left }

func (s *ComparationOperationExprContext) GetRight() IExprContext { return s.right }

func (s *ComparationOperationExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *ComparationOperationExprContext) SetRight(v IExprContext) { s.right = v }

func (s *ComparationOperationExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparationOperationExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ComparationOperationExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ComparationOperationExprContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserEQUALS, 0)
}

func (s *ComparationOperationExprContext) NOT_EQUALS() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserNOT_EQUALS, 0)
}

func (s *ComparationOperationExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterComparationOperationExpr(s)
	}
}

func (s *ComparationOperationExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitComparationOperationExpr(s)
	}
}

func (s *ComparationOperationExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitComparationOperationExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructAsArgumentContext struct {
	ExprContext
}

func NewStructAsArgumentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructAsArgumentContext {
	var p = new(StructAsArgumentContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StructAsArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructAsArgumentContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *StructAsArgumentContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLPAREN, 0)
}

func (s *StructAsArgumentContext) StructCallList() IStructCallListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructCallListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructCallListContext)
}

func (s *StructAsArgumentContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRPAREN, 0)
}

func (s *StructAsArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterStructAsArgument(s)
	}
}

func (s *StructAsArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitStructAsArgument(s)
	}
}

func (s *StructAsArgumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitStructAsArgument(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArithmeticOperationExprContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewArithmeticOperationExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArithmeticOperationExprContext {
	var p = new(ArithmeticOperationExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ArithmeticOperationExprContext) GetOp() antlr.Token { return s.op }

func (s *ArithmeticOperationExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ArithmeticOperationExprContext) GetLeft() IExprContext { return s.left }

func (s *ArithmeticOperationExprContext) GetRight() IExprContext { return s.right }

func (s *ArithmeticOperationExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *ArithmeticOperationExprContext) SetRight(v IExprContext) { s.right = v }

func (s *ArithmeticOperationExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticOperationExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArithmeticOperationExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ArithmeticOperationExprContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserDIVIDE, 0)
}

func (s *ArithmeticOperationExprContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserMULTIPLY, 0)
}

func (s *ArithmeticOperationExprContext) MODULO() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserMODULO, 0)
}

func (s *ArithmeticOperationExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserPLUS, 0)
}

func (s *ArithmeticOperationExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserMINUS, 0)
}

func (s *ArithmeticOperationExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterArithmeticOperationExpr(s)
	}
}

func (s *ArithmeticOperationExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitArithmeticOperationExpr(s)
	}
}

func (s *ArithmeticOperationExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitArithmeticOperationExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type RelationalOperationExprContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewRelationalOperationExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelationalOperationExprContext {
	var p = new(RelationalOperationExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *RelationalOperationExprContext) GetOp() antlr.Token { return s.op }

func (s *RelationalOperationExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *RelationalOperationExprContext) GetLeft() IExprContext { return s.left }

func (s *RelationalOperationExprContext) GetRight() IExprContext { return s.right }

func (s *RelationalOperationExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *RelationalOperationExprContext) SetRight(v IExprContext) { s.right = v }

func (s *RelationalOperationExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalOperationExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *RelationalOperationExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *RelationalOperationExprContext) GREATER() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserGREATER, 0)
}

func (s *RelationalOperationExprContext) GREATER_EQUALS() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserGREATER_EQUALS, 0)
}

func (s *RelationalOperationExprContext) LESS() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLESS, 0)
}

func (s *RelationalOperationExprContext) LESS_EQUALS() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLESS_EQUALS, 0)
}

func (s *RelationalOperationExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterRelationalOperationExpr(s)
	}
}

func (s *RelationalOperationExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitRelationalOperationExpr(s)
	}
}

func (s *RelationalOperationExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitRelationalOperationExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type DigitExprContext struct {
	ExprContext
}

func NewDigitExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DigitExprContext {
	var p = new(DigitExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *DigitExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DigitExprContext) DIGIT_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserDIGIT_PRIMITIVE, 0)
}

func (s *DigitExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterDigitExpr(s)
	}
}

func (s *DigitExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitDigitExpr(s)
	}
}

func (s *DigitExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitDigitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExprContext struct {
	ExprContext
	right IExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NotExprContext) GetRight() IExprContext { return s.right }

func (s *NotExprContext) SetRight(v IExprContext) { s.right = v }

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) NEGATION_OPERATOR() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserNEGATION_OPERATOR, 0)
}

func (s *NotExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExprContext struct {
	ExprContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLPAREN, 0)
}

func (s *ParenExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRPAREN, 0)
}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitParenExpr(s)
	}
}

func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitParenExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallFunctionExprContext struct {
	ExprContext
}

func NewCallFunctionExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallFunctionExprContext {
	var p = new(CallFunctionExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CallFunctionExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallFunctionExprContext) CallFunctionStmt() ICallFunctionStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallFunctionStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallFunctionStmtContext)
}

func (s *CallFunctionExprContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserSEMICOLON, 0)
}

func (s *CallFunctionExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterCallFunctionExpr(s)
	}
}

func (s *CallFunctionExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitCallFunctionExpr(s)
	}
}

func (s *CallFunctionExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitCallFunctionExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanExprContext struct {
	ExprContext
}

func NewBooleanExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanExprContext {
	var p = new(BooleanExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BooleanExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanExprContext) TRU() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserTRU, 0)
}

func (s *BooleanExprContext) FAL() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserFAL, 0)
}

func (s *BooleanExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterBooleanExpr(s)
	}
}

func (s *BooleanExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitBooleanExpr(s)
	}
}

func (s *BooleanExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitBooleanExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *SFGrammarParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 64
	p.EnterRecursionRule(localctx, 64, SFGrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(733)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 84, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(706)
			p.Match(SFGrammarParserNEGATION_OPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(707)

			var _x = p.expr(17)

			localctx.(*NotExprContext).right = _x
		}

	case 2:
		localctx = NewNegExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(708)
			p.Match(SFGrammarParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(709)

			var _x = p.expr(16)

			localctx.(*NegExprContext).right = _x
		}

	case 3:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(710)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(711)
			p.expr(0)
		}
		{
			p.SetState(712)
			p.Match(SFGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewStructAsArgumentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(714)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(715)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(716)
			p.StructCallList()
		}
		{
			p.SetState(717)
			p.Match(SFGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewCallFunctionExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(719)
			p.CallFunctionStmt()
		}
		p.SetState(721)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 82, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(720)
				p.Match(SFGrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 6:
		localctx = NewCallBackExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(723)
			p.CallBack()
		}
		p.SetState(725)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 83, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(724)
				p.Match(SFGrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 7:
		localctx = NewEmbeddedFunctionExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(727)
			p.EmbbededFunc()
		}

	case 8:
		localctx = NewDigitExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(728)
			p.Match(SFGrammarParserDIGIT_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewStringExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(729)
			p.Match(SFGrammarParserSTRING_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewIdExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(730)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewNilExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(731)
			p.Match(SFGrammarParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		localctx = NewBooleanExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(732)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SFGrammarParserTRU || _la == SFGrammarParserFAL) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(752)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 86, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(750)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 85, p.GetParserRuleContext()) {
			case 1:
				localctx = NewArithmeticOperationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticOperationExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SFGrammarParserRULE_expr)
				p.SetState(735)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(736)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ArithmeticOperationExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2017612633061982208) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ArithmeticOperationExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(737)

					var _x = p.expr(16)

					localctx.(*ArithmeticOperationExprContext).right = _x
				}

			case 2:
				localctx = NewArithmeticOperationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticOperationExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SFGrammarParserRULE_expr)
				p.SetState(738)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(739)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ArithmeticOperationExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SFGrammarParserPLUS || _la == SFGrammarParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ArithmeticOperationExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(740)

					var _x = p.expr(15)

					localctx.(*ArithmeticOperationExprContext).right = _x
				}

			case 3:
				localctx = NewComparationOperationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ComparationOperationExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SFGrammarParserRULE_expr)
				p.SetState(741)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(742)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ComparationOperationExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SFGrammarParserEQUALS || _la == SFGrammarParserNOT_EQUALS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ComparationOperationExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(743)

					var _x = p.expr(14)

					localctx.(*ComparationOperationExprContext).right = _x
				}

			case 4:
				localctx = NewRelationalOperationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*RelationalOperationExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SFGrammarParserRULE_expr)
				p.SetState(744)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(745)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelationalOperationExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-63)) & ^0x3f) == 0 && ((int64(1)<<(_la-63))&15) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelationalOperationExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(746)

					var _x = p.expr(13)

					localctx.(*RelationalOperationExprContext).right = _x
				}

			case 5:
				localctx = NewLogicalOperationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LogicalOperationExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SFGrammarParserRULE_expr)
				p.SetState(747)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(748)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*LogicalOperationExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SFGrammarParserAND || _la == SFGrammarParserOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*LogicalOperationExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(749)

					var _x = p.expr(12)

					localctx.(*LogicalOperationExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(754)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 86, p.GetParserRuleContext())
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	STRING() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	CHAR() antlr.TerminalNode
	ID_PRIMITIVE() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SFGrammarParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFGrammarParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) INT() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserINT, 0)
}

func (s *TypeContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserFLOAT, 0)
}

func (s *TypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserSTRING, 0)
}

func (s *TypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserBOOL, 0)
}

func (s *TypeContext) CHAR() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCHAR, 0)
}

func (s *TypeContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFGrammarListener); ok {
		listenerT.ExitType(s)
	}
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFGrammarParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SFGrammarParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(755)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2199023255614) != 0) {
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

func (p *SFGrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 32:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SFGrammarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 11)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
