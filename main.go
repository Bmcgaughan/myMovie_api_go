package main

import (
	"api_go/db"
	"api_go/routes"
	"context"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	r := gin.Default()
	r.SetTrustedProxies(nil)
	routes.SetupRoutes(r)

	// Connect to the database
	client := db.ConnectDB()
	defer client.Disconnect(context.Background())
	r.Run(":" + port)
}
