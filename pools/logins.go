package pools

import "sync"

type Login struct {
	Username  string
	SessionID string
}

var logins sync.Map

func AddLogin(login *Login) {
	logins.Store(login.Username, login)
}

func RemoveLogin(login *Login) {
	logins.Delete(login.Username)
}
