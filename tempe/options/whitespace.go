package opt

import (
	"fmt"
	"strings"

	"github.com/tniedbala/tempe-go/tempe/api"
)

var WhitespacePositions = []WhitespacePosition{
	UpperLeft,
	UpperRight,
	InnerLeft,
	InnerRight,
	LowerLeft,
	LowerRight,
}

type WhitespaceStatementType int

const (
	Assignment WhitespaceStatementType = iota
	Comment
)

func (w WhitespaceStatementType) String() string {
	switch w {
	case Assignment:
		return "Assignment"
	case Comment:
		return "Comment"
	default:
		return "WhitespaceStatementType"
	}
}

type WhitespacePosition int

const (
	UpperLeft WhitespacePosition = iota
	InnerLeft
	LowerLeft
	UpperRight
	InnerRight
	LowerRight
)

func (w WhitespacePosition) String() string {
	switch w {
	case UpperLeft:
		return "UpperLeft"
	case UpperRight:
		return "UpperRight"
	case InnerLeft:
		return "InnerLeft"
	case InnerRight:
		return "InnerRight"
	case LowerLeft:
		return "LowerLeft"
	case LowerRight:
		return "LowerRight"
	default:
		return "WhitespacePosition"
	}
}

type WhitespaceRule int

const (
	Keep WhitespaceRule = iota
	Trim
	TrimLine
	TrimSpace
)

func (w WhitespaceRule) String() string {
	switch w {
	case Keep:
		return "Keep"
	case Trim:
		return "Trim"
	case TrimLine:
		return "TrimLine"
	case TrimSpace:
		return "TrimSpace"
	default:
		return "WhitespaceRule"
	}
}

type whitespaceTypes map[WhitespaceStatementType]WhitespaceRule

type whitespacePositions map[WhitespacePosition]WhitespaceRule

type whitespaceOptions struct {
	types     whitespaceTypes
	positions whitespacePositions
}

func (w whitespaceOptions) String() string {
	var b strings.Builder
	b.WriteString("{")
	indent := strings.Repeat(" ", 8)
	for k, v := range w.types {
		b.WriteString(fmt.Sprintf("\n%s%s: %s", indent, k, v))
	}
	for _, position := range WhitespacePositions {
		rule := w.positions[position]
		b.WriteString(fmt.Sprintf("\n%s%s: %s", indent, position, rule))
	}
	b.WriteString("\n    }")
	return b.String()
}

func defaultWhitespaceOptions() whitespaceOptions {
	return whitespaceOptions{
		types: whitespaceTypes{
			Assignment: Trim,
			Comment:    Trim,
		},
		positions: whitespacePositions{
			UpperLeft:  TrimSpace,
			InnerLeft:  TrimSpace,
			LowerLeft:  Trim,
			UpperRight: Trim,
			InnerRight: Trim,
			LowerRight: TrimSpace,
		},
	}
}

func WhitespaceOption[T WhitespacePosition | WhitespaceStatementType](whitespace T) api.Option {
	return Option(func(opts *Options) (value any, err error) {
		var ok bool
		switch ws := any(whitespace).(type) {
		case WhitespaceStatementType:
			value, ok = opts.whitespace.types[ws]
		case WhitespacePosition:
			value, ok = opts.whitespace.positions[ws]
		default:
			value, ok = nil, false
		}
		if !ok {
			return value, fmt.Errorf("option not found: %v", whitespace)
		}
		return value, nil
	})
}

func Whitespace[T WhitespacePosition | WhitespaceStatementType](rule WhitespaceRule, whitespace ...T) api.Option {
	return Option(func(opts *Options) (value any, err error) {
		// if no whitespace args provided, apply rule to all whitespace
		if len(whitespace) == 0 {
			for ws := range opts.whitespace.types {
				opts.whitespace.types[ws] = rule
			}
			for ws := range opts.whitespace.positions {
				opts.whitespace.positions[ws] = rule
			}
			return nil, nil
		}
		// apply rule to target whitespace
		for _, value := range whitespace {
			switch ws := any(value).(type) {
			case WhitespaceStatementType:
				opts.whitespace.types[ws] = rule
			case WhitespacePosition:
				opts.whitespace.positions[ws] = rule
			default:
				return nil, fmt.Errorf("invalid value passed to SetWhitespace: %v", value)
			}
		}
		return nil, nil
	})
}
