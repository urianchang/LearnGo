package main
import "fmt"

/* ====
  Variadic Functions:
    There is a special form available for the last parameter in a Go
    function. By using '...' before the type name of the last parameter,
    you can indicate that it takes zero or more of those parameters.

==== */

func add(args ...int) int {
  total := 0
  for _, v := range args {
    total += v
  }
  return total
}

func main() {
  fmt.Println(add(1, 2, 3))
}

/* ====
  This is how fmt.Println function is implemented:
    func Println(a ...interface{}) (n int, err error)
==== */

/* ====
  We can also pass in a slice of int's:
    func main() {
      xs := []int{ 1, 2, 3 }
      fmt.Println(add(xs...))
    }
==== */
