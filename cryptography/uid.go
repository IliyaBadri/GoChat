package cryptography

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	return uuid.New().String()
}

func GenerateURID() string {
	safeToken := GenerateSafeToken(16)
	timestamp := time.Now().Unix()
	token := fmt.Sprintf("%d-%s", timestamp, safeToken)
	return token
}

func GenerateSafeToken(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	byteSlice := make([]byte, length)
	_, err := rand.Read(byteSlice)
	if err != nil {
		log.Printf("[-] There was an error generating a safe token with length: %s", fmt.Sprint(length))
		return ""
	}

	for index, the_byte := range byteSlice {
		byteSlice[index] = charset[the_byte%byte(len(charset))]
	}

	return string(byteSlice)
}
