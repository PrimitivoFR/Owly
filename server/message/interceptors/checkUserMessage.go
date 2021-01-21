package interceptors

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"github.com/primitivofr/owly/server/message/messagepb"

	"github.com/olivere/elastic"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// IsUserAuthorOfMessage allows to check wether the given user is the author of the message.
// It requires an elastic client, the id of the chatroom, the concerned message, and the id of the user.
// Will return an error if the user is not authorized.
func IsUserAuthorOfMessage(client *elastic.Client, chatroomID string, messageID string, userID string) error {
	doc, err := client.Get().
		Index(chatroomID).
		Id(messageID).
		Type("_doc").
		Do(context.Background())

	if err != nil {
		log.Printf("Error while getting message in ES %v", err)
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while getting  message in ES %v", err),
		)
	}
	var message messagepb.Message
	data := *doc.Source
	json.Unmarshal(data, &message)

	// Not authorized
	if message.AuthorUUID != userID {
		log.Printf("This user %v is not the author of the message %v. He can't do anything it", userID, messageID)
		return status.Errorf(
			codes.PermissionDenied,
			fmt.Sprintf("This user %v is not the author of the message %v. He can't do anything with it", userID, messageID),
		)
	}

	return nil
}
