package main

import "fmt"

func performTask(taskName string) {
	fmt.Printf("Task: %s has started.\n", taskName)
	defer fmt.Printf("Task %s has completed\n", taskName)
}

// func main() {
// 	performTask("Task 1")
// 	performTask("Task 2")
// 	performTask("Task 3")
// }
