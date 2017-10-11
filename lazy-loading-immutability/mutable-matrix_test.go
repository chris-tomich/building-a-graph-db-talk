package main

import "testing"

func TestMutableMatrixMultiplication(t *testing.T) {
	m1 := NewMutableMatrix(
		[][]int{
			{2, 3, 4},
			{1, 0, 0},
		},
	)

	m2 := NewMutableMatrix(
		[][]int{
			{0, 1000},
			{1, 100},
			{0, 10},
		},
	)

	m1.MatrixMultiply(m2)
	m1EqualsM2 := m1.Equals(NewMutableMatrix(
		[][]int{
			{3, 2340},
			{0, 1000},
		},
	))

	if !m1EqualsM2 {
		t.Fail()
	}

	m3 := NewMutableMatrix(
		[][]int{
			{2, 3, 4},
		},
	)

	m4 := NewMutableMatrix(
		[][]int{
			{0, 1000},
			{1, 100},
			{0, 10},
		},
	)

	m3.MatrixMultiply(m4)
	m3EqualsM4 := m3.Equals(NewMutableMatrix(
		[][]int{
			{3, 2340},
		},
	))

	if !m3EqualsM4 {
		t.Fail()
	}

	m5 := NewMutableMatrix(
		[][]int{
			{2, 3, 4},
			{1, 0, 0},
		},
	)

	m6 := NewMutableMatrix(
		[][]int{
			{0},
			{1},
			{0},
		},
	)

	m5.MatrixMultiply(m6)
	m5EqualsM6 := m5.Equals(NewMutableMatrix(
		[][]int{
			{3},
			{0},
		},
	))

	if !m5EqualsM6 {
		t.Fail()
	}
}
