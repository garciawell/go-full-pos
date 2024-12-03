package usecase

import (
	"github.com/garciawell/go-full-pos/00-challenge-clean-arch/internal/entity"
)

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: orderRepository,
	}
}

func (u *ListOrderUseCase) Execute() ([]entity.Order, error) {
	return u.OrderRepository.FindAll()
}
