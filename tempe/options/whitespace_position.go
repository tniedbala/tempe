package options

import (
	"fmt"
	"iter"
	"strings"
)

var wsPositions = map[string]WhitespacePosition{}

func init() {
	for position := range WhitespacePositions() {
		wsPositions[position.String()] = position
	}
}

type WhitespacePosition int

const (
	Leading WhitespacePosition = iota
	Trailing
)

func WhitespacePositions() iter.Seq[WhitespacePosition] {
	return func(yield func(WhitespacePosition) bool) {
		for pos := Leading; pos <= Trailing; pos++ {
			if !yield(pos) {
				return
			}
		}
	}
}

func (w *WhitespacePosition) UnmarshalJSON(data []byte) error {
	key := strings.ToLower(strings.Trim(string(data), `"`))
	position, ok := wsPositions[key]
	if !ok {
		return fmt.Errorf(`invalid syntax element: "%s"`, key)
	}
	*w = position
	return nil
}

func (w WhitespacePosition) String() string {
	switch w {
	case Leading:
		return "leading"
	case Trailing:
		return "trailing"
	default:
		return "undefined"
	}
}

func (w WhitespacePosition) FromString(str string) (WhitespacePosition, bool) {
	pos, ok := wsPositions[strings.ToLower(str)]
	return pos, ok
}

func (w WhitespacePosition) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, w)), nil
}
