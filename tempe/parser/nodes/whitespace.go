package nodes

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	opt "github.com/tniedbala/tempe-go/tempe/options"
)

var (
	regexLeadingWs  *regexp.Regexp
	regexTrailingWs *regexp.Regexp
	regexLine       *regexp.Regexp
	regexSpace      *regexp.Regexp
)

func init() {
	regexLeadingWs = regexp.MustCompile(`^\r?\n[ \t]*`)
	regexTrailingWs = regexp.MustCompile(`[ \t]*\r?\n$`)
	regexLine = regexp.MustCompile(`\r?\n`)
	regexSpace = regexp.MustCompile(`[ \t]+`)
}

type WhitespaceFlag int

const (
	None WhitespaceFlag = iota
	Keep
	Trim
)

func (w WhitespaceFlag) String() string {
	switch w {
	case None:
		return "none"
	case Keep:
		return "keep"
	case Trim:
		return "trim"
	default:
		return "undefined"
	}
}

func (w WhitespaceFlag) MarshalJSON() ([]byte, error) {
	if w == None {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, w)), nil
}

type Whitespace struct {
	Position opt.WhitespacePosition `json:"position"`
	Flag     WhitespaceFlag         `json:"flag"`
	Text     string                 `json:"text"`
}

func DefaultWhitespace(position opt.WhitespacePosition) Whitespace {
	return Whitespace{position, None, ""}
}

func NewWhitespace(position opt.WhitespacePosition, text string) Whitespace {
	text, flag := extractWhitespace(position, text), getWhitespaceFlag(text)
	return Whitespace{position, flag, text}
}

func (s Whitespace) Render(rule opt.WhitespaceRule, w io.StringWriter) error {
	if _, err := w.WriteString(s.ToString(rule)); err != nil {
		return err
	}
	return nil
}

func (s Whitespace) ToString(rule opt.WhitespaceRule) string {
	if s.Flag == Keep {
		return s.Text
	} else if s.Flag == Trim {
		return ""
	}
	switch rule {
	case opt.Keep:
		return s.Text
	case opt.Trim:
		return ""
	case opt.TrimLine:
		return stripLine(s.Text)
	case opt.TrimSpace:
		return stripSpace(s.Text)
	default:
		return s.Text
	}
}

func stripLine(text string) string {
	return regexLine.ReplaceAllString(text, "")
}

func stripSpace(text string) string {
	return regexSpace.ReplaceAllString(text, "")
}

func getWhitespaceFlag(text string) WhitespaceFlag {
	if strings.Contains(text, "+") {
		return Keep
	} else if strings.Contains(text, "-") {
		return Trim
	} else {
		return None
	}
}

func extractWhitespace(position opt.WhitespacePosition, text string) string {
	if position == opt.Leading {
		return regexLeadingWs.FindString(text)
	} else {
		return regexTrailingWs.FindString(text)
	}
}
