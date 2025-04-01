package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Message struct {
	Name string
	Body string
	Time time.Time
}

func main() {
	m := Message{"Alice", "Body", time.Time{}}
	b, _ := json.Marshal(m)

	fmt.Println(string(b))
}
