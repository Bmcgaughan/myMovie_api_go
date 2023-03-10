package auth

import (
	config "api_go/config"
	h "api_go/helpers"
	"api_go/models"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type SiteUser struct {
	Username string `json:"username" binding:"required,min=5,alphanum"`
	Password string `json:"password" binding:"required,min=5"`
}

func (su *SiteUser) hashPassword() error {
	cost, err := strconv.Atoi(os.Getenv("COST"))
	if err != nil {
		return err
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(su.Password), cost)
	if err != nil {
		return err
	}
	su.Password = string(bytes)
	return nil
}

func CreateUser(c *gin.Context) {
	var user SiteUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = user.hashPassword()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
	}

	userModel := models.User{UserName: user.Username, Password: user.Password}

	userCreated, err := h.CreateUser(config.MainConfig.MongoClient, userModel)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "User already exists",
		})
		return
	}
	c.JSON(http.StatusCreated, userCreated)
}

func LoginUser(c *gin.Context) {
	var user SiteUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userModel, err := h.GetUser(config.MainConfig.MongoClient, user.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	// return jwt token
	token, err := generateToken(userModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	loginTimeErr := h.LastLogin(config.MainConfig.MongoClient, userModel.UserName)
	if loginTimeErr != nil {
		log.Printf("Error adding Login Time to User %s at %s", userModel.UserName, time.Now())
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  userModel,
		"token": token,
	})
}

func generateToken(user models.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["subject"] = user.UserName
	claims["exp"] = time.Now().Add(time.Hour * 24 * 15).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateJWT(c *gin.Context) (string, error) {
	tokenString := c.Request.Header.Get("Authorization")
	bearerToken := strings.TrimPrefix(tokenString, "Bearer ")
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("bad Signing Method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	return token.Claims.(jwt.MapClaims)["subject"].(string), nil
}

