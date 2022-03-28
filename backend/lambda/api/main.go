package main

import (
	"context"
	"sync"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/holmes89/ibis/internal/database"
	v1 "github.com/holmes89/ibis/internal/handlers/rest/v1"
	"github.com/holmes89/ibis/internal/handlers/rest/v1/service"
)

var (
	muxAdapter *gorillamux.GorillaMuxAdapter
	once       sync.Once
)

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return muxAdapter.ProxyWithContext(ctx, request)
}

func setup() {
	db := database.NewConnection()
	gameRouter := v1.NewGameApiController(service.NewGameServicer(db))
	userRouter := v1.NewUserApiController(service.NewUserServicer(db))
	router := v1.NewRouter(gameRouter, userRouter)
	muxAdapter = gorillamux.New(router)
}

func main() {
	once.Do(setup)
	lambda.Start(handleRequest)
}
