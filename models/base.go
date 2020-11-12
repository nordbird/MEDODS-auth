package models

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"medods-auth/utils"
	"os"
	"time"
)

type DBTokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserSession struct {
	GUID   string        `json:"guid"`
	Tokens []DBTokenPair `json:"tokens"`
}

var db *mongo.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Print(err)
	}

	connectionUri := os.Getenv("connectionUri")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		connectionUri,
	))
	if err != nil {
		log.Fatal(err)
	}
	db = client
}

func GetDB() *mongo.Client {
	return db
}

func GetCollection() *mongo.Collection {
	return GetDB().Database("test").Collection("tokens")
}

func CreateDBTokenPair(accessTokenString string, refreshTokenString string) (DBTokenPair, error) {
	var pair DBTokenPair
	var err error

	pair.AccessToken = accessTokenString
	pair.RefreshToken, err = utils.GetHash(refreshTokenString)
	return pair, err
}
