package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

var resurectedFileWrapper *FileWrapper

type FileWrapper struct {
	File *os.File
}

func main() {
	file, _ := os.CreateTemp("", "example.txt")
	wrapper := &FileWrapper{File: file}

	runtime.SetFinalizer(wrapper, func(f *FileWrapper) {
		fmt.Println("Closing file")
		resurectedFileWrapper = f
		f.File.Close()
	})

	wrapper = nil

	runtime.GC()                // Suggest GC run
	time.Sleep(1 * time.Second) // Give time for finalizer to execute

	if resurectedFileWrapper != nil {
		fmt.Println("File wrapper has been resurected. Oops")
	}

	resurectedFileWrapper = nil
	runtime.GC()                // Suggest GC run
	time.Sleep(1 * time.Second) // Give time for finalizer to execute
}
