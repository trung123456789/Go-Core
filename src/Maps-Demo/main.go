package main

import "fmt"

func main() {
	// Declare map
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	// Use variable as key
	employee := "jamie"
	fmt.Println("Salary of", employee, "is", personSalary[employee]) // Salary of jamie is 15000
	// Add items
	personSalary["mike"] = 9000 // map[jamie:15000 mike:9000 steve:12000]
	// Deleting items
	delete(personSalary, "steve") // map[jamie:15000 mike:9000]
	// Length of map
	fmt.Println("length is", len(personSalary))
	// Maps are reference types
	newPersonSalary := personSalary
	newPersonSalary["mike"] = 18000
}
