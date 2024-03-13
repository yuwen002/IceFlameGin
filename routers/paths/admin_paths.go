package paths

// 定义 admin 模块的路径常量
const (
	AdminRoot                    = "/admin"
	Admin404                     = "/page-not-found"
	AdminLogin                   = "/login"
	AdminHandleLogin             = "/login"
	AdminRegister                = "/register"
	AdminHandleRegister          = "/register"
	AdminForgotPassword          = "/forgot-password"
	AdminHandleForgotPassword    = "/forgot-password"
	AdminPasswordRecovery        = "/password-recovery"
	AdminHandlePasswordRecovery  = "/password-recovery"
	AdminDashboard               = "/dashboard"
	AdminChangeOwnPassword       = "/change-own-password"
	AdminHandleChangeOwnPassword = "/change-own-password"
	AdminLogout                  = "/logout"
	AdminChangeMasterInfo        = "/change-master-info"
	AdminHandleChangeMasterInfo  = "/change-master-info"

	AdminCreateMasterRole       = "/master-role/create"
	AdminHandleCreateMasterRole = "/master-role/create"
	AdminEditMasterRole         = "/master-role/edit"
	AdminHandleEditMasterRole   = "/master-role/edit"
	AdminListMasterRole         = "/master-role/list"
)
