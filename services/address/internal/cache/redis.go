package cache

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type RedisCache struct {
	cli    *redis.Client
	ttlCEP time.Duration
}

func NewRedisFromEnv() (*RedisCache, error) {
	addr := getenv("REDIS_ADDR", "localhost:6379")
	pwd := os.Getenv("REDIS_PASSWORD")
	db, _ := strconv.Atoi(getenv("REDIS_DB", "0"))
	ttl, _ := strconv.Atoi(getenv("REDIS_TTL_CEP_SECONDS", "86400"))

	cli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := cli.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis ping: %w", err)
	}

	return &RedisCache{
		cli:    cli,
		ttlCEP: time.Duration(ttl) * time.Second,
	}, nil
}

func (r *RedisCache) GetCEP(ctx context.Context, cep string) (string, bool, error) {
	val, err := r.cli.Get(ctx, keyCEP(cep)).Result()
	if err == redis.Nil {
		return "", false, nil
	}

	if err != nil {
		return "", false, err
	}
	return val, true, nil
}

func (r *RedisCache) SetCEP(ctx context.Context, cep string, json string) error {
	return r.cli.Set(ctx, keyCEP(cep), json, r.ttlCEP).Err()
}

func keyCEP(cep string) string { return "cep: " + cep }

func getenv(k, fallback string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return fallback
}
