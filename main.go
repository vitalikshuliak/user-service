package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vitalikshuliak/user-service/api"
	"github.com/vitalikshuliak/user-service/database"
)

func main() {
	// Initialize the database connection
	db := database.NewConnection()
	defer db.Close()
	print("Successfully connected to DB!\n")

	// Create a new router instance
	router := mux.NewRouter()

	// Register API routes
	api.RegisterRoutes(router, db)

	// Start the server and listen for requests
	go func() {
		if err := http.ListenAndServe(":8080", router); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("Server is running on port 8080")

	select {}
}
