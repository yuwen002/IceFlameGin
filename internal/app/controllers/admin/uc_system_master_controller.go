package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	dto "ice_flame_gin/internal/app/dto/d_uc_center"
	services "ice_flame_gin/internal/app/services/s_uc_center"
	"ice_flame_gin/internal/app/validators"
	"ice_flame_gin/internal/system"
	"ice_flame_gin/routers/paths"
	"net/http"
)

var UcSystemMaster = cUcSystemMaster{
	pageNotFound: paths.AdminRoot + paths.Admin404,
}

type cUcSystemMaster struct {
	pageNotFound string
}

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

	// 获取错误信息
	errMsg := system.GetFlashedData(ctx, "error")
	cookie, err := ctx.Request.Cookie("admin_tel")
	if err == nil {
		tel = cookie.Value
	}
	system.Render(ctx, "admin/login.html", gin.H{
		"title":   "后台登入",
		"tel":     tel,
		"checked": tel != "",
		"error":   errMsg,
	})
}

// HandleLogin
//
// @Title HandleLogin
// @Description: 后台登入页,处理验证页面
// @Author liuxingyu
// @Date 2024-02-06 23:35:18
// @receiver c
// @param ctx
func (c *cUcSystemMaster) HandleLogin(ctx *gin.Context) {
	var form validators.AdminLoginForm
	path := paths.AdminRoot + paths.AdminLogin
	fmt.Println(form)
	// 解析表单数据
	if err := ctx.ShouldBind(&form); err != nil {
		// 获取验证错误信息
		errMsg := system.GetValidationErrorMsg(err, form)
		system.AddFlashData(ctx, errMsg, "error")
		system.RedirectGet(ctx, path)
		return
	}

	// 调用登录服务进行验证
	output := services.NewUcSystemMasterService().LoginTelPassword(dto.LoginTelPasswordInput{
		Tel:      form.Tel,
		Password: form.Password,
	})
	// 验证不通过，渲染带有错误信息的登录页面
	if output.Code != 0 {
		system.AddFlashData(ctx, "用户名密码错误", "error")
		system.RedirectGet(ctx, path)
		return
	}

	// 设置cookie
	cookieName := "admin_tel"
	// 将用户的用户名保存到cookie中
	if form.Remember {
		ctx.SetCookie(cookieName, form.Tel, 24*3600, "/", "", false, true)
	} else {
		// 删除cookie
		ctx.SetCookie(cookieName, "", -1, "/", "", false, true)
	}

	// 输出登入成功的JSON
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "登入成功",
	})
}

// Register
//
// @Title Register
// @Description: 管理员注册显示页面
// @Author liuxingyu
// @Date 2024-02-09 01:02:45
// @receiver c
// @param ctx
func (c *cUcSystemMaster) Register(ctx *gin.Context) {
	// 从会话中获取错误信息
	var errMsg map[string]interface{}
	err := system.GetDataFromFlash(ctx, "err", &errMsg)
	if err != nil {
		system.RedirectGet(ctx, c.pageNotFound)
		return
	}

	var form validators.AdminRegisterForm
	err = system.GetDataFromFlash(ctx, "form", &form)
	if err != nil {
		system.RedirectGet(ctx, c.pageNotFound)
		return
	}

	fail := system.GetFlashedData(ctx, "fail")

	system.Render(ctx, "admin/register.html", gin.H{
		"title": "管理员注册",
		"error": errMsg,
		"fail":  fail,
		"form":  form,
	})
}

// HandleRegister
//
// @Title HandleRegister
// @Description: 管理员注册页,处理验证页面
// @Author liuxingyu
// @Date 2024-02-09 01:03:20
// @receiver c
// @param ctx
func (c *cUcSystemMaster) HandleRegister(ctx *gin.Context) {
	var form validators.AdminRegisterForm
	path := paths.AdminRoot + paths.AdminRegister
	if err := ctx.ShouldBind(&form); err != nil {
		// 获取验证错误信息
		errMsg := system.GetValidationErrors(err, form)
		// 将错误信息存储到会话中
		errFlash := system.AddDataToFlash(ctx, errMsg, "err")
		if errFlash != nil {
			system.RedirectGet(ctx, c.pageNotFound)
			return
		}
		errFlash = system.AddDataToFlash(ctx, form, "form")
		if errFlash != nil {
			system.RedirectGet(ctx, c.pageNotFound)
			return
		}
		// 跳转注册页
		system.RedirectGet(ctx, path)
		return
	}

	//  注册用户
	output := services.NewUcSystemMasterService().Register(dto.RegisterSystemMasterInput{
		Password: form.Password,
		Tel:      form.Tel,
		Name:     form.Name,
		Email:    form.Email,
	})

	// 用户注册失败
	if output.Code != 0 {
		system.AddFlashData(ctx, output.Message, "fail")
		system.RedirectGet(ctx, path)
		return
	}

	// 跳转到登入页面
	system.RedirectGet(ctx, paths.AdminRoot+paths.AdminLogin)
}

// ForgotPassword
//
// @Title ForgotPassword
// @Description: 忘记密码
// @Author liuxingyu
// @Date 2024-02-15 01:50:20
// @receiver c
// @param ctx
func (c *cUcSystemMaster) ForgotPassword(ctx *gin.Context) {
	// 从会话中获取错误信息
	var errMsg map[string]interface{}
	err := system.GetDataFromFlash(ctx, "err", &errMsg)
	if err != nil {
		system.RedirectGet(ctx, c.pageNotFound)
		return
	}

	var form validators.AdminForgotPassword
	err = system.GetDataFromFlash(ctx, "form", &form)
	if err != nil {
		system.RedirectGet(ctx, c.pageNotFound)
		return
	}

	msg := system.GetFlashedData(ctx, "msg")
	fmt.Println(msg)

	system.Render(ctx, "admin/forgot_password.html", gin.H{
		"title": "忘记密码",
		"error": errMsg,
		"form":  form,
		"msg":   msg,
	})
}

// HandleForgotPassword
//
// @Title HandleForgotPassword
// @Description: 忘记密码处理页面，并发送邮件
// @Author liuxingyu
// @Date 2024-02-15 19:45:21
// @receiver c
// @param ctx
func (c *cUcSystemMaster) HandleForgotPassword(ctx *gin.Context) {
	var form validators.AdminForgotPassword
	var path string
	path = paths.AdminRoot + paths.AdminForgotPassword
	if err := ctx.ShouldBind(&form); err != nil {
		// 获取验证错误信息
		errMsg := system.GetValidationErrors(err, form)
		// 将错误信息存储到会话中
		errFlash := system.AddDataToFlash(ctx, errMsg, "err")
		if errFlash != nil {
			system.RedirectGet(ctx, c.pageNotFound)
			return
		}
		errFlash = system.AddDataToFlash(ctx, form, "form")
		if errFlash != nil {
			system.RedirectGet(ctx, c.pageNotFound)
			return
		}

		// 跳转忘记密码页
		system.RedirectGet(ctx, path)
		return
	}

	output := services.NewUcSystemMasterService().ForgotPassword(dto.ForgotPasswordSystemMasterInput{
		Email: "yuwen002@163.com",
	})
	fmt.Println(output)
	if output.Code == 1 {
		system.AddFlashData(ctx, output.Message, "msg")
		// 跳转忘记密码页
		system.RedirectGet(ctx, path)
		return
	}

	system.AddFlashData(ctx, "发送邮件成功，请查看邮件。", "msg")
	// 跳转忘记密码页
	system.RedirectGet(ctx, path)
	return
}
