package main

import (
	"fmt"
	"sync"
	"time"
)

func scheduleTask(task func(), wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		task()
	}()
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	task1 := func() {
		sum := 0
		for i := 0; i < 100; i++ {
			sum = i * i
		}
		fmt.Println("sum is:", sum)
	}
	task2 := func() {
		time.Sleep(time.Second * 2)
		fmt.Println("waited for 2s")
	}
	task3 := func() {
		time.Sleep(time.Second * 3)
		fmt.Println("waited for 3s")
	}
	tasks := []func(){task1, task2, task3}

	for _, t := range tasks {
		scheduleTask(t, &wg)
	}

	wg.Wait()
	fmt.Println("All tasks completed")
	fmt.Println(time.Since(start))
}
