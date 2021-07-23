package main

import "fmt"

func main() {
  //: var x string = "Hello World"
  var x string
  x = "first"
  fmt.Println(x)
  x += " second"
  fmt.Println(x)
  //: Short way of declaring and assigning variables
  a := "hello"
  b := "hello"
  c := "world"
  fmt.Println(a == b)
  fmt.Println(a == c)
  //: Declaring constant
  const d = "Hello World"
  //: Declaring multiple variables
  var (
    e = 1
    f = 2
    g = 3
  )
  fmt.Println(e+f+g)
}
