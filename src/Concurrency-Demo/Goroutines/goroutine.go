package goroutines

import (
	"fmt"
	"time"
)

// Hello function
func Hello() {
	fmt.Println("Hello world goroutine")
}

// Numbers function
func Numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

// Alphabets function
func Alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
