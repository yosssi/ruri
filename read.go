package ruri

import "io/ioutil"

// readFile reads the file and returns the slice of bytes.
func readFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}
