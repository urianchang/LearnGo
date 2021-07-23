package main

import (
	"fmt"
	"github.com/urianchang/LearnGo/internal/books/intro_to_programming_in_go/11_packages/math"
)

func main() {
    xs := []float64{1, 2, 3, 4}
    fmt.Println(math.Average(xs))
    fmt.Println(math.Min(xs))
    fmt.Println(math.Max(xs))
}
