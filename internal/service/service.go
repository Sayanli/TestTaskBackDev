package service

import (
	"context"

	"github.com/Sayanli/TestTaskBackDev/internal/entity"
	"github.com/Sayanli/TestTaskBackDev/internal/repository"
)

type Auth interface {
	CreateUser(ctx context.Context, guid string) (entity.Token, error)
	RefreshToken(ctx context.Context, guid string, refreshToken string) (entity.Token, error)
}

type Service struct {
	Auth
}

func NewService(r *repository.Repository, secret string) *Service {
	return &Service{
		Auth: NewAuthService(r, secret),
	}
}
