package common_mongo

import (
	"context"
	"primitivofr/owly/server/common/models"
)

func checkUserCollectionConnectivity() error {
	if UserCollection == nil {
		errSetup := SetupMongoDB()
		if errSetup != nil {
			return errSetup
		}
	}
	return nil
}

func InsertOneUserCollection(user_mongo models.UserMongo) error {

	// var c = &MongoCollection{UserCollection}
	// var m MongoORM = c
	// errInsert := m.ORMInsertOne(user_mongo)
	// return errInsert

	if err := checkUserCollectionConnectivity(); err != nil {
		return err
	}

	_, errInsert := UserCollection.InsertOne(context.Background(), user_mongo)
	return errInsert
}
