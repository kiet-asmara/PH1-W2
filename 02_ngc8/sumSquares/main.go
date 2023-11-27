package main

import (
	"fmt"
	"math"
	"sync"
)

var mu sync.Mutex

func calculateSumOfSquares(n int, sum *int, wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	*sum += int(math.Pow(float64(n), 2))
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0

	for _, n := range numbers {
		wg.Add(1)
		go calculateSumOfSquares(n, &sum, &wg)
	}

	wg.Wait()

	fmt.Println("Sum of squares:", sum)
}
