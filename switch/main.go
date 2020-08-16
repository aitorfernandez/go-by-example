package main

import "fmt"

func main() {
	foods := map[string]interface{}{
		"bacon": "delicious",
		"eggs": struct {
			source string
			price  float64
		}{"chicken", 1.75},
		"steak": true,
	}

	// type switch
	for k, v := range foods {
		// switch c := v.(type) {
		switch v.(type) {
		case string:
			fmt.Println(">", k, v)
		case float64:
			fmt.Println(">", k, v)
		default:
			// https://golang.org/pkg/fmt/
			fmt.Printf("Not sure what type item %q is, but I think it might be %T\n", k, v)
		}
	}
}
