package main

import "fmt"

func main() {
	// `for` is the only looping construct

	// 1: single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// 2: intial / condition / after
	for j := 7; j <= 9; j++ {
		if j == 8 {
			continue
		}
		fmt.Println(j)
	}

	// 3: for without a condition will loop until `break` or `return`
	for {
		fmt.Println("to infinity and beyond!")
		break
	}

}