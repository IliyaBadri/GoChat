package database

import (
	"database/sql"
	"gochat/cryptography"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func UsernameExists(username string) bool {
	the_database := GetDatabase()
	var exists bool
	query := `SELECT 1 FROM users WHERE username = ? LIMIT 1`

	err := the_database.QueryRow(query, username).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		log.Printf("[-] There was an error checking for an existing username in database: %s", err)
		log.Println("[-] Terminating.")
		os.Exit(1)
		return false
	}

	return exists
}

func EmailExists(email string) bool {
	the_database := GetDatabase()
	var exists bool
	query := `SELECT 1 FROM users WHERE email = ? LIMIT 1`

	err := the_database.QueryRow(query, email).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		log.Printf("[-] There was an error checking for an existing email in database: %s", err)
		log.Println("[-] Terminating.")
		os.Exit(1)
		return false
	}

	return exists
}

func InsertUser(username string, email string, password string) {
	passwordHash := cryptography.HashPassword(password)
	the_database := GetDatabase()
	query := `INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)`

	_, err := the_database.Exec(query, username, email, passwordHash)
	if err != nil {
		log.Printf("[-] There was an error adding a user to the database: %s", err)
		log.Println("[-] Terminating.")
		os.Exit(1)
		return
	}
}
