// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024

const (
	// OpHALT stop the machine.
	OpHALT = iota
	// OpNOOP short for NO OPeration.
	OpNOOP
	// OpINCA increment.
	OpINCA
	// OpDECA decrement.
	OpDECA
)

// Gmachine keeps information about our computer system model.
type Gmachine struct {
	P      uint64 // Program Counter.
	A      uint64 // Accumulator register.
	Memory []uint64
}

// New returns a new G-machine.
func New() *Gmachine {
	mem := make([]uint64, DefaultMemSize)
	return &Gmachine{
		P:      0,
		A:      0,
		Memory: mem,
	}
}

// Run increments the P register.
func (m *Gmachine) Run() {
	for {
		op := m.Memory[m.P]
		m.P++
		switch op {
		case OpHALT:
			return
		case OpINCA:
			m.A++
		case OpDECA:
			m.A--
		}
	}
}
