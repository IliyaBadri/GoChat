package handlers

import (
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
	for {
		messageType, data, err := connection.ReadMessage()
		if err != nil {
			log.Printf("[-] WebSocket read message error: (%s) : %s", connectionRemoteAddress, err)
			break
		}

		err = connection.WriteMessage(messageType, data)
		if err != nil {
			log.Printf("[-] WebSocket write message error: (%s) : %s", connectionRemoteAddress, err)
			break
		}
	}
}
