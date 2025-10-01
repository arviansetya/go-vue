package domain

type Product struct {
	ID          int64   `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name" gorm:"not null;unique"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Category    string  `json:"category"`
}

// Product Repository defines the methods for product data access.
// ProductRepository adalah interface — memisahkan logika bisnis dari infrastruktur.
type ProductRepository interface {
	CreateProduct(product *Product) error
	GetProductByID(id int64) (*Product, error)
	GetAllProducts() ([]Product, error)
	UpdateProduct(product *Product) error
	DeleteProduct(id int64) error
	// SearchProductsByName(name string) ([]Product, error)
}
