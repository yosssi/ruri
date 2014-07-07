package ruri

// ParseFile parses the file and returns the HTML source code.
func ParseFile(path string, opts *Options) (string, error) {
	b, err := readFile(path)
	if err != nil {
		return "", err
	}

	r, err := parse(string(b))
	if err != nil {
		return "", err
	}

	return r.HTML(), nil
}
