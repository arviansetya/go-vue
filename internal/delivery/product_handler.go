package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"go-vue/internal/domain"
	"go-vue/internal/infrastructure"
	"go-vue/internal/usecase"
)

type ProductHandler struct {
	usecase *usecase.ProductUsecase
}

func NewProductHandler(db *gorm.DB) *ProductHandler {
	repo := infrastructure.NewProductRepository(db)
	uc := usecase.NewProductUsecase(repo)
	return &ProductHandler{usecase: uc}
}

// Register Product routes
func (h *ProductHandler) RegisterRoutes(e *echo.Echo) {
	e.POST("/products", h.CreateProduct)
	e.GET("/products/:id", h.GetProductByID)
	e.GET("/products", h.GetAllProducts)
	e.PUT("/products/:id", h.UpdateProduct)
	e.DELETE("/products/:id", h.DeleteProduct)
}

// CreateProduct handles POST /products
func (h *ProductHandler) CreateProduct(c echo.Context) error {
	var product domain.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	if err := h.usecase.CreateProduct(&product); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create product"})
	}

	// 🧹 Hapus cache
	infrastructure.RedisClient.Del(infrastructure.Ctx, "products:list")

	return c.JSON(http.StatusCreated, product)
}

// GetProductByID handles GET /products/:id
func (h *ProductHandler) GetProductByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}
	product, err := h.usecase.GetProductByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}
	return c.JSON(http.StatusOK, product)
}

// GetAllProducts handles GET /products
func (h *ProductHandler) GetAllProducts(c echo.Context) error {
	// 🔍 Coba ambil dari Redis dulu
	cacheKey := "products:list"
	val, err := infrastructure.RedisClient.Get(infrastructure.Ctx, cacheKey).Result()
	if err == nil {
		// ✅ Cache hit
		var products []domain.Product
		json.Unmarshal([]byte(val), &products)
		return c.JSON(http.StatusOK, products)
	}

	// ❌ Cache miss → ambil dari DB
	products, err := h.usecase.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve products"})
	}

	// 💾 Simpan ke Redis (expire 5 menit)
	data, _ := json.Marshal(products)
	infrastructure.RedisClient.Set(
		infrastructure.Ctx,
		cacheKey,
		data,
		5*time.Minute,
	)

	return c.JSON(http.StatusOK, products)
}

// UpdateProduct handles PUT /products/:id
func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}
	var product domain.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	product.ID = id
	if err := h.usecase.UpdateProduct(&product); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update product"})
	}

	// 🧹 Hapus cache
	infrastructure.RedisClient.Del(infrastructure.Ctx, "products:list")

	return c.JSON(http.StatusOK, product)
}

// DeleteProduct handles DELETE /products/:id
func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}
	if err := h.usecase.DeleteProduct(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete product"})
	}

	// 🧹 Hapus cache
	infrastructure.RedisClient.Del(infrastructure.Ctx, "products:list")

	return c.NoContent(http.StatusNoContent)
}
