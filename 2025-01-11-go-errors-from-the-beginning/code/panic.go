package main

import "fmt"

func main() {
	doSomething1()

	fmt.Println("finished main")
}

func doSomething1() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("caught a panic")
		}
	}()

	doSomething2()
}

func doSomething2() {
	panic("panic")
}
