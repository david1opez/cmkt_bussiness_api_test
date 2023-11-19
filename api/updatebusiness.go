package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"campusmarket/models"
	"campusmarket/mongo"
	"campusmarket/middleware"

	"go.mongodb.org/mongo-driver/bson"
)

type UpdateBusinessPayload struct {
	Id  string	`bson:"businessId,omitempty"`
	Field		string	`bson:"field,omitempty"`
	Value		any		`bson:"value,omitempty"`
}

func UpdateBusiness(w http.ResponseWriter, r *http.Request) {
	if middleware.Authorize(r.Header.Get("X-CMKT-Authorization")) == http.StatusUnauthorized {
		http.Error(w, "Unauthorized Request", http.StatusUnauthorized)
		return
	}
	
	if r.Method == "POST" {
		// Get Request Data
		decoder := json.NewDecoder(r.Body)
		var reqBody UpdateBusinessPayload

		err := decoder.Decode(&reqBody)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			panic(err)
		}

		// Validate Data
		validData, err := models.ValidateData(reqBody.Field, reqBody.Value)

		if err != nil {
			errorMessage := "Error Validating Data: '" + err.Error() + "'"
			http.Error(w, errorMessage, http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		// How can i print "Error Validating Data: " and the error message in golang?
		
		if !validData {
			errorMessage := "Invalid Data: '" + err.Error() + "'"
			http.Error(w, errorMessage, http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		// Connect to Mongo
		user := os.Getenv("MONGO_USER")
		password := os.Getenv("MONGO_PASSWORD")
		host := os.Getenv("MONGO_HOST")

		client := mongo.Connect(user, password, host)
		defer mongo.Disconnect(client)

		// Update Business Info
		coll := mongo.GetCollection(client, "businesses")
		
		mongo.UpdateById(coll, reqBody.Id, bson.D{{Key: reqBody.Field, Value: reqBody.Value}})
	}
}