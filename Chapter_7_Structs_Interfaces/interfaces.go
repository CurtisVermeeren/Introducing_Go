package main

import (
		"fmt"
		"math"
)

type Shape interface {
	// Define a method set that a type must have to implement the interface
	area() float64
}

// Create a circle struct with 3 floats
type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

// Create a methof for type Circle to calculate area
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// Take 0 or more shapes using the interface and adds their areas
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main(){
	c := Circle{x:10,y:5,r:2}
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(totalArea(&c, &r))
}
