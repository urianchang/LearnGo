package main

import (
	"fmt"
	"math"
)

const STRING string = "cannot overwrite"

func main() {
	fmt.Println(STRING)

	const num = 5000
	const val = 5e10 / num
	fmt.Println(val)	// arbitrary precision

	// numeric constant has no type until given one
	fmt.Println(int32(val))
	fmt.Println(math.Cos(val))
}