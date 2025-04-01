package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting work...")

	process("Task1")
	process("Task2")
	process("Task3")

	fmt.Println("All work completed!")
}

func process(task string) {
	fmt.Println("Started", task)
	time.Sleep(time.Second)
	fmt.Println("Finished", task)
}
