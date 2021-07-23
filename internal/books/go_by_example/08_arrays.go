package main

import "fmt"

func main() {
	// By default, an array is zero-valued
	var a [6]float64
	fmt.Println("emp: ", a)

	// Get, set, and list
	a[4] = 1000
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("size:", len(a))

	// Declare and initialize
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b:", b)

	// [row][column]
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i+j
		}
	}
	fmt.Println("2D:", twoD)
}