package routers

import (
	"github.com/gin-gonic/gin"
	"ice_flame_gin/internal/app/controllers/admin"
	"ice_flame_gin/internal/app/middlewares"
	"ice_flame_gin/internal/system"
	"ice_flame_gin/routers/paths"
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
		r.GET(paths.AdminPasswordRecovery, admin.UcSystemMaster.RecoverPassword)
		r.POST(paths.AdminHandlePasswordRecovery, admin.UcSystemMaster.HandleRecoverPassword)

		r.Use(middlewares.MasterAuthMiddleware())
		{
			r.GET(paths.AdminLogout, admin.UcSystemMaster.Logout)
			r.GET(paths.AdminDashboard, admin.UcSystemMaster.Dashboard)
			r.GET(paths.AdminChangeOwnPassword, admin.UcSystemMaster.ChangeOwnPassword)
			r.POST(paths.AdminHandleChangeOwnPassword, admin.UcSystemMaster.HandleChangeOwnPassword)
			r.GET(paths.AdminChangeMasterInfo, admin.UcSystemMaster.ChangeMasterInfo)
			r.POST(paths.AdminHandleChangeMasterInfo, admin.UcSystemMaster.HandleChangeMasterInfo)

			// 管理员角色信息
			r.GET(paths.AdminCreateMasterRole, admin.UcSystemMasterRole.CreateMasterRole)
			r.POST(paths.AdminHandleCreateMasterRole, admin.UcSystemMasterRole.HandleCreateMasterRole)
			r.GET(paths.AdminEditMasterRole, admin.UcSystemMasterRole.EditMasterRole)
			r.POST(paths.AdminHandleAjaxEditMasterRole, admin.UcSystemMasterRole.HandleAjaxEditMasterRole)
			r.GET(paths.AdminListMasterRole, admin.UcSystemMasterRole.ListMasterRole)
			r.GET(paths.AdminAjaxListMasterRole, admin.UcSystemMasterRole.AjaxListMasterRole)

			// 管理员角色关联
			r.GET(paths.AdminCreateMasterRoleRelation, admin.UcSystemMasterRole.CreateMasterRoleRelation)
			r.POST(paths.AdminHandleCreateMasterRoleRelation, admin.UcSystemMasterRole.HandleCreateMasterRoleRelation)
			r.GET(paths.AdminListMasterRoleRelation, admin.UcSystemMasterRole.ListMasterRoleRelation)
			r.GET(paths.AdminAjaxListMasterRoleRelation, admin.UcSystemMasterRole.AjaxListMasterRoleRelation)
			r.GET(paths.AdminEditMasterRoleRelation, admin.UcSystemMasterRole.EditMasterRoleRelation)
			r.POST(paths.AdminHandleAjaxEditMasterRoleRelation, admin.UcSystemMasterRole.HandleAjaxEditMasterRoleRelation)

			// 管理员
			r.GET(paths.AdminCreateMaster, admin.UcSystemMaster.CreateSystemMaster)
			r.POST(paths.AdminHandleCreateMaster, admin.UcSystemMaster.HandleCreateSystemMaster)
			r.GET(paths.AdminListMaster, admin.UcSystemMaster.ListSystemMaster)
			r.GET(paths.AdminAjaxListMaster, admin.UcSystemMaster.AjaxListSystemMaster)
			r.GET(paths.AdminEditMaster, admin.UcSystemMaster.EditMaster)
			r.POST(paths.AdminHandleAjaxEditMaster, admin.UcSystemMaster.HandleAjaxEditMaster)

			// 访问类型
			r.GET(paths.AdminCreateVisitCategory, admin.UcSystemMasterVisit.CreateVisitCategory)
			r.POST(paths.AdminHandleCreateVisitCategory, admin.UcSystemMasterVisit.HandleCreateVisitCategory)
			r.GET(paths.AdminListVisitCategory, admin.UcSystemMasterVisit.ListVisitCategory)
			r.GET(paths.AdminAjaxListVisitCategory, admin.UcSystemMasterVisit.AjaxListVisitCategory)
			r.GET(paths.AdminEditVisitCategory, admin.UcSystemMasterVisit.EditVisitCategory)
			r.POST(paths.AdminHandleAjaxVisitCategory, admin.UcSystemMasterVisit.AjaxEditVisitCategory)
			r.GET(paths.AdminListVisitorLogs, admin.UcSystemMasterVisit.ListVisitorLogs)
			r.GET(paths.AdminAjaxListVisitorLogs, admin.UcSystemMasterVisit.AjaxListVisitorLogs)
		}

	}
	// 设置一个处理 404 中间件函数
	r.Use(func(c *gin.Context) {
		system.Render(c, "admin/error/404.html", nil)
	})
}
