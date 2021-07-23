package main  //: Refers to function name

import "fmt"

func main() {
  fmt.Println(true && true)
  fmt.Println(true && false)
  fmt.Println(false && false)
  fmt.Println(true || true)
  fmt.Println(true || false)
  fmt.Println(false || false)
  fmt.Println(!true)
}
