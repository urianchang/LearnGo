package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (l + w)
}

type Shape interface {
	area() float64
	perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalPerimeter(shapes ...Shape) float64 {
	var perimeter float64
	for _, s := range shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}

func main() {
	rec := Rectangle{x1: 0, y1: 0, x2: 10, y2: 10}
	cir := Circle{x: 0, y: 0, r: 5}

	fmt.Println("Area")
	fmt.Println(rec.area())      // 100
	fmt.Println(cir.area())      // 78.53981633974483

	fmt.Println("\nPerimeter")
	fmt.Println(rec.perimeter()) // 40
	fmt.Println(cir.perimeter()) // 31.41592653589793

	fmt.Println("\nTotal")
	fmt.Println(totalArea(&rec, &cir)) // 178.53981633974485
	fmt.Println(totalPerimeter(&rec, &cir))	// 71.41592653589794
}
