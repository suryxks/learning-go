package main

import (
	"fmt"
)

// A struct is a collection of fields

type Vertex struct {
	X int
	Y int
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	i, j := 42, 2701
	p := &i         // pointer to i
	fmt.Println(*p) // read i through the ponter
	*p = 21         // set i through the pointer

	fmt.Println(i) // see the new value 21
	p = &j         //point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) //see the new value of j

	// --------STRUCTS-----------------

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X) //4
	p1 := &v
	// struct fields can be accessed through a struct pointer.
	// accessing X here can be done with p.X instead of (*p).X

	fmt.Println(p1.X == v.X) // true
	p1.X = 1e9
	fmt.Println(v) //{1000000000 2}

	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p2 = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, p2, v2, v3) // {1 2} &{1 2} {1 0} {0 0}
	// -----------------Arrays--------------------------------

	// [n]T is an array of n values of type T

	var a [2]string
	a[0] = "HEllo"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Array --> fixed size , Slice ---> dynamically sized
	var s []int = primes[1:4] // from index 1 to 3
	fmt.Println(s)            // [3,5,7]

	// A slice does not store any data, it just describes a section of an underlying array.
	// changing the elements of a slice modifies the corresponding elements of its underlying array.
	// other slices that share the same underlying array will see those changes.

	names := [4]string{
		"Surya",
		"Aman",
		"sherwin",
		"aravindan",
	}
	fmt.Println(names)   //["Surya","Aman","sherwin","aravindan"]
	slice1 := names[0:2] //["Surya","Aman"]
	slice2 := names[1:3] //["Aman","sherwin"]

	fmt.Println(slice1, slice2)
	slice2[0] = "ASDF"
	fmt.Println(slice1, slice2)
	fmt.Println(names)

	//----slice literals
	// similar to array literals but without the length.

	//[2]bool{true,false} ---> array literal
	//[]bool {true,true}  ----> slice literal

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	sz := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{3, false},
		{4, false},
	}
	fmt.Println(sz)
	// when slicing you may omit the high or low bounds to use their defaults instead.

	// the default is 0 for low bound and the length of the slice for the highbound

	// a slice has both length and capacity.
	// length---> number of elements it contains use -> len(s)
	// capacity ----> number of elements in the underlying array coundting from the first element in slice use -> cap(s).

	sample := []int{2, 3, 5, 7, 11, 13}
	sample = sample[1:4] // [3,5,7]
	fmt.Println(sample)
	printSlice(sample) // len 3 cap 5

	sample = sample[:2] // [3,5]
	fmt.Println(sample)
	printSlice(sample)

	sample = sample[1:] //[5]
	fmt.Println(sample)
	printSlice(sample)

	// slices can be created using built in make function

	sliceCreatedUsingMake := make([]int, 5)
	fmt.Println(sliceCreatedUsingMake)

	// range form of for loop iterates over a slice or map

	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d=%d\n", i, v)
	}
}
