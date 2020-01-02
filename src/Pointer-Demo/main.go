package main

import "fmt"

func main() {
	// Declaring pointers
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)

	// Zero value of a pointer
	zeroV := 25
	var pZero *int
	if pZero == nil {
		fmt.Println("pZero is", pZero)
		pZero = &zeroV
		fmt.Println("pZero after initialization is", pZero)
	}

	//Creating pointers using the new function
	pSize := new(int)
	fmt.Printf("pSize value is %d, type is %T, address is %v\n", *pSize, pSize, pSize)
	*pSize = 85
	fmt.Println("New pSize value is", *pSize)

	// Dereferencing a pointer
	increase := 255
	pIncrease := &increase
	fmt.Println("address of increase is", pIncrease)
	fmt.Println("value of increase is", *pIncrease)
	*pIncrease++
	fmt.Println("new value of increase is", increase)

	// Passing pointer to a function
	before := 58 // before: 58
	p := &before
	change(p) // before: 55

	// Returning pointer from a function
	iPointer := reutrnPointer()
	fmt.Println("Value of iPointer", *iPointer)

	// Do not pass a pointer to an array as a argument to a function. Use slice instead.
	arr := [3]int{89, 90, 91}
	modify(&arr)
	fmt.Println(arr)

	// Go does not support pointer arithmetic
	// bArr := [...]int{109, 110, 111}
	// pArr := &bArr
	// pArr++
}

func change(val *int) {
	*val = 55
}

func reutrnPointer() *int {
	i := 5
	return &i
}

func modify(arr *[3]int) {
	arr[0] = 90
}
