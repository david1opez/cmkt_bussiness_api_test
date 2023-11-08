package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"campusmarket/mongo"
	"campusmarket/models"
)

func loadEnv() {
	if err := godotenv.Load(".env.dev"); err != nil {
		fmt.Println("Error loading .env file")
		return
	}
}

func main() {
	loadEnv()

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
}