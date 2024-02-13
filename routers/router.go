package routers

import "github.com/gin-gonic/gin"

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
	return r
}
