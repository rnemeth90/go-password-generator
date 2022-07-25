package generator

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateRandomString generates a random secure string
func GenerateRandomString(length int64) (string, error) {
	password, err := generatePassword(length)
	if err != nil {
		return "", err
	}
	return password, err
}

// GeneratePBCKS5 generates a PBCKS5 compliant password
func GeneratePBKDF2() (string, error) {
	var password string

	// generate a password using the PBKDF2 algorithm

	return password, nil
}

func generatePassword(input int64) (string, error) {
	b, err := generateBytes(input)
	return base64.URLEncoding.EncodeToString(b), err
}

// This function returns (secure) random bytes
func generateBytes(input int64) ([]byte, error) {
	b := make([]byte, input)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
