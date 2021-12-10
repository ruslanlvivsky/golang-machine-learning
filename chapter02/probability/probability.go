package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

func main() {
	// 카이 제곱 검정 통계량
	// https://ko.wikipedia.org/wiki/%EC%B9%B4%EC%9D%B4%EC%A0%9C%EA%B3%B1_%EA%B2%80%EC%A0%95
	observed := []float64{
		260, 135, 105,
	}

	totalObserved := 500.0

	expected := []float64{
		totalObserved * 0.6,
		totalObserved * 0.25,
		totalObserved * 0.15,
	}

	chiSquare := stat.ChiSquare(observed, expected)
	fmt.Printf("\n카이-제곱 검정 통계량: %0.2f\n", chiSquare)

	chiDist := distuv.ChiSquared{
		K:   2.0,
		Src: nil,
	}

	pValue := chiDist.Prob(chiSquare)
	fmt.Printf("\np-value: %0.4f\n", pValue)
	// 편차의 결과가 우연히 발생했을 확률 0.01%
	// 5% 의 임계 값을 사용하는 경우 귀무 가설을 배제하고 대립 가설 채택해야 한다
	// 유의확률
	// https://ko.wikipedia.org/wiki/%EC%9C%A0%EC%9D%98_%ED%99%95%EB%A5%A0
}
