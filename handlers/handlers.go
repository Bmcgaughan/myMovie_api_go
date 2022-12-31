package handlers

import (
	"api_go/auth"
	config "api_go/config"
	"api_go/db"
	h "api_go/helpers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	popularCacheKey  = "popular"
	trendingCacheKey = "trending"
)

// healthcheck endpoint
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

// /tv endpoint
func GetTV(c *gin.Context) {
	_, err := auth.ValidateJWT(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	shows, err := h.GetAllTV(config.MainConfig.MongoClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, shows)
}

// /tv/:title endpoint
func GetTVByTitle(c *gin.Context) {
	_, err := auth.ValidateJWT(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	title := c.Param("title")

	show, err := h.GetTVByTitle(config.MainConfig.MongoClient, title)
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

	user, err := h.AddFavorite(config.MainConfig.MongoClient, username, tvID)
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

	user, err := h.RemoveFavorite(config.MainConfig.MongoClient, username, tvID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error removing movie from favorites",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

// /tv/popular endpoint
func GetPopularTV(c *gin.Context) {
	_, err := auth.ValidateJWT(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	//check redis cache for popular movies and return if exists
	popularShows, err := db.GetCache(popularCacheKey)
	if err == nil {
		log.Println("Cache hit for popular movies")
		c.JSON(http.StatusOK, popularShows)
		return
	}

	shows, err := h.GetPopularTV(config.MainConfig.MongoClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	//set popular movies in redis cache
	err = db.SetCache(popularCacheKey, shows)
	if err != nil {
		log.Println("Error setting popular movies in cache")
	}

	c.JSON(http.StatusOK, shows)
}

// /tv/trending endpoint
func GetTrendingTV(c *gin.Context) {
	_, err := auth.ValidateJWT(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	//check redis cache for trending movies and return if exists
	trendingShows, err := db.GetCache(trendingCacheKey)
	if err == nil {
		log.Println("Cache hit for trending movies")
		c.JSON(http.StatusOK, trendingShows)
		return
	}

	shows, err := h.GetTrendingTV(config.MainConfig.MongoClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	//set trending movies in redis cache
	err = db.SetCache(trendingCacheKey, shows)
	if err != nil {
		log.Println("Error setting trending movies in cache")
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

	shows, err := h.GetRecommendedTV(config.MainConfig.MongoClient, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, shows)
}

// /tv/foryou endpoint
func GetTVForYou(c *gin.Context) {
	username, err := auth.ValidateJWT(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	shows, err := h.GetMostRecommendedTV(config.MainConfig.MongoClient, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, shows)
}

func GetUser(c *gin.Context) {
	username := c.Param("username")

	authedUser, err := auth.ValidateJWT(c)
	if err != nil || authedUser != username {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	user, err := h.GetUserDetails(config.MainConfig.MongoClient, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, user)
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

	shows, err := h.SearchTV(config.MainConfig.MongoClient, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, shows)
}
