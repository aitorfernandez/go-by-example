// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Add returns the result of adding all numbers together.
func Add(a, b float64, nums ...float64) (float64, error) {
	result := a + b
	if result > math.MaxFloat64 {
		return 0, fmt.Errorf("sum is greater than %v", math.MaxFloat64)
	}
	for _, n := range nums {
		result += n
	}
	return result, nil
}

// Subtract takes `n` numbers and returns the result subtracting
// the second and subsequent numbers from the first.
func Subtract(a, b float64, nums ...float64) float64 {
	result := a - b
	for _, n := range nums {
		result -= n
	}
	return result
}

// Multiply takes `n` numbers and returns the result of the
// multiplication them all together.
func Multiply(a, b float64, nums ...float64) float64 {
	result := a * b
	for _, n := range nums {
		result *= n
	}
	return result
}

// Divide takes 'n' numbers and returns the result of dividing
// the first by the second and subsequent numbers, or an error
// if any of the divisors is zero.
func Divide(a, b float64, nums ...float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("bad input: division by zero is undefined %f, %f", a, b)
	}
	result := a / b
	for _, n := range nums {
		if n == 0 {
			return 0, fmt.Errorf("bad input: division by zero is undefined %f, %f", result, n)
		}
		result /= n
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

// Evaluate accepts a string expresion  as an input operation and returns the result.
func Evaluate(str string) (float64, error) {
	str = strings.ReplaceAll(str, " ", "")
	pos := strings.IndexAny(str, "*+/-")
	if pos < 0 {
		return 0, fmt.Errorf("operator invalid in %s", str)
	}

	values := strings.Split(str, string(str[pos]))
	var nn []float64
	for _, val := range values {
		n, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return 0, nil
		}
		nn = append(nn, n)
	}

	switch str[pos] {
	case '+':
		return Add(nn[0], nn[1], nn[2:]...)
	case '-':
		return Subtract(nn[0], nn[1], nn[2:]...), nil
	case '/':
		return Divide(nn[0], nn[1], nn[2:]...)
	case '*':
		return Multiply(nn[0], nn[1], nn[2:]...), nil
	default:
		return 0, fmt.Errorf("expresion is invalid %s", string(str[pos]))
	}
}
