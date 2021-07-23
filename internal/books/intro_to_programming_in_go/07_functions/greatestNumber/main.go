package main

import "fmt"


func greatestNumber(nums ...int) int {
	greatest := nums[0]
	for _, value := range nums {
		if value > greatest {
			greatest = value
		}
	}
	return greatest
}


func main() {
	arr := []int{ 5, 2, 1, 6, 7, 10, 42, 38 }
	fmt.Println(greatestNumber(arr...))
}

