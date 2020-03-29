package endpoints

// Client holds all measurement data that is related to the client side.
type Client struct {
	Country   string  `json:"country"`
	IP        string  `json:"ip"`
	Isp       string  `json:"isp"`
	Isprating float32 `json:"isp-rating"`
	Rating    float32 `json:"rating"`
}

// Server holds all measurement data that is related to the server side.
type Server struct {
	City    string  `json:"city"`
	Country string  `json:"country"`
	Host    string  `json:"host"`
	ID      int     `json:"id"`
	Latency float32 `json:"latency"`
	Sponsor string  `json:"sponsor"`
	URL     string  `json:"url"`
	URL2    string  `json:"url2"`
}

// Item holds all data related to one data item, including Client and Server data.
type Item struct {
	Download  int     `json:"download"`
	Upload    int     `json:"upload"`
	Ping      int     `json:"ping"`
	Timestamp string  `json:"timestamp"`
	Client    *Client `json:"client"`
	Server    *Server `json:"server"`
}
