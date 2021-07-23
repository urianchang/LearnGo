package main

import "fmt"

func main() {

  //: Create array of 5 floats (shorter syntax)
  x := [5]float64{ 98, 93, 77, 82, 83 }

  //: Create variable to keep track of total
  var total float64 = 0

  //: Add up all the variables
  //: '_' is current position in array and 'value' is x[i]
  for _, value := range x {
    total += value
  }

  //: Print average - note the type conversion
  fmt.Println(total / float64(len(x)))
}
