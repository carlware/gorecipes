package shapes

// Shape any shape interface
type Shape interface {
	Area() float64
}

// ShapeArea calculate any shape area
func ShapeArea(s Shape) float64 {
	return s.Area()
}
