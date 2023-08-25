package repository

import (
	"context"

	"github.com/Sayanli/TestTaskBackDev/internal/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type Auth interface {
	Create(ctx context.Context, user entity.User) error
	RefreshToken(ctx context.Context, user entity.User) error
	GetByGuid(ctx context.Context, guid string) (entity.User, error)
	IsUserExists(ctx context.Context, guid string) (bool, error)
}

type Repository struct {
	Auth
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Auth: NewAuthMongo(db),
	}
}
