package admin

import (
	"github.com/gin-gonic/gin"
	dto "ice_flame_gin/internal/app/dto/d_uc_center"
	services "ice_flame_gin/internal/app/services/s_uc_center"
	"ice_flame_gin/internal/app/validators"
	"ice_flame_gin/internal/system"
	"ice_flame_gin/routers/paths"
	"net/http"
	"sync"
)

var UcSystemMaster = cUcSystemMaster{
	pageNotFound: paths.AdminRoot + paths.Admin404,
}

type cUcSystemMaster struct {
	pageNotFound string
	usedTokens   map[string]int64
	mu           sync.Mutex
}

// Login
//
// @Title 后台登入页
// @Description 后台登入显示页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-01 10:56:34
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMaster) Login(c *gin.Context) {
	tel := ""
	// 获取“记住我”cookie中保存的用户名

	// 获取错误信息
	errMsg := system.GetFlashedData(c, "error")
	cookie, err := c.Request.Cookie("admin_tel")
	if err == nil {
		tel = cookie.Value
	}
	system.Render(c, "admin/login.html", gin.H{
		"title":   "后台登入",
		"tel":     tel,
		"checked": tel != "",
		"error":   errMsg,
	})
	return
}

// HandleLogin
//
// @Title HandleLogin
// @Description: 后台登入页,处理验证页面
// @Author liuxingyu
// @Date 2024-02-06 23:35:18
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMaster) HandleLogin(c *gin.Context) {
	var form validators.AdminLoginForm
	path := paths.AdminRoot + paths.AdminLogin
	// 解析表单数据
	if err := c.ShouldBind(&form); err != nil {
		// 获取验证错误信息
		errMsg := system.GetValidationErrorMsg(err, form)
		system.AddFlashData(c, errMsg, "error")
		system.RedirectGet(c, path)
		return
	}

	// 调用登录服务进行验证
	output := services.NewUcSystemMasterService().LoginTelPassword(dto.LoginTelPasswordSystemMasterInput{
		Tel:      form.Tel,
		Password: form.Password,
	})
	// 验证不通过，渲染带有错误信息的登录页面
	if output.Code != 0 {
		system.AddFlashData(c, "用户名密码错误", "error")
		system.RedirectGet(c, path)
		return
	}

	token, ok := output.Data.(dto.SystemMasterTokenOutput)
	if !ok {
		system.AddFlashData(c, "登入失败", "error")
		system.RedirectGet(c, path)
		return
	}

	// 将token写入session
	system.SetSession(c, "ice-flame-master", token)

	// 设置cookie
	cookieName := "admin_tel"
	// 将用户的用户名保存到cookie中
	if form.Remember {
		c.SetCookie(cookieName, form.Tel, 24*3600, "/", "", false, true)
	} else {
		// 删除cookie
		c.SetCookie(cookieName, "", -1, "/", "", false, true)
	}

	// 输出登入成功的JSON
	c.JSON(http.StatusOK, gin.H{
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
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMaster) Register(c *gin.Context) {
	// 从会话中获取错误信息
	var errMsg map[string]interface{}
	err := system.GetDataFromFlash(c, "err", &errMsg)
	if err != nil {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	var form validators.AdminRegisterForm
	err = system.GetDataFromFlash(c, "form", &form)
	if err != nil {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	fail := system.GetFlashedData(c, "fail")

	system.Render(c, "admin/register.html", gin.H{
		"title": "管理员注册",
		"error": errMsg,
		"fail":  fail,
		"form":  form,
	})
	return
}

// HandleRegister
//
// @Title HandleRegister
// @Description: 管理员注册页,处理验证页面
// @Author liuxingyu
// @Date 2024-02-09 01:03:20
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMaster) HandleRegister(c *gin.Context) {
	var form validators.AdminRegisterForm
	path := paths.AdminRoot + paths.AdminRegister
	if err := c.ShouldBind(&form); err != nil {
		// 获取验证错误信息
		errMsg := system.GetValidationErrors(err, form)
		// 将错误信息存储到会话中
		errFlash := system.AddDataToFlash(c, errMsg, "err")
		if errFlash != nil {
			system.RedirectGet(c, ctrl.pageNotFound)
			return
		}
		errFlash = system.AddDataToFlash(c, form, "form")
		if errFlash != nil {
			system.RedirectGet(c, ctrl.pageNotFound)
			return
		}
		// 跳转注册页
		system.RedirectGet(c, path)
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
		system.AddFlashData(c, output.Message, "fail")
		system.RedirectGet(c, path)
		return
	}

	// 跳转到登入页面
	system.RedirectGet(c, paths.AdminRoot+paths.AdminLogin)
	return
}

// ForgotPassword
//
// @Title ForgotPassword
// @Description: 忘记密码
// @Author liuxingyu
// @Date 2024-02-15 01:50:20
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMaster) ForgotPassword(c *gin.Context) {
	// 从会话中获取错误信息
	var errMsg map[string]interface{}
	err := system.GetDataFromFlash(c, "err", &errMsg)
	if err != nil {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	var form validators.AdminForgotPassword
	err = system.GetDataFromFlash(c, "form", &form)
	if err != nil {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	msg := system.GetFlashedData(c, "msg")

	system.Render(c, "admin/forgot_password.html", gin.H{
		"title": "忘记密码",
		"error": errMsg,
		"form":  form,
		"msg":   msg,
	})
	return
}

// HandleForgotPassword
//
// @Title HandleForgotPassword
// @Description: 忘记密码处理页面，并发送邮件
// @Author liuxingyu
// @Date 2024-02-15 19:45:21
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMaster) HandleForgotPassword(c *gin.Context) {
	var form validators.AdminForgotPassword
	var path string
	path = paths.AdminRoot + paths.AdminForgotPassword
	if err := c.ShouldBind(&form); err != nil {
		// 获取验证错误信息
		errMsg := system.GetValidationErrors(err, form)
		// 将错误信息存储到会话中
		errFlash := system.AddDataToFlash(c, errMsg, "err")
		if errFlash != nil {
			system.RedirectGet(c, ctrl.pageNotFound)
			return
		}
		errFlash = system.AddDataToFlash(c, form, "form")
		if errFlash != nil {
			system.RedirectGet(c, ctrl.pageNotFound)
			return
		}

		// 跳转忘记密码页
		system.RedirectGet(c, path)
		return
	}

	output := services.NewUcSystemMasterService().ForgotPassword(dto.ForgotPasswordInput{
		Email: form.Email,
	})
	if output.Code == 1 {
		system.AddFlashData(c, output.Message, "msg")
		// 跳转忘记密码页
		system.RedirectGet(c, path)
		return
	}

	system.AddFlashData(c, "发送邮件成功，请查看邮件。", "msg")
	// 跳转忘记密码页
	system.RedirectGet(c, path)
	return
}

// PasswordRecovery
//
// @Title PasswordRecovery
// @Description: 密码恢复,设置新密码
// @Author liuxingyu
// @Date 2024-02-24 01:10:50
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMaster) PasswordRecovery(c *gin.Context) {
	title := "重置密码"
	// 获取URL参数的值
	token := c.Query("token")
	output := services.NewUcSystemMasterService().DecryptToken(token)
	if output.Code == 1 {
		system.Render(c, "admin/password_recovery.html", gin.H{
			"title": title,
			"fail":  output.Message,
		})
		return
	}

	// 从会话中获取错误信息
	var errMsg map[string]interface{}
	err := system.GetDataFromFlash(c, "err", &errMsg)
	if err != nil {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	fail := system.GetFlashedData(c, "fail")

	system.Render(c, "admin/password_recovery.html", gin.H{
		"title": title,
		"token": token,
		"error": errMsg,
		"fail":  fail,
	})
	return
}

// HandlePasswordRecovery
//
// @Title HandlePasswordRecovery
// @Description: 密码恢复,设置新密码，处理页面
// @Author liuxingyu
// @Date 2024-02-25 22:53:33
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMaster) HandlePasswordRecovery(c *gin.Context) {
	var form validators.AdminPasswordRecovery

	if err := c.ShouldBind(&form); err != nil {
		// 获取验证错误信息
		errMsg := system.GetValidationErrors(err, form)
		// 将错误信息存储到会话中
		errFlash := system.AddDataToFlash(c, errMsg, "err")
		if errFlash != nil {
			system.RedirectGet(c, ctrl.pageNotFound)
			return
		}

		// 获取 referer，即为 POST 请求前的 URL
		referer := c.Request.Referer()
		system.RedirectGet(c, referer)
		return
	}

	//  重置密码
	output := services.NewUcSystemMasterService().PasswordRecovery(form.Token, form.Password)
	if output.Code == 1 {
		system.AddFlashData(c, output.Message, "fail")
		// 获取 referer，即为 POST 请求前的 URL
		referer := c.Request.Referer()
		system.RedirectGet(c, referer)
		return
	}
}
