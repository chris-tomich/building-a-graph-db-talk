package simpleimmutability

import (
	"testing"
	"fmt"
)

func TestMutableMatrixMultiplication(t *testing.T) {
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

	m1.MatrixMultiply(m2)
	m1EqualsM2 := m1.Equals(New(
		[][]int{
			{3, 2340},
			{0, 1000},
		},
	))

	if !m1EqualsM2 {
		fmt.Println("here 1")
		t.Fail()
	}

	m3 := New(
		[][]int{
			{2, 3, 4},
		},
	)

	m4 := New(
		[][]int{
			{0, 1000},
			{1, 100},
			{0, 10},
		},
	)

	m3.MatrixMultiply(m4)
	m3EqualsM4 := m3.Equals(New(
		[][]int{
			{3, 2340},
		},
	))

	if !m3EqualsM4 {
		fmt.Println("here 2")
		t.Fail()
	}

	m5 := New(
		[][]int{
			{2, 3, 4},
			{1, 0, 0},
		},
	)

	m6 := New(
		[][]int{
			{0},
			{1},
			{0},
		},
	)

	m5.MatrixMultiply(m6)
	m5EqualsM6 := m5.Equals(New(
		[][]int{
			{3},
			{0},
		},
	))

	if !m5EqualsM6 {
		fmt.Println(m5)
		t.Fail()
	}
}
