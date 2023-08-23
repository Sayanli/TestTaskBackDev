package service

import (
	"context"

	"github.com/Sayanli/TestTaskBackDev/internal/repository"
)

type User interface {
	CreateUser(ctx context.Context, guid string) (Tokens, error)
	RefreshToken(ctx context.Context, guid string, refreshToken string) (Tokens, error)
}

type Service struct {
	User
}

func NewService(r *repository.Repository, secret string) *Service {
	return &Service{
		User: NewUserService(r, secret),
	}
}
