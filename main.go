package main

import (
	config "api_go/config"
	"api_go/db"
	"api_go/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadConfig()

	port := config.MainConfig.Port
	r := gin.Default()
	r.Use(gin.Logger())
	r.SetTrustedProxies(nil)
	routes.SetupRoutes(r)

	mongo := db.ConnectDB()
	redisClient := db.ConnectRedis()

	config.MainConfig.MongoClient = mongo
	config.MainConfig.RedisClient = redisClient

	r.Run(":" + port)
}
