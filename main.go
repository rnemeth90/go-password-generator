package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var length int64 = 16

	arguments := os.Args

	switch len(arguments) {
	case 2:
		length, _ = strconv.ParseInt(os.Args[1], 10, 64)
		if length <= 0 {
			length = 8
		}
	}

	password, err := generatePassword(length)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(password[0:length])
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
