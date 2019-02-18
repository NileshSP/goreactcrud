package main

import (
	"os"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	listen = flag.String("listen", GetPort(), "listen address")
	dir    = flag.String("dir", ".", "directory to serve")
)

// GetPort represents a request to return the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "8081"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Health check endpoint initialized")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"alive": true}`)

	fmt.Println("Health check endpoint completed")
}

func handleRequests() {
	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/", healthCheck).Methods("GET")
	myrouter.HandleFunc("/healthcheck", healthCheck).Methods("GET")
	myrouter.HandleFunc("/users", AllUsers).Methods("GET")
	myrouter.HandleFunc("/adduser/{name}/{email}", NewUser).Methods("POST")
	myrouter.HandleFunc("/deleteuser/{name}/{email}", DeleteUser).Methods("DELETE")
	myrouter.HandleFunc("/updateuser/{name}/{email}", UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(*listen, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(myrouter)))
}

func main() {
	fmt.Println("Go ORM server initialized")
	//InitialMigrations()
	handleRequests()
}
