package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./file.txt")
	if err != nil {
		fmt.Println("falling back to default file")

		content, _ = os.ReadFile("./default.txt")
	}
	fmt.Print(string(content))
}
