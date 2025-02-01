package parser

import (
	"fmt"
)

type SyntaxError struct {
	line    int
	column  int
	message string
}

func NewSyntaxError(line, column int, message string) SyntaxError {
	return SyntaxError{
		line:    line,
		column:  column,
		message: message,
	}
}

func (e SyntaxError) Error() string {
	return fmt.Sprintf("Parser Error at line %d, column %d: %s", e.line, e.column, e.message)
}
