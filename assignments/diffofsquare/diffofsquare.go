package main

import (
	"fmt"
	"math"
)

func sumSqr(maxNum float64) float64 {
	var i, sum float64
	for i = 1; i <= maxNum; i++ {
		sum += math.Pow(i, 2)
	}
	return sum
}
func sqrSum(maxNum float64) float64 {
	var i, sum float64
	for i = 1; i <= maxNum; i++ {
		sum += i
	}
	return math.Pow(sum, 2)
}
func main() {
	diff := sqrSum(10) - sumSqr(10)
	fmt.Printf("%f - %f = %f", sqrSum(10), sumSqr(10), diff)
}
