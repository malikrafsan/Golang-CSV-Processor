package utils

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func SumAll[E Number](numbers ...E) E {
	var sum E
	for _, number := range numbers {
		sum += number
	}
	return sum
}
