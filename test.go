package main

import (
	"fmt"
	"log"
	"net"
	"time"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"github.com/racoon63/go-bandwidth-monitor-api/internal/pkg/config"
)

func main() {

	conf := config.readEnvVals()

	for {
		_, err := net.Dial("tcp", "172.19.0.2:27017")
		if err != nil {
			log.Println("Could not reach database. Wait until database becomes available.")
			time.Sleep(3)
		} else {
			log.Println("Connection test successful")
			break
		}
	}
	fmt.Println("Noice it worked!")
}
