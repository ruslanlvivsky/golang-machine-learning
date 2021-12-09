# 행렬, 확률, 통계

## 벡터

- 벡터로만 작업하는 경우 float 으로 이루어진 슬라이스에 빠르고 가벼운 연산이 필요하기 때문에 floats 가 좋은 선택이다
- 행렬과 벡터 모두 사용해야하는 경우 벡터/행렬 에 대한 좀 더 넓은 범위의 기능에 접근해야 하기 때문에 mat 을 사용하는 것이 더 나은 선택이다

```go
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
```

## 행렬

```go
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
```

