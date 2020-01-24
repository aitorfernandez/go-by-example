package main

import "fmt"

type position struct {
	X, Y int
}

func main() {
	// has type position
	posA := position{1, 2}
	// Y:0 is implicit
	posB := position{X: 1}
	// X:0 and Y:0
	posC := position{}
	// has type *position
	p := &position{10, 10}

	fmt.Println(posA, posB, posC, p)

	fmt.Printf("%T, %T, %v, %v\n", posA, p, p.X, p.Y)

	// (*p).X = 20 // works but no necessary
	p.X = 20
	fmt.Printf("%T, %v, %v\n", p, p.X, p.Y)

	// Struct instanceand assigning values.
	// var p position
	// c.X = 10
	// c.Y = 10

	// Comparing structs.
	x := position{1, 2}
	y := position{3, 4}
	z := position{3, 4}
	fmt.Println(x == y)
	fmt.Println(y == z)

	// Comparing different structs with same props.
	type point struct {
		X, Y int
		// Y, X int // can't compare with position X, Y
	}

	myPoint := point{1, 2}
	// fmt.Println(x == myPoint) // mismatched types
	fmt.Println(point(x) == myPoint)
}
