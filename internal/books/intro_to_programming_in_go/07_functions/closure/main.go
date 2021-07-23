package main
import "fmt"

/* ====
  One way to use closure is by writing a function which returns
  another function which - when called - can generate a sequence
  of numbers.
==== */

func makeEvenGenerator() func() uint {
  i := uint(0)
  return func() (ret uint) {
    ret = i
    i += 2
    return
  }
}

func main() {
  //: It is possible to create functions within functions
  // add := func(x, y int) int {
  //   return x + y
  // }
  // fmt.Println(add(1, 1))

  //: Local functions have access to local variables -- a closure!
  // z := 0
  // increment := func() int {
  //   z++
  //   return z
  // }
  // fmt.Println(increment())
  // fmt.Println(increment())

  //:Generate even numbers with closures
  nextEven := makeEvenGenerator()
  fmt.Println(nextEven()) // 0
  fmt.Println(nextEven()) // 2
  fmt.Println(nextEven()) // 4
}
