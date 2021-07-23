package main

import "fmt"

func half(num int) (int, bool) {
	halved := num / 2
	isEven := true
	if (num % 2) != 0 {
		isEven = false
	}
	return halved, isEven
}

func main() {
	fmt.Print("Enter integer: ")
	var input int
	fmt.Scanf("%d", &input)
	fmt.Println(half(input))
}
