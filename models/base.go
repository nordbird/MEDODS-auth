package models

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var db *mongo.Client

func init() {

	// Load config
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

// возвращает дескриптор объекта DB
func GetDB() *mongo.Client {
	return db
}
