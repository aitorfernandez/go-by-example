package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Integer to String.
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

	fmt.Println(strconv.FormatInt(int64(x), 2))
	// or
	fmt.Printf("%b\n", x)

	// String to Integer.
	// z, _ := strconv.Atoi(y)
	z, _ := strconv.ParseInt("123", 10, 64)
	fmt.Println("base 10, up to 64 bits", z)
}
