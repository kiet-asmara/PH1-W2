package main

import (
	"fmt"
	"sync"
)

func fibonacci(wg *sync.WaitGroup, res *int, n int) {
	defer wg.Done()

	var fibo []int

	// inpuf first two numbers
	fibo = append(fibo, 0)
	fibo = append(fibo, 1)

	for i := 0; i <= n; i++ {
		if i == 0 || i == 1 {

		} else {
			fibo = append(fibo, 0)
			fibo[i] = fibo[i-1] + fibo[i-2]
		}
		// fmt.Printf("bilangan fibonacci ke %d adalah %d\n", i+1, fibo[i])
	}
	*res = fibo[n-1]
	// fmt.Printf("bilangan fibonacci ke %d adalah %d\n", n, fibo[n-1])
}

// no go routine
// func main() {
// 	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 	fmt.Println("Counting numbers")

// 	for _, n := range numbers {
// 		fibonacci(n)
// 	}

// 	fmt.Println("Finished counting")

// }

// with go routine
// func main() {
// 	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 	for _, n := range numbers {
// 		go fibonacci(n)
// 	}

// 	fmt.Println("Counting numbers")

// 	time.Sleep(1 * time.Second)

// 	fmt.Println("Finished counting")

// }

// with confinement
func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := make([]int, len(numbers))
	var wg sync.WaitGroup

	fmt.Println("input numbers:", numbers)

	for i, n := range numbers {
		wg.Add(1)
		go fibonacci(&wg, &result[i], n)
	}

	wg.Wait()

	fmt.Println("fibonacci result:", result)

}
