package db

import (
	"cyaniccerulean.com/inventory/v2/internal/model"
)

// Function to get one entry from the database by ID
func (d Database) GetEntryByID(entryID string) (model.Entry, error) {
	var entry model.Entry

	// Query the Entry table to get the entry by ID
	err := d.db.QueryRow("SELECT ID, Name, Description, Note FROM Entry WHERE ID = ?", entryID).Scan(&entry.ID, &entry.Name, &entry.Description, &entry.Note)
	if err != nil {
		return model.Entry{}, err
	}

	// Query the Rectangle table to get all rectangles associated with the entry
	rows, err := d.db.Query("SELECT X, Y FROM Rectangle WHERE EntryID = ?", entryID)
	if err != nil {
		return model.Entry{}, err
	}
	defer rows.Close()

	// Iterate over the rows to build the slice of rectangles
	for rows.Next() {
		var rect model.Rectangle
		err := rows.Scan(&rect.X, &rect.Y)
		if err != nil {
			return model.Entry{}, err
		}
		entry.Rectangles = append(entry.Rectangles, rect)
	}

	return entry, nil
}

// Function to get all entries from the database
func (d Database) GetAllEntries() ([]model.Entry, error) {
	var entries []model.Entry

	// Query all entries from the Entry table
	rows, err := d.db.Query("SELECT ID, Name, Description, Note FROM Entry")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows to build the slice of entries
	for rows.Next() {
		var entry model.Entry
		err := rows.Scan(&entry.ID, &entry.Name, &entry.Description, &entry.Note)
		if err != nil {
			return nil, err
		}

		// Query the Rectangle table to get all rectangles associated with the entry
		rects, err := d.db.Query("SELECT X, Y FROM Rectangle WHERE EntryID = ?", entry.ID)
		if err != nil {
			return nil, err
		}
		defer rects.Close()

		// Iterate over the rows to build the slice of rectangles for each entry
		for rects.Next() {
			var rect model.Rectangle
			err := rects.Scan(&rect.X, &rect.Y)
			if err != nil {
				return nil, err
			}
			entry.Rectangles = append(entry.Rectangles, rect)
		}

		entries = append(entries, entry)
	}

	return entries, nil
}
