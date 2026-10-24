package main

import (
	"fmt"

	"github.com/michaelzhan1/clustering/internal/load"
	"github.com/mpraski/clusters"
)

func main() {
	irisData, irisTargets, err := load.LoadIris()

	c, err := clusters.DBSCAN(8, 1.5, 4, clusters.EuclideanDistance)
	if err != nil {
		panic(err)
	}

	err = c.Learn(irisData)
	if err != nil {
		panic(err)
	}

	preds := c.Guesses()
	fmt.Println(preds)
	fmt.Println(irisTargets)
}
