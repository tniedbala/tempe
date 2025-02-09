package nodes

import (
	"fmt"
	"io"

	"github.com/tniedbala/tempe-go/tempe/api"
	opt "github.com/tniedbala/tempe-go/tempe/options"
)

type WhitespaceSyntax struct {
	Syntax   opt.SyntaxElement `json:"syntax"`
	Leading  Whitespace        `json:"leading"`
	Trailing Whitespace        `json:"trailing"`
}

func NewWhitespaceSyntax(elem opt.SyntaxElement) WhitespaceSyntax {
	return WhitespaceSyntax{
		Syntax:   elem,
		Leading:  DefaultWhitespace(opt.Leading),
		Trailing: DefaultWhitespace(opt.Trailing),
	}
}

func (s WhitespaceSyntax) Render(options api.Options, w io.StringWriter) error {
	for _, ws := range []Whitespace{s.Leading, s.Trailing} {
		rule := s.GetWhitespaceRule(ws.Position, options)
		if err := ws.Render(rule, w); err != nil {
			return err
		}
	}
	return nil
}

func (s WhitespaceSyntax) GetWhitespaceRule(position opt.WhitespacePosition, options api.Options) opt.WhitespaceRule {
	defaultRule := opt.Keep
	value, ok := options.Get(opt.WhitespaceOption(s.Syntax, position))
	if !ok {
		return defaultRule
	}
	rule, ok := value.(opt.WhitespaceRule)
	if !ok {
		return defaultRule
	}
	return rule
}

func (s WhitespaceSyntax) String() string {
	return fmt.Sprintf(`WhitespaceSyntax{%s}`, s.Syntax)
}
