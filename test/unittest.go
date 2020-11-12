package main

import (
	"log"
	"medods-auth/controllers"
	"medods-auth/models"
)

func main() {
	guid := "Bob"

	accessTokenA, refreshTokenA, err := models.CreateWebTokenPair(guid)
	if err != nil {
		log.Fatal(err)
	}

	tokenPairA, err := models.CreateDBTokenPair(accessTokenA, refreshTokenA)
	if err != nil {
		log.Fatal(err)
	}

	accessTokenB, refreshTokenB, err := models.CreateWebTokenPair(guid)
	if err != nil {
		log.Fatal(err)
	}

	tokenPairB, err := models.CreateDBTokenPair(accessTokenB, refreshTokenB)
	if err != nil {
		log.Fatal(err)
	}

	accessTokenC, refreshTokenC, err := models.CreateWebTokenPair(guid)
	if err != nil {
		log.Fatal(err)
	}

	tokenPairC, err := models.CreateDBTokenPair(accessTokenC, refreshTokenC)
	if err != nil {
		log.Fatal(err)
	}

	err = controllers.AddDBTokenPair(guid, tokenPairA)
	if err != nil {
		log.Fatal(err)
	}
	err = controllers.AddDBTokenPair(guid, tokenPairB)
	if err != nil {
		log.Fatal(err)
	}
	err = controllers.AddDBTokenPair(guid, tokenPairC)
	if err != nil {
		log.Fatal(err)
	}

	err = controllers.DeleteDBTokenPair(guid, tokenPairB)
	if err != nil {
		log.Fatal(err)
	}

	err = controllers.RefreshDBTokenPair(guid, tokenPairC, tokenPairB)
	if err != nil {
		log.Fatal(err)
	}

	err = controllers.DeleteAllDBTokenPair(guid)
	if err != nil {
		log.Fatal(err)
	}

}
