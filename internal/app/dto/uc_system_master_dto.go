package dto

import "ice_flame_gin/internal/app/models/model"

//	LoginTelPasswordSystemMasterInput
//	@Description: 用户电话密码登入
//
// @Author liuxingyu
// @Date 2024-02-08 01:03:37
type LoginTelPasswordSystemMasterInput struct {
	Tel      string
	Password string
}

// SystemMasterTokenOutput
// @Description: 登入后输出token
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-27 17:24:51
type SystemMasterTokenOutput struct {
	Token     string
	AccountID uint32
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

// ChangePasswordInput
//
// @Description: 修改密码
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-01 23:26:01
type ChangePasswordInput struct {
	ID       uint32
	Password string
}

// ChangeOwnPasswordInput
//
// @Description: 修改用户自己的密码
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-01 23:36:12
type ChangeOwnPasswordInput struct {
	ID          uint32
	NewPassword string
	OldPassword string
}

// ChangeMasterInfoInput
// @Description: 修改用户的个人信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-06 16:29:54
type ChangeMasterInfoInput struct {
	ID    uint32
	Email string
	Name  string
	Tel   string
}

// CreateSystemMasterInput
//
// @Description: 后台新建管理员
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-23 00:16:51
type CreateSystemMasterInput struct {
	Tel      string
	Name     string
	Email    string
	Password string
}

// ListSystemMasterInput
//
// @Description: 后台管理员列表输入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-24 22:41:48
type ListSystemMasterInput struct {
	Order  string
	Start  int
	Length int
}

// ListSystemMasterOutput
//
// @Description: 后台管理员列表输出
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-24 22:59:43
type ListSystemMasterOutput struct {
	List  []*model.UcSystemMaster
	Total int64
}
