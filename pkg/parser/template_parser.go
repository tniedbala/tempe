package parser

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/tniedbala/tempe-go/pkg/parser/base"
)

type TemplateParser struct {
	src        string
	lexer      *base.TemplateLexer
	tokens     *antlr.CommonTokenStream
	baseParser *base.TemplateParser
	tree        base.ITemplateContext
}

func NewTemplateParser() *TemplateParser {
	return &TemplateParser{}
}

func (p *TemplateParser) Parse(src string) (*ParseTree, error) {
	p.src = src
	inputStream := antlr.NewInputStream(src)
	p.lexer = base.NewTemplateLexer(inputStream)
	p.tokens = antlr.NewCommonTokenStream(p.lexer, 0)
	p.baseParser = base.NewTemplateParser(p.tokens)
	// p.baseParser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	if p.baseParser.HasError() {
		ex := p.baseParser.SynErr
		token := ex.GetOffendingToken()
		return nil, NewSyntaxError(token.GetLine(), token.GetColumn(), ex.GetMessage())
	}
	p.tree = p.baseParser.Template()
	visitor := TemplateVisitor{}
	nodes := visitor.VisitTemplate(p.tree)
	return NewParseTree(nodes, nil), nil
}
