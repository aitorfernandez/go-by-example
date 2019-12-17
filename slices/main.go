package main

import "fmt"

func main() {
	// A slice literal is like an array literal without the length
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[3:len(s)]
	printSlice(s)

	// Slice the slice to give it zero length
	s = s[:0]
	printSlice(s)

	// append
	s = append(s, 1)
	printSlice(s)
	s = append(s, 3, 5, 7, 9)
	printSlice(s)
	s = append(s, []int{2, 4, 6, 8}...)
	printSlice(s)

	printValues([]int{99, 98, 97, 96}...)

	// with make
	a := make([]int, 0, 15) // len(0), cap(15)
	printSlice(a)

	// slice of slice
	cells := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
	}

	for i, l := 0, len(cells); i < l; i++ {
		fmt.Printf("%v", cells[i])
	}
}

func printSlice(s []int) {
	fmt.Printf("len %v, cap %v, %v\n", len(s), cap(s), s)
}

func printValues(is ...int) {
	for i := 0; i < len(is); i++ {
		fmt.Println(is[i])
	}
}
