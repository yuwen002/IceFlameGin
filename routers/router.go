package routers

import "github.com/gin-gonic/gin"

// RegisterRouters
//
// @Title RegisterRouters
// @Description: 注册路由
// @Author liuxingyu
// @Date 2024-01-22 21:15:35
// @return *gin.Engine
func RegisterRouters() *gin.Engine {
	r := gin.Default()
	// 注册管理后台路由
	setupAdminRoutes(r)
	return r
}
