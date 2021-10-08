package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../data/Iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reader := csv.NewReader(f)

	// 라인당 필드 수를 모르는 경우
	// FieldsPerRecord를 음수로 설정해 각 행의 필드 수를 얻을 수 있다
	reader.FieldsPerRecord = -1

	/*
		한 번에 읽어들이기
		rawCSVData, err := reader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
	*/

	//	한 행씩 읽어들이기
	var rawCSVData [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// 예상치 못한 필드 때문에 오류가 발생하는 경우
		if err != nil {
			log.Fatal(err)
			continue
		}

		rawCSVData = append(rawCSVData, record)
	}

	fmt.Println(rawCSVData)
}
