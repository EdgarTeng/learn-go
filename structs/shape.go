package structs

import "math"

// Rectangle is one of shapes
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle is one of shapes
type Circle struct {
	Radius float64
}

// Primeter is used for compute of shape's primeter
func Primeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Height) * 2
}

// Area of circle
func Area(circle Circle) float64 {
	return circle.Radius * circle.Radius * math.Pi
}
