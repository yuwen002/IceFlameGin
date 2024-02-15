package dto

//	LoginTelPasswordInput
//	@Description: 用户电话密码登入
//
// @Author liuxingyu
// @Date 2024-02-08 01:03:37
type LoginTelPasswordInput struct {
	Tel      string
	Password string
}

//	RegisterSystemMasterInput
//	@Description: 创建管理员新用户
//
// @Author liuxingyu
// @Date 2024-02-08 23:41:24
type RegisterSystemMasterInput struct {
	Password string
	Tel      string
	Name     string
	Email    string
}

// ForgotPasswordSystemMasterInput
//
// @Description: 用户忘记密码
// @Author liuxingyu
// @Date 2024-02-15 19:48:44
type ForgotPasswordSystemMasterInput struct {
	Email string
}
