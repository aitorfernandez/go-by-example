package calculator_test

import (
	"math"
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
			got, err := calc.Add(v.a, v.b)
			if err != nil {
				t.Fatalf("error status %w", err)
			}
			if got != v.output {
				t.Errorf("%s, got %f, want %f", k, got, v.output)
			}
		}
	})

	t.Run("Variadic", func(t *testing.T) {
		var want float64 = 25
		got, err := calc.Add(4, 6, 7, 8)
		if err != nil {
			t.Fatalf("error status %w", err)
		}
		if got != want {
			t.Errorf("got %f, want %f for 4, 6, 7, 8", got, want)
		}
	})

	t.Run("overflow", func(t *testing.T) {
		got, err := calc.Add(math.MaxFloat64, math.MaxFloat64)
		if err == nil {
			t.Error("error is nil, want an error for overflow results")
		}
		if got != 0 {
			t.Errorf("got %f, want 0 for overflow", got)
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

	t.Run("Variadic", func(t *testing.T) {
		var want float64 = -4
		if got := calc.Subtract(12, 6, 5, 5); got != want {
			t.Errorf("got %f, want %f for 12, 6, 5, 5", got, want)
		}
	})
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

	t.Run("Variadic with 10 inputs", func(t *testing.T) {
		var want float64 = 11289600
		if got := calc.Multiply(12, 6, 2, 4, 7, 8, 2, 5, 7, 5); got != want {
			t.Errorf("got %f, want %f for 12, 6, 2, 4, 7, 8, 5, 7, 5", got, want)
		}
	})
}

func TestDivide(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		a, b, want  float64
		errExpected bool
	}{
		"20 divided by 4 equals 5": {
			a:           20,
			b:           4,
			want:        5,
			errExpected: false,
		},
		"Division by zero": {
			a:           6,
			b:           0,
			want:        0,
			errExpected: true,
		},
		// "Expected error and don't expect one": {
		// 	a:           2,
		// 	b:           0,
		// 	errExpected: false,
		// },
		"don't compare data value if error": {
			a:           3,
			b:           0,
			want:        999,
			errExpected: true,
		},
	}

	for k, v := range tests {
		got, err := calc.Divide(v.a, v.b)
		// if err != nil {
		// 	errReceived = err != nil
		// }
		errReceived := (err != nil)

		// Ignore the data value if error.
		if errReceived != v.errExpected {
			t.Fatalf("%s, %f divide %f, unexpected error status %v", k, v.a, v.b, err)
		}
		if !errReceived && got != v.want {
			t.Errorf("%s, %f divide %f want %f, got %f", k, v.a, v.b, v.want, got)
		}
	}

	t.Run("Variadic", func(t *testing.T) {
		var want float64 = 1
		got, err := calc.Divide(12, 6, 2)
		if err != nil {
			t.Fatalf("error status %w", err)
		}
		if got != want {
			t.Errorf("got %f, want %f for 12, 6, 2", got, want)
		}
	})
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		a, output   float64
		errExpected bool
	}{
		"Square root of 4 is 2": {
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

func TestEvaluate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input  string
		output float64
	}{
		{"2*2", 4},
		{"1 + 1.5", 2.5},
		{"18    /    6", 3},
		{"100-0.1", 99.9},
		{"2+4+5", 11},
	}

	for _, test := range tests {
		got, err := calc.Evaluate(test.input)
		if err != nil {
			t.Fatalf("error status %w", err)
		}
		if got != test.output {
			t.Errorf("got %f, want %f", got, test.output)
		}
	}

	t.Run("Invalid expresion", func(t *testing.T) {
		got, err := calc.Evaluate("2#2")
		if err == nil {
			t.Error("error is nil, want an error for invalid expresions")
		}
		if got != 0 {
			t.Errorf("got %v, want 0 for invalid expresions", got)
		}
	})
}

func BenchmarkAddRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := rand.Float64()
		y := rand.Float64()

		result := x + y
		got, err := calc.Add(x, y)
		if err != nil {
			b.Fatalf("error status %w", err)
		}
		if got != result {
			b.Errorf("unexpected result %f, %f, %f", x, y, result)
		}
	}
}
