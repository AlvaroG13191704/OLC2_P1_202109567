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
		"", "'Int'", "'Float'", "'String'", "'Bool'", "'Character'", "'print'",
		"'true'", "'false'", "'var'", "'let'", "", "", "", "'!'", "'('", "')'",
		"':'", "','", "';'", "'='", "'?'", "'+'", "'-'", "'*'", "'/'", "'%'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "STRING", "BOOL", "CHAR", "PRINT", "TRU", "FAL",
		"DECLARATION_1", "DECLARATION_2", "DIGIT_PRIMITIVE", "STRING_PRIMITIVE",
		"ID_PRIMITIVE", "NEGATION_OPERATOR", "LPAREN", "RPAREN", "COLON", "COMMA",
		"SEMICOLON", "IS_", "QUESTION_MARK", "PLUS", "MINUS", "MULTIPLY", "DIVIDE",
		"MODULO", "WHITESPACE", "MULTI_COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"INT", "FLOAT", "STRING", "BOOL", "CHAR", "PRINT", "TRU", "FAL", "DECLARATION_1",
		"DECLARATION_2", "DIGIT_PRIMITIVE", "STRING_PRIMITIVE", "ID_PRIMITIVE",
		"NEGATION_OPERATOR", "LPAREN", "RPAREN", "COLON", "COMMA", "SEMICOLON",
		"IS_", "QUESTION_MARK", "PLUS", "MINUS", "MULTIPLY", "DIVIDE", "MODULO",
		"WHITESPACE", "MULTI_COMMENT", "LINE_COMMENT", "ESC_SEQ",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 29, 208, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 4, 10, 120, 8, 10, 11, 10,
		12, 10, 121, 1, 10, 1, 10, 4, 10, 126, 8, 10, 11, 10, 12, 10, 127, 3, 10,
		130, 8, 10, 1, 11, 1, 11, 5, 11, 134, 8, 11, 10, 11, 12, 11, 137, 9, 11,
		1, 11, 1, 11, 1, 12, 1, 12, 5, 12, 143, 8, 12, 10, 12, 12, 12, 146, 9,
		12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 4, 26, 175, 8, 26, 11, 26,
		12, 26, 176, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 185, 8, 27,
		10, 27, 12, 27, 188, 9, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1,
		28, 1, 28, 1, 28, 5, 28, 199, 8, 28, 10, 28, 12, 28, 202, 9, 28, 1, 28,
		1, 28, 1, 29, 1, 29, 1, 29, 1, 186, 0, 30, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5,
		11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29,
		15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47,
		24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 0, 1, 0, 7, 1, 0, 48, 57,
		1, 0, 34, 34, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95,
		95, 97, 122, 4, 0, 9, 10, 13, 13, 32, 32, 92, 92, 2, 0, 10, 10, 13, 13,
		7, 0, 32, 33, 35, 35, 43, 43, 45, 46, 58, 58, 64, 64, 91, 93, 214, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0,
		0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0,
		0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0,
		0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1,
		0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55,
		1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 1, 61, 1, 0, 0, 0, 3, 65, 1, 0, 0, 0, 5,
		71, 1, 0, 0, 0, 7, 78, 1, 0, 0, 0, 9, 83, 1, 0, 0, 0, 11, 93, 1, 0, 0,
		0, 13, 99, 1, 0, 0, 0, 15, 104, 1, 0, 0, 0, 17, 110, 1, 0, 0, 0, 19, 114,
		1, 0, 0, 0, 21, 119, 1, 0, 0, 0, 23, 131, 1, 0, 0, 0, 25, 140, 1, 0, 0,
		0, 27, 147, 1, 0, 0, 0, 29, 149, 1, 0, 0, 0, 31, 151, 1, 0, 0, 0, 33, 153,
		1, 0, 0, 0, 35, 155, 1, 0, 0, 0, 37, 157, 1, 0, 0, 0, 39, 159, 1, 0, 0,
		0, 41, 161, 1, 0, 0, 0, 43, 163, 1, 0, 0, 0, 45, 165, 1, 0, 0, 0, 47, 167,
		1, 0, 0, 0, 49, 169, 1, 0, 0, 0, 51, 171, 1, 0, 0, 0, 53, 174, 1, 0, 0,
		0, 55, 180, 1, 0, 0, 0, 57, 194, 1, 0, 0, 0, 59, 205, 1, 0, 0, 0, 61, 62,
		5, 73, 0, 0, 62, 63, 5, 110, 0, 0, 63, 64, 5, 116, 0, 0, 64, 2, 1, 0, 0,
		0, 65, 66, 5, 70, 0, 0, 66, 67, 5, 108, 0, 0, 67, 68, 5, 111, 0, 0, 68,
		69, 5, 97, 0, 0, 69, 70, 5, 116, 0, 0, 70, 4, 1, 0, 0, 0, 71, 72, 5, 83,
		0, 0, 72, 73, 5, 116, 0, 0, 73, 74, 5, 114, 0, 0, 74, 75, 5, 105, 0, 0,
		75, 76, 5, 110, 0, 0, 76, 77, 5, 103, 0, 0, 77, 6, 1, 0, 0, 0, 78, 79,
		5, 66, 0, 0, 79, 80, 5, 111, 0, 0, 80, 81, 5, 111, 0, 0, 81, 82, 5, 108,
		0, 0, 82, 8, 1, 0, 0, 0, 83, 84, 5, 67, 0, 0, 84, 85, 5, 104, 0, 0, 85,
		86, 5, 97, 0, 0, 86, 87, 5, 114, 0, 0, 87, 88, 5, 97, 0, 0, 88, 89, 5,
		99, 0, 0, 89, 90, 5, 116, 0, 0, 90, 91, 5, 101, 0, 0, 91, 92, 5, 114, 0,
		0, 92, 10, 1, 0, 0, 0, 93, 94, 5, 112, 0, 0, 94, 95, 5, 114, 0, 0, 95,
		96, 5, 105, 0, 0, 96, 97, 5, 110, 0, 0, 97, 98, 5, 116, 0, 0, 98, 12, 1,
		0, 0, 0, 99, 100, 5, 116, 0, 0, 100, 101, 5, 114, 0, 0, 101, 102, 5, 117,
		0, 0, 102, 103, 5, 101, 0, 0, 103, 14, 1, 0, 0, 0, 104, 105, 5, 102, 0,
		0, 105, 106, 5, 97, 0, 0, 106, 107, 5, 108, 0, 0, 107, 108, 5, 115, 0,
		0, 108, 109, 5, 101, 0, 0, 109, 16, 1, 0, 0, 0, 110, 111, 5, 118, 0, 0,
		111, 112, 5, 97, 0, 0, 112, 113, 5, 114, 0, 0, 113, 18, 1, 0, 0, 0, 114,
		115, 5, 108, 0, 0, 115, 116, 5, 101, 0, 0, 116, 117, 5, 116, 0, 0, 117,
		20, 1, 0, 0, 0, 118, 120, 7, 0, 0, 0, 119, 118, 1, 0, 0, 0, 120, 121, 1,
		0, 0, 0, 121, 119, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 129, 1, 0, 0,
		0, 123, 125, 5, 46, 0, 0, 124, 126, 7, 0, 0, 0, 125, 124, 1, 0, 0, 0, 126,
		127, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 130,
		1, 0, 0, 0, 129, 123, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 22, 1, 0,
		0, 0, 131, 135, 5, 34, 0, 0, 132, 134, 8, 1, 0, 0, 133, 132, 1, 0, 0, 0,
		134, 137, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136,
		138, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 138, 139, 5, 34, 0, 0, 139, 24,
		1, 0, 0, 0, 140, 144, 7, 2, 0, 0, 141, 143, 7, 3, 0, 0, 142, 141, 1, 0,
		0, 0, 143, 146, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0,
		145, 26, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 147, 148, 5, 33, 0, 0, 148,
		28, 1, 0, 0, 0, 149, 150, 5, 40, 0, 0, 150, 30, 1, 0, 0, 0, 151, 152, 5,
		41, 0, 0, 152, 32, 1, 0, 0, 0, 153, 154, 5, 58, 0, 0, 154, 34, 1, 0, 0,
		0, 155, 156, 5, 44, 0, 0, 156, 36, 1, 0, 0, 0, 157, 158, 5, 59, 0, 0, 158,
		38, 1, 0, 0, 0, 159, 160, 5, 61, 0, 0, 160, 40, 1, 0, 0, 0, 161, 162, 5,
		63, 0, 0, 162, 42, 1, 0, 0, 0, 163, 164, 5, 43, 0, 0, 164, 44, 1, 0, 0,
		0, 165, 166, 5, 45, 0, 0, 166, 46, 1, 0, 0, 0, 167, 168, 5, 42, 0, 0, 168,
		48, 1, 0, 0, 0, 169, 170, 5, 47, 0, 0, 170, 50, 1, 0, 0, 0, 171, 172, 5,
		37, 0, 0, 172, 52, 1, 0, 0, 0, 173, 175, 7, 4, 0, 0, 174, 173, 1, 0, 0,
		0, 175, 176, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177,
		178, 1, 0, 0, 0, 178, 179, 6, 26, 0, 0, 179, 54, 1, 0, 0, 0, 180, 181,
		5, 47, 0, 0, 181, 182, 5, 42, 0, 0, 182, 186, 1, 0, 0, 0, 183, 185, 9,
		0, 0, 0, 184, 183, 1, 0, 0, 0, 185, 188, 1, 0, 0, 0, 186, 187, 1, 0, 0,
		0, 186, 184, 1, 0, 0, 0, 187, 189, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 189,
		190, 5, 42, 0, 0, 190, 191, 5, 47, 0, 0, 191, 192, 1, 0, 0, 0, 192, 193,
		6, 27, 0, 0, 193, 56, 1, 0, 0, 0, 194, 195, 5, 47, 0, 0, 195, 196, 5, 47,
		0, 0, 196, 200, 1, 0, 0, 0, 197, 199, 8, 5, 0, 0, 198, 197, 1, 0, 0, 0,
		199, 202, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201,
		203, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0, 203, 204, 6, 28, 0, 0, 204, 58,
		1, 0, 0, 0, 205, 206, 5, 92, 0, 0, 206, 207, 7, 6, 0, 0, 207, 60, 1, 0,
		0, 0, 9, 0, 121, 127, 129, 135, 144, 176, 186, 200, 1, 6, 0, 0,
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
	SFGrammarLexerPRINT             = 6
	SFGrammarLexerTRU               = 7
	SFGrammarLexerFAL               = 8
	SFGrammarLexerDECLARATION_1     = 9
	SFGrammarLexerDECLARATION_2     = 10
	SFGrammarLexerDIGIT_PRIMITIVE   = 11
	SFGrammarLexerSTRING_PRIMITIVE  = 12
	SFGrammarLexerID_PRIMITIVE      = 13
	SFGrammarLexerNEGATION_OPERATOR = 14
	SFGrammarLexerLPAREN            = 15
	SFGrammarLexerRPAREN            = 16
	SFGrammarLexerCOLON             = 17
	SFGrammarLexerCOMMA             = 18
	SFGrammarLexerSEMICOLON         = 19
	SFGrammarLexerIS_               = 20
	SFGrammarLexerQUESTION_MARK     = 21
	SFGrammarLexerPLUS              = 22
	SFGrammarLexerMINUS             = 23
	SFGrammarLexerMULTIPLY          = 24
	SFGrammarLexerDIVIDE            = 25
	SFGrammarLexerMODULO            = 26
	SFGrammarLexerWHITESPACE        = 27
	SFGrammarLexerMULTI_COMMENT     = 28
	SFGrammarLexerLINE_COMMENT      = 29
)
