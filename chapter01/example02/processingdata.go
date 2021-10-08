package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type CSVRecord struct {
	SepalLengthCm float64
	SepalWidthCm  float64
	PetalLengthCm float64
	PetalWidthCm  float64
	Species       string
	ParseError    error
}

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

	//	한 행씩 읽어들이기
	var csvData []CSVRecord
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		record = record[1:]
		var csvRecord CSVRecord
		for idx, value := range record {
			if idx == 4 {
				if value == "" {
					log.Printf("Unexpected type in column %d\n", idx)
					csvRecord.ParseError = fmt.Errorf("Empty string value")
					break
				}

				csvRecord.Species = value
				continue
			}

			var floatValue float64
			// 레코드 값이 float으로 읽히지 않는 경우
			// 로그에 기록하고 구문 분석 처리 루프를 중단한다
			if floatValue, err = strconv.ParseFloat(value, 64); err != nil {
				log.Printf("Unexpected type in column %d\n", idx)
				csvRecord.ParseError = fmt.Errorf("Could not parse float")
				break
			}

			switch idx {
			case 0:
				csvRecord.SepalLengthCm = floatValue
			case 1:
				csvRecord.SepalWidthCm = floatValue
			case 2:
				csvRecord.PetalLengthCm = floatValue
			case 3:
				csvRecord.PetalWidthCm = floatValue
			}
		}

		if csvRecord.ParseError == nil {
			csvData = append(csvData, csvRecord)
		}
	}

	fmt.Println(csvData)
}
