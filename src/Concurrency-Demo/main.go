package main

import (
	channels "Concurrency-Demo/Channels"
	goroutines "Concurrency-Demo/Goroutines"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("------------Sstart Goroutine--------------")
	go goroutines.Hello()
	go goroutines.Numbers()
	go goroutines.Alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("------------Terminated Goroutine--------------")

	fmt.Println("------------Start Channel--------------")
	fmt.Println("------------Declare Channel--------------")
	channels.DeclareChannel()

	fmt.Println("-----------Simple example Channel--------------")
	done := make(chan bool)
	go channels.Hello(done)
	abc := <-done
	fmt.Println(abc)

	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go channels.CalcSquares(number, sqrch)
	go channels.CalcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println(squares)
	fmt.Println(cubes)
	fmt.Println("Final output", squares+cubes)

	fmt.Println("-----Closing channels and for range loops on channels-------")
	ch := make(chan int)
	go channels.Producer(ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}

	fmt.Println("-----Buffered Channels-------")
	chBuf := make(chan int, 2)
	go channels.Write(chBuf)
	time.Sleep(2 * time.Second)
	for v := range chBuf {
		fmt.Println("read value", v, "from chBuf")
		time.Sleep(2 * time.Second)

	}

	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)

}

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}
