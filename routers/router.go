package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterRouters
//
// @Title RegisterRouters
// @Description: 注册路由
// @Author liuxingyu
// @Date 2024-01-22 21:15:35
// @param r
// @return *gin.Engine
func RegisterRouters(r *gin.Engine) *gin.Engine {
	// 注册管理后台路由
	setupAdminRoutes(r)
	// 设置一个处理 404 中间件函数
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "admin/error/404.html", nil)
	})
	return r
}
