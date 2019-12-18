package main

import (
	"fmt"
	"math"
)

type circle struct {
	Radius float64
}

type rectangle struct {
	Side float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return c.Radius * c.Radius
}

func (r rectangle) area() float64 {
	return math.Pi * r.Side
}

func totalArea(shapes ...shape) float64 {
	var area float64
	for _, s := range shapes {
		fmt.Printf("%T\n", s)
		area += s.area()
	}
	return area
}

func main() {
	c := circle{1}
	r := rectangle{4}
	fmt.Println(totalArea(&c, &r))
}

// type animal struct {
// 	sound string
// }

// type dog struct {
// 	animal
// }

// type cat struct {
// 	animal
// }

// // Empty interfaces
// type animals interface{}

// func main() {
// 	dog := dog{animal{"woof"}}
// 	cat := cat{animal{"meow"}}

// 	// without empty interface type
// 	// animals := []interface{}{dog, cat}
// 	// or
// 	farm := []animals{dog, cat}

// 	for _, a := range farm {
// 		printValue(a)
// 	}
// }

// // Any value
// func printValue(v interface{}) {
// 	fmt.Println(v)
// }
