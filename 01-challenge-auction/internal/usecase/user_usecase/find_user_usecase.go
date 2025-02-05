package user_usecase

import (
	"context"

	"github.com/garciawell/go-challenge-auction/internal/entity/user_entity"
	internal_error "github.com/garciawell/go-challenge-auction/internal/internal_erro"
)

func NewUserUseCase(userRepository user_entity.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}

type UserUseCaseInterface interface {
	FindUserById(ctx context.Context, userId string) (*UserOutputDTO, *internal_error.InternalError)
}

type UserUseCase struct {
	UserRepository user_entity.UserRepositoryInterface
}

type UserOutputDTO struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (u *UserUseCase) FindUserById(ctx context.Context, userId string) (*UserOutputDTO, *internal_error.InternalError) {
	userEntity, err := u.UserRepository.FindUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &UserOutputDTO{
		Id:   userEntity.Id,
		Name: userEntity.Name,
	}, nil
}
