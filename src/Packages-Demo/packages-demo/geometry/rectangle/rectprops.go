package rectangle

import (
	"fmt"
	"math"
)

/*
 * init function added
 */
func init() {
	fmt.Println("rectangle package initialized")
}
func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

func Perimeter(len, wid float64) float64 {
	perimeter := math.Sqrt((len * len) + (wid * wid))
	return perimeter
}
