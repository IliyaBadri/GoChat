package globals

type Request struct {
	ID      string
	Session *Session
	Data    []byte
}
