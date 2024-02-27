package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MasterAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerValue := c.Request.Header.Get("Authorization")
		fmt.Println(headerValue)
		c.JSON(http.StatusOK, gin.H{"message": "这是自定义中间件"})
	}
}
