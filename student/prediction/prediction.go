package prediction

import (
	"guess-it-1/student/prediction/calculations"
	"math"
)

func PredictRange(slice []int) []int {
	result := []int{}
	mean := int(math.Round(calculations.CalculateAverage(slice)))
	stdDev := int(math.Round(calculations.CalculateStdDev(slice)))
	min := mean - stdDev
	max := mean + stdDev
	result = append(result, min)
	result = append(result, max)

	return result
}
