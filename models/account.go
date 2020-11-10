package models

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"strings"
	"time"
)

type AccessClaims struct {
	GUID string `json:"guid"`
	jwt.StandardClaims
}

type RefreshClaims struct {
	AccessToken string `json:"access_token"`
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

func CreateAccessToken(guid string) (string, error) {
	accessExpirationTime := time.Now().Add(3 * 24 * time.Hour)

	accessClaims := AccessClaims{
		GUID: guid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessExpirationTime.Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, accessClaims)
	accessTokenString, err := accessToken.SignedString(os.Getenv("jwt_salt"))
	if err != nil {
		return "Creation access token failed", err
	}

	return accessTokenString, err
}

func IsValidAccessToken(tokenString string) (bool, string) {
	claims := &AccessClaims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return os.Getenv("jwt_salt"), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, "Access token signature is invalid"
		}
		return false, err.Error()
	}

	if !tkn.Valid {
		return false, "Access token is invalid"
	}

	return true, "Access token is valid"
}

func CreateRefreshToken(accessTokenString string) (string, error) {
	refreshExpirationTime := time.Now().Add(6 * time.Hour)

	refreshClaims := RefreshClaims{
		AccessToken: accessTokenString,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpirationTime.Unix(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS512, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(os.Getenv("jwt_salt"))
	if err != nil {
		return "Creation refresh token failed", err
	}

	return refreshTokenString, err
}

func IsValidRefreshToken(refreshTokenString string, accessTokenString string) (bool, string) {
	claims := &RefreshClaims{}

	tkn, err := jwt.ParseWithClaims(refreshTokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return os.Getenv("jwt_salt"), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, "Refresh token signature is invalid"
		}
		return false, err.Error()
	}

	if !tkn.Valid {
		return false, "Refresh token is invalid"
	}

	if strings.Compare(accessTokenString, claims.AccessToken) != 0 {
		return false, "Refresh token is not suitable"
	}

	return true, "Refresh token is valid"
}

func CreateTokenPair(guid string) (string, string, error) {
	accessTokenString, err := CreateAccessToken(guid)
	if err != nil {
		return "", "", err
	}

	refreshTokenString, err := CreateRefreshToken(accessTokenString)
	if err != nil {
		return "", "", err
	}
	return accessTokenString, refreshTokenString, err
}
