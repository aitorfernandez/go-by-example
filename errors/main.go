package main

import (
	"errors"
	"fmt"
	"math"
)

// Curstom
var NegativeErr = errors.New("negative number")

// Or even more customs
type NegativeMathErr struct {
	code int
	err  error
}

func (n *NegativeMathErr) Error() string {
	fmt.Errorf("negative number with code %v and message %v", n.code, n.err)
}

func main() {
	fmt.Println("vim-go")
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// return 0, errors.New("negative number")
		// return 0, NegativeErr
		// return 0, fmt.Errorf("negative number: %v", f)
		m := fmt.Errorf("negative number: %v", f)
		return 0, &NegativeMathErr{"505", m}
	}
	return math.Sqrt(f), nil
}
