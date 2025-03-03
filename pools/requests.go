package pools

import (
	"gochat/globals"
	"gochat/signals"
	"sync"
)

var requests sync.Map

func AddRequest(request *globals.Request) {
	requests.Store(request.ID, request)
	signals.Session.Signal()
}

func RemoveRequest(request *globals.Request) {
	requests.Delete(request.ID)
	signals.Session.Signal()
}
