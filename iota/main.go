package main

import "fmt"

// iota simplify definitions of incrementing numbers
const (
	x = iota
	y
	z
)

type timezone int

const (
	// EST iota: 0, EST: -5
	EST timezone = -(5 + iota)
	// CST iota: 1, CST: -6
	CST
	// MST iota: 2, MST: -7
	MST
	// PST iota: 3, MST: -8
	PST
)

const (
	// Reset iota, monday = 0
	monday = iota
	tuesday
	// ...
	// ...
)

const (
	on, off = iota, iota
)

func main() {
	fmt.Println("- Iota")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Printf("EST %v\n", EST)
	fmt.Printf("CST %v\n", CST)
	fmt.Printf("MST %v\n", MST)
	fmt.Printf("PST %v\n", PST)

	fmt.Printf("monday %v\n", monday)
	fmt.Printf("tuesday %v\n", tuesday)

	fmt.Printf("on off %v %v\n", on, off)
}
