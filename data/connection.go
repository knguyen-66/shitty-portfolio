package data

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDatabase(driverName string, connectionString string) error {
	db, err := sql.Open(driverName, connectionString)
	if err != nil {
		return err
	}
	DB = db
	return nil
}

func CloseDatabase() {
	DB.Close()
}
