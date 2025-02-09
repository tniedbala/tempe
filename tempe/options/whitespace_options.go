package options

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tniedbala/tempe-go/tempe/api"
)

func Whitespace(property string, rules ...WhitespaceRule) api.Option {
	if len(rules) == 0 {
		return getWhitespace(property)
	}
	return setWhitespace(property, rules...)
}

func WhitespaceOption(elem SyntaxElement, position WhitespacePosition) api.Option {
	return Option(func(options *Options) (any, error) {
		rules, ok := options.whitespace.properties[elem]
		if !ok {
			return nil, fmt.Errorf("whitespace rules not found for %s", elem)
		}
		rule, ok := rules[position]
		if !ok {
			return nil, fmt.Errorf("whitespace rule not found for %s.%s", elem, position)
		}
		return rule, nil
	})
}

func getWhitespace(property string) api.Option {
	return Option(func(options *Options) (any, error) {
		rule, ok := options.whitespace.GetRule(property)
		if !ok {
			return nil, fmt.Errorf(`invalid whitespace property: "%s"`, property)
		}
		return rule, nil
	})
}

func setWhitespace(property string, rules ...WhitespaceRule) api.Option {
	return Option(func(options *Options) (any, error) {
		if len(rules) != 1 {
			return nil, fmt.Errorf("only a single whitespace rule may be provided; got %d: %s", len(rules), rules)
		}
		return nil, options.whitespace.SetRule(property, rules[0])
	})
}

type WhitespaceOptions struct {
	properties WhitespaceProperties
}

func DefaultWhitespaceOptions() *WhitespaceOptions {
	return &WhitespaceOptions{properties: DefaultWhitespaceProperties()}
}

func (w *WhitespaceOptions) GetRule(property string) (rule WhitespaceRule, ok bool) {
	properties := strings.Split(property, ".")
	if len(properties) != 2 {
		return rule, false
	}
	var elem SyntaxElement
	var position WhitespacePosition
	elem, ok = elem.FromString(properties[0])
	if !ok {
		return rule, ok
	}
	position, ok = position.FromString(properties[1])
	if !ok {
		return rule, ok
	}
	rule, ok = w.properties[elem][position]
	return rule, ok
}

func (w *WhitespaceOptions) SetRule(property string, rule WhitespaceRule) error {
	properties := strings.Split(property, ".")
	switch len(properties) {
	case 0:
		w.update("", "", rule)
	case 1:
		return w.update(properties[0], "", rule)
	case 2:
		return w.update(properties[0], properties[1], rule)
	default:
		return fmt.Errorf("error")
	}
	return nil
}

func (w *WhitespaceOptions) update(syntax string, position string, rule WhitespaceRule) error {
	if w.isWildcard(syntax) {
		w.updateAll(rule)
		return nil
	}
	var elem SyntaxElement
	var pos WhitespacePosition
	var ok bool
	elem, ok = elem.FromString(syntax)
	if !ok {
		return fmt.Errorf(`unrecognized whitespace syntax "%s"`, syntax)
	}
	if w.isWildcard(position) {
		w.updateSyntax(elem, rule)
		return nil
	}
	pos, ok = pos.FromString(position)
	if !ok {
		return fmt.Errorf(`unrecognized whitespace position "%s"`, position)
	}
	w.properties[elem][pos] = rule
	return nil
}

func (w *WhitespaceOptions) updateAll(rule WhitespaceRule) {
	for elem := range w.properties {
		w.updateSyntax(elem, rule)
	}
}

func (w *WhitespaceOptions) updateSyntax(elem SyntaxElement, rule WhitespaceRule) {
	for position := range w.properties[elem] {
		w.properties[elem][position] = rule
	}
}

func (w *WhitespaceOptions) isWildcard(str string) bool {
	str = strings.TrimSpace(str)
	return str == "" || str == "*"
}

func (w *WhitespaceOptions) String() string {
	b, err := json.MarshalIndent(w.properties, "", "  ")
	if err != nil {
		return fmt.Sprintf(`Whitespace{error: "%s"}`, err)
	}
	return fmt.Sprintf("Whitespace%s", string(b))
}

func (w *WhitespaceOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(w.properties)
}

func (w *WhitespaceOptions) UnmarshalJSON(data []byte) error {
	var properties WhitespaceProperties
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	// ensure any missing rules are set to default value
	for elem, rules := range DefaultWhitespaceProperties() {
		_, ok := properties[elem]
		if !ok {
			properties[elem] = WhitespaceRuleMap{}
		}
		for position, rule := range rules {
			_, ok = properties[elem][position]
			if !ok {
				properties[elem][position] = rule
			}
		}
	}
	*w = WhitespaceOptions{properties}
	return nil
}