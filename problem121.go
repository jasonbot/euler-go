package main

import (
	"fmt"
)

func ProbabilityAtN(iteration uint64) float64 {
	return 1. / (float64(iteration) + 2.)
}

func ProbabilityArrayForIterationCount(rounds uint64) []float64 {
	probability_array := make([]float64, rounds)
	for index := uint64(0); index < rounds; index++ {
		probability_array[index] = ProbabilityAtN(index)
	}

	return probability_array
}

func main() {
	fmt.Printf("%v\n", ProbabilityArrayForIterationCount(15))
}
