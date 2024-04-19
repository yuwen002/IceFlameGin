package routers

import (
	"github.com/flosch/pongo2/v6"
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
			system.Render(c, "admin/error/404.html", pongo2.Context{})
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
			r.GET(paths.AdminEditMaster, admin.UcSystemMaster.EditSystemMaster)
			r.POST(paths.AdminHandleAjaxEditMaster, admin.UcSystemMaster.HandleAjaxEditSystemMaster)
			r.POST(paths.AdminHandleAjaxEditStatusMaster, admin.UcSystemMaster.HandleAjaxEditStatusSystemMaster)
			r.GET(paths.AdminEditPasswordMaster, admin.UcSystemMaster.EditPasswordSystemMaster)
			r.POST(paths.AdminHandleAjaxEditPasswordMaster, admin.UcSystemMaster.HandleAjaxEditPasswordSystemMaster)

			// 访问类型
			r.GET(paths.AdminCreateVisitCategory, admin.UcSystemMasterVisit.CreateVisitCategory)
			r.POST(paths.AdminHandleCreateVisitCategory, admin.UcSystemMasterVisit.HandleCreateVisitCategory)
			r.GET(paths.AdminListVisitCategory, admin.UcSystemMasterVisit.ListVisitCategory)
			r.GET(paths.AdminAjaxListVisitCategory, admin.UcSystemMasterVisit.AjaxListVisitCategory)
			r.GET(paths.AdminEditVisitCategory, admin.UcSystemMasterVisit.EditVisitCategory)
			r.POST(paths.AdminHandleAjaxVisitCategory, admin.UcSystemMasterVisit.AjaxEditVisitCategory)
			r.GET(paths.AdminListVisitorLogs, admin.UcSystemMasterVisit.ListVisitorLogs)
			r.GET(paths.AdminAjaxListVisitorLogs, admin.UcSystemMasterVisit.AjaxListVisitorLogs)

			// 单页信息
			r.GET(paths.AdminCreateSinglePage, admin.SinglePage.CreateSinglePage)
			r.POST(paths.AdminHandleCreateSinglePage, admin.SinglePage.HandleCreateSinglePage)
			r.GET(paths.AdminListSinglePage, admin.SinglePage.ListSinglePage)
			r.GET(paths.AdminAjaxListSinglePage, admin.SinglePage.AjaxListSinglePage)
			r.GET(paths.AdminEditSinglePage, admin.SinglePage.EditSinglePage)
			r.POST(paths.AdminHandleAjaxEditSinglePage, admin.SinglePage.HandleAjaxEditSinglePage)
			r.POST(paths.AdminHandleAjaxEditStatusSinglePage, admin.SinglePage.HandleAjaxEditStatusSinglePage)
			r.POST(paths.AdminHandleAjaxDeleteSinglePage, admin.SinglePage.DeleteSinglePage)

			// 文件上传
			r.POST(paths.AdminUploadFile, admin.SysUploadFile.UploadFile)
			r.POST(paths.AdminUploadImage, admin.SysUploadFile.UploadImage)

			// 文章分类
			r.GET(paths.AdminCreateArticleCategory, admin.Article.CreateArticleCategory)
			r.POST(paths.AdminHandleCreateArticleCategory, admin.Article.HandleCreateArticleCategory)
			r.GET(paths.AdminListArticleCategory, admin.Article.ListArticleCategory)
			r.GET(paths.AdminAjaxListArticleCategory, admin.Article.AjaxListArticleCategory)
			r.GET(paths.AdminEditArticleCategory, admin.Article.EditArticleCategory)
			r.POST(paths.AdminHandleAjaxEditArticleCategory, admin.Article.HandleAjaxEditArticleCategory)
			r.POST(paths.AdminHandelAjaxEditStatusArticleCategory, admin.Article.HandleAjaxEditStatusArticleCategory)
			r.POST(paths.AdminHandelAjaxDeleteArticleCategory, admin.Article.HandelAjaxDeleteArticleCategory)

			// 文章频道
			r.GET(paths.AdminCreateArticleChannel, admin.Article.CreateArticleChannel)
			r.POST(paths.AdminHandCreateArticleChannel, admin.Article.HandelCreateArticleChannel)
			r.GET(paths.AdminListArticleChannel, admin.Article.ListArticleChannel)
			r.GET(paths.AdminAjaxListArticleChannel, admin.Article.AjaxListArticleChannel)
		}

	}
	// 设置一个处理 404 中间件函数
	r.Use(func(c *gin.Context) {
		system.Render(c, "admin/error/404.html", nil)
	})
}
