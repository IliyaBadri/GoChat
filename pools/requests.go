package pools

import (
	"gochat/signals"
	"sync"
)

type Request struct {
	ID      string
	Session *Session
	Data    []byte
}

var requests sync.Map

func AddRequest(request *Request) {
	requests.Store(request.ID, request)
	signals.Request.Signal()
}

func RemoveRequest(request *Request) {
	requests.Delete(request.ID)
	signals.Request.Signal()
}
