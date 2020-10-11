package common_mongo

import (
	"primitivofr/owly/server/common/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneUserCollection(user_mongo models.UserMongo) (*mongo.InsertOneResult, error) {

	var m MongoORM = &MongoEntity{&UserCollection}

	return m.ORMInsertOne(user_mongo)

}
