package infrastructure

import (
	"fmt"
	"github.com/BambooTuna/letustalk/backend/config"
	"github.com/go-redis/redis"
)

func RedisConnect(db int) *redis.Client {
	redisAddr := fmt.Sprintf("%s:%s",
		config.FetchEnvValue("REDIS_HOST", "127.0.0.1"),
		config.FetchEnvValue("REDIS_PORT", "6379"),
	)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       db,
	})
	return redisClient
}
