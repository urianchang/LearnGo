# Functions
A function is an independent section of code that maps zero or more input
parameters to zero or more output parameters.

Functions start with the keyword `func`, followed by the function's name. The
parameters (inputs) of the function are defined like this:
`name type, name type, ...`. After the parameters, we put the return type.

```
func average(xs []float64) float64 {
    panic("Not Implemented")
}
```

Functions are built up in a "stack". Each time we call a function we push it
onto the call stack and each time we return from a function we pop the last
function off of the stack.

Go is capable of returning multiple values from a function. Multiple values are
often used to return an error value along with the result, or a boolean to
indicate success.

## Variadic Functions
By using `...` before the type name of the last parameter, you can indicate that
it takes zero or more of those parameters.

```
func add(args ...int) int {
     total := 0
     for _, v := range args {
           total += v
     }
     return total
}
func main() {
     fmt.Println(add(1,2,3))
}
```

We can also pass a slice by following the slice with `...`:

```
func main() {
     xs := []int{1,2,3}
     fmt.Println(add(xs...))
}
```

## Closure
It is possible to create functions inside of functions. When you create a local
function, it has access to other local variables.

One way to use closure is by writing a function which returns another function.

```
func makeEvenGenerator() func() uint {
     i := uint(0)
     return func() (ret uint) {
           ret = i
           i += 2
           return
     }
}
func main() {
     nextEven := makeEvenGenerator()
     fmt.Println(nextEven()) // 0
     fmt.Println(nextEven()) // 2
     fmt.Println(nextEven()) // 4
}
```

## Recursion
A function can call itself.

```
func factorial(x uint) uint {
     if x == 0 {
        return 1
     }
     return x * factorial(x-1)
}
```

## Defer
`defer` schedules a function call to be run after the function completes.

```
package main

import "fmt"

func first() {
     fmt.Println("1st")
}
func second() {
     fmt.Println("2nd")
}
func main() {
     defer second()
     first()
}
```

`defer` is often used when resources need to be freed in some way. For example,
when we open a file, we need to make sure to close it later.

```
f, _ := os.Open(filename)
defer f.Close()
```

This has 3 advantages:

1. Keeps the `Close` call near `Open` so it's easier to understand
2. If our function had multiple return statements, `Close` will happen before
both of them
3. Deferred functions are run even if a run-time panic occurs

## Panic & Recover
`panic` causes a run-time error. We can handle a run-time panic with the
`recover` function: it stops the panic and returns the value that was passed to
the call to `panic`.

```
package main

import "fmt"

func main() {
     defer func() {
           str := recover()
           fmt.Println(str)
     }()
     panic("PANIC")
}
```

A panic generally indicates a programmer error or an exceptional condition that
there's no easy way to recover from.
