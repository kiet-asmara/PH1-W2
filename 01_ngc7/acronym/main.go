package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func acronym(t string) {
	array := regexp.MustCompile("[\\-\\,\\.\\ \\s]+").Split(t, -1)

	res := ""
	for _, word := range array {
		res += string(word[0])
	}

	fmt.Printf("%s - %s\n", t, strings.ToUpper(res))
}

// with go routine
func main() {
	words := []string{
		"Thank George it's Friday!",
		"John Doe",
		"As Soon as Possible",
		"Liquid-crystal Display",
	}

	for _, w := range words {
		go acronym(w)
	}

	fmt.Println("Making acronyms...")

	time.Sleep(1 * time.Second)

	fmt.Println("Finished task.")
}

// without go routine
// func main() {
// 	words := []string{
// 		"Thank George it's Friday!",
// 		"John Doe",
// 		"As Soon as Possible",
// 		"Liquid-crystal Display",
// 	}

// 	fmt.Println("Making acronyms...")

// 	for _, w := range words {
// 		acronym(w)
// 	}

// 	fmt.Println("Finished task.")
// }
