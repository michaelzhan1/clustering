package main

import (
	"github.com/michaelzhan1/clustering/internal/load"
	"github.com/michaelzhan1/clustering/internal/plot"
	"github.com/mpraski/clusters"
)

func main() {
	irisData, _, err := load.LoadIris()

	c, err := clusters.DBSCAN(8, 1.5, 4, clusters.EuclideanDistance)
	if err != nil {
		panic(err)
	}

	err = c.Learn(irisData)
	if err != nil {
		panic(err)
	}

	preds := c.Guesses()
	// fmt.Println(preds)
	// fmt.Println(irisTargets)

	plot.Scatter("out.png", irisData, preds, "test")
}
