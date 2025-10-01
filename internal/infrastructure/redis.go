// internal/infrastructure/redis.go
package infrastructure

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	Ctx         = context.Background()
)

// InitRedisClient menginisialisasi koneksi ke Redis
func InitRedisClient() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "192.168.1.113:6379", // sesuai Redis Windows kamu
		Password: "",                   // default tanpa password
		DB:       0,                    // gunakan database 0
	})

	// Test koneksi
	_, err := RedisClient.Ping(Ctx).Result()
	return err
}
