package calculations

import (
	"math"
)

func CalculateAverage(slice []int) float64 {
	var average float64
	for _, num := range slice {
		average += float64(num)
	}
	lengthOfSlice := float64(len(slice))
	average = math.Round(average / lengthOfSlice)
	return average
}

func CalculateVariance(slice []int) float64 {
	var variance float64
	average := CalculateAverage(slice)
	for _, num := range slice {
		variance += (float64(num) - average) * (float64(num) - average)
	}
	lengthOfSlice := float64(len(slice))
	variance = variance / (lengthOfSlice)
	return variance
}

func CalculateStdDev(slice []int) float64 {
	var stdDev float64
	variance := CalculateVariance(slice)
	stdDev = math.Sqrt(variance)
	return stdDev
}
