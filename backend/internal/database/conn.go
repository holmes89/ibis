package database

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/rs/zerolog/log"
)

var (
	tableName = "ibis"
	notfound  *types.ResourceNotFoundException
)

func init() {
	if table := os.Getenv("DYNAMODB_TABLE"); table != "" {
		tableName = table
	}
}

// Conn is the connection to the Dynamodb
type Conn struct {
	db *dynamodb.Client
}

// NewConnection will create new connection to the Dynamodb
func NewConnection() *Conn {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := loadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("unable to load conifg")
	}
	// Create DynamoDB client
	svc := dynamodb.NewFromConfig(cfg)

	return &Conn{
		db: svc,
	}
}

func loadConfig() (aws.Config, error) {
	if conn := os.Getenv("DYNAMODB_ENDPOINT"); conn != "" {
		log.Warn().Msg("using local database connection")
		return config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(("us-east-1")),
			config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://dynamo:8000", SigningRegion: "us-east-1"}, nil
			})))
	}
	return config.LoadDefaultConfig(context.Background())
}
