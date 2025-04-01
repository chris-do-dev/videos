Error handling is an essential element of robust software. Most programming languages deal with errors by throwing an exception that is caught further up the callstack.

```javascript
try {
	dosomething()
catch (e)
	log("something went wrong")
}
```

```java
try {
	dosomething()
}
catch(Exception e) {
	logger.log(Level.WARNING, "Something went wrong");
}
```

```python
try:
	dosomething()
except Exception:
	print("Something went wrong")
```

Go, on the other hand, does things a little differently
```go
result, err := dosomething()
if err != nil {
	fmt.Println("something went wrong")
}
```

In Go, errors are values, so instead of throwing an exception when something goes wrong, an error is initialised and either handled immediately or returned from the function and handled by the caller. People who are used to these other languages can find this aspect of Go a little unsettling, but let me tell you why I think this is a better way of dealing with errors.

Consider this Typescript code:
```typescript
try {
	dosomething()
catch (e)
	log("something went wrong in dosomething")
}
```

If something goes wrong, we know it went wrong in the dosomething function, since it's the only one being wrapped in the try-catch block. But what if we had more than one function

```typescript
try {
	var result1 = dosomething()
	dosomethingelse(result1)
catch (e)
	log("something went wrong")
}
```

now, if an exception is thrown from either dosomething or dosomethingelse, it will be caught and logged, but we won't know which function encountered the error, preventing us adding useful context to the error message.

We could solve this by wrapping each function in its own try catch block
```typescript
try {
	var result1 = dosomething()
catch (e)
	log("something went wrong")
}
try {
	dosomethingelse(result1)
catch (e)
	log("something else went wrong")
}
```
allowing us to log a different message depending on which function encountered the error, but this has come at the price of code that is more verbose and harder to read.

Additionally, The try-catch style of error handling can lead to errors being handled far from where they were generated. Our dosomething and dosomethingelse functions might be calling additional functions which can throw errors of their own. Catching these errors mulltiple stack frames away from where they were generated is another way we are prevented adding useful context to the error message.

Go's style of error handling is designed to solve these problems. When calling a function in go, many will return an error value as their last result, which can be checked by the caller.
```go
result, err := dosomething()
if err != nil {
	fmt.Println("something went wrong")
}
```

If the error is non-nil, this indicates something went wrong, and we handle this error inside an if statement immediately after the function. This is an example of a design pattern called Locality of Behaviour, which is the idea that the behaviour of a piece of code should be fully understood just by looking at that piece of code.

A link to this can be found in the description (https://htmx.org/essays/locality-of-behaviour/).

Let's look at an example of error handling in Go.

We'll start with our main package boilerplate
```go
package main

import (

)

func main() {

}
```

then we'll read the contents of a file
```go
package main

import (
	"os"
)

func main() {
	content, err := os.ReadFile("./file.txt")
}
```

when choosing a name for the error variable, err is idiomatic, but not a hard rule.

Since the ReadFile function is returning an error, we will want to handle it. We'll add our `nil` check,

```go
package main

import (
	"os"
)

func main() {
	content, err := os.ReadFile("./file.txt")
	if err != nil {

	}
}

```

log that we've encountered an error and what we're doing about it

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./file.txt")
	if err != nil {
		fmt.Println("falling back to default file")
	}
}
```

then read the default file like we said in our log message
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./file.txt")
	if err != nil {
		fmt.Println("falling back to default file")

		content, _ = os.ReadFile("./default.txt")
	}
}
```

When reading the default file, we assign the error to the blank assignment, that's the underscore, indicating to anyone reading the code that we don't care about its value. We can get away with ignoring the error here because we are confident that the default file will exist.

Finally, lets print our content
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./file.txt")
	if err != nil {
		fmt.Println("falling back to default file")

		content, _ = os.ReadFile("./default.txt")
	}
	fmt.Print(string(content))
}
```

Before we run this code, we need to create our default file and give it some content
```sh
echo "default file contents" > default.txt
```

In this example, we've been able to not only catch and handle an error, but we've been able to recover from it without leaving the current function. Each function's error is handled immediately after the function in question, which gives two main benefits.

The first is about maintaining the flow of execution through the code. When errors are handled by either logging them as soon as they happen, or returning them to the calling function, we maintain control over the flow of execution through our code. When an error occurs in Go, we get to choose what happens next.

The other benefit is stylstic. When writing Go code, we separate errors into two type: those we anticipate and those we don't. Those we anticipate are the ones we return, and good Go code is written to handle any error that the developer can anticipate. For example, here are some of the errors anticipated by the `os` package that could have occurred when we were reading our files earlier. 

```go
var (
	ErrPermission = fs.ErrPermission // "permission denied"
	ErrExist      = fs.ErrExist      // "file already exists"
	ErrNotExist   = fs.ErrNotExist   // "file does not exist"
	ErrClosed     = fs.ErrClosed     // "file already closed"
)
```

When we attempted to open the non-existent file, we received an `ErrNotExist` error. When we're writing packages for use by other developers, we should implement errors to a similar level of detail.

The other type of errors are the unanticipated ones. When they occur, the program is said to panic. These can include programmer error, where the programmer hasn't provided valid input to a function, such as providing a `nil` driver to a database function.

```go
conn, _ := sql.Open("libsql", "file:"+filename)

// we passed an invalid conn here because we didn't check the previous error
Exec(ctx, conn, "CREATE TABLE records")
```

or if a mandatory dependency cannot be satisfied, such a a regexp not compiling.
```go
_, err := regexp.MustCompile("[")
```

It's possible to cause your own panic using the panic function, and to recover it using the recover function. Let's look at how that's done.

We'll start again with our main package boilerplate:
```go
package main

func main() {

}
```

Next we'll add a function that panics
```go
package main

func main() {
	doSomething1()

	fmt.Println("finished main")
}

func doSomething1() {
	doSomething2()
}

func doSomething2() {
	panic("panic")
}
```

Before finally recovering from the panic in a deferred function
```go
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
```

We need to place our `recover` in a deferred function because once the program panics, deferred functions are the only ones that are run.

If we run this code, we'll see the panic being caught by the deferred function before control being returned to `main` to allow the program to finish running.

```
caught a panic
finished main
```

Some are tempted to use this pattern to emulate the try-catch behaviour from other languages. This is inadvisable for two reasons. The first and most important one is that it is places error handling far away from the code that generated it, loosing the readability and control gained from handling an error returned from a function on the very next line.

The second is inefficiency. A panic is intended to cause the entire program to exit, with the whole stack being unwound in search of the source of the panic. Since the program isn't expecting to continue running, this isn't intended to be efficient and that time will add up if you use this technique liberally throughout your code.

I'd love to hear from you in the comments if you found this useful. Also, left me know if there's another aspect of Go you'd like to hear about.