package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// A goroutine is  a lightweight thread managed by Go runtime

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // closes the chanel c
}

// select statement lets a goroutine wait on multiple communication operations.

// A select blocks until one of its cases can run,then it executes that case. it chooses one at random if multiple are ready

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}

}

func main() {
	// go say("world")
	// say("hello")

	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)

	//channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <-   this will overflow the buffer and create a deadlock
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	d := make(chan int, 10)
	fmt.Println(cap(d))
	go fibonacci(cap(d), d)
	for i := range d {
		fmt.Println(i)
	}

	channel := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)

		}
		quit <- 0
	}()
	fibonacci2(c, quit)

}
