package handlers

import (
	"api_go/auth"
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

// GetAllTV is a handler for the /movies endpoint
func GetTV(c *gin.Context) {
	_, err := auth.ValidateJWT(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	shows, err := h.GetAllTV(db.Client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, shows)
}

// GetMovieByTitle is a handler for the /movies/:title endpoint
func GetTVByTitle(c *gin.Context) {
	_, err := auth.ValidateJWT(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	title := c.Param("title")

	show, err := h.GetTVByTitle(db.Client, title)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Movie not found",
		})
		return
	}
	c.JSON(http.StatusOK, show)
}

// AddToFavorites is a handler for the /users/:username/favorites/:movieID endpoint
func AddToFavorites(c *gin.Context) {
	username := c.Param("username")
	tvID := c.Param("movieID")

	authedUser, err := auth.ValidateJWT(c)
	if err != nil || authedUser != username {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	user, err := h.AddFavorite(db.Client, username, tvID)
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
	tvID := c.Param("movieID")

	authedUser, err := auth.ValidateJWT(c)
	if err != nil || authedUser != username {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	user, err := h.RemoveFavorite(db.Client, username, tvID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error removing movie from favorites",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetPopularTV is a handler for the /movies/popular endpoint
func GetPopularTV(c *gin.Context) {
	_, err := auth.ValidateJWT(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	shows, err := h.GetPopularTV(db.Client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, shows)
}

// GetTrendingTV is a handler for the /movies/trending endpoint
func GetTrendingTV(c *gin.Context) {
	_, err := auth.ValidateJWT(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	shows, err := h.GetTrendingTV(db.Client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, shows)
}

// GetRecommendedTV is a handler for the /movies/recommended/:id endpoint
func GetRecommendedTV(c *gin.Context) {
	_, err := auth.ValidateJWT(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	id := c.Param("id")

	shows, err := h.GetRecommendedTV(db.Client, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, shows)
}

// SearchTV is a handler for the /search/:query endpoint
func SearchTV(c *gin.Context) {
	_, err := auth.ValidateJWT(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	query := c.Param("query")

	shows, err := h.SearchTV(db.Client, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, shows)
}
