package calculator_test

import (
	"math/rand"
	"testing"

	calc "github.com/aitorfernandez/go-by-example/ftl-fundamentals"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	t.Run("Addition", func(t *testing.T) {
		tests := map[string]struct {
			a, b, output float64
		}{
			"1 plus 1 equals 2": {
				a:      1,
				b:      1,
				output: 2,
			},
			"Adding 0 does't change": {
				a:      1,
				b:      0,
				output: 1,
			},
		}

		for k, v := range tests {
			got := calc.Add(v.a, v.b)
			if got != v.output {
				t.Errorf("%s, got %f, want %f", k, got, v.output)
			}
		}
	})

	t.Run("Addition Properties", func(t *testing.T) {
		tests := map[string]struct {
			a, b float64
		}{
			"Adition is commutative": {
				a: calc.Add(3, 6),
				b: calc.Add(6, 3),
			},
			"Associativity": {
				a: calc.Add(2, 4) + 5,
				b: 2 + calc.Add(4, 5),
			},
		}

		for k, v := range tests {
			if v.a != v.b {
				t.Errorf("%s, want same result %f, %f", k, v.a, v.b)
			}
		}
	})
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		a, b, output float64
	}{
		"2 minus 1 equals 1": {
			a:      2,
			b:      1,
			output: 1,
		},
		"4 minus 6 equals negative 2": {
			a:      4,
			b:      6,
			output: -2,
		},
	}

	for k, v := range tests {
		got := calc.Subtract(v.a, v.b)
		if got != v.output {
			t.Errorf("%s, got %f, want %f", k, got, v.output)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		a, b, output float64
	}{
		"2 times three equals 6": {
			a:      2,
			b:      3,
			output: 6,
		},
		"1st mult arithmetic axiom 0": {
			a:      5,
			b:      0,
			output: 0,
		},
		"2nd mult arithmetic axiom -1": {
			a:      -1,
			b:      -1,
			output: 1,
		},
	}

	for k, v := range tests {
		got := calc.Multiply(v.a, v.b)
		if got != v.output {
			t.Errorf("%s, got %f, want %f", k, got, v.output)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		a, b, output float64
		errExpected  bool
	}{
		"20 divided by 4 equals 5": {
			a:           20,
			b:           4,
			output:      5,
			errExpected: false,
		},
		"Division by zero": {
			a:           6,
			b:           0,
			output:      0,
			errExpected: true,
		},
	}

	for k, v := range tests {
		got, err := calc.Divide(v.a, v.b)
		if err != nil && !v.errExpected {
			t.Fatalf("%s, error status %w, %v", k, err, v.errExpected)
		}
		if got != v.output {
			t.Errorf("%s, got %f, want %f", k, got, v.output)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		a, output   float64
		errExpected bool
	}{
		"The square root": {
			a:           4,
			output:      2,
			errExpected: false,
		},
		"With negative numbers": {
			a:           -1,
			output:      0,
			errExpected: true,
		},
	}

	for k, v := range tests {
		got, err := calc.Sqrt(v.a)
		if err != nil && !v.errExpected {
			t.Fatalf("%s, error status %w, %v", k, err, v.errExpected)
		}
		if got != v.output {
			t.Errorf("%s, got %f, want %f", k, got, v.output)
		}
	}
}

func BenchmarkAddRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := rand.Float64()
		y := rand.Float64()

		result := x + y
		if calc.Add(x, y) != result {
			b.Errorf("unexpected result %f, %f, %f", x, y, result)
		}
	}
}
