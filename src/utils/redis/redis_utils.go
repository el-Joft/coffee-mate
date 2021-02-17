package redis

import (
	"context"
	"errors"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/jpoles1/gopherbadger/logging"
)

// Database ->
type Database struct {
	Client *redis.Client
}

// Repository ->
func Repository() *Database {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	return NewRedisDB(redisHost, redisPort, redisPassword)
}

var (
	// ErrNil ->
	ErrNil = errors.New("no matching record found in redis database")
	// Ctx ->
	Ctx = context.TODO()
)

// NewRedisDB ->
func NewRedisDB(host, port, password string) *Database {
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})
	if err := client.Ping(Ctx).Err(); err != nil {
		logging.Error("Failed to connect to redis: %s", err)
	}

	return &Database{
		Client: client,
	}
}
