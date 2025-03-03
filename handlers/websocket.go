package handlers

import (
	"encoding/json"
	"gochat/globals"
	"gochat/network"
	"gochat/pools"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func HandleWebSocket(responseWriter http.ResponseWriter, request *http.Request) {
	websocketUpgrader := websocket.Upgrader{
		CheckOrigin: func(checkOriginRequest *http.Request) bool {
			return true
		},
	}
	connection, err := websocketUpgrader.Upgrade(responseWriter, request, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	connectionRemoteAddress := connection.RemoteAddr()
	log.Printf("[+] New WebSocket connection established: (%s)", connectionRemoteAddress)

	_, identificationRequestData, err := connection.ReadMessage()
	if err != nil {
		log.Printf("[-] WebSocket read message error: (%s) : %s", connectionRemoteAddress, err)
		connection.Close()
		return
	}

	var identificationMessage network.IdentificationMessage
	err = json.Unmarshal(identificationRequestData, &identificationMessage)
	if err != nil {
		log.Println("[-] WebSocket JSON decode error:", err)
		connection.Close()
		return
	}

	session := globals.Session{UserID: identificationMessage.UserID, Connection: connection}

	pools.AddSession(&session)

	for {
		_, incomingRequestData, err := connection.ReadMessage()
		if err != nil {
			log.Printf("[-] WebSocket read message error: (%s) : %s", connectionRemoteAddress, err)
			break
		}

		var IncomingRequestMessage network.IncomingRequestMessage
		err = json.Unmarshal(incomingRequestData, &IncomingRequestMessage)
		if err != nil {
			log.Println("[-] WebSocket JSON decode error:", err)
			break
		}

		request := globals.Request{Session: &session, Data: data}
		err = json.Unmarshal(data, &request.Message)
	}
	pools.RemoveSession(&session)
	connection.Close()
}
