package handler

import (
	"fmt"
	"net/http"
	"os"

	"campusmarket/mongo"
	"campusmarket/models"
)

func Handler(w http.ResponseWriter, r *http.Request) {
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

	fmt.Fprintf(w, "<h1>Campus Market Business API</h1>")
}
