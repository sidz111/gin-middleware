package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func MiddlewareName() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := uuid.New().String()
		c.Set("UserId", id)
		fmt.Println("Middleware Started")
		c.Next()
		fmt.Println("middleware Stoped")
	}
}

var requests = make(map[string]int)

func RateLimitter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		requests[ip]++
		if requests[ip] > 3 {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "too mane requests",
				"data":  requests,
			})
		}
	}
}
