package main

import (
	"fmt"
	"runtime"
	"weak"
)

const BYTE_SIZE = 1024

type Object []byte

func (b Object) String() string {
	return fmt.Sprintf("Object size: %d", len(b)/BYTE_SIZE)
}

func main() {
	initialHeap := heapSize()

	objectSize := 1000
	// Multiple by 1024 to make the size of the heap substantial. Without this we
	// can't see a noticeable difference in the size
	strong := make(Object, objectSize*BYTE_SIZE)
	weak := weak.Make(&strong)

	fmt.Println("Before garbgage collection")
	fmt.Println("   strong =", strong)
	fmt.Println("   weak =", weak.Value())

	runtime.GC()
	// runtime.KeepAlive(b)

	fmt.Println("After garbgage collection")
	fmt.Println("   weak =", weak.Value())

	fmt.Printf("heap size diff = %d KB\n", heapSize()-initialHeap)
}

func heapSize() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc / BYTE_SIZE
}

pointer := cache.Get(objectId)
if pointer == nil {
  pointer = origin.Get(objectId)
}
