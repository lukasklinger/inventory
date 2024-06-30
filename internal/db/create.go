package db

import (
	"cyaniccerulean.com/inventory/v2/internal/model"
)

// Function to create a new entry in the database
func (d Database) CreateEntry(entry model.Entry) error {
	// Start a database transaction
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}

	// Prepare the statement for inserting into Entry table
	stmtEntry, err := tx.Prepare("INSERT INTO Entry(ID, Name, Description, Note) VALUES(?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmtEntry.Close()

	// Insert the entry into the Entry table
	_, err = stmtEntry.Exec(entry.ID, entry.Name, entry.Description, entry.Note)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Prepare the statement for inserting into Rectangle table
	stmtRect, err := tx.Prepare("INSERT INTO Rectangle(EntryID, X, Y) VALUES(?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmtRect.Close()

	// Insert rectangles associated with the entry into the Rectangle table
	for _, rect := range entry.Rectangles {
		_, err = stmtRect.Exec(entry.ID, rect.X, rect.Y)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction if everything was successful
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
