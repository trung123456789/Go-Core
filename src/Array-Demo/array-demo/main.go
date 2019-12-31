package main

import "fmt"

func main() {
	fmt.Println("---------------Declare variable-------------------")
	var a [3]int // => [0 0 0]
	a[0] = 12    // => [12 0 0]
	fmt.Println("length of a: ", len(a))
	fmt.Println("a: ", a)
	b := [3]int{12, 78, 50} // => [12 78 50]
	fmt.Println("b: ", b)
	c := [3]int{12} // => [12 0 0]

	fmt.Println("c :", c)
	d := [...]int{12, 78, 50, 40} // => [12 78 50, 40]
	fmt.Println("d: ", d)

	fmt.Println("---------------Add value to array-------------------")
	arrAppend := []int{}
	arrAppend = append(arrAppend, 11)
	arrAppend = append(arrAppend, 11)
	fmt.Println(arrAppend)

	fmt.Println("---------------Assign variable-------------------")
	// arr1: [USA China India Germany France]
	arr1 := [...]string{"USA", "China", "India", "Germany", "France"}
	fmt.Println("arr1: ", arr1)
	// arr2: [Singapore China India Germany France]
	arr2 := arr1 // a copy of a is assigned to b
	fmt.Println("arr2: ", arr2)

	fmt.Println("---------------User array in function-------------------")
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function: ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function: ", num)

	fmt.Println("---------------Loop array with len-------------------")
	sum := 0
	for i := 0; i < len(num); i++ { //looping from 0 to the length of the array
		fmt.Printf("%d th element of a is %d\n", i, num[i])
		sum += num[i]
	}
	fmt.Println("sum: ", sum)
	sum = 0
	fmt.Println("---------------Loop array with range-------------------")
	for i, v := range num { //range returns both the index and value
		fmt.Printf("%d the element of a is %d\n", i, v)
		sum += v
	}
	fmt.Println("sum: ", sum)
	sum = 0
	fmt.Println("---------------ignore the index or value-------------------")
	for _, v := range num {
		sum += v
	}
	fmt.Println("sum: ", sum)

	fmt.Println("---------------Multidimensional arrays-------------------")
	multiArr1 := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	}
	printarray(multiArr1)
	var multiArr2 [3][2]string
	multiArr2[0][0] = "apple"
	multiArr2[0][1] = "samsung"
	multiArr2[1][0] = "microsoft"
	multiArr2[1][1] = "google"
	multiArr2[2][0] = "AT&T"
	multiArr2[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printarray(multiArr2)
}

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function: ", num)

}

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}
