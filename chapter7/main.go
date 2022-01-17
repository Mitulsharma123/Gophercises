package main

import "fmt"
import "calci"

/*
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}
*/

func main() {
	list := []float64{98, 93, 77, 82, 83}
	
	avg := calci.Average(list)
	min := calci.Minimum(list)
	max := calci.Maximum(list)

	fmt.Println("Average:", avg)
	fmt.Println("Minimum", min)
	fmt.Println("Maximum", max)
}
