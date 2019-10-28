// Package grains includes a solution for the "Grains" problem in the Go track on https://exercism.io.
package grains

import "errors"

const boardSize = 64

// Square computes the number of grains in the i-th square on the chessboard.
func Square(i int) (uint64, error) {
	if i <= 0 || i > 64 {
		return 0, errors.New("input must a positive number between 1 and 64")
	}

	return 1 << uint(i-1), nil
}

// Total computes the total number of grains in the chessboard.
func Total() uint64 {
	return 1<<boardSize - 1
}
