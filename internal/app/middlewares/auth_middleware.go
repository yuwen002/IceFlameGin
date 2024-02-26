package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MasterAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "这是自定义中间件"})
	}
}
