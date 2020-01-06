package main

import (
	goroutines "Concurrency-Demo/Goroutines"
	"fmt"
	"time"
)

func main() {
	go goroutines.Hello()
	go goroutines.Numbers()
	go goroutines.Alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main function")
}
