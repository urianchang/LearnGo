package main
import "fmt"

//: Change return type to contain multiple types
func f() (int, int) {
  //: Change expression after return to contain multiple expressions
  return 5, 6
}

func main() {
  //: Change assignment statement so multiple values are on left side
  x, y := f()
}

/* ====
  Multiple values are often used to return an error value along with
  the result (x, err := f()), or a boolean to indicate success
  (x, ok := f()).
==== */
