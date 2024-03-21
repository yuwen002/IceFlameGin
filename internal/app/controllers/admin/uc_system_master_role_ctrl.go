package admin

import (
	"fmt"
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
		output := services.NewUcSystemMasterRoleService().CreateMasterRole(dto.SystemMasterRoleInput{
			Name:   form.Name,
			Remark: form.Remark,
		})

		if output.Code == 1 {
			system.AddFlashData(c, output.Message, "fail")
		} else {
			system.AddFlashData(c, "添加角色信息成功", "success")
		}
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
	start, err := utils.ToInt(c.DefaultQuery("start", "0"))
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
	if err != nil {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

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

// HandleAjaxEditMasterRole
//
// @Title HandleAjaxEditMasterRole
// @Description: Ajax处理编辑用户角色请求
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-17 02:18:18
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMasterRole) HandleAjaxEditMasterRole(c *gin.Context) {
	// 获取要编辑的用户角色ID
	id := c.Query("id")
	uint32ID, err := utils.ToUint32(id)
	if err != nil {
		c.JSON(http.StatusOK, &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var form validators.AdminRole
	if err := c.ShouldBind(&form); err != nil {
		// 获取验证错误信息
		errMsg := system.GetValidationErrors(err, form)
		c.JSON(http.StatusOK, &system.SysResponse{
			Code:    2,
			Message: "验证错误",
			Data:    errMsg,
		})
		return
	}
	// 根据 ID 更新用户角色信息
	output := services.NewUcSystemMasterRoleService().ChangeMasterRoleById(uint32ID, dto.SystemMasterRoleInput{
		Name:   form.Name,
		Remark: form.Remark,
	})
	if output.Code == 1 {
		c.JSON(http.StatusOK, &system.SysResponse{
			Code:    1,
			Message: output.Message,
			Data:    nil,
		})
		return
	}

	// 更新成功后，可以跳转到用户角色列表页面或显示成功信息
	c.JSON(http.StatusOK, &system.SysResponse{
		Code:    0,
		Message: "Success",
		Data:    nil,
	})
	return
}

// CreateMasterRoleRelation
//
// @Title CreateMasterRoleRelation
// @Description: 渲染创建管理员角色关联页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-20 15:29:52
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMasterRole) CreateMasterRoleRelation(c *gin.Context) {
	// 从会话中获取成功信息
	success := system.GetFlashedData(c, "success")

	// 从会话中获取错误信息
	nonFieldError := system.GetFlashedData(c, "non_field_error")
	var errMsg map[string]interface{}
	err := system.GetDataFromFlash(c, "err_msg", &errMsg)
	if err != nil {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	var form validators.AdminRoleRelation
	err = system.GetDataFromFlash(c, "form", &form)
	if err != nil {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	// 角色信息列表
	output := services.NewUcSystemMasterRoleService().ShowMasterRoleAll()
	if output.Code == 1 {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	roles, ok := output.Data.(map[uint32]string)
	if !ok {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	// 管理员信息列表
	output = services.NewUcSystemMasterService().ShowMasterAll()
	if output.Code == 1 {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	masters, ok := output.Data.(map[uint32]string)
	if !ok {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	system.Render(c, "admin/system_master_role_relation/create.html", pongo2.Context{
		"title":           "新建管理员角色关联",
		"roles":           roles,
		"masters":         masters,
		"success":         success,
		"non_field_error": nonFieldError,
		"err_msg":         errMsg,
		"form":            form,
	})
	return
}

func (ctrl *cUcSystemMasterRole) HandleCreateMasterRoleRelation(c *gin.Context) {
	var form validators.AdminRoleRelation

	if err := c.ShouldBind(&form); err != nil {
		fmt.Println("err:", err)
		fmt.Println("form:", form)
		// 获取验证错误信息
		errMsg := system.GetValidationErrors(err, form)
		// 将错误信息存储到会话中
		if _, ok := errMsg["non-field-error"]; ok {
			system.AddFlashData(c, errMsg["non-field-error"], "non_field_error")
		} else {
			errFlash := system.AddDataToFlash(c, errMsg, "err_msg")
			if errFlash != nil {
				system.RedirectGet(c, ctrl.pageNotFound)
				return
			}

			errFlash = system.AddDataToFlash(c, form, "form")
			if errFlash != nil {
				system.RedirectGet(c, ctrl.pageNotFound)
				return
			}
		}
	} else {
		accountID, errUint := utils.ToUint32(form.AccountID)
		if errUint != nil {
			system.RedirectGet(c, ctrl.pageNotFound)
			return
		}

		roleID, errUint := utils.ToUint32(form.RoleID)

		output := services.NewUcSystemMasterRoleService().CreateMasterRoleRelation(dto.SystemMasterRoleRelationInput{
			AccountId: accountID,
			RoleId:    roleID,
		})
		if output.Code == 1 {
			system.AddFlashData(c, "添加管理员角色绑定失败", "fail")
		} else {
			system.AddFlashData(c, "添加管理员角色绑定成功", "success")
		}
	}

	//system.RedirectGet(c, paths.AdminRoot+paths.AdminCreateMasterRoleRelation)
}

// ListMasterRoleRelation
//
// @Title ListMasterRoleRelation
// @Description: 管理员角色关联显示页
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-19 23:17:26
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMasterRole) ListMasterRoleRelation(c *gin.Context) {
	system.Render(c, "admin/system_master_role_relation/list.html", pongo2.Context{
		"title": "管理员角色关联列表",
	})
}

// AjaxListMasterRoleRelation
//
// @Title AjaxListMasterRoleRelation
// @Description: Ajax获取管理员角色关联
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-19 23:47:02
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMasterRole) AjaxListMasterRoleRelation(c *gin.Context) {
	start, err := utils.ToInt(c.DefaultQuery("start", "0"))
	if err != nil {
		start = 0
	}

	length, err := utils.ToInt(c.DefaultQuery("length", "10"))
	if err != nil {
		length = 10
	}
	fmt.Println(length)

	output := services.NewUcSystemMasterRoleService().ShowMasterRoleRelation(dto.ListSystemMasterRoleRelationInput{
		Order:  "id desc",
		Start:  start,
		Length: length,
	})

	if output.Code == 1 {
		system.EmptyJSON(c)
		return
	}

	data := output.Data.(dto.ListSystemMasterRoleRoleRelationOutput)
	c.JSON(http.StatusOK, gin.H{
		"draw":            c.Query("draw"),
		"data":            data.List,
		"recordsTotal":    data.Total,
		"recordsFiltered": data.Total,
	})

}
