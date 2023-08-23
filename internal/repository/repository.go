package repository

import (
	"context"

	"github.com/Sayanli/TestTaskBackDev/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type User interface {
	Create(ctx context.Context, user domain.User) error
	RefreshToken(ctx context.Context, user domain.User) error
	GetByGuid(ctx context.Context, guid string) (domain.User, error)
	CheckDublicateUser(ctx context.Context, guid string) (bool, error)
}

type Repository struct {
	User
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		User: NewUserMongo(db),
	}
}
