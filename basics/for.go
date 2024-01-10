package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// go has only one looping construct, the for loop
// the basic for loop contains three compoentns seperated by semicolons:
// 1) init statement: executed before first iteration
// 2) condition experssion
// 3) post statement: executed at the end of every iteration

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// if statement can start with a short statement to execute before the condition
// variables declared in the statement are only in scope until the end of if.

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v

	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}
func printOneToNReverse(n int) {
	fmt.Println("--Print start--")
	for i := 0; i < n; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("--print end--")
}
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum) //45

	// the init and post statement are optional
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum) // 1440

	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Print("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS x.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("%s.\n", os)
	}
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Print("today")
	case today + 1:
		fmt.Print("tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}

	// switch without a condition is the same as switch true
	// this construct can be a clean way to write long if-then-else chains

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening")
	}
	// defer statment defers the execution of a function until the surrounding function returns

	defer fmt.Println("world")
	fmt.Println("hello")

	// hello
	//world

	// STACKING defers --> defered function calls are pushed onto a stack.
	// When a function returns, its deferred calls are executed in lifo
	printOneToNReverse(5) // --> prints 4 3 2 1 0
}
