package options

import (
	"fmt"
	"iter"
	"strings"
)

var syntaxElements = map[string]SyntaxElement{}

func init() {
	for elem := range SyntaxElements() {
		syntaxElements[elem.String()] = elem
	}
}

type SyntaxElement int

const (
	Assignment SyntaxElement = iota
	Comment
	For
	EndFor
	If
	ElseIf
	Else
	EndIf
)

func SyntaxElements() iter.Seq[SyntaxElement] {
	return func(yield func(SyntaxElement) bool) {
		for elem := Assignment; elem <= EndIf; elem++ {
			if !yield(elem) {
				return
			}
		}
	}
}

func (s SyntaxElement) String() string {
	switch s {
	case Assignment:
		return "assignment"
	case Comment:
		return "comment"
	case For:
		return "for"
	case EndFor:
		return "endfor"
	case If:
		return "if"
	case ElseIf:
		return "elseif"
	case Else:
		return "else"
	case EndIf:
		return "endif"
	default:
		return "undefined"
	}
}

func (s SyntaxElement) FromString(str string) (SyntaxElement, bool) {
	elem, ok := syntaxElements[strings.ToLower(str)]
	return elem, ok
}

func (s SyntaxElement) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, s)), nil
}
