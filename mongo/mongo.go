package mongo

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"campusmarket/models"
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

	err = client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

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

func UpdateById(collection *mongo.Collection, id string, data primitive.D) {
	update := bson.D{{Key: "$set", Value: data}}

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	result, err := collection.UpdateByID(context.TODO(), objectId, update)

	if err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Println("No error")
	}

	if result != nil {
		fmt.Printf("UpsertedID: %v\n", result.UpsertedID)
	} else {
		fmt.Println("No Result")
	}
}

func PaginatedFind(collection *mongo.Collection, limit int, page int,) ([]models.Business, error) {
	result := make([]models.Business, 0)
 
	l := int64(limit)
	skip := int64(page * limit - limit)
	fOpt := options.FindOptions{Limit: &l, Skip: &skip}

	curr, err := collection.Find(context.TODO(), bson.D{}, &fOpt)

	if err != nil {
	   return result, err
	}
 
	for curr.Next(context.TODO()) {
	   var el models.Business

	   if err := curr.Decode(&el); err != nil {
		  fmt.Println(err)
	   }
 
	   result = append(result, el)
	}
 
	return result, nil
 }