package product

type ProductUseCase struct {
	ProductRepository ProductRepositoryInterface
}

func NewProductUseCase(repository ProductRepositoryInterface) *ProductUseCase {
	return &ProductUseCase{repository}
}

func (u *ProductUseCase) GetProduct(id int) (*Product, error) {
	return u.ProductRepository.GetProduct(id)
}
