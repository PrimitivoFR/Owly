package common_mongo

import (
	"primitivofr/owly/server/common/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneChatroomCollection(chatroomMongo models.Chatroom) (*mongo.InsertOneResult, error) {

	var m MongoORM = &MongoEntity{&ChatroomCollection}

	return m.ORMInsertOne(chatroomMongo)

}
