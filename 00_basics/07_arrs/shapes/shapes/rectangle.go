package shapes

// Rectangle struct
type Rectangle struct {
	width  float64
	height float64
}

// NewRectangle create a rectangle
func NewRectangle(w float64, h float64) *Rectangle {
	return &Rectangle{width: w, height: h}

}

// Area computes area
func (r *Rectangle) Area() float64 {
	return r.width * r.height
}
