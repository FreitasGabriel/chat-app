package service

import (
	"net/http"
	"strings"
	"time"

	"github.com/FreitasGabriel/chat-app/config/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (u *userService) CreateJWTToken(email, password string, JWTExpiresIn int, JWTSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    email,
		"password": password,
		"exp":      time.Now().Add(time.Second * time.Duration(JWTExpiresIn)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(JWTSecret))
	if err != nil {
		logger.Error("error to create a signed jwt", err)
		return "", err
	}

	return tokenString, nil
}

func ValidateJWTToken(c *gin.Context) {
	JWTSecret := c.Value("jwt_secret").(string)
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		logger.Error("token not found", nil)
		c.JSON(http.StatusUnauthorized, "authorization token not provided")
		c.Abort()
		return
	}

	token, err := jwt.Parse(RemoveBearerPrefix(tokenString), func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})

	if err != nil {
		logger.Error("error to validate jwt", err)
		c.JSON(http.StatusUnauthorized, err)
		c.Abort()
		return
	}

	if !token.Valid {
		logger.Error("invalid jwt token", err)
		c.JSON(http.StatusUnauthorized, err)
		c.Abort()
		return
	}

	logger.Info("user authenticated")
}

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}
	return token
}
