package nodes

import (
	"io"

	"github.com/tniedbala/tempe-go/tempe/api"
	"github.com/tniedbala/tempe-go/tempe/opt"
)

type StatementPosition int

const (
	Upper StatementPosition = iota
	Inner
	Lower
)

type whitespaceMap map[opt.WhitespacePosition]Whitespace

type Statement struct {
	whitespace whitespaceMap
}

func NewStatement() Statement {
	whitespace := whitespaceMap{}
	for _, position := range opt.WhitespacePositions {
		whitespace[position] = DefaultWhitespace()
	} 
	return Statement{whitespace}
}

func (s *Statement) RenderWhitespace(position StatementPosition, opts api.Options, w io.StringWriter) error {
	var leftWs, rightWs string
	switch position {
	case Upper:
		leftWs, rightWs = s.GetWhitespace(opts, opt.UpperLeft), s.GetWhitespace(opts, opt.UpperRight)
	case Inner:
		leftWs, rightWs = s.GetWhitespace(opts, opt.InnerLeft), s.GetWhitespace(opts, opt.InnerRight)
	case Lower:
		leftWs, rightWs = s.GetWhitespace(opts, opt.LowerLeft), s.GetWhitespace(opts, opt.LowerRight)
	}
	_, err := w.WriteString(leftWs + rightWs)
	return err
}

func (s *Statement) SetWhitespace(position StatementPosition, openStmt, closeStmt string) {
	switch position {
	case Upper:
		s.setWs(opt.UpperLeft, openStmt)
		s.setWs(opt.UpperRight, closeStmt)
	case Inner:
		s.setWs(opt.InnerLeft, openStmt)
		s.setWs(opt.InnerRight, closeStmt)
	case Lower:
		s.setWs(opt.LowerLeft, openStmt)
		s.setWs(opt.LowerRight, closeStmt)
	}
}

func (s *Statement) setWs(position opt.WhitespacePosition, text string) {
	s.whitespace[position] = NewWhitespace(position, text)
}

func (s *Statement) GetWhitespace(opts api.Options, position opt.WhitespacePosition) string {
	ws, ok := s.whitespace[position]
	if !ok {
		return ""
	}
	// if +/- is provided in statement then this overrides whitespace options
	if ws.flag == Keep {
		return ws.text
	} else if ws.flag == Strip {
		return ""
	}
	// otherwise return whitespace per whitespace options
	rule, ok := opts.Get(opt.WhitespaceOption(position))
	if !ok {
		return ws.text
	}
	switch rule {
	case opt.Trim:
		return ""
	case opt.TrimLine:
		return stripLine(ws.text)
	case opt.TrimSpace:
		return stripSpace(ws.text)
	default:
		return ws.text
	}
}
