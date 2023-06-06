package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CSVTable struct {
	gorm.Model
	FileID string
	Data   string
	Sums   string
}

type DBConn struct {
	Conn *gorm.DB
}

var dbConn *DBConn

func newDBConn() (*DBConn, error) {
	dsn := "root@tcp(127.0.0.1:3306)/csv_processor?charset=utf8&parseTime=True"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&CSVTable{})

	return &DBConn{
		Conn: db,
	}, nil
}

func GetDBConn() (*DBConn, error) {
	if dbConn != nil {
		return dbConn, nil
	}

	dbConn, err := newDBConn()
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}
