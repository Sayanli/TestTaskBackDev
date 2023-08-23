package repository

import (
	"context"

	"github.com/Sayanli/TestTaskBackDev/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Collection
}

func NewUserMongo(db *mongo.Database) *UserRepository {
	return &UserRepository{
		db: db.Collection("users"),
	}
}

func (r *UserRepository) Create(ctx context.Context, user domain.User) error {
	_, err := r.db.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) RefreshToken(ctx context.Context, user domain.User) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"guid": user.Guid}, bson.M{"$set": bson.M{"refresh_token": user.RefreshToken}})
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetByGuid(ctx context.Context, guid string) (domain.User, error) {
	var user domain.User
	err := r.db.FindOne(ctx, bson.M{"guid": guid}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserRepository) CheckDublicateUser(ctx context.Context, guid string) (bool, error) {
	count, err := r.db.CountDocuments(ctx, bson.M{"guid": guid})
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
