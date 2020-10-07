package common_mongo

import (
	"context"
	"primitivofr/owly/server/common/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func checkChatroomCollection() error {
	if ChatroomCollection == nil {
		errSetup := SetupMongoDB()
		if errSetup != nil {
			return errSetup
		}
	}
	return nil
}

func InsertOneChatroomCollection(chatroomMongo models.Chatroom) (*mongo.InsertOneResult, error) {

	if err := checkChatroomCollection(); err != nil {
		return nil, err
	}

	// TODO : check what to do with the return, instead of making it "_"
	res, errInsert := ChatroomCollection.InsertOne(context.Background(), chatroomMongo)
	return res, errInsert
}
