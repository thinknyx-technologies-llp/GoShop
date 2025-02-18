package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

var DB *sql.DB // Use *sql.DB instead of *gorm.DB

// Connect initializes and returns the *sql.DB instance
func Connect() *sql.DB {
	dsn := "root:Madhuri123!@#@tcp(localhost:3306)/shop"
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	// Check the connection
	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	return DB
}
