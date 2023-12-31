package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"os"

	"campusmarket/middleware"
	"campusmarket/mongo"
)

type GetPreviewsRequest struct {
	Page int `bson:"page,omitempty"`
}

func GetPreviews(w http.ResponseWriter, r *http.Request) {
	if middleware.Authorize(r.Header.Get("X-CMKT-Authorization")) == http.StatusUnauthorized {
		http.Error(w, "Unauthorized Request", http.StatusUnauthorized)
		return
	}

	if r.Method == "GET" {
		// Get Request Query Data
		pageQuery := r.URL.Query().Get("page")

		// Validate Data
		page, err := strconv.Atoi(pageQuery)

		if err != nil {
			errorMessage := "Invalid number provided: '" + err.Error() + "'"
			http.Error(w, errorMessage, http.StatusBadRequest)
			return
		}

		// Connect to Mongo
		user := os.Getenv("MONGO_USER")
		password := os.Getenv("MONGO_PASSWORD")
		host := os.Getenv("MONGO_HOST")

		client := mongo.Connect(user, password, host)
		defer mongo.Disconnect(client)

		//Get Previews
		coll := mongo.GetCollection(client, "businesses")

		results, err := mongo.PaginatedFind(coll, 7, page)

		if err != nil {
			errorMessage := "Error on PaginatedFind: '" + err.Error() + "'"
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	}
}