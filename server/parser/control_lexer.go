// Code generated from Control.g4 by ANTLR 4.13.0. DO NOT EDIT.

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

type ControlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ControlLexerLexerStaticData struct {
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

func controllexerLexerInit() {
	staticData := &ControlLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'='", "'println'", "'('", "')'", "'print'", "'if'", "'{'", "'}'",
		"'while'", "'for'", "';'", "'!'", "'*'", "'/'", "'+'", "'-'", "'>='",
		"'>'", "'<='", "'<'", "'=='", "'!='", "'&&'", "'||'", "'true'", "'false'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "INT", "ID", "STRING", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "INT", "ID", "STRING", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 30, 173, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1,
		20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 26, 4, 26, 145, 8, 26, 11, 26, 12, 26, 146, 1, 27, 1, 27,
		5, 27, 151, 8, 27, 10, 27, 12, 27, 154, 9, 27, 1, 28, 1, 28, 1, 28, 1,
		28, 5, 28, 160, 8, 28, 10, 28, 12, 28, 163, 9, 28, 1, 28, 1, 28, 1, 29,
		4, 29, 168, 8, 29, 11, 29, 12, 29, 169, 1, 29, 1, 29, 0, 0, 30, 1, 1, 3,
		2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12,
		25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21,
		43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30,
		1, 0, 5, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65,
		90, 95, 95, 97, 122, 3, 0, 10, 10, 13, 13, 34, 34, 2, 0, 9, 10, 32, 32,
		177, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0,
		0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1,
		0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23,
		1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0,
		31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0,
		0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0,
		0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0,
		0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 1, 61, 1,
		0, 0, 0, 3, 63, 1, 0, 0, 0, 5, 71, 1, 0, 0, 0, 7, 73, 1, 0, 0, 0, 9, 75,
		1, 0, 0, 0, 11, 81, 1, 0, 0, 0, 13, 84, 1, 0, 0, 0, 15, 86, 1, 0, 0, 0,
		17, 88, 1, 0, 0, 0, 19, 94, 1, 0, 0, 0, 21, 98, 1, 0, 0, 0, 23, 100, 1,
		0, 0, 0, 25, 102, 1, 0, 0, 0, 27, 104, 1, 0, 0, 0, 29, 106, 1, 0, 0, 0,
		31, 108, 1, 0, 0, 0, 33, 110, 1, 0, 0, 0, 35, 113, 1, 0, 0, 0, 37, 115,
		1, 0, 0, 0, 39, 118, 1, 0, 0, 0, 41, 120, 1, 0, 0, 0, 43, 123, 1, 0, 0,
		0, 45, 126, 1, 0, 0, 0, 47, 129, 1, 0, 0, 0, 49, 132, 1, 0, 0, 0, 51, 137,
		1, 0, 0, 0, 53, 144, 1, 0, 0, 0, 55, 148, 1, 0, 0, 0, 57, 155, 1, 0, 0,
		0, 59, 167, 1, 0, 0, 0, 61, 62, 5, 61, 0, 0, 62, 2, 1, 0, 0, 0, 63, 64,
		5, 112, 0, 0, 64, 65, 5, 114, 0, 0, 65, 66, 5, 105, 0, 0, 66, 67, 5, 110,
		0, 0, 67, 68, 5, 116, 0, 0, 68, 69, 5, 108, 0, 0, 69, 70, 5, 110, 0, 0,
		70, 4, 1, 0, 0, 0, 71, 72, 5, 40, 0, 0, 72, 6, 1, 0, 0, 0, 73, 74, 5, 41,
		0, 0, 74, 8, 1, 0, 0, 0, 75, 76, 5, 112, 0, 0, 76, 77, 5, 114, 0, 0, 77,
		78, 5, 105, 0, 0, 78, 79, 5, 110, 0, 0, 79, 80, 5, 116, 0, 0, 80, 10, 1,
		0, 0, 0, 81, 82, 5, 105, 0, 0, 82, 83, 5, 102, 0, 0, 83, 12, 1, 0, 0, 0,
		84, 85, 5, 123, 0, 0, 85, 14, 1, 0, 0, 0, 86, 87, 5, 125, 0, 0, 87, 16,
		1, 0, 0, 0, 88, 89, 5, 119, 0, 0, 89, 90, 5, 104, 0, 0, 90, 91, 5, 105,
		0, 0, 91, 92, 5, 108, 0, 0, 92, 93, 5, 101, 0, 0, 93, 18, 1, 0, 0, 0, 94,
		95, 5, 102, 0, 0, 95, 96, 5, 111, 0, 0, 96, 97, 5, 114, 0, 0, 97, 20, 1,
		0, 0, 0, 98, 99, 5, 59, 0, 0, 99, 22, 1, 0, 0, 0, 100, 101, 5, 33, 0, 0,
		101, 24, 1, 0, 0, 0, 102, 103, 5, 42, 0, 0, 103, 26, 1, 0, 0, 0, 104, 105,
		5, 47, 0, 0, 105, 28, 1, 0, 0, 0, 106, 107, 5, 43, 0, 0, 107, 30, 1, 0,
		0, 0, 108, 109, 5, 45, 0, 0, 109, 32, 1, 0, 0, 0, 110, 111, 5, 62, 0, 0,
		111, 112, 5, 61, 0, 0, 112, 34, 1, 0, 0, 0, 113, 114, 5, 62, 0, 0, 114,
		36, 1, 0, 0, 0, 115, 116, 5, 60, 0, 0, 116, 117, 5, 61, 0, 0, 117, 38,
		1, 0, 0, 0, 118, 119, 5, 60, 0, 0, 119, 40, 1, 0, 0, 0, 120, 121, 5, 61,
		0, 0, 121, 122, 5, 61, 0, 0, 122, 42, 1, 0, 0, 0, 123, 124, 5, 33, 0, 0,
		124, 125, 5, 61, 0, 0, 125, 44, 1, 0, 0, 0, 126, 127, 5, 38, 0, 0, 127,
		128, 5, 38, 0, 0, 128, 46, 1, 0, 0, 0, 129, 130, 5, 124, 0, 0, 130, 131,
		5, 124, 0, 0, 131, 48, 1, 0, 0, 0, 132, 133, 5, 116, 0, 0, 133, 134, 5,
		114, 0, 0, 134, 135, 5, 117, 0, 0, 135, 136, 5, 101, 0, 0, 136, 50, 1,
		0, 0, 0, 137, 138, 5, 102, 0, 0, 138, 139, 5, 97, 0, 0, 139, 140, 5, 108,
		0, 0, 140, 141, 5, 115, 0, 0, 141, 142, 5, 101, 0, 0, 142, 52, 1, 0, 0,
		0, 143, 145, 7, 0, 0, 0, 144, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146,
		144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 54, 1, 0, 0, 0, 148, 152, 7,
		1, 0, 0, 149, 151, 7, 2, 0, 0, 150, 149, 1, 0, 0, 0, 151, 154, 1, 0, 0,
		0, 152, 150, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 56, 1, 0, 0, 0, 154,
		152, 1, 0, 0, 0, 155, 161, 5, 34, 0, 0, 156, 160, 8, 3, 0, 0, 157, 158,
		5, 34, 0, 0, 158, 160, 5, 34, 0, 0, 159, 156, 1, 0, 0, 0, 159, 157, 1,
		0, 0, 0, 160, 163, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 161, 162, 1, 0, 0,
		0, 162, 164, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 164, 165, 5, 34, 0, 0, 165,
		58, 1, 0, 0, 0, 166, 168, 7, 4, 0, 0, 167, 166, 1, 0, 0, 0, 168, 169, 1,
		0, 0, 0, 169, 167, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 171, 1, 0, 0,
		0, 171, 172, 6, 29, 0, 0, 172, 60, 1, 0, 0, 0, 6, 0, 146, 152, 159, 161,
		169, 1, 6, 0, 0,
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

// ControlLexerInit initializes any static state used to implement ControlLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewControlLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ControlLexerInit() {
	staticData := &ControlLexerLexerStaticData
	staticData.once.Do(controllexerLexerInit)
}

// NewControlLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewControlLexer(input antlr.CharStream) *ControlLexer {
	ControlLexerInit()
	l := new(ControlLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ControlLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Control.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ControlLexer tokens.
const (
	ControlLexerT__0   = 1
	ControlLexerT__1   = 2
	ControlLexerT__2   = 3
	ControlLexerT__3   = 4
	ControlLexerT__4   = 5
	ControlLexerT__5   = 6
	ControlLexerT__6   = 7
	ControlLexerT__7   = 8
	ControlLexerT__8   = 9
	ControlLexerT__9   = 10
	ControlLexerT__10  = 11
	ControlLexerT__11  = 12
	ControlLexerT__12  = 13
	ControlLexerT__13  = 14
	ControlLexerT__14  = 15
	ControlLexerT__15  = 16
	ControlLexerT__16  = 17
	ControlLexerT__17  = 18
	ControlLexerT__18  = 19
	ControlLexerT__19  = 20
	ControlLexerT__20  = 21
	ControlLexerT__21  = 22
	ControlLexerT__22  = 23
	ControlLexerT__23  = 24
	ControlLexerT__24  = 25
	ControlLexerT__25  = 26
	ControlLexerINT    = 27
	ControlLexerID     = 28
	ControlLexerSTRING = 29
	ControlLexerWS     = 30
)