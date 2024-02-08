package routers

import (
	"github.com/gin-gonic/gin"
	"ice_flame_gin/internal/app/controllers/admin"
)

// setupAdminRoutes
//
// @Title 后台路由
// @Description: 后台路由
// @Author liuxingyu
// @Date 2024-01-21 22:28:27
// @param router
func setupAdminRoutes(router *gin.Engine) {
	r := router.Group("/admin")
	{
		r.GET("/", admin.UcSystemMaster.Login)
		r.GET("/login", admin.UcSystemMaster.Login)
		r.POST("/login", admin.UcSystemMaster.HandleLogin)
		r.GET("register", admin.UcSystemMaster.Register)
		r.POST("register", admin.UcSystemMaster.HandleRegister)
	}
}
