package main

import "fmt"

//: Take in a slice of float64's and return one float64
  //: Params are defined like this: name type, name type, ...
func average(xs []float64) float64 {
  // panic("Not implemented")  //: panic is a built-in function that causes a run time error

  //: Imported from main function
  total := 0.0
  for _, v := range xs {
    total += v
  }

  return total / float64(len(xs))
}

//: Computes average of a series of numbers
func main() {
  someOtherName := []float64{ 98, 93, 77, 82, 83 }

  //: Imported to separate function (above)
  // total := 0.0
  // for _, v := range xs {
  //   total += v
  // }
  //
  // fmt.Println(total / float64(len(xs)))

  fmt.Println(average(someOtherName))

  /* ====
    Functions don't have access to anything in the calling function
        func f() {
          fmt.Println(x)
        }
        func main() {
          x := 5
          f()
        }
  ==== */

  /* ====
    Functions are built up in a "stack"
        func main() {
          fmt.Println(f1())
        }
        func f1() int {
          return f2()
        }
        func f2() int {
          return 1
        }
    Each time we call a function, we push it onto the call stack and
    each time we return from a function we pop the last function off
    the stack.
  ==== */

  /* ====
    We can name the return type:
        func f2() (r int) {
          r = 1
          return
        }
  ==== */

}
