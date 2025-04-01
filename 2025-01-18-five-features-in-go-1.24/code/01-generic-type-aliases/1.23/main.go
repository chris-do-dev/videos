package main

import "fmt"

type Set[T comparable] = map[T]bool

func main() {
	set := Set[string]{"one": true, "two": true}

	fmt.Println("'one' in set:", set["one"])
	fmt.Println("'six' in set:", set["six"])
	fmt.Printf("set is %T\n", set)
}
