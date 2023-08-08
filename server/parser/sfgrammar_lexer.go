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
		"'else'", "'print'", "'true'", "'false'", "'var'", "'let'", "", "",
		"", "'!'", "'('", "')'", "':'", "','", "';'", "'='", "'?'", "'+'", "'-'",
		"'*'", "'/'", "'%'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "STRING", "BOOL", "CHAR", "IF", "ELSE", "PRINT",
		"TRU", "FAL", "DECLARATION_VAR", "DECLARATION_LET", "DIGIT_PRIMITIVE",
		"STRING_PRIMITIVE", "ID_PRIMITIVE", "NEGATION_OPERATOR", "LPAREN", "RPAREN",
		"COLON", "COMMA", "SEMICOLON", "IS_", "QUESTION_MARK", "PLUS", "MINUS",
		"MULTIPLY", "DIVIDE", "MODULO", "WHITESPACE", "MULTI_COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"INT", "FLOAT", "STRING", "BOOL", "CHAR", "IF", "ELSE", "PRINT", "TRU",
		"FAL", "DECLARATION_VAR", "DECLARATION_LET", "DIGIT_PRIMITIVE", "STRING_PRIMITIVE",
		"ID_PRIMITIVE", "NEGATION_OPERATOR", "LPAREN", "RPAREN", "COLON", "COMMA",
		"SEMICOLON", "IS_", "QUESTION_MARK", "PLUS", "MINUS", "MULTIPLY", "DIVIDE",
		"MODULO", "WHITESPACE", "MULTI_COMMENT", "LINE_COMMENT", "ESC_SEQ",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 31, 220, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 4, 12, 132, 8,
		12, 11, 12, 12, 12, 133, 1, 12, 1, 12, 4, 12, 138, 8, 12, 11, 12, 12, 12,
		139, 3, 12, 142, 8, 12, 1, 13, 1, 13, 5, 13, 146, 8, 13, 10, 13, 12, 13,
		149, 9, 13, 1, 13, 1, 13, 1, 14, 1, 14, 5, 14, 155, 8, 14, 10, 14, 12,
		14, 158, 9, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 4, 28, 187,
		8, 28, 11, 28, 12, 28, 188, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 5,
		29, 197, 8, 29, 10, 29, 12, 29, 200, 9, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 5, 30, 211, 8, 30, 10, 30, 12, 30, 214,
		9, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 198, 0, 32, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43,
		22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61,
		31, 63, 0, 1, 0, 7, 1, 0, 48, 57, 1, 0, 34, 34, 3, 0, 65, 90, 95, 95, 97,
		122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 4, 0, 9, 10, 13, 13, 32, 32,
		92, 92, 2, 0, 10, 10, 13, 13, 7, 0, 32, 33, 35, 35, 43, 43, 45, 46, 58,
		58, 64, 64, 91, 93, 226, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21,
		1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0,
		29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0,
		0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0,
		0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0,
		0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1,
		0, 0, 0, 0, 61, 1, 0, 0, 0, 1, 65, 1, 0, 0, 0, 3, 69, 1, 0, 0, 0, 5, 75,
		1, 0, 0, 0, 7, 82, 1, 0, 0, 0, 9, 87, 1, 0, 0, 0, 11, 97, 1, 0, 0, 0, 13,
		100, 1, 0, 0, 0, 15, 105, 1, 0, 0, 0, 17, 111, 1, 0, 0, 0, 19, 116, 1,
		0, 0, 0, 21, 122, 1, 0, 0, 0, 23, 126, 1, 0, 0, 0, 25, 131, 1, 0, 0, 0,
		27, 143, 1, 0, 0, 0, 29, 152, 1, 0, 0, 0, 31, 159, 1, 0, 0, 0, 33, 161,
		1, 0, 0, 0, 35, 163, 1, 0, 0, 0, 37, 165, 1, 0, 0, 0, 39, 167, 1, 0, 0,
		0, 41, 169, 1, 0, 0, 0, 43, 171, 1, 0, 0, 0, 45, 173, 1, 0, 0, 0, 47, 175,
		1, 0, 0, 0, 49, 177, 1, 0, 0, 0, 51, 179, 1, 0, 0, 0, 53, 181, 1, 0, 0,
		0, 55, 183, 1, 0, 0, 0, 57, 186, 1, 0, 0, 0, 59, 192, 1, 0, 0, 0, 61, 206,
		1, 0, 0, 0, 63, 217, 1, 0, 0, 0, 65, 66, 5, 73, 0, 0, 66, 67, 5, 110, 0,
		0, 67, 68, 5, 116, 0, 0, 68, 2, 1, 0, 0, 0, 69, 70, 5, 70, 0, 0, 70, 71,
		5, 108, 0, 0, 71, 72, 5, 111, 0, 0, 72, 73, 5, 97, 0, 0, 73, 74, 5, 116,
		0, 0, 74, 4, 1, 0, 0, 0, 75, 76, 5, 83, 0, 0, 76, 77, 5, 116, 0, 0, 77,
		78, 5, 114, 0, 0, 78, 79, 5, 105, 0, 0, 79, 80, 5, 110, 0, 0, 80, 81, 5,
		103, 0, 0, 81, 6, 1, 0, 0, 0, 82, 83, 5, 66, 0, 0, 83, 84, 5, 111, 0, 0,
		84, 85, 5, 111, 0, 0, 85, 86, 5, 108, 0, 0, 86, 8, 1, 0, 0, 0, 87, 88,
		5, 67, 0, 0, 88, 89, 5, 104, 0, 0, 89, 90, 5, 97, 0, 0, 90, 91, 5, 114,
		0, 0, 91, 92, 5, 97, 0, 0, 92, 93, 5, 99, 0, 0, 93, 94, 5, 116, 0, 0, 94,
		95, 5, 101, 0, 0, 95, 96, 5, 114, 0, 0, 96, 10, 1, 0, 0, 0, 97, 98, 5,
		105, 0, 0, 98, 99, 5, 102, 0, 0, 99, 12, 1, 0, 0, 0, 100, 101, 5, 101,
		0, 0, 101, 102, 5, 108, 0, 0, 102, 103, 5, 115, 0, 0, 103, 104, 5, 101,
		0, 0, 104, 14, 1, 0, 0, 0, 105, 106, 5, 112, 0, 0, 106, 107, 5, 114, 0,
		0, 107, 108, 5, 105, 0, 0, 108, 109, 5, 110, 0, 0, 109, 110, 5, 116, 0,
		0, 110, 16, 1, 0, 0, 0, 111, 112, 5, 116, 0, 0, 112, 113, 5, 114, 0, 0,
		113, 114, 5, 117, 0, 0, 114, 115, 5, 101, 0, 0, 115, 18, 1, 0, 0, 0, 116,
		117, 5, 102, 0, 0, 117, 118, 5, 97, 0, 0, 118, 119, 5, 108, 0, 0, 119,
		120, 5, 115, 0, 0, 120, 121, 5, 101, 0, 0, 121, 20, 1, 0, 0, 0, 122, 123,
		5, 118, 0, 0, 123, 124, 5, 97, 0, 0, 124, 125, 5, 114, 0, 0, 125, 22, 1,
		0, 0, 0, 126, 127, 5, 108, 0, 0, 127, 128, 5, 101, 0, 0, 128, 129, 5, 116,
		0, 0, 129, 24, 1, 0, 0, 0, 130, 132, 7, 0, 0, 0, 131, 130, 1, 0, 0, 0,
		132, 133, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134,
		141, 1, 0, 0, 0, 135, 137, 5, 46, 0, 0, 136, 138, 7, 0, 0, 0, 137, 136,
		1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0,
		0, 0, 140, 142, 1, 0, 0, 0, 141, 135, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0,
		142, 26, 1, 0, 0, 0, 143, 147, 5, 34, 0, 0, 144, 146, 8, 1, 0, 0, 145,
		144, 1, 0, 0, 0, 146, 149, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 147, 148,
		1, 0, 0, 0, 148, 150, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 150, 151, 5, 34,
		0, 0, 151, 28, 1, 0, 0, 0, 152, 156, 7, 2, 0, 0, 153, 155, 7, 3, 0, 0,
		154, 153, 1, 0, 0, 0, 155, 158, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 156,
		157, 1, 0, 0, 0, 157, 30, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 159, 160, 5,
		33, 0, 0, 160, 32, 1, 0, 0, 0, 161, 162, 5, 40, 0, 0, 162, 34, 1, 0, 0,
		0, 163, 164, 5, 41, 0, 0, 164, 36, 1, 0, 0, 0, 165, 166, 5, 58, 0, 0, 166,
		38, 1, 0, 0, 0, 167, 168, 5, 44, 0, 0, 168, 40, 1, 0, 0, 0, 169, 170, 5,
		59, 0, 0, 170, 42, 1, 0, 0, 0, 171, 172, 5, 61, 0, 0, 172, 44, 1, 0, 0,
		0, 173, 174, 5, 63, 0, 0, 174, 46, 1, 0, 0, 0, 175, 176, 5, 43, 0, 0, 176,
		48, 1, 0, 0, 0, 177, 178, 5, 45, 0, 0, 178, 50, 1, 0, 0, 0, 179, 180, 5,
		42, 0, 0, 180, 52, 1, 0, 0, 0, 181, 182, 5, 47, 0, 0, 182, 54, 1, 0, 0,
		0, 183, 184, 5, 37, 0, 0, 184, 56, 1, 0, 0, 0, 185, 187, 7, 4, 0, 0, 186,
		185, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 188, 189,
		1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 191, 6, 28, 0, 0, 191, 58, 1, 0,
		0, 0, 192, 193, 5, 47, 0, 0, 193, 194, 5, 42, 0, 0, 194, 198, 1, 0, 0,
		0, 195, 197, 9, 0, 0, 0, 196, 195, 1, 0, 0, 0, 197, 200, 1, 0, 0, 0, 198,
		199, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 199, 201, 1, 0, 0, 0, 200, 198,
		1, 0, 0, 0, 201, 202, 5, 42, 0, 0, 202, 203, 5, 47, 0, 0, 203, 204, 1,
		0, 0, 0, 204, 205, 6, 29, 0, 0, 205, 60, 1, 0, 0, 0, 206, 207, 5, 47, 0,
		0, 207, 208, 5, 47, 0, 0, 208, 212, 1, 0, 0, 0, 209, 211, 8, 5, 0, 0, 210,
		209, 1, 0, 0, 0, 211, 214, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 212, 213,
		1, 0, 0, 0, 213, 215, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 215, 216, 6, 30,
		0, 0, 216, 62, 1, 0, 0, 0, 217, 218, 5, 92, 0, 0, 218, 219, 7, 6, 0, 0,
		219, 64, 1, 0, 0, 0, 9, 0, 133, 139, 141, 147, 156, 188, 198, 212, 1, 6,
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
	SFGrammarLexerDECLARATION_VAR   = 11
	SFGrammarLexerDECLARATION_LET   = 12
	SFGrammarLexerDIGIT_PRIMITIVE   = 13
	SFGrammarLexerSTRING_PRIMITIVE  = 14
	SFGrammarLexerID_PRIMITIVE      = 15
	SFGrammarLexerNEGATION_OPERATOR = 16
	SFGrammarLexerLPAREN            = 17
	SFGrammarLexerRPAREN            = 18
	SFGrammarLexerCOLON             = 19
	SFGrammarLexerCOMMA             = 20
	SFGrammarLexerSEMICOLON         = 21
	SFGrammarLexerIS_               = 22
	SFGrammarLexerQUESTION_MARK     = 23
	SFGrammarLexerPLUS              = 24
	SFGrammarLexerMINUS             = 25
	SFGrammarLexerMULTIPLY          = 26
	SFGrammarLexerDIVIDE            = 27
	SFGrammarLexerMODULO            = 28
	SFGrammarLexerWHITESPACE        = 29
	SFGrammarLexerMULTI_COMMENT     = 30
	SFGrammarLexerLINE_COMMENT      = 31
)
