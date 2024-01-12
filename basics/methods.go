package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// we can declare a method on non-struct types, to
// In this example we see a numeric type Myfloat with Abs2 method

type MyFloat float64

func (f MyFloat) Abs2() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// method is  function with special receiver argument.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// methods can also be declared with pointer receivers.
// this means that receiver type has literal syntax *T for some type T.
// T cannot be itself be a pointer such as *int.

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//There are two reasons to use a pointer receiver
// 1) the method can modify the value that its receiver points to.
// 2) to avoid copying the value on each method call. This can be more efficient if the reciver is a large struct.

func (v *Vertex) absp() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// An interface type is defined as a set of method signatures.
func main() {
	v := Vertex{3, 4}
	f := MyFloat(-math.Sqrt2)
	fmt.Println(v.Abs())
	fmt.Println(f.Abs2())
	v.Scale(2)
	fmt.Println(v)
}
