package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"campusmarket/middleware"
	"campusmarket/mongo"
)

func GetPost(w http.ResponseWriter, r *http.Request) {
	if middleware.Authorize(r.Header.Get("X-CMKT-Authorization")) == http.StatusUnauthorized {
		http.Error(w, "Unauthorized Request", http.StatusUnauthorized)
		return
	}
	
	if r.Method == "GET" {
		// Get Request Query Data
		id := r.URL.Query().Get("id")

		// Connect to Mongo
		user := os.Getenv("MONGO_USER")
		password := os.Getenv("MONGO_PASSWORD")
		host := os.Getenv("MONGO_HOST")

		client := mongo.Connect(user, password, host)
		defer mongo.Disconnect(client)

		coll := mongo.GetCollection(client, "businesses")

		result, err := mongo.FindOne(coll, id)

		if err != nil {
			if err == mongo.DocumentNotFoundError {
				http.Error(w, "Document Not Found", http.StatusNotFound)
				return
			} else {
				http.Error(w, "Error on PaginatedFind", http.StatusInternalServerError)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}