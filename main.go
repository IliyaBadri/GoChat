package main

import (
	"gochat/handlers"
	"log"
	"net/http"
	"os"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func main() {
	if !PathExists("./static") {
		log.Fatal("[-] No static path was found! Terminating.")
		return
	}

	staticServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", staticServer)
	http.HandleFunc("/api/session/", handlers.HandleSession)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
