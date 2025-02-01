package parser

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/antlr4-go/antlr/v4"
	"github.com/tniedbala/tempe-go/tempe/parser/base"
)

func strLen(text string) int {
	return utf8.RuneCountInString(text)
}

func alignRight(text string, width int, char string) string {
	charCount := strLen(text)
	if charCount > width {
		return text
	}
	padLeft := strings.Repeat(char, width-charCount)
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
