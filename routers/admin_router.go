package routers

import "github.com/gin-gonic/gin"

// setupAdminRoutes
//
// @Title 后台路由
// @Description: 后台路由
// @Author liuxingyu
// @Date 2024-01-21 22:28:27
// @param router
func setupAdminRoutes(router *gin.Engine) {
	admin := router.Group("/admin")
	{
		admin.GET("/", adminIndex)
		admin.GET("/users", adminUsers)
		admin.POST("/users", adminCreateUser)
		// 添加更多后台路由...
	}
}
