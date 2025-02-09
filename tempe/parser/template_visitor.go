package parser

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/tniedbala/tempe-go/tempe/api"
	"github.com/tniedbala/tempe-go/tempe/parser/base"
	nd "github.com/tniedbala/tempe-go/tempe/parser/nodes"
	opt "github.com/tniedbala/tempe-go/tempe/options"
)

type TemplateVisitor struct {
	*base.BaseTemplateParserVisitor
}

func (v TemplateVisitor) VisitTemplate(ctx base.ITemplateContext) *nd.TemplateNodesCollection {
	return v.VisitNodes(ctx.AllNode())
}

func (v TemplateVisitor) VisitNodes(nodes []base.INodeContext) *nd.TemplateNodesCollection {
	n := nd.NewTemplateNodesCollection()
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

func (v TemplateVisitor) VisitText(ctx base.ITextContext) *nd.Text {
	return nd.NewText(ctx.GetText())
}

func (v TemplateVisitor) VisitExpression(ctx base.IExpressionContext) *nd.Expression {
	return nd.NewExpression(ctx.Expr().GetText())
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

func (v TemplateVisitor) VisitAssignment(ctx base.IAssignmentContext) *nd.Assignment {
	assignment := nd.NewAssignment()
	for _, v := range ctx.AllVar_() {
		name, expr := v.Name().GetText(), v.Expr().GetText()
		assignment.Append(name, expr)
	}
	openStmt, closeStmt := v.visitEndStatement(ctx)
	assignment.SetWhitespace(openStmt, closeStmt)
	return assignment
}

func (v TemplateVisitor) VisitForLoop(ctx base.IForLoopContext) *nd.ForLoop {
	startFor, endFor := ctx.StartFor(), ctx.EndFor()
	varName := startFor.Name().GetText()
	indexName, index := "", startFor.Index()
	if index != nil {
		indexName = index.GetText()
	}
	expr := startFor.Expr().GetText()
	body := v.VisitNodes(ctx.AllNode())
	forLoop := nd.NewForLoop(indexName, varName, expr, body)
	enterLeading, enterTrailing := v.visitStartStatement(startFor)
	exitLeading, exitTrailing := v.visitEndStatement(endFor)
	forLoop.SetWhitespace(opt.For, enterLeading, enterTrailing)
	forLoop.SetWhitespace(opt.EndFor, exitLeading, exitTrailing)
	return forLoop
}

func (v TemplateVisitor) VisitIfStatement(ctx base.IIfStatementContext) *nd.IfStatement {
	clauses := nd.NewTemplateNodesCollection()
	startIf, endIf := ctx.StartIf(), ctx.EndIf()
	clause := v.VisitIfClause(startIf, nd.If)
	clauses.Append(clause)
	for _, elseIf := range ctx.AllElseIf() {
		clause := v.VisitIfClause(elseIf, nd.ElseIf)
		clauses.Append(clause)
	}
	if elseCtx := ctx.Else_(); elseCtx != nil {
		clause := v.VisitElseClause(elseCtx)
		clauses.Append(clause)
	}
	ifStmt := nd.NewIfStatement(clauses)
	openStmt, closeStmt := v.visitEndStatement(endIf)
	ifStmt.SetWhitespace(openStmt, closeStmt)
	return ifStmt
}

func (v TemplateVisitor) VisitIfClause(ctx ifClauseContext, kind nd.IfClauseKind) *nd.IfClause {
	expr, body := ctx.Expr().GetText(), v.VisitNodes(ctx.AllNode())
	clause := nd.NewIfClause(kind, expr, body)
	openStmt, closeStmt := v.visitStartStatement(ctx)
	clause.SetWhitespace(openStmt, closeStmt)
	return clause
}

func (v TemplateVisitor) VisitElseClause(ctx base.IElseContext) *nd.IfClause {
	body := v.VisitNodes(ctx.AllNode())
	clause := nd.NewIfClause(nd.Else, "", body)
	openStmt, closeStmt := v.visitEndStatement(ctx)
	clause.SetWhitespace(openStmt, closeStmt)
	return clause
}

func (v TemplateVisitor) visitStartStatement(ctx startStatement) (string, string) {
	return ctx.OPEN_STMT().GetText(), ctx.CLOSE_STMT_EXPR().GetText()
}

func (v TemplateVisitor) visitEndStatement(ctx endStatement) (string, string) {
	return ctx.OPEN_STMT().GetText(), ctx.CLOSE_STMT().GetText()
}

func (v TemplateVisitor) VisitErrorNode(ctx antlr.ErrorNode) api.TemplateNode {
	// TODO: improve error handling
	panic(fmt.Sprintf("ERROR: %v", ctx))
}

func (v TemplateVisitor) VisitChildren(ctx antlr.RuleNode) *nd.TemplateNodesCollection {
	nodes := nd.NewTemplateNodesCollection()
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
