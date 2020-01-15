package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Lorem ipsum dolor sit amet, consectetur adipiscing elit"

	fmt.Println(len(str))

	// Iterate
	for _, c := range str {
		fmt.Println(string(c))
	}

	// Reverse
	reverse := func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}

	fmt.Println(reverse(str))

	// Slicing
	fmt.Println(str[:11])         // Lorem ipsum
	fmt.Println(str[len(str)-4:]) // elit

	// Count the number of times appear a character
	fmt.Println(strings.Count(str, "o"))

	// Replace
	fmt.Println(strings.ReplaceAll(str, "o", "e"))

	// Replace in range
	fmt.Println(strings.ReplaceAll(str[:10], "o", "e") + str[10:])

	// Uppercase
	fmt.Println(strings.ToUpper(str))

	// Lowercase
	fmt.Println(strings.ToLower(str))

	// Split
	fmt.Println(strings.Split(str, "sit"))

	// Contains
	fmt.Println(strings.Contains(str, "Lorem"))
}
