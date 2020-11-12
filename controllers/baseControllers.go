package controllers

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"medods-auth/models"
	"medods-auth/utils"
	"strings"
)

func GetUserSession(guid string) (models.UserSession, error) {
	var session models.UserSession
	collection := models.GetCollection()

	filter := bson.D{{"guid", guid}}
	err := collection.FindOne(context.TODO(), filter).Decode(&session)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.UserSession{GUID: guid}, nil
		}
		return session, err
	}

	return session, err
}

func SaveUserSession(session models.UserSession) error {
	var dbSession mongo.Session
	var err error

	collection := models.GetCollection()

	dbSession, err = models.GetDB().StartSession()
	if err != nil {
		return err
	}
	err = dbSession.StartTransaction()
	if err != nil {
		return err
	}
	err = mongo.WithSession(context.TODO(), dbSession, func(sc mongo.SessionContext) error {
		opts := options.Replace().SetUpsert(true)
		filter := bson.D{{"guid", session.GUID}}
		result, err := collection.ReplaceOne(context.TODO(), filter, session, opts)
		if err != nil {
			return err
		}
		if result.MatchedCount != 0 {
			fmt.Println("matched and replaced an existing document")
		}
		if result.UpsertedCount != 0 {
			fmt.Printf("inserted a new document with ID %v\n", result.UpsertedID)
		}
		err = dbSession.CommitTransaction(sc)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	dbSession.EndSession(context.TODO())

	return nil
}

func AddDBTokenPair(guid string, tokenPair models.DBTokenPair) error {
	session, err := GetUserSession(guid)
	if err != nil {
		return err
	}

	session.Tokens = append(session.Tokens, tokenPair)
	err = SaveUserSession(session)

	utils.PrintStruct(session)

	return err
}

func RefreshDBTokenPair(guid string, oldTokenPair models.DBTokenPair, newTokenPair models.DBTokenPair) error {
	session, err := GetUserSession(guid)
	if err != nil {
		return err
	}

	for i := 0; i < len(session.Tokens); i++ {
		if strings.Compare(session.Tokens[i].RefreshToken, oldTokenPair.RefreshToken) == 0 {
			session.Tokens[i] = newTokenPair
		}
	}

	err = SaveUserSession(session)

	utils.PrintStruct(session)

	return err
}

func DeleteDBTokenPair(guid string, tokenPair models.DBTokenPair) error {
	session, err := GetUserSession(guid)
	if err != nil {
		return err
	}

	for i := 0; i < len(session.Tokens); i++ {
		if strings.Compare(session.Tokens[i].RefreshToken, tokenPair.RefreshToken) == 0 {
			session.Tokens[i] = session.Tokens[len(session.Tokens)-1]
			session.Tokens = session.Tokens[:len(session.Tokens)-1]
			break
		}
	}

	err = SaveUserSession(session)

	utils.PrintStruct(session)

	return err
}

func DeleteAllDBTokenPair(guid string) error {
	session, err := GetUserSession(guid)
	if err != nil {
		return err
	}

	session.Tokens = nil

	err = SaveUserSession(session)

	utils.PrintStruct(session)

	return err
}
