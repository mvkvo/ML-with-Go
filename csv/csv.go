package csv

import (
	"encoding/csv"
	"mime/multipart"
	"strconv"
	"web/utils"
)

const (
	SQUARE_FTM = 0.09290304
)

func Read_CSV(file multipart.File) ([][]string, error) {
	data := csv.NewReader(file)
	records, err := data.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func Create_slice(records [][]string) []utils.Inputs {
	length := len(records[0])
	features := make([]utils.Inputs, length)
	for i, record := range records {
		for j, r := range record {
			if i == 0 {
				features[j].Header = r
			} else {
				features[j].Values = append(features[j].Values, r)
			}
		}
	}
	return features
}

func String_to_Float(value string) (float_value float64, err error) {
	float_value, err = strconv.ParseFloat(value, 64)
	return float_value, err
}
