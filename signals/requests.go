package signals

import "sync"

var Request = sync.NewCond(&sync.Mutex{})
