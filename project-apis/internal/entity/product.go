package entity

import (
	"errors"
	"time"

	"github.com/garciawell/go-full-pos/apis/pkg/entity"
)

var (
	ErrIDsRequired   = errors.New("id is required")
	ErrIDsInvalid    = errors.New("invalid ID")
	ErrNamesRequired = errors.New("name is required")
	ErrPriceRequired = errors.New("price is required")
	ErrPriceIndalid  = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		Name:      name,
		Price:     price,
		ID:        entity.NewID(),
		CreatedAt: time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrIDsInvalid

	}
	if p.Name == "" {
		return ErrNamesRequired
	}
	if p.Price == 0 {
		return ErrPriceRequired
	}
	if p.Price < 0 {
		return ErrPriceIndalid
	}
	return nil
}
