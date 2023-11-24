package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

func score(w string, wg *sync.WaitGroup) int {
	score := 0
	for _, l := range w {
		wg.Add(1)

		go func(l rune) {
			defer wg.Done()
			mu.Lock()
			switch l {
			case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
				score += 1
			case 'd', 'g':
				score += 2
			case 'b', 'c', 'm', 'p':
				score += 3
			case 'f', 'h', 'v', 'w', 'y':
				score += 4
			case 'k':
				score += 5
			case 'j', 'x':
				score += 8
			case 'q', 'z':
				score += 10
			}
			mu.Unlock()
		}(l)
	}
	wg.Wait()
	return score
}

func main() {
	var wg sync.WaitGroup

	input := "exampleword"

	res := score(input, &wg)

	fmt.Println(res)

}
