package common_mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoORM interface {
	ORMInsertOne(document interface{}) error
	CheckConnectivity() error
}

type MongoCollection struct {
	*mongo.Collection
}

func (collection *MongoCollection) ORMInsertOne(document interface{}) error {

	log.Println("collection", collection)

	if err := collection.CheckConnectivity(); err != nil {
		return err
	}

	// TODO : check what to do with the return, instead of making it "_"
	_, errInsert := collection.InsertOne(context.Background(), document)

	return errInsert
}

func (collection *MongoCollection) CheckConnectivity() error {

	log.Println("check co")

	if collection == nil {
		errSetup := SetupMongoDB()
		if errSetup != nil {
			return errSetup
		}
	}
	return nil
}
