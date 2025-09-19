package domain

// User represents a user in the system.
// Entity ini mencerminkan struktur tabel di database.
type User struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"not null;unique"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// User Repository defines the methods for user data access.
//UserRepository adalah interface — memisahkan logika bisnis dari infrastruktur.
type UserRepository interface {
	CreateUser(user *User) error
	GetUserByID(id int64) (*User, error)
	GetAllUsers() ([]User, error)
	UpdateUser(user *User) error
	DeleteUser(id int64) error
}
