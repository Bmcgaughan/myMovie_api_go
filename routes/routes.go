package routes

import (
	"api_go/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/healthcheck", handlers.HealthCheck)
	r.GET("/movies", handlers.GetMovies)
	r.GET("/movies/:title", handlers.GetMovieByTitle)

	//add to useres favorites
	r.POST("/users/:username/favorites/:movieID", handlers.AddToFavorites)

	//remove from users favorites
	r.DELETE("/users/:username/favorites/:movieID", handlers.RemoveFromFavorites)

}
