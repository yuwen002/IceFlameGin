package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ice_flame_gin/internal/app/controllers/admin"
	"ice_flame_gin/internal/app/middlewares"
	"ice_flame_gin/internal/system"
	"ice_flame_gin/routers/paths"
	"net/http"
)

// setupAdminRoutes
//
// @Title 后台路由
// @Description: 后台路由
// @Author liuxingyu
// @Date 2024-01-21 22:28:27
// @param router
func setupAdminRoutes(router *gin.Engine) {
	r := router.Group(paths.AdminRoot)
	{
		r.GET(paths.Admin404, func(c *gin.Context) {
			c.AbortWithStatus(404)
		})
		r.GET("/", admin.UcSystemMaster.Login)
		r.GET(paths.AdminLogin, admin.UcSystemMaster.Login)
		r.POST(paths.AdminHandleLogin, admin.UcSystemMaster.HandleLogin)
		r.GET(paths.AdminRegister, admin.UcSystemMaster.Register)
		r.POST(paths.AdminHandleRegister, admin.UcSystemMaster.HandleRegister)
		r.GET(paths.AdminForgotPassword, admin.UcSystemMaster.ForgotPassword)
		r.POST(paths.AdminHandleForgotPassword, admin.UcSystemMaster.HandleForgotPassword)
		r.GET(paths.AdminPasswordRecovery, admin.UcSystemMaster.PasswordRecovery)
		r.POST(paths.AdminHandlePasswordRecovery, admin.UcSystemMaster.HandlePasswordRecovery)

		r.Use(middlewares.MasterAuthMiddleware())
		{
			r.GET(paths.AdminDashboard, admin.UcSystemMaster.Dashboard)
			r.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "系统首页")
				// 获取主用户的 ID 和用户信息
				masterID, _ := c.Get("master_id")
				userInfo, _ := c.Get("user_info")
				fmt.Println(masterID)
				fmt.Println(userInfo)
			})
		}

	}
	// 设置一个处理 404 中间件函数
	r.Use(func(c *gin.Context) {
		system.Render(c, "admin/error/404.html", nil)
	})
}
