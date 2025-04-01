package main

import (
	"fmt"
	"log"
	"os"
)

const (
	ROOT_DIR = "/var/lib/data/myapp/"
)

func main() {
	root, err := os.OpenRoot(ROOT_DIR)
	if err != nil {
		log.Fatal(err)
	}
	file, err := root.Open("/home/devops/.ssh/id_ed25519")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("printing file")
	}
}
