package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// we can declare a method on non-struct types, to
// In this example we see a numeric type Myfloat with Abs method

type MyFloat float64

func (f MyFloat) Abs() float64 {
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
// A value of interface type can hold any value that implements those methods.

type Abser interface {
	Abs() float64
}
type I interface {
	M()
}
type T struct {
	S string
}

// this method means that type T implements the interface I
// but we dont need to explicitly declare that it does so
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

// Under the hood , interface values can be thought of as  a tuple of a value and a concrete type:
// (value , type)
// a nil interface value holds neither value nor concrete type
func (f MyFloat) M() {
	fmt.Println(f)
}
func describe(i interface{}) {
	fmt.Printf("(%v,%T)\n", i, i)
}
func main() {
	v := Vertex{3, 4}
	f := MyFloat(-math.Sqrt2)
	fmt.Println(v.Abs())
	fmt.Println(f.Abs())
	v.Scale(2)
	fmt.Println(v)

	var a Abser
	a = f
	a = &v
	a = v
	fmt.Println(a.Abs())
	var i I
	i = &T{"Hello"}
	i.M()
	i = MyFloat(math.Pi)
	describe(i)
	i.M()

	// the interface that specifies zero methids is known as empty interface

	var emptyInterface interface{}
	emptyInterface = 42
	describe(emptyInterface) //(42,int)

	emptyInterface = "hello"
	describe(emptyInterface) // (hello,string)
}
