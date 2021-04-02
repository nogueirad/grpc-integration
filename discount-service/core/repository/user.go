package repository

import (
	"discount-service/core/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUserById(userID string) (domain.User, error)
}

type userRepository struct {
	usersCollection *mongo.Collection
}

func NewUserRepository(usersCollection *mongo.Collection ) UserRepository {
	return &userRepository{usersCollection}
}

func (c *userRepository) GetUserById(userID string) (domain.User, error) {
	user := domain.User{
		FirstName:"zé",
	}

	return user, nil
}
