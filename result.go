package ruri

// result represents a parse result.
type result struct {
	elements []element
}

// HTML generates and returns an HTML source code.
func (r *result) HTML() string {
	return ""
}

// newResult creates and returns a result.
func newResult(elements []element) *result {
	return &result{
		elements: elements,
	}
}
