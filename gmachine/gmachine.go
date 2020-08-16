// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024

const (
	// HALT stop the machine.
	HALT = iota
	// NOOP short for NO OPeration.
	NOOP
)

// Gmachine keeps information about our computer system model.
type Gmachine struct {
	// Program Counter.
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

// Run increments the P register.
func (m *Gmachine) Run() {
	m.P++
}
