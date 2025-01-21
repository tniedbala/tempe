// Code generated from /workspace/grammar/TemplateParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package base // TemplateParser
import "github.com/antlr4-go/antlr/v4"

type BaseTemplateParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTemplateParserVisitor) VisitTemplate(ctx *TemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTemplateParserVisitor) VisitNode(ctx *NodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTemplateParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTemplateParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTemplateParserVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTemplateParserVisitor) VisitVar(ctx *VarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTemplateParserVisitor) VisitForLoop(ctx *ForLoopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTemplateParserVisitor) VisitStartFor(ctx *StartForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTemplateParserVisitor) VisitEndFor(ctx *EndForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTemplateParserVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTemplateParserVisitor) VisitStartIf(ctx *StartIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTemplateParserVisitor) VisitElseIf(ctx *ElseIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTemplateParserVisitor) VisitElse(ctx *ElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTemplateParserVisitor) VisitEndIf(ctx *EndIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTemplateParserVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTemplateParserVisitor) VisitIndex(ctx *IndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTemplateParserVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTemplateParserVisitor) VisitText(ctx *TextContext) interface{} {
	return v.VisitChildren(ctx)
}
