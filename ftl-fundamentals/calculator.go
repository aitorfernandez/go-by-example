// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func hasEnoughNums(nums ...float64) (float64, bool) {
	switch len(nums) {
	case 0:
		return 0, false
	case 1:
		return nums[0], false
	default:
		return 0, true
	}
}

// Add returns the result of adding all numbers together.
func Add(nums ...float64) float64 {
	if val, ok := hasEnoughNums(nums...); !ok {
		return val
	}

	var result float64
	for _, n := range nums {
		result += n
	}
	return result
}

// Subtract takes `n` numbers and returns the result subtracting
// the second and subsequent numbers from the first.
func Subtract(nums ...float64) float64 {
	if val, ok := hasEnoughNums(nums...); !ok {
		return val
	}

	var result float64 = nums[0]
	for i := 1; i < len(nums); i++ {
		result -= nums[i]
	}
	return result
}

// Multiply takes `n` numbers and returns the result of the
// multiplication them all together.
func Multiply(nums ...float64) float64 {
	if val, ok := hasEnoughNums(nums...); !ok {
		return val
	}

	var result float64 = nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		result *= nums[i]
	}
	return result
}

// Divide takes 'n' numbers and returns the result of dividing
// the first by the second and subsequent numbers, or an error
// if any of the divisors is zero.
func Divide(nums ...float64) (float64, error) {
	if val, ok := hasEnoughNums(nums...); !ok {
		return val, nil
	}

	var result float64 = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			return 0, fmt.Errorf("bad input: division by zero is undefined %v", nums)
		}

		result /= nums[i]
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
	str = strings.ReplaceAll(str, " ", "")
	pos := strings.IndexAny(str, "*+/-")

	parse := func(s string) float64 {
		i, _ := strconv.ParseFloat(s, 64)
		return i
	}

	nums := []float64{parse(str[:pos]), parse(str[pos+1:])}

	switch str[pos] {
	case '+':
		return Add(nums...)
	case '-':
		return Subtract(nums...)
	case '/':
		res, err := Divide(nums...)
		if err != nil {
			return 0
		}
		return res
	case '*':
		return Multiply(nums...)
	default:
		return 0
	}
}
