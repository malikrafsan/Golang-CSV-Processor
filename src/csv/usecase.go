package csv

import (
	"context"
	"encoding/csv"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/malikrafsan/Golang-CSV-Processor/src/utils"
)

func SumRowCSV(ctx context.Context, row *[]string) (int64, error) {
	intSlice, err := utils.StrSliceToIntSlice(row)
	if err != nil {
		return 0, err
	}
	return utils.SumAll(intSlice...), nil
}

func SumCSV(ctx context.Context, csvReader *[][]string) ([]int64, error) {
	leng := len(*csvReader)

	sums := make([]int64, leng)
	for i, line := range *csvReader {
		sumRow, err := SumRowCSV(ctx, &line)
		if err != nil {
			return nil, err
		}
		sums[i] = sumRow
	}

	return sums, nil
}

func ProcessFile(ctx context.Context, file *multipart.FileHeader, fileID string) error {
	// open file as csv
	csvFile, err := file.Open()
	if err != nil {
		return err
	}

	// read csv file
	csvReader, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return err
	}

	for _, row := range csvReader {
		sumRow, err := SumRowCSV(ctx, &row)
		if err != nil {
			return err
		}

		strRow := strings.Join(row, ",")
		err = StoreCSV(ctx, fileID, strRow, strconv.FormatInt(sumRow, 10))
		if err != nil {
			return err
		}
	}

	return nil
}

func GetProcessedFile(ctx context.Context, fileID string) (GetProcessedCSVData, error) {
	rows, err := FetchCSV(ctx, fileID)
	if err != nil {
		return GetProcessedCSVData{}, err
	}

	sums := make([]string, len(rows))
	for idx, row := range rows {
		sums[idx] = row.Sums
	}

	return GetProcessedCSVData{
		FileID: fileID,
		Sum:    strings.Join(sums, ","),
	}, nil

}
