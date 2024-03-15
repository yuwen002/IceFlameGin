package admin

import (
	"github.com/flosch/pongo2/v6"
	"github.com/gin-gonic/gin"
	dto "ice_flame_gin/internal/app/dto/d_uc_center"
	"ice_flame_gin/internal/app/models/model"
	services "ice_flame_gin/internal/app/services/s_uc_center"
	"ice_flame_gin/internal/app/validators"
	"ice_flame_gin/internal/pkg/utils"
	"ice_flame_gin/internal/system"
	"ice_flame_gin/routers/paths"
	"net/http"
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

// CreateMasterRole
//
// @Title CreateMasterRole
// @Description: 渲染创建用户角色页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-08 00:59:47
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMasterRole) CreateMasterRole(c *gin.Context) {
	// 从会话中获取成功信息
	success := system.GetFlashedData(c, "success")
	// 从会话中获取错误信息
	var errMsg map[string]interface{}
	err := system.GetDataFromFlash(c, "err_msg", &errMsg)
	if err != nil {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	system.Render(c, "admin/system_master_role/create.html", pongo2.Context{
		"title":   "创建用户角色",
		"success": success,
		"err_msg": errMsg,
	})
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
		// 获取验证错误信息
		errMsg := system.GetValidationErrors(err, form)
		// 将错误信息存储到会话中
		errFlash := system.AddDataToFlash(c, errMsg, "err_msg")
		if errFlash != nil {
			system.RedirectGet(c, ctrl.pageNotFound)
			return
		}

		system.SetOldInput(c, "name", form.Name)
		system.SetOldInput(c, "remark", form.Remark)

	} else {
		system.AddFlashData(c, "添加角色信息成功", "success")
	}

	system.RedirectGet(c, paths.AdminRoot+paths.AdminCreateMasterRole)
}

// ListMasterRole
//
// @Title ListMasterRole
// @Description: 获取角色显示页
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-15 11:08:55
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMasterRole) ListMasterRole(c *gin.Context) {
	// 渲染用户角色列表页面
	system.Render(c, "admin/system_master_role/list.html", pongo2.Context{
		"title": "用户角色列表",
	})
}

// AjaxListMasterRole
//
// @Title AjaxListMasterRole
// @Description: Ajax获取角色列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-15 11:08:15
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMasterRole) AjaxListMasterRole(c *gin.Context) {
	start, err := utils.ToInt(c.DefaultQuery("start", "1"))
	if err != nil {
		start = 0
	}

	length, err := utils.ToInt(c.DefaultQuery("length", "10"))
	if err != nil {
		length = 10
	}

	output := services.NewUcSystemMasterRoleService().ShowMasterRole(dto.ListSystemMasterRoleInput{
		Order:  "id desc",
		Start:  start,
		Length: length,
	})

	if output.Code == 1 {
		system.EmptyJSON(c)
		return
	}

	data := output.Data.(dto.ListSystemMasterRoleOutput)
	c.JSON(http.StatusOK, gin.H{
		"draw":            c.Query("draw"),
		"data":            data.List,
		"recordsTotal":    data.Total,
		"recordsFiltered": data.Total,
	})
}

// EditMasterRole
//
// @Title EditMasterRole
// @Description: 渲染创编辑户角色页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-15 16:13:32
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMasterRole) EditMasterRole(c *gin.Context) {
	id, err := utils.ToInt(c.Query("id"))
	if err != nil {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	uint32ID, err := utils.ToUint32(id)

	output := services.NewUcSystemMasterRoleService().GetMasterRoleById(uint32ID)
	if output.Code == 1 {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	role, ok := output.Data.(*model.UcSystemMasterRole)
	if !ok {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	// 渲染编辑用户角色页面
	system.Render(c, "admin/system_master_role/edit.html", pongo2.Context{
		"title": "编辑用户角色页面",
		"role":  role,
		"id":    id,
	})
}

func (ctrl *cUcSystemMasterRole) HandleEditMasterRole(c *gin.Context) {
	// 获取要编辑的用户角色ID
	//id := c.Param("id")

	var form validators.AdminRole
	if err := c.ShouldBind(&form); err != nil {
		// 获取验证错误信息
		errMsg := system.GetValidationErrors(err, form)
		// 将错误信息存储到会话中
		errFlash := system.AddDataToFlash(c, errMsg, "err_msg")
		if errFlash != nil {
			system.RedirectGet(c, ctrl.pageNotFound)
			return
		}

		system.SetOldInput(c, "name", form.Name)
		system.SetOldInput(c, "remark", form.Remark)

	} else {
		// 根据 ID 更新用户角色信息
		// 更新成功后，可以跳转到用户角色列表页面或显示成功信息
		system.AddFlashData(c, "更新角色信息成功", "success")
	}

	system.RedirectGet(c, paths.AdminRoot+paths.AdminEditMasterRole)
}
