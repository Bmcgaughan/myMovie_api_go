package routes

import (
	"api_go/auth"
	"api_go/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/healthcheck", handlers.HealthCheck)
	r.GET("/tv", handlers.GetTV)
	r.GET("/tv/:title", handlers.GetTVByTitle)
	r.GET("/tv/popular", handlers.GetPopularTV)
	r.GET("tv/trending", handlers.GetTrendingTV)
	r.GET("tv/recommended/:id", handlers.GetRecommendedTV)
	r.GET("tv/foryou", handlers.GetTVForYou)
	r.GET("search/:query", handlers.SearchTV)

	r.GET("users/:username", handlers.GetUser)

	r.POST("/users", auth.CreateUser)
	r.POST("login", auth.LoginUser)

	//add to useres favorites
	r.POST("/users/:username/favorites/:movieID", handlers.AddToFavorites)

	//remove from users favorites
	r.DELETE("/users/:username/favorites/:movieID", handlers.RemoveFromFavorites)

}
