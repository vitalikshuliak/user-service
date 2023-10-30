package api

import (
	"github.com/gorilla/mux"
	"github.com/vitalikshuliak/user-service/database"
)

func RegisterRoutes(r *mux.Router, db *database.DB) {
	// r.Use(utils.AuthMiddleware)
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", ReadUser).Methods("GET")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
}
