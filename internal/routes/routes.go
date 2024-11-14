package routes

import "github.com/gin-gonic/gin"

func Init(c *gin.RouterGroup) {
	c.GET("/heakth", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
