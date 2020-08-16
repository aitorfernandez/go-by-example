package main

import (
	"fmt"
	"sort"
)

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
	// if _, ok := m["top"]; ok == false {
	// or
	if _, ok := m["top"]; !ok {
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

	// maps cannot be compared to each other, the only legal comparison is with nil.
	equal := func(x, y map[string]int) bool {
		if len(x) != len(y) {
			return false
		}
		for k, xv := range x {
			if yv, ok := y[k]; !ok || yv != xv {
				return false
			}
		}
		return true
	}
	// true
	fmt.Println("equal", equal(map[string]int{"a": 42}, map[string]int{"a": 42}))

	// map[string]interface{}
	foods := map[string]interface{}{
		"bacon": "delicious",
		"eggs": struct {
			source string
			price  float64
		}{"chicken", 1.75},
		"steak": true,
	}
	fmt.Println("foods", foods)

	// Ordering map keys
	fmt.Println("Ordering map keys")

	var dishes []string
	for k := range foods {
		dishes = append(dishes, k)
	}
	sort.Strings(dishes)

	for _, d := range dishes {
		fmt.Printf("%s, %v\n", d, foods[d])
	}

	// Ordering map values
	fmt.Println("Ordering map values")

	menu := map[string]int{
		"bacon": 4,
		"eggs":  4,
		"steak": 9,
	}

	// var swap map[int]string // panic: assignment to entry in nil map
	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range menu {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})

	for _, kv := range ss {
		fmt.Printf("%s, %d\n", kv.Key, kv.Value)
	}

	// Clean map
	fmt.Println("Clean map")
	menu = map[string]int{}

	fmt.Println(">map", menu)
}
