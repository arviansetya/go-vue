package delivery

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"go-vue/internal/domain"
	"go-vue/internal/infrastructure"
	"go-vue/internal/usecase"
)

type UserHandler struct {
	usecase *usecase.UserUsecase
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	repo := infrastructure.NewUserRepository(db)
	uc := usecase.NewUserUsercase(repo)
	return &UserHandler{usecase: uc}
}

// 💡 Penjelasan:
// Ini adalah lapisan delivery/handler — menangani HTTP request.
// Memanggil usecase untuk logika bisnis.
// Bisa ditambahkan middleware, validasi input, dll di sini.
// Contoh endpoint handler:

// Register User routes
func (h *UserHandler) RegisterRoutes(e *echo.Echo) {
	e.POST("/users", h.CreateUser)
	e.GET("/users/:id", h.GetUserByID)
	e.GET("/users", h.GetAllUsers)
	e.PUT("/users/:id", h.UpdateUser)
	e.DELETE("/users/:id", h.DeleteUser)
}

// CreateUser handles POST /users
func (h *UserHandler) CreateUser(c echo.Context) error {
	var user domain.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	if err := h.usecase.CreateUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}
	return c.JSON(http.StatusCreated, user)
}

// GetUserByID handles GET /users/:id
func (h *UserHandler) GetUserByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}
	user, err := h.usecase.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}
	return c.JSON(http.StatusOK, user)
}

// GetAllUsers handles GET /users
func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.usecase.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve users"})
	}
	return c.JSON(http.StatusOK, users)
}

// UpdateUser handles PUT /users/:id
func (h *UserHandler) UpdateUser(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}
	var user domain.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	user.ID = id
	if err := h.usecase.UpdateUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user"})
	}
	return c.JSON(http.StatusOK, user)
}

// DeleteUser handles DELETE /users/:id
func (h *UserHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}
	if err := h.usecase.DeleteUser(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
	}
	return c.NoContent(http.StatusNoContent)
}

// 💡 Penjelasan:
// Handler ini menggunakan Echo framework untuk routing dan request handling.
// Setiap handler method mengonversi request ke domain.User, memanggil usecase, dan mengembalikan response.
// Bisa ditambahkan validasi input, middleware autentikasi, dll di sini sesuai kebutuhan.
