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
// @Date 2024-02-13 03:29:05
// @param ctx
// @param path
func RedirectGet(ctx *gin.Context, path string) {
	ctx.Redirect(http.StatusMovedPermanently, path)
}

// RedirectPost
//
// @Title RedirectPost
// @Description: 创建通用的 POST 重定向函数
// @Author liuxingyu
// @Date 2024-02-13 03:28:36
// @param ctx
// @param path
func RedirectPost(ctx *gin.Context, path string) {
	ctx.Redirect(http.StatusFound, path)
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
