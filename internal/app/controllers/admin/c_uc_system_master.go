package admin

import (
	"github.com/gin-gonic/gin"
	"ice_flame_gin/internal/app/validators"
	"ice_flame_gin/internal/system"
)

var UcSystemMaster = cUcSystemMaster{}

type cUcSystemMaster struct{}

func (c *cUcSystemMaster) Login(ctx *gin.Context) {
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
