package endpoints

import (
	"net/http"
)

// GetItemValue retrieves a single value from a data item from storage backend and writes it to HTTP response handler
func GetItemValue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
