package globals

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Session struct {
	UserID     string
	Connection *websocket.Conn
}

var sessionPool sync.Map

func AddSession(session *Session) {
	sessionPool.Store(session.UserID, session)
	SessionsUpdateSignal.Signal()
}

func RemoveSession(session *Session) {
	sessionPool.Delete(session.Connection)
	SessionsUpdateSignal.Signal()
}
