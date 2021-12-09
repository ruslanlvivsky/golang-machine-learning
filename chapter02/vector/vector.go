package main

import (
	"fmt"
	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

func main() {
	var myVector []float64

	myVector = append(myVector, 11.0)
	myVector = append(myVector, 5.2)

	fmt.Println(myVector)

	vectorA := []float64{11.0, 5.2, -1.3}
	vectorB := []float64{-7.2, 4.2, 5.1}

	dotProduct := floats.Dot(vectorA, vectorB)
	fmt.Printf("백터 내적: %0.2f\n", dotProduct)

	floats.Scale(1.5, vectorA)
	fmt.Printf("A에 1.5 곱셈: %v\n", vectorA)

	normB := floats.Norm(vectorB, 2)
	fmt.Printf("벡터 B의 놈/길이: %0.2f\n", normB)

	vectorC := mat.NewVecDense(3, []float64{11.0, 5.2, -1.3})
	vectorD := mat.NewVecDense(3, []float64{-7.2, 4.2, 5.1})

	dotProduct = mat.Dot(vectorC, vectorD)
	fmt.Printf("백터 내적: %0.2f\n", dotProduct)

	vectorC.ScaleVec(1.5, vectorC)
	fmt.Printf("C에 1.5 곱셈: %v\n", vectorC)

	normD := blas64.Nrm2(vectorD.RawVector())
	fmt.Printf("벡터 B의 놈/길이: %0.2f\n", normD)
}
