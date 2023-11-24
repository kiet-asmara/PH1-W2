package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func aritmatika(letter string, n1 string, n2 string) (hasil float64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from:", r)
		}
	}()

	var IsNumber = regexp.MustCompile(`^[0-9]+$`).MatchString
	var IsChoice = regexp.MustCompile(`^[a-d]+$`).MatchString

	if !IsNumber(n1) || !IsNumber(n2) || (letter == "d" && n2 == "0") || !IsChoice(letter) {
		panic("input contains invalid characters")
	}

	i1, _ := strconv.Atoi(n1)
	i2, _ := strconv.Atoi(n2)

	f1, f2 := float64(i1), float64(i2)
	var res float64

	switch letter {
	case "a":
		res = f1 + f2
	case "b":
		res = f1 - f2
	case "c":
		res = f1 * f2
	case "d":
		res = f1 / f2
	}
	return res, nil
}

func main() {
	fmt.Println("pilih operasi aritmatika:")
	fmt.Println("> penjumlahan (+) a")
	fmt.Println("> pengurangan (-) b")
	fmt.Println("> perkalian (*) c")
	fmt.Println("> pembagian (/) d")

	var choice, n1, n2 string
	fmt.Scanln(&choice)
	fmt.Scanln(&n1)
	fmt.Scanln(&n2)
	fmt.Println(aritmatika(choice, n1, n2))

}
