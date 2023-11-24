package main

import (
	"fmt"
	"math"
)

func calculateSumOfSquares(numbers []int) int {

	sum := 0
	for _, n := range numbers {
		sum += int(math.Pow(float64(n), 2))
	}

	return sum
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(calculateSumOfSquares(numbers))
}
