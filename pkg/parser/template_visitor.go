package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/tniedbala/tempe/pkg/api"
	"github.com/tniedbala/tempe/pkg/parser/base"
)

type TemplateVisitor struct {
	*base.BaseTemplateParserVisitor
}

func (v TemplateVisitor) VisitTemplate(ctx base.ITemplateContext) *TemplateNodesCollection {
	return v.VisitNodes(ctx.AllNode())
}

func (v TemplateVisitor) VisitNodes(nodes []base.INodeContext) *TemplateNodesCollection {
	n := NewTemplateNodesCollection()
	for _, node := range nodes {
		n.Append(v.VisitNode(node))
	}
	return n
}

func (v TemplateVisitor) VisitNode(ctx base.INodeContext) api.TemplateNode {
	if text := ctx.Text(); text != nil {
		return v.VisitText(text)
	}
	if expr := ctx.Expression(); expr != nil {
		return v.VisitExpression(expr)
	}
	if stmt := ctx.Statement(); stmt != nil {
		return v.VisitStatement(stmt)
	}
	return v.VisitChildren(ctx)
}

func (v TemplateVisitor) VisitText(ctx base.ITextContext) *Text {
	return NewText(ctx.GetText())
}

func (v TemplateVisitor) VisitExpression(ctx base.IExpressionContext) *Expression {
	return NewExpression(ctx.Expr().GetText())
}

func (v TemplateVisitor) VisitStatement(ctx base.IStatementContext) api.TemplateNode {
	if assign := ctx.Assignment(); assign != nil {
		return v.VisitAssignment(assign)
	}
	if forLoop := ctx.ForLoop(); forLoop != nil {
		return v.VisitForLoop(forLoop)
	}
	if ifStatement := ctx.IfStatement(); ifStatement != nil {
		return v.VisitIfStatement(ifStatement)
	}
	return v.VisitChildren(ctx)
}

func (v TemplateVisitor) VisitAssignment(ctx base.IAssignmentContext) *Assignment {
	assignment := NewAssignment()
	for _, v := range ctx.AllVar_() {
		name, expr := v.Name().GetText(), v.Expr().GetText()
		assignment.Append(name, expr)
	}
	return assignment
}

func (v TemplateVisitor) VisitForLoop(ctx base.IForLoopContext) *ForLoop {
	startFor := ctx.StartFor()
	varName := startFor.Name().GetText()
	indexName, index := "", startFor.Index()
	if index != nil {
		indexName = index.GetText()
	}
	expr := startFor.Expr().GetText()
	body := v.VisitNodes(ctx.AllNode())
	return NewForLoop(indexName, varName, expr, body)
}

func (v TemplateVisitor) VisitIfStatement(ctx base.IIfStatementContext) *IfStatement {
	clauses := NewTemplateNodesCollection()
	startIf := ctx.StartIf()
	expr, body := startIf.Expr().GetText(), v.VisitNodes(startIf.AllNode())
	clauses.Append(NewIfClause(If, expr, body))
	for _, elseIf := range ctx.AllElseIf() {
		expr, body = elseIf.Expr().GetText(), v.VisitNodes(elseIf.AllNode())
		clauses.Append(NewIfClause(ElseIf, expr, body))
	}
	if elseCtx := ctx.Else_(); elseCtx != nil {
		body = v.VisitNodes(elseCtx.AllNode())
		clauses.Append(NewIfClause(Else, "", body))
	}
	return NewIfStatement(clauses)
}

func (v TemplateVisitor) VisitErrorNode(ctx antlr.ErrorNode) api.TemplateNode {
	// TODO: improve error handling
	panic(fmt.Sprintf("ERROR: %v", ctx))
}

func (v TemplateVisitor) VisitChildren(ctx antlr.RuleNode) *TemplateNodesCollection {
	nodes := NewTemplateNodesCollection()
	for _, child := range ctx.GetChildren() {
		nodes.Append(v.Visit(child))
	}
	return nodes
}

func (v TemplateVisitor) Visit(tree antlr.Tree) api.TemplateNode {
	switch ctx := tree.(type) {
	case base.ITemplateContext:
		return v.VisitTemplate(ctx)
	case base.INodeContext:
		return v.VisitNode(ctx)
	case base.ITextContext:
		return v.VisitText(ctx)
	case base.IExpressionContext:
		return v.VisitExpression(ctx)
	case base.IStatementContext:
		return v.VisitStatement(ctx)
	case base.IAssignmentContext:
		return v.VisitAssignment(ctx)
	case base.IForLoopContext:
		return v.VisitForLoop(ctx)
	case base.IIfStatementContext:
		return v.VisitIfStatement(ctx)
	case antlr.ErrorNode:
		return v.VisitErrorNode(ctx)
	default:
		// TODO: improve error message
		panic(fmt.Sprintf("invalid parser rule: %v", tree))
	}
}
