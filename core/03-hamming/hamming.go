// Package hamming includes a solution for the "Hamming" problem in the Go track on https://exercism.io.
package hamming

import "errors"

// Distance calculates the Hamming distance between two sequences.
func Distance(a, b string) (int, error) {
	var aRune = []rune(a)
	var bRune = []rune(b)

	if len(aRune) != len(bRune) {
		return 0, errors.New("the two sequences have different lengths")
	}

	diff := 0

	for i := 0; i < len(aRune); i++ {
		if aRune[i] != bRune[i] {
			diff++
		}
	}

	return diff, nil
}
