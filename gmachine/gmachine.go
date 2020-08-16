// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024

// Gmachine keeps information about our computer system model.
type Gmachine struct {
	P      uint64
	Memory []uint64
}

// New returns a new G-machine.
func New() *Gmachine {
	mem := make([]uint64, DefaultMemSize)
	return &Gmachine{
		P:      0,
		Memory: mem,
	}
}
