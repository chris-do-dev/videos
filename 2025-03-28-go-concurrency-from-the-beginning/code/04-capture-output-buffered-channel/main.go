package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Starting work...")

	tasks := []string{"Task1", "Task2", "Task3"}

	resultChan := make(chan string, len(tasks))

	var wg sync.WaitGroup
	wg.Add(len(tasks))

	for _, task := range tasks {
		go func() {
			defer wg.Done()
			process(task, resultChan)
		}()
	}

	go func() {
		defer close(resultChan)
		wg.Wait()
	}()

	for result := range resultChan {
		fmt.Println(result)
	}

	fmt.Println("Finished work.")
}

func process(task string, results chan<- string) {
	fmt.Println("Started", task)
	time.Sleep(time.Second)
	results <- fmt.Sprint("Finished ", task)
}
