package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"campusmarket/middleware"
	"campusmarket/models"
	"campusmarket/mongo"
)

func RegisterBusiness(w http.ResponseWriter, r *http.Request) {
	if middleware.Authorize(r.Header.Get("X-CMKT-Authorization")) == http.StatusUnauthorized {
		http.Error(w, "Unauthorized Request", http.StatusUnauthorized)
		return
	}

	if r.Method == "POST" {
		// Get Request Data
		decoder := json.NewDecoder(r.Body)
		var reqBody models.Business

		err := decoder.Decode(&reqBody)

		if err != nil {
			http.Error(w, "Failed To Decode Request Body", http.StatusBadRequest)
			fmt.Println(err)
			return
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
			http.Error(w, "Failed To Create Business Object", http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		mongo.InsertOne(coll, newB)
	}
}