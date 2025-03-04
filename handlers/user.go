package handlers

import (
	"encoding/json"
	"fmt"
	"gochat/cryptography"
	"gochat/database"
	"gochat/network"
	"gochat/pools"
	"gochat/validators"
	"io"
	"net/http"
)

func sendSignupError(responseWriter http.ResponseWriter, errorMessage string, statusCode int) {
	responseWriter.WriteHeader(statusCode)
	errorResponse := network.SignupResponse{
		Successful: false,
		Error:      errorMessage,
		SessionID:  "",
	}
	err := json.NewEncoder(responseWriter).Encode(errorResponse)
	if err != nil {
		http.Error(responseWriter, "Failed to encode JSON.", http.StatusInternalServerError)
	}
}

func User(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	if request.Method == http.MethodPost {
		body, err := io.ReadAll(request.Body)
		println(body)
		if err != nil {
			sendSignupError(responseWriter, "Failed to read request body.", http.StatusBadRequest)
			request.Body.Close()
			return
		}

		var signupRequest network.SignupRequest
		err = json.Unmarshal(body, &signupRequest)
		if err != nil {
			sendSignupError(responseWriter, "Invalid JSON.", http.StatusBadRequest)
			request.Body.Close()
			return
		}

		isEmail := validators.IsEmail(signupRequest.Email)

		if !isEmail {
			sendSignupError(responseWriter, "Invalid Email.", http.StatusBadRequest)
			request.Body.Close()
			return
		}

		isUsername, usernameValidationErrorMessage := validators.IsUsername(signupRequest.Username)

		if !isUsername {
			sendSignupError(responseWriter, fmt.Sprintf("Invalid Username: %s", usernameValidationErrorMessage), http.StatusBadRequest)
			request.Body.Close()
			return
		}

		isPassword, passwordValidationErrorMessage := validators.IsPassword(signupRequest.Password)
		if !isPassword {
			sendSignupError(responseWriter, fmt.Sprintf("Invalid Password: %s", passwordValidationErrorMessage), http.StatusBadRequest)
			request.Body.Close()
			return
		}

		usernameExists := database.UsernameExists(signupRequest.Username)
		if usernameExists {
			sendSignupError(responseWriter, "Username already exists.", http.StatusConflict)
			request.Body.Close()
			return
		}

		emailExists := database.EmailExists(signupRequest.Email)
		if emailExists {
			sendSignupError(responseWriter, "Email already exists.", http.StatusConflict)
			request.Body.Close()
			return
		}

		sessionID := cryptography.GenerateUUID()

		database.InsertUser(signupRequest.Username, signupRequest.Email, signupRequest.Password)

		login := pools.Login{Username: signupRequest.Username, SessionID: sessionID}
		pools.AddLogin(&login)

		responseWriter.WriteHeader(http.StatusOK)
		response := network.SignupResponse{
			Successful: true,
			Error:      "",
			SessionID:  sessionID,
		}
		err = json.NewEncoder(responseWriter).Encode(response)
		if err != nil {
			http.Error(responseWriter, "Failed to encode JSON.", http.StatusInternalServerError)
		}
		request.Body.Close()
		return
	} else {
		sendSignupError(responseWriter, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}
}
