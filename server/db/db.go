package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "aarush"
		password = "postgres"
		dbname   = "testdb"
	)

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL!")

	// Create users table if it doesn't exist
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL
		);
	`)
	if err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	} else {
		fmt.Println("🛠️ Users table initialized.")
	}

	// Create files table if it doesn't exist
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS files (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			filename TEXT NOT NULL,
			content TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW()
		);
	`)
	if err != nil {
		log.Fatalf("Failed to create files table: %v", err)
	} else {
		fmt.Println("🛠️ Files table initialized.")
	}

	fmt.Println("🛠️ Tables initialized.")
}

func RetrieveSessionID(userID string) (string, error) {
	var sessionID string
	query := `SELECT session_id FROM users WHERE id = $1`

	err := DB.QueryRow(query, userID).Scan(&sessionID)
	if err != nil {
		if err == sql.ErrNoRows {
			// No row found with this id
			return "", nil
		}
		return "", fmt.Errorf("failed to retrieve session ID: %v", err)
	}

	return sessionID, nil
}
