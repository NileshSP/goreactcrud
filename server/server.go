package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

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
		fmt.Println("INFO: Go server - No PORT environment variable detected, defaulting to " + port)
	} else {
		fmt.Println("INFO: Go server - PORT environment variable " + port)
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

	// server data endpoints with json response using '/api' prefix
	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/api/", healthCheck).Methods("GET")
	myrouter.HandleFunc("/api/healthcheck", healthCheck).Methods("GET")

	// For go-memdb
	InitializeInMemoryDB()
	myrouter.HandleFunc("/api/users", GetPersons).Methods("GET")
	myrouter.HandleFunc("/api/adduser/{name}/{email}", AddPerson).Methods("POST")
	myrouter.HandleFunc("/api/deleteuser/{name}/{email}", DeletePerson).Methods("DELETE")
	myrouter.HandleFunc("/api/updateuser/{name}/{email}", UpdatePerson).Methods("PUT")

	// For in sqlite
	// InitializeSqlite();
	// myrouter.HandleFunc("/api/users", AllUsers).Methods("GET")
	// myrouter.HandleFunc("/api/adduser/{name}/{email}", NewUser).Methods("POST")
	// myrouter.HandleFunc("/api/deleteuser/{name}/{email}", DeleteUser).Methods("DELETE")
	// myrouter.HandleFunc("/api/updateuser/{name}/{email}", UpdateUser).Methods("PUT")

	// serve static files(from react) for '/' prefix
	myrouter.PathPrefix("/").Handler(http.FileServer(http.Dir("./client/build")))
	log.Fatal(http.ListenAndServe(*listen, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(myrouter)))
}

func main() {
	fmt.Println("Go ORM server initialized")
	handleRequests()
}
