package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"tree/controllers"
	"tree/repositories"
)

// Basic OK route for healthcheck
func ok(w http.ResponseWriter, _ *http.Request) {
	_, err := io.WriteString(w, "ok")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var port string
	if port = os.Getenv("PORT"); port == "" { // get port from env var
		port = "3002" // set port 3002 if env var not specified
	}

	var environment string
	if environment = os.Getenv("ENV"); environment == "" { // get port from env var
		environment = "PROD" // set port 3002 if env var not specified
	}

	// Create a new router to serve routes
	router := mux.NewRouter()

	router.HandleFunc("/tree/ok", ok).Methods("GET")
	if environment != "CI" {
		repositories.InitDatabase()
	}
	controllers.HandleRoutes(router)

	fmt.Println("Server is running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
