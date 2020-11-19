package common_elastic

import (
	"context"
	"fmt"
	"log"

	"github.com/olivere/elastic"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DoesChatroomIndexExist checks if chatroom index exists in ES.
// Will return an error if not
// Needs an olivere elastic Client.
func DoesChatroomIndexExist(olivereClient *elastic.Client, chatroomID string) error {

	exist, err := olivereClient.IndexExists(chatroomID).Do(context.Background())

	if err != nil {
		log.Printf("Error occured while checking if index %v exists. Error: %v", chatroomID, err)
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error occured while checking if index %v exists. Error: %v", chatroomID, err),
		)
	}

	if !exist {
		log.Printf("Chatroom index %v not found", chatroomID)
		return status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Chatroom index %v not found", chatroomID),
		)
	}

	return nil

}
