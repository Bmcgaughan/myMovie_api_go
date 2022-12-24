package handlers

import (
	"api_go/db"
	h "api_go/helpers"

	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck is a handler for the healthcheck endpoint
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

// GetAllMovies is a handler for the /movies endpoint
func GetMovies(c *gin.Context) {
	movies, err := h.GetAllMovies(db.Client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, movies)
}

// GetMovieByTitle is a handler for the /movies/:title endpoint
func GetMovieByTitle(c *gin.Context) {
	title := c.Param("title")

	movie, err := h.GetMovieByTitle(db.Client, title)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Movie not found",
		})
		return
	}
	c.JSON(http.StatusOK, movie)
}

// AddToFavorites is a handler for the /users/:username/favorites/:movieID endpoint
func AddToFavorites(c *gin.Context) {
	username := c.Param("username")
	movieID := c.Param("movieID")

	user, err := h.AddFavoriteMovie(db.Client, username, movieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error adding movie to favorites",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

// RemoveFromFavorites is a handler for the /users/:username/favorites/:movieID endpoint
func RemoveFromFavorites(c *gin.Context) {
	username := c.Param("username")
	movieID := c.Param("movieID")

	user, err := h.RemoveFavoriteMovie(db.Client, username, movieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error removing movie from favorites",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}
