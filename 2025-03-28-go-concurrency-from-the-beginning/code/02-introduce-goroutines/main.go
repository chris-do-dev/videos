package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Starting work...")

	tasks := []string{"Task1", "Task2", "Task3"}
	var wg sync.WaitGroup

	wg.Add(len(tasks))

	for _, task := range tasks {
		go func() {
			defer wg.Done()
			process(task)
		}()
	}

	wg.Wait()

	fmt.Println("All work completed!")
}

func process(name string) {
	fmt.Println("Started", name)
	time.Sleep(time.Second)
	fmt.Println("Finished", name)
}
