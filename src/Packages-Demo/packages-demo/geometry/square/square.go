//square.go
package square

import "fmt"

/*
 * init function added
 */
func init() {
	fmt.Println("square package initialized")
}
func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

func Perimeter(len float64) float64 {
	perimeter := len * 4
	return perimeter
}
