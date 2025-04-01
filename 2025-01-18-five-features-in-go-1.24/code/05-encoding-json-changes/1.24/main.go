package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Message struct {
	Name string
	Body string
	Time time.Time `json:",omitzero"`
}

func main() {
	m := Message{"Alice", "Hello", time.Time{}}
	b, _ := json.Marshal(m)

	fmt.Println(string(b))
}

func (t T) IsZero() bool

type Status [T][]T
