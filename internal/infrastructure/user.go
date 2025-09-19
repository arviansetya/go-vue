package infrastructure

import "go-vue/internal/domain"

type User struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"not null;unique"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// Konversi dari domain ke infrastruktur (untuk simpan ke DB)
func (u *User) FromDomain(user *domain.User) {
	u.ID = user.ID
	u.Username = user.Username
	u.Email = user.Email
	u.Password = user.Password
	u.Role = user.Role
}

// Konversi dari infrastruktur ke domain (untuk kembalikan ke service/use case)
func (u *User) ToDomain() *domain.User {
	return &domain.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		Role:     u.Role,
	}
}

//Penjelasan:

// Ini adalah model GORM — punya tag gorm:"..." untuk mapping ke tabel.
// Fungsi FromDomain() dan ToDomain() untuk konversi antara lapisan — agar domain tetap bersih.
