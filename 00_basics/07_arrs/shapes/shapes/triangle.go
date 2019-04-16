package shapes

// Triangle struct
type Triangle struct {
	base   float64
	height float64
}

// NewTriangle create a new triangule
func NewTriangle(b float64, h float64) *Triangle {
	return &Triangle{base: b, height: h}
}

// Area computes the triangule total area
func (t *Triangle) Area() float64 {
	return (t.base * t.height) / 2
}
