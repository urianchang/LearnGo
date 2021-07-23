package main

import "fmt"

func main() {
  //: Two ways of writing for-loop (only type of loop in Go)
  i := 1
  for i <= 10 {
    fmt.Println(i)
    i += 1
  }

  for i := 1; i <= 10; i++ {
    if i%2 == 0 {
      fmt.Println(i, "even")
    } else {
      fmt.Println(i, "odd")
    }
  }

  //: Switch cases
  for i := 0; i <= 4; i++ {
    switch i {
      case 0: fmt.Println("Zero")
      case 1: fmt.Println("One")
      case 2: fmt.Println("Two")
      case 3: fmt.Println("Three")
      default: fmt.Println("Unknown number")
    }
  }
}
