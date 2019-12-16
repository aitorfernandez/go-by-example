package main

import "fmt"

// Closure helps to limit the scope of variables used by multiple functions
func wrapper() func() int {
	i := 10
	return func() int {
		i--
		return i
	}
}

func main() {
	i := 0
	add := func() int {
		i++
		return i
	}

	fmt.Println(add())
	fmt.Println(add())

	sub := wrapper()
	fmt.Println(sub())
	fmt.Println(sub())
}
