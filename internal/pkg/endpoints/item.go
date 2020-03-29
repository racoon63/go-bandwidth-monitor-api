package endpoints

import (
	"net/http"
)

// GetItem retrieves data Item from storage backend and writes it to HTTP response handler.
func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
