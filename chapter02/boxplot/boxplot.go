package main

import (
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

	p := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	p.Title.Text = "Box plots"
	p.Y.Label.Text = "Values"

	w := vg.Points(50)

	for idx, colName := range irisDF.Names() {
		if colName != "Species" && colName != "Id" {
			v := make(plotter.Values, irisDF.Nrow())
			for i, floatVal := range irisDF.Col(colName).Float() {
				v[i] = floatVal
			}

			b, err := plotter.NewBoxPlot(w, float64(idx), v)
			if err != nil {
				log.Fatal(err)
			}
			p.Add(b)
		}
	}

	p.NominalX("", "SepalLengthCm", "SepalWidthCm", "PetalLengthCm", "PetalWidthCm")
	if err := p.Save(6*vg.Inch, 8*vg.Inch, "boxplots.png"); err != nil {
		log.Fatal(err)
	}
}
