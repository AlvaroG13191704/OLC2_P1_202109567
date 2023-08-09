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
		"'else'", "'print'", "'true'", "'false'", "'nil'", "'var'", "'let'",
		"", "", "", "'!'", "'('", "')'", "':'", "','", "';'", "'='", "'+='",
		"'-='", "'?'", "'+'", "'-'", "'*'", "'/'", "'%'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "STRING", "BOOL", "CHAR", "IF", "ELSE", "PRINT",
		"TRU", "FAL", "NIL", "DECLARATION_VAR", "DECLARATION_LET", "DIGIT_PRIMITIVE",
		"STRING_PRIMITIVE", "ID_PRIMITIVE", "NEGATION_OPERATOR", "LPAREN", "RPAREN",
		"COLON", "COMMA", "SEMICOLON", "IS_", "PLUS_IS", "MINUS_IS", "QUESTION_MARK",
		"PLUS", "MINUS", "MULTIPLY", "DIVIDE", "MODULO", "WHITESPACE", "MULTI_COMMENT",
		"LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"start", "block", "stmts", "printstmt", "declaration", "type_declaration",
		"assignment", "exprList", "expr", "type",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 34, 119, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 1,
		0, 1, 0, 1, 1, 5, 1, 25, 8, 1, 10, 1, 12, 1, 28, 9, 1, 1, 2, 1, 2, 3, 2,
		32, 8, 2, 1, 2, 1, 2, 3, 2, 36, 8, 2, 1, 2, 1, 2, 3, 2, 40, 8, 2, 3, 2,
		42, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 3, 4, 67, 8, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 3, 6, 80, 8, 6, 1, 7, 1, 7, 1, 7, 5, 7, 85, 8, 7, 10, 7, 12,
		7, 88, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 104, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 5, 8, 112, 8, 8, 10, 8, 12, 8, 115, 9, 8, 1, 9, 1, 9, 1, 9, 0, 1,
		16, 10, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 0, 5, 1, 0, 12, 13, 1, 0, 9,
		10, 1, 0, 29, 31, 1, 0, 27, 28, 1, 0, 1, 5, 128, 0, 20, 1, 0, 0, 0, 2,
		26, 1, 0, 0, 0, 4, 41, 1, 0, 0, 0, 6, 43, 1, 0, 0, 0, 8, 66, 1, 0, 0, 0,
		10, 68, 1, 0, 0, 0, 12, 79, 1, 0, 0, 0, 14, 81, 1, 0, 0, 0, 16, 103, 1,
		0, 0, 0, 18, 116, 1, 0, 0, 0, 20, 21, 3, 2, 1, 0, 21, 22, 5, 0, 0, 1, 22,
		1, 1, 0, 0, 0, 23, 25, 3, 4, 2, 0, 24, 23, 1, 0, 0, 0, 25, 28, 1, 0, 0,
		0, 26, 24, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 3, 1, 0, 0, 0, 28, 26, 1,
		0, 0, 0, 29, 31, 3, 6, 3, 0, 30, 32, 5, 22, 0, 0, 31, 30, 1, 0, 0, 0, 31,
		32, 1, 0, 0, 0, 32, 42, 1, 0, 0, 0, 33, 35, 3, 8, 4, 0, 34, 36, 5, 22,
		0, 0, 35, 34, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 42, 1, 0, 0, 0, 37, 39,
		3, 12, 6, 0, 38, 40, 5, 22, 0, 0, 39, 38, 1, 0, 0, 0, 39, 40, 1, 0, 0,
		0, 40, 42, 1, 0, 0, 0, 41, 29, 1, 0, 0, 0, 41, 33, 1, 0, 0, 0, 41, 37,
		1, 0, 0, 0, 42, 5, 1, 0, 0, 0, 43, 44, 5, 8, 0, 0, 44, 45, 5, 18, 0, 0,
		45, 46, 3, 14, 7, 0, 46, 47, 5, 19, 0, 0, 47, 7, 1, 0, 0, 0, 48, 49, 3,
		10, 5, 0, 49, 50, 5, 16, 0, 0, 50, 51, 5, 20, 0, 0, 51, 52, 3, 18, 9, 0,
		52, 53, 5, 23, 0, 0, 53, 54, 3, 16, 8, 0, 54, 67, 1, 0, 0, 0, 55, 56, 3,
		10, 5, 0, 56, 57, 5, 16, 0, 0, 57, 58, 5, 20, 0, 0, 58, 59, 3, 18, 9, 0,
		59, 60, 5, 26, 0, 0, 60, 67, 1, 0, 0, 0, 61, 62, 3, 10, 5, 0, 62, 63, 5,
		16, 0, 0, 63, 64, 5, 23, 0, 0, 64, 65, 3, 16, 8, 0, 65, 67, 1, 0, 0, 0,
		66, 48, 1, 0, 0, 0, 66, 55, 1, 0, 0, 0, 66, 61, 1, 0, 0, 0, 67, 9, 1, 0,
		0, 0, 68, 69, 7, 0, 0, 0, 69, 11, 1, 0, 0, 0, 70, 71, 5, 16, 0, 0, 71,
		72, 5, 23, 0, 0, 72, 80, 3, 16, 8, 0, 73, 74, 5, 16, 0, 0, 74, 75, 5, 24,
		0, 0, 75, 80, 3, 16, 8, 0, 76, 77, 5, 16, 0, 0, 77, 78, 5, 25, 0, 0, 78,
		80, 3, 16, 8, 0, 79, 70, 1, 0, 0, 0, 79, 73, 1, 0, 0, 0, 79, 76, 1, 0,
		0, 0, 80, 13, 1, 0, 0, 0, 81, 86, 3, 16, 8, 0, 82, 83, 5, 21, 0, 0, 83,
		85, 3, 16, 8, 0, 84, 82, 1, 0, 0, 0, 85, 88, 1, 0, 0, 0, 86, 84, 1, 0,
		0, 0, 86, 87, 1, 0, 0, 0, 87, 15, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 89, 90,
		6, 8, -1, 0, 90, 91, 5, 17, 0, 0, 91, 104, 3, 16, 8, 10, 92, 93, 5, 28,
		0, 0, 93, 104, 3, 16, 8, 9, 94, 95, 5, 18, 0, 0, 95, 96, 3, 16, 8, 0, 96,
		97, 5, 19, 0, 0, 97, 104, 1, 0, 0, 0, 98, 104, 5, 14, 0, 0, 99, 104, 5,
		15, 0, 0, 100, 104, 5, 16, 0, 0, 101, 104, 5, 11, 0, 0, 102, 104, 7, 1,
		0, 0, 103, 89, 1, 0, 0, 0, 103, 92, 1, 0, 0, 0, 103, 94, 1, 0, 0, 0, 103,
		98, 1, 0, 0, 0, 103, 99, 1, 0, 0, 0, 103, 100, 1, 0, 0, 0, 103, 101, 1,
		0, 0, 0, 103, 102, 1, 0, 0, 0, 104, 113, 1, 0, 0, 0, 105, 106, 10, 8, 0,
		0, 106, 107, 7, 2, 0, 0, 107, 112, 3, 16, 8, 9, 108, 109, 10, 7, 0, 0,
		109, 110, 7, 3, 0, 0, 110, 112, 3, 16, 8, 8, 111, 105, 1, 0, 0, 0, 111,
		108, 1, 0, 0, 0, 112, 115, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113, 114,
		1, 0, 0, 0, 114, 17, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 116, 117, 7, 4,
		0, 0, 117, 19, 1, 0, 0, 0, 11, 26, 31, 35, 39, 41, 66, 79, 86, 103, 111,
		113,
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
	SFGrammarParserELSE              = 7
	SFGrammarParserPRINT             = 8
	SFGrammarParserTRU               = 9
	SFGrammarParserFAL               = 10
	SFGrammarParserNIL               = 11
	SFGrammarParserDECLARATION_VAR   = 12
	SFGrammarParserDECLARATION_LET   = 13
	SFGrammarParserDIGIT_PRIMITIVE   = 14
	SFGrammarParserSTRING_PRIMITIVE  = 15
	SFGrammarParserID_PRIMITIVE      = 16
	SFGrammarParserNEGATION_OPERATOR = 17
	SFGrammarParserLPAREN            = 18
	SFGrammarParserRPAREN            = 19
	SFGrammarParserCOLON             = 20
	SFGrammarParserCOMMA             = 21
	SFGrammarParserSEMICOLON         = 22
	SFGrammarParserIS_               = 23
	SFGrammarParserPLUS_IS           = 24
	SFGrammarParserMINUS_IS          = 25
	SFGrammarParserQUESTION_MARK     = 26
	SFGrammarParserPLUS              = 27
	SFGrammarParserMINUS             = 28
	SFGrammarParserMULTIPLY          = 29
	SFGrammarParserDIVIDE            = 30
	SFGrammarParserMODULO            = 31
	SFGrammarParserWHITESPACE        = 32
	SFGrammarParserMULTI_COMMENT     = 33
	SFGrammarParserLINE_COMMENT      = 34
)

// SFGrammarParser rules.
const (
	SFGrammarParserRULE_start            = 0
	SFGrammarParserRULE_block            = 1
	SFGrammarParserRULE_stmts            = 2
	SFGrammarParserRULE_printstmt        = 3
	SFGrammarParserRULE_declaration      = 4
	SFGrammarParserRULE_type_declaration = 5
	SFGrammarParserRULE_assignment       = 6
	SFGrammarParserRULE_exprList         = 7
	SFGrammarParserRULE_expr             = 8
	SFGrammarParserRULE_type             = 9
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
		p.SetState(20)
		p.Block()
	}
	{
		p.SetState(21)
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
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&78080) != 0 {
		{
			p.SetState(23)
			p.Stmts()
		}

		p.SetState(28)
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
	Printstmt() IPrintstmtContext
	SEMICOLON() antlr.TerminalNode
	Declaration() IDeclarationContext
	Assignment() IAssignmentContext

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

func (s *StmtsContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SFGrammarParserSEMICOLON, 0)
}

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

	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SFGrammarParserPRINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(29)
			p.Printstmt()
		}
		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserSEMICOLON {
			{
				p.SetState(30)
				p.Match(SFGrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case SFGrammarParserDECLARATION_VAR, SFGrammarParserDECLARATION_LET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(33)
			p.Declaration()
		}
		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserSEMICOLON {
			{
				p.SetState(34)
				p.Match(SFGrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case SFGrammarParserID_PRIMITIVE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(37)
			p.Assignment()
		}
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SFGrammarParserSEMICOLON {
			{
				p.SetState(38)
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
	p.EnterRule(localctx, 6, SFGrammarParserRULE_printstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(SFGrammarParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(44)
		p.Match(SFGrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.ExprList()
	}
	{
		p.SetState(46)
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
	p.EnterRule(localctx, 8, SFGrammarParserRULE_declaration)
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeValueDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(48)
			p.Type_declaration()
		}
		{
			p.SetState(49)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.Type_()
		}
		{
			p.SetState(52)
			p.Match(SFGrammarParserIS_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.expr(0)
		}

	case 2:
		localctx = NewTypeOptionalValueDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(55)
			p.Type_declaration()
		}
		{
			p.SetState(56)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)
			p.Match(SFGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(58)
			p.Type_()
		}
		{
			p.SetState(59)
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
			p.SetState(61)
			p.Type_declaration()
		}
		{
			p.SetState(62)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.Match(SFGrammarParserIS_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
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
	p.EnterRule(localctx, 10, SFGrammarParserRULE_type_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
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

func (p *SFGrammarParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SFGrammarParserRULE_assignment)
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewValueAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.Match(SFGrammarParserIS_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)
			p.expr(0)
		}

	case 2:
		localctx = NewPlusAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.Match(SFGrammarParserPLUS_IS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.expr(0)
		}

	case 3:
		localctx = NewMinusAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(76)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Match(SFGrammarParserMINUS_IS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
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
	p.EnterRule(localctx, 14, SFGrammarParserRULE_exprList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.expr(0)
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SFGrammarParserCOMMA {
		{
			p.SetState(82)
			p.Match(SFGrammarParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.expr(0)
		}

		p.SetState(88)
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
	_startState := 16
	p.EnterRecursionRule(localctx, 16, SFGrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SFGrammarParserNEGATION_OPERATOR:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(90)
			p.Match(SFGrammarParserNEGATION_OPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)

			var _x = p.expr(10)

			localctx.(*NotExprContext).right = _x
		}

	case SFGrammarParserMINUS:
		localctx = NewNegExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(92)
			p.Match(SFGrammarParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(93)

			var _x = p.expr(9)

			localctx.(*NegExprContext).right = _x
		}

	case SFGrammarParserLPAREN:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(94)
			p.Match(SFGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.expr(0)
		}
		{
			p.SetState(96)
			p.Match(SFGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SFGrammarParserDIGIT_PRIMITIVE:
		localctx = NewDigitExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(98)
			p.Match(SFGrammarParserDIGIT_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SFGrammarParserSTRING_PRIMITIVE:
		localctx = NewStringExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(99)
			p.Match(SFGrammarParserSTRING_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SFGrammarParserID_PRIMITIVE:
		localctx = NewIdExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(100)
			p.Match(SFGrammarParserID_PRIMITIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SFGrammarParserNIL:
		localctx = NewNilExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(101)
			p.Match(SFGrammarParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SFGrammarParserTRU, SFGrammarParserFAL:
		localctx = NewBooleanExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(102)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SFGrammarParserTRU || _la == SFGrammarParserFAL) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(111)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewArithmeticOperationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticOperationExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SFGrammarParserRULE_expr)
				p.SetState(105)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(106)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ArithmeticOperationExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3758096384) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ArithmeticOperationExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(107)

					var _x = p.expr(9)

					localctx.(*ArithmeticOperationExprContext).right = _x
				}

			case 2:
				localctx = NewArithmeticOperationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticOperationExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SFGrammarParserRULE_expr)
				p.SetState(108)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(109)

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
					p.SetState(110)

					var _x = p.expr(8)

					localctx.(*ArithmeticOperationExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 18, SFGrammarParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62) != 0) {
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
	case 8:
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
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
