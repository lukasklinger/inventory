package routes

import (
	"net/http"
)

// return 200 OK as a health check
func HandleHealthEndpoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
