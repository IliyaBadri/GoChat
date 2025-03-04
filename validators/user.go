package validators

import (
	"net/mail"
	"strings"
	"unicode"
)

func IsEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsPassword(password string) (bool, string) {
	var hasUpper, hasLower, hasDigit, hasSpecial bool
	specialChars := "!@#$%^&*()-_=+[]{}|;:'\",.<>?/`~"

	if len(password) < 8 || len(password) > 64 {
		return false, "Password must be between 8 and 64 characters long."
	}

	for _, character := range password {
		switch {
		case unicode.IsUpper(character):
			hasUpper = true
		case unicode.IsLower(character):
			hasLower = true
		case unicode.IsDigit(character):
			hasDigit = true
		case strings.ContainsRune(specialChars, character):
			hasSpecial = true
		case unicode.IsSpace(character):
			return false, "Password must not contain spaces."
		}
	}

	if !hasUpper {
		return false, "Password must contain at least one uppercase letter."
	}
	if !hasLower {
		return false, "Password must contain at least one lowercase letter."
	}
	if !hasDigit {
		return false, "Password must contain at least one digit."
	}
	if !hasSpecial {
		return false, "Password must contain at least one special character."
	}

	return true, ""
}

func IsUsername(username string) (bool, string) {
	if len(username) < 4 || len(username) > 20 {
		return false, "Username must be between 4 and 20 characters long."
	}

	if !unicode.IsLetter(rune(username[0])) {
		return false, "Username must start with a letter."
	}

	for _, character := range username {
		if !(unicode.IsLetter(character) || unicode.IsDigit(character) || character == '_') {
			return false, "Username must only contain uppercase and lowercase letters and digits and underscores."
		}
	}

	return true, ""
}
