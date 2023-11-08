package handler
 
import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"campusmarket/mongo"
)

func loadEnv() {
	if err := godotenv.Load(".env.dev"); err != nil {
        fmt.Println("Error loading .env file")
        return
    }
}
 
func Handler(w http.ResponseWriter, r *http.Request) {
	loadEnv()

	user := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")

	client := mongo.Connect(user, password, host)
	defer mongo.Disconnect(client)

	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}
