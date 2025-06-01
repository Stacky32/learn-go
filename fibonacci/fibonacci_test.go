package fibonacci

import (
	"fmt"
	"testing"
)

func TestFibRecursive(t *testing.T) {
	FibFunctionTest(t, FibRecursive)
}

func TestFibRecursiveCached(t *testing.T) {
	FibFunctionTest(t, FibRecursiveCached)
}

func FibFunctionTest(t *testing.T, fibFunction func(int) int64) {
	var tests = []struct {
		n    int
		want int64
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
		{7, 21},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Fib(%d)", tt.n)
		t.Run(testName, func(t *testing.T) {
			actual := fibFunction(tt.n)
			if actual != tt.want {
				t.Errorf("Expected %d but received %d", tt.want, actual)
			}
		})
	}
}

func BenchmarkFibFunctions(b *testing.B) {
	b.Run("FibRecursive", BenchmarkFibFunction)
	b.Run("FibRecursiveCached", BenchmarkFibFunctionRecursive)
}

func BenchmarkFibFunction(b *testing.B) {
	FibFunctionBenchmarker(b, FibRecursive)
}

func BenchmarkFibFunctionRecursive(b *testing.B) {
	FibFunctionBenchmarker(b, FibRecursiveCached)
}

func FibFunctionBenchmarker(b *testing.B, f func(int) int64) {
	for n := 3; n < 15; n++ {
		testName := fmt.Sprintf("(n = %d) FibRecursive", n)
		b.Run(testName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FibRecursive(n)
			}
		})
	}
}
