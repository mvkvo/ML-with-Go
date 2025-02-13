package utils

import (
	"math"
	"strconv"
)

type Dataset struct {
	Inputs []Inputs `json:"inputs"`
	Err    error    `json:"err"`
}

type Inputs struct {
	Header string   `json:"header"`
	Values []string `json:"values"`
}

func FloatToString(input []float64) string {
	var s string
	for _, i := range input {
		s += strconv.FormatFloat(i, 'f', 2, 64)
	}
	return s
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

/*
	 func CleanData(input_x, input_y []float64) (cleaned_inputs []InputValues) {
		qa, _ := stats.QuartileOutliers(input_x)
		qp, _ := stats.QuartileOutliers(input_y)

		var x_outlines = make(stats.Float64Data, 0)
		var y_outlines = make(stats.Float64Data, 0)

		x_outlines = append(append(qa.Mild, qa.Extreme...), x_outlines...)
		y_outlines = append(append(qp.Mild, qp.Extreme...), y_outlines...)

		for _, o := range x_outlines {
			if index := Index(input_x, o); index > 0 {
				input_x[index] = calc.Mean(input_x)
				input_y[index] = calc.Mean(input_y)
			}
		}

		for _, o := range y_outlines {
			if index := Index(input_y, o); index > 0 {
				input_x[index] = calc.Mean(input_x)
				input_y[index] = calc.Mean(input_y)
			}
		}

		return
	}
*/
func Index(s []float64, v float64) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

func String_to_float_slice(slice []string) []float64 {
	slice_float := make([]float64, len(slice))
	for i, s := range slice {
		s_float, _ := strconv.ParseFloat(s, 64)
		slice_float[i] = s_float
	}
	return slice_float
}
