package basics

import "fmt"

func add(x int, y int) int {
	return x + y
}

// when two or more consecutive named function parameters share a type
//
//	you can omit the type from all but the last
//
// here x int, y int --> x,y int
func add1(x, y int) int {
	return x + y
}

//functions can return any number of results
// this function returns two strings

func swap(x, y string) (string, string) {
	return y, x
}

//naked returns --> in this case where the the return statement without
// argunments ruturns the named return values x,y

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// var statement declares a list of variables; as in fucntion argument lists, the type is last
// var can be used both at function and package level
var c, python, java bool

// var declarations can include intializers , one per variable
// if an intializer is present, the type can be omitted; the variable will take the type of the initializer

var j, k int = 1, 2

// constants are declared like variables but with the const keyword
// constants can be character , string , boolean, or numeric values
// constants cannot be declared using the := syntax
const Pi = 3.14

func Basic() {
	fmt.Println(add(10, 7))   //17
	fmt.Println((add(10, 7))) //17
	a, b := swap("hello", "World")
	fmt.Println(a, b)      // "World hello"
	fmt.Println(split(17)) //7,10;

	var i int
	fmt.Println(i, c, python, java) //0 false false false

	var cpp, js, golang = true, false, "Yessss"
	fmt.Println(j, k, cpp, js, golang) // 1 2 true false Yessss

	const Truth = true
	fmt.Println(Pi, Truth) // 3.14 true
}
