package main

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	GUID    string `json:"guid"`
	Access  jwt.StandardClaims
	Refresh string `json:"refresh"`
}

type UserSession struct {
	GUID   string   `json:"guid"`
	Tokens []Claims `json:"tokens"`
}

func main() {

}
