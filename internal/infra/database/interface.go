package database

import "github.com/garciawell/go-full-pos/apis/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(emailId string) (*entity.User, error)
}
