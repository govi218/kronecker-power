package main

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func setSubMatrix(M *mat.Dense, Sub mat.Matrix, i, j, length int) {
	Mi, Mj := 0, 0
	if i != 0 {
		Mi = i + length - 1
	} else {
		Mi = 0
	}
	if j != 0 {
		Mj = j + length - 1
	} else {
		Mj = 0
	}
	Mrows, _ := M.Dims()

	if Mi+length > Mrows || Mj+length > Mrows {
		return
	}
	for row := Mi; row < Mi+length; row++ {
		for col := Mj; col < Mj+length; col++ {
			(*M).Set(row, col, Sub.At(row-Mi, col-Mj))
		}
	}
}

func kroneckerProduct(A, B mat.Matrix) (C *mat.Dense) {
	aRows, _ := A.Dims()
	bRows, _ := B.Dims()

	C = mat.NewDense(bRows*2, bRows*2, nil)
	temp := mat.NewDense(bRows, bRows, nil)

	for i := 0; i < aRows; i++ {
		for j := 0; j < aRows; j++ {
			temp.Scale(A.At(i, j), B)
			setSubMatrix(C, temp, i, j, bRows)
		}
	}

	return
}

func channelCombine(n int64) (C *mat.Dense) {
	numProducts := int(math.Log2(float64(n)))
	k := make([]float64, 4)
	k[0] = float64(1)
	k[1] = float64(0)
	k[2] = float64(1)
	k[3] = float64(1)

	arikanMatrix := mat.NewDense(2, 2, k)
	C = arikanMatrix
	for i := 0; i < numProducts; i++ {
		C = kroneckerProduct(arikanMatrix, C)
		println("C at i=", i)
		matPrint(C)
	}
	return
}

func main() {
	channelCombine(8)
}
