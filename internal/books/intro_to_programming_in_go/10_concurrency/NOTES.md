# Concurrency
Making progress on more than one task simultaneously is known as concurrency.

## Goroutines
A goroutine is a function that is capable of running concurrently with other
functions. To create a goroutine, we use the keyword `go` followed by a function
invocation:

```
package main

import "fmt"

func f(n int) {
     for i := 0; i < 10; i++ {
           fmt.Println(n, ":", i)
     }
}

func main() {
     go f(0)
     var input string
     fmt.Scanln(&input)
}
```

There are two goroutines. The first is implicit and it is the main function
itself. The second goroutine is created when we call `go f(0)`. Normally, when
a function is invoked, the program will execute all the statements in a function
and then return to the next line following invocation. With a goroutine, we
return immediately to the next line and don't wait for the function to complete.

Goroutines are lightweight and we can easily create multiples of them.

## Channels
Channels provide a way for two goroutines to communicate with one another and
synchronize their execution.

```
package main

import (
     "fmt"
     "time"
)

func pinger(c chan string) {
     for i := 0; ; i++ {
        c <- "ping"
     }
}
func printer(c chan string) {
     for {
           msg := <- c
           fmt.Println(msg)
           time.Sleep(time.Second * 1)
     }
}
func main() {
     var c chan string = make(chan string)

     go pinger(c)
     go printer(c)

     var input string
     fmt.Scanln(&input)
}
```

A channel type is represented with the keyword `chan` followed by the type of
the things that are passed on the channel. The `<-` operator is used to send
and receive messages on the channel.

Using a channel like this synchronizes the two goroutines. When `pinger`
attempts to send a message on the channel, it will wait until `printer` is ready
to receive the message (blocking).

## Channel Direction
We can specify a direction on a channel type thus restricting it to either
sending or receiving.

```
func pinger(c chan<- string)
```

Now, `c` can only be sent to. Attempting to receive from c will result in a
compiler error.

A channel that doesn't have these restrictions is known as bi-directional. A
bi-directional channel can be passed to a function that takes send-only or
receive-only channels, but the reverse is not true.

## Select
Go has a special statement called `select` which works like a `switch` but for
channels. The `select` statement is often used to implement a timeout. The
default case happens immediately if none of the channels are ready.

## Buffered Channels
Possible to pass a second parameter to the make function when creating a
channel:

```
c := make(chan int, 1)
```

This creates a buffered channel with a capacity of 1. Normally channels are
synchronous; both sides of the channel will wait until the other side is ready.
A buffered channel is asynchronous; sending or receiving a message will not wait
unless the channel is already full.
