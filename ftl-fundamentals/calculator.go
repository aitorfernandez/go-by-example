// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Add returns the result of adding all numbers together.
func Add(nums ...float64) float64 {
	var result float64
	for _, n := range nums {
		result += n
	}

	return result
}

// Subtract takes `n` numbers and returns the result of subtracting
// them all together.
func Subtract(nums ...float64) float64 {
	var result float64 = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		result -= nums[i]
	}
	return result
}

// Multiply takes `n` numbers and returns the result of the
// multiplication them all together.
func Multiply(nums ...float64) float64 {
	var result float64 = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		result *= nums[i]
	}
	return result
}

// Divide takes 'n' numbers and returns the result of the division
// them all together.
func Divide(nums ...float64) (float64, error) {
	var result float64 = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if nums[i] == 0 {
			return 0, fmt.Errorf("any dividend divided by zero is undefined %f", nums[i])
		}

		result = result / nums[i]
	}

	return result, nil
}

// Sqrt returns the square root of a number.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("only nonnegative values are permited %f", a)
	}

	return math.Sqrt(a), nil
}

// Str accepts a string as an input calc operation and returns the result.
func Str(str string) float64 {
	str = strings.Replace(str, " ", "", -1)
	i := strings.IndexAny(str, "*+/-")

	parse := func(s string) float64 {
		i, _ := strconv.ParseFloat(s, 64)
		return i
	}

	nums := []float64{parse(str[:i]), parse(str[i+1:])}

	switch str[i] {
	case '+':
		return Add(nums...)
	case '-':
		return Subtract(nums...)
	case '/':
		res, _ := Divide(nums...)
		return res
	case '*':
		return Multiply(nums...)
	}

	return 0
}
