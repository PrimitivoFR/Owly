package common_elastic

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/olivere/elastic"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Init inits an elasticsearch client with official elasticsearch package
func Init() (*elasticsearch.Client, error) {

	elasticHost := os.Getenv("ELASTIC_HOST")
	// Declare an Elasticsearch configuration
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://" + elasticHost + ":9200",
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

// InitOlivereES init an elasticearch client with Olivere elasticsearch package
func InitOlivereES() (*elastic.Client, error) {
	elasticHost := os.Getenv("ELASTIC_HOST")
	client, err := elastic.NewClient(
		elastic.SetSniff(true),
		elastic.SetURL("http://"+elasticHost+":9200"),
		elastic.SetHealthcheckInterval(5*time.Second),
	)

	if err != nil {
		log.Printf("Elasticsearch connection error %v", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Elasticsearch connection error %v", err),
		)
	}

	return client, nil
}
