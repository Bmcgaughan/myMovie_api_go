package routes

import (
	"api_go/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/healthcheck", handlers.HealthCheck)

}
