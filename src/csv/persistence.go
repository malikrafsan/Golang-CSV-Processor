package csv

import (
	"context"

	"github.com/malikrafsan/Golang-CSV-Processor/src/db"
)

func StoreCSV(ctx context.Context, fileID string, data string, sums string) error {
	row := db.CSVTable{
		FileID: fileID,
		Data:   data,
		Sums:   sums,
	}

	conn, err := db.GetDBConn()
	if err != nil {
		return err
	}

	result := conn.Conn.Create(&row)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func FetchCSV(ctx context.Context, fileID string) ([]db.CSVTable, error) {
	conn, err := db.GetDBConn()
	if err != nil {
		return nil, err
	}

	var rows []db.CSVTable
	result := conn.Conn.Where("file_id = ?", fileID).Find(&rows)
	if result.Error != nil {
		return nil, result.Error
	}

	return rows, nil
}
