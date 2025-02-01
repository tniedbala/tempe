package parser

import (
	"strings"
)

const FormatStringMaxLength int = 60

func formatText(text string) string {
	str := strings.Replace(text, "\n", `\n`, -1)
	if len(str) > FormatStringMaxLength {
		str = str[:FormatStringMaxLength-3] + "..."
	}
	return str
}
