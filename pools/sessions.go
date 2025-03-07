package pools

import (
	"gochat/signals"
	"sync"

	"github.com/gorilla/websocket"
)

type Session struct {
	ID         string
	Connection *websocket.Conn
}

var sessions sync.Map

func AddSession(session *Session) {
	sessions.Store(session.ID, session)
	signals.Session.Signal()
}

func RemoveSession(session *Session) {
	sessions.Delete(session.ID)
	signals.Session.Signal()
}
