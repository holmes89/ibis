package database

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/holmes89/ibis/internal"
	"github.com/rs/zerolog/log"
)

var (
	_ internal.GameRepo = (*Conn)(nil)
)

func (conn *Conn) ListGames(ctx context.Context) ([]internal.Game, error) {
	params := &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		KeyConditionExpression: aws.String("ID = :key"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":key": &types.AttributeValueMemberS{Value: "games"},
		},
		ScanIndexForward: aws.Bool(false),
	}

	var games = make([]internal.Game, 0)

	resp, err := conn.db.Query(ctx, params)
	if err != nil {
		if errors.As(err, &notfound) {
			log.Warn().Msg("no devices found")
			return games, nil
		}
		log.Error().Err(err).Msg("unable to find all devices")
		return nil, errors.New("unable to fetch all devices")
	}
	if err := attributevalue.UnmarshalListOfMaps(resp.Items, &games); err != nil {
		log.Error().Err(err).Msg("unable to unmarshal devices")
		return nil, errors.New("unable to fetch all devices")
	}

	log.Info().Int("entries", len(games)).Msg("found games")
	return games, nil
}
