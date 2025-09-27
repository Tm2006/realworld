package db

import (
	"database/sql"
	_ "embed"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	//go:embed schema.sql
	schema string
)

func InitDB(dbPath string) (*sql.DB, error) {
	// Open or create the database file
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Check if the file is new or empty
	info, err := os.Stat(dbPath)
	if os.IsNotExist(err) || (err == nil && info.Size() == 0) {
		_, err = db.Exec(schema)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
