package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MiddlewareName() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Middleware Started")
		c.Next()
		fmt.Println("middleware Stoped")
	}
}
