package simpleimmutability

import "testing"

func TestImmutableMatrixMultiplication(t *testing.T) {
	m1 := New(
		[][]int{
			{2, 3, 4},
			{1, 0, 0},
		},
	)

	m2 := New(
		[][]int{
			{0, 1000},
			{1, 100},
			{0, 10},
		},
	)

	m3, _ := m1.MatrixMultiply(m2)
	isCorrect := m3.Equals(New(
		[][]int{
			{3, 2340},
			{0, 1000},
		},
	))

	if !isCorrect {
		t.Fail()
	}

	m4 := New(
		[][]int{
			{2, 3, 4},
		},
	)

	m5 := New(
		[][]int{
			{0, 1000},
			{1, 100},
			{0, 10},
		},
	)

	m6, _ := m4.MatrixMultiply(m5)
	isCorrect = m6.Equals(New(
		[][]int{
			{3, 2340},
		},
	))

	if !isCorrect {
		t.Fail()
	}

	m7 := New(
		[][]int{
			{2, 3, 4},
			{1, 0, 0},
		},
	)

	m8 := New(
		[][]int{
			{0},
			{1},
			{0},
		},
	)

	m9, _ := m7.MatrixMultiply(m8)
	isCorrect = m9.Equals(New(
		[][]int{
			{3},
			{0},
		},
	))

	if !isCorrect {
		t.Fail()
	}
}
