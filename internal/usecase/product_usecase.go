package usecase

import "go-vue/internal/domain"

type ProductUsecase struct {
	repo domain.ProductRepository
}

func NewProductUsecase(repo domain.ProductRepository) *ProductUsecase {
	return &ProductUsecase{repo: repo}
}

func (p *ProductUsecase) CreateProduct(product *domain.Product) error {
	return p.repo.CreateProduct(product)
}

func (p *ProductUsecase) GetProductByID(id int64) (*domain.Product, error) {
	return p.repo.GetProductByID(id)
}

func (p *ProductUsecase) GetAllProducts() ([]domain.Product, error) {
	return p.repo.GetAllProducts()
}

func (p *ProductUsecase) UpdateProduct(product *domain.Product) error {
	return p.repo.UpdateProduct(product)
}

func (p *ProductUsecase) DeleteProduct(id int64) error {
	return p.repo.DeleteProduct(id)
}
