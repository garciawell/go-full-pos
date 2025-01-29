package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/garciawell/go-challenge-auction/config/logger"
	"github.com/garciawell/go-challenge-auction/internal/entity/user_entity"
	internal_error "github.com/garciawell/go-challenge-auction/internal/internal_erro"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserEntityMongo struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{Collection: db.Collection("users")}
}

func (r *UserRepository) FindUserById(ctx context.Context, id string) (*user_entity.User, *internal_error.InternalError) {

	var userEntityMongo UserEntityMongo
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&userEntityMongo)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(fmt.Sprintf("user not found with this id = %s", id), err)
			return nil, internal_error.NewNotFoundError(fmt.Sprintf("user not found with this id = %d", id))
		}

		logger.Error(fmt.Sprintf("Error trying to find user with this id = %s", id), err)
		return nil, internal_error.NewNotFoundError(fmt.Sprintf("Error trying to find user with this id  = %d", id))
	}

	return &user_entity.User{
		Id:   userEntityMongo.Id,
		Name: userEntityMongo.Name,
	}, nil
}
