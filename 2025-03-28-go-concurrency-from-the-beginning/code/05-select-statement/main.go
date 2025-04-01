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

	for i, task := range tasks {
		go func() {
			process(task, resultChan, i+1)
		}()
	}

	for range tasks {
		result := <-resultChan
		fmt.Println(result)
	}

	fmt.Println("Finished work.")
}

func process(task string, results chan<- string, taskDuration int) {
	fmt.Println("Started", task)

	timeout := time.After(time.Duration(2.5 * float32(time.Second)))
	taskDone := time.After(time.Duration(taskDuration) * time.Second)

	select {
	case <-timeout:
		results <- fmt.Sprint("Timedout ", task)
	case <-taskDone:
		results <- fmt.Sprint("Finished ", task)
	}
}
