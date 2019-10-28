// Package twofer includes a solution for the "Two Fer" problem in the Go track on https://exercism.io
package twofer

import "fmt"

// ShareWith returns "One for X, one for me.", where X is the input or "you" (if the input is empty).
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
