package common_mongo

import (
	"context"
	"primitivofr/owly/server/common/models"
)

func InsertOneUserCollection(user_mongo models.UserMongo) error {

	if UserCollection == nil {
		errSetup := SetupMongoDB()
		if errSetup != nil {
			return errSetup
		}
	}

	// TODO : check what to do with the return, instead of making it "_"
	_, errInsert := UserCollection.InsertOne(context.Background(), user_mongo)
	return errInsert
}
