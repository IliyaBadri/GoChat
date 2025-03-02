package pools

import (
	"gochat/globals"
	"gochat/signals"
	"sync"
)

var sessions sync.Map

func AddSession(session *globals.Session) {
	sessions.Store(session.UserID, session)
	signals.Session.Signal()
}

func RemoveSession(session *globals.Session) {
	sessions.Delete(session.Connection)
	signals.Session.Signal()
}
