package main

import "fmt"

// Go functions can be written to work on multipe types using type parameters.
// the type parameters of a function appear between barckets, before the functions arguments

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

// in addition to generic functions, Go also supports genenric types.
// A type can be parameterized with a type parameter, which could be useful for implementing generic data structures.

type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	//Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15)) //2

	//Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello")) //-1

}
