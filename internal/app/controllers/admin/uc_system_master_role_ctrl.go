package admin

import (
	"fmt"
	"github.com/flosch/pongo2/v6"
	"github.com/gin-gonic/gin"
	"ice_flame_gin/internal/app/validators"
	"ice_flame_gin/internal/system"
	"ice_flame_gin/routers/paths"
)

// cUcSystemMasterRole
//
// @Description: 用户角色控制器
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-08 00:54:06
type cUcSystemMasterRole struct {
	pageNotFound string
}

// UcSystemMasterRole 初始化用户角色控制器
var UcSystemMasterRole = cUcSystemMasterRole{
	pageNotFound: paths.AdminRoot + paths.Admin404,
}

// ShowCreateMasterRole
//
// @Title ShowCreateMasterRole
// @Description: 渲染创建用户角色页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-08 00:59:47
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMasterRole) ShowCreateMasterRole(c *gin.Context) {
	// 从会话中获取成功信息
	success := system.GetFlashedData(c, "success")
	fmt.Println(success)
	// 从会话中获取错误信息
	var errMsg map[string]interface{}
	err := system.GetDataFromFlash(c, "err_msg", &errMsg)
	if err != nil {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	system.Render(c, "admin/dashboard/create_master_role.html", pongo2.Context(gin.H{
		"title":   "创建用户角色页面",
		"success": success,
		"err_msg": errMsg,
	}))
	return
}

// HandleCreateMasterRole
//
// @Title HandleCreateMasterRole
// @Description: 处理创建用户角色请求
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-08 01:03:29
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMasterRole) HandleCreateMasterRole(c *gin.Context) {
	var form validators.AdminRole
	if err := c.ShouldBind(&form); err != nil {
		system.SetOldInput(c, "name", form.Name)
		system.SetOldInput(c, "remark", form.Remark)

		system.RedirectGet(c, paths.AdminRoot+paths.AdminCreateMasterRole)
		return
	}
}
