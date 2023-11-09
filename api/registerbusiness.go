package handler

import (
	"fmt"
	"net/http"
	"os"

	"campusmarket/mongo"
	"campusmarket/models"
)

func RegisterBusiness(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		user := os.Getenv("MONGO_USER")
		password := os.Getenv("MONGO_PASSWORD")
		host := os.Getenv("MONGO_HOST")

		client := mongo.Connect(user, password, host)
		defer mongo.Disconnect(client)

		coll := mongo.GetCollection(client, "businesses")

		newB, err := models.NewBusiness(models.Business{
			Name: "",
			Title: "",
			Verified: false,
			Images: [3]string{"", "", ""},
			SmallImages: [3]string{"", "", ""},
			Rating: 0,
			Description: "",
			Location: "",
			Schedule: "",
		})

		if err != nil {
			panic(err.Error())
		}

		mongo.InsertOne(coll, newB)
	} else if r.Method == "GET" {
		fmt.Fprintf(w, "<h1>Register Business</h1>")
	}
}