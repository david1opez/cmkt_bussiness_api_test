package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"campusmarket/middleware"
)

type GetPreviewsRequest struct {
	Page int32 `bson:"page,omitempty"`
}

func GetPreviews(w http.ResponseWriter, r *http.Request) {
	// if middleware.Authorize(r.Header.Get("X-CMKT-Authorization")) == http.StatusUnauthorized {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// }

	if r.Method == "GET" {
		// Get Request Data
		decoder := json.NewDecoder(r.Body)
		var reqBody GetPreviewsRequest

		err := decoder.Decode(&reqBody)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			panic(err)
		}

		fmt.Fprintf(w, "<h1>Get Previews</h1>")

		// Validate Data

		// Get Previews
	}
}
