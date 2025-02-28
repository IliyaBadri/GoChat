package main

import (
	"fmt"
	"gochat/handlers"
	"log"
	"net/http"
	"os"
	"path/filepath"
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
	serverPort := 8080
	staticPath, _ := filepath.Abs("./static")
	if !PathExists(staticPath) {
		log.Printf("[-] No static path was found at: %s\n[-] Terminating.", staticPath)
		os.Exit(1)
		return
	}
	staticServer := http.FileServer(http.Dir(staticPath))
	http.Handle("/", staticServer)
	http.HandleFunc("/ws", handlers.HandleWebSocket)
	log.Printf("[+] Started web server on http://localhost:%s/", fmt.Sprint(serverPort))
	log.Print(http.ListenAndServe(":8080", nil))
	os.Exit(1)
}
