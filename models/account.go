package models

import (
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type Claims struct {
	GUID string `json:"guid"`
	jwt.StandardClaims
}

type WebRequest struct {
	GUID         string `json:"guid"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type WebResponse struct {
	Message string                 `json:"message"`
	Payload map[string]interface{} `json:"payload"`
}

func CreateToken(guid string, expirationDuration time.Duration) (string, error) {
	expirationTime := time.Now().Add(expirationDuration)

	claims := Claims{
		GUID: guid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(os.Getenv("jwt_salt"))
	if err != nil {
		return "Creation access token failed", err
	}

	tokenString = base64.StdEncoding.EncodeToString([]byte(tokenString))

	return tokenString, err
}

func IsValidToken(tokenString string) (bool, string) {
	tokenData, err := base64.StdEncoding.DecodeString(tokenString)
	if err != nil {
		return false, err.Error()
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(string(tokenData), claims, func(token *jwt.Token) (interface{}, error) {
		return os.Getenv("jwt_salt"), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, "Access token signature is invalid"
		}
		return false, err.Error()
	}

	if !token.Valid {
		return false, "Access token is invalid"
	}

	return true, "Access token is valid"
}

func CreateTokenPair(guid string) (string, string, error) {
	accessTokenString, err := CreateToken(guid, 5*time.Minute)
	if err != nil {
		return "", "", err
	}

	refreshTokenString, err := CreateToken(guid, 7*24*time.Hour)
	if err != nil {
		return "", "", err
	}
	return accessTokenString, refreshTokenString, err
}
