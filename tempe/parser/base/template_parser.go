// Code generated from /workspace/grammar/TemplateParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package base // TemplateParser
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

type TemplateParser struct {
	*antlr.BaseParser
}

var TemplateParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func templateparserParserInit() {
	staticData := &TemplateParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'{{'", "", "", "", "'}}'", "", "", "", "", "", "", "", "", "",
		"", "", "','",
	}
	staticData.SymbolicNames = []string{
		"", "OPEN_EXPR", "OPEN_STMT", "LINE", "TEXT", "CLOSE_EXPR", "EXPR",
		"CLOSE_STMT", "FOR", "IN", "ENDFOR", "IF", "ELSEIF", "ELSE", "ENDIF",
		"NAME", "EQ", "COMMA", "WS", "CLOSE_STMT_EXPR", "STMT_EXPR", "SEMICOLON",
		"ASSIGN_EXPR",
	}
	staticData.RuleNames = []string{
		"template", "node", "expression", "statement", "assignment", "var",
		"forLoop", "startFor", "endFor", "ifStatement", "startIf", "elseIf",
		"else", "endIf", "name", "index", "expr", "text",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 22, 168, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 5, 0, 38, 8, 0, 10, 0, 12, 0, 41, 9,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 48, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 3, 3, 3, 57, 8, 3, 1, 4, 1, 4, 4, 4, 61, 8, 4, 11, 4, 12,
		4, 62, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 5, 6, 74,
		8, 6, 10, 6, 12, 6, 77, 9, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		5, 7, 86, 8, 7, 10, 7, 12, 7, 89, 9, 7, 1, 7, 1, 7, 5, 7, 93, 8, 7, 10,
		7, 12, 7, 96, 9, 7, 1, 7, 1, 7, 3, 7, 100, 8, 7, 1, 7, 4, 7, 103, 8, 7,
		11, 7, 12, 7, 104, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9,
		1, 9, 5, 9, 117, 8, 9, 10, 9, 12, 9, 120, 9, 9, 1, 9, 3, 9, 123, 8, 9,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 132, 8, 10, 10, 10,
		12, 10, 135, 9, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 142, 8, 11,
		10, 11, 12, 11, 145, 9, 11, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 151, 8,
		12, 10, 12, 12, 12, 154, 9, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 0, 0, 18, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 0, 2, 3, 0, 6, 6,
		20, 20, 22, 22, 1, 0, 3, 4, 165, 0, 39, 1, 0, 0, 0, 2, 47, 1, 0, 0, 0,
		4, 49, 1, 0, 0, 0, 6, 56, 1, 0, 0, 0, 8, 58, 1, 0, 0, 0, 10, 66, 1, 0,
		0, 0, 12, 71, 1, 0, 0, 0, 14, 80, 1, 0, 0, 0, 16, 110, 1, 0, 0, 0, 18,
		114, 1, 0, 0, 0, 20, 126, 1, 0, 0, 0, 22, 136, 1, 0, 0, 0, 24, 146, 1,
		0, 0, 0, 26, 155, 1, 0, 0, 0, 28, 159, 1, 0, 0, 0, 30, 161, 1, 0, 0, 0,
		32, 163, 1, 0, 0, 0, 34, 165, 1, 0, 0, 0, 36, 38, 3, 2, 1, 0, 37, 36, 1,
		0, 0, 0, 38, 41, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40,
		42, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 42, 43, 5, 0, 0, 1, 43, 1, 1, 0, 0,
		0, 44, 48, 3, 4, 2, 0, 45, 48, 3, 6, 3, 0, 46, 48, 3, 34, 17, 0, 47, 44,
		1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 47, 46, 1, 0, 0, 0, 48, 3, 1, 0, 0, 0,
		49, 50, 5, 1, 0, 0, 50, 51, 3, 32, 16, 0, 51, 52, 5, 5, 0, 0, 52, 5, 1,
		0, 0, 0, 53, 57, 3, 8, 4, 0, 54, 57, 3, 12, 6, 0, 55, 57, 3, 18, 9, 0,
		56, 53, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 56, 55, 1, 0, 0, 0, 57, 7, 1, 0,
		0, 0, 58, 60, 5, 2, 0, 0, 59, 61, 3, 10, 5, 0, 60, 59, 1, 0, 0, 0, 61,
		62, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 64, 1, 0, 0,
		0, 64, 65, 5, 7, 0, 0, 65, 9, 1, 0, 0, 0, 66, 67, 3, 28, 14, 0, 67, 68,
		5, 16, 0, 0, 68, 69, 3, 32, 16, 0, 69, 70, 5, 21, 0, 0, 70, 11, 1, 0, 0,
		0, 71, 75, 3, 14, 7, 0, 72, 74, 3, 2, 1, 0, 73, 72, 1, 0, 0, 0, 74, 77,
		1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 78, 1, 0, 0, 0,
		77, 75, 1, 0, 0, 0, 78, 79, 3, 16, 8, 0, 79, 13, 1, 0, 0, 0, 80, 81, 5,
		2, 0, 0, 81, 99, 5, 8, 0, 0, 82, 100, 3, 28, 14, 0, 83, 87, 3, 30, 15,
		0, 84, 86, 5, 18, 0, 0, 85, 84, 1, 0, 0, 0, 86, 89, 1, 0, 0, 0, 87, 85,
		1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 90, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0,
		90, 94, 5, 17, 0, 0, 91, 93, 5, 18, 0, 0, 92, 91, 1, 0, 0, 0, 93, 96, 1,
		0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 97, 1, 0, 0, 0, 96,
		94, 1, 0, 0, 0, 97, 98, 3, 28, 14, 0, 98, 100, 1, 0, 0, 0, 99, 82, 1, 0,
		0, 0, 99, 83, 1, 0, 0, 0, 100, 102, 1, 0, 0, 0, 101, 103, 5, 18, 0, 0,
		102, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 104,
		105, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 107, 5, 9, 0, 0, 107, 108,
		3, 32, 16, 0, 108, 109, 5, 19, 0, 0, 109, 15, 1, 0, 0, 0, 110, 111, 5,
		2, 0, 0, 111, 112, 5, 10, 0, 0, 112, 113, 5, 7, 0, 0, 113, 17, 1, 0, 0,
		0, 114, 118, 3, 20, 10, 0, 115, 117, 3, 22, 11, 0, 116, 115, 1, 0, 0, 0,
		117, 120, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119,
		122, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 121, 123, 3, 24, 12, 0, 122, 121,
		1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 125, 3, 26,
		13, 0, 125, 19, 1, 0, 0, 0, 126, 127, 5, 2, 0, 0, 127, 128, 5, 11, 0, 0,
		128, 129, 3, 32, 16, 0, 129, 133, 5, 19, 0, 0, 130, 132, 3, 2, 1, 0, 131,
		130, 1, 0, 0, 0, 132, 135, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134,
		1, 0, 0, 0, 134, 21, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 136, 137, 5, 2,
		0, 0, 137, 138, 5, 12, 0, 0, 138, 139, 3, 32, 16, 0, 139, 143, 5, 19, 0,
		0, 140, 142, 3, 2, 1, 0, 141, 140, 1, 0, 0, 0, 142, 145, 1, 0, 0, 0, 143,
		141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 23, 1, 0, 0, 0, 145, 143, 1,
		0, 0, 0, 146, 147, 5, 2, 0, 0, 147, 148, 5, 13, 0, 0, 148, 152, 5, 7, 0,
		0, 149, 151, 3, 2, 1, 0, 150, 149, 1, 0, 0, 0, 151, 154, 1, 0, 0, 0, 152,
		150, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 25, 1, 0, 0, 0, 154, 152, 1,
		0, 0, 0, 155, 156, 5, 2, 0, 0, 156, 157, 5, 14, 0, 0, 157, 158, 5, 7, 0,
		0, 158, 27, 1, 0, 0, 0, 159, 160, 5, 15, 0, 0, 160, 29, 1, 0, 0, 0, 161,
		162, 5, 15, 0, 0, 162, 31, 1, 0, 0, 0, 163, 164, 7, 0, 0, 0, 164, 33, 1,
		0, 0, 0, 165, 166, 7, 1, 0, 0, 166, 35, 1, 0, 0, 0, 14, 39, 47, 56, 62,
		75, 87, 94, 99, 104, 118, 122, 133, 143, 152,
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

// TemplateParserInit initializes any static state used to implement TemplateParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTemplateParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TemplateParserInit() {
	staticData := &TemplateParserParserStaticData
	staticData.once.Do(templateparserParserInit)
}

// NewTemplateParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTemplateParser(input antlr.TokenStream) *TemplateParser {
	TemplateParserInit()
	this := new(TemplateParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &TemplateParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "TemplateParser.g4"

	return this
}

// TemplateParser tokens.
const (
	TemplateParserEOF             = antlr.TokenEOF
	TemplateParserOPEN_EXPR       = 1
	TemplateParserOPEN_STMT       = 2
	TemplateParserLINE            = 3
	TemplateParserTEXT            = 4
	TemplateParserCLOSE_EXPR      = 5
	TemplateParserEXPR            = 6
	TemplateParserCLOSE_STMT      = 7
	TemplateParserFOR             = 8
	TemplateParserIN              = 9
	TemplateParserENDFOR          = 10
	TemplateParserIF              = 11
	TemplateParserELSEIF          = 12
	TemplateParserELSE            = 13
	TemplateParserENDIF           = 14
	TemplateParserNAME            = 15
	TemplateParserEQ              = 16
	TemplateParserCOMMA           = 17
	TemplateParserWS              = 18
	TemplateParserCLOSE_STMT_EXPR = 19
	TemplateParserSTMT_EXPR       = 20
	TemplateParserSEMICOLON       = 21
	TemplateParserASSIGN_EXPR     = 22
)

// TemplateParser rules.
const (
	TemplateParserRULE_template    = 0
	TemplateParserRULE_node        = 1
	TemplateParserRULE_expression  = 2
	TemplateParserRULE_statement   = 3
	TemplateParserRULE_assignment  = 4
	TemplateParserRULE_var         = 5
	TemplateParserRULE_forLoop     = 6
	TemplateParserRULE_startFor    = 7
	TemplateParserRULE_endFor      = 8
	TemplateParserRULE_ifStatement = 9
	TemplateParserRULE_startIf     = 10
	TemplateParserRULE_elseIf      = 11
	TemplateParserRULE_else        = 12
	TemplateParserRULE_endIf       = 13
	TemplateParserRULE_name        = 14
	TemplateParserRULE_index       = 15
	TemplateParserRULE_expr        = 16
	TemplateParserRULE_text        = 17
)

// ITemplateContext is an interface to support dynamic dispatch.
type ITemplateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllNode() []INodeContext
	Node(i int) INodeContext

	// IsTemplateContext differentiates from other interfaces.
	IsTemplateContext()
}

type TemplateContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemplateContext() *TemplateContext {
	var p = new(TemplateContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_template
	return p
}

func InitEmptyTemplateContext(p *TemplateContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_template
}

func (*TemplateContext) IsTemplateContext() {}

func NewTemplateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemplateContext {
	var p = new(TemplateContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TemplateParserRULE_template

	return p
}

func (s *TemplateContext) GetParser() antlr.Parser { return s.parser }

func (s *TemplateContext) EOF() antlr.TerminalNode {
	return s.GetToken(TemplateParserEOF, 0)
}

func (s *TemplateContext) AllNode() []INodeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INodeContext); ok {
			len++
		}
	}

	tst := make([]INodeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INodeContext); ok {
			tst[i] = t.(INodeContext)
			i++
		}
	}

	return tst
}

func (s *TemplateContext) Node(i int) INodeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INodeContext); ok {
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

	return t.(INodeContext)
}

func (s *TemplateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemplateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TemplateParserVisitor:
		return t.VisitTemplate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TemplateParser) Template() (localctx ITemplateContext) {
	localctx = NewTemplateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TemplateParserRULE_template)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30) != 0 {
		{
			p.SetState(36)
			p.Node()
		}

		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(42)
		p.Match(TemplateParserEOF)
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

// INodeContext is an interface to support dynamic dispatch.
type INodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	Statement() IStatementContext
	Text() ITextContext

	// IsNodeContext differentiates from other interfaces.
	IsNodeContext()
}

type NodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNodeContext() *NodeContext {
	var p = new(NodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_node
	return p
}

func InitEmptyNodeContext(p *NodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_node
}

func (*NodeContext) IsNodeContext() {}

func NewNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NodeContext {
	var p = new(NodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TemplateParserRULE_node

	return p
}

func (s *NodeContext) GetParser() antlr.Parser { return s.parser }

func (s *NodeContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NodeContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *NodeContext) Text() ITextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *NodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TemplateParserVisitor:
		return t.VisitNode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TemplateParser) Node() (localctx INodeContext) {
	localctx = NewNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TemplateParserRULE_node)
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TemplateParserOPEN_EXPR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(44)
			p.Expression()
		}

	case TemplateParserOPEN_STMT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)
			p.Statement()
		}

	case TemplateParserLINE, TemplateParserTEXT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(46)
			p.Text()
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_EXPR() antlr.TerminalNode
	Expr() IExprContext
	CLOSE_EXPR() antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TemplateParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) OPEN_EXPR() antlr.TerminalNode {
	return s.GetToken(TemplateParserOPEN_EXPR, 0)
}

func (s *ExpressionContext) Expr() IExprContext {
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

func (s *ExpressionContext) CLOSE_EXPR() antlr.TerminalNode {
	return s.GetToken(TemplateParserCLOSE_EXPR, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TemplateParserVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TemplateParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TemplateParserRULE_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Match(TemplateParserOPEN_EXPR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Expr()
	}
	{
		p.SetState(51)
		p.Match(TemplateParserCLOSE_EXPR)
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
	Assignment() IAssignmentContext
	ForLoop() IForLoopContext
	IfStatement() IIfStatementContext

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
	p.RuleIndex = TemplateParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TemplateParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Assignment() IAssignmentContext {
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

func (s *StatementContext) ForLoop() IForLoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForLoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForLoopContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TemplateParserVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TemplateParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TemplateParserRULE_statement)
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)
			p.Assignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.ForLoop()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(55)
			p.IfStatement()
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

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_STMT() antlr.TerminalNode
	CLOSE_STMT() antlr.TerminalNode
	AllVar_() []IVarContext
	Var_(i int) IVarContext

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
	p.RuleIndex = TemplateParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TemplateParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) OPEN_STMT() antlr.TerminalNode {
	return s.GetToken(TemplateParserOPEN_STMT, 0)
}

func (s *AssignmentContext) CLOSE_STMT() antlr.TerminalNode {
	return s.GetToken(TemplateParserCLOSE_STMT, 0)
}

func (s *AssignmentContext) AllVar_() []IVarContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarContext); ok {
			len++
		}
	}

	tst := make([]IVarContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarContext); ok {
			tst[i] = t.(IVarContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentContext) Var_(i int) IVarContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarContext); ok {
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

	return t.(IVarContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TemplateParserVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TemplateParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TemplateParserRULE_assignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(TemplateParserOPEN_STMT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == TemplateParserNAME {
		{
			p.SetState(59)
			p.Var_()
		}

		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(64)
		p.Match(TemplateParserCLOSE_STMT)
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

// IVarContext is an interface to support dynamic dispatch.
type IVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	EQ() antlr.TerminalNode
	Expr() IExprContext
	SEMICOLON() antlr.TerminalNode

	// IsVarContext differentiates from other interfaces.
	IsVarContext()
}

type VarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarContext() *VarContext {
	var p = new(VarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_var
	return p
}

func InitEmptyVarContext(p *VarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_var
}

func (*VarContext) IsVarContext() {}

func NewVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarContext {
	var p = new(VarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TemplateParserRULE_var

	return p
}

func (s *VarContext) GetParser() antlr.Parser { return s.parser }

func (s *VarContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *VarContext) EQ() antlr.TerminalNode {
	return s.GetToken(TemplateParserEQ, 0)
}

func (s *VarContext) Expr() IExprContext {
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

func (s *VarContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(TemplateParserSEMICOLON, 0)
}

func (s *VarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TemplateParserVisitor:
		return t.VisitVar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TemplateParser) Var_() (localctx IVarContext) {
	localctx = NewVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TemplateParserRULE_var)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Name()
	}
	{
		p.SetState(67)
		p.Match(TemplateParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(68)
		p.Expr()
	}
	{
		p.SetState(69)
		p.Match(TemplateParserSEMICOLON)
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

// IForLoopContext is an interface to support dynamic dispatch.
type IForLoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StartFor() IStartForContext
	EndFor() IEndForContext
	AllNode() []INodeContext
	Node(i int) INodeContext

	// IsForLoopContext differentiates from other interfaces.
	IsForLoopContext()
}

type ForLoopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForLoopContext() *ForLoopContext {
	var p = new(ForLoopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_forLoop
	return p
}

func InitEmptyForLoopContext(p *ForLoopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_forLoop
}

func (*ForLoopContext) IsForLoopContext() {}

func NewForLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForLoopContext {
	var p = new(ForLoopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TemplateParserRULE_forLoop

	return p
}

func (s *ForLoopContext) GetParser() antlr.Parser { return s.parser }

func (s *ForLoopContext) StartFor() IStartForContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStartForContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStartForContext)
}

func (s *ForLoopContext) EndFor() IEndForContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEndForContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEndForContext)
}

func (s *ForLoopContext) AllNode() []INodeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INodeContext); ok {
			len++
		}
	}

	tst := make([]INodeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INodeContext); ok {
			tst[i] = t.(INodeContext)
			i++
		}
	}

	return tst
}

func (s *ForLoopContext) Node(i int) INodeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INodeContext); ok {
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

	return t.(INodeContext)
}

func (s *ForLoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForLoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForLoopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TemplateParserVisitor:
		return t.VisitForLoop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TemplateParser) ForLoop() (localctx IForLoopContext) {
	localctx = NewForLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TemplateParserRULE_forLoop)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.StartFor()
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(72)
				p.Node()
			}

		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(78)
		p.EndFor()
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

// IStartForContext is an interface to support dynamic dispatch.
type IStartForContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_STMT() antlr.TerminalNode
	FOR() antlr.TerminalNode
	IN() antlr.TerminalNode
	Expr() IExprContext
	CLOSE_STMT_EXPR() antlr.TerminalNode
	Name() INameContext
	Index() IIndexContext
	COMMA() antlr.TerminalNode
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode

	// IsStartForContext differentiates from other interfaces.
	IsStartForContext()
}

type StartForContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartForContext() *StartForContext {
	var p = new(StartForContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_startFor
	return p
}

func InitEmptyStartForContext(p *StartForContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_startFor
}

func (*StartForContext) IsStartForContext() {}

func NewStartForContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartForContext {
	var p = new(StartForContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TemplateParserRULE_startFor

	return p
}

func (s *StartForContext) GetParser() antlr.Parser { return s.parser }

func (s *StartForContext) OPEN_STMT() antlr.TerminalNode {
	return s.GetToken(TemplateParserOPEN_STMT, 0)
}

func (s *StartForContext) FOR() antlr.TerminalNode {
	return s.GetToken(TemplateParserFOR, 0)
}

func (s *StartForContext) IN() antlr.TerminalNode {
	return s.GetToken(TemplateParserIN, 0)
}

func (s *StartForContext) Expr() IExprContext {
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

func (s *StartForContext) CLOSE_STMT_EXPR() antlr.TerminalNode {
	return s.GetToken(TemplateParserCLOSE_STMT_EXPR, 0)
}

func (s *StartForContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *StartForContext) Index() IIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *StartForContext) COMMA() antlr.TerminalNode {
	return s.GetToken(TemplateParserCOMMA, 0)
}

func (s *StartForContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(TemplateParserWS)
}

func (s *StartForContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(TemplateParserWS, i)
}

func (s *StartForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartForContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartForContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TemplateParserVisitor:
		return t.VisitStartFor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TemplateParser) StartFor() (localctx IStartForContext) {
	localctx = NewStartForContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TemplateParserRULE_startFor)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.Match(TemplateParserOPEN_STMT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
		p.Match(TemplateParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(82)
			p.Name()
		}

	case 2:
		{
			p.SetState(83)
			p.Index()
		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TemplateParserWS {
			{
				p.SetState(84)
				p.Match(TemplateParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(89)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(90)
			p.Match(TemplateParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TemplateParserWS {
			{
				p.SetState(91)
				p.Match(TemplateParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(96)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(97)
			p.Name()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == TemplateParserWS {
		{
			p.SetState(101)
			p.Match(TemplateParserWS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(106)
		p.Match(TemplateParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Expr()
	}
	{
		p.SetState(108)
		p.Match(TemplateParserCLOSE_STMT_EXPR)
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

// IEndForContext is an interface to support dynamic dispatch.
type IEndForContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_STMT() antlr.TerminalNode
	ENDFOR() antlr.TerminalNode
	CLOSE_STMT() antlr.TerminalNode

	// IsEndForContext differentiates from other interfaces.
	IsEndForContext()
}

type EndForContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndForContext() *EndForContext {
	var p = new(EndForContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_endFor
	return p
}

func InitEmptyEndForContext(p *EndForContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_endFor
}

func (*EndForContext) IsEndForContext() {}

func NewEndForContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndForContext {
	var p = new(EndForContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TemplateParserRULE_endFor

	return p
}

func (s *EndForContext) GetParser() antlr.Parser { return s.parser }

func (s *EndForContext) OPEN_STMT() antlr.TerminalNode {
	return s.GetToken(TemplateParserOPEN_STMT, 0)
}

func (s *EndForContext) ENDFOR() antlr.TerminalNode {
	return s.GetToken(TemplateParserENDFOR, 0)
}

func (s *EndForContext) CLOSE_STMT() antlr.TerminalNode {
	return s.GetToken(TemplateParserCLOSE_STMT, 0)
}

func (s *EndForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndForContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndForContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TemplateParserVisitor:
		return t.VisitEndFor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TemplateParser) EndFor() (localctx IEndForContext) {
	localctx = NewEndForContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TemplateParserRULE_endFor)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.Match(TemplateParserOPEN_STMT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.Match(TemplateParserENDFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.Match(TemplateParserCLOSE_STMT)
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

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StartIf() IStartIfContext
	EndIf() IEndIfContext
	AllElseIf() []IElseIfContext
	ElseIf(i int) IElseIfContext
	Else_() IElseContext

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TemplateParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) StartIf() IStartIfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStartIfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStartIfContext)
}

func (s *IfStatementContext) EndIf() IEndIfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEndIfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEndIfContext)
}

func (s *IfStatementContext) AllElseIf() []IElseIfContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseIfContext); ok {
			len++
		}
	}

	tst := make([]IElseIfContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseIfContext); ok {
			tst[i] = t.(IElseIfContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) ElseIf(i int) IElseIfContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseIfContext); ok {
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

	return t.(IElseIfContext)
}

func (s *IfStatementContext) Else_() IElseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TemplateParserVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TemplateParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TemplateParserRULE_ifStatement)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.StartIf()
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(115)
				p.ElseIf()
			}

		}
		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(121)
			p.Else_()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(124)
		p.EndIf()
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

// IStartIfContext is an interface to support dynamic dispatch.
type IStartIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_STMT() antlr.TerminalNode
	IF() antlr.TerminalNode
	Expr() IExprContext
	CLOSE_STMT_EXPR() antlr.TerminalNode
	AllNode() []INodeContext
	Node(i int) INodeContext

	// IsStartIfContext differentiates from other interfaces.
	IsStartIfContext()
}

type StartIfContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartIfContext() *StartIfContext {
	var p = new(StartIfContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_startIf
	return p
}

func InitEmptyStartIfContext(p *StartIfContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_startIf
}

func (*StartIfContext) IsStartIfContext() {}

func NewStartIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartIfContext {
	var p = new(StartIfContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TemplateParserRULE_startIf

	return p
}

func (s *StartIfContext) GetParser() antlr.Parser { return s.parser }

func (s *StartIfContext) OPEN_STMT() antlr.TerminalNode {
	return s.GetToken(TemplateParserOPEN_STMT, 0)
}

func (s *StartIfContext) IF() antlr.TerminalNode {
	return s.GetToken(TemplateParserIF, 0)
}

func (s *StartIfContext) Expr() IExprContext {
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

func (s *StartIfContext) CLOSE_STMT_EXPR() antlr.TerminalNode {
	return s.GetToken(TemplateParserCLOSE_STMT_EXPR, 0)
}

func (s *StartIfContext) AllNode() []INodeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INodeContext); ok {
			len++
		}
	}

	tst := make([]INodeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INodeContext); ok {
			tst[i] = t.(INodeContext)
			i++
		}
	}

	return tst
}

func (s *StartIfContext) Node(i int) INodeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INodeContext); ok {
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

	return t.(INodeContext)
}

func (s *StartIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartIfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartIfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TemplateParserVisitor:
		return t.VisitStartIf(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TemplateParser) StartIf() (localctx IStartIfContext) {
	localctx = NewStartIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TemplateParserRULE_startIf)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(TemplateParserOPEN_STMT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
		p.Match(TemplateParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
		p.Expr()
	}
	{
		p.SetState(129)
		p.Match(TemplateParserCLOSE_STMT_EXPR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(130)
				p.Node()
			}

		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
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

// IElseIfContext is an interface to support dynamic dispatch.
type IElseIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_STMT() antlr.TerminalNode
	ELSEIF() antlr.TerminalNode
	Expr() IExprContext
	CLOSE_STMT_EXPR() antlr.TerminalNode
	AllNode() []INodeContext
	Node(i int) INodeContext

	// IsElseIfContext differentiates from other interfaces.
	IsElseIfContext()
}

type ElseIfContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfContext() *ElseIfContext {
	var p = new(ElseIfContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_elseIf
	return p
}

func InitEmptyElseIfContext(p *ElseIfContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_elseIf
}

func (*ElseIfContext) IsElseIfContext() {}

func NewElseIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfContext {
	var p = new(ElseIfContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TemplateParserRULE_elseIf

	return p
}

func (s *ElseIfContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfContext) OPEN_STMT() antlr.TerminalNode {
	return s.GetToken(TemplateParserOPEN_STMT, 0)
}

func (s *ElseIfContext) ELSEIF() antlr.TerminalNode {
	return s.GetToken(TemplateParserELSEIF, 0)
}

func (s *ElseIfContext) Expr() IExprContext {
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

func (s *ElseIfContext) CLOSE_STMT_EXPR() antlr.TerminalNode {
	return s.GetToken(TemplateParserCLOSE_STMT_EXPR, 0)
}

func (s *ElseIfContext) AllNode() []INodeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INodeContext); ok {
			len++
		}
	}

	tst := make([]INodeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INodeContext); ok {
			tst[i] = t.(INodeContext)
			i++
		}
	}

	return tst
}

func (s *ElseIfContext) Node(i int) INodeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INodeContext); ok {
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

	return t.(INodeContext)
}

func (s *ElseIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseIfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TemplateParserVisitor:
		return t.VisitElseIf(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TemplateParser) ElseIf() (localctx IElseIfContext) {
	localctx = NewElseIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TemplateParserRULE_elseIf)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(TemplateParserOPEN_STMT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.Match(TemplateParserELSEIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.Expr()
	}
	{
		p.SetState(139)
		p.Match(TemplateParserCLOSE_STMT_EXPR)
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
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(140)
				p.Node()
			}

		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
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

// IElseContext is an interface to support dynamic dispatch.
type IElseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_STMT() antlr.TerminalNode
	ELSE() antlr.TerminalNode
	CLOSE_STMT() antlr.TerminalNode
	AllNode() []INodeContext
	Node(i int) INodeContext

	// IsElseContext differentiates from other interfaces.
	IsElseContext()
}

type ElseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseContext() *ElseContext {
	var p = new(ElseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_else
	return p
}

func InitEmptyElseContext(p *ElseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_else
}

func (*ElseContext) IsElseContext() {}

func NewElseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseContext {
	var p = new(ElseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TemplateParserRULE_else

	return p
}

func (s *ElseContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseContext) OPEN_STMT() antlr.TerminalNode {
	return s.GetToken(TemplateParserOPEN_STMT, 0)
}

func (s *ElseContext) ELSE() antlr.TerminalNode {
	return s.GetToken(TemplateParserELSE, 0)
}

func (s *ElseContext) CLOSE_STMT() antlr.TerminalNode {
	return s.GetToken(TemplateParserCLOSE_STMT, 0)
}

func (s *ElseContext) AllNode() []INodeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INodeContext); ok {
			len++
		}
	}

	tst := make([]INodeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INodeContext); ok {
			tst[i] = t.(INodeContext)
			i++
		}
	}

	return tst
}

func (s *ElseContext) Node(i int) INodeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INodeContext); ok {
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

	return t.(INodeContext)
}

func (s *ElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TemplateParserVisitor:
		return t.VisitElse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TemplateParser) Else_() (localctx IElseContext) {
	localctx = NewElseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TemplateParserRULE_else)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		p.Match(TemplateParserOPEN_STMT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(147)
		p.Match(TemplateParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.Match(TemplateParserCLOSE_STMT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(149)
				p.Node()
			}

		}
		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
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

// IEndIfContext is an interface to support dynamic dispatch.
type IEndIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_STMT() antlr.TerminalNode
	ENDIF() antlr.TerminalNode
	CLOSE_STMT() antlr.TerminalNode

	// IsEndIfContext differentiates from other interfaces.
	IsEndIfContext()
}

type EndIfContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndIfContext() *EndIfContext {
	var p = new(EndIfContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_endIf
	return p
}

func InitEmptyEndIfContext(p *EndIfContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_endIf
}

func (*EndIfContext) IsEndIfContext() {}

func NewEndIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndIfContext {
	var p = new(EndIfContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TemplateParserRULE_endIf

	return p
}

func (s *EndIfContext) GetParser() antlr.Parser { return s.parser }

func (s *EndIfContext) OPEN_STMT() antlr.TerminalNode {
	return s.GetToken(TemplateParserOPEN_STMT, 0)
}

func (s *EndIfContext) ENDIF() antlr.TerminalNode {
	return s.GetToken(TemplateParserENDIF, 0)
}

func (s *EndIfContext) CLOSE_STMT() antlr.TerminalNode {
	return s.GetToken(TemplateParserCLOSE_STMT, 0)
}

func (s *EndIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndIfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndIfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TemplateParserVisitor:
		return t.VisitEndIf(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TemplateParser) EndIf() (localctx IEndIfContext) {
	localctx = NewEndIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, TemplateParserRULE_endIf)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(TemplateParserOPEN_STMT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Match(TemplateParserENDIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)
		p.Match(TemplateParserCLOSE_STMT)
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

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_name
	return p
}

func InitEmptyNameContext(p *NameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_name
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TemplateParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) NAME() antlr.TerminalNode {
	return s.GetToken(TemplateParserNAME, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TemplateParserVisitor:
		return t.VisitName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TemplateParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, TemplateParserRULE_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Match(TemplateParserNAME)
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

// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode

	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_index
	return p
}

func InitEmptyIndexContext(p *IndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_index
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TemplateParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) NAME() antlr.TerminalNode {
	return s.GetToken(TemplateParserNAME, 0)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TemplateParserVisitor:
		return t.VisitIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TemplateParser) Index() (localctx IIndexContext) {
	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, TemplateParserRULE_index)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(TemplateParserNAME)
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

	// Getter signatures
	EXPR() antlr.TerminalNode
	STMT_EXPR() antlr.TerminalNode
	ASSIGN_EXPR() antlr.TerminalNode

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
	p.RuleIndex = TemplateParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TemplateParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) EXPR() antlr.TerminalNode {
	return s.GetToken(TemplateParserEXPR, 0)
}

func (s *ExprContext) STMT_EXPR() antlr.TerminalNode {
	return s.GetToken(TemplateParserSTMT_EXPR, 0)
}

func (s *ExprContext) ASSIGN_EXPR() antlr.TerminalNode {
	return s.GetToken(TemplateParserASSIGN_EXPR, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TemplateParserVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TemplateParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, TemplateParserRULE_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&5242944) != 0) {
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

// ITextContext is an interface to support dynamic dispatch.
type ITextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TEXT() antlr.TerminalNode
	LINE() antlr.TerminalNode

	// IsTextContext differentiates from other interfaces.
	IsTextContext()
}

type TextContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextContext() *TextContext {
	var p = new(TextContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_text
	return p
}

func InitEmptyTextContext(p *TextContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TemplateParserRULE_text
}

func (*TextContext) IsTextContext() {}

func NewTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextContext {
	var p = new(TextContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TemplateParserRULE_text

	return p
}

func (s *TextContext) GetParser() antlr.Parser { return s.parser }

func (s *TextContext) TEXT() antlr.TerminalNode {
	return s.GetToken(TemplateParserTEXT, 0)
}

func (s *TextContext) LINE() antlr.TerminalNode {
	return s.GetToken(TemplateParserLINE, 0)
}

func (s *TextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TemplateParserVisitor:
		return t.VisitText(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TemplateParser) Text() (localctx ITextContext) {
	localctx = NewTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, TemplateParserRULE_text)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TemplateParserLINE || _la == TemplateParserTEXT) {
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
