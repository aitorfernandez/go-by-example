package main

import "fmt"

func reset(z *int) {
	fmt.Printf("z address in reset %p\n", z)
	*z = 0
}

func main() {
	i := 43

	fmt.Println("i - ", i)
	fmt.Println("i's memory address - ", &i)

	// point to i to the memory address where an int is stored
	p := &i
	// set i through the pointer
	*p = 21
	fmt.Printf("read i through the pointer %v\n", *p)
	fmt.Println(&i, p) // same address

	y := 5
	fmt.Printf("y address in main %v\n", &y)
	reset(&y)
	fmt.Println(y)
}
