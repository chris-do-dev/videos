Let's look at five more feature being introduced in Go 1.24. If you missed the last video covering another five features of Go 1.24, you can find a link to it in the description.
# Directory-limited filesystem access
Directory-limited filesystem access will be a feature in Go 1.24.

Consider this code that reads a file from the filesystem.
```go
package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	ROOT_DIR = "/var/lib/data/myapp/"
)

func main() {
	file, err := OpenFile("/home/devops/.ssh/id_ed25519")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("printing file")
	}
}

func OpenFile(filename string) (*os.File, error) {
	if ok := strings.HasPrefix(filename, ROOT_DIR); !ok {
		return nil, fmt.Errorf("path escapes from parent")
	}

	return os.Open(filename)
}

```
The important thing to notice here is the guard clause at the top of the OpenFile function, checking if the file is within the application's data directory. Only if the file is underneath the applications root directory will it be read, otherwise an error is returned. In our case, this prevented an SSH private key being read, a scenario that might occur if the user is able to specify arbitrary files to read.

If we run this code and attempt to open the SSH key, we will get the desired error.
```
$ go run main.go
path escapes from parent
```

In Go 1.24, this can be simplified using the new `os.Root` type. Let's rewrite our code to make use of it.

```go
package main

import (
	"fmt"
	"os"
)

const (
	ROOT_DIR = "/var/lib/data/myapp/"
)

func main() {
	root, _ := os.OpenRoot(ROOT_DIR)
	file, err := root.Open("/home/devops/.ssh/id_ed25519")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("printing file")
	}
}

```

We'll remove our OpenFile method, since it only exists to encapsulate our guard clause. We'll move the os.Open function to where our OpenFile function used to be, and create a root object by opening the applications root directory using the new `os.OpenRoot` function. Finally, we'll use this root object to open our file instead of `os`.

Running this code to read the SSH key gives us the same error message from before, but without having to implement it ourselves.
```
$ go run main.go
openat /home/devops/.ssh/id_ed25519: path escapes from parent
```
# Improved finalisers
Improved finalisers will be added to Go 1.24.

In Go, a finaliser is a function that is run when an object becomes unreachable from the code and the garbage collector is about to reclaim its memory. The point of them is to perform any additional steps required in the cleanup of the object that wouldn't be carried out by simply reclaiming the memory.

Let's write some code that shows how this works.

We'll start with our main package boilerplate before defining a custom type called `FileWrapper` that wraps an `os.File`. We'll create a temporary file and wrap it in our custom file wrapper. Now we'll set a finaliser function on our wrapped file in which we'll print some useful output before closing the wrapped file.

Now we need the file wrapper to go out of scope which we can simulate this by setting its pointer to `nil`. Finally we run the garbage collector and sleep for a second to allow the finaliser time to run.

As expected, we see the output from our function, indicating the wrapped file has been closed successfully.
```
$ go run main.go
Closing file
```

While this simple example may look fine, finalisers have a number of issues, including:
- a finaliser can only be called once. The way they work is the garbage collector runs them before removing them from the object, so if an object gets resurrected during finalisation, the finaliser will not be able to clean it up a second time, causing a memory leak.
- since objects with finalisers require two runs of the garbage collector to be released, it can be some time before the resources associated with that object are cleaned up.
- you can only attach a single finaliser to an object, so if you have multiple files and connections to close, you need to pile all that logic into a single function, which will be quite a mess and difficult to maintain.

Go 1.24 adds the `AddCleanup` function to the `runtime` library, with the recommendation that all new code use it in place of `SetFinalizer`.

`AddCleanup` behaves in the same way as `SetFinalizer` by being attached to an object and running after the garbage collector runs, but with some notable improvements over `SetFinalizer`:
- multiple cleanup functions can be attached to a single object, allowing you to split the cleanup of different resources into different functions.
- the chances of an object being resurrected are reduced since `Addcleanup` takes both the pointer and underlying resource as arguments, and panics of they are the same.

`AddCleanup` still shares some limitations with `SetFinalizer`, including:
- since cleanups are only run by the garbage collector, this couples them to memory utilisation, and doesn't take into account the utilisation of other scarce resources such as file descriptors.
- the cleanup function isn't guaranteed to run as soon as object becomes unreachable, rather at some point after. It is therefore not a good idea to rely on it for timely deallocation of memory.

So despite finalisers being improved in Go 1.24, they should still only be used sparingly as a safety net, and deferred functions should continue to be relied on for releasing resources.
# New interfaces - TextAppender and BinaryAppender
Go 1.24 introduces two new interfaces to the encoding package that are implemented by many more packages in the standard library: `TextAppender` and `BinaryAppender`. Since so many packages implement these interfaces, we'll just use one package to examine them.

We'll use the `time package` since it's one of the few that implements both interfaces. We'll just look at the `TextAppender`, since learning about that will automatically teach us about the `BinaryAppender`.

These two interfaces each define a single function: append text for the text appender and append binary for the binary appender.
```go
type TextAppender interface {
	AppendText(b []byte) ([]byte, error)
}

type BinaryAppender interface {
	AppendBinary(b []byte) ([]byte, error)
}
```

The purpose of these interfaces is to append text or binary representations of a type to a byte slice. Let's look at how that works with the `time.AppendText` function.

We'll start with some main package boilerplate and define a string to which we want our time appended. We'll append the time to this string in two different ways before printing it.

First, we'll define the current time and print it using the `fmt` package. Quite straightforward.

Now we'll use the new `AppendText` function to create a single string that we can print. 

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	prefix := "The time is "

	now := time.Now()
	fmt.Println(prefix, now)

	timeString, _ := now.AppendText([]byte(prefix))
	fmt.Println(string(timeString))
}
```

Running this code will print the same time twice, albeit with slight differences in formatting that can be ironed out if you ever do this in production.
```
$ go run main.go
The time is  2025-01-25 15:06:13.958818 +0000 GMT m=+0.000151542
The time is 2025-01-25T15:06:13.958818Z
```
# GOAUTH environment variable
Go 1.24 introduced a new environment variable called `GOAUTH` to make it easier to get modules from private repositories. It is a semi-colon separated list of authentication commands what will be tried one at a time until one works.

Go supports multiple methods for authenticating with private module repositories. Prior to Go 1.24, you couldn't tell Go which one to use, it would just use the first one that worked. The `GOAUTH` variable will allow you to specify your desired authentication method, giving developers more control over the development process.

Let's take a look at the possible values of this variable before looking at some examples of how to use it.

There are four possible values for `GOAUTH`:
- `off`, which disables authentication entirely
- `netrc`, which uses credentials specified in a `.netrc` file in your home directory
- `git <dir>`, which uses credentials stored in the git credential helper
- and `command` which specifies an arbitrary command to execute which returns headers that are attached to HTTP requests.

You can chain all of these together in any order you want, separated by semicolons.

Let's see how we use each of these values.
## `off`
The `off` option is a straightforward option that disables authentication. If we try to get a private module with this option set, we will get an authentication error.
```
$ GOAUTH=off go get github.com/bin-devops/private-go-module
go: module github.com/bin-devops/private-go-module: git ls-remote -q origin in /Users/chris/go/pkg/mod/cache/vcs/f5e2f31253ea14b02e56a70f91c883505fb1139901da3ebb845689d853d2a864: exit status 128:
        fatal: Cannot prompt because user interactivity has been disabled.
        fatal: could not read Username for 'https://github.com': terminal prompts disabled
Confirm the import path was entered correctly.
If this is a private repository, see https://golang.org/doc/faq#git_https for additional information.
```
## `netrc`
The `netrc` option tells `go get` to use credentials specified in a `.netrc` file found in the home directory.

If you're not familiar with this file, it's a general purpose file used by many applications to authenticate with services. For every hostname there is a username and password to authenticate with that host.
```netrc
machine github.com login bin-devops password <github_password>
machine gitlab.com login bin-devops password <gitlab_password>
```

If we now try to get a module with valid credentials set for the github.com machine, we get a new error.

```
$ GOAUTH=netrc go get github.com/bin-devops/private-go-module
github.com/bin-devops/private-go-module@v0.0.0-20250123114509-c8fbadafab75: verifying module: github.com/bin-devops/private-go-module@v0.0.0-20250123114509-c8fbadafab75: reading https://sum.golang.org/lookup/github.com/bin-devops/private-go-module@v0.0.0-20250123114509-c8fbadafab75: 404 Not Found
        server response:
        not found: github.com/bin-devops/private-go-module@v0.0.0-20250123114509-c8fbadafab75: invalid version: git ls-remote -q origin in /tmp/gopath/pkg/mod/cache/vcs/f5e2f31253ea14b02e56a70f91c883505fb1139901da3ebb845689d853d2a864: exit status 128:
                fatal: could not read Username for 'https://github.com': terminal prompts disabled
        Confirm the import path was entered correctly.
        If this is a private repository, see https://golang.org/doc/faq#git_https for additional information.
```

This is because our private module does not have a checksum that can be used to validate its authenticity, so we need to tell `go get` to ignore that by adding our github url to the `GOPRIVATE` variable.

```
$ GOPRIVATE=github.com/bin-devops GOAUTH=netrc go get github.com/bin-devops/private-go-module
go: added github.com/bin-devops/private-go-module v0.0.0-20250123114509-c8fbadafab75
```
## `git <dir>`
The `git` option specifies an absolute path to a git repository containing a `.git-credentials` file in its root. This repository has its credential helper set to this file, allowing the credentials to be used by git. Let's see how that works.

We'll create a new directory for our demo and move into it. We'll then disable authenticity checking of Go modules as well as the git terminal prompt. This last one will prevent Git asking us for a username and password in the terminal.

Next, we'll create a git repo, set its credential helper to be a local file called `.git-credentials` before adding our credentials to the credentials file. Finally, we'll initialise our module and get our private module.
```
mkdir goauth-demo
cd goauth-demo

export GOPROXY=direct
export GOSUMDB=off
export GIT_TERMINAL_PROMPT=0

git init
git config credential.helper 'store --file=.git-credentials'

echo "https://bin-devops:<token>@github.com" >> .git-credentials

go1.24rc1 mod init goauth_demo
```

```
$ go get github.com/bin-devops/private-go-module
go: added github.com/bin-devops/private-go-module v0.0.0-20250123114509-c8fbadafab75
```
## `command`
The final option for `GOAUTH` is a command to invoke a custom authenticator. The output of this command is a set os HTTP headers to be attached to requests to the private repository.

You can specify any command you want, provided its output is in the following format, which returns this example.

```
Response      = { CredentialSet } .
CredentialSet = URLLine { URLLine } BlankLine { HeaderLine } BlankLine .
URLLine       = /* URL that starts with "https://" */ '\n' .
HeaderLine    = /* HTTP Request header */ '\n' .
BlankLine     = '\n' .
```

```
https://example.com/
https://example.net/api/

Authorization: Basic <token>

```
Here, we have an authorisation header with a bearer token generated by the command, which allows authentication with a private repository without leaving hard-coded credentials lying at rest in a netrc file.
# go/types iterator functions
Various types in the `go/types` package are getting new functions that return an iterator over their internal data. Let's look at an example using the `Union` type.

We'll start with our main package boilerplate before defining an array of terms which we'll use to create a union.

Now in Go 1.23, if we wanted to iterate over these terms, we would use a C-style for loop with the length of the union at the upper bound. We'd then use the `Term` function to access the term at the given index.

In Go 1.24, we just range over the new `Terms` function.

Here is the full list of new functions being introduced to this package.

```
Interface.EmbeddedTypes
Interface.ExplicitMethods
Interface.Methods
MethodSet.Methods
Named.Methods
Scope.Children
Struct.Fields
Tuple.Variables
TypeList.Types
TypeParamList.TypeParams
Union.Terms
```
# Outro
That's five more features of the upcoming Go 1.24. If you made it this far and found this video useful, please consider liking the video. It really helps the channel. If you didn't see my first video on five features of Go 1.24, why not check it out. You'll find a link to it in the description.
