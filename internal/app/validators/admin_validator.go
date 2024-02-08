package validators

//	AdminLoginForm
//	@Description: 后台登入验证Form
//
// @Author liuxingyu
// @Date 2024-02-09 00:56:10
type AdminLoginForm struct {
	Tel      string `form:"tel" binding:"required,min=5,max=32" msg:"用户名不能为空|用户名长度不能小于5个字符|用户名长度不能超过32个字符"`
	Password string `form:"password" binding:"required,min=5,max=32" msg:"密码不能为空|密码长度不能小于5个字符|密码长度不能超过32个字符"`
	Remember string `form:"remember"`
}

//	AdminRegisterForm
//	@Description: 后台用户注册
//
// @Author liuxingyu
// @Date 2024-02-09 00:56:05
type AdminRegisterForm struct {
	Name           string `form:"name" binding:"required,min=2,max=20" msg:"姓名不能为空|姓名长度不能小于2个字符|姓名长度不能超过20个字符"`
	Tel            string `form:"tel" binding:"required,min=5,max=32" msg:"电话不能为空|电话长度不能小于5个字符|电话长度不能超过15个字符"`
	Password       string `form:"password" binding:"required,min=5,max=32" msg:"密码不能为空|密码长度不能小于5个字符|密码长度不能超过32个字符"`
	RetypePassword string `form:"retype_password" binding:"required,eqfield=Password" msg:"请重复输入密码|两次输入的密码不一致"`
	Terms          string `form:"terms" binding:"required" msg:"请同意条款"`
}
