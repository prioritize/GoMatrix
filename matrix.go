package matrix

type Dims struct {
	x, y int
}
type MatrixInfo struct {
	dim Dims
}
type Matrix interface {
	Multiply()
	ScalarMultiply()
	Transpose()
	Sum()
	Add()
	Invert()
	SubMat()
	IsSquare()
	IsOrthogonal()
	Trace()
	Determinant()
}

func New(dim Dims) Matrix {

}
