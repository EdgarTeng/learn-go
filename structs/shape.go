package structs

import "math"

// Shape is a interface of all shapes
type Shape interface {
	Primeter() float64
	Area() float64
}

// Rectangle is one of shapes
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle is one of shapes
type Circle struct {
	Radius float64
}

// Primeter of Rectangle
func (r *Rectangle) Primeter() float64 {
	return (r.Width + r.Height) * 2
}

// Area of circle
func (c *Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
