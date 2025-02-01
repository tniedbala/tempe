package parser

import (
	"regexp"
	"strings"

	"github.com/tniedbala/tempe-go/tempe/opt"
)

var (
	regexWsLeft  *regexp.Regexp
	regexWsRight *regexp.Regexp
	regexLine    *regexp.Regexp
	regexSpace   *regexp.Regexp
)

func init() {
	regexWsLeft = regexp.MustCompile(`^\r?\n[ \t]*`)
	regexWsRight = regexp.MustCompile(`[ \t]*\r?\n$`)
	regexLine = regexp.MustCompile(`\r?\n`)
	regexSpace = regexp.MustCompile(`[ \t]+`)
}

type WhitespaceFlag int

const (
	None WhitespaceFlag = iota
	Keep
	Strip
)

type Whitespace struct {
	flag WhitespaceFlag
	text string
}

func DefaultWhitespace() Whitespace {
	return Whitespace{None, ""}
}

func NewWhitespace(position opt.WhitespacePosition, text string) Whitespace {
	text, flag := extractWhitespace(position, text), getWhitespaceFlag(text)
	return Whitespace{flag, text}
}

func getWhitespaceFlag(text string) WhitespaceFlag {
	if strings.Contains(text, "+") {
		return Keep
	} else if strings.Contains(text, "-") {
		return Strip
	} else {
		return None
	}
}

func extractWhitespace(position opt.WhitespacePosition, text string) string {
	switch position {
	case opt.UpperLeft, opt.InnerLeft, opt.LowerLeft:
		return regexWsLeft.FindString(text)
	case opt.UpperRight, opt.InnerRight, opt.LowerRight:
		return regexWsRight.FindString(text)
	default:
		return ""
	}
}

func stripLine(text string) string {
	return regexLine.ReplaceAllString(text, "")
}

func stripSpace(text string) string {
	return regexSpace.ReplaceAllString(text, "")
}
