package signals

import "sync"

var Session = sync.NewCond(&sync.Mutex{})
