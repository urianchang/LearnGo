package main

import "fmt"

func swap(xPtr *int, yPtr *int) {
	temp := *xPtr
	*xPtr = *yPtr
	*yPtr = temp
}
func main() {
	fmt.Print("Enter integer for x: ")
	var x int
	fmt.Scanln(&x)

	fmt.Print("Enter integer for y: ")
	var y int
	fmt.Scanln(&y)

	swap(&x, &y)
	fmt.Println(x, y)
}
