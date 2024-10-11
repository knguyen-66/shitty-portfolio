package data

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDatabase() error {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		return err
	}
	DB = db
	return nil
}

func CloseDatabase() {
	DB.Close()
}
