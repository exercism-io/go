// Package isogram includes a solution for the "Isogram" problem in the Go track on https://exercism.io.
package isogram

import "unicode"

// IsIsogram computes if a given word or phase is an isogram.
func IsIsogram(s string) bool {
	var cache [26]bool
	asRunes := []rune(s)

	for i := 0; i < len(asRunes); i++ {
		if !unicode.IsLetter(asRunes[i]) {
			continue
		}

		lower := unicode.ToLower(asRunes[i])
		index := int(lower - 'a')

		if cache[index] {
			return false
		}

		cache[index] = true
	}

	return true
}
