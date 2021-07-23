package main

import "fmt"

func main() {
  
  //: Example of a slice
  var x[]float64

  //: Make a slice - 2nd param is length, 3rd param is capacity
  y := make([]float64, 5, 10)

  //: Another way to create slices is to use [low : high]
  arr := []float64{1, 2, 3, 4, 5}
  z := arr[0:5]

  fmt.Println("initialize x, y, z:", x, y, z)

  //: Slice Function - append
  slice1 := []int{1, 2, 3}
  slice2 := append(slice1, 4, 5)

  fmt.Println("append... slice1, slice2:", slice1, slice2)

  //: Slice Function - copy
  slice3 := make([]int, 2)
  copy(slice3, slice1)

  fmt.Println("copy... slice1, slice3:", slice1, slice3)

}
