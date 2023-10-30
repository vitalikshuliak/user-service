package api

import (
	"encoding/json"
	"net/http"

	"github.com/vitalikshuliak/user-service/database"
	"github.com/vitalikshuliak/user-service/models"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	// Save the uesr to the database
	err := database.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func ReadUser(w http.ResponseWriter, r *http.Request) {
	// Get the ID of the user from the request URL or query parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// Retrieve the user from the database
	user, err := database.ReadUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the user as JSON response
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	/// Get the user ID from the request URL
	vars := mux.Vars(r)
	id := vars["id"]

	// Decode the request body into a user object
	var updatedUser models.User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call the function in the database package to update the user by ID
	err = database.UpdateUser(id, &updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success status
	w.WriteHeader(http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the request URL
	vars := mux.Vars(r)
	id := vars["id"]

	// Call the function in the database package to delete the user by ID
	err := database.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success status
	w.WriteHeader(http.StatusOK)
}
