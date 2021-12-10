package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"log"
	"os"
)

func main() {
	irisFile, err := os.Open("../../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	irisDF := dataframe.ReadCSV(irisFile)

	for _, colName := range irisDF.Names() {
		if colName != "Species" {
			v := make(plotter.Values, irisDF.Nrow())
			for i, floatVal := range irisDF.Col(colName).Float() {
				v[i] = floatVal
			}

			p := plot.New()
			p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

			h, err := plotter.NewHist(v, 16)
			if err != nil {
				log.Fatal(err)
			}

			// 막대 그래프 정규화
			h.Normalize(1)

			p.Add(h)

			if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
				log.Fatal(err)
			}
		}
	}
}
