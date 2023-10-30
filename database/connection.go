package database

import (
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	sqlDB *sql.DB
}

func (db *DB) Prepare(query string) (*sql.Stmt, error) {
	return db.sqlDB.Prepare(query)
}

func (db *DB) Close() {
	if db.sqlDB != nil {
		db.sqlDB.Close()
	}
}

type User struct {
	ID      int `gorm:"primaryKey"`
	Name    string
	Surname string
	Phone   string
	Address string
}

func NewConnection() *DB {
	// Open a connection to your database
	dsn := "newuser:password@tcp(testDB:3306)/userdb?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Get the underlying *sql.DB object from *gorm.DB
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to ensure a successful connection
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Check if "users" table exists
	if !db.Migrator().HasTable(&User{}) {
		// Create the "users" table
		err = db.AutoMigrate(&User{})
		if err != nil {
			log.Fatal(err)
		}
		print("Table Created!\n")
	}

	return &DB{sqlDB}
}
