package cryptography

import (
	"encoding/base64"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("[-] There was an error while hashing a password.")
		log.Println("[-] Terminating.")
		os.Exit(1)
		return ""
	}

	return base64.StdEncoding.EncodeToString(hashedBytes)
}

func VerifyPassword(password string, hash string) bool {
	hashBytes, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		log.Println("[-] There was an error while decoding a hash from base64 string.")
		log.Println("[-] Not verifing the hash.")
		return false
	}
	err = bcrypt.CompareHashAndPassword(hashBytes, []byte(password))
	return err == nil
}
