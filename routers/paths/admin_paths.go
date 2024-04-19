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

	AdminCreateMaster                 = "/master/create"
	AdminHandleCreateMaster           = "/master/create"
	AdminListMaster                   = "/master/list"
	AdminAjaxListMaster               = "/master/ajax/list"
	AdminEditMaster                   = "/master/edit"
	AdminHandleAjaxEditMaster         = "/master/ajax/edit"
	AdminHandleAjaxEditStatusMaster   = "/master/ajax/edit/status"
	AdminEditPasswordMaster           = "/master/edit/password"
	AdminHandleAjaxEditPasswordMaster = "/master/ajax/edit/password"

	AdminCreateVisitCategory       = "/visit-category/create"
	AdminHandleCreateVisitCategory = "/visit-category/create"
	AdminListVisitCategory         = "/visit-category/list"
	AdminAjaxListVisitCategory     = "/visit-category/ajax/list"
	AdminEditVisitCategory         = "/visit-category/edit"
	AdminHandleAjaxVisitCategory   = "/visit-category/ajax/edit"
	AdminListVisitorLogs           = "/visitor-logs/list"
	AdminAjaxListVisitorLogs       = "/visitor-logs/ajax/list"

	AdminCreateSinglePage               = "/single-page/create"
	AdminHandleCreateSinglePage         = "/single-page/create"
	AdminListSinglePage                 = "/single-page/list"
	AdminAjaxListSinglePage             = "/single-page/ajax/list"
	AdminEditSinglePage                 = "/single-page/edit"
	AdminHandleAjaxEditSinglePage       = "/single-page/ajax/edit"
	AdminHandleAjaxEditStatusSinglePage = "/single-page/ajax/edit/status"
	AdminHandleAjaxDeleteSinglePage     = "/single-page/ajax/delete"

	AdminUploadFile  = "/upload/file"
	AdminUploadImage = "/upload/image"

	AdminCreateArticleCategory               = "/article-category/create"
	AdminHandleCreateArticleCategory         = "/article-category/create"
	AdminListArticleCategory                 = "/article-category/list"
	AdminAjaxListArticleCategory             = "/article-category/ajax/list"
	AdminEditArticleCategory                 = "/article-category/edit"
	AdminHandleAjaxEditArticleCategory       = "/article-category/ajax/edit"
	AdminHandelAjaxEditStatusArticleCategory = "/article-category/ajax/edit/status"
	AdminHandelAjaxDeleteArticleCategory     = "/article-category/ajax/delete"

	AdminCreateArticleChannel     = "/article-channel/create"
	AdminHandCreateArticleChannel = "/article-channel/create"
	AdminListArticleChannel       = "/article-channel/list"
	AdminAjaxListArticleChannel   = "/article-channel/ajax/list"
	AdminEditArticle
)
