package admin

import (
	"github.com/gin-gonic/gin"
	"ice_flame_gin/internal/app/repositories"
	"ice_flame_gin/internal/app/validators"
	"ice_flame_gin/internal/system"
)

var UcSystemMaster = cUcSystemMaster{}

type cUcSystemMaster struct{}

// Login
//
// @Title 后台登入页
// @Description 后台登入显示页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-01 10:56:34
// @receiver c
// @param ctx
func (c *cUcSystemMaster) Login(ctx *gin.Context) {
	repositories.NewUcAccountRepository().GetById()
	system.Render(ctx, "admin/login.html", gin.H{"title": "后台登入"})
}

func (c *cUcSystemMaster) HandleLogin(ctx *gin.Context) {
	var form validators.AdminLoginForm
	if err := ctx.ShouldBind(&form); err != nil {
		// 获取验证错误信息
		errMsg := system.GetValidationErrorMsg(err, form)
		system.Render(ctx, "admin/login.html", gin.H{
			"title": "后台登入",
			"error": errMsg,
		})
		return
	}
}
