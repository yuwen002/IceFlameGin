package admin

import (
	"github.com/flosch/pongo2/v6"
	"github.com/gin-gonic/gin"
	"ice_flame_gin/internal/app/dto"
	"ice_flame_gin/internal/app/services"
	"ice_flame_gin/internal/app/validators"
	"ice_flame_gin/internal/pkg/utils"
	"ice_flame_gin/internal/system"
	"ice_flame_gin/routers/paths"
	"net/http"
)

// cUcSystemMasterVisitor
// @Description: 用户访问类型
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 15:09:53
type cUcSystemMasterVisitor struct {
	pageNotFound string
}

// UcSystemMasterVisitor 初始化
var UcSystemMasterVisitor = cUcSystemMasterVisitor{
	pageNotFound: paths.AdminRoot + paths.Admin404,
}

// CreateVisitorCategory
//
// @Title CreateVisitorCategory
// @Description: 渲染新建用户访问类型分类
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 20:22:03
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMasterVisitor) CreateVisitorCategory(c *gin.Context) {
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

	system.Render(c, "admin/system_master_visitor/create.html", pongo2.Context{
		"title":   "新建访问类型",
		"success": success,
		"fail":    fail,
		"err_msg": errMsg,
	})
}

// HandleCreateVisitorCategory
//
// @Title HandleCreateVisitorCategory
// @Description: 处理新建用户访问类型分类请求
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 20:22:33
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMasterVisitor) HandleCreateVisitorCategory(c *gin.Context) {
	var form validators.AdminVisitorCategory

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
		output := services.NewUcSystemMasterVisitor().CreateVisitorCategory(dto.SystemMasterVisitorCategoryInput{
			Title: form.Title,
		})

		if output.Code == 1 {
			system.AddFlashData(c, output.Message, "fail")
		} else {
			system.AddFlashData(c, "添加访问类型分类信息成功", "success")
		}
	}

	system.RedirectGet(c, paths.AdminRoot+paths.AdminCreateVisitorCategory)
}

// ListVisitorCategory
//
// @Title ListVisitorCategory
// @Description: 渲染用户访问类型分类列表页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 20:25:12
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMasterVisitor) ListVisitorCategory(c *gin.Context) {
	// 渲染用户访问类型分类列表页面
	system.Render(c, "admin/system_master_visitor/list.html", pongo2.Context{
		"title": "访问类型分类列表",
	})
}

// AjaxListVisitorCategory
//
// @Title AjaxListVisitorCategory
// @Description: Ajax获取用户访问类型分类列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 22:12:38
// @receiver ctrl
// @param c
func (ctrl *cUcSystemMasterVisitor) AjaxListVisitorCategory(c *gin.Context) {
	start, err := utils.ToInt(c.DefaultQuery("start", "0"))
	if err != nil {
		start = 0
	}

	length, err := utils.ToInt(c.DefaultQuery("length", "10"))
	if err != nil {
		length = 10
	}

	output := services.NewUcSystemMasterVisitor().ShowVisitorCategory(dto.ListSystemMasterVisitorCategoryInput{
		Order:  "id desc",
		Start:  start,
		Length: length,
	})

	if output.Code == 1 {
		system.EmptyJSON(c)
		return
	}

	data := output.Data.(dto.ListSystemMasterVisitorCategoryOutput)
	c.JSON(http.StatusOK, gin.H{
		"draw":            c.Query("draw"),
		"data":            data.List,
		"recordsTotal":    data.Total,
		"recordsFiltered": data.Total,
	})
}
