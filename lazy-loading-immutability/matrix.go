package lazyloadingimmutability

const MatrixWidth int = 810
const MatrixHeight int = 810

type Matrix interface {
	Width() int
	Height() int
	Get(int, int) int
	Equals(Matrix) bool
	Add(Matrix) (Matrix, error)
	Subtract(Matrix) (Matrix, error)
	ScalarMultiply(s int) Matrix
	Transpose() Matrix
	MatrixMultiply(Matrix) (Matrix, error)
}
