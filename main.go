package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sidz111/middlewade-practice1/middlewares"
)

func main() {
	r := gin.Default()

	r.Use(middlewares.MiddlewareName())

	r.GET("hello", middlewares.MiddlewareName(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})

	})
	r.Run(":8080")
}
