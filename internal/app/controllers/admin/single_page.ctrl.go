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
	"reflect"
)

// cSinglePage
//
// @Description: 单页信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-31 23:15:13
type cSinglePage struct {
	pageNotFound string
}

// SinglePage 初始化单页信息控制器
var SinglePage = cSinglePage{
	pageNotFound: paths.AdminRoot + paths.Admin404,
}

// CreateSinglePage
//
// @Title CreateSinglePage
// @Description: 渲染新建单页页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-30 00:25:30
// @receiver ctrl
// @param c
func (ctrl *cSinglePage) CreateSinglePage(c *gin.Context) {
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

	system.Render(c, "admin/single_page/create.html", pongo2.Context{
		"title":   "新建单页信息",
		"success": success,
		"fail":    fail,
		"err_msg": errMsg,
	})
	return
}

// HandleCreateSinglePage
//
// @Title HandleCreateSinglePage
// @Description: 处理新建单页页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-30 00:26:19
// @receiver ctrl
// @param c
func (ctrl *cSinglePage) HandleCreateSinglePage(c *gin.Context) {
	var form validators.SinglePageForm
	if err := c.ShouldBind(&form); err != nil {
		// 获取验证错误信息
		errMsg := system.GetValidationErrorMsg(err, form)
		// 将错误信息存储到会话中
		errFlash := system.AddDataToFlash(c, errMsg, "err_msg")
		if errFlash != nil {
			system.RedirectGet(c, ctrl.pageNotFound)
			return
		}

		// 写入表单提交信息
		v := reflect.ValueOf(form).Elem()
		t := v.Type()
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldName := field.Tag.Get("form")
			fieldValue := v.Field(i).String()
			system.SetOldInput(c, fieldName, fieldValue)
		}
		//system.SetOldInput(c, "title", form.Title)
		//system.SetOldInput(c, "description", form.Description)
		//system.SetOldInput(c, "keyword", form.Keyword)
		//system.SetOldInput(c, "content", form.Content)
		//system.SetOldInput(c, "thumbnail", form.Thumbnail)
		//system.SetOldInput(c, "click", form.Click)
		//system.SetOldInput(c, "status", form.Status)
	} else {
		click, errUInt32 := utils.ToUint32(form.Click)
		if errUInt32 != nil {
			click = 0
		}

		status, errUInt32 := utils.ToUint32(form.Status)
		if errUInt32 != nil {
			status = 0
		}

		output := services.NewSinglePageService().CreateSinglePage(dto.SinglePageInput{
			Title:       form.Title,
			Description: form.Description,
			Keyword:     form.Keyword,
			Content:     form.Content,
			Thumbnail:   form.Thumbnail,
			Click:       click,
			Status:      status,
		})

		if output.Code == 1 {
			system.AddFlashData(c, output.Message, "fail")
		} else {
			system.AddFlashData(c, "添加单页信息成功", "success")
		}
	}

	_ = services.NewUcSystemMasterVisitService().WriteSystemMasterVisitorLogs(c, 1, 5, 0, "添加单页信息")
	system.RedirectGet(c, paths.AdminRoot+paths.AdminCreateSinglePage)
	return
}

// ListSinglePage
//
// @Title ListSinglePage
// @Description: 渲染单页信息列表页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-02 00:11:03
// @receiver ctrl
// @param c
func (ctrl *cSinglePage) ListSinglePage(c *gin.Context) {
	// 渲染单页信息列表页面
	system.Render(c, "admin/single_page/list.html", pongo2.Context{
		"title": "单页信息列表",
	})
}

// AjaxListSinglePage
//
// @Title AjaxListSinglePage
// @Description: Ajax获取单页信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-02 00:11:22
// @receiver ctrl
// @param c
func (ctrl *cSinglePage) AjaxListSinglePage(c *gin.Context) {
	start, err := utils.ToInt(c.DefaultQuery("start", "0"))
	if err != nil {
		start = 0
	}

	length, err := utils.ToInt(c.DefaultQuery("length", "10"))
	if err != nil {
		length = 10
	}

	output := services.NewSinglePageService().ShowSinglePage(dto.ListSinglePageInput{
		Order:  "id desc",
		Start:  start,
		Length: length,
	})

	if output.Code == 1 {
		system.EmptyJSON(c)
		return
	}

	data := output.Data.(dto.ListSinglePageOutput)
	c.JSON(http.StatusOK, gin.H{
		"draw":            c.Query("draw"),
		"data":            data.List,
		"recordsTotal":    data.Total,
		"recordsFiltered": data.Total,
	})
}

// EditSinglePage
//
// @Title EditSinglePage
// @Description: 渲染管单页信息页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-03 16:41:35
// @receiver ctrl
// @param c
func (ctrl *cSinglePage) EditSinglePage(c *gin.Context) {
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

	output := services.NewSinglePageService().GetSinglePageByID(uint32ID)
	if output.Code == 1 {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	singlePage, ok := output.Data.(*model.SinglePage)
	if !ok {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	// 渲染单页信息列表页面
	system.Render(c, "admin/single_page/edit.html", pongo2.Context{
		"title":      "单页信息编辑",
		"singlePage": singlePage,
		"id":         uint32ID,
	})
}

// HandleAjaxEditSinglePage
//
// @Title HandleAjaxEditSinglePage
// @Description: Ajax处理编辑管单页信息请求
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-04 23:45:49
// @receiver ctrl
// @param c
func (ctrl *cSinglePage) HandleAjaxEditSinglePage(c *gin.Context) {
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

	var form validators.SinglePageForm
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

	click, errUInt32 := utils.ToUint32(form.Click)
	if errUInt32 != nil {
		click = 0
	}

	status, errUInt32 := utils.ToUint32(form.Status)
	if errUInt32 != nil {
		status = 0
	}

	// 根据 ID 更新单页信息
	output := services.NewSinglePageService().ChangeSinglePageByID(uint32ID, dto.SinglePageInput{
		Title:       form.Title,
		Description: form.Description,
		Keyword:     form.Keyword,
		Content:     form.Content,
		Thumbnail:   form.Thumbnail,
		Click:       click,
		Status:      status,
	})
	if output.Code == 1 {
		c.JSON(http.StatusOK, &system.SysResponse{
			Code:    1,
			Message: output.Message,
			Data:    nil,
		})
		return
	}

	_ = services.NewUcSystemMasterVisitService().WriteSystemMasterVisitorLogs(c, 1, 5, 0, "编辑单页信息")
	// 更新成功后，可以跳转到用户角色列表页面或显示成功信息
	c.JSON(http.StatusOK, &system.SysResponse{
		Code:    0,
		Message: "Success",
		Data:    nil,
	})
	return
}

func (ctrl *cSinglePage) HandleAjaxEditStatusSinglePage(c *gin.Context) {
	id := c.PostForm("id")
	uint32ID, err := utils.ToUint32(id)
	if err != nil {
		c.JSON(http.StatusOK, &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	status := c.PostForm("status")
	uint32Status, err := utils.ToUint32(status)
	if err != nil {
		c.JSON(http.StatusOK, &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// 根据 ID 更新单页信息
	output := services.NewSinglePageService().ChangeSinglePageByID(uint32ID, dto.SinglePageInput{
		Status: uint32Status,
	})
	if output.Code == 1 {
		c.JSON(http.StatusOK, &system.SysResponse{
			Code:    1,
			Message: output.Message,
			Data:    nil,
		})
		return
	}

	_ = services.NewUcSystemMasterVisitService().WriteSystemMasterVisitorLogs(c, 1, 5, 0, "编辑单页信息发布状态")
	// 更新成功后，可以跳转到用户角色列表页面或显示成功信息
	c.JSON(http.StatusOK, &system.SysResponse{
		Code:    0,
		Message: "Success",
		Data:    nil,
	})
	return
}

func (ctrl *cSinglePage) DeleteSinglePage(c *gin.Context) {

}
