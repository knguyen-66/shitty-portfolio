package data

import (
	"database/sql"
	"os"
	"path/filepath"

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

func InitDatabaseDefault() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	return InitDatabase("sqlite3", filepath.Join(dir, "data/app.db"))
}

func CloseDatabase() {
	if DB != nil {
		DB.Close()
	}
}
