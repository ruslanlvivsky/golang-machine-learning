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



### 데이터 프레임을 활용해 CSV 데이터 조작하기 

```go
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
```



## JSON

### JSON 파싱과 출력

```go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const citiBikeURL = "https://gbfs.citibikenyc.com/gbfs/en/station_status.json"

type stationData struct {
	LastUpdated int `json:"last_updated"`
	TTL         int `json:"ttl"`

	Data struct {
		Stations []station `json:"stations"`
	} `json:"data"`
}

type station struct {
	ID                     string `json:"station_id"`
	NumBikesAvailable      int    `json:"num_bikes_available"`
	NumEbikesAvailable     int    `json:"num_ebikes_available"`
	LastReported           int    `json:"last_reported"`
	NumDocksAvailable      int    `json:"num_docks_available"`
	EightdHasAvailableKeys bool   `json:"eightd_has_available_keys"`
	StationStatus          string `json:"station_status"`
	IsRenting              int    `json:"is_renting"`
	LegacyID               string `json:"legacy_id"`
	IsInstalled            int    `json:"is_installed"`
	NumDocksDisabled       int    `json:"num_docks_disabled"`
	NumBikesDisabled       int    `json:"num_bikes_disabled"`
	IsReturning            int    `json:"is_returning"`
}

func main() {
	response, err := http.Get(citiBikeURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var sd stationData
	err = json.Unmarshal(body, &sd)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n\n", sd.Data.Stations[0])

	outputData, err := json.Marshal(sd)
	if err != nil {
		log.Fatal(err)
	}

	if err = ioutil.WriteFile("citibike.json", outputData, 0644); err != nil {
		log.Fatal(err)
	}

}

```



## 캐싱

머신 러닝 모델이나 분석 애플리케이션이 계속해서 API 를 요청하는 것이 아닌 메모리나 디스크에 저장하여 API 요청 회수를 줄이고 속도를 높이는 방법

```go
func main() {
	// 메모리 캐싱
	c := cache.New(5*time.Minute, 30*time.Second)
	c.Set("mykey", "myvalue", cache.DefaultExpiration)

	v, found := c.Get("mykey")
	if found {
		fmt.Printf("key: mykey, value: %s\n", v)
	}

	// 디스크 로컬 캐싱
	db, err := bolt.Open("embedded.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	if err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		err = b.Put([]byte("mykey"), []byte("myvalue"))
		return err
	}); err != nil {
		log.Fatal(err)
	}

	if err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		cursor := b.Cursor()
		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
			fmt.Printf("key: %s, value: %s\n", key, value)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}
```





