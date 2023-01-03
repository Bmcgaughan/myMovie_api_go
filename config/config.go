package config

import (
	"os"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type APIConfig struct {
	MongoClient   *mongo.Client
	RedisClient   *redis.Client
	Port          string `mapstructure:"PORT"`
	MongoDBURI    string `mapstructure:"MONGODB_URI"`
	RedisURI      string `mapstructure:"REDIS_URI"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	APIKey        string `mapstructure:"API_KEY"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	ImageBase     string `mapstructure:"IMAGE_BASE"`
	AllowedOrigin string `mapstructure:"ALLOWED_ORIGINS"`
}

var MainConfig APIConfig

func LoadConfig() {

	MainConfig = APIConfig{
		Port:          os.Getenv("PORT"),
		MongoDBURI:    os.Getenv("MONGODB_URI"),
		RedisURI:      os.Getenv("REDIS_URI"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		APIKey:        os.Getenv("API_KEY"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
		ImageBase:     os.Getenv("IMAGE_BASE"),
		AllowedOrigin: os.Getenv("ALLOWED_ORIGINS"),
	}
}
