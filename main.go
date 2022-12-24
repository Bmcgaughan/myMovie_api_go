package main

import (
	"api_go/db"
	"api_go/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	r := gin.Default()
	r.SetTrustedProxies(nil)
	routes.SetupRoutes(r)

	db.ConnectDB()

	r.Run(":" + port)
}
