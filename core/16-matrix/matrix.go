// Package matrix includes a solution for the "Matrix" problem in the Go track on https://exercism.io.
package matrix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Matrix represents a matrix
type Matrix struct {
	data  [][]int
	nRows int
	nCols int
}

// New creates a 2D matrix from string
func New(input string) (*Matrix, error) {
	m := new(Matrix)
	rows := strings.Split(input, "\n")
	m.nRows = len(rows)
	m.data = make([][]int, m.nRows)

	for r, row := range rows {
		elements := strings.Split(strings.TrimSpace(row), " ")
		if len(elements) == 0 {
			return nil, fmt.Errorf("row %d is empty", r)
		}
		if m.nCols != 0 && len(elements) != m.nCols {
			return nil, errors.New("Uneven rows")
		}
		if m.nCols == 0 {
			m.nCols = len(elements)
		}
		m.data[r] = make([]int, len(elements))
		for e, element := range elements {
			i, err := strconv.Atoi(element)
			if err != nil {
				return nil, err
			}
			m.data[r][e] = i
		}
	}
	return m, nil
}

// Rows returns the rows of the matrix
func (m *Matrix) Rows() [][]int {
	cpy := make([][]int, m.nRows)
	for i := range cpy {
		cpy[i] = make([]int, m.nCols)
		copy(cpy[i], m.data[i])
	}

	return cpy
}

// Cols returns the columns of a matrix
func (m *Matrix) Cols() [][]int {
	cols := make([][]int, m.nCols)
	for i := 0; i < m.nCols; i++ {
		cols[i] = make([]int, m.nRows)
	}
	for r, row := range m.data {
		for c := range row {
			cols[c][r] = row[c]
		}
	}
	return cols
}

// Set sets a value in the matrix given the row and column.
func (m *Matrix) Set(r, c, val int) bool {
	if r < 0 || r >= m.nRows || c < 0 || c >= m.nCols {
		return false
	}
	m.data[r][c] = val
	return true
}
