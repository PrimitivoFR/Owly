package common_mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoORM interface {
	ORMInsertOne(document interface{}) (*mongo.InsertOneResult, error)
	CheckConnectivity() error
}

type MongoEntity struct {
	c **mongo.Collection
}

func (collection *MongoEntity) ORMInsertOne(document interface{}) (*mongo.InsertOneResult, error) {

	if err := collection.CheckConnectivity(); err != nil {
		return nil, err
	}

	collec := collection.c
	entity := *collec

	// TODO : check what to do with the return, instead of making it "_"
	res, errInsert := entity.InsertOne(context.Background(), document)

	return res, errInsert
}

func (collection *MongoEntity) CheckConnectivity() error {

	collec := collection.c
	entity := *collec

	if collection == nil || collection.c == nil || entity == nil {
		errSetup := SetupMongoDB()
		if errSetup != nil {
			return errSetup
		}
	}
	return nil
}
