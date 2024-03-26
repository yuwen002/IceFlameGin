package admin

import (
	"github.com/flosch/pongo2/v6"
	"github.com/gin-gonic/gin"
	"ice_flame_gin/internal/app/dto"
	"ice_flame_gin/internal/app/models/model"
	"ice_flame_gin/internal/app/services"
	"ice_flame_gin/internal/app/validators"
	"ice_flame_gin/internal/pkg/utils"
	"ice_flame_gin/internal/system"
	"ice_flame_gin/routers/paths"
	"net/http"
)

// cUcSystemMasterVisit
// @Description: 用户访问类型
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 15:09:53
type cUcSystemMasterVisit struct {
	pageNotFound string
}

// UcSystemMasterVisit 初始化
var UcSystemMasterVisit = cUcSystemMasterVisit{
	pageNotFound: paths.AdminRoot + paths.Admin404,
}

// CreateVisitCategory
//
// @Title CreateVisitCategory
// @Description: 渲染新建用户访问类型分类
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 20:22:03
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMasterVisit) CreateVisitCategory(c *gin.Context) {
	// 从会话中获取成功信息
	success := system.GetFlashedData(c, "success")
	// 从会话中获取错误信息
	fail := system.GetFlashedData(c, "fail")

	var errMsg map[string]interface{}
	err := system.GetDataFromFlash(c, "err_msg", &errMsg)
	if err != nil {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	system.Render(c, "admin/system_master_visit/create.html", pongo2.Context{
		"title":   "新建访问类型",
		"success": success,
		"fail":    fail,
		"err_msg": errMsg,
	})
}

// HandleCreateVisitCategory
//
// @Title HandleCreateVisitCategory
// @Description: 处理新建用户访问类型分类请求
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 20:22:33
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMasterVisit) HandleCreateVisitCategory(c *gin.Context) {
	var form validators.AdminVisitCategory

	if err := c.ShouldBind(&form); err != nil {
		// 获取验证错误信息
		errMsg := system.GetValidationErrors(err, form)
		// 将错误信息存储到会话中
		errFlash := system.AddDataToFlash(c, errMsg, "err_msg")
		if errFlash != nil {
			system.RedirectGet(c, ctrl.pageNotFound)
			return
		}

		system.SetOldInput(c, "title", form.Title)
	} else {
		output := services.NewUcSystemMasterVisit().CreateVisitCategory(dto.SystemMasterVisitCategoryInput{
			Title: form.Title,
		})

		if output.Code == 1 {
			system.AddFlashData(c, output.Message, "fail")
		} else {
			system.AddFlashData(c, "添加访问类型分类信息成功", "success")
		}
	}

	system.RedirectGet(c, paths.AdminRoot+paths.AdminCreateVisitCategory)
}

// ListVisitCategory
//
// @Title ListVisitCategory
// @Description: 渲染用户访问类型分类列表页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 20:25:12
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMasterVisit) ListVisitCategory(c *gin.Context) {
	// 渲染用户访问类型分类列表页面
	system.Render(c, "admin/system_master_visit/list.html", pongo2.Context{
		"title": "访问类型分类列表",
	})
}

// AjaxListVisitCategory
//
// @Title AjaxListVisitCategory
// @Description: Ajax获取用户访问类型分类列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 22:12:38
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMasterVisit) AjaxListVisitCategory(c *gin.Context) {
	start, err := utils.ToInt(c.DefaultQuery("start", "0"))
	if err != nil {
		start = 0
	}

	length, err := utils.ToInt(c.DefaultQuery("length", "10"))
	if err != nil {
		length = 10
	}

	output := services.NewUcSystemMasterVisit().ShowVisitCategory(dto.ListSystemMasterVisitCategoryInput{
		Order:  "id desc",
		Start:  start,
		Length: length,
	})

	if output.Code == 1 {
		system.EmptyJSON(c)
		return
	}

	data := output.Data.(dto.ListSystemMasterVisitCategoryOutput)
	c.JSON(http.StatusOK, gin.H{
		"draw":            c.Query("draw"),
		"data":            data.List,
		"recordsTotal":    data.Total,
		"recordsFiltered": data.Total,
	})
}

// EditVisitCategory
//
// @Title EditVisitCategory
// @Description: 渲染用户访问类型分类编辑页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-26 14:41:49
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMasterVisit) EditVisitCategory(c *gin.Context) {
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

	output := services.NewUcSystemMasterVisit().GetVisitCategoryByID(uint32ID)
	visit, ok := output.Data.(*model.UcSystemMasterVisitCategory)
	if !ok {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	// 渲染编辑用户访问类型分类页面
	system.Render(c, "admin/system_master_visit/edit.html", pongo2.Context{
		"title": "编辑用户访问类型分类",
		"id":    uint32ID,
		"visit": visit,
	})
	return
}

// AjaxEditVisitCategory
//
// @Title AjaxEditVisitCategory
// @Description: Ajax处理编辑用户访问类型分类请求
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-26 14:59:50
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMasterVisit) AjaxEditVisitCategory(c *gin.Context) {
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

	var form validators.AdminVisitCategory
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

	// 根据 ID 更新访问类型分类
	output := services.NewUcSystemMasterVisit().ChangeVisitCategoryByID(uint32ID, dto.SystemMasterVisitCategoryInput{
		Title: form.Title,
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
