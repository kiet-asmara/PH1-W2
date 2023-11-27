package main

import (
	"fmt"
	"math"
)

func SumSquare() int {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}

	sum *= sum
	return sum
}

func SquareSum() int {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += int(math.Pow(float64(i), 2))
	}
	return sum
}

func main() {
	ch := make(chan int, 2)

	go func() {
		ch <- SumSquare()
	}()
	go func() {
		ch <- SquareSum()
	}()

	res := <-ch
	res -= <-ch

	fmt.Println("Square of sum - sum of squares =", res)

}
