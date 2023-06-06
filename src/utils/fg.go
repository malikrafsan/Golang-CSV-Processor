package utils

type mapFunc[TIn any, TOut any] func(TIn) TOut

func Map[TIn any, TOut any](s []TIn, f mapFunc[TIn, TOut]) []TOut {
	result := make([]TOut, len(s))
	for i := range s {
		result[i] = f(s[i])
	}
	return result
}
