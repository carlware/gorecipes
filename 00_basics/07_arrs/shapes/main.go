package main

import (
	"fmt"
	"shapes/shapes/shapes"
)

func main() {

	r := shapes.NewRectangle(9, 6)
	t := shapes.NewTriangle(3, 6)

	fmt.Println("Area of myRectangle: ", shapes.ShapeArea(r))
	fmt.Println("Area of myTriangle: ", shapes.ShapeArea(t))

}
