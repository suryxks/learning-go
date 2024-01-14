package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, 2*v)
	case string:
		fmt.Printf("%q is %v bytes long \n", v, len(v))
	default:
		fmt.Printf("I dont know about type %T! \n", v)
	}
}

type Person struct {
	Name string
	Age  int
}

// one of the most common interfaces is Stringer defined by fmt package
// type Stringer interface{
// 	String() string
// }

// A stringer is a type that can describe itself as a string.
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	var i interface{} = "hello"
	// a type assertion provides access to an interface value's underlying concrete value
	// t:=i.(T)
	//If i does not hold a T, the statement will trigger a panic.

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // panic
	fmt.Println(f)

	// A type switch is a construct that permits several type assertions in series.
	do(21)
	do("hello")
	do(true)

	a := Person{"Surya", 22}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
