package repository

import (
	"context"

	"github.com/Sayanli/TestTaskBackDev/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository struct {
	db *mongo.Collection
}

func NewAuthMongo(db *mongo.Database) *AuthRepository {
	return &AuthRepository{
		db: db.Collection("users"),
	}
}

func (r *AuthRepository) Create(ctx context.Context, user entity.User) error {
	_, err := r.db.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthRepository) RefreshToken(ctx context.Context, user entity.User) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"guid": user.Guid}, bson.M{"$set": bson.M{"refresh_token": user.RefreshToken}})
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthRepository) GetByGuid(ctx context.Context, guid string) (entity.User, error) {
	var user entity.User
	err := r.db.FindOne(ctx, bson.M{"guid": guid}).Decode(&user)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *AuthRepository) IsUserExists(ctx context.Context, guid string) (bool, error) {
	count, err := r.db.CountDocuments(ctx, bson.M{"guid": guid})
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
