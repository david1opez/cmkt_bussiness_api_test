package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"campusmarket/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type UpdateBusinessPayload struct {
	BusinessId  string	`bson:"businessId,omitempty"`
	Field		string	`bson:"field,omitempty"`
	Value		any	`bson:"value,omitempty"`
}

func UpdateBusiness(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Get Request Data
		decoder := json.NewDecoder(r.Body)
		var reqBody UpdateBusinessPayload

		err := decoder.Decode(&reqBody)

		if err != nil {
			panic(err)
		}

		// Validate Data

		// Connect to Mongo
		user := os.Getenv("MONGO_USER")
		password := os.Getenv("MONGO_PASSWORD")
		host := os.Getenv("MONGO_HOST")

		client := mongo.Connect(user, password, host)
		defer mongo.Disconnect(client)

		// Update Business Info
		coll := mongo.GetCollection(client, "businesses")
		
		mongo.UpdateOne(coll, reqBody.BusinessId, bson.D{{Key: reqBody.Field, Value: reqBody.Value}})
	} else if r.Method == "GET" {
		fmt.Fprintf(w, "<h1>Update Business</h1>")
	}
}