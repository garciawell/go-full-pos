package product

type ProductUseCase struct {
	ProductRepository *ProductRepository
}

func NewProductUseCase(repository *ProductRepository) *ProductUseCase {
	return &ProductUseCase{repository}
}

func (u *ProductUseCase) GetProduct(id int) (*Product, error) {
	return u.ProductRepository.GetProduct(id)
}
