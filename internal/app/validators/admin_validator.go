package validators

//	AdminLoginForm
//	@Description: 后台登入验证Form
//
// @Author liuxingyu
// @Date 2024-02-09 00:56:10
type AdminLoginForm struct {
	Tel      string `form:"tel" binding:"required,min=5,max=32" msg:"用户名不能为空|用户名长度不能小于5个字符|用户名长度不能超过32个字符"`
	Password string `form:"password" binding:"required,min=5,max=32" msg:"密码不能为空|密码长度不能小于5个字符|密码长度不能超过32个字符"`
	Remember bool   `form:"remember"`
}

//	AdminRegisterForm
//	@Description: 后台用户注册
//
// @Author liuxingyu
// @Date 2024-02-09 00:56:05
type AdminRegisterForm struct {
	Name           string `form:"name" binding:"required,min=2,max=20" msg:"姓名不能为空|姓名长度不能小于2个字符|姓名长度不能超过20个字符"`
	Email          string `form:"email" binding:"required,email" msg:"电子邮箱不能为空|电子邮箱格式不正确"`
	Tel            string `form:"tel" binding:"required,min=5,max=32" msg:"电话不能为空|电话长度不能小于5个字符|电话长度不能超过15个字符"`
	Password       string `form:"password" binding:"required,min=5,max=32" msg:"密码不能为空|密码长度不能小于5个字符|密码长度不能超过32个字符"`
	RetypePassword string `form:"retype_password" binding:"required,eqfield=Password" msg:"请重复输入密码|两次输入的密码不一致"`
	Terms          string `form:"terms" binding:"required" msg:"请同意条款"`
}

// AdminForgotPassword
//
// @Description: 用户忘记密码
// @Author liuxingyu
// @Date 2024-02-15 01:49:53
type AdminForgotPassword struct {
	Email string `form:"email" binding:"required,email" msg:"电子邮箱不能为空|电子邮箱格式不正确"`
}

// AdminPasswordRecovery
// @Description: 用户充值密码
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-23 11:41:55
type AdminPasswordRecovery struct {
	Password       string `form:"password" binding:"required,min=5,max=32" msg:"密码不能为空|密码长度不能小于5个字符|密码长度不能超过32个字符"`
	RetypePassword string `form:"retype_password" binding:"required,eqfield=Password" msg:"请重复输入密码|两次输入的密码不一致"`
	Token          string `form:"token"`
}

// AdminChangeOwnPassword
// @Description: 管理员修改自己的密码
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-01 17:41:17
type AdminChangeOwnPassword struct {
	OldPassword       string `form:"old_password" binding:"required,min=5,max=32" msg:"密码不能为空|密码长度不能小于5个字符|密码长度不能超过32个字符"`
	NewPassword       string `form:"new_password" binding:"required,min=5,max=32" msg:"密码不能为空|密码长度不能小于5个字符|密码长度不能超过32个字符"`
	RetypeNewPassword string `form:"retype_new_password" binding:"required,eqfield=NewPassword" msg:"请重复输入密码|两次输入的密码不一致"`
}

// AdminChangeMasterInfo
// @Description: 管理员修改个人信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-06 16:57:07
type AdminChangeMasterInfo struct {
	Name  string `form:"name" binding:"required,min=2,max=20" msg:"姓名不能为空|姓名长度不能小于2个字符|姓名长度不能超过20个字符"`
	Email string `form:"email" binding:"required,email" msg:"电子邮箱不能为空|电子邮箱格式不正确"`
	Tel   string `form:"tel" binding:"required,min=5,max=32" msg:"电话不能为空|电话长度不能小于5个字符|电话长度不能超过15个字符"`
}

// AdminRole
//
// @Description: 管理员角色信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-08 22:03:17
type AdminRole struct {
	Name   string `form:"name" binding:"required,min=2,max=20" msg:"角色名称不能为空|角色名称长度不能小于2个字符|角色名称长度不能超过20个字符"`
	Remark string `form:"remark"`
}

// AdminRoleRelation
// @Description: 管理员角色绑定
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-20 16:17:20
type AdminRoleRelation struct {
	AccountID uint32 `form:"account_id" binding:"required,positive_number" msg:"管理员名称不能为空|管理员ID信息错误"`
	RoleID    uint32 `form:"role_id" binding:"required,positive_number" msg:"角色名称不能为空|角色ID信息错误"`
}
