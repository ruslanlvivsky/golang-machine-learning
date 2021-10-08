# 데이터 수집 및 구성

## 장점

데이터 수집, 해석(파싱), 구성 측면에서 무결성을 높은 수준으로 유지하는 기회를 제공한다.

1. 예상되는 유형을 확인 하고 작업을 수행하기
   - 동적 타입을 지원하는 언어에 반해 명시적으로 예상되는 타입으로 데이터 구분을 분석하고 오류를 처리할 수 있다

2. 데이터 입출력 표준화 및 단순화 하기
   - 특정 유형의 데이터를 다루는 다양한 서드파티를 사용하지만 stdlib 을 중심으로 사용해 입출력을 표준화한다

3. 데이터 버전 관리하기
   - 코드와 데이터의 버전을 관리하여 일관된 결과를 만든다

## CSV 파일

### CSV 데이터 읽기

```go
	f, err := os.Open("./data/Iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reader := csv.NewReader(f)
```

### 예상하지 못한 필드 처리하기

```go
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
```

### 예상하지 못한 타입 처리하기

```go
			// 붓꽃 타입 확인
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
```



