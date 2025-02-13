package calc

import "math"

func StdDev(values []float64, mean float64) float64 {
	var sum float64
	for _, val := range values {
		sum += math.Pow(DevOfMean(float64(val), mean), 2)
	}
	return math.Sqrt(sum)
}

func DevOfMean(val, mean float64) float64 {
	return val - mean
}

func Mean(values []float64) float64 {
	sum := Sum(values)
	return sum / float64(len(values))
}

func Sum(values []float64) float64 {
	total := 0.0
	for _, value := range values {
		total += value
	}
	return total
}
