package gmachine_test

import (
	"testing"

	"github.com/aitorfernandez/go-by-example/gmachine"
)

func TestNew(t *testing.T) {
	t.Parallel()

	g := gmachine.New()

	wantMemSize := gmachine.DefaultMemSize
	gotMemSize := len(g.Memory)
	if wantMemSize != gotMemSize {
		t.Errorf("want %d words of memory, got %d", wantMemSize, gotMemSize)
	}

	var wantP uint64 = 0
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}

	var wantA uint64 = 0
	if wantA != g.P {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}

	var wantMemValue uint64 = 0
	gotMemValue := g.Memory[gmachine.DefaultMemSize-1]
	if wantMemValue != gotMemValue {
		t.Errorf("want last memory location to contain %d, got %d", wantMemValue, gotMemValue)
	}
}

func TestOpHALT(t *testing.T) {
	t.Parallel()

	g := gmachine.New()
	g.Run()

	var wantP uint64 = 1
	if wantP != g.P {
		t.Errorf("want P == %d, got %d", wantP, g.P)
	}
}

func TestOpNOOP(t *testing.T) {
	t.Parallel()

	g := gmachine.New()
	g.RunProgram([]uint64{
		gmachine.OpNOOP,
	})

	var wantP uint64 = 2
	if wantP != g.P {
		t.Errorf("want P == %d, got %d", wantP, g.P)
	}
}

func TestOpINCA(t *testing.T) {
	t.Parallel()

	g := gmachine.New()
	g.RunProgram([]uint64{
		gmachine.OpINCA,
	})

	var wantA uint64 = 1
	if wantA != g.A {
		t.Errorf("want A == %d, got %d", wantA, g.A)
	}
}

func TestOpDECA(t *testing.T) {
	t.Parallel()

	g := gmachine.New()
	g.Memory[0] = 2
	g.Run()

	var wantA uint64 = 1
	if wantA != g.A {
		t.Errorf("want A == %d, got %d", wantA, g.A)
	}
}

func TestSub2(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input  uint64
		output uint64
	}{
		{2, 0},
		{7, 5},
		{9, 7},
	}

	for _, test := range tests {
		g := gmachine.New()
		g.RunProgram([]uint64{
			gmachine.OpSETA, test.input,
			gmachine.OpDECA,
			gmachine.OpDECA,
		})

		if test.output != g.A {
			t.Errorf("want A == %d, got %d", test.output, g.A)
		}
	}
}

func TestOpSETA(t *testing.T) {
	t.Parallel()

	g := gmachine.New()
	g.RunProgram([]uint64{
		gmachine.OpSETA, 9,
	})

	var wantA uint64 = 9
	if wantA != g.A {
		t.Errorf("want A == %d, got %d", wantA, g.A)
	}

	var wantP uint64 = 3
	if wantP != g.P {
		t.Errorf("want P == %d, got %d", wantP, g.P)
	}
}
