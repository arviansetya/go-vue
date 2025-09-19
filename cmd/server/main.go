package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-vue/internal/delivery"
	"go-vue/internal/infrastructure"
)

func main() {
	// Load environment variables
	godotenv.Load()

	// Connect to PostgreSQL
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	// Auto-migrate schema (buat tabel jika belum ada)
	err = db.AutoMigrate(&infrastructure.User{})
	if err != nil {
		log.Fatal("❌ Failed to migrate database:", err)
	}

	log.Println("✅ Database connected and migrated successfully")

	// Setup Echo server
	e := echo.New()

	// Register user routes
	userHandler := delivery.NewUserHandler(db)
	userHandler.RegisterRoutes(e)

	// Start server
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server starting on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}
