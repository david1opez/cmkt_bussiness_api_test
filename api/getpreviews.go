package handler

import (
	"fmt"
	"net/http"
)

type GetPreviewsRequest struct {
	Page int `bson:"page,omitempty"`
}

func GetPreviews(w http.ResponseWriter, r *http.Request) {
	if middleware.Authorize(r.Header.Get("X-CMKT-Authorization")) == http.StatusUnauthorized {
		w.WriteHeader(http.StatusUnauthorized)
	}

	if r.Method == "GET" {
		// Get Request Query Data
		page := r.URL.Query().Get("page")

		fmt.Println(page)

		fmt.Fprintf(w, "<h1>Get Previews</h1>")

		// Validate Data

		// Get Previews
	}
}
