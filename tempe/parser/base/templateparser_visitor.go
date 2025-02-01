// Code generated from /workspace/grammar/TemplateParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package base // TemplateParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by TemplateParser.
type TemplateParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TemplateParser#template.
	VisitTemplate(ctx *TemplateContext) interface{}

	// Visit a parse tree produced by TemplateParser#node.
	VisitNode(ctx *NodeContext) interface{}

	// Visit a parse tree produced by TemplateParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by TemplateParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by TemplateParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by TemplateParser#var.
	VisitVar(ctx *VarContext) interface{}

	// Visit a parse tree produced by TemplateParser#forLoop.
	VisitForLoop(ctx *ForLoopContext) interface{}

	// Visit a parse tree produced by TemplateParser#startFor.
	VisitStartFor(ctx *StartForContext) interface{}

	// Visit a parse tree produced by TemplateParser#endFor.
	VisitEndFor(ctx *EndForContext) interface{}

	// Visit a parse tree produced by TemplateParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by TemplateParser#startIf.
	VisitStartIf(ctx *StartIfContext) interface{}

	// Visit a parse tree produced by TemplateParser#elseIf.
	VisitElseIf(ctx *ElseIfContext) interface{}

	// Visit a parse tree produced by TemplateParser#else.
	VisitElse(ctx *ElseContext) interface{}

	// Visit a parse tree produced by TemplateParser#endIf.
	VisitEndIf(ctx *EndIfContext) interface{}

	// Visit a parse tree produced by TemplateParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by TemplateParser#index.
	VisitIndex(ctx *IndexContext) interface{}

	// Visit a parse tree produced by TemplateParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by TemplateParser#text.
	VisitText(ctx *TextContext) interface{}
}
