package routes

import (
	"encoding/json"
	"net/http"

	"cyaniccerulean.com/inventory/v2/internal/db"
	"cyaniccerulean.com/inventory/v2/internal/model"
	"github.com/google/uuid"
)

type API struct {
	db *db.Database
}

// initialize API routes by injecting a database connection
func InitAPI(db *db.Database) API {
	return API{db}
}

// handle incoming requests based on method
func (a API) APIHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.GetEntriesHandler(w, r)
	case http.MethodPost:
		a.CreateEntryHandler(w, r)
	case http.MethodPut:
		a.ModifyEntryHandler(w, r)
	case http.MethodDelete:
		a.DeleteEntryHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (a API) CreateEntryHandler(w http.ResponseWriter, r *http.Request) {
	var entry model.Entry

	// Decode the JSON body of the request into the entry struct
	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		http.Error(w, "Failed to decode JSON body", http.StatusBadRequest)
		return
	}

	// override ID to UUID
	id, err := uuid.NewV7()
	if err != nil {
		http.Error(w, "Failed to generate ID for entry", http.StatusInternalServerError)
		return
	}

	entry.ID = id.String()

	// Call the createEntry function to add the entry to the database
	err = a.db.CreateEntry(entry)
	if err != nil {
		http.Error(w, "Failed to create entry in the database", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Entry created successfully", "id": entry.ID})
}

func (a API) DeleteEntryHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	// Execute the DELETE query
	err := a.db.DeleteEntry(id)
	if err != nil {
		http.Error(w, "Failed to delete entry", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Entry deleted successfully"})
}

func (a API) GetEntriesHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := a.db.GetAllEntries()
	if err != nil {
		http.Error(w, "Failed to fetch entries", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entries)
}

func (a API) ModifyEntryHandler(w http.ResponseWriter, r *http.Request) {
	var entry model.Entry

	// Decode the JSON body of the request into the entry struct
	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		http.Error(w, "Failed to decode JSON body", http.StatusBadRequest)
		return
	}

	// delete the old entry
	err = a.db.DeleteEntry(entry.ID)
	if err != nil {
		http.Error(w, "Failed to delete entry in the database", http.StatusInternalServerError)
		return
	}

	// Call the createEntry function to add the entry to the database
	err = a.db.CreateEntry(entry)
	if err != nil {
		http.Error(w, "Failed to create entry in the database", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Entry created successfully"})
}
