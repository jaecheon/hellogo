package gogeneric

import "fmt"

type Number interface {
	int64 |
	float64
}

func Gogen() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first": 35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n", SumInts(ints), SumFloats(floats))
}

func SumInts(m map[string]int64) int64 {
	var s int64

	for _, v := range m {
		s += v
	}

	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64

	for _, v := range m {
		s += v
	}

	return s
}

// func SumIntsOrFloats(k comparable, v int64, float64 BadExpr, m map[K]V) V {
// 	var s V

// 	for _, v := range m {
// 		s += v
// 	}

// 	return s
// }

// func SumNumbers(K comparable, V Number, _ m, _ K, V BadExpr, _ BadExpr, V BadExpr, _, v BadExpr, m BadExpr, _ v,
// 	_ BadExpr, s BadExpr, _ BadExpr)