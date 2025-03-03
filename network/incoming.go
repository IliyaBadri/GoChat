package network

type IncomingRequestMessage struct {
	RequestType string `json:"requestType"`
	RequestJson string `json:"requestJson"`
}
