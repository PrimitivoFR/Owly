package common_mongo

import (
	"primitivofr/owly/server/common/models"

	"go.mongodb.org/mongo-driver/mongo"
	"context"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/bson"
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

func PopUserInChatroomCollection(userID string, chatroomID string) (error) {

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

	if ChatroomCollection == nil {
        errSetup := SetupMongoDB()
        if errSetup != nil {
            return nil, errSetup
        }
    }

	chatroomOID, errOID := primitive.ObjectIDFromHex(chatroomID)
    if errOID != nil {
		return nil, errOID
	}

	// find chatroom by OID
    var chatroomResult models.Chatroom
    findErr := ChatroomCollection.FindOne(
        context.Background(),
        bson.M{"_id": chatroomOID},
	).Decode(&chatroomResult)
	
	if findErr != nil {
		return nil, findErr
	}

	return &chatroomResult, nil
}

func DeleteOneChatroomCollection(chatroomID string) (error) {

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