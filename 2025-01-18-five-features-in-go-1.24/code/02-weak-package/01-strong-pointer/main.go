package main

import (
	"fmt"
	"runtime"
	"weak"
)

const KILOBYTE_SIZE = 1024

type Object []byte

func (b Object) String() string {
	return fmt.Sprintf("Object size: %d", len(b)/KILOBYTE_SIZE)
}

func main() {
	initialHeap := heapSize()

	objectSize := 1000
	strong := make(Object, objectSize*KILOBYTE_SIZE)
	weak := weak.Make(&strong)

	fmt.Println("Before garbgage collection")
	fmt.Println("   strong =", strong)
	fmt.Println("   weak =", weak.Value())

	runtime.GC()

	fmt.Println("After garbgage collection")

	fmt.Println("   weak =", weak.Value())

	fmt.Printf("heap size diff = %d KB\n", heapSize()-initialHeap)
}

func heapSize() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc / KILOBYTE_SIZE
}
