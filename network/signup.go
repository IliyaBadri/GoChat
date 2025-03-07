package network

type SignupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponse struct {
	Successful bool   `json:"successful"`
	Error      string `json:"error"`
	SessionID  string `json:"sessionID"`
}
