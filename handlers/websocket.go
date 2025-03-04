package handlers

import (
	"encoding/json"
	"gochat/cryptography"
	"gochat/network"
	"gochat/pools"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func WebSocket(responseWriter http.ResponseWriter, request *http.Request) {

	websocketUpgrader := websocket.Upgrader{
		CheckOrigin: func(checkOriginRequest *http.Request) bool {
			return true
		},
	}

	connection, err := websocketUpgrader.Upgrade(responseWriter, request, nil)
	if err != nil {
		log.Printf("[-] WebSocket upgrade error: %s", err)
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
		log.Printf("[-] WebSocket JSON decode error: (%s) : %s", connectionRemoteAddress, err)
		connection.Close()
		return
	}

	session := pools.Session{ID: identificationMessage.SessionID, Connection: connection}
	pools.AddSession(&session)

	for {
		_, incomingRequestData, err := connection.ReadMessage()
		if err != nil {
			log.Printf("[-] WebSocket read message error: (%s) : %s", connectionRemoteAddress, err)
			break
		}

		requestId := cryptography.GenerateURID()
		websocketRequest := pools.Request{ID: requestId, Session: &session, Data: incomingRequestData}
		pools.AddRequest(&websocketRequest)
	}

	pools.RemoveSession(&session)
	connection.Close()
}
