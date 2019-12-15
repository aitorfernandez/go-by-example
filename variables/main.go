package main

import (
	"fmt"
	"math/cmplx"
)

var (
	// Variables with initializers
	b bool   = true
	s string = "I'm a String"
	// int  int8  int16  int32  int64
	i int = 1
	// uint uint8 uint16 uint32 uint64 uintptr
	maxInt uint64 = 1<<64 - 1
	// float32 float64
	f float32 = 4.12
	// complex64 complex128
	c complex128 = cmplx.Sqrt(-5 + 12i)

	// zero value
	x bool
	y float64
	z string
)

func title(t string) {
	fmt.Printf("\n-> %v\n", t)
}

func main() {
	title("Basic Types")
	fmt.Printf("Type: %T Value: %v\n", b, b)
	fmt.Printf("Type: %T Value: %v\n", s, s)
	fmt.Printf("Type: %T Value: %v\n", i, i)
	fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", f, f)
	fmt.Printf("Type: %T Value: %v\n", c, c)

	title("Zero Value")
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)

	title("Inside function and many at once")
	// var declaration with implicit type
	var j, k int = 1, 2
	// short assignment
	javascript, python := true, false
	fmt.Printf("%v, %v, %v, %v\n", j, k, javascript, python)

	title("Conversions")
	fmt.Printf("%T is %v\n", i, i)
	convI := float32(i)
	fmt.Printf("conversion to float64 %T is %v\n", convI, convI)
	convU := uint(convI)
	fmt.Printf("conversion to uint %T is %v\n", convU, convU)

	title("Inferred")
	v := 42.3
	v = 1
	fmt.Printf("v is of type %T, %v\n", v, v)
}
