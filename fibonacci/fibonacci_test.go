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

func TestFibExplicitBinet(t *testing.T) {
	FibFunctionTest(t, FibExplicitBinet)
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
	testFunctions := map[string](func(int) int64){
		"FibRecursive":       FibRecursive,
		"FibRecursiveCached": FibRecursiveCached,
		"FibExplicitBinet":   FibExplicitBinet,
	}

	for testName, testFunc := range testFunctions {
		b.Run(testName, func(b *testing.B) {
			FibFunctionBenchmarker(b, testFunc)
		})
	}
}

func FibFunctionBenchmarker(b *testing.B, f func(int) int64) {
	for n := range 20 {
		testName := fmt.Sprintf("(n = %d)", n)
		b.Run(testName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f(n)
			}
		})
	}
}
