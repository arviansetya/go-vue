package infrastructure

import (
	"go-vue/internal/domain"

	"gorm.io/gorm"
)

// UserRepository implements domain.UserRepository using GORM.
type UserRepository struct {
	DB *gorm.DB
}

// Konstruktor untuk UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser creates a new user in the database.
func (r *UserRepository) CreateUser(user *domain.User) error {
	model := &User{}
	model.FromDomain(user)
	return r.DB.Create(model).Error
}

// GetUserByID retrieves a user by ID from the database.
func (r *UserRepository) GetUserByID(id int64) (*domain.User, error) {
	var model User
	if err := r.DB.First(&model, id).Error; err != nil {
		return nil, err
	}
	return model.ToDomain(), nil
}

// GetAllUsers retrieves all users from the database.
func (r *UserRepository) GetAllUsers() ([]domain.User, error) {
	var models []User
	if err := r.DB.Find(&models).Error; err != nil {
		return nil, err
	}
	var users []domain.User
	for _, model := range models {
		users = append(users, *model.ToDomain())
	}
	return users, nil
}

// UpdateUser updates an existing user in the database.
func (r *UserRepository) UpdateUser(user *domain.User) error {
	model := &User{}
	model.FromDomain(user)
	return r.DB.Save(model).Error
}

// DeleteUser deletes a user by ID from the database.
func (r *UserRepository) DeleteUser(id int64) error {
	return r.DB.Delete(&User{}, id).Error
}

// 💡 Penjelasan:

// Ini implementasi dari UserRepository interface.
// Menggunakan GORM untuk query ke PostgreSQL.
// Setiap method mengonversi antara domain.User dan infrastructure.User.
