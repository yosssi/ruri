package ruri

// result represents a parse result.
type result struct {
	elements []element
}

// HTML generates and returns an HTML source code.
func (r *result) HTML() string {
	return ""
}
