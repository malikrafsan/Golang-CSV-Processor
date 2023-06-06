package utils

func SumAll(numbers ...int64) int64 {
	var sum int64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
