package main

import "fmt"

func main() {
	// Old school for loop
	// for init; condition; post {
	// }

	i := 0
	// C's while
	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}

	// with Labels
out:
	for i := 0; i < 300; i++ {
		// []byte deals with bytes, that is, only 1 byte (8 bits) at a time
		fmt.Printf("%v, %v, %v, \n", i, string(i), []byte(string(i)))
		if i == 250 {
			break out
		}
	}
}
