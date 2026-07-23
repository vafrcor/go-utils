package math

import "math"

func RoundToDecimal(x float64, places int) float64 {
	pow := math.Pow(10, float64(places))
	return math.Round(x*pow) / pow
}

func Percentage(count int64, total int64, precission int) float64 {
	if total == 0 {
		return 0.0
	}
	return RoundToDecimal(float64(count)/float64(total)*100, precission)
}

func CeilInt(v float64) int {
	return int(math.Ceil(v))
}
