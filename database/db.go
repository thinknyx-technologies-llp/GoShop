package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB
var DBConnected = false // Flag to track DB connection status

func Connect() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠ Could not load .env file, using default values")
	}

	// Fetch DB credentials from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Construct DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open MySQL connection
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Database connection failed:", err)
		DBConnected = false
		return
	}

	// Check if the database is reachable
	if err = DB.Ping(); err != nil {
		log.Println("Database is unreachable:", err)
		DBConnected = false
		return
	}

	DBConnected = true
	fmt.Println("Database connected successfully!")
}
