// Entrypoint for API
package main

import (
 	"log"
 	"net/http"
 	"os"
	"github.com/gorilla/handlers"
	"rest-server/store"
)

func main() {
	
	log.Println("************ Demo for JIO Assignment ************")
	 
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT is not set....must be set")
		port = "8000"
		log.Println("Default Port : " + port)
	}
    log.Println("Port is set to : " + port)
	
	router := store.NewRouter() // create routes
	
	// These two lines are important in order to allow access from the front-end side to the methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"}) 
 	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	// Launch server with CORS validations
 	log.Fatal(http.ListenAndServe(":" + port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}