package main

import "errors"

// MutableMatrix is a matrix with mutating operations.
type MutableMatrix struct {
	matrix [][]int
}

// NewMutableMatrix creates a new matrix with the given initial values.
func NewMutableMatrix(matrix [][]int) *MutableMatrix {
	return &MutableMatrix{matrix: matrix}
}

// NewEmptyMutableMatrix createas a new empty matrix with the given dimensions.
func NewEmptyMutableMatrix(width int, height int) (*MutableMatrix, error) {
	if width == 0 || height == 0 {
		return nil, errors.New("width and height must both be non-zero")
	}

	m := &MutableMatrix{
		matrix: make([][]int, height),
	}

	for i := 0; i < len(m.matrix); i++ {
		m.matrix[i] = make([]int, width)
	}

	return m, nil
}

// Width returns the number of columns in the matrix.
func (m *MutableMatrix) Width() int {
	return len(m.matrix[0])
}

// Height returns the number of rows in the matrix.
func (m *MutableMatrix) Height() int {
	return len(m.matrix)
}

// Get returns the integer at the provided coordinates.
func (m *MutableMatrix) Get(row int, col int) int {
	return m.matrix[row][col]
}

// Equals will compare a matrix against this matrix and return if they are equal.
func (m *MutableMatrix) Equals(m2 Matrix) bool {
	if m.Height() != m2.Height() {
		return false
	}

	if m.Width() != m2.Width() {
		return false
	}

	for r := 0; r < len(m.matrix); r++ {
		for c := 0; c < len(m.matrix[r]); c++ {
			if m.matrix[r][c] != m2.Get(r, c) {
				return false
			}
		}
	}

	return true
}

// Add will add the values of a matrix to this matrix.
func (m *MutableMatrix) Add(m2 Matrix) (Matrix, error) {
	if m.Height() != m2.Height() {
		return nil, errors.New("width of both matrices are not the same")
	}

	if m.Width() != m2.Width() {
		return nil, errors.New("height of both matrices are not the same")
	}

	for r := 0; r < len(m.matrix); r++ {
		for c := 0; c < len(m.matrix[r]); c++ {
			m.matrix[r][c] = m.matrix[r][c] + m2.Get(r ,c)
		}
	}

	return m, nil
}

// Subtract will subtract the values of a matrix from this matrix.
func (m *MutableMatrix) Subtract(m2 Matrix) (Matrix, error) {
	if m.Height() != m2.Height() {
		return nil, errors.New("width of both matrices are not the same")
	}

	if m.Width() != m2.Width() {
		return nil, errors.New("height of both matrices are not the same")
	}

	for r := 0; r < len(m.matrix); r++ {
		for c := 0; c < len(m.matrix[r]); c++ {
			m.matrix[r][c] = m.matrix[r][c] - m2.Get(r, c)
		}
	}

	return m, nil
}

// ScalarMultiply will multiply this matrix by a given scalar value.
func (m *MutableMatrix) ScalarMultiply(s int) Matrix {
	for r := 0; r < len(m.matrix); r++ {
		for c := 0; c < len(m.matrix[r]); c++ {
			m.matrix[r][c] = m.matrix[r][c] * s
		}
	}

	return m
}

// Transpose will transpose this matrix.
func (m *MutableMatrix) Transpose() Matrix {
	t := make([][]int, len(m.matrix[0]))

	for rt := 0; rt < len(t); rt++ {
		for ct := 0; ct < len(m.matrix); ct++ {
			if t[rt] == nil {
				t[rt] = make([]int, len(m.matrix))
			}
			t[rt][ct] = m.matrix[ct][rt]
		}
	}

	m.matrix = t

	return m
}

// MatrixMultiply will multiple the given matrix against this matrix.
func (m *MutableMatrix) MatrixMultiply(m2 Matrix) (Matrix, error) {
	if m.Width() != m2.Height() {
		return nil, errors.New("the dimensions of the matrices are incompatible, try transposing one first")
	}

	n := make([][]int, m.Height())

	for i := 0; i < m.Height(); i++ {
		n[i] = make([]int, m2.Width())
	}

	for rm := 0; rm < m.Height(); rm++ {
		for cm2 := 0; cm2 < m2.Width(); cm2++ {
			product := 0
			for cm := 0; cm < m.Width(); cm++ {
				product = product + m.matrix[rm][cm]*m2.Get(cm, cm2)
			}
			n[rm][cm2] = product
		}
	}

	m.matrix = n

	return m, nil
}
