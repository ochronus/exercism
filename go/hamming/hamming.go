package hamming

import "errors"

// Distance calculates the Hamming distance of two strings and returns it
// returns an error if the strings don't have equal length
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strings must be of the same length")
	}
	var hammingDiff = 0
	for i, x := range a {
		if string(x) != string(b[i]) {
			hammingDiff++
		}
	}
	return hammingDiff, nil
}
