package main

import "log/slog"

func main() {
	ints := map[string]int64{
		"a": 3,
		"b": 7,
		"c": 2,
	}

	floats := map[string]float64{
		"a": 3.12,
		"b": 6.88,
		"c": 2.54,
	}

	intSum := SumInt64(ints)
	floatSum := SumFloat64(floats)

	slog.Info("Non generic methods", "int_totals", intSum, "float_totals", floatSum)

	intSum = Sum(ints)
	floatSum = Sum(floats)

	slog.Info("Generic methods", "int_totals", intSum, "float_totals", floatSum)

	intSum = Sum(ints)
	floatSum = Sum(floats)

	slog.Info("Generic methods with type constraint interface", "int_totals", intSum, "float_totals", floatSum)
}

func SumInt64(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}

	return s
}

func SumFloat64(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}

	return s
}

func Sum[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}

type Number64 interface {
	int64 | float64
}

func SumNumber64[K comparable, V Number64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}
