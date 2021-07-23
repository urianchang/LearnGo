# Pointers
When we call a function that takes an argument, that argument is copied to the
function.

```
func zero(x int) {
    x=0
}
func main() {
     x := 5
     zero(x)
     fmt.Println(x) // x is still 5
}
```

To modify the original variable `x`, we'll need to use a pointer.

```
func zero(xPtr *int) {
     *xPtr = 0
}
func main() {
     x := 5
     zero(&x)
     fmt.Println(x) // x is 0
}
```

A pointer references a location in memory where a value is stored rather than
the value itself. By using a pointer `*int`, the `zero` function is able to
modify the original variable.

## * and & operators
A pointer is represented using the `*` character followed by the type of the
stored value. `*` is also used to "dereference" pointer variables, giving access
to the value the pointer points to.

The `&` operator is used to find the address of a variable. `&x` returns a
`*int` because `x` is an `int`.

## new
Another way to get a pointer is to use `new` function:

```
func one(xPtr *int) {
     *xPtr = 1
}
func main() {
     xPtr := new(int)
     one(xPtr)
     fmt.Println(*xPtr) // x is 1
}
```

`new` takes a type as an argument, allocates enough memory to fit a value of
that type and returns a pointer to it. In some programming languages, there is a
significant difference between using `new` and `&`, with great care being needed
to delete anything created with `new`. Because Go is a garbage collected
language, memory is cleaned up automatically when nothing refers to it anymore.

Pointers are rarely used with Go's built-in types, but they are extrememly
useful when paired with structs.
