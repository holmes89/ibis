package service

import (
	"context"

	"github.com/holmes89/ibis/internal"
	v1 "github.com/holmes89/ibis/internal/handlers/rest/v1"
)

var (
	_ v1.UserApiServicer = (*UserServicer)(nil)
)

type UserServicer struct {
	repo internal.UserRepo
}

func NewUserServicer(repo internal.UserRepo) *UserServicer {
	return &UserServicer{
		repo: repo,
	}
}

func (s *UserServicer) GetUserByName(ctx context.Context, id string) (resp v1.ImplResponse, err error) {
	return resp, err
}
func (s *UserServicer) ListUserFriends(ctx context.Context, id string) (resp v1.ImplResponse, err error) {
	return resp, err
}
func (s *UserServicer) ListUsers(ctx context.Context) (resp v1.ImplResponse, err error) {
	return resp, err
}
