package main

import "fmt"

func main() {
  x := make(map[string]int)

  x["key"] = 10
  x["key2"] = 11

  fmt.Println(x["key"])   //: maps are not sequential
  fmt.Println(len(x))   //: Length of map can change

  //: Delete: (map, key)
  delete(x, "key")
}
