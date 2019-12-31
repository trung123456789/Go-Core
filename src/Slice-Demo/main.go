package main

import "fmt"

func main() {
	fmt.Println("---------------Declare Slice------------------")

	var mySlice1 = make([]int, 4, 7)
	fmt.Println("len: ", len(mySlice1), "cap: ", cap(mySlice1))
	fmt.Println(mySlice1)

	var mySlice2 = make([]int, 7)
	fmt.Println("len: ", len(mySlice2), "cap: ", cap(mySlice2))

	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] //creates a slice from a[1] to a[3] : [77 78 79]
	fmt.Println(b)

	c := []int{6, 7, 8, 8, 9} //creates and array and returns a slice reference: [6 7 8 8 9]
	fmt.Println(c)

	fmt.Println("---------------Modify Slice------------------")
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)
	fmt.Println("****************************************")
	numa := [3]int{78, 79, 80}
	nums1 := numa[:] //creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)

	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
	fruitslice = fruitslice[:cap(fruitslice)]                                        //re-slicing furitslice till its capacity
	fmt.Println("After re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))

	fmt.Println("---------------Get value from slice and not chage value original slice------------------")
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	countriesCpy[1] = "trung"
	return countriesCpy
}
