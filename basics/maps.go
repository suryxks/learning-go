package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var map2 = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// in go functions or values too. They can be passed aroud like jsut other values.
// function values may be used as function arguments and return values.
func main() {
	m = make(map[string]Vertex)
	m["Bell Lab"] = Vertex{40.68433, -74.39967}
	fmt.Println(m["Bell Lab"])
	fmt.Println(map2)
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(math.Pow))
	fmt.Println(compute(hypot))

}
