package main

import (
	"fmt"
	"math/rand"
)

func Factorial(ch <-chan int) chan int {
	num := <-ch
	total := 1
	for i := 2; i <= num; i++ {
		total *= i
	}

	chOut := make(chan int, 1)

	defer close(chOut)

	chOut <- total
	return chOut
}

func main() {
	n := rand.Intn(99) + 1
	fmt.Println(n)

	chIn := make(chan int, 1)
	chIn <- n

	defer close(chIn)

	fmt.Println(<-Factorial(chIn))

}
