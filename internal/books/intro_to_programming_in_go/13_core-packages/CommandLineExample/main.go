package main

import (
	"flag"
	"fmt"
	"math/rand"
)

// To run: `go run main.go -max=100`
func main() {
	// Define flags
	maxp := flag.Int("max", 6, "the max value")

	// Parse
	flag.Parse()
	// Any additional non-flag arguments
	// Use flag.Args(), which returns []string

	// Generate a number between 0 and max
	fmt.Println(rand.Intn(*maxp))
}
