package globals

import "sync"

var SessionsUpdateSignal = sync.NewCond(&sync.Mutex{})
