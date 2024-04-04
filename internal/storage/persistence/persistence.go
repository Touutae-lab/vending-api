package persistence

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

func NewPostgresConnectionPool() *sql.DB {
	// Construct the connection string
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	// Open the connection
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Printf("Failed to open database connection: %v", err)
		panic("Failed to open database connection")
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Printf("Failed to establish a connection with the database: %v", err)
		panic("Failed to establish a connection with the database")
	}

	// Set up a connection pool
	db.SetMaxOpenConns(25)                 // Set max open connections
	db.SetMaxIdleConns(25)                 // Set max idle connections
	db.SetConnMaxLifetime(5 * time.Minute) // Set max connection lifetime

	return db
}
