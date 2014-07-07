package ruri

import "strings"

// line represents a line of codes.
type line string

// empty returns true if the line is empty.
func (ln line) empty() bool {
	return strings.TrimSpace(string(ln)) == ""
}

// newLine creates and returns a line.
func newLine(s string) line {
	return line(s)
}
