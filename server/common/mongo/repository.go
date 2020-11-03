package common_mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoORM interface {
	ORMInsertOne(document interface{}) (*mongo.InsertOneResult, error)
	ORMFindOneById(id interface{}) (*mongo.SingleResult, error)
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

func (collection *MongoEntity) ORMFindOneById(id interface{}) (*mongo.SingleResult, error) {
	// id could be an ObjectID or a simple string ID. No conversion is done here

	if err := collection.CheckConnectivity(); err != nil {
		return nil, err
	}

	collec := collection.c
	entity := *collec

	// find document by _id
	result := entity.FindOne(
		context.Background(),
		bson.M{"_id": id},
	)

	return result, nil

}
