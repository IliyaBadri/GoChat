package handlers

import (
	"encoding/json"
	"gochat/globals"
	"gochat/messages"
	"gochat/pools"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func HandleWebSocket(responseWriter http.ResponseWriter, request *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(checkOriginRequest *http.Request) bool {
			return true
		},
	}
	connection, err := upgrader.Upgrade(responseWriter, request, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer connection.Close()
	connectionRemoteAddress := connection.RemoteAddr()
	log.Printf("[+] New WebSocket connection established: (%s)", connectionRemoteAddress)

	messageType, data, err := connection.ReadMessage()

	if err != nil {
		log.Printf("[-] WebSocket read message error: (%s) : %s", connectionRemoteAddress, err)
		return
	}

	var identificationMessage messages.IdentificationMessage
	err = json.Unmarshal(data, &identificationMessage)
	if err != nil {
		log.Println("[-] WebSocket JSON decode error:", err)
		connection.Close()
		return
	}

	session := globals.Session{UserID: identificationMessage.UserID, Connection: connection}

	pools.AddSession(&session)

	for {
		err = connection.WriteMessage(messageType, data)
		if err != nil {
			log.Printf("[-] WebSocket write message error: (%s) : %s", connectionRemoteAddress, err)
			break
		}
	}
}
