//geometry.go
package main

import (
	"Packages-Demo/packages-demo/geometry/circle"
	"Packages-Demo/packages-demo/geometry/rectangle"
	"Packages-Demo/packages-demo/geometry/square"
	"fmt"
)

func main() {
	var rectLen, rectWidth float64 = 6, 7
	var radius float64 = 5

	// rectangle
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	fmt.Printf("diagonal of the rectangle %.2f ", rectangle.Perimeter(rectLen, rectWidth))

	// Circle
	fmt.Printf("Area of circle %.2f\n", circle.Area(radius))
	fmt.Printf("Perimeter of the ciccle %.2f ", circle.Perimeter(radius))

	// Square
	fmt.Printf("area of square %.2f\n", square.Area(rectLen, rectWidth))
	fmt.Printf("Perimeter of the square %.2f ", square.Perimeter(rectLen))
}
