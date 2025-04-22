package handler

import (
	"net/http"

	rl "github.com/FreitasGabriel/chat-app/config/redis"
	"github.com/gin-gonic/gin"
)

func RateLimitMiddleware(rl *rl.RateLimit, next gin.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		clientIP := c.RemoteIP()
		if !rl.Allow(clientIP) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			return
		}
	})
}
