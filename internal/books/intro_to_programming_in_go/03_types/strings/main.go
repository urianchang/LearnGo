package main

import "fmt"

func main() {
  //: Strings
  fmt.Println(`string with backticks`)
  fmt.Println("string with double quotes")
  fmt.Println(len("test"))
  fmt.Println("Question"[0])  //: Character is represented by byte
  fmt.Println("Hello" + "World")
}
