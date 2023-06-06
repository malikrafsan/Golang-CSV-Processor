package utils

import (
	"strconv"
)

func StrSliceToIntSlice(s *[]string) ([]int64, error) {
	var intSlice []int64
	for _, str := range *s {
		num, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, err
		}

		intSlice = append(intSlice, num)
	}
	return intSlice, nil
}

func IntSliceToStrSlice(s *[]int64) ([]string, error) {
	var strSlice []string
	for _, num := range *s {
		str := strconv.FormatInt(num, 10)
		strSlice = append(strSlice, str)
	}
	return strSlice, nil
}
