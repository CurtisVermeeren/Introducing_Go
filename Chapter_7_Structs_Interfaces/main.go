package main
import ("fmt"; "math")

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

func main() {
	// Create a new circleArea
	// Name fields aren't required if order is known
	c := Circle{x:10,y:5,r:2}

	// Access fields
	fmt.Println(c.x,c.y,c.r)

	// Get the area of a Circle
	fmt.Println(c.area())

	// Get the area of a Rectangle
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

}
