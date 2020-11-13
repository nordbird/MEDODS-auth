package utils

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func GetHash(text string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckHash(text, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))
}

func PrintStruct(x interface{}) {
	jsonStr, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("%s\n", string(jsonStr))
}
