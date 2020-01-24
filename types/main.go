package main

import "fmt"

type Celsius float64

const FreezingC Celsius = 0

func (c Celsius) String() string {
	return fmt.Sprintf("%gC\n", c)
}

func main() {
	fmt.Println(FreezingC.String())
}
