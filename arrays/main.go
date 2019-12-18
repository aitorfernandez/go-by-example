package main

import "fmt"

func main() {
	// a := [7]int{3, 5, 7, 9, 11, 13, 17}
	a := [...]int{3, 5, 7, 9, 11, 13, 17}
	fmt.Println(a)

	// Specific elements
	langs := [4]string{0: "Go", 3: "Javascript"}
	fmt.Println(langs)

	// old school
	// for i := 0; i < len(a); i++ {
	for i, l := 0, len(a); i < l; i++ {
		fmt.Println(a[i])
	}

	for _, v := range a {
		fmt.Println(v)
	}

	// multidimensional array
	arr := [2][2]int{
		{3, 5},
		{7, 9},
	}

	for _, r := range arr {
		for _, c := range r {
			fmt.Print(c, " ")
		}
		fmt.Println()
	}
}
