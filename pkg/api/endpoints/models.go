package endpoints

import "go.mongodb.org/mongo-driver/bson/primitive"

// APIInfo holds all api related information.
type APIInfo struct {
	Version     string `json:"version"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// Client holds all measurement data that is related to the client side.
type Client struct {
	Country   string `json:"country,omitempty" bson:"country"`
	IP        string `json:"ip,omitempty" bson:"ip"`
	Isp       string `json:"isp,omitempty" bson:"isp"`
	Isprating string `json:"isp-rating,omitempty" bson:"isp-rating"`
	Rating    string `json:"rating,omitempty" bson:"rating"`
}

// Server holds all measurement data that is related to the server side.
type Server struct {
	City    string  `json:"city,omitempty" bson:"city"`
	Country string  `json:"country,omitempty" bson:"country"`
	Host    string  `json:"host,omitempty" bson:"host"`
	ID      string  `json:"id,omitempty" bson:"id"`
	Latency float64 `json:"latency,omitempty" bson:"latency"`
	Sponsor string  `json:"sponsor,omitempty" bson:"sponsor"`
	URL     string  `json:"url,omitempty" bson:"url"`
	URL2    string  `json:"url2,omitempty" bson:"url2"`
}

// Item holds all data related to one data item, including Client and Server data.
type Item struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Download  float64            `json:"download,omitempty" bson:"download,omitempty"`
	Upload    float64            `json:"upload,omitempty" bson:"upload,omitempty"`
	Ping      float64            `json:"ping,omitempty" bson:"ping,omitempty"`
	Timestamp string             `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	Client    Client             `json:"client,omitempty" bson:"client,omitempty"`
	Server    Server             `json:"server,omitempty" bson:"server,omitempty"`
}
