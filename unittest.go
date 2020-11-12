package main

import (
	"./controllers"
	"./models"
	"log"
)

func main() {
	guid := "Bob"

	accessTokenA, refreshTokenA, err := models.CreateTokenPair(guid)
	if err != nil {
		log.Fatal(err)
	}

	tokenPairA, err := models.CreateDBTokenPair(accessTokenA, refreshTokenA)
	if err != nil {
		log.Fatal(err)
	}

	accessTokenB, refreshTokenB, err := models.CreateTokenPair(guid)
	if err != nil {
		log.Fatal(err)
	}

	tokenPairB, err := models.CreateDBTokenPair(accessTokenB, refreshTokenB)
	if err != nil {
		log.Fatal(err)
	}

	accessTokenC, refreshTokenC, err := models.CreateTokenPair(guid)
	if err != nil {
		log.Fatal(err)
	}

	tokenPairC, err := models.CreateDBTokenPair(accessTokenC, refreshTokenC)
	if err != nil {
		log.Fatal(err)
	}

	err = controllers.AddTokenPair(guid, tokenPairA)
	if err != nil {
		log.Fatal(err)
	}
	err = controllers.AddTokenPair(guid, tokenPairB)
	if err != nil {
		log.Fatal(err)
	}
	err = controllers.AddTokenPair(guid, tokenPairC)
	if err != nil {
		log.Fatal(err)
	}

	err = controllers.DeleteTokenPair(guid, tokenPairB)
	if err != nil {
		log.Fatal(err)
	}

	err = controllers.RefreshTokenPair(guid, tokenPairC, tokenPairB)
	if err != nil {
		log.Fatal(err)
	}

	err = controllers.DeleteAllTokenPair(guid)
	if err != nil {
		log.Fatal(err)
	}

}
