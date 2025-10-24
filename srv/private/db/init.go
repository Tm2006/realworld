package db

import (
	"database/sql"
	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

var (
	//go:embed schema.sql
	schema string
)

func InitDB(dbPath string) (*sql.DB, error) {
	// Open database connection
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Check if database needs initialization
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name='users'").Scan(&count)
	if err != nil || count == 0 {
		// Initialize database schema
		_, err = db.Exec(schema)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
