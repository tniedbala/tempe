package parser

import (
	"fmt"
	"strings"
	"unicode/utf8"

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

func strLen(text string) int {
	return utf8.RuneCountInString(text)
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

func alignRight(text string, width int, char string) string {
	charCount := strLen(text)
	if charCount > width {
		return text
	}
	padLeft := strings.Repeat(char, width - charCount)
	return fmt.Sprintf("%s%s", padLeft, text)
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
