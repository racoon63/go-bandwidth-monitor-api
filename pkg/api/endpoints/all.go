package endpoints

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	"time"
)

type d struct {
	items []bson.M
}

func openConnection() mongo.Client {

	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://root:example@localhost:27017/?connect=direct"))

	// Set context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connect to MongoDB
	err := client.Connect(ctx)

	if err != nil {
		log.Print(err)
	}

	// Check the connection
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Print(err)
	}

	log.Print("Connected to MongoDB")
	return *client
}

func closeConnection(client *mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Print(err)
	}
	log.Println("Connection to MongoDB closed")
}

// GetAllData tries to retrieve all data from storage backend and writes it to HTTP handler.
func GetAllData(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	client := openConnection()

	var data []Item
	collection := client.Database("bwm").Collection("data")
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var item Item
		cursor.Decode(&item)
		data = append(data, item)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	closeConnection(&client)
	json.NewEncoder(response).Encode(data)
}
