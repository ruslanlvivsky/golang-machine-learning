package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
	"log"
	"os"
)

func main() {
	irisFile, err := os.Open("../../../data/Iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	irisDF := dataframe.ReadCSV(irisFile)

	sepalLength := irisDF.Col("SepalLengthCm").Float()

	minVal := floats.Min(sepalLength)
	maxVal := floats.Max(sepalLength)
	rangeVal := maxVal - minVal
	varianceVal := stat.Variance(sepalLength, nil)
	stdDevVal := stat.StdDev(sepalLength, nil)

	inds := make([]int, len(sepalLength))
	floats.Argsort(sepalLength, inds)

	quant25 := stat.Quantile(0.25, stat.Empirical, sepalLength, nil)
	quant50 := stat.Quantile(0.50, stat.Empirical, sepalLength, nil)
	quant75 := stat.Quantile(0.75, stat.Empirical, sepalLength, nil)

	fmt.Printf("\n꽃받침 길이 요약 통계:\n")
	fmt.Printf("최댓값: %0.2f\n", maxVal)
	fmt.Printf("최솟값: %0.2f\n", minVal)
	fmt.Printf("범위값: %0.2f\n", rangeVal)
	fmt.Printf("분산: %0.2f\n", varianceVal)
	fmt.Printf("표준 편차: %0.2f\n", stdDevVal)
	fmt.Printf("25분위: %0.2f\n", quant25)
	fmt.Printf("50분위: %0.2f\n", quant50)
	fmt.Printf("75분위: %0.2f\n", quant75)

}
