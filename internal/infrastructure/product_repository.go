package infrastructure

import "go-vue/internal/domain"

type Product struct {
	ID          int64   `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name" gorm:"not null;unique"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Category    string  `json:"category"`
}

func (p *Product) ToDomain() *domain.Product {
	return &domain.Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
		Category:    p.Category,
	}
}

func (p *Product) FromDomain(product *domain.Product) {
	p.ID = product.ID
	p.Name = product.Name
	p.Description = product.Description
	p.Price = product.Price
	p.Stock = product.Stock
	p.Category = product.Category
}
