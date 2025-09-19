package usecase

import "go-vue/internal/domain"

// UserUsercase implements the business logic for users.

type UserUsecase struct {
	repo domain.UserRepository
}

func NewUserUsercase(repo domain.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

// 💡 Penjelasan:
// Ini adalah lapisan use case/service.
// Memiliki metode yang sesuai dengan UserRepository interface.
// Meneruskan panggilan ke repository.
// Bisa ditambahkan logika bisnis tambahan di sini jika diperlukan.

func (u *UserUsecase) CreateUser(user *domain.User) error {
	return u.repo.CreateUser(user)
}

func (u *UserUsecase) GetUserByID(id int64) (*domain.User, error) {
	return u.repo.GetUserByID(id)
}

func (u *UserUsecase) GetAllUsers() ([]domain.User, error) {
	return u.repo.GetAllUsers()
}

func (u *UserUsecase) UpdateUser(user *domain.User) error {
	return u.repo.UpdateUser(user)
}

func (u *UserUsecase) DeleteUser(id int64) error {
	return u.repo.DeleteUser(id)
}

// 💡 Penjelasan:

// Usecase adalah tempat logika bisnis — misal: validasi, aturan bisnis, dll.
// Saat ini masih sederhana — langsung panggil repository.
// Nanti bisa ditambahi: cek duplikat, validasi umur, dll.
