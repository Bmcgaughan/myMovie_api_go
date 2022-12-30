package config

import (
	"log"

	"github.com/go-redis/redis/v8"
	viper "github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type APIConfig struct {
	MongoClient *mongo.Client
	RedisClient *redis.Client
	Config      Config
}

type Config struct {
	Port          string
	MongoURI      string
	RedisURI      string
	RedisPassword string
	ImageBase     string
	Cost          int
	ApiKey        string
	JWTSecret     string
}

var MainConfig APIConfig

func LoadConfig() {
	// load config from env variables and return a Config struct
	viper.AutomaticEnv()

	viper.SetDefault("port", "8080")
	viper.SetDefault("mongo_uri", "")
	viper.SetDefault("redis_uri", "")
	viper.SetDefault("redis_password", "")
	viper.SetDefault("image_base", "")
	viper.SetDefault("cost", 10)

	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

}
