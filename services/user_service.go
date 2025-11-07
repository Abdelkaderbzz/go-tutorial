package services

import (
	"context"
	"myapp/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	client *mongo.Client
	dbName string
}

func NewUserService(client *mongo.Client, dbName string) *UserService {
	return &UserService{
		client: client,
		dbName: dbName,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user models.User) (interface{}, error) {
	collection := s.client.Database(s.dbName).Collection("users")
	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}