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
		"", "'Int'", "'Float'", "'String'", "'Bool'", "'Char'", "'print'", "'true'",
		"'false'", "", "", "", "'!'", "'('", "')'", "'+'", "'-'", "'*'", "'/'",
		"'%'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "STRING", "BOOL", "CHAR", "PRINT", "TRU", "FAL",
		"DIGIT_PRIMITIVE", "STRING_PRIMITIVE", "ID_PRIMITIVE", "NEGATION_OPERATOR",
		"LPAREN", "RPAREN", "PLUS", "MINUS", "MULTIPLY", "DIVIDE", "MODULO",
		"WHITESPACE", "MULTI_COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"INT", "FLOAT", "STRING", "BOOL", "CHAR", "PRINT", "TRU", "FAL", "DIGIT_PRIMITIVE",
		"STRING_PRIMITIVE", "ID_PRIMITIVE", "NEGATION_OPERATOR", "LPAREN", "RPAREN",
		"PLUS", "MINUS", "MULTIPLY", "DIVIDE", "MODULO", "WHITESPACE", "MULTI_COMMENT",
		"LINE_COMMENT", "ESC_SEQ",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 22, 171, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 8, 4, 8, 93, 8, 8, 11, 8, 12, 8, 94, 1, 8, 1, 8, 4, 8, 99,
		8, 8, 11, 8, 12, 8, 100, 3, 8, 103, 8, 8, 1, 9, 1, 9, 5, 9, 107, 8, 9,
		10, 9, 12, 9, 110, 9, 9, 1, 9, 1, 9, 1, 10, 1, 10, 5, 10, 116, 8, 10, 10,
		10, 12, 10, 119, 9, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 4,
		19, 138, 8, 19, 11, 19, 12, 19, 139, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20,
		1, 20, 5, 20, 148, 8, 20, 10, 20, 12, 20, 151, 9, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 162, 8, 21, 10, 21,
		12, 21, 165, 9, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 149, 0, 23, 1,
		1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 0, 1, 0, 7, 1, 0, 48, 57, 1, 0, 34, 34, 3, 0, 65, 90,
		95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 4, 0, 9, 10, 13,
		13, 32, 32, 92, 92, 2, 0, 10, 10, 13, 13, 7, 0, 32, 33, 35, 35, 43, 43,
		45, 46, 58, 58, 64, 64, 91, 93, 177, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0,
		0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0,
		0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0,
		0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0,
		0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1,
		0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43,
		1, 0, 0, 0, 1, 47, 1, 0, 0, 0, 3, 51, 1, 0, 0, 0, 5, 57, 1, 0, 0, 0, 7,
		64, 1, 0, 0, 0, 9, 69, 1, 0, 0, 0, 11, 74, 1, 0, 0, 0, 13, 80, 1, 0, 0,
		0, 15, 85, 1, 0, 0, 0, 17, 92, 1, 0, 0, 0, 19, 104, 1, 0, 0, 0, 21, 113,
		1, 0, 0, 0, 23, 120, 1, 0, 0, 0, 25, 122, 1, 0, 0, 0, 27, 124, 1, 0, 0,
		0, 29, 126, 1, 0, 0, 0, 31, 128, 1, 0, 0, 0, 33, 130, 1, 0, 0, 0, 35, 132,
		1, 0, 0, 0, 37, 134, 1, 0, 0, 0, 39, 137, 1, 0, 0, 0, 41, 143, 1, 0, 0,
		0, 43, 157, 1, 0, 0, 0, 45, 168, 1, 0, 0, 0, 47, 48, 5, 73, 0, 0, 48, 49,
		5, 110, 0, 0, 49, 50, 5, 116, 0, 0, 50, 2, 1, 0, 0, 0, 51, 52, 5, 70, 0,
		0, 52, 53, 5, 108, 0, 0, 53, 54, 5, 111, 0, 0, 54, 55, 5, 97, 0, 0, 55,
		56, 5, 116, 0, 0, 56, 4, 1, 0, 0, 0, 57, 58, 5, 83, 0, 0, 58, 59, 5, 116,
		0, 0, 59, 60, 5, 114, 0, 0, 60, 61, 5, 105, 0, 0, 61, 62, 5, 110, 0, 0,
		62, 63, 5, 103, 0, 0, 63, 6, 1, 0, 0, 0, 64, 65, 5, 66, 0, 0, 65, 66, 5,
		111, 0, 0, 66, 67, 5, 111, 0, 0, 67, 68, 5, 108, 0, 0, 68, 8, 1, 0, 0,
		0, 69, 70, 5, 67, 0, 0, 70, 71, 5, 104, 0, 0, 71, 72, 5, 97, 0, 0, 72,
		73, 5, 114, 0, 0, 73, 10, 1, 0, 0, 0, 74, 75, 5, 112, 0, 0, 75, 76, 5,
		114, 0, 0, 76, 77, 5, 105, 0, 0, 77, 78, 5, 110, 0, 0, 78, 79, 5, 116,
		0, 0, 79, 12, 1, 0, 0, 0, 80, 81, 5, 116, 0, 0, 81, 82, 5, 114, 0, 0, 82,
		83, 5, 117, 0, 0, 83, 84, 5, 101, 0, 0, 84, 14, 1, 0, 0, 0, 85, 86, 5,
		102, 0, 0, 86, 87, 5, 97, 0, 0, 87, 88, 5, 108, 0, 0, 88, 89, 5, 115, 0,
		0, 89, 90, 5, 101, 0, 0, 90, 16, 1, 0, 0, 0, 91, 93, 7, 0, 0, 0, 92, 91,
		1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0,
		95, 102, 1, 0, 0, 0, 96, 98, 5, 46, 0, 0, 97, 99, 7, 0, 0, 0, 98, 97, 1,
		0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0,
		101, 103, 1, 0, 0, 0, 102, 96, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 18,
		1, 0, 0, 0, 104, 108, 5, 34, 0, 0, 105, 107, 8, 1, 0, 0, 106, 105, 1, 0,
		0, 0, 107, 110, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0,
		109, 111, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 111, 112, 5, 34, 0, 0, 112,
		20, 1, 0, 0, 0, 113, 117, 7, 2, 0, 0, 114, 116, 7, 3, 0, 0, 115, 114, 1,
		0, 0, 0, 116, 119, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118, 1, 0, 0,
		0, 118, 22, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 120, 121, 5, 33, 0, 0, 121,
		24, 1, 0, 0, 0, 122, 123, 5, 40, 0, 0, 123, 26, 1, 0, 0, 0, 124, 125, 5,
		41, 0, 0, 125, 28, 1, 0, 0, 0, 126, 127, 5, 43, 0, 0, 127, 30, 1, 0, 0,
		0, 128, 129, 5, 45, 0, 0, 129, 32, 1, 0, 0, 0, 130, 131, 5, 42, 0, 0, 131,
		34, 1, 0, 0, 0, 132, 133, 5, 47, 0, 0, 133, 36, 1, 0, 0, 0, 134, 135, 5,
		37, 0, 0, 135, 38, 1, 0, 0, 0, 136, 138, 7, 4, 0, 0, 137, 136, 1, 0, 0,
		0, 138, 139, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140,
		141, 1, 0, 0, 0, 141, 142, 6, 19, 0, 0, 142, 40, 1, 0, 0, 0, 143, 144,
		5, 47, 0, 0, 144, 145, 5, 42, 0, 0, 145, 149, 1, 0, 0, 0, 146, 148, 9,
		0, 0, 0, 147, 146, 1, 0, 0, 0, 148, 151, 1, 0, 0, 0, 149, 150, 1, 0, 0,
		0, 149, 147, 1, 0, 0, 0, 150, 152, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 152,
		153, 5, 42, 0, 0, 153, 154, 5, 47, 0, 0, 154, 155, 1, 0, 0, 0, 155, 156,
		6, 20, 0, 0, 156, 42, 1, 0, 0, 0, 157, 158, 5, 47, 0, 0, 158, 159, 5, 47,
		0, 0, 159, 163, 1, 0, 0, 0, 160, 162, 8, 5, 0, 0, 161, 160, 1, 0, 0, 0,
		162, 165, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164,
		166, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 166, 167, 6, 21, 0, 0, 167, 44,
		1, 0, 0, 0, 168, 169, 5, 92, 0, 0, 169, 170, 7, 6, 0, 0, 170, 46, 1, 0,
		0, 0, 9, 0, 94, 100, 102, 108, 117, 139, 149, 163, 1, 6, 0, 0,
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
	SFGrammarLexerDIGIT_PRIMITIVE   = 9
	SFGrammarLexerSTRING_PRIMITIVE  = 10
	SFGrammarLexerID_PRIMITIVE      = 11
	SFGrammarLexerNEGATION_OPERATOR = 12
	SFGrammarLexerLPAREN            = 13
	SFGrammarLexerRPAREN            = 14
	SFGrammarLexerPLUS              = 15
	SFGrammarLexerMINUS             = 16
	SFGrammarLexerMULTIPLY          = 17
	SFGrammarLexerDIVIDE            = 18
	SFGrammarLexerMODULO            = 19
	SFGrammarLexerWHITESPACE        = 20
	SFGrammarLexerMULTI_COMMENT     = 21
	SFGrammarLexerLINE_COMMENT      = 22
)
