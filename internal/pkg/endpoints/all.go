package endpoints

import (
	// "encoding/json"
	"net/http"
	// "github.com/gorilla/mux"
)

// GetAllData tries to retrieve all data from storage backend and writes it to HTTP handler.
func GetAllData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// params := mux.Vars(r)
	// if val, ok := params["startts"]; ok {
	// 	// startts := val
	// }
	// if val, ok := params["stopts"]; ok {
	// 	// stopts := val
	// }

	// json.NewEncoder(w).Encode(item)
}
