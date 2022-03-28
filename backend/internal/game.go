package internal

import "context"

type Game struct {
	Type        string `json:"-" dynamodbav:"ID"`
	ID          string `json:"id" dynamodbav:"SK"`
	Title       string `json:"title" dynamodbav:"title"`
	Path        string `json:"path,omitempty" dynamodbav:"path"`
	Image       string `json:"image" dynamodbav:"image"`
	Platform    string `json:"platform" dynamodbav:"platform"`
	Description string `json:"description" dynamodbav:"description"`
}

type GameRepo interface {
	ListGames(context.Context) ([]Game, error)
}
