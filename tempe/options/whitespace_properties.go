package options

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"maps"
	"slices"
)

//go:embed whitespace_properties.json
var DefaultWhitespacePropertiesJson string

func DefaultWhitespaceProperties() WhitespaceProperties {
	var properties WhitespaceProperties
	if err := json.Unmarshal([]byte(DefaultWhitespacePropertiesJson), &properties); err == nil {
		return properties
	}
	properties = WhitespaceProperties{}
	for elem := range SyntaxElements() {
		for position := range WhitespacePositions() {
			properties[elem][position] = Trim
		}
	}
	for _, elem := range []SyntaxElement{For, If} {
		properties[elem][Leading] = TrimSpace
	}
	for _, elem := range []SyntaxElement{EndFor, EndIf} {
		properties[elem][Trailing] = TrimSpace
	}
	return properties
}

type WhitespaceRuleMap map[WhitespacePosition]WhitespaceRule

type WhitespaceProperties map[SyntaxElement]WhitespaceRuleMap

func (w WhitespaceProperties) MarshalJSON() ([]byte, error) {
	obj := map[string]map[string]string{}
	for _, elem := range slices.Sorted(maps.Keys(w)) {
		elemKey := elem.String()
		obj[elemKey] = map[string]string{}
		for _, position := range slices.Sorted(maps.Keys(w[elem])) {
			rule := w[elem][position]
			obj[elemKey][position.String()] = rule.String()
		}
	}
	return json.Marshal(obj)
}

func (w *WhitespaceProperties) UnmarshalJSON(data []byte) error {
	var obj map[string]any
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}
	var ok bool
	properties := WhitespaceProperties{}
	for key, value := range obj {
		var elem SyntaxElement
		elem, ok = elem.FromString(key)
		if !ok {
			return fmt.Errorf(`unrecognized syntax element: "%s"`, key)
		}
		values, ok := value.(map[string]any)
		if !ok {
			return fmt.Errorf(`invalid whitespace properties: "%s": %s`, key, value)
		}
		properties[elem] = WhitespaceRuleMap{}
		for k, v := range values {
			var position WhitespacePosition
			position, ok = position.FromString(k)
			if !ok {
				return fmt.Errorf(`unrecognized whitespace position: "%s"`, k)
			}
			ruleStr, ok := v.(string)
			if !ok {
				return fmt.Errorf(`invalid whitespace rule: %s`, v)
			}
			var rule WhitespaceRule
			rule, ok = rule.FromString(ruleStr)
			if !ok {
				return fmt.Errorf(`unrecognized whitespace rule: "%s"`, ruleStr)
			}
			properties[elem][position] = rule
		}
	}
	*w = properties
	return nil
}