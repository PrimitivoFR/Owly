package common_mongo

import (
	"github.com/primitivofr/owly/server/common/models"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneChatroomCollection(chatroomMongo models.Chatroom) (*mongo.InsertOneResult, error) {

	var m MongoORM = &MongoEntity{&ChatroomCollection}

	return m.ORMInsertOne(chatroomMongo)

}

func IsChatroomOwner(userID string, chatroomID string) (bool, error) {

	if ChatroomCollection == nil {
		errSetup := SetupMongoDB()
		if errSetup != nil {
			return false, errSetup
		}
	}

	OID, errOID := primitive.ObjectIDFromHex(chatroomID)
	if errOID != nil {
		return false, errOID
	}

	var chartoomResult models.Chatroom
	errFind := ChatroomCollection.FindOne(context.Background(), bson.M{"_id": OID}).Decode(&chartoomResult)
	if errFind != nil {
		return false, errFind
	}

	userIsChatroomOwner := userID == chartoomResult.Owner

	return userIsChatroomOwner, nil
}

func ChangeChatroomOwner(newOwnerId string, chatroomId string) error {
	if ChatroomCollection == nil {
		errSetup := SetupMongoDB()
		if errSetup != nil {
			return errSetup
		}
	}

	chatroomOID, errOID := primitive.ObjectIDFromHex(chatroomId)
	if errOID != nil {
		return errOID
	}

	// find chatroom by OID
	var chatroomResult models.Chatroom
	findErr := ChatroomCollection.FindOne(
		context.Background(),
		bson.M{"_id": chatroomOID},
	).Decode(&chatroomResult)

	if findErr != nil {
		return findErr
	}

	// update chatroom in DB with new chatroom owner
	_, updateChatroomErr := ChatroomCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": chatroomOID},
		bson.M{"$set": bson.M{"owner": newOwnerId}},
	)
	if updateChatroomErr != nil {
		return updateChatroomErr
	}

	return nil // no error, everything went well
}

func PopUserInChatroomCollection(userID string, chatroomID string) error {

	if ChatroomCollection == nil {
		errSetup := SetupMongoDB()
		if errSetup != nil {
			return errSetup
		}
	}

	chatroomOID, errOID := primitive.ObjectIDFromHex(chatroomID)
	if errOID != nil {
		return errOID
	}

	// find chatroom by OID
	var chatroomResult models.Chatroom
	findErr := ChatroomCollection.FindOne(
		context.Background(),
		bson.M{"_id": chatroomOID},
	).Decode(&chatroomResult)

	if findErr != nil {
		return findErr
	}

	// remove user from the user list in the chatroom object
	for i, id := range chatroomResult.Users {
		if id == userID {
			chatroomResult.Users = append(chatroomResult.Users[:i], chatroomResult.Users[i+1:]...)
			break
		}
	}

	// update chatroom in DB
	_, updateChatroomERR := ChatroomCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": chatroomOID},
		bson.M{"$set": bson.M{"users": chatroomResult.Users}},
	)
	if updateChatroomERR != nil {
		return updateChatroomERR
	}

	return nil // no error, everything went well
}

func FindOneChatroomCollection(chatroomID string) (*models.Chatroom, error) {
	// We need ID here, it will automatically convert to ObjectID

	var m MongoORM = &MongoEntity{&ChatroomCollection}

	chatroomOID, errOID := primitive.ObjectIDFromHex(chatroomID)
	if errOID != nil {
		return nil, errOID
	}

	// find chatroom by OID
	result, err := m.ORMFindOneById(chatroomOID)
	if err != nil {
		return nil, err
	}

	var chatroomResult models.Chatroom

	if result == nil {
		return &chatroomResult, nil
	}

	err = result.Decode(&chatroomResult)

	if err != nil {
		return nil, err
	}

	return &chatroomResult, nil
}

func DeleteOneChatroomCollection(chatroomID string) error {

	if ChatroomCollection == nil {
		errSetup := SetupMongoDB()
		if errSetup != nil {
			return errSetup
		}
	}

	chatroomOID, errOID := primitive.ObjectIDFromHex(chatroomID)
	if errOID != nil {
		return errOID
	}

	_, deleteErr := ChatroomCollection.DeleteOne(
		context.Background(),
		bson.M{"_id": chatroomOID},
	)
	if deleteErr != nil {
		return deleteErr
	}

	return nil // no error, everything went well
}
