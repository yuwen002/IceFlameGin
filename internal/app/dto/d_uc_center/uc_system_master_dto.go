package dto

//	LoginTelPasswordSystemMasterInput
//	@Description: 用户电话密码登入
//
// @Author liuxingyu
// @Date 2024-02-08 01:03:37
type LoginTelPasswordSystemMasterInput struct {
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

// ForgotPasswordInput
//
// @Description: 用户忘记密码
// @Author liuxingyu
// @Date 2024-02-15 19:48:44
type ForgotPasswordInput struct {
	Email string
}

// PasswordRecoveryOutput
//
// @Description: 重置密码
// @Author liuxingyu
// @Date 2024-02-25 00:21:21
type PasswordRecoveryOutput struct {
	Email     string `json:"email"`
	Timestamp int64  `json:"timestamp"`
}

// EncryptTokenOutput
//
// @Description: 密文输出
// @Author liuxingyu
// @Date 2024-02-25 00:50:33
type EncryptTokenOutput struct {
	CiphertextHex string `json:"ciphertext_hex"`
}
