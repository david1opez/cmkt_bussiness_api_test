package mongo

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(user string, password string, host string) *mongo.Client {
	const protocol = "mongodb+srv://"

	uri := protocol + user + ":" + password + host

	serverAPI := options.ServerAPI(options.ServerAPIVersion1) // Use the SetServerAPIOptions() method to set the Stable API version to 1
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI).SetCompressors([]string{"snappy"})
	client, err := mongo.Connect(context.TODO(), opts) // Create a new client and connect to the server

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	var result bson.M // Send a ping to confirm a successful connection

	fmt.Println("Ran result bson.M")

	err = client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result)

	fmt.Println("Aajmjofasjafsjafsk")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client
}

func Disconnect(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func GetCollection(client *mongo.Client, collName string) *mongo.Collection {
	db := os.Getenv("DATABASE")
	coll := client.Database(db).Collection(collName)

	return coll
}

func InsertOne(collection *mongo.Collection, document interface{}) {
	result, err := collection.InsertOne(context.TODO(), document)

	if err != nil {
		panic(err)
	}

	if result != nil {
		fmt.Printf("result.InsertedID: %v\n", result.InsertedID)
	}
}
