package user_entity

import (
	"context"

	internal_error "github.com/garciawell/go-challenge-auction/internal/internal_erro"
)

type User struct {
	Id   string
	Name string
}

type UserRepositoryInterface interface {
	FindUserById(ctx context.Context, userId string) (*User, *internal_error.InternalError)
}
