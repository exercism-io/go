// Package raindrops includes a solution for the "Raindrops" problem in the Go track on https://exercism.io.
package raindrops

import "strconv"

// Convert converts a number to string, the contents of which depend on the number's factors.
func Convert(n int) string {
	s := ""

	if n%3 == 0 {
		s += "Pling"
	}

	if n%5 == 0 {
		s += "Plang"
	}

	if n%7 == 0 {
		s += "Plong"
	}

	if s == "" {
		s = strconv.Itoa(n)
	}

	return s
}
