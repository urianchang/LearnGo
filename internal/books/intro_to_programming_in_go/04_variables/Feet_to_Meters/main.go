package main

import "fmt"

func main() {
	fmt.Print("Enter feet: ")
	var feet float64
	fmt.Scanf("%f", &feet)
	meters := feet * 0.3048
	output_str := fmt.Sprintf("%.4f", meters)
	fmt.Printf("Meters: %s\n", output_str)
}