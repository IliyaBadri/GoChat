package handlers

import (
	"encoding/json"
	"net/http"
)

type MessageResponse struct {
	Status  int    `json:"status"`
	Type    int    `json:"type"`
	Message string `json:"message"`
}

func HandleSession(responseWriter http.ResponseWriter, request *http.Request) {

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)

	switch request.Method {
	case http.MethodPost:
		error := request.ParseMultipartForm(10 << 20)

		if error != nil {
			response := MessageResponse{
				Status:  1,
				Type:    0,
				Message: "Internal Error",
			}

			json.NewEncoder(responseWriter).Encode(response)
		}
	}

}
