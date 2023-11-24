package main

import (
	"bufio"
	"fmt"
	"os"
)

func processFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := 0
	for scanner.Scan() {
		lines++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lines, nil
}

func main() {
	filename := "data.txt"
	lines, err := processFile(filename)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Printf("Number of lines in %s: %d \n", filename, lines)
}
