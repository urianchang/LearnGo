# Control Structures

## For
The `for` statement allows us to repeat a list of statements (a block) multiple
times.

```
package main

import "fmt"

func main() {
    i := 1
    for i <= 10 {
        fmt.Println(i)
        i = i + 1
    }
}
```

Go only has one type of loop that can be used in a variety of different ways.
The `for` loop can also be written like this:

```
func main() {
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
}
```

## If
An `if` statement is similar to a `for` statement in that it has a condition
followed by a block. If statements also have optional `else if` and `else`
parts.

```
if i % 2 == 0 {
    // divisible by 2
} else if i % 3 == 0 {
    // divisible by 3
} else if i % 4 == 0 {
    // divisible by 4
}
```

## Switch
Starts with the keyword `switch` followed by an expression and then a series of
`case`s. The value of the expression is compared to the expression following
each `case` keyword.

```
switch i {
    case 0: fmt.Println("Zero")
    case 1: fmt.Println("One")
    case 2: fmt.Println("Two")
    case 3: fmt.Println("Three")
    case 4: fmt.Println("Four")
    case 5: fmt.Println("Five")
    default: fmt.Println("Unknown Number")
}
```
