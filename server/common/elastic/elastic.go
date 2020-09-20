package common_elastic

import (
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Init() (*elasticsearch.Client, error) {
	// Declare an Elasticsearch configuration
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://elasticsearch:9200",
		},
	}

	// Connect to ElasticSearch
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Printf("Elasticsearch connection error %v", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Elasticsearch connection error: %v", err),
		)
	}

	return client, nil
}
