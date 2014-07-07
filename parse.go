package ruri

import (
	"fmt"
	"strings"
)

// parse parses the string and returns the result.
func parse(s string) (*result, error) {
	var elements []element

	// Split the string to lines.
	lines := strings.Split(formatLF(s), "\n")

	i, l := 0, len(lines)

	// Ignore the last empty line.
	if lines[l-1] == "" {
		l--
	}

	for i < l {
		// Fetch the line.
		ln := newLine(lines[i])
		i++

		// Ignore the empty line.
		if ln.empty() {
			continue
		}

		fmt.Println(ln)
	}

	return newResult(elements), nil
}

// formatLF replaces the line feed codes with LF and returns the result.
func formatLF(s string) string {
	return strings.Replace(strings.Replace(s, "\r\n", "\n", -1), "\r", "\n", -1)
}
