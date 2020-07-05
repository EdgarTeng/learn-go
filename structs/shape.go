package structs

// Rectangle is one of shapes
type Rectangle struct{
	Width float64
	Height float64
}

// Primeter is used for compute of shape's primeter
func Primeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Height) * 2
}
