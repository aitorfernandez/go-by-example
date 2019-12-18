package main

import "fmt"

type pos struct {
	X, Y int
}

type location map[string]*pos

func main() {
	// A nil map has no keys, nor can keys be added
	// var m map[string]int
	// fmt.Println(m)
	// fmt.Println(m == nil)
	// m["answers"] = 42

	// Mutation maps
	// var m = make(map[string]pos)
	// m := make(map[string]pos)
	// Composite literal
	// m := map[string]pos{}
	// Or
	// m := map[string]pos{
	// 	"top": pos{1, 0},
	// }
	m := make(location)
	fmt.Println(m)

	m["center"] = &pos{0, 0}
	fmt.Println(m["center"])

	fmt.Println(len(m))

	for key, value := range m {
		fmt.Println("Key: ", key, "Value: ", value)
	}

	delete(m, "top")
	v, ok := m["top"]
	fmt.Println("The value:", v, "Exist?", ok)

	// To test for a key without retrieving the value
	if _, ok := m["top"]; ok == false {
		fmt.Println("Value doesn't exist")
	}

	// no error
	delete(m, "bottom")
	m["bottom"] = &pos{-1, 0}
	if _, ok := m["bottom"]; ok {
		fmt.Println("Created")
	}

	m["bottom"].Y = 2
	fmt.Println(m["bottom"])
}
