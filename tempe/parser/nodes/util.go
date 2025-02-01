package nodes

import (
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/tniedbala/tempe-go/tempe/parser/base"
)

type IfClauseKind int

const (
	If IfClauseKind = iota
	ElseIf
	Else
)

func (k IfClauseKind) String() string {
	switch k {
	case If:
		return "if"
	case ElseIf:
		return "elseif"
	case Else:
		return "else"
	default:
		return "IfClause"
	}
}

const TextAlignWidth int = 60
const FormatStringMaxLength int = 60

func formatText(text string) string {
	str := replaceNewline(text)
	if len(str) > FormatStringMaxLength {
		str = str[:FormatStringMaxLength-3] + "..."
	}
	return str
}

func replaceNewline(text string) string {
	text = strings.Replace(text, "\r", `\r`, -1)
	return strings.Replace(text, "\n", `\n`, -1)
}

func replaceSpace(text string) string {
	text = strings.Replace(text, "\t", `\t`, -1)
	return strings.Replace(text, " ", "␣", -1)
}

func replaceWhitespace(text string) string {
	text = replaceNewline(text)
	return replaceSpace(text)
}

type startStatement interface {
	OPEN_STMT() antlr.TerminalNode
	CLOSE_STMT_EXPR() antlr.TerminalNode
}

type endStatement interface {
	OPEN_STMT() antlr.TerminalNode
	CLOSE_STMT() antlr.TerminalNode
}

type ifClauseContext interface {
	startStatement
	Expr() base.IExprContext
	AllNode() []base.INodeContext
}
