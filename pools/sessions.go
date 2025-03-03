package pools

import (
	"gochat/globals"
	"gochat/signals"
	"sync"
)

var sessions sync.Map

func AddSession(session *globals.Session) {
	sessions.Store(session.ID, session)
	signals.Session.Signal()
}

func RemoveSession(session *globals.Session) {
	sessions.Delete(session.ID)
	signals.Session.Signal()
}
