package parser

import (
	"fmt"
	"strings"
	"unicode/utf8"
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
