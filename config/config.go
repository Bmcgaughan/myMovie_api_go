package config

import (
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	viper "github.com/spf13/viper"
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
	Cost          int    `mapstructure:"COST"`
}

var MainConfig APIConfig

func LoadConfig() {
	// load .env file from root if exists
	err := godotenv.Load("settings.env")
	if err != nil {
		log.Println("No .env file found")
	}

	v := viper.New()
	v.BindEnv("PORT")
	v.BindEnv("MONGODB_URI")
	v.BindEnv("REDIS_URI")
	v.BindEnv("REDIS_PASSWORD")
	v.BindEnv("API_KEY")
	v.BindEnv("JWT_SECRET")
	v.BindEnv("IMAGE_BASE")
	v.BindEnv("COST")

	//load viper into Config struct
	err = v.Unmarshal(&MainConfig)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

}
