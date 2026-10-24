package load

import (
	"encoding/csv"
	"os"
	"strconv"
)

func LoadIris() ([][]float64, []string, error) {
	f, err := os.Open("data/iris/iris.data")
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// 5 tokens: 4 features and 1 target
	featureCount := 4

	features := make([][]float64, len(records))
	targets := make([]string, len(records))
	for i, record := range records {
		features[i] = make([]float64, featureCount)
		for j := range featureCount {
			features[i][j], err = strconv.ParseFloat(record[j], 64)
			if err != nil {
				return nil, nil, err
			}
		}
		targets[i] = record[featureCount]
	}

	return features, targets, nil
}
