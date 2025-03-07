package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func Initialize(databasePath string) {
	newDatabase, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		log.Printf("[-] There was an error initializing the database: %s", err)
		log.Println("[-] Terminating.")
		os.Exit(1)
		return
	}
	err = newDatabase.Ping()
	if err != nil {
		log.Printf("[-] There was an error pinging the database: %s", err)
		os.Exit(1)
		return
	}
	database = newDatabase
}

func GetDatabase() *sql.DB {
	if database == nil {
		log.Println("[-] The database has been accessed before being initialized.")
		log.Println("[-] Terminating.")
		os.Exit(1)
		return nil
	}
	return database
}

func CreateTables() {
	the_database := GetDatabase()
	tableCreationQueries := []string{
		`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		email TEXT NOT NULL,
		password_hash TEXT NOT NULL
		);`,
	}

	for _, query := range tableCreationQueries {
		_, err := the_database.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Table 'users' created successfully!")
}
