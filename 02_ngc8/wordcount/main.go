package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func countWords(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	count := 0

	for scanner.Scan() {
		text := strings.Fields(scanner.Text())
		count += len(text)
	}

	return count
}

func main() {
	filenames := []string{
		"Files1.txt",
		"Files2.txt",
		"Files3.txt",
	}

	var wg sync.WaitGroup

	wg.Add(len(filenames))

	wordCounts := make([]int, len(filenames))

	for i, filename := range filenames {
		go func(fn string, i int) {
			defer wg.Done()
			count := countWords(fn)
			wordCounts[i] = count
		}(filename, i)
	}

	wg.Wait()

	for i, count := range wordCounts {
		fmt.Printf("%s: %d words\n", filenames[i], count)
	}
}
