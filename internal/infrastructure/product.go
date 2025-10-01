package infrastructure

import (
	"go-vue/internal/domain"

	"gorm.io/gorm"
)

// ProductRepository implements domain.ProductRepository using GORM.
type ProductRepository struct {
	DB *gorm.DB
}

// Konstruktor untuk ProductRepository
func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

// CreateProduct creates a new product in the database.
func (r *ProductRepository) CreateProduct(product *domain.Product) error {
	model := &Product{}
	model.FromDomain(product)
	return r.DB.Create(model).Error
}

// GetProductByID retrieves a product by ID from the database.
func (r *ProductRepository) GetProductByID(id int64) (*domain.Product, error) {
	var model Product
	if err := r.DB.First(&model, id).Error; err != nil {
		return nil, err
	}
	return model.ToDomain(), nil
}

// GetAllProducts retrieves all products from the database.
func (r *ProductRepository) GetAllProducts() ([]domain.Product, error) {
	var models []Product
	if err := r.DB.Find(&models).Error; err != nil {
		return nil, err
	}
	var products []domain.Product
	for _, model := range models {
		products = append(products, *model.ToDomain())
	}
	return products, nil
}

// UpdateProduct updates an existing product in the database.
func (r *ProductRepository) UpdateProduct(product *domain.Product) error {
	model := &Product{}
	model.FromDomain(product)
	return r.DB.Save(model).Error
}

// DeleteProduct deletes a product by ID from the database.
func (r *ProductRepository) DeleteProduct(id int64) error {
	return r.DB.Delete(&Product{}, id).Error
}
