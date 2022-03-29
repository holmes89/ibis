package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/holmes89/ibis/internal"
	v1 "github.com/holmes89/ibis/internal/handlers/rest/v1"
	"github.com/rs/zerolog/log"
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
	games, err := s.repo.ListGames(ctx)
	if err != nil {
		log.Error().Err(err).Msg("unable to find games")
		return resp, errors.New("unable to find games")
	}
	resp.Code = http.StatusOK
	resp.Body = games
	return resp, err
}

func (s *GameServicer) ListUserGameInstances(ctx context.Context, id string) (resp v1.ImplResponse, err error) {
	return resp, err
}
