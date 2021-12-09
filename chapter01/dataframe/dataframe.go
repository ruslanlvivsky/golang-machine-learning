package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"log"
	"os"
)

func main() {
	irisFile, err := os.Open("../data/Iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	irisDF := dataframe.ReadCSV(irisFile)
	fmt.Println(irisDF)

	// 데이터 프레임의 필터를 생성
	filter := dataframe.F{
		Colname:    "Species",
		Comparator: "==",
		Comparando: "Iris-versicolor",
	}

	versicolorDF := irisDF.Filter(filter)
	err = versicolorDF.Error()
	if err != nil {
		log.Fatal(err)
	}

	versicolorDF = irisDF.Filter(filter).Select([]string{
		"SepalWidthCm",
		"Species",
	})

	// 처음 세 개의 결과만 선택
	versicolorDF = irisDF.Filter(filter).Select([]string{
		"SepalWidthCm",
		"Species",
	}).Subset([]int{0, 1, 2})

	fmt.Println(versicolorDF)
}
