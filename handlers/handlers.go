package handlers

import (
	"api_go/auth"
	"api_go/db"
	h "api_go/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// healthcheck endpoint
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

// /movies endpoint
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

// /movies/:title endpoint
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

// /users/:username/favorites/:movieID endpoint
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

// /users/:username/favorites/:movieID endpoint
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

// /movies/popular endpoint
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

// /movies/trending endpoint
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

// /tv/recommended/:id endpoint
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

// /search/:query endpoint
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
