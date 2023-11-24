package main

import "fmt"

func fizzBuzz(in chan int) {
	sum := 0
	oddCount := 0
	evenCount := 0
	for i := 1; i <= 100; i++ {
		num := <-in
		if num%15 == 0 {
			fmt.Println(num, "FizzBuzz")
		} else if num%5 == 0 {
			fmt.Println(num, "Buzz")
		} else if num%3 == 0 {
			fmt.Println(num, "Fizz")
		} else {
			fmt.Println(num)
		}
		sum += num
		if num%2 != 0 {
			oddCount += 1
		} else {
			evenCount += 1
		}
	}
	fmt.Println("jumlah tambah semua angka:", sum)
	fmt.Println("total jumlah angka ganjil:", oddCount)
	fmt.Println("total jumlah angka genap:", evenCount)

}

func main() {
	inChannel := make(chan int, 100)

	for i := 1; i <= 100; i++ {
		inChannel <- i
	}

	fizzBuzz(inChannel)
}
