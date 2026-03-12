package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sidz111/middlewade-practice1/middlewares"
)

func main() {
	r := gin.Default()

	r.Use(middlewares.MiddlewareName())

	r.GET("hello", middlewares.MiddlewareName(), func(c *gin.Context) {
		token := c.GetHeader("authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "Unauthorized",
			})
			return
		}
		c.Next()
		c.JSON(200, gin.H{
			"message": "hello",
			"token":   token,
		})

	})
	r.Run(":8080")
}
