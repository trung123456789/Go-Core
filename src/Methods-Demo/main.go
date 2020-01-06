package main

import (
	"fmt"
)

// Employee struct
type Employee struct {
	name     string
	salary   int
	currency string
}

type rectangle struct {
	length int
	width  int
}

func main() {
	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalary() //Calling displaySalary() method of Employee type
	displaySalary(emp1)

	// Pointer Receivers vs Value Receivers
	emp2 := Employee{
		name:     "Mark Andrew",
		salary:   5000,
		currency: "$",
	}
	fmt.Printf("Employee name before change: %s", emp2.name)
	emp2.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", emp2.name)

	fmt.Printf("\n\nEmployee age before change: %d", emp2.salary)
	(&emp2).changeSalary(5001)
	fmt.Printf("\nEmployee age after change: %d", emp2.salary)

	r := rectangle{
		length: 10,
		width:  5,
	}
	p := &r //pointer to r
	perimeter(p)
	p.perimeter()

	/*
	   cannot use r (type rectangle) as type *rectangle in argument to perimeter
	*/
	// perimeter(r)

	r.perimeter() //calling pointer receiver with a value

}

func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func displaySalary(e Employee) {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

/*
Method with value receiver
*/
func (e Employee) changeName(newName string) {
	e.name = newName
}

/*
Method with pointer receiver
*/
func (e *Employee) changeSalary(newSalary int) {
	e.salary = newSalary
}

func perimeter(r *rectangle) {
	fmt.Println("perimeter function output:", 2*(r.length+r.width))

}

func (r *rectangle) perimeter() {
	fmt.Println("perimeter method output:", 2*(r.length+r.width))
}
