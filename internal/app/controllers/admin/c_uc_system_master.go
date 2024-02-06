package admin

import (
	"github.com/gin-gonic/gin"
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
	tel := ""
	// 获取“记住我”cookie中保存的用户名
	cookie, err := ctx.Request.Cookie("admin_tel")
	if err == nil {
		tel = cookie.Value
	}
	system.Render(ctx, "admin/login.html", gin.H{"title": "后台登入", "tel": tel, "checked": tel != ""})
}

// HandleLogin
//
// @Title HandleLogin
// @Description:
// @Author liuxingyu
// @Date 2024-02-06 23:35:18
// @receiver c
// @param ctx
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

	// 设置cookie
	cookieName := "admin_tel"
	// 将用户的用户名保存到cookie中
	if form.Remember {
		ctx.SetCookie(cookieName, form.Tel, 24*3600, "/", "", false, true)
	} else {
		ctx.SetCookie(cookieName, "", -1, "/", "", false, true)
	}

}
