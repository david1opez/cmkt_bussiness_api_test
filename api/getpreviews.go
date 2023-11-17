package handler

import (
	"fmt"
	"net/http"

	"campusmarket/middleware"
)

type GetPreviewsRequest struct {
	Page int `bson:"page,omitempty"`
}

func GetPreviews(w http.ResponseWriter, r *http.Request) {
	if middleware.Authorize(r.Header.Get("X-CMKT-Authorization")) == http.StatusUnauthorized {
		w.WriteHeader(http.StatusUnauthorized)
		panic("Unauthorized")
	}

	if r.Method == "GET" {
		// Get Request Query Data
		page := r.URL.Query().Get("page")

		// Validate Data

		// Connect to Mongo
		user := os.Getenv("MONGO_USER")
		password := os.Getenv("MONGO_PASSWORD")
		host := os.Getenv("MONGO_HOST")

		client := mongo.Connect(user, password, host)
		defer mongo.Disconnect(client)

		//Get Previews
		coll := mongo.GetCollection(client, "businesses")

		results := mongo.PaginatedFind(coll, 7, page)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)

		fmt.Fprintf(w, "<h1>Get Previews</h1>")
	}
}
