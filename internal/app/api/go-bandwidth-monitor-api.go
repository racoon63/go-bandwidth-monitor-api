package api

import (
	"errors"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/racoon63/go-bandwidth-monitor-api/internal/pkg/config"
	"github.com/racoon63/go-bandwidth-monitor-api/internal/pkg/endpoints"
	//"github.com/racoon63/go-bandwidth-monitor-api/internal/pkg/storage"
)

type conf struct{}

func getIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Println(err)
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			log.Println(err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue
			}
			return ip.String(), nil
		}
	}
	return "127.0.0.1", errors.New("Can not fetch IP address. Fallback to localhost address")
}

// Init checks if all conditions met to start the API.
func Init() {
	log.Print("Initialize API...")

	conf := config.Read()

	log.Print("Verify database connectivity...")
	for {
		log.Println("Trying to connect to database...")
		_, err := net.Dial("tcp", conf.Address+":"+conf.Port)
		if err != nil {
			log.Println(err)
			log.Println("Could not reach database. Waiting 60 seconds and retry...")
			time.Sleep(60 * time.Second)
		} else {
			log.Println("Connection verification was successful!")
			break
		}
	}
}

// Run starts the API and handles the incoming requests.
func Run() {

	apiPath := "/api/v1"
	dataPath := apiPath + "/data"
	dataItemPath := dataPath + "/{ts}"
	dataItemValuePath := dataItemPath + "/{ping|download|upload}"

	r := mux.NewRouter()

	r.HandleFunc(apiPath, endpoints.GetInfo).Methods("GET")
	r.HandleFunc(dataPath, endpoints.GetAllData).Methods("GET")
	r.HandleFunc(dataItemPath, endpoints.GetItem).Methods("GET")
	r.HandleFunc(dataItemValuePath, endpoints.GetItemValue).Methods("GET")

	myIP, _ := getIP()

	log.Print("Starting API on ", myIP, ":", "8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
