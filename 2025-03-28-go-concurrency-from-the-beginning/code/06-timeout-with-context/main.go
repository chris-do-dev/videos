package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting work...")

	tasks := []string{"Task1", "Task2", "Task3"}

	resultChan := make(chan string)
	defer close(resultChan)

	timeout := time.Duration(2.5 * float32(time.Second))

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for i, task := range tasks {
		go func() {
			process(ctx, task, resultChan, i+1)
		}()
	}

	for range tasks {
		result := <-resultChan
		fmt.Println(result)
	}

	fmt.Println("Finished work.")
}

func process(ctx context.Context, task string, resultChan chan<- string, taskDuration int) {
	fmt.Println("Started", task)

	taskDone := time.After(time.Duration(taskDuration) * time.Second)

	select {
	case <-ctx.Done():
		resultChan <- fmt.Sprint("Timedout ", task)
	case <-taskDone:
		resultChan <- fmt.Sprint("Finished ", task)
	}
}
