package handler
 
import (
	"fmt"
	"net/http"
	"os"

	"campusmarket/mongo"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")

	client := mongo.Connect(user, password, host)
	defer mongo.Disconnect(client)

	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}
