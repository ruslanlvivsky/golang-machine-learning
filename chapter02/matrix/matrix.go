package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"log"
	"math"
)

func main() {
	data := []float64{1.2, -5.7, -2.4, 7.3}
	a := mat.NewDense(2, 2, data)

	fa := mat.Formatted(a, mat.Prefix(" "))
	fmt.Printf("A = %v\n\n", fa)

	val := a.At(0, 1)
	fmt.Printf("val = %v\n\n", val)
	col := mat.Col(nil, 0, a)
	fmt.Printf("col = %v\n\n", col)
	row := mat.Row(nil, 0, a)
	fmt.Printf("row = %v\n\n", row)

	a.Set(0, 1, 11.2)
	a.SetRow(0, []float64{14.3, -4.2})
	a.SetCol(0, []float64{1.7, -0.3})

	fmt.Printf("A = %v\n\n", fa)

	// 연산
	a = mat.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})
	b := mat.NewDense(3, 3, []float64{8, 9, 10, 1, 4, 2, 9, 0, 2})

	c := mat.NewDense(3, 2, []float64{3, 2, 1, 4, 0, 8})

	d := mat.NewDense(3, 3, nil)
	d.Add(a, b)
	fd := mat.Formatted(d, mat.Prefix(" "))
	fmt.Printf("d = a + b = %0.4v\n\n", fd)

	f := mat.NewDense(3, 2, nil)
	f.Mul(a, c)
	ff := mat.Formatted(f, mat.Prefix(" "))
	fmt.Printf("f = a * c = %0.4v\n\n", ff)

	g := mat.NewDense(3, 3, nil)
	g.Pow(a, 5)
	fg := mat.Formatted(g, mat.Prefix(" "))
	fmt.Printf("g = a^5 = %0.4v\n\n", fg)

	h := mat.NewDense(3, 3, nil)
	sqrt := func(_, _ int, v float64) float64 { return math.Sqrt(v) }
	h.Apply(sqrt, a)
	fh := mat.Formatted(h, mat.Prefix(" "))
	fmt.Printf("h = sqrt(a) = %0.4v\n\n", fh)

	// 전치 행렬
	ft := mat.Formatted(a.T(), mat.Prefix(" "))
	fmt.Printf("a^T = %0.4v\n\n", ft)

	// ad-bc
	deta := mat.Det(a)
	fmt.Printf("det(a) = %0.2v\n\n", deta)

	// 역행렬
	aInverse := mat.NewDense(3, 3, nil)
	if err := aInverse.Inverse(a); err != nil {
		log.Fatal(err)
	}
	fi := mat.Formatted(aInverse, mat.Prefix(" "))
	fmt.Printf("a^-1 = %0.2v\n\n", fi)
}
