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
		"structStmts", "declarationStructs", "structCallList", "declaration",
		"type_declaration", "assignment", "ifstmt", "switchStmt", "caseBlock",
		"defaultBlock", "whileStmt", "forStmt", "forRange", "guardStmt", "functionStmt",
		"listFunctionParams", "callFunctionStmt", "listCallFunctionStmt", "callBack",
		"embbededFunc", "printstmt", "exprList", "intstmt", "floatstmt", "stringstmt",
		"expr", "type",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 71, 681, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 71, 8, 1, 10, 1, 12, 1,
		74, 9, 1, 1, 2, 1, 2, 3, 2, 78, 8, 2, 1, 2, 1, 2, 3, 2, 82, 8, 2, 1, 2,
		1, 2, 3, 2, 86, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 3, 2, 98, 8, 2, 1, 2, 1, 2, 3, 2, 102, 8, 2, 1, 2, 3, 2, 105,
		8, 2, 1, 3, 1, 3, 3, 3, 109, 8, 3, 1, 3, 1, 3, 3, 3, 113, 8, 3, 1, 3, 1,
		3, 3, 3, 117, 8, 3, 1, 3, 3, 3, 120, 8, 3, 3, 3, 122, 8, 3, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 5, 5, 131, 8, 5, 10, 5, 12, 5, 134, 9, 5,
		1, 6, 1, 6, 3, 6, 138, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 169,
		8, 7, 3, 7, 171, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8,
		180, 8, 8, 10, 8, 12, 8, 183, 9, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		3, 9, 215, 8, 9, 3, 9, 217, 8, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 251, 8, 11, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 263, 8,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 273,
		8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 279, 8, 13, 10, 13, 12, 13, 282,
		9, 13, 1, 13, 3, 13, 285, 8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 320, 8, 17, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 20, 3, 20, 336, 8, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 3, 20, 344, 8, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 351,
		8, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 360, 8,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 366, 8, 20, 1, 21, 1, 21, 1, 21,
		1, 21, 3, 21, 372, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5,
		21, 380, 8, 21, 10, 21, 12, 21, 383, 9, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		3, 21, 389, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 397,
		8, 21, 10, 21, 12, 21, 400, 9, 21, 1, 21, 1, 21, 1, 21, 3, 21, 405, 8,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 412, 8, 21, 1, 21, 5, 21,
		415, 8, 21, 10, 21, 12, 21, 418, 9, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3,
		21, 424, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		3, 21, 434, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 440, 8, 21, 10, 21,
		12, 21, 443, 9, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 449, 8, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 459, 8, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 5, 21, 465, 8, 21, 10, 21, 12, 21, 468, 9, 21,
		1, 21, 1, 21, 1, 21, 3, 21, 473, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 3, 21, 482, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21,
		488, 8, 21, 10, 21, 12, 21, 491, 9, 21, 3, 21, 493, 8, 21, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 503, 8, 22, 1, 23, 3,
		23, 506, 8, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 513, 8, 23, 1,
		23, 1, 23, 1, 23, 5, 23, 518, 8, 23, 10, 23, 12, 23, 521, 9, 23, 1, 23,
		3, 23, 524, 8, 23, 1, 23, 1, 23, 1, 23, 3, 23, 529, 8, 23, 1, 23, 5, 23,
		532, 8, 23, 10, 23, 12, 23, 535, 9, 23, 3, 23, 537, 8, 23, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 578, 8, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 3, 24, 585, 8, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 3, 24, 594, 8, 24, 1, 25, 1, 25, 1, 25, 3, 25, 599, 8,
		25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 5, 27, 609,
		8, 27, 10, 27, 12, 27, 612, 9, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3,
		31, 640, 8, 31, 1, 31, 1, 31, 3, 31, 644, 8, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 657, 8, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 5, 31, 674, 8, 31, 10, 31, 12, 31, 677,
		9, 31, 1, 32, 1, 32, 1, 32, 0, 1, 62, 33, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
		54, 56, 58, 60, 62, 64, 0, 8, 1, 0, 34, 35, 1, 0, 31, 32, 1, 0, 58, 60,
		1, 0, 56, 57, 1, 0, 61, 62, 1, 0, 63, 66, 1, 0, 67, 68, 2, 0, 1, 5, 41,
		41, 758, 0, 66, 1, 0, 0, 0, 2, 72, 1, 0, 0, 0, 4, 104, 1, 0, 0, 0, 6, 121,
		1, 0, 0, 0, 8, 123, 1, 0, 0, 0, 10, 132, 1, 0, 0, 0, 12, 135, 1, 0, 0,
		0, 14, 170, 1, 0, 0, 0, 16, 172, 1, 0, 0, 0, 18, 216, 1, 0, 0, 0, 20, 218,
		1, 0, 0, 0, 22, 250, 1, 0, 0, 0, 24, 272, 1, 0, 0, 0, 26, 274, 1, 0, 0,
		0, 28, 288, 1, 0, 0, 0, 30, 293, 1, 0, 0, 0, 32, 297, 1, 0, 0, 0, 34, 319,
		1, 0, 0, 0, 36, 321, 1, 0, 0, 0, 38, 327, 1, 0, 0, 0, 40, 365, 1, 0, 0,
		0, 42, 492, 1, 0, 0, 0, 44, 502, 1, 0, 0, 0, 46, 536, 1, 0, 0, 0, 48, 593,
		1, 0, 0, 0, 50, 598, 1, 0, 0, 0, 52, 600, 1, 0, 0, 0, 54, 605, 1, 0, 0,
		0, 56, 613, 1, 0, 0, 0, 58, 618, 1, 0, 0, 0, 60, 623, 1, 0, 0, 0, 62, 656,
		1, 0, 0, 0, 64, 678, 1, 0, 0, 0, 66, 67, 3, 2, 1, 0, 67, 68, 5, 0, 0, 1,
		68, 1, 1, 0, 0, 0, 69, 71, 3, 4, 2, 0, 70, 69, 1, 0, 0, 0, 71, 74, 1, 0,
		0, 0, 72, 70, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 3, 1, 0, 0, 0, 74, 72,
		1, 0, 0, 0, 75, 77, 3, 18, 9, 0, 76, 78, 5, 51, 0, 0, 77, 76, 1, 0, 0,
		0, 77, 78, 1, 0, 0, 0, 78, 105, 1, 0, 0, 0, 79, 81, 3, 22, 11, 0, 80, 82,
		5, 51, 0, 0, 81, 80, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 105, 1, 0, 0,
		0, 83, 85, 3, 50, 25, 0, 84, 86, 5, 51, 0, 0, 85, 84, 1, 0, 0, 0, 85, 86,
		1, 0, 0, 0, 86, 105, 1, 0, 0, 0, 87, 105, 3, 24, 12, 0, 88, 105, 3, 26,
		13, 0, 89, 105, 3, 32, 16, 0, 90, 105, 3, 34, 17, 0, 91, 105, 3, 38, 19,
		0, 92, 105, 3, 6, 3, 0, 93, 105, 3, 40, 20, 0, 94, 105, 3, 52, 26, 0, 95,
		97, 3, 44, 22, 0, 96, 98, 5, 51, 0, 0, 97, 96, 1, 0, 0, 0, 97, 98, 1, 0,
		0, 0, 98, 105, 1, 0, 0, 0, 99, 101, 3, 48, 24, 0, 100, 102, 5, 51, 0, 0,
		101, 100, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 105, 1, 0, 0, 0, 103,
		105, 3, 8, 4, 0, 104, 75, 1, 0, 0, 0, 104, 79, 1, 0, 0, 0, 104, 83, 1,
		0, 0, 0, 104, 87, 1, 0, 0, 0, 104, 88, 1, 0, 0, 0, 104, 89, 1, 0, 0, 0,
		104, 90, 1, 0, 0, 0, 104, 91, 1, 0, 0, 0, 104, 92, 1, 0, 0, 0, 104, 93,
		1, 0, 0, 0, 104, 94, 1, 0, 0, 0, 104, 95, 1, 0, 0, 0, 104, 99, 1, 0, 0,
		0, 104, 103, 1, 0, 0, 0, 105, 5, 1, 0, 0, 0, 106, 108, 5, 20, 0, 0, 107,
		109, 5, 51, 0, 0, 108, 107, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 122,
		1, 0, 0, 0, 110, 112, 5, 21, 0, 0, 111, 113, 5, 51, 0, 0, 112, 111, 1,
		0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 122, 1, 0, 0, 0, 114, 116, 5, 22, 0,
		0, 115, 117, 3, 62, 31, 0, 116, 115, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0,
		117, 119, 1, 0, 0, 0, 118, 120, 5, 51, 0, 0, 119, 118, 1, 0, 0, 0, 119,
		120, 1, 0, 0, 0, 120, 122, 1, 0, 0, 0, 121, 106, 1, 0, 0, 0, 121, 110,
		1, 0, 0, 0, 121, 114, 1, 0, 0, 0, 122, 7, 1, 0, 0, 0, 123, 124, 5, 7, 0,
		0, 124, 125, 5, 41, 0, 0, 125, 126, 5, 45, 0, 0, 126, 127, 3, 10, 5, 0,
		127, 128, 5, 46, 0, 0, 128, 9, 1, 0, 0, 0, 129, 131, 3, 12, 6, 0, 130,
		129, 1, 0, 0, 0, 131, 134, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 132, 133,
		1, 0, 0, 0, 133, 11, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 135, 137, 3, 14,
		7, 0, 136, 138, 5, 51, 0, 0, 137, 136, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0,
		138, 13, 1, 0, 0, 0, 139, 140, 3, 20, 10, 0, 140, 141, 5, 41, 0, 0, 141,
		142, 5, 49, 0, 0, 142, 143, 3, 64, 32, 0, 143, 144, 5, 52, 0, 0, 144, 145,
		3, 62, 31, 0, 145, 171, 1, 0, 0, 0, 146, 147, 3, 20, 10, 0, 147, 148, 5,
		41, 0, 0, 148, 149, 5, 49, 0, 0, 149, 150, 3, 64, 32, 0, 150, 171, 1, 0,
		0, 0, 151, 152, 3, 20, 10, 0, 152, 153, 5, 41, 0, 0, 153, 154, 5, 52, 0,
		0, 154, 155, 3, 62, 31, 0, 155, 171, 1, 0, 0, 0, 156, 157, 3, 20, 10, 0,
		157, 158, 5, 41, 0, 0, 158, 159, 5, 49, 0, 0, 159, 160, 5, 47, 0, 0, 160,
		161, 3, 64, 32, 0, 161, 162, 5, 48, 0, 0, 162, 168, 5, 52, 0, 0, 163, 164,
		5, 47, 0, 0, 164, 165, 3, 54, 27, 0, 165, 166, 5, 48, 0, 0, 166, 169, 1,
		0, 0, 0, 167, 169, 5, 41, 0, 0, 168, 163, 1, 0, 0, 0, 168, 167, 1, 0, 0,
		0, 169, 171, 1, 0, 0, 0, 170, 139, 1, 0, 0, 0, 170, 146, 1, 0, 0, 0, 170,
		151, 1, 0, 0, 0, 170, 156, 1, 0, 0, 0, 171, 15, 1, 0, 0, 0, 172, 173, 5,
		41, 0, 0, 173, 174, 5, 49, 0, 0, 174, 181, 3, 62, 31, 0, 175, 176, 5, 50,
		0, 0, 176, 177, 5, 41, 0, 0, 177, 178, 5, 49, 0, 0, 178, 180, 3, 62, 31,
		0, 179, 175, 1, 0, 0, 0, 180, 183, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 181,
		182, 1, 0, 0, 0, 182, 17, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 184, 185, 3,
		20, 10, 0, 185, 186, 5, 41, 0, 0, 186, 187, 5, 49, 0, 0, 187, 188, 3, 64,
		32, 0, 188, 189, 5, 52, 0, 0, 189, 190, 3, 62, 31, 0, 190, 217, 1, 0, 0,
		0, 191, 192, 3, 20, 10, 0, 192, 193, 5, 41, 0, 0, 193, 194, 5, 49, 0, 0,
		194, 195, 3, 64, 32, 0, 195, 196, 5, 55, 0, 0, 196, 217, 1, 0, 0, 0, 197,
		198, 3, 20, 10, 0, 198, 199, 5, 41, 0, 0, 199, 200, 5, 52, 0, 0, 200, 201,
		3, 62, 31, 0, 201, 217, 1, 0, 0, 0, 202, 203, 3, 20, 10, 0, 203, 204, 5,
		41, 0, 0, 204, 205, 5, 49, 0, 0, 205, 206, 5, 47, 0, 0, 206, 207, 3, 64,
		32, 0, 207, 208, 5, 48, 0, 0, 208, 214, 5, 52, 0, 0, 209, 210, 5, 47, 0,
		0, 210, 211, 3, 54, 27, 0, 211, 212, 5, 48, 0, 0, 212, 215, 1, 0, 0, 0,
		213, 215, 5, 41, 0, 0, 214, 209, 1, 0, 0, 0, 214, 213, 1, 0, 0, 0, 215,
		217, 1, 0, 0, 0, 216, 184, 1, 0, 0, 0, 216, 191, 1, 0, 0, 0, 216, 197,
		1, 0, 0, 0, 216, 202, 1, 0, 0, 0, 217, 19, 1, 0, 0, 0, 218, 219, 7, 0,
		0, 0, 219, 21, 1, 0, 0, 0, 220, 221, 5, 41, 0, 0, 221, 222, 5, 52, 0, 0,
		222, 251, 3, 62, 31, 0, 223, 224, 5, 41, 0, 0, 224, 225, 5, 53, 0, 0, 225,
		251, 3, 62, 31, 0, 226, 227, 5, 41, 0, 0, 227, 228, 5, 54, 0, 0, 228, 251,
		3, 62, 31, 0, 229, 230, 5, 41, 0, 0, 230, 231, 5, 47, 0, 0, 231, 232, 3,
		62, 31, 0, 232, 233, 5, 48, 0, 0, 233, 234, 5, 52, 0, 0, 234, 235, 3, 62,
		31, 0, 235, 251, 1, 0, 0, 0, 236, 237, 5, 41, 0, 0, 237, 238, 5, 47, 0,
		0, 238, 239, 3, 62, 31, 0, 239, 240, 5, 48, 0, 0, 240, 241, 5, 54, 0, 0,
		241, 242, 3, 62, 31, 0, 242, 251, 1, 0, 0, 0, 243, 244, 5, 41, 0, 0, 244,
		245, 5, 47, 0, 0, 245, 246, 3, 62, 31, 0, 246, 247, 5, 48, 0, 0, 247, 248,
		5, 53, 0, 0, 248, 249, 3, 62, 31, 0, 249, 251, 1, 0, 0, 0, 250, 220, 1,
		0, 0, 0, 250, 223, 1, 0, 0, 0, 250, 226, 1, 0, 0, 0, 250, 229, 1, 0, 0,
		0, 250, 236, 1, 0, 0, 0, 250, 243, 1, 0, 0, 0, 251, 23, 1, 0, 0, 0, 252,
		253, 5, 6, 0, 0, 253, 254, 3, 62, 31, 0, 254, 255, 5, 45, 0, 0, 255, 256,
		3, 2, 1, 0, 256, 262, 5, 46, 0, 0, 257, 258, 5, 16, 0, 0, 258, 259, 5,
		45, 0, 0, 259, 260, 3, 2, 1, 0, 260, 261, 5, 46, 0, 0, 261, 263, 1, 0,
		0, 0, 262, 257, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 273, 1, 0, 0, 0,
		264, 265, 5, 6, 0, 0, 265, 266, 3, 62, 31, 0, 266, 267, 5, 45, 0, 0, 267,
		268, 3, 2, 1, 0, 268, 269, 5, 46, 0, 0, 269, 270, 5, 16, 0, 0, 270, 271,
		3, 24, 12, 0, 271, 273, 1, 0, 0, 0, 272, 252, 1, 0, 0, 0, 272, 264, 1,
		0, 0, 0, 273, 25, 1, 0, 0, 0, 274, 275, 5, 17, 0, 0, 275, 276, 3, 62, 31,
		0, 276, 280, 5, 45, 0, 0, 277, 279, 3, 28, 14, 0, 278, 277, 1, 0, 0, 0,
		279, 282, 1, 0, 0, 0, 280, 278, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281,
		284, 1, 0, 0, 0, 282, 280, 1, 0, 0, 0, 283, 285, 3, 30, 15, 0, 284, 283,
		1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 287, 5, 46,
		0, 0, 287, 27, 1, 0, 0, 0, 288, 289, 5, 18, 0, 0, 289, 290, 3, 62, 31,
		0, 290, 291, 5, 49, 0, 0, 291, 292, 3, 2, 1, 0, 292, 29, 1, 0, 0, 0, 293,
		294, 5, 19, 0, 0, 294, 295, 5, 49, 0, 0, 295, 296, 3, 2, 1, 0, 296, 31,
		1, 0, 0, 0, 297, 298, 5, 23, 0, 0, 298, 299, 3, 62, 31, 0, 299, 300, 5,
		45, 0, 0, 300, 301, 3, 2, 1, 0, 301, 302, 5, 46, 0, 0, 302, 33, 1, 0, 0,
		0, 303, 304, 5, 24, 0, 0, 304, 305, 5, 41, 0, 0, 305, 306, 5, 27, 0, 0,
		306, 307, 3, 36, 18, 0, 307, 308, 5, 45, 0, 0, 308, 309, 3, 2, 1, 0, 309,
		310, 5, 46, 0, 0, 310, 320, 1, 0, 0, 0, 311, 312, 5, 24, 0, 0, 312, 313,
		5, 41, 0, 0, 313, 314, 5, 27, 0, 0, 314, 315, 3, 62, 31, 0, 315, 316, 5,
		45, 0, 0, 316, 317, 3, 2, 1, 0, 317, 318, 5, 46, 0, 0, 318, 320, 1, 0,
		0, 0, 319, 303, 1, 0, 0, 0, 319, 311, 1, 0, 0, 0, 320, 35, 1, 0, 0, 0,
		321, 322, 3, 62, 31, 0, 322, 323, 5, 28, 0, 0, 323, 324, 5, 28, 0, 0, 324,
		325, 5, 28, 0, 0, 325, 326, 3, 62, 31, 0, 326, 37, 1, 0, 0, 0, 327, 328,
		5, 29, 0, 0, 328, 329, 3, 62, 31, 0, 329, 330, 5, 16, 0, 0, 330, 331, 5,
		45, 0, 0, 331, 332, 3, 2, 1, 0, 332, 333, 5, 46, 0, 0, 333, 39, 1, 0, 0,
		0, 334, 336, 5, 8, 0, 0, 335, 334, 1, 0, 0, 0, 335, 336, 1, 0, 0, 0, 336,
		337, 1, 0, 0, 0, 337, 338, 5, 25, 0, 0, 338, 339, 5, 41, 0, 0, 339, 340,
		5, 43, 0, 0, 340, 343, 5, 44, 0, 0, 341, 342, 5, 26, 0, 0, 342, 344, 3,
		64, 32, 0, 343, 341, 1, 0, 0, 0, 343, 344, 1, 0, 0, 0, 344, 345, 1, 0,
		0, 0, 345, 346, 5, 45, 0, 0, 346, 347, 3, 2, 1, 0, 347, 348, 5, 46, 0,
		0, 348, 366, 1, 0, 0, 0, 349, 351, 5, 8, 0, 0, 350, 349, 1, 0, 0, 0, 350,
		351, 1, 0, 0, 0, 351, 352, 1, 0, 0, 0, 352, 353, 5, 25, 0, 0, 353, 354,
		5, 41, 0, 0, 354, 355, 5, 43, 0, 0, 355, 356, 3, 42, 21, 0, 356, 359, 5,
		44, 0, 0, 357, 358, 5, 26, 0, 0, 358, 360, 3, 64, 32, 0, 359, 357, 1, 0,
		0, 0, 359, 360, 1, 0, 0, 0, 360, 361, 1, 0, 0, 0, 361, 362, 5, 45, 0, 0,
		362, 363, 3, 2, 1, 0, 363, 364, 5, 46, 0, 0, 364, 366, 1, 0, 0, 0, 365,
		335, 1, 0, 0, 0, 365, 350, 1, 0, 0, 0, 366, 41, 1, 0, 0, 0, 367, 368, 5,
		41, 0, 0, 368, 369, 5, 41, 0, 0, 369, 371, 5, 49, 0, 0, 370, 372, 5, 37,
		0, 0, 371, 370, 1, 0, 0, 0, 371, 372, 1, 0, 0, 0, 372, 373, 1, 0, 0, 0,
		373, 381, 3, 64, 32, 0, 374, 375, 5, 50, 0, 0, 375, 376, 5, 41, 0, 0, 376,
		377, 5, 41, 0, 0, 377, 378, 5, 49, 0, 0, 378, 380, 3, 64, 32, 0, 379, 374,
		1, 0, 0, 0, 380, 383, 1, 0, 0, 0, 381, 379, 1, 0, 0, 0, 381, 382, 1, 0,
		0, 0, 382, 493, 1, 0, 0, 0, 383, 381, 1, 0, 0, 0, 384, 385, 5, 38, 0, 0,
		385, 386, 5, 41, 0, 0, 386, 388, 5, 49, 0, 0, 387, 389, 5, 37, 0, 0, 388,
		387, 1, 0, 0, 0, 388, 389, 1, 0, 0, 0, 389, 390, 1, 0, 0, 0, 390, 398,
		3, 64, 32, 0, 391, 392, 5, 50, 0, 0, 392, 393, 5, 38, 0, 0, 393, 394, 5,
		41, 0, 0, 394, 395, 5, 49, 0, 0, 395, 397, 3, 64, 32, 0, 396, 391, 1, 0,
		0, 0, 397, 400, 1, 0, 0, 0, 398, 396, 1, 0, 0, 0, 398, 399, 1, 0, 0, 0,
		399, 493, 1, 0, 0, 0, 400, 398, 1, 0, 0, 0, 401, 402, 5, 41, 0, 0, 402,
		404, 5, 49, 0, 0, 403, 405, 5, 37, 0, 0, 404, 403, 1, 0, 0, 0, 404, 405,
		1, 0, 0, 0, 405, 406, 1, 0, 0, 0, 406, 416, 3, 64, 32, 0, 407, 408, 5,
		50, 0, 0, 408, 409, 5, 41, 0, 0, 409, 411, 5, 49, 0, 0, 410, 412, 5, 37,
		0, 0, 411, 410, 1, 0, 0, 0, 411, 412, 1, 0, 0, 0, 412, 413, 1, 0, 0, 0,
		413, 415, 3, 64, 32, 0, 414, 407, 1, 0, 0, 0, 415, 418, 1, 0, 0, 0, 416,
		414, 1, 0, 0, 0, 416, 417, 1, 0, 0, 0, 417, 493, 1, 0, 0, 0, 418, 416,
		1, 0, 0, 0, 419, 420, 5, 41, 0, 0, 420, 421, 5, 41, 0, 0, 421, 423, 5,
		49, 0, 0, 422, 424, 5, 37, 0, 0, 423, 422, 1, 0, 0, 0, 423, 424, 1, 0,
		0, 0, 424, 425, 1, 0, 0, 0, 425, 426, 5, 47, 0, 0, 426, 427, 3, 64, 32,
		0, 427, 441, 5, 48, 0, 0, 428, 429, 5, 50, 0, 0, 429, 430, 5, 41, 0, 0,
		430, 431, 5, 41, 0, 0, 431, 433, 5, 49, 0, 0, 432, 434, 5, 37, 0, 0, 433,
		432, 1, 0, 0, 0, 433, 434, 1, 0, 0, 0, 434, 435, 1, 0, 0, 0, 435, 436,
		5, 47, 0, 0, 436, 437, 3, 64, 32, 0, 437, 438, 5, 48, 0, 0, 438, 440, 1,
		0, 0, 0, 439, 428, 1, 0, 0, 0, 440, 443, 1, 0, 0, 0, 441, 439, 1, 0, 0,
		0, 441, 442, 1, 0, 0, 0, 442, 493, 1, 0, 0, 0, 443, 441, 1, 0, 0, 0, 444,
		445, 5, 38, 0, 0, 445, 446, 5, 41, 0, 0, 446, 448, 5, 49, 0, 0, 447, 449,
		5, 37, 0, 0, 448, 447, 1, 0, 0, 0, 448, 449, 1, 0, 0, 0, 449, 450, 1, 0,
		0, 0, 450, 451, 5, 47, 0, 0, 451, 452, 3, 64, 32, 0, 452, 466, 5, 48, 0,
		0, 453, 454, 5, 50, 0, 0, 454, 455, 5, 38, 0, 0, 455, 456, 5, 41, 0, 0,
		456, 458, 5, 49, 0, 0, 457, 459, 5, 37, 0, 0, 458, 457, 1, 0, 0, 0, 458,
		459, 1, 0, 0, 0, 459, 460, 1, 0, 0, 0, 460, 461, 5, 47, 0, 0, 461, 462,
		3, 64, 32, 0, 462, 463, 5, 48, 0, 0, 463, 465, 1, 0, 0, 0, 464, 453, 1,
		0, 0, 0, 465, 468, 1, 0, 0, 0, 466, 464, 1, 0, 0, 0, 466, 467, 1, 0, 0,
		0, 467, 493, 1, 0, 0, 0, 468, 466, 1, 0, 0, 0, 469, 470, 5, 41, 0, 0, 470,
		472, 5, 49, 0, 0, 471, 473, 5, 37, 0, 0, 472, 471, 1, 0, 0, 0, 472, 473,
		1, 0, 0, 0, 473, 474, 1, 0, 0, 0, 474, 475, 5, 47, 0, 0, 475, 476, 3, 64,
		32, 0, 476, 489, 5, 48, 0, 0, 477, 478, 5, 50, 0, 0, 478, 479, 5, 41, 0,
		0, 479, 481, 5, 49, 0, 0, 480, 482, 5, 37, 0, 0, 481, 480, 1, 0, 0, 0,
		481, 482, 1, 0, 0, 0, 482, 483, 1, 0, 0, 0, 483, 484, 5, 47, 0, 0, 484,
		485, 3, 64, 32, 0, 485, 486, 5, 48, 0, 0, 486, 488, 1, 0, 0, 0, 487, 477,
		1, 0, 0, 0, 488, 491, 1, 0, 0, 0, 489, 487, 1, 0, 0, 0, 489, 490, 1, 0,
		0, 0, 490, 493, 1, 0, 0, 0, 491, 489, 1, 0, 0, 0, 492, 367, 1, 0, 0, 0,
		492, 384, 1, 0, 0, 0, 492, 401, 1, 0, 0, 0, 492, 419, 1, 0, 0, 0, 492,
		444, 1, 0, 0, 0, 492, 469, 1, 0, 0, 0, 493, 43, 1, 0, 0, 0, 494, 495, 5,
		41, 0, 0, 495, 496, 5, 43, 0, 0, 496, 503, 5, 44, 0, 0, 497, 498, 5, 41,
		0, 0, 498, 499, 5, 43, 0, 0, 499, 500, 3, 46, 23, 0, 500, 501, 5, 44, 0,
		0, 501, 503, 1, 0, 0, 0, 502, 494, 1, 0, 0, 0, 502, 497, 1, 0, 0, 0, 503,
		45, 1, 0, 0, 0, 504, 506, 5, 36, 0, 0, 505, 504, 1, 0, 0, 0, 505, 506,
		1, 0, 0, 0, 506, 507, 1, 0, 0, 0, 507, 508, 5, 41, 0, 0, 508, 509, 5, 49,
		0, 0, 509, 519, 3, 62, 31, 0, 510, 512, 5, 50, 0, 0, 511, 513, 5, 36, 0,
		0, 512, 511, 1, 0, 0, 0, 512, 513, 1, 0, 0, 0, 513, 514, 1, 0, 0, 0, 514,
		515, 5, 41, 0, 0, 515, 516, 5, 49, 0, 0, 516, 518, 3, 62, 31, 0, 517, 510,
		1, 0, 0, 0, 518, 521, 1, 0, 0, 0, 519, 517, 1, 0, 0, 0, 519, 520, 1, 0,
		0, 0, 520, 537, 1, 0, 0, 0, 521, 519, 1, 0, 0, 0, 522, 524, 5, 36, 0, 0,
		523, 522, 1, 0, 0, 0, 523, 524, 1, 0, 0, 0, 524, 525, 1, 0, 0, 0, 525,
		533, 3, 62, 31, 0, 526, 528, 5, 50, 0, 0, 527, 529, 5, 36, 0, 0, 528, 527,
		1, 0, 0, 0, 528, 529, 1, 0, 0, 0, 529, 530, 1, 0, 0, 0, 530, 532, 3, 62,
		31, 0, 531, 526, 1, 0, 0, 0, 532, 535, 1, 0, 0, 0, 533, 531, 1, 0, 0, 0,
		533, 534, 1, 0, 0, 0, 534, 537, 1, 0, 0, 0, 535, 533, 1, 0, 0, 0, 536,
		505, 1, 0, 0, 0, 536, 523, 1, 0, 0, 0, 537, 47, 1, 0, 0, 0, 538, 539, 5,
		41, 0, 0, 539, 540, 5, 28, 0, 0, 540, 541, 5, 9, 0, 0, 541, 542, 5, 43,
		0, 0, 542, 543, 3, 62, 31, 0, 543, 544, 5, 43, 0, 0, 544, 594, 1, 0, 0,
		0, 545, 546, 5, 41, 0, 0, 546, 547, 5, 28, 0, 0, 547, 548, 5, 11, 0, 0,
		548, 549, 5, 43, 0, 0, 549, 594, 5, 44, 0, 0, 550, 551, 5, 41, 0, 0, 551,
		552, 5, 28, 0, 0, 552, 553, 5, 12, 0, 0, 553, 554, 5, 43, 0, 0, 554, 555,
		5, 15, 0, 0, 555, 556, 5, 49, 0, 0, 556, 557, 3, 62, 31, 0, 557, 558, 5,
		44, 0, 0, 558, 594, 1, 0, 0, 0, 559, 560, 5, 41, 0, 0, 560, 561, 5, 28,
		0, 0, 561, 562, 5, 13, 0, 0, 562, 563, 5, 43, 0, 0, 563, 594, 5, 44, 0,
		0, 564, 565, 5, 41, 0, 0, 565, 566, 5, 28, 0, 0, 566, 594, 5, 10, 0, 0,
		567, 568, 5, 41, 0, 0, 568, 569, 5, 47, 0, 0, 569, 570, 3, 62, 31, 0, 570,
		571, 5, 48, 0, 0, 571, 594, 1, 0, 0, 0, 572, 573, 5, 14, 0, 0, 573, 574,
		5, 28, 0, 0, 574, 577, 5, 41, 0, 0, 575, 576, 5, 52, 0, 0, 576, 578, 3,
		62, 31, 0, 577, 575, 1, 0, 0, 0, 577, 578, 1, 0, 0, 0, 578, 594, 1, 0,
		0, 0, 579, 580, 5, 41, 0, 0, 580, 581, 5, 28, 0, 0, 581, 584, 5, 41, 0,
		0, 582, 583, 5, 52, 0, 0, 583, 585, 3, 62, 31, 0, 584, 582, 1, 0, 0, 0,
		584, 585, 1, 0, 0, 0, 585, 594, 1, 0, 0, 0, 586, 587, 5, 41, 0, 0, 587,
		588, 5, 28, 0, 0, 588, 589, 5, 41, 0, 0, 589, 590, 5, 43, 0, 0, 590, 591,
		3, 42, 21, 0, 591, 592, 5, 44, 0, 0, 592, 594, 1, 0, 0, 0, 593, 538, 1,
		0, 0, 0, 593, 545, 1, 0, 0, 0, 593, 550, 1, 0, 0, 0, 593, 559, 1, 0, 0,
		0, 593, 564, 1, 0, 0, 0, 593, 567, 1, 0, 0, 0, 593, 572, 1, 0, 0, 0, 593,
		579, 1, 0, 0, 0, 593, 586, 1, 0, 0, 0, 594, 49, 1, 0, 0, 0, 595, 599, 3,
		56, 28, 0, 596, 599, 3, 58, 29, 0, 597, 599, 3, 60, 30, 0, 598, 595, 1,
		0, 0, 0, 598, 596, 1, 0, 0, 0, 598, 597, 1, 0, 0, 0, 599, 51, 1, 0, 0,
		0, 600, 601, 5, 30, 0, 0, 601, 602, 5, 43, 0, 0, 602, 603, 3, 54, 27, 0,
		603, 604, 5, 44, 0, 0, 604, 53, 1, 0, 0, 0, 605, 610, 3, 62, 31, 0, 606,
		607, 5, 50, 0, 0, 607, 609, 3, 62, 31, 0, 608, 606, 1, 0, 0, 0, 609, 612,
		1, 0, 0, 0, 610, 608, 1, 0, 0, 0, 610, 611, 1, 0, 0, 0, 611, 55, 1, 0,
		0, 0, 612, 610, 1, 0, 0, 0, 613, 614, 5, 1, 0, 0, 614, 615, 5, 43, 0, 0,
		615, 616, 3, 62, 31, 0, 616, 617, 5, 44, 0, 0, 617, 57, 1, 0, 0, 0, 618,
		619, 5, 2, 0, 0, 619, 620, 5, 43, 0, 0, 620, 621, 3, 62, 31, 0, 621, 622,
		5, 44, 0, 0, 622, 59, 1, 0, 0, 0, 623, 624, 5, 3, 0, 0, 624, 625, 5, 43,
		0, 0, 625, 626, 3, 62, 31, 0, 626, 627, 5, 44, 0, 0, 627, 61, 1, 0, 0,
		0, 628, 629, 6, 31, -1, 0, 629, 630, 5, 42, 0, 0, 630, 657, 3, 62, 31,
		17, 631, 632, 5, 57, 0, 0, 632, 657, 3, 62, 31, 16, 633, 634, 5, 43, 0,
		0, 634, 635, 3, 62, 31, 0, 635, 636, 5, 44, 0, 0, 636, 657, 1, 0, 0, 0,
		637, 639, 3, 44, 22, 0, 638, 640, 5, 51, 0, 0, 639, 638, 1, 0, 0, 0, 639,
		640, 1, 0, 0, 0, 640, 657, 1, 0, 0, 0, 641, 643, 3, 48, 24, 0, 642, 644,
		5, 51, 0, 0, 643, 642, 1, 0, 0, 0, 643, 644, 1, 0, 0, 0, 644, 657, 1, 0,
		0, 0, 645, 657, 3, 50, 25, 0, 646, 647, 5, 41, 0, 0, 647, 648, 5, 43, 0,
		0, 648, 649, 3, 16, 8, 0, 649, 650, 5, 44, 0, 0, 650, 657, 1, 0, 0, 0,
		651, 657, 5, 39, 0, 0, 652, 657, 5, 40, 0, 0, 653, 657, 5, 41, 0, 0, 654,
		657, 5, 33, 0, 0, 655, 657, 7, 1, 0, 0, 656, 628, 1, 0, 0, 0, 656, 631,
		1, 0, 0, 0, 656, 633, 1, 0, 0, 0, 656, 637, 1, 0, 0, 0, 656, 641, 1, 0,
		0, 0, 656, 645, 1, 0, 0, 0, 656, 646, 1, 0, 0, 0, 656, 651, 1, 0, 0, 0,
		656, 652, 1, 0, 0, 0, 656, 653, 1, 0, 0, 0, 656, 654, 1, 0, 0, 0, 656,
		655, 1, 0, 0, 0, 657, 675, 1, 0, 0, 0, 658, 659, 10, 15, 0, 0, 659, 660,
		7, 2, 0, 0, 660, 674, 3, 62, 31, 16, 661, 662, 10, 14, 0, 0, 662, 663,
		7, 3, 0, 0, 663, 674, 3, 62, 31, 15, 664, 665, 10, 13, 0, 0, 665, 666,
		7, 4, 0, 0, 666, 674, 3, 62, 31, 14, 667, 668, 10, 12, 0, 0, 668, 669,
		7, 5, 0, 0, 669, 674, 3, 62, 31, 13, 670, 671, 10, 11, 0, 0, 671, 672,
		7, 6, 0, 0, 672, 674, 3, 62, 31, 12, 673, 658, 1, 0, 0, 0, 673, 661, 1,
		0, 0, 0, 673, 664, 1, 0, 0, 0, 673, 667, 1, 0, 0, 0, 673, 670, 1, 0, 0,
		0, 674, 677, 1, 0, 0, 0, 675, 673, 1, 0, 0, 0, 675, 676, 1, 0, 0, 0, 676,
		63, 1, 0, 0, 0, 677, 675, 1, 0, 0, 0, 678, 679, 7, 7, 0, 0, 679, 65, 1,
		0, 0, 0, 65, 72, 77, 81, 85, 97, 101, 104, 108, 112, 116, 119, 121, 132,
		137, 168, 170, 181, 214, 216, 250, 262, 272, 280, 284, 319, 335, 343, 350,
		359, 365, 371, 381, 388, 398, 404, 411, 416, 423, 433, 441, 448, 458, 466,
		472, 481, 489, 492, 502, 505, 512, 519, 523, 528, 533, 536, 577, 584, 593,
		598, 610, 639, 643, 656, 673, 675,
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
	SFGrammarParserRULE_structCallList       = 8
	SFGrammarParserRULE_declaration          = 9
	SFGrammarParserRULE_type_declaration     = 10
	SFGrammarParserRULE_assignment           = 11
	SFGrammarParserRULE_ifstmt               = 12
	SFGrammarParserRULE_switchStmt           = 13
	SFGrammarParserRULE_caseBlock            = 14
	SFGrammarParserRULE_defaultBlock         = 15
	SFGrammarParserRULE_whileStmt            = 16
	SFGrammarParserRULE_forStmt              = 17
	SFGrammarParserRULE_forRange             = 18
	SFGrammarParserRULE_guardStmt            = 19
	SFGrammarParserRULE_functionStmt         = 20
	SFGrammarParserRULE_listFunctionParams   = 21
	SFGrammarParserRULE_callFunctionStmt     = 22
	SFGrammarParserRULE_listCallFunctionStmt = 23
	SFGrammarParserRULE_callBack             = 24
	SFGrammarParserRULE_embbededFunc         = 25
	SFGrammarParserRULE_printstmt            = 26
	SFGrammarParserRULE_exprList             = 27
	SFGrammarParserRULE_intstmt              = 28
	SFGrammarParserRULE_floatstmt            = 29
	SFGrammarParserRULE_stringstmt           = 30
	SFGrammarParserRULE_expr                 = 31
	SFGrammarParserRULE_type                 = 32
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
		p.SetState(66)
		p.Block()
	}
	{
		p.SetState(67)
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
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2252239684046) != 0 {
		{
			p.SetState(69)
			p.Stmts()
		}

		p.SetState(74)
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

	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)
			p.Declaration()
		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserSEMICOLON {
			{
				p.SetState(76)
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
			p.SetState(79)
			p.Assignment()
		}
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserSEMICOLON {
			{
				p.SetState(80)
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
			p.SetState(83)
			p.EmbbededFunc()
		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserSEMICOLON {
			{
				p.SetState(84)
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
			p.SetState(87)
			p.Ifstmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(88)
			p.SwitchStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(89)
			p.WhileStmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(90)
			p.ForStmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(91)
			p.GuardStmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(92)
			p.TransferStmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(93)
			p.FunctionStmt()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(94)
			p.Printstmt()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(95)
			p.CallFunctionStmt()
		}
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserSEMICOLON {
			{
				p.SetState(96)
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
			p.SetState(99)
			p.CallBack()
		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserSEMICOLON {
			{
				p.SetState(100)
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
			p.SetState(103)
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

	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SFGrammarParserBREAK:
		localctx = NewBreakStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(106)
			p.Match(SFGrammarParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserSEMICOLON {
			{
				p.SetState(107)
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
			p.SetState(110)
			p.Match(SFGrammarParserCONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserSEMICOLON {
			{
				p.SetState(111)
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
			p.SetState(114)
			p.Match(SFGrammarParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(116)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(115)
				p.expr(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserSEMICOLON {
			{
				p.SetState(118)
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
		p.SetState(123)
		p.Match(SFGrammarParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Match(SFGrammarParserID_PRIMITIVE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.Match(SFGrammarParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)
		p.StructBlock()
	}
	{
		p.SetState(127)
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
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SFGrammarParserDECLARATION_VAR || _la == SFGrammarParserDECLARATION_LET {
		{
			p.SetState(129)
			p.StructStmts()
		}

		p.SetState(134)
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

func (s *StructStmtsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructStmtsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.DeclarationStructs()
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SFGrammarParserSEMICOLON {
		{
			p.SetState(136)
			p.Match(SFGrammarParserSEMICOLON)
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
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStructDeclarationWithValueAndTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(139)
			p.Type_declaration()
		}
		{
			p.SetState(140)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.Type_()
		}
		{
			p.SetState(143)
			p.Match(SFGrammarParserIS_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)
			p.expr(0)
		}

	case 2:
		localctx = NewStructDeclarationWithoutValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(146)
			p.Type_declaration()
		}
		{
			p.SetState(147)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(148)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)
			p.Type_()
		}

	case 3:
		localctx = NewStructDeclarationImplicitValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(151)
			p.Type_declaration()
		}
		{
			p.SetState(152)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)
			p.Match(SFGrammarParserIS_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.expr(0)
		}

	case 4:
		localctx = NewStructDeclarationVectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(156)
			p.Type_declaration()
		}
		{
			p.SetState(157)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(159)
			p.Match(SFGrammarParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.Type_()
		}
		{
			p.SetState(161)
			p.Match(SFGrammarParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)
			p.Match(SFGrammarParserIS_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case SFGrammarParserLBRACKET:
			{
				p.SetState(163)
				p.Match(SFGrammarParserLBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(164)
				p.ExprList()
			}
			{
				p.SetState(165)
				p.Match(SFGrammarParserRBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case SFGrammarParserID_PRIMITIVE:
			{
				p.SetState(167)
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
	p.EnterRule(localctx, 16, SFGrammarParserRULE_structCallList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Match(SFGrammarParserID_PRIMITIVE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.Match(SFGrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(174)
		p.expr(0)
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SFGrammarParserCOMMA {
		{
			p.SetState(175)
			p.Match(SFGrammarParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)
			p.expr(0)
		}

		p.SetState(183)
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

func (s *ValueDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitValueDeclaration(s)

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
	p.EnterRule(localctx, 18, SFGrammarParserRULE_declaration)
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeValueDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(184)
			p.Type_declaration()
		}
		{
			p.SetState(185)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Type_()
		}
		{
			p.SetState(188)
			p.Match(SFGrammarParserIS_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.expr(0)
		}

	case 2:
		localctx = NewTypeOptionalValueDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(191)
			p.Type_declaration()
		}
		{
			p.SetState(192)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)
			p.Type_()
		}
		{
			p.SetState(195)
			p.Match(SFGrammarParserQUESTION_MARK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewValueDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(197)
			p.Type_declaration()
		}
		{
			p.SetState(198)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)
			p.Match(SFGrammarParserIS_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)
			p.expr(0)
		}

	case 4:
		localctx = NewVectorDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(202)
			p.Type_declaration()
		}
		{
			p.SetState(203)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(205)
			p.Match(SFGrammarParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(206)
			p.Type_()
		}
		{
			p.SetState(207)
			p.Match(SFGrammarParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(208)
			p.Match(SFGrammarParserIS_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case SFGrammarParserLBRACKET:
			{
				p.SetState(209)
				p.Match(SFGrammarParserLBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(210)
				p.ExprList()
			}
			{
				p.SetState(211)
				p.Match(SFGrammarParserRBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case SFGrammarParserID_PRIMITIVE:
			{
				p.SetState(213)
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
	p.EnterRule(localctx, 20, SFGrammarParserRULE_type_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
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
	p.EnterRule(localctx, 22, SFGrammarParserRULE_assignment)
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewValueAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(220)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(221)
			p.Match(SFGrammarParserIS_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(222)
			p.expr(0)
		}

	case 2:
		localctx = NewPlusAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(223)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)
			p.Match(SFGrammarParserPLUS_IS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(225)
			p.expr(0)
		}

	case 3:
		localctx = NewMinusAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(226)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(227)
			p.Match(SFGrammarParserMINUS_IS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(228)
			p.expr(0)
		}

	case 4:
		localctx = NewVectorAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(229)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(230)
			p.Match(SFGrammarParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)
			p.expr(0)
		}
		{
			p.SetState(232)
			p.Match(SFGrammarParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)
			p.Match(SFGrammarParserIS_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(234)
			p.expr(0)
		}

	case 5:
		localctx = NewVectorMinusAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
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
			p.Match(SFGrammarParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.expr(0)
		}
		{
			p.SetState(239)
			p.Match(SFGrammarParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(240)
			p.Match(SFGrammarParserMINUS_IS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(241)
			p.expr(0)
		}

	case 6:
		localctx = NewVectorPlusAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
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
			p.Match(SFGrammarParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(245)
			p.expr(0)
		}
		{
			p.SetState(246)
			p.Match(SFGrammarParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(247)
			p.Match(SFGrammarParserPLUS_IS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(248)
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
	p.EnterRule(localctx, 24, SFGrammarParserRULE_ifstmt)
	var _la int

	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIfElseStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(252)
			p.Match(SFGrammarParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(253)
			p.expr(0)
		}
		{
			p.SetState(254)
			p.Match(SFGrammarParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)
			p.Block()
		}
		{
			p.SetState(256)
			p.Match(SFGrammarParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserELSE {
			{
				p.SetState(257)
				p.Match(SFGrammarParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(258)
				p.Match(SFGrammarParserLBRACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(259)
				p.Block()
			}
			{
				p.SetState(260)
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
			p.SetState(264)
			p.Match(SFGrammarParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(265)
			p.expr(0)
		}
		{
			p.SetState(266)
			p.Match(SFGrammarParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(267)
			p.Block()
		}
		{
			p.SetState(268)
			p.Match(SFGrammarParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(269)
			p.Match(SFGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(270)
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
	p.EnterRule(localctx, 26, SFGrammarParserRULE_switchStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(274)
		p.Match(SFGrammarParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(275)
		p.expr(0)
	}
	{
		p.SetState(276)
		p.Match(SFGrammarParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SFGrammarParserCASE {
		{
			p.SetState(277)
			p.CaseBlock()
		}

		p.SetState(282)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SFGrammarParserDEFAULT {
		{
			p.SetState(283)
			p.DefaultBlock()
		}

	}
	{
		p.SetState(286)
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
	p.EnterRule(localctx, 28, SFGrammarParserRULE_caseBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(288)
		p.Match(SFGrammarParserCASE)
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
		p.Match(SFGrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(291)
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
	p.EnterRule(localctx, 30, SFGrammarParserRULE_defaultBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.Match(SFGrammarParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(294)
		p.Match(SFGrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(295)
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
	p.EnterRule(localctx, 32, SFGrammarParserRULE_whileStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		p.Match(SFGrammarParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(298)
		p.expr(0)
	}
	{
		p.SetState(299)
		p.Match(SFGrammarParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(300)
		p.Block()
	}
	{
		p.SetState(301)
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

func (s *ForRangeExprContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
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
	p.EnterRule(localctx, 34, SFGrammarParserRULE_forStmt)
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForRangeExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(303)
			p.Match(SFGrammarParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(304)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(305)
			p.Match(SFGrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(306)
			p.ForRange()
		}
		{
			p.SetState(307)
			p.Match(SFGrammarParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(308)
			p.Block()
		}
		{
			p.SetState(309)
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
			p.SetState(311)
			p.Match(SFGrammarParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(312)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(313)
			p.Match(SFGrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(314)
			p.expr(0)
		}
		{
			p.SetState(315)
			p.Match(SFGrammarParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(316)
			p.Block()
		}
		{
			p.SetState(317)
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
	p.EnterRule(localctx, 36, SFGrammarParserRULE_forRange)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(321)

		var _x = p.expr(0)

		localctx.(*ForRangeContext).left = _x
	}
	{
		p.SetState(322)
		p.Match(SFGrammarParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(323)
		p.Match(SFGrammarParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(324)
		p.Match(SFGrammarParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(325)

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
	p.EnterRule(localctx, 38, SFGrammarParserRULE_guardStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(327)
		p.Match(SFGrammarParserGUARD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(328)
		p.expr(0)
	}
	{
		p.SetState(329)
		p.Match(SFGrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(330)
		p.Match(SFGrammarParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(331)
		p.Block()
	}
	{
		p.SetState(332)
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

func (s *FunctionWithParamsContext) MUTATING() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserMUTATING, 0)
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

func (s *FunctionWithoutParamsContext) MUTATING() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserMUTATING, 0)
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
	p.EnterRule(localctx, 40, SFGrammarParserRULE_functionStmt)
	var _la int

	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFunctionWithoutParamsContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(335)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserMUTATING {
			{
				p.SetState(334)
				p.Match(SFGrammarParserMUTATING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(337)
			p.Match(SFGrammarParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(340)
			p.Match(SFGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(343)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserARROW_FUNCTION {
			{
				p.SetState(341)
				p.Match(SFGrammarParserARROW_FUNCTION)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(342)
				p.Type_()
			}

		}
		{
			p.SetState(345)
			p.Match(SFGrammarParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(346)
			p.Block()
		}
		{
			p.SetState(347)
			p.Match(SFGrammarParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewFunctionWithParamsContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(350)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserMUTATING {
			{
				p.SetState(349)
				p.Match(SFGrammarParserMUTATING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(352)
			p.Match(SFGrammarParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(353)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(354)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(355)
			p.ListFunctionParams()
		}
		{
			p.SetState(356)
			p.Match(SFGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(359)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserARROW_FUNCTION {
			{
				p.SetState(357)
				p.Match(SFGrammarParserARROW_FUNCTION)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(358)
				p.Type_()
			}

		}
		{
			p.SetState(361)
			p.Match(SFGrammarParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(362)
			p.Block()
		}
		{
			p.SetState(363)
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

func (s *ListFunctionParamsEIVectorContext) AllLBRACKET() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserLBRACKET)
}

func (s *ListFunctionParamsEIVectorContext) LBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLBRACKET, i)
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

func (s *ListFunctionParamsEIVectorContext) AllRBRACKET() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserRBRACKET)
}

func (s *ListFunctionParamsEIVectorContext) RBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRBRACKET, i)
}

func (s *ListFunctionParamsEIVectorContext) AllINOUT() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserINOUT)
}

func (s *ListFunctionParamsEIVectorContext) INOUT(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserINOUT, i)
}

func (s *ListFunctionParamsEIVectorContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserCOMMA)
}

func (s *ListFunctionParamsEIVectorContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOMMA, i)
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

func (s *ListFunctionParamsNEIVectorContext) AllLBRACKET() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserLBRACKET)
}

func (s *ListFunctionParamsNEIVectorContext) LBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLBRACKET, i)
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

func (s *ListFunctionParamsNEIVectorContext) AllRBRACKET() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserRBRACKET)
}

func (s *ListFunctionParamsNEIVectorContext) RBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRBRACKET, i)
}

func (s *ListFunctionParamsNEIVectorContext) AllINOUT() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserINOUT)
}

func (s *ListFunctionParamsNEIVectorContext) INOUT(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserINOUT, i)
}

func (s *ListFunctionParamsNEIVectorContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserCOMMA)
}

func (s *ListFunctionParamsNEIVectorContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOMMA, i)
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

func (s *ListFunctionParamsBEIVectorContext) AllLBRACKET() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserLBRACKET)
}

func (s *ListFunctionParamsBEIVectorContext) LBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLBRACKET, i)
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

func (s *ListFunctionParamsBEIVectorContext) AllRBRACKET() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserRBRACKET)
}

func (s *ListFunctionParamsBEIVectorContext) RBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRBRACKET, i)
}

func (s *ListFunctionParamsBEIVectorContext) AllINOUT() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserINOUT)
}

func (s *ListFunctionParamsBEIVectorContext) INOUT(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserINOUT, i)
}

func (s *ListFunctionParamsBEIVectorContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserCOMMA)
}

func (s *ListFunctionParamsBEIVectorContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserCOMMA, i)
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
	p.EnterRule(localctx, 42, SFGrammarParserRULE_listFunctionParams)
	var _la int

	p.SetState(492)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		localctx = NewListFunctionParamsEIContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
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
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(369)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(371)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserINOUT {
			{
				p.SetState(370)
				p.Match(SFGrammarParserINOUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(373)
			p.Type_()
		}
		p.SetState(381)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SFGrammarParserCOMMA {
			{
				p.SetState(374)
				p.Match(SFGrammarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(375)
				p.Match(SFGrammarParserID_PRIMITIVE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(376)
				p.Match(SFGrammarParserID_PRIMITIVE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(377)
				p.Match(SFGrammarParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(378)
				p.Type_()
			}

			p.SetState(383)
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
			p.SetState(384)
			p.Match(SFGrammarParserNOT_PARAM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(385)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(386)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(388)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserINOUT {
			{
				p.SetState(387)
				p.Match(SFGrammarParserINOUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(390)
			p.Type_()
		}
		p.SetState(398)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SFGrammarParserCOMMA {
			{
				p.SetState(391)
				p.Match(SFGrammarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(392)
				p.Match(SFGrammarParserNOT_PARAM)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(393)
				p.Match(SFGrammarParserID_PRIMITIVE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(394)
				p.Match(SFGrammarParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(395)
				p.Type_()
			}

			p.SetState(400)
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
			p.SetState(401)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(402)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(404)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserINOUT {
			{
				p.SetState(403)
				p.Match(SFGrammarParserINOUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(406)
			p.Type_()
		}
		p.SetState(416)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SFGrammarParserCOMMA {
			{
				p.SetState(407)
				p.Match(SFGrammarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(408)
				p.Match(SFGrammarParserID_PRIMITIVE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(409)
				p.Match(SFGrammarParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(411)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SFGrammarParserINOUT {
				{
					p.SetState(410)
					p.Match(SFGrammarParserINOUT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(413)
				p.Type_()
			}

			p.SetState(418)
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
			p.SetState(419)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(420)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(421)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(423)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserINOUT {
			{
				p.SetState(422)
				p.Match(SFGrammarParserINOUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(425)
			p.Match(SFGrammarParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(426)
			p.Type_()
		}
		{
			p.SetState(427)
			p.Match(SFGrammarParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(441)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SFGrammarParserCOMMA {
			{
				p.SetState(428)
				p.Match(SFGrammarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(429)
				p.Match(SFGrammarParserID_PRIMITIVE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(430)
				p.Match(SFGrammarParserID_PRIMITIVE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(431)
				p.Match(SFGrammarParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(433)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SFGrammarParserINOUT {
				{
					p.SetState(432)
					p.Match(SFGrammarParserINOUT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(435)
				p.Match(SFGrammarParserLBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(436)
				p.Type_()
			}
			{
				p.SetState(437)
				p.Match(SFGrammarParserRBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(443)
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
			p.SetState(444)
			p.Match(SFGrammarParserNOT_PARAM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(445)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(446)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(448)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserINOUT {
			{
				p.SetState(447)
				p.Match(SFGrammarParserINOUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(450)
			p.Match(SFGrammarParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(451)
			p.Type_()
		}
		{
			p.SetState(452)
			p.Match(SFGrammarParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(466)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SFGrammarParserCOMMA {
			{
				p.SetState(453)
				p.Match(SFGrammarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(454)
				p.Match(SFGrammarParserNOT_PARAM)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(455)
				p.Match(SFGrammarParserID_PRIMITIVE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(456)
				p.Match(SFGrammarParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(458)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SFGrammarParserINOUT {
				{
					p.SetState(457)
					p.Match(SFGrammarParserINOUT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(460)
				p.Match(SFGrammarParserLBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(461)
				p.Type_()
			}
			{
				p.SetState(462)
				p.Match(SFGrammarParserRBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(468)
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
		{
			p.SetState(474)
			p.Match(SFGrammarParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(475)
			p.Type_()
		}
		{
			p.SetState(476)
			p.Match(SFGrammarParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(489)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SFGrammarParserCOMMA {
			{
				p.SetState(477)
				p.Match(SFGrammarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(478)
				p.Match(SFGrammarParserID_PRIMITIVE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(479)
				p.Match(SFGrammarParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(481)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SFGrammarParserINOUT {
				{
					p.SetState(480)
					p.Match(SFGrammarParserINOUT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(483)
				p.Match(SFGrammarParserLBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(484)
				p.Type_()
			}
			{
				p.SetState(485)
				p.Match(SFGrammarParserRBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(491)
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
	p.EnterRule(localctx, 44, SFGrammarParserRULE_callFunctionStmt)
	p.SetState(502)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCallFunctionWithoutParamsContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(494)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(495)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(496)
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
			p.SetState(497)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(498)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(499)
			p.ListCallFunctionStmt()
		}
		{
			p.SetState(500)
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
	p.EnterRule(localctx, 46, SFGrammarParserRULE_listCallFunctionStmt)
	var _la int

	p.SetState(536)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 54, p.GetParserRuleContext()) {
	case 1:
		localctx = NewListCallFunctionStmtEIContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(505)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserREFERENCE {
			{
				p.SetState(504)
				p.Match(SFGrammarParserREFERENCE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(507)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(508)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(509)
			p.expr(0)
		}
		p.SetState(519)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SFGrammarParserCOMMA {
			{
				p.SetState(510)
				p.Match(SFGrammarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(512)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SFGrammarParserREFERENCE {
				{
					p.SetState(511)
					p.Match(SFGrammarParserREFERENCE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(514)
				p.Match(SFGrammarParserID_PRIMITIVE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(515)
				p.Match(SFGrammarParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(516)
				p.expr(0)
			}

			p.SetState(521)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		localctx = NewListCallFunctionStmtNEIContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(523)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserREFERENCE {
			{
				p.SetState(522)
				p.Match(SFGrammarParserREFERENCE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(525)
			p.expr(0)
		}
		p.SetState(533)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SFGrammarParserCOMMA {
			{
				p.SetState(526)
				p.Match(SFGrammarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(528)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SFGrammarParserREFERENCE {
				{
					p.SetState(527)
					p.Match(SFGrammarParserREFERENCE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(530)
				p.expr(0)
			}

			p.SetState(535)
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

func (s *AppendVectorContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserLPAREN)
}

func (s *AppendVectorContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLPAREN, i)
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

func (s *RemoveLastVectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitRemoveLastVector(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructFunctionContext struct {
	CallBackContext
}

func NewStructFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructFunctionContext {
	var p = new(StructFunctionContext)

	InitEmptyCallBackContext(&p.CallBackContext)
	p.parser = parser
	p.CopyAll(ctx.(*CallBackContext))

	return p
}

func (s *StructFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructFunctionContext) AllID_PRIMITIVE() []antlr.TerminalNode {
	return s.GetTokens(SFGrammarParserID_PRIMITIVE)
}

func (s *StructFunctionContext) ID_PRIMITIVE(i int) antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, i)
}

func (s *StructFunctionContext) DOT() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserDOT, 0)
}

func (s *StructFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLPAREN, 0)
}

func (s *StructFunctionContext) ListFunctionParams() IListFunctionParamsContext {
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

func (s *StructFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRPAREN, 0)
}

func (s *StructFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitStructFunction(s)

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

func (s *StructAttributeContext) DOT() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserDOT, 0)
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

func (s *IsEmptyVectorContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLPAREN, 0)
}

func (s *IsEmptyVectorContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRPAREN, 0)
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
	p.EnterRule(localctx, 48, SFGrammarParserRULE_callBack)
	p.SetState(593)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAppendVectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(538)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(539)
			p.Match(SFGrammarParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(540)
			p.Match(SFGrammarParserAPPEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(541)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(542)
			p.expr(0)
		}
		{
			p.SetState(543)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewRemoveLastVectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
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
			p.Match(SFGrammarParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(547)
			p.Match(SFGrammarParserREMOVELAST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(548)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(549)
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
			p.SetState(550)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(551)
			p.Match(SFGrammarParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(552)
			p.Match(SFGrammarParserREMOVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(553)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(554)
			p.Match(SFGrammarParserAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(555)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(556)
			p.expr(0)
		}
		{
			p.SetState(557)
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
			p.SetState(559)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(560)
			p.Match(SFGrammarParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(561)
			p.Match(SFGrammarParserISEMPTY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(562)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(563)
			p.Match(SFGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewCountVectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
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
			p.Match(SFGrammarParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(566)
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
			p.SetState(567)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(568)
			p.Match(SFGrammarParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(569)
			p.expr(0)
		}
		{
			p.SetState(570)
			p.Match(SFGrammarParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewSelfFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(572)
			p.Match(SFGrammarParserSELF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(573)
			p.Match(SFGrammarParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(574)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(577)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 55, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(575)
				p.Match(SFGrammarParserIS_)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(576)
				p.expr(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 8:
		localctx = NewStructAttributeContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(579)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(580)
			p.Match(SFGrammarParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(581)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(584)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(582)
				p.Match(SFGrammarParserIS_)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(583)
				p.expr(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 9:
		localctx = NewStructFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(586)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(587)
			p.Match(SFGrammarParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(588)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(589)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(590)
			p.ListFunctionParams()
		}
		{
			p.SetState(591)
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
	p.EnterRule(localctx, 50, SFGrammarParserRULE_embbededFunc)
	p.SetState(598)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SFGrammarParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(595)
			p.Intstmt()
		}

	case SFGrammarParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(596)
			p.Floatstmt()
		}

	case SFGrammarParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(597)
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
	p.EnterRule(localctx, 52, SFGrammarParserRULE_printstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(600)
		p.Match(SFGrammarParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(601)
		p.Match(SFGrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(602)
		p.ExprList()
	}
	{
		p.SetState(603)
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
	p.EnterRule(localctx, 54, SFGrammarParserRULE_exprList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(605)
		p.expr(0)
	}
	p.SetState(610)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SFGrammarParserCOMMA {
		{
			p.SetState(606)
			p.Match(SFGrammarParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(607)
			p.expr(0)
		}

		p.SetState(612)
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
	p.EnterRule(localctx, 56, SFGrammarParserRULE_intstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(613)
		p.Match(SFGrammarParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(614)
		p.Match(SFGrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(615)
		p.expr(0)
	}
	{
		p.SetState(616)
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
	p.EnterRule(localctx, 58, SFGrammarParserRULE_floatstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(618)
		p.Match(SFGrammarParserFLOAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(619)
		p.Match(SFGrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(620)
		p.expr(0)
	}
	{
		p.SetState(621)
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
	p.EnterRule(localctx, 60, SFGrammarParserRULE_stringstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(623)
		p.Match(SFGrammarParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(624)
		p.Match(SFGrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(625)
		p.expr(0)
	}
	{
		p.SetState(626)
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

func (s *ComparationOperationExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitComparationOperationExpr(s)

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

func (s *ArithmeticOperationExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitArithmeticOperationExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructCallExprContext struct {
	ExprContext
}

func NewStructCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructCallExprContext {
	var p = new(StructCallExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StructCallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructCallExprContext) ID_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserID_PRIMITIVE, 0)
}

func (s *StructCallExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserLPAREN, 0)
}

func (s *StructCallExprContext) StructCallList() IStructCallListContext {
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

func (s *StructCallExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserRPAREN, 0)
}

func (s *StructCallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFGrammarVisitor:
		return t.VisitStructCallExpr(s)

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
	_startState := 62
	p.EnterRecursionRule(localctx, 62, SFGrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(656)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 62, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(629)
			p.Match(SFGrammarParserNEGATION_OPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(630)

			var _x = p.expr(17)

			localctx.(*NotExprContext).right = _x
		}

	case 2:
		localctx = NewNegExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(631)
			p.Match(SFGrammarParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(632)

			var _x = p.expr(16)

			localctx.(*NegExprContext).right = _x
		}

	case 3:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(633)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(634)
			p.expr(0)
		}
		{
			p.SetState(635)
			p.Match(SFGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewCallFunctionExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(637)
			p.CallFunctionStmt()
		}
		p.SetState(639)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(638)
				p.Match(SFGrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 5:
		localctx = NewCallBackExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(641)
			p.CallBack()
		}
		p.SetState(643)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 61, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(642)
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
		localctx = NewEmbeddedFunctionExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(645)
			p.EmbbededFunc()
		}

	case 7:
		localctx = NewStructCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(646)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(647)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(648)
			p.StructCallList()
		}
		{
			p.SetState(649)
			p.Match(SFGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewDigitExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(651)
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
			p.SetState(652)
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
			p.SetState(653)
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
			p.SetState(654)
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
			p.SetState(655)
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
	p.SetState(675)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 64, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(673)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 63, p.GetParserRuleContext()) {
			case 1:
				localctx = NewArithmeticOperationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticOperationExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SFGrammarParserRULE_expr)
				p.SetState(658)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(659)

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
					p.SetState(660)

					var _x = p.expr(16)

					localctx.(*ArithmeticOperationExprContext).right = _x
				}

			case 2:
				localctx = NewArithmeticOperationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticOperationExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SFGrammarParserRULE_expr)
				p.SetState(661)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(662)

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
					p.SetState(663)

					var _x = p.expr(15)

					localctx.(*ArithmeticOperationExprContext).right = _x
				}

			case 3:
				localctx = NewComparationOperationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ComparationOperationExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SFGrammarParserRULE_expr)
				p.SetState(664)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(665)

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
					p.SetState(666)

					var _x = p.expr(14)

					localctx.(*ComparationOperationExprContext).right = _x
				}

			case 4:
				localctx = NewRelationalOperationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*RelationalOperationExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SFGrammarParserRULE_expr)
				p.SetState(667)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(668)

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
					p.SetState(669)

					var _x = p.expr(13)

					localctx.(*RelationalOperationExprContext).right = _x
				}

			case 5:
				localctx = NewLogicalOperationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LogicalOperationExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SFGrammarParserRULE_expr)
				p.SetState(670)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(671)

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
					p.SetState(672)

					var _x = p.expr(12)

					localctx.(*LogicalOperationExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(677)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 64, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 64, SFGrammarParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(678)
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
	case 31:
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
