package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

// initialize SQLite database
func Initialize(ctx context.Context, path string) (*Database, error) {
	// try to open database file
	// Create the file if it does not exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return nil, fmt.Errorf("error creating new file: %w", err)
		}
		file.Close()
	}

	// Open the database
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	// create table for entries
	_, err = db.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS Entry (ID TEXT PRIMARY KEY, Name TEXT, Description TEXT, Note TEXT);")
	if err != nil {
		return nil, err
	}

	// create table for rectangle coordinates
	_, err = db.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS Rectangle (EntryID TEXT, X INTEGER, Y INTEGER, FOREIGN KEY(EntryID) REFERENCES Entry(ID));")
	return &Database{db: db}, err
}

// close file
func (d Database) Close() error {
	return d.db.Close()
}
