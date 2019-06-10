package matrix

import "fmt"

type Dims struct {
	x, y int
}
type MatrixOperations interface {
	Multiply(Matrix)
	ScalarMultiply(int)
	Transpose()
	Sum()
	Add(Matrix)
	Invert()
	SubMat(Dims)
	IsSquare()
	IsOrthogonal()
	Trace()
	Determinant()
	SetRow(int, []int)
	SetCol(int, []int)
}

type Matrix struct {
	data [][]int
	dim  Dims
}

func (d *Dims) SetDims(x, y int) {
	d.x = x
	d.y = y
}

// func ()

func New(dim Dims) Matrix {
	var mat Matrix
	mat.dim = dim
	mat.data = make([][]int, dim.y)
	for i := range mat.data {
		mat.data[i] = make([]int, dim.x)
	}
	return mat
}
func (m Matrix) Multiply(inMat Matrix) Matrix {
	var dim Dims
	// var outChannel chan int
	outMat := New(dim)
	if m.dim.x != inMat.dim.y {
		fmt.Println("matrix dimensions do not support multiplication")
		return outMat
	}
	dim.y = m.dim.y
	dim.x = inMat.dim.x
	for i, v := range m.data {
		for j := range v {
			outMat.data[i][j] += m.data[i][j] * inMat.data[j][i]
		}
		// go func() {
		// 	for _, w := range v {
		// 		fmt.Println(w)
		// 	}
		// }()
	}
	return outMat
}
func (m *Matrix) SetRow(index int, row []int) bool {
	if len(m.data[index]) == len(row) {
		m.data[index] = row
		return true
	}
	return false
}
func (m *Matrix) SetCol(index int, col []int) bool {
	if len(m.data) == len(col) {

		for i := range m.data[index] {
			m.data[index][i] = col[i]
		}
		return true
	}
	return false
}
func (m Matrix) PrettyPrint() {
	for _, v := range m.data {
		fmt.Println(v)
	}
}
func (m Matrix) PrintSpecific(x, y int) {
	fmt.Println(m.data[y][x])
}
