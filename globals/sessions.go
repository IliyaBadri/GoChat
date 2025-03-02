package globals

import (
	"github.com/gorilla/websocket"
)

type Session struct {
	UserID     string
	Connection *websocket.Conn
}
