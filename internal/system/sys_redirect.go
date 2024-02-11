package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RedirectGet
//
// @Title RedirectGet
// @Description: 创建通用的 GET 重定向函数
// @Author liuxingyu
// @Date 2024-02-11 23:43:24
// @param target
// @return gin.HandlerFunc
func RedirectGet(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, target)
	}
}

// RedirectPost
//
// @Title RedirectPost
// @Description: 创建通用的 POST 重定向函数
// @Author liuxingyu
// @Date 2024-02-11 23:43:43
// @param target
// @return gin.HandlerFunc
func RedirectPost(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Redirect(http.StatusFound, target)
	}
}

// Redirect
//
// @Title Redirect 路由重定向
// @Description:
// @Author liuxingyu
// @Date 2024-02-11 23:44:57
// @param c
// @param target
func Redirect(c *gin.Context, target string) {
	c.Request.URL.Path = target
	c.Request.Method = http.MethodGet
	c.Next()
}
