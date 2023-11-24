package main

import (
	"fmt"
	"regexp"
)

func anagramCheck(w1 string, w2 string) (an bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("input tidak valid")
		}
	}()

	var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

	if !IsLetter(w1) || !IsLetter(w2) || len(w1) > 10 || len(w2) > 10 {
		panic("input contains invalid characters")
	}

	// anagram logic
	lenS := len(w1)
	lenT := len(w2)
	res := false

	if lenS != lenT {
		res = false
	}

	anagramMap := make(map[string]int)

	for i := 0; i < lenS; i++ {
		anagramMap[string(w1[i])]++
	}

	for i := 0; i < lenT; i++ {
		anagramMap[string(w2[i])]--
	}

	for i := 0; i < lenS; i++ {
		if anagramMap[string(w1[i])] != 0 {
			res = false
		}
	}

	return res, nil
}

func main() {
	// fmt.Println(anagramCheck("ttht%rt", "tthtrt"))
	res, err := anagramCheck("test", "tse5")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

}
