package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

type FileWrapper struct {
	File *os.File
}

func main() {
	file, _ := os.CreateTemp("", "example.txt")
	wrapper := &FileWrapper{File: file}

	runtime.SetFinalizer(wrapper, func(f *FileWrapper) {
		fmt.Println("Closing file")
		f.File.Close()
	})

	wrapper = nil

	runtime.GC()
	time.Sleep(1 * time.Second)
}
