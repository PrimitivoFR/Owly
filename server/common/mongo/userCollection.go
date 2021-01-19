package common_mongo

import (
	"primitivofr/owly/server-common/models"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneUserCollection(user_mongo models.UserMongo) (*mongo.InsertOneResult, error) {

	var m MongoORM = &MongoEntity{&UserCollection}

	return m.ORMInsertOne(user_mongo)

}

func FindOneUserCollection(id string) (*models.UserMongo, error) {
	var m MongoORM = &MongoEntity{&UserCollection}

	result, err := m.ORMFindOneById(id)
	if err != nil {
		return nil, err
	}

	var userResult models.UserMongo

	if result == nil {
		return &userResult, nil
	}

	err = result.Decode(&userResult)

	if err != nil {
		return nil, err
	}

	return &userResult, nil

}

func PopChatroomInUserCollection(userID string, chatroomID string) error {

	if UserCollection == nil {
		errSetup := SetupMongoDB()
		if errSetup != nil {
			return errSetup
		}
	}

	// find user by its ID
	var userResult models.UserMongo
	findErr := UserCollection.FindOne(
		context.Background(),
		bson.M{"_id": userID},
	).Decode(&userResult)

	if findErr != nil {
		return findErr
	}

	// remove chatroom from the chatroom list in the user object
	for i, id := range userResult.Chatrooms {
		if id == chatroomID {
			userResult.Chatrooms = append(userResult.Chatrooms[:i], userResult.Chatrooms[i+1:]...)
			break
		}
	}

	// update user in DB
	_, updateUserERR := UserCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": userID}, // filter
		bson.M{"$set": bson.M{"chatrooms": userResult.Chatrooms}}, // update
	)
	if updateUserERR != nil {
		return updateUserERR
	}

	return nil // no error, everything went well
}
