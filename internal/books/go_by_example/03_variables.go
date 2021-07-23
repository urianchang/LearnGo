package main

import "fmt"

func main() {

	// Different ways of initializing variables
	var a = "a"				// type is inferred
	var b, c int = 1, 2		// type is declared
	var d float64			// zero-valued
	e := true				// shorthand to declare and initialize

	fmt.Println(a, b, c, d, e)
}