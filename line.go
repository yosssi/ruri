package ruri

import "strings"

const (
	unicodeSpace = 32

	indentTop = 0
)

// line represents a line of codes.
type line struct {
	no     int
	s      string
	indent int
}

// empty returns true if the line is empty.
func (ln line) empty() bool {
	return strings.TrimSpace(ln.s) == ""
}

// topIndent returns true if the line's indent is the top indent level.
func (ln line) topIndent() bool {
	return ln.indent == indentTop
}

// newLine creates and returns a line.
func newLine(no int, s string) *line {
	return &line{
		no:     no,
		s:      s,
		indent: indent(s),
	}
}

// indent returns the line's indent.
func indent(s string) int {
	var i int

	for _, b := range s {
		if b != unicodeSpace {
			break
		}
		i++
	}

	return i / 2
}
