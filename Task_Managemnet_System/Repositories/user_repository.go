package repositories

import (
	"context"
	"errors"

	"github.com/BemnetMussa/Backend_A2SV/Task_Managemnet_System/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	CountUsers(ctx context.Context) (int64, error)
	CreateUser(ctx context.Context, user domain.User) error
	PromoteUserByEmail(ctx context.Context, email string) error
}

type MongoUserRepository struct {
	Collection *mongo.Collection
}

func NewMongoUserRepository(collection *mongo.Collection) *MongoUserRepository {
	return &MongoUserRepository{Collection: collection}
}

func (r *MongoUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	err := r.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MongoUserRepository) CountUsers(ctx context.Context) (int64, error) {
	count, err := r.Collection.CountDocuments(ctx, bson.M{})
	return count, err
}

func (r *MongoUserRepository) CreateUser(ctx context.Context, user domain.User) error {
	_, err := r.Collection.InsertOne(ctx, user)
	return err
}

func (r *MongoUserRepository) PromoteUserByEmail(ctx context.Context, email string) error {
	filter := bson.M{"email": email}
	update := bson.M{"$set": bson.M{"role": "admin"}}

	result, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New("failed to promote user")
	}
	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}
