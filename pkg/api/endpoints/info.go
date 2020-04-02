package endpoints

import (
	"encoding/json"
	"net/http"
)

func newInfo(version string, description string, url string) *APIInfo {
	var info APIInfo

	info.Version = version
	info.Description = description
	info.URL = url

	return &info
}

// GetInfo creates new APIInfo object and writes it as json to HTTP response handler.
func GetInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	apiInfo := newInfo("v0.0.1", "This API is a frontend for the Bandwidth-Monitor service. Follow url to get more information.", "https://github.com/racoon63/bandwidth-monitor")
	json.NewEncoder(w).Encode(apiInfo)
}
