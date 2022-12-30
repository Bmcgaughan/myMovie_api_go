package db

import (
	config "api_go/config"
	"api_go/models"
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	redis "github.com/go-redis/redis/v8"
)

const (
	redisTTL = 60 * 60 * 24 * 2 // 2 days
)

func ConnectRedis() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"), // use default Addr
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Panic(err)
	}

	log.Println("Connected to Redis!")
	return redisClient
}

func GetCache(key string) ([]*models.Movie, error) {
	val, err := config.MainConfig.RedisClient.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}

	var shows []*models.Movie
	jsonErr := json.Unmarshal([]byte(val), &shows)
	if jsonErr != nil {
		log.Println("Error unmarshalling popular movies")
		return nil, jsonErr
	}

	return shows, nil
}

func SetCache(key string, shows *[]models.Movie) error {
	ttl := redisTTL

	value, marshalErr := json.Marshal(shows)
	if marshalErr != nil {
		log.Println("Error marshalling popular movies")
		return marshalErr
	}

	err := config.MainConfig.RedisClient.Set(context.Background(), key, value, time.Duration(ttl)*time.Second)
	if err != nil {
		log.Println("Error setting popular movies")
		return err.Err()
	}

	return nil
}
