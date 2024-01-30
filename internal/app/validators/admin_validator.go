package validators

type AdminLoginForm struct {
	Username string `form:"username" binding:"required,min=5,max=32" msg:"用户名不能为空|用户名长度不能小于5个字符|用户名长度不能超过32个字符"`
	Password string `form:"password" binding:"required,min=5,max=32" msg:"密码不能为空|密码长度不能小于5个字符|密码长度不能超过32个字符"`
}
