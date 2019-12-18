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
	fmt.Printf("%T, %v, %v", p, p.X, p.Y)
}
