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
func GeneratePBKDF2() (string, error) {
	log.SetPrefix("GeneratePBKDF2():")
	var builder strings.Builder

	for i := 0; i <= 4; i++ {
		builder.WriteString(getRandomWord())
		number := strconv.Itoa(getRandomNumber())
		builder.WriteString(number)
	}
	// next, we need to build a string
	// the string should have 5 random words,
	// separated by 4 random numbers

	// use getRandomWord() and getRandomInt()
	// to generate a random string like:
	// example: trilogy3lectured5chubbiness1toads9breaths
	// we may use a stringbuilder for this

	// generate a password using the PBKDF2 algorithm

	return builder.String(), nil
}

func generatePassword(input int64) (string, error) {
	log.SetPrefix("generatePassword():")
	b, err := generateBytes(input)
	return base64.URLEncoding.EncodeToString(b), err
}

// This function returns (secure) random bytes
func generateBytes(input int64) ([]byte, error) {
	log.SetPrefix("generateBytes():")
	b := make([]byte, input)
	_, err := crypto.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetRandomWord does something
func getRandomWord() string {

	rand.Seed(time.Now().Unix())
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

	// get a random number, max len(lines)
	n := rand.Intn(len(lines))

	// return a random word
	return lines[n]
}

func getRandomNumber() int {

	log.SetPrefix("getRandomNumber(): ")

	rand.Seed(time.Now().Unix())
	// generate a random number and return it
	n := rand.Intn(10)

	return n
}
