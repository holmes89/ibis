package service

import (
	"context"

	"github.com/holmes89/ibis/internal"
	v1 "github.com/holmes89/ibis/internal/handlers/rest/v1"
)

var (
	_ v1.GameApiServicer = (*GameServicer)(nil)
)

type GameServicer struct {
	repo internal.GameRepo
}

func NewGameServicer(repo internal.GameRepo) *GameServicer {
	return &GameServicer{
		repo: repo,
	}
}

func (s *GameServicer) ListGames(ctx context.Context) (resp v1.ImplResponse, err error) {
	return resp, err
}

func (s *GameServicer) ListUserGameInstances(ctx context.Context, id string) (resp v1.ImplResponse, err error) {
	return resp, err
}
