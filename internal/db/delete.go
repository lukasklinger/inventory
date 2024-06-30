package db

// Function to delete an entry from the database by ID
func (d Database) DeleteEntry(entryID string) error {
	// Start a database transaction
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}

	// Prepare the statement to delete from Rectangle table
	stmtRect, err := tx.Prepare("DELETE FROM Rectangle WHERE EntryID = ?")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmtRect.Close()

	// Delete rectangles associated with the entry
	_, err = stmtRect.Exec(entryID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Prepare the statement to delete from Entry table
	stmtEntry, err := tx.Prepare("DELETE FROM Entry WHERE ID = ?")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmtEntry.Close()

	// Delete the entry from the Entry table
	_, err = stmtEntry.Exec(entryID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction if everything was successful
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
