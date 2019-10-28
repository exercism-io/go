//Package luhn includes a solution for the "Luhn" problem in the Go track on https://exercism.io.
package luhn

import "unicode"

// Valid computes if a string is valid per the Luhn formula.
func Valid(s string) bool {
	asRunes := []rune(s)
	n := len(asRunes)
	nDigits := 0
	sum := 0

	for i := n - 1; i >= 0; i-- {
		if unicode.IsSpace(asRunes[i]) {
			continue
		}

		nDigits++

		if asRunes[i] < '0' || asRunes[i] > '9' {
			// Invalid characters
			return false
		}

		d := int(asRunes[i] - '0')

		if nDigits%2 == 0 {
			d *= 2

			if d > 9 {
				d -= 9
			}
		}

		sum += d
	}

	return sum%10 == 0 && nDigits > 1
}
