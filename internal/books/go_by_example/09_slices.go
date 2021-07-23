package main

import "fmt"

// More info: https://blog.golang.org/slices-intro
func main() {
	// Slices are typed only by the elements (not number of elements)
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// set and get like arrays
	s[0] = "1"
	s[1] = "2"
	s[2] = "3"
	fmt.Println("set", s)
	fmt.Println("get", s[2])
	fmt.Println("len", len(s))

	// slices support append, which return a new slice
	s = append(s, "a")
	s = append(s, "b")
	fmt.Println("append", s)

	// copy
	newSlice := make([]string, len(s))
	copy(newSlice, s)
	fmt.Println("copy", newSlice)

	// slicing a slice
	fmt.Println("slice", newSlice[2:5])

	// declare and initialize in a single line
	t := []string{"d", "e", "f"}
	fmt.Println("declare+init", t)

	// slices can be composed of multi-dimensional data;
	// length of inner slices can vary
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D", twoD)
}
