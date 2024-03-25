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

	AdminCreateMasterRole         = "/master-role/create"
	AdminHandleCreateMasterRole   = "/master-role/create"
	AdminEditMasterRole           = "/master-role/edit"
	AdminHandleAjaxEditMasterRole = "/master-role/ajax/edit"
	AdminListMasterRole           = "/master-role/list"
	AdminAjaxListMasterRole       = "/master-role/ajax/list"

	AdminCreateMasterRoleRelation         = "/master-role-relation/create"
	AdminHandleCreateMasterRoleRelation   = "/master-role-relation/create"
	AdminListMasterRoleRelation           = "/master-role-relation/list"
	AdminAjaxListMasterRoleRelation       = "/master-role-relation/ajax/list"
	AdminEditMasterRoleRelation           = "/master-role-relation/edit"
	AdminHandleAjaxEditMasterRoleRelation = "/master-role-relation/ajax/edit"

	AdminCreateMaster         = "/master/create"
	AdminHandleCreateMaster   = "/master/create"
	AdminListMaster           = "/master/list"
	AdminAjaxListMaster       = "/master/ajax/list"
	AdminEditMaster           = "/master/edit"
	AdminHandleAjaxEditMaster = "/master/ajax/edit"

	AdminCreateVisitorCategory       = "/visitor-category/create"
	AdminHandleCreateVisitorCategory = "/visitor-category/create"
)
