package circle

import (
	"fmt"
	"math"
)

/*
 * init function added
 */
func init() {
	fmt.Println("circle package initialized")
}

// Area function
func Area(radius float64) float64 {
	area := 2 * 3.14 * radius
	return area
}

// Perimeter function
func Perimeter(radius float64) float64 {
	perimeter := 3.14 * math.Pow(radius, 2)
	return perimeter
}
