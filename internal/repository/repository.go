package repository

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

func InitDatabase(databaseURL string) *sql.DB {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("Failed to open a database connection: %v", err)
	}
	return db
}

func InitRedis(redisURL string) *redis.Client {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("Failed to parse Redis URL: %v", err)
	}
	client := redis.NewClient(opt)
	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	return client
}
