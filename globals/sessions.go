package globals

import (
	"github.com/gorilla/websocket"
)

type Session struct {
	ID         string
	Connection *websocket.Conn
}
