package options

import (
	"fmt"
	"iter"
	"strings"
)

var wsRules = map[string]WhitespaceRule{}

func init() {
	for rule := range WhitespaceRules() {
		wsRules[rule.String()] = rule
	}
}

type WhitespaceRule int

const (
	Keep WhitespaceRule = iota
	Trim
	TrimSpace
	TrimLine
)

func WhitespaceRules() iter.Seq[WhitespaceRule] {
	return func(yield func(WhitespaceRule) bool) {
		for rule := Keep; rule <= TrimLine; rule++ {
			if !yield(rule) {
				return
			}
		}
	}
}

func (w *WhitespaceRule) UnmarshalJSON(data []byte) error {
	key := strings.ToLower(strings.Trim(string(data), `"`))
	rule, ok := wsRules[key]
	if !ok {
		return fmt.Errorf(`invalid syntax element: "%s"`, key)
	}
	*w = rule
	return nil
}

func (w WhitespaceRule) String() string {
	switch w {
	case Keep:
		return "keep"
	case Trim:
		return "trim"
	case TrimSpace:
		return "trimspace"
	case TrimLine:
		return "trimline"
	default:
		return "undefined"
	}
}

func (w WhitespaceRule) FromString(str string) (WhitespaceRule, bool) {
	rule, ok := wsRules[strings.ToLower(str)]
	return rule, ok
}

func (w WhitespaceRule) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, w)), nil
}
