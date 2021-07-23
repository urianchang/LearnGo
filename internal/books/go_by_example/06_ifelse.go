package main

import "fmt"

func main() {
	if 1+1 == 2 {
		fmt.Println("1+1 is equal to 2")
	}

	// A statement can precede conditionals
	if a := 1; a < 1 {
		fmt.Println("a is less than 1")
	} else if a > 1 {
		fmt.Println("a is greater than 1")
	} else {
		fmt.Println("a is equal to 1")
	}

	// NOTE: There is no ternary if, so full if/else will need to be used.
}
