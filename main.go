package main

import (
	config "api_go/config"
	"api_go/db"
	"api_go/routes"
	"log"
	"os"
	"strings"
	"time"

	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	enviro := os.Getenv("ENVIRO")

	config.LoadConfig()

	port := config.MainConfig.Port
	r := gin.Default()
	if enviro == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
	log.Println(allowedOrigins)
	log.Println(config.MainConfig)
	// set cors policy
	r.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(gin.Logger())

	routes.SetupRoutes(r)

	mongo := db.ConnectDB()
	redisClient := db.ConnectRedis()

	config.MainConfig.MongoClient = mongo
	config.MainConfig.RedisClient = redisClient

	r.Run(":" + port)
}
