// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns the result of the multiplication.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the result of the division
// and a value of type error if undefined result.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("any dividend divided by zero is undefined %f, %f", a, b)
	}

	return (a / b), nil
}

// Sqrt returns the square root of a number.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("only nonnegative values are permited %f", a)
	}

	return math.Sqrt(a), nil
}
