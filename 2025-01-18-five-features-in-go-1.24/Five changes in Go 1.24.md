Let's look at five features of the upcoming 1.24 release of Go.
# Generic Type Aliases
Generic type aliases will be introduced in Go 1.24.

There are two ways to create a custom type in Go - type definition and type alias
```go
// 'type definition', where the types of Status and int are not equal
type Status int 

// 'type alias', where the types of Status and int are equal
type Status = int
```

Let's look at some code that illustrates the difference between the two.

We'll start with main package boilerplate
```go
package main

func main() {

}
```

before defining two new custom types: one as a definition and the other as an alias
```go
package main

type StatusDefinition int

type StatusAlias = int

func main() {

}
```

next we'll declare a variable of each of our custom types,
```go
package main

type StatusDefinition int

type StatusAlias = int

func main() {
    var def StatusDefinition



    var alias StatusAlias


}
```

initialise them,
```go
package main

type StatusDefinition int

type StatusAlias = int

func main() {
    var def StatusDefinition
	def = 1


    var alias StatusAlias
	alias = 1

}
```

before finally printing each variables type
```go
package main

import (
	"fmt"
	"reflect"
)

type StatusDefinition int

type StatusAlias = int

func main() {
	var def StatusDefinition
	def = 1
	fmt.Println("Type of StatusDefinition: ", reflect.TypeOf(def))

	var alias StatusAlias
	alias = 1
	fmt.Println("Type of StatusAlias: ", reflect.TypeOf(alias))
}
```

If we now run this code we'll see the difference between the two:
```
Type of StatusDefinition:  main.StatusDefinition
Type of StatusAlias:  int
```

the type alias is just an int, whereas the type definition is it's own type

So far so good, but let's introduce templating to our custom types. We'll start by making our two types generic
```go
package main

import (
	"fmt"
	"reflect"
)

type StatusDefinition[T comparable] []T

type StatusAlias[T comparable] = []T

func main() {
	var def StatusDefinition
	def = 1
	fmt.Println("Type of StatusDefinition: ", reflect.TypeOf(def))

	var alias StatusAlias
	alias = 1
	fmt.Println("Type of StatusAlias: ", reflect.TypeOf(alias))
}
```

before updating the declaration
```go
package main

import (
	"fmt"
	"reflect"
)

type StatusDefinition[T comparable] []T

type StatusAlias[T comparable] = []T

func main() {
	var def StatusDefinition[int]
	def = 1
	fmt.Println("Type of StatusDefinition: ", reflect.TypeOf(def))

	var alias StatusAlias[int]
	alias = 1
	fmt.Println("Type of StatusAlias: ", reflect.TypeOf(alias))
}
```

and finally the initialisation
```go
package main

import (
	"fmt"
	"reflect"
)

type StatusDefinition[T comparable] []T

type StatusAlias[T comparable] = []T

func main() {
	var def StatusDefinition[int]
	def = append(def, 1)
	fmt.Println("Type of StatusDefinition: ", reflect.TypeOf(def))

	var alias StatusAlias[int]
	alias = append(alias, 1)
	fmt.Println("Type of StatusAlias: ", reflect.TypeOf(alias))
}
```

If we try to run this, however, we get an error telling us a generic type cannot be an alias
```
generic type cannot be alias
```

Before Go 1.24, only type definitions could be generic, but Go 1.24 allows type aliases to also be generic.

```go
package main

import (
	"fmt"
	"reflect"
)

type StatusDefinition[T comparable] []T

type StatusAlias[T comparable] = []T

func main() {
	var def StatusDefinition[int]
	def = append(def, 1)
	fmt.Println("Type of StatusDefinition: ", reflect.TypeOf(def))

	var alias StatusAlias[int]
	alias = append(alias, 1)
	fmt.Println("Type of StatusAlias: ", reflect.TypeOf(alias))
}
```

If we now run our generic code with Go 1.24, it works as expected
```
Type of StatusDefinition:  main.StatusDefinition[int]
Type of StatusAlias:  []int
```
# Weak pointers
Weak pointers will be introduced in Go 1.24.

Go's garbage collector will reclaim the memory of any object that becomes unreachable from your code. Let's see how that works.

We'll begin with main package boilerplate
```go
package main

import (

)

func main() {

}
```

we'll create a function that returns the current heap size
```go
package main

import (
	"runtime"
)

const BYTE_SIZE = 1024

func main() {
	initialHeap := heapSize()
}

func heapSize() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc / BYTE_SIZE
}
```

We'll define a custom type
```go
package main

import (
	"fmt"
	"runtime"
)

const KILOBYTE_SIZE = 1024

type Object []byte

func (b Object) String() string {
	return fmt.Sprintf("Object size: %d", len(b)/KILOBYTE_SIZE)
}

func main() {
	initialHeap := heapSize()
}

func heapSize() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc / KILOBYTE_SIZE
}
```

and create a strong pointer to an object
```go
package main

import (
	"fmt"
	"runtime"
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
}

func heapSize() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc / BYTE_SIZE
}
```

next we'll print the pointer
```go
package main

import (
	"fmt"
	"runtime"
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
	
	fmt.Println("Before garbgage collection")
	fmt.Println("   strong =", strong)
}

func heapSize() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc / BYTE_SIZE
}
```

trigger the garbage collector
```go
package main

import (
	"fmt"
	"runtime"
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
	
	fmt.Println("Before garbgage collection")
	fmt.Println("   strong =", strong)
	
	runtime.GC()
}

func heapSize() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc / BYTE_SIZE
}
```

print the pointer again
```go
package main

import (
	"fmt"
	"runtime"
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

	fmt.Println("Before garbgage collection")
	fmt.Println("   strong =", strong)

	runtime.GC()

	fmt.Println("After garbgage collection")
	fmt.Println("   strong =", strong)
}

func heapSize() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc / BYTE_SIZE
}
```

then compare the heap size
```go
package main

import (
	"fmt"
	"runtime"
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

	fmt.Println("Before garbgage collection")
	fmt.Println("   strong =", strong)

	runtime.GC()

	fmt.Println("After garbgage collection")
	fmt.Println("   strong =", strong)

	fmt.Printf("heap size diff = %d KB\n", heapSize()-initialHeap)
}

func heapSize() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc / BYTE_SIZE
}
```

when we run this code we see that the pointer is present before and after garbage collection
```
$ go run main
Before garbgage collection
   strong = Object size: 1000
After garbgage collection
   strong = Object size: 1000
heap size diff = 1012 KB
```

However, if we remove any mention of the strong pointer after garbage collection, 
```go
package main

import (
	"fmt"
	"runtime"
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

	fmt.Println("Before garbgage collection")
	fmt.Println("   strong =", strong)

	runtime.GC()

	fmt.Printf("heap size diff = %d KB\n", heapSize()-initialHeap)
}

func heapSize() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc / BYTE_SIZE
}
```

we see that the heap size is 1000KB smaller than before.
```
$ go run main
Before garbgage collection
   strong = Object size: 1000
heap size diff = 10 KB
```

This has happened because we removed the use of the strong pointer after garbage collection. This meant the pointer was unreachable from any code after the garbage collection run, so it was removed.

Let's now look at how weak pointers affect this. We'll return to printing the strong pointer after garbage collection,
```go
package main

import (
	"fmt"
	"runtime"
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

	fmt.Println("Before garbgage collection")
	fmt.Println("   strong =", strong)

	runtime.GC()

	fmt.Println("After garbgage collection")
	fmt.Println("   strong =", strong)

	fmt.Printf("heap size diff = %d KB\n", heapSize()-initialHeap)
}

func heapSize() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc / BYTE_SIZE
}
```

we'll add a weak pointer,
```go
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

	runtime.GC()

	fmt.Println("After garbgage collection")
	fmt.Println("   strong =", strong)

	fmt.Printf("heap size diff = %d KB\n", heapSize()-initialHeap)
}

func heapSize() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc / BYTE_SIZE
}
```

which we'll print alongside our strong pointer.
```go
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
	fmt.Println("   strong =", strong)
	fmt.Println("   weak =", weak.Value())

	fmt.Printf("heap size diff = %d KB\n", heapSize()-initialHeap)
}

func heapSize() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc / BYTE_SIZE
}
```

Running this give us output similar to our original strong pointer output:
```
$ go run main
Before garbgage collection
   strong = Object size: 1000
   weak = Object size: 1000
After garbgage collection
   strong = Object size: 1000
   weak = Object size: 1000
heap size diff = 1012 KB
```

both the strong and weak pointers are pointer to our object, and the heap size shows the object is still allocated.

Now if we remove the strong pointer from after the garbage collection, but keep the weak pointer,

```go
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
	return m.Alloc / BYTE_SIZE
}
```

we see that the object has been deallocated from the heap, despite the weak pointer still being present
```
$ go run main
Before garbgage collection
   strong = Object size: 1000
   weak = Object size: 1000
After garbgage collection
   weak = <nil>
heap size diff = 12 KB
```

This has happened because the weak pointer was not counted towards the reachability of the object, only the strong pointer was. 

An example use case of weak pointers is for a cache, where we want objects to be able to expire and be removed despite still having a pointer. We can perform a `nil` check on the pointer to see if the object is still cached, or if we need to refresh it.
# Managing tools in go.mod
Go 1.24 eliminates the need for `tools.go`, allowing us to track versions of our tools in `go.mod`.

In Go 1.23 or lower, if you wanted to managed your tools, such as linters, in your project using go.mod, you needed to do a little trick. Let me show you.

We'll initialise a new go module using version 1.23
```
$ go version
go version go1.23.4 darwin/arm64

$ go mod init tools_pattern
go: creating new go.mod: module tools_pattern

$ cat go.mod
module tools_pattern

go 1.23.4
```

Now we'll add `golint` to our project
```
$ go get golang.org/x/lint/golint
go: added golang.org/x/lint v0.0.0-20241112194109-818c5a804067
go: added golang.org/x/tools v0.0.0-20200130002326-2f3ba24bd6e7
```

If we look in `go.mod` we'll see two dependencies added:
```
$ cat go.mod
module tools_pattern

go 1.23.4

require (
	golang.org/x/lint v0.0.0-20241112194109-818c5a804067 // indirect
	golang.org/x/tools v0.0.0-20200130002326-2f3ba24bd6e7 // indirect
)
```

but they're both indirect. The next time we run `go mod tidy`, we lose them
```
$ go mod tidy
$ cat go.mod
module tools_pattern

go 1.23.4
```

This is because the only reference to those dependencies was in the `golint` binary, which isn't included in our package via an `import` statement.

Before Go 1.24, we would work around this by faking the import using a `tools.go` file.
```go
package main

import (
	_ "golang.org/x/lint/golint"
)
```

We would include this file in our main package, and import our tools to the blank assignment. If we then run `go mod tidy`, we'll see it downloading our tools

```
$ go mod tidy
go: finding module for package golang.org/x/lint/golint
go: found golang.org/x/lint/golint in golang.org/x/lint v0.0.0-20241112194109-818c5a804067
```

And if we look at `go.mod` we'll see `lint` is no longer an indirect dependency, since it's directly imported by `golint` in our `tools.go`.
```
$ cat go.mod
module tools_pattern

go 1.23.4

require golang.org/x/lint v0.0.0-20241112194109-818c5a804067

require golang.org/x/tools v0.0.0-20200130002326-2f3ba24bd6e7 // indirect
```

With Go 1.24, we no longer need the `tools.go` since tools can be tracked in `go.mod` by adding the `-tool` flag to the `go get` command:
```
$ go get -tool golang.org/x/lint/golint
go: added golang.org/x/lint v0.0.0-20241112194109-818c5a804067
go: added golang.org/x/tools v0.0.0-20200130002326-2f3ba24bd6e7
```

If we then look in our `go.mod`
```
$ cat go.mod
module tools_pattern

go 1.24rc1

tool golang.org/x/lint/golint

require (
	golang.org/x/lint v0.0.0-20241112194109-818c5a804067 // indirect
	golang.org/x/tools v0.0.0-20200130002326-2f3ba24bd6e7 // indirect
)
```

we'll see a new `tool` directive added, allowing us to track our external tools directly in `go.mod` without the workaround of a `tools.go`.
# New functions in bytes/strings packages
Go 1.24 introduces new functions to the bytes and strings packages for working with iterators. Both packages introduce the same functions, so I go over them using the strings package.

All of these functions have the benefit of being able to iterate over a slice of data without having to allocate memory for the entire slice - only the element pointed to by the iterator is allocated.

Lines returns an iterator over the newline-terminated lines in the string.

For example, given a multi-line string, in Go 1.23 or lower, we would split the string and iterate over the resulting slice whereas in Go 1.24, we only need to call the Lines function
```go
	s := `line 1\nline 2\nline 3`

	fmt.Println(" In Go 1.23")
	for _, line := range strings.Split(s, "\n") {
		fmt.Println("   ", line)
	}

	fmt.Println(" In Go 1.24")
	for line := range strings.Lines(s) {
		fmt.Print("   ", line)
	}
```

SplitSeq returns an iterator over all substrings of s separated by the given separator. It can be used in place of Split
```go
	s := "line 1+line 2+line 3"

	fmt.Println(" In Go 1.23")
	for _, line := range strings.Split(s, "+") {
		fmt.Println("   ", line)
	}

	fmt.Println(" In Go 1.24")
	for line := range strings.SplitSeq(s, "+") {
		fmt.Println("   ", line)
	}
```

SplitAfterSeq returns an iterator over substrings of a string, split after each instance of the separator and is equivalent to the SplitAfter function
```go
	s := "line 1+line 2+line 3"

	fmt.Println(" In Go 1.23")
	for _, line := range strings.SplitAfter(s, "+") {
		fmt.Println("   ", line)
	}

	fmt.Println(" In Go 1.24")
	for line := range strings.SplitAfterSeq(s, "+") {
		fmt.Println("   ", line)
	}
```

FieldsSeq is an iterator-based equivalent to Fields, splitting the string around whitespace characters
```go
	s := " line1     line2 line3   "

	fmt.Println(" In Go 1.23")
	for _, line := range strings.Fields(s) {
		fmt.Println("   ", line)
	}

	fmt.Println(" In Go 1.24")
	for line := range strings.FieldsSeq(s) {
		fmt.Println("   ", line)
	}
```

and finally FieldsFuncSeq is an iterator-based equivalent to FieldsFunc, splitting a string into fields based on the evaluation of a function
```go
	s := "foo bar+baz asd!"

	f := func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}

	fmt.Println(" In Go 1.23")
	for _, field := range strings.FieldsFunc(s, f) {
		fmt.Println(" ", field)
	}

	fmt.Println(" In Go 1.24")
	for field := range strings.FieldsFuncSeq(s, f) {
		fmt.Println(" ", field)
	}
```
# encoding/json changes
Go 1.24 introduces an `omitzero` tag to remove zero values when marshalling json.

Consider this code
```go
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
	m := Message{"Alice", "Hello", time.Time{}}
	b, _ := json.Marshal(m)

	fmt.Println(string(b))
}
```

It's a simple example of turning a struct into json, and if we run it we can see our json object
```
$ go run main.go
{"Name":"Alice","Body":"Hello","Time":"0001-01-01T00:00:00Z"}
```

Notice the Time value in the printed json. Our message struct had an uninitialized time in it, which got translated to this zero representation of time when marshalled.

In order to remove this empty time, we can use the new `omitzero` tag in our struct like so
```go
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
	m := Message{"Alice", "Body", time.Time{}}
	b, _ := json.Marshal(m)

	fmt.Println(string(b))
}
```

and when we run this, we no longer have a time field in the json
```
$ go run main.go
{"Name":"Alice","Body":"Hello"}
```

Zero-ness is determined by implementing an `IsZero()` function on the type.

The existing `omitempty` tag doesn't work with uninitialised time values since they are not technically empty.

I'd love to know in the comments if you're excited for any of these new features, or if there are new features you want me to take a look at in a later video.