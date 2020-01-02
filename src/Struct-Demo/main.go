package main

import (
	"computer"
	"fmt"
)

// Employee struct
type Employee struct {
	firstName, lastName string
	age, salary         int
}

// Anonymous struct
type Anonymous struct {
	string
	int
}

// Address struct
type Address struct {
	city, state string
}

// Person struct
type Person struct {
	name    string
	age     int
	address Address
}

// PersonPromoted struct
type PersonPromoted struct {
	name string
	age  int
	Address
}

// Spec struct
type Spec struct { //exported struct
	Maker string //exported field
	model string //unexported field
	Price int    //exported field
}

// Name struct
type Name struct {
	firstName string
	lastName  string
}

// SpecExported type
type SpecExported computer.SpecExported

func main() {
	//creating structure using field names
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}
	fmt.Println("emp1: ", emp1)
	//creating structure without using field names
	emp2 := Employee{"Thomas", "Paul", 29, 800}
	fmt.Println("emp2: ", emp2)
	// Creating anonymous structures
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}
	fmt.Println("emp3: ", emp3)

	// Zero value of a structure
	var emp4 Employee //zero valued structure
	fmt.Println("Employee 4", emp4)

	// Accessing individual fields of a struct
	emp6 := Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: $%d", emp6.salary)

	// Anonymous fields
	ano := Anonymous{"Naveen", 50}
	fmt.Println(ano)
	ano.string = "naveen new"
	ano.int = 500
	fmt.Println(ano)

	// Nested structs
	var p Person
	p.name = "Naveen"
	p.age = 50
	p.address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("-------------Nested structs----------------")
	fmt.Println(p)
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.address.city)
	fmt.Println("State:", p.address.state)

	// Promoted fields
	var pPromoted PersonPromoted
	pPromoted.name = "Naveen"
	pPromoted.age = 50
	pPromoted.Address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("------------Promoted fields----------------")
	fmt.Println(p)
	fmt.Println("Name:", pPromoted.name)
	fmt.Println("Age:", pPromoted.age)
	fmt.Println("City:", pPromoted.city)   //city is promoted field
	fmt.Println("State:", pPromoted.state) //state is promoted field

	// Exported Structs and Fields
	fmt.Println("--------------Struct in package similar--------------")
	var spec Spec
	spec.Maker = "apple"
	spec.Price = 50000
	spec.model = "Mac Mini"
	fmt.Println("Spec:", spec)

	fmt.Println("--------------Struct in package similar--------------")
	var specE SpecExported
	specE.Maker = "apple"
	specE.Price = 50000
	// specE.model = "Mac Mini"
	fmt.Println("specE:", specE)

	// Structs Equality
	name1 := Name{"Steve", "Jobs"}
	name2 := Name{"Steve", "Jobs"}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := Name{firstName: "Steve", lastName: "Jobs"}
	name4 := Name{}
	name4.firstName = "Steve"
	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}

}
