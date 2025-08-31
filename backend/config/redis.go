package config

import (
	"CurrencyExchangeApp/global"
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func initRedis() {
	//获取redis客户端
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	//测试连通性
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis, got error: %v", err)
	}

	global.RedisDB = RedisClient
}
