package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/montanaflynn/stats"
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

	meanVal := stat.Mean(sepalLength, nil)

	modeVal, modeCount := stat.Mode(sepalLength, nil)

	medianVal, err := stats.Median(sepalLength)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n꽃받침 길이 요약 통계:\n")
	fmt.Printf("평균값: %0.2f\n", meanVal)
	fmt.Printf("최빈값: %0.2f\n", modeVal)
	fmt.Printf("최빈값 개수: %d\n", int(modeCount))
	fmt.Printf("최빈값: %0.2f\n", medianVal)
}
