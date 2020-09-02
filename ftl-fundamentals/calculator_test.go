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

	t.Run("Variadic", func(t *testing.T) {
		var want float64 = 25
		nums := []float64{4, 6, 7, 8}
		if got := calc.Add(nums...); got != want {
			t.Errorf("got %f, want %f for trailing arguments %v", got, want, nums)
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
		nums := []float64{12, 6, 5, 5}
		if got := calc.Subtract(nums...); got != want {
			t.Errorf("got %f, want %f for trailing arguments %v", got, want, nums)
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
		nums := []float64{12, 6, 2, 4, 7, 8, 2, 5, 7, 5}
		if got := calc.Multiply(nums...); got != want {
			t.Errorf("got %f, want %f for trailing arguments %v", got, want, nums)
		}
	})
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

	t.Run("Variadic", func(t *testing.T) {
		var want float64 = 1
		nums := []float64{12, 6, 2}
		if got, _ := calc.Divide(nums...); got != want {
			t.Errorf("got %f, want %f for trailing arguments %v", got, want, nums)
		}
	})
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

func TestStr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input  string
		output float64
	}{
		{"2*2", 4},
		{"1 + 1.5", 2.5},
		{"18    /    6", 3},
		{"100-0.1", 99.9},
	}

	for _, test := range tests {
		if got := calc.Str(test.input); got != test.output {
			t.Errorf("got %f, want %f", got, test.output)
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
