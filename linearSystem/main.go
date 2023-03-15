package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewDense(2, 2, []float64{1, 2, 3, -5})
	ia := mat.NewDense(2, 2, nil)
	b := mat.NewDense(2, 1, []float64{5, 4})
	s := mat.NewDense(2, 1, nil)

	err := ia.Inverse(a)
	if err != nil {
		panic(err)
	}

	s.Mul(ia, b)

	fa := mat.Formatted(a, mat.Prefix("    "), mat.Squeeze())
	fb := mat.Formatted(b, mat.Prefix("    "), mat.Squeeze())
	fs := mat.Formatted(s, mat.Prefix("          "), mat.Squeeze())

	fmt.Printf("\nA = %.2v\n\n", fa)
	fmt.Printf("\nB = %.2v\n\n", fb)
	fmt.Printf("\nsolucao = %.2v\n\n", fs)
}
