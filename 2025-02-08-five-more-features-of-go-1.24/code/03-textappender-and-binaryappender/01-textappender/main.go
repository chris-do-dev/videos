package main

import (
	"fmt"
	"time"
)

func main() {
	prefix := "The time is "

	now := time.Now()
	fmt.Println(prefix, now)

	timeTextAppended, _ := now.AppendText([]byte(prefix))
	fmt.Println(string(timeTextAppended))
}
