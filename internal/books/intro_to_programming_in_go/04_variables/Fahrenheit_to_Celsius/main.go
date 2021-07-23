package main

import "fmt"

func main() {
	fmt.Print("Enter Fahrenheit: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := (input - 32) * 5/9
	output_str := fmt.Sprintf("%.2f", output)
	fmt.Printf("Celsius: %s\n", output_str)
}
