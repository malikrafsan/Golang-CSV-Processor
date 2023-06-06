package csv

import "github.com/malikrafsan/Golang-CSV-Processor/src/utils"

func SumCSV(csvReader *[][]string) ([]int64, error) {
	leng := len(*csvReader)

	sums := make([]int64, leng)
	for i, line := range *csvReader {
		intSlice, err := utils.StrSliceToIntSlice(&line)
		if err != nil {
			return nil, err
		}

		sums[i] = utils.SumAll(intSlice...)
	}

	return sums, nil
}
