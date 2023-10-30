package database

import (
	"database/sql"

	"github.com/vitalikshuliak/user-service/models"
)

func CreateUser(user *models.User) error {
	db := NewConnection()
	// Prepare the SQL statement to insert a new user
	stmt, err := db.Prepare("INSERT INTO users (name, phone) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the prepared statement with the user data
	_, err = stmt.Exec(user.Name, user.Phone)
	if err != nil {
		return err
	}
	print("Successfully User Created!\n")

	return nil
}

func ReadUser(id string) (*models.User, error) {
	println("Reading...")
	db := NewConnection()
	// Prepare the SQL statement to select a user by ID
	stmt, err := db.Prepare("SELECT name, phone FROM users WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute the prepared statement with the user ID
	row := stmt.QueryRow(id)

	// Create a new user object to hold the retrieved data
	user := &models.User{}

	// Scan the row and populate the user object with the retrieved data
	err = row.Scan(&user.Name, &user.Phone)
	if err == sql.ErrNoRows {
		return nil, nil // User not found
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUser(id string, user *models.User) error {
	println(("Updating..."))
	db := NewConnection()
	// Prepare the SQL statement to update a user by ID
	stmt, err := db.Prepare("UPDATE users SET name = ?, phone = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the prepared statement with the updated user data and ID
	_, err = stmt.Exec(user.Name, user.Phone, id)
	if err != nil {
		return err
	}
	println("Successfully User Updated!")
	return nil
}

func DeleteUser(id string) error {
	println("Deleting...")
	db := NewConnection()
	// Prepare the SQL statement to delete a user by ID
	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the prepared statement with the user ID
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	println("Successfully User Deleted!")

	return nil
}
