package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/racoon63/go-bandwidth-monitor-api/internal/pkg/config"
	"github.com/racoon63/go-bandwidth-monitor-api/internal/pkg/endpoints"
	//"github.com/racoon63/go-bandwidth-monitor-api/internal/pkg/storage"
)

// Run starts the API and handles the incoming requests.
func Run() {

	var sconf config.ServiceConf
	sconf.GetServiceConf()

	apiPath := "/api/v1"
	dataPath := apiPath + "/data"
	dataItemPath := dataPath + "/{ts}"
	dataItemValuePath := dataItemPath + "/{ping|download|upload}"

	r := mux.NewRouter()

	r.HandleFunc(apiPath, endpoints.GetInfo).Methods("GET")
	r.HandleFunc(dataPath, endpoints.GetAllData).Methods("GET")
	r.HandleFunc(dataItemPath, endpoints.GetItem).Methods("GET")
	r.HandleFunc(dataItemValuePath, endpoints.GetItemValue).Methods("GET")

	log.Print("Starting API on ", sconf.Address, ":", sconf.Port)
	log.Fatal(http.ListenAndServe(":"+sconf.Port, r))
}
