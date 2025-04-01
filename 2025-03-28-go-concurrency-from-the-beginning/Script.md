Go's concurrency model is composed of only a few simple parts that combine to
create sophisticated yet elegant flows. In this video, we'll explore concurrency
in Go from the ground up using examples of working code, each building on the
last.

# 01 Non-concurrent

We'll start with the non-concurrent code you see now and evolve it one example
at a time to see how each of Go's concurrency components build on top of one
another. This simple example shows three processes being run sequentially, each
taking a second. If we run this code we'll see each task start, finish a second
later, then the next task immediately start.

We can begin adding concurrency to this by introducing `goroutines` to process
each of our tasks in parallel.

# 02 Introduce goroutines

We'll start by removing the three process function invocations and replacing
them with a range over a slice of task strings, inside which we have an
anonymous function and inside that is our process function.

The go keyword in front of our anonymous function is how we tell go to run the
given function concurrently. We could give the go keyword our process function
without the anonymous function, but we're going to need it shortly.

If we run this code now, we'll see our start and complete messages with nothing
from the process function. This is because the main function continued running
after it started the goroutines and returned before they could complete their
work. To prevent this we're going to introduce another concurrency feature
called wait groups.

We'll create a wait group variable from the sync package and call its Wait
function after the goroutines are started. This will make the application wait
until the goroutines have all completed before continuing execution. We'll tell
the wait group how many goroutines it needs to wait for, then have each
goroutine tell the wait group it's done.

Let's run this and see how it looks. This one runs much quicker, with each task
starting at the same time and finishing at the same time. We've just tripled the
speed of our application.

This example doesn't have any result being returned to the main function. If we
did want that, we need to introduce another of Go's concurrency features:
channels.

# 03 Capture Output Unbuffered Channel

We'll start by removing the wait group from our code, create a channel using the
make function and defer the closure of the channel. We'll then update the
signature of our process function to take the channel as an argument, and inside
the process function we'll replace our final print statement with the same text
being written to the channel.

After creating our goroutines we'll range over our tasks again, this time
creating a variable called result, reading the contents of our channel and
finally printing it.

If we run this code, we get the exact same output as last time, but the code is
quite different. Let's take a look.

The key thing to note here is that returning a result from a goroutine doesn't
use a normal return statement. Instead the channel allows us to pass data
between goroutines.

The arrow operator is used to interact with a channel. The arrow on the left
means we're reading from the channel, and on the right means we're writing to
it. When passing a channel to a function, using the arrow in the signature is an
optional way to define the channel as read or write only. Omitting the arrow and
passing a bidirectional channel is like using the any type - unclear and
potentially dangerous.

In our previous code, we had a wait group ensuring the caller function didn't
exit before all the goroutines were complete. We still have that same problem
when using channels, but we solve it differently.

When we write some code to read from a channel, execution of the code will block
until there is something in the channel to be read. Values in a channel are
removed as they are read. When writing to a channel, the same thing happens in
reverse - if there's a value in the channel waiting to be read the function will
block until the channel is empty.

These single value channels are one of two types of channel Go has to offer,
known as unbuffered channels. The other type, buffered channels, allows us to
write multiple values into a queue.

# 04 Capture Output Buffered Channels

For this example, we'll start by modifying the creation of our channel to
specify a size. This is the number of values is can hold before further attempts
to write to it block.

We'll reintroduce the wait group from our earlier example. It'll behave the same
as it did previously with one difference. We're going to place the wait function
inside another goroutine along with a deferred call to close our channel.
Placing these two lines inside a goroutine ensure our channel gets closed once
the goroutines have told the wait group they are done.

With the goroutines complete and the channel closed, we can range over the
channel and print each result. Note we don't need the arrow syntax this time.
Running this code gives us the same output as our previous two examples. Next,
we'll introduce control flow for our goroutines and channels using the select
statement.

# 05 Select Statement

We'll modify the range over our tasks to get the current index and pass it to
the process function as the task duration after incrementing it.

We'll use the After function from the time package to create two channel
variables in our process function - timeout and taskDone. The After function
takes a duration and returns a channel which will have a value written to it
once the duration has elapsed.

We'll add a select statement to our process function to handle the timeout
channel. A select statement is a switch statement that deal exclusively with
channels. It evaluates all cases and executes the one that has a value on the
channel waiting to be read. If multiple cases have channels waiting to be read,
it picks one at random. If none of the channels have a value, it either executes
the default statement, which we're not including here, or blocks until there is
something to read.

We'll add two cases to our select: one for if the timeout is reach and one for
if the task duration is reached. We don't care about what's on either channel,
so we'll just send our timeout or finished message to the results channel.

Returning to our main function, we'll range over our tasks and read the value
from the channel before printing it.

If we run this code now, we'll see something slightly different from before. All
three tasks start as normal, and the first two finish as normal, but the third
one doesn't, and instead we are told the timeout has been reached.

We've implemented a timeout here using functions from the time package, but
there's another way we can time out a goroutine, and that's using a context.

# 06 Timeout with Context

We'll move our timeout variable into the main function and convert it from a
channel to a duration. We'll then create a context and a cancel function with
the timeout variable. We'll defer running the cancel function to ensure all
running goroutines get cancelled and cleaned up if they're still running when
the main function exits or panics.

We'll modify our process function to take the context as an argument. Inside the
process function, we'll modify our select statement to replace the timeout
channel with the done function of the context. The channel returned by this
function will have a value written to it when the timeout is reached.

Running this code give us the same output as the previous example where we used
the timeout channel.

# 07 Mutexes

goroutines, channels and select statements combine to form one of Go's two ways
of writing concurrent code, and ideal if you want to coordinate different
goroutines as they pass and transform data between themselves. If you want to
share a single piece of data between multiple go routines then a mutex is a
better option.

We'll illustrate this with a new example.

Let's create a new type called inventory manager which will have a map of
strings to integers as the inventory, and a mutex. We'll see how the mutex works
shortly as we complete this example.

We'll attach two functions to our inventory manager type - update and read. When
we update a value, we first lock the mutex, defer the unlock then write our new
value.

Read is very similar. We lock the mutex, defer the unlock, read the specified
value and return it.

When we lock the mutex, any other goroutine that tries to access the inventory
will block until the mutex is unlock.

Here is the equivalent code written using channels. The mutex-based approach is
much shorter and elegant than the channel based one.

And that's an overview of Go's concurrency model. If you made it this far and
found it useful, why not leave a like on the video. It helps a lot more than you
might think.
