package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting work...")

	tasks := []string{"Task1", "Task2", "Task3"}

	resultChan := make(chan string)
	defer close(resultChan)

	for _, task := range tasks {
		go func() {
			process(task, resultChan)
		}()
	}

	for range tasks {
		result := <-resultChan
		fmt.Println(result)
	}

	fmt.Println("All work completed!")
}

func process(task string, result chan<- string) {
	fmt.Println("Started", task)
	time.Sleep(time.Second)
	result <- fmt.Sprint("Finished ", task)
}
