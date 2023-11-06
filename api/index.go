package handler
 
import (
  "fmt"
  "net/http"
  "context"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func loadEnv() {
	if err := godotenv.Load(".env.dev"); err != nil {
        fmt.Println("Error loading .env file")
        return
    }
}

func connectToMongo(user string, password string, host string) {
  const protocol = "mongodb+srv://"

	uri := protocol + user + ":" + password + host

	fmt.Println("Connecting to Mongo...")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1) // Use the SetServerAPIOptions() method to set the Stable API version to 1
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI).SetCompressors([]string{"snappy"})
	client, err := mongo.Connect(context.TODO(), opts) // Create a new client and connect to the server

	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	var result bson.M // Send a ping to confirm a successful connection

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
 
func Handler(w http.ResponseWriter, r *http.Request) {
  loadEnv()

	user := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")

	connectToMongo(user, password, host)
  
  fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}
