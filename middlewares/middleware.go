package middlewares

import (
	"fmt"

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
