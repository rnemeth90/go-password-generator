package generator

import (
	crypto "crypto/rand"
	"encoding/base64"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// GenerateRandomString generates a random secure string
func GenerateRandomString(length int64) (string, error) {
	log.SetPrefix("GenerateRandomString():")

	password, err := generatePassword(length)
	if err != nil {
		return "", err
	}
	return password, err
}

// GeneratePBCKS5 generates a PBCKS5 compliant password
func GenerateMemorablePassword() (string, error) {
	// example: trilogy3lectured5chubbiness1toads9breaths
	rand.Seed(time.Now().Unix())
	log.SetPrefix("GenerateMemorablePassword():")

	var builder strings.Builder

	for i := 0; i <= 4; i++ {
		builder.WriteString(getRandomWord())
		number := strconv.Itoa(getRandomNumber())
		builder.WriteString(number)
	}
	return builder.String(), nil
}

func generatePassword(input int64) (string, error) {
	log.SetPrefix("generatePassword():")
	b, err := generateBytes(input)
	return base64.URLEncoding.EncodeToString(b), err
}

func generateBytes(input int64) ([]byte, error) {
	log.SetPrefix("generateBytes():")
	b := make([]byte, input)
	_, err := crypto.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func getRandomWord() string {

	var lines []string
	// set our log prefix
	log.SetPrefix("getRandomWord(): ")

	// open the file
	words, err := ioutil.ReadFile("./words.txt")
	if err != nil {
		log.Fatal("Error opening the file", words)
	}

	// split the []byte into []string
	lines = strings.Split(string(words), "\n")
	n := rand.Intn(len(lines))
	// return a random word
	return lines[n]
}

func getRandomNumber() int {
	log.SetPrefix("getRandomNumber(): ")
	n := rand.Intn(10)
	return n
}
