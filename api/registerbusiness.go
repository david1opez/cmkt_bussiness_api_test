package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"campusmarket/models"
	"campusmarket/mongo"
)

func RegisterBusiness(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Get Request Data
		decoder := json.NewDecoder(r.Body)
		var reqBody models.Business

		err := decoder.Decode(&reqBody)

		if err != nil {
			panic(err)
		}

		// Validate Data
		newBusiness, err := models.NewBusiness(reqBody)
		
		// Connect to Mongo
		user := os.Getenv("MONGO_USER")
		password := os.Getenv("MONGO_PASSWORD")
		host := os.Getenv("MONGO_HOST")

		client := mongo.Connect(user, password, host)
		defer mongo.Disconnect(client)

		//Insert Data
		coll := mongo.GetCollection(client, "businesses")

		newB, err := models.NewBusiness(*newBusiness)

		if err != nil {
			panic(err.Error())
		}

		mongo.InsertOne(coll, newB)
	} else if r.Method == "GET" {
		fmt.Fprintf(w, "<h1>Register Business</h1>")
	}
}