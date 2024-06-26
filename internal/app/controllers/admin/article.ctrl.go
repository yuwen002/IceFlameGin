package admin

import (
	"fmt"
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

// cArticle
//
// @Description: 文章管理
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-10 00:13:44
type cArticle struct {
	pageNotFound string
}

// Article 初始化文章管理控制器
var Article = &cArticle{
	pageNotFound: paths.AdminRoot + paths.Admin404,
}

// CreateArticleCategory
//
// @Title CreateArticleCategory
// @Description: 渲染新建文章分类页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-10 00:17:11
// @receiver ctrl
// @param c
func (ctrl *cArticle) CreateArticleCategory(c *gin.Context) {
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

	output := services.NewArticleService().ShowArticleCategoryByFID(0)
	if output.Code == 1 {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	fidSelect, ok := output.Data.([]*dto.SelectOptionOutput)
	if !ok {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	system.Render(c, "admin/article_category/create.html", pongo2.Context{
		"title":      "新建文章分类信息",
		"success":    success,
		"fail":       fail,
		"err_msg":    errMsg,
		"fid_select": fidSelect,
	})
	return
}

// HandleCreateArticleCategory
//
// @Title HandleCreateArticleCategory
// @Description: 处理新建文章分类页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-10 00:18:49
// @receiver ctrl
// @param c
func (ctrl *cArticle) HandleCreateArticleCategory(c *gin.Context) {
	var form validators.ArticleCategoryForm
	if err := c.ShouldBind(&form); err != nil {
		// 获取验证错误信息
		errMsg := system.GetValidationErrors(err, form)
		// 将错误信息存储到会话中
		errFlash := system.AddDataToFlash(c, errMsg, "err_msg")
		if errFlash != nil {
			system.RedirectGet(c, ctrl.pageNotFound)
			return
		}

		// 写入表单提交信息
		v := reflect.ValueOf(&form).Elem()
		t := v.Type()
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldName := field.Tag.Get("form")
			fieldValue := v.Field(i).String()
			system.SetOldInput(c, fieldName, fieldValue)
		}
		system.RedirectGet(c, paths.AdminRoot+paths.AdminCreateArticleCategory)
		return
	}

	fid, errUInt32 := utils.ToUint32(form.Fid)
	if errUInt32 != nil {
		fid = 0
	}

	sort, errUInt32 := utils.ToUint32(form.Sort)
	if errUInt32 != nil {
		sort = 0
	}

	status, errUInt32 := utils.ToUint32(form.Status)
	if errUInt32 != nil {
		status = 0
	}

	output := services.NewArticleService().CreateArticleCategory(&dto.ArticleCategoryInput{
		Fid:    fid,
		Name:   form.Name,
		Remark: form.Remark,
		Sort:   sort,
		Status: status,
	})

	if output.Code == 1 {
		system.AddFlashData(c, output.Message, "fail")
	} else {
		system.AddFlashData(c, "添加文章分类信息成功", "success")
	}

	_ = services.NewUcSystemMasterVisitService().WriteSystemMasterVisitorLogs(c, 1, 5, 0, "添加文章分类")
	system.RedirectGet(c, paths.AdminRoot+paths.AdminCreateArticleCategory)
	return
}

// AjaxListArticleCategoryJson
//
// @Title AjaxListArticleCategoryJson
// @Description: Ajax 通过一级分类 ID 查询对应的二级分类
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-06-01 00:34:56
// @receiver ctrl
// @param c
func (ctrl *cArticle) AjaxListArticleCategoryJson(c *gin.Context) {
	firstCategory := c.Query("category_id")
	toFirstCategory, err := utils.ToUint32(firstCategory)
	if err != nil {
		system.EmptyJSON(c)
		return
	}
	// 查询分类信息
	output := services.NewArticleService().ShowArticleCategoryByFID(toFirstCategory)
	if output.Code == 1 {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	fidSelect, ok := output.Data.([]*dto.SelectOptionOutput)
	if !ok {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	c.JSON(http.StatusOK, &system.SysResponse{
		Code:    0,
		Message: "success",
		Data:    fidSelect,
	})
}

// ListArticleCategory
//
// @Title ListArticleCategory
// @Description: 渲染文章分类信息列表页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-14 00:58:17
// @receiver ctrl
// @param c
func (ctrl *cArticle) ListArticleCategory(c *gin.Context) {
	// 渲染文章分类信息列表页面
	system.Render(c, "admin/article_category/list.html", pongo2.Context{
		"title": "文章分类信息列表",
	})
}

// AjaxListArticleCategory
//
// @Title AjaxListArticleCategory
// @Description: Ajax获取文章分类信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-14 00:57:55
// @receiver ctrl
// @param c
func (ctrl *cArticle) AjaxListArticleCategory(c *gin.Context) {
	start, err := utils.ToInt(c.DefaultQuery("start", "0"))
	if err != nil {
		start = 0
	}

	length, err := utils.ToInt(c.DefaultQuery("length", "10"))
	if err != nil {
		length = 10
	}

	output := services.NewArticleService().ShowArticleCategory(dto.ListArticleCategoryInput{
		Order:  "id desc",
		Start:  start,
		Length: length,
	})

	if output.Code == 1 {
		system.EmptyJSON(c)
		return
	}

	data := output.Data.(dto.ListArticleCategoryOutput)
	c.JSON(http.StatusOK, gin.H{
		"draw":            c.Query("draw"),
		"data":            data.List,
		"recordsTotal":    data.Total,
		"recordsFiltered": data.Total,
	})
}

// EditArticleCategory
//
// @Title EditArticleCategory
// @Description: 渲染文章分类信息编辑页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-14 22:40:14
// @receiver ctrl
// @param c
func (ctrl *cArticle) EditArticleCategory(c *gin.Context) {
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

	// 查询分类信息
	output := services.NewArticleService().GetArticleCategoryByID(uint32ID)
	if output.Code == 1 {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	category, ok := output.Data.(*model.ArticleCategory)
	if !ok {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	output = services.NewArticleService().ShowArticleCategoryByFID(0)
	if output.Code == 1 {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	fidSelect, ok := output.Data.([]*dto.SelectOptionOutput)
	if !ok {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	// 渲染文章分类信息编辑页面
	system.Render(c, "admin/article_category/edit.html", pongo2.Context{
		"title":      "文章分类信息编辑",
		"category":   category,
		"fid_select": fidSelect,
		"id":         uint32ID,
	})
}

// HandleAjaxEditArticleCategory
//
// @Title HandleAjaxEditArticleCategory
// @Description: Ajax处理编辑管文章分类信息请求
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-14 22:54:49
// @receiver ctrl
// @param c
func (ctrl *cArticle) HandleAjaxEditArticleCategory(c *gin.Context) {
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

	var form validators.ArticleCategoryForm
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

	fid, errUInt32 := utils.ToUint32(form.Fid)
	if errUInt32 != nil {
		fid = 0
	}

	sort, errUInt32 := utils.ToUint32(form.Sort)
	if errUInt32 != nil {
		sort = 0
	}

	status, errUInt32 := utils.ToUint32(form.Status)
	if errUInt32 != nil {
		status = 0
	}

	// 更新信息
	output := services.NewArticleService().ChangeArticleCategoryByID(uint32ID, dto.ArticleCategoryInput{
		Fid:    fid,
		Name:   form.Name,
		Remark: form.Remark,
		Sort:   sort,
		Status: status,
	})
	if output.Code == 1 {
		c.JSON(http.StatusOK, &system.SysResponse{
			Code:    1,
			Message: output.Message,
			Data:    nil,
		})
		return
	}

	_ = services.NewUcSystemMasterVisitService().WriteSystemMasterVisitorLogs(c, 1, 5, 0, "编辑文章分类信息")
	// 更新成功后，可以跳转到用户角色列表页面或显示成功信息
	c.JSON(http.StatusOK, &system.SysResponse{
		Code:    0,
		Message: "Success",
		Data:    nil,
	})
	return
}

// HandleAjaxEditStatusArticleCategory
//
// @Title HandleAjaxEditStatusArticleCategory
// @Description: Ajax处理编辑管文章分类状态信息请求
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-21 16:53:22
// @receiver ctrl
// @param c
func (ctrl *cArticle) HandleAjaxEditStatusArticleCategory(c *gin.Context) {
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

	// 更新信息
	output := services.NewArticleService().ChangeArticleCategoryByID(uint32ID, dto.ArticleCategoryInput{
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

	_ = services.NewUcSystemMasterVisitService().WriteSystemMasterVisitorLogs(c, 1, 5, 0, "编辑文章分类信息状态")
	// 更新成功后，可以跳转到文章分类列表页面或显示成功信息
	c.JSON(http.StatusOK, &system.SysResponse{
		Code:    0,
		Message: "Success",
		Data:    nil,
	})
	return
}

// HandelAjaxDeleteArticleCategory
//
// @Title HandelAjaxDeleteArticleCategory
// @Description: Ajax处理删除管文章分类信息请求
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-21 16:52:13
// @receiver ctrl
// @param c
func (ctrl *cArticle) HandelAjaxDeleteArticleCategory(c *gin.Context) {
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

	output := services.NewArticleService().DeleteArticleCategoryByID(uint32ID)
	if output.Code == 1 {
		c.JSON(http.StatusOK, &system.SysResponse{
			Code:    1,
			Message: output.Message,
			Data:    nil,
		})
		return
	}

	_ = services.NewUcSystemMasterVisitService().WriteSystemMasterVisitorLogs(c, 1, 5, 0, "删除文章分类信息")
	// 更新成功后，可以跳转到文章分类列表页面或显示成功信息
	c.JSON(http.StatusOK, &system.SysResponse{
		Code:    0,
		Message: "Success",
		Data:    nil,
	})
	return
}

// CreateArticleChannel
//
// @Title CreateArticleChannel
// @Description: 渲染新建文章频道页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-17 23:34:13
// @receiver ctrl
// @param c
func (ctrl *cArticle) CreateArticleChannel(c *gin.Context) {
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

	system.Render(c, "admin/article_channel/create.html", pongo2.Context{
		"title":   "新建文章频道信息",
		"success": success,
		"fail":    fail,
		"err_msg": errMsg,
	})
	return
}

// HandelCreateArticleChannel
//
// @Title HandelCreateArticleChannel
// @Description: 处理新建文章频道页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-17 23:40:37
// @receiver ctrl
// @param c
func (ctrl *cArticle) HandelCreateArticleChannel(c *gin.Context) {
	var form validators.ArticleChannelForm
	if err := c.ShouldBind(&form); err != nil {
		// 获取验证错误信息
		errMsg := system.GetValidationErrors(err, form)
		// 将错误信息存储到会话中
		errFlash := system.AddDataToFlash(c, errMsg, "err_msg")
		if errFlash != nil {
			system.RedirectGet(c, ctrl.pageNotFound)
			return
		}

		// 写入表单提交信息
		v := reflect.ValueOf(&form).Elem()
		t := v.Type()
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldName := field.Tag.Get("form")
			fieldValue := v.Field(i).String()
			system.SetOldInput(c, fieldName, fieldValue)
		}
		system.RedirectGet(c, paths.AdminRoot+paths.AdminCreateArticleChannel)
		return
	}

	sort, errUInt32 := utils.ToUint32(form.Sort)
	if errUInt32 != nil {
		sort = 0
	}

	status, errUInt32 := utils.ToUint32(form.Status)
	if errUInt32 != nil {
		status = 0
	}

	output := services.NewArticleService().CreateArticleChannel(&dto.ArticleChannelInput{
		Name:   form.Name,
		Remark: form.Remark,
		Sort:   sort,
		Status: status,
	})

	if output.Code == 1 {
		system.AddFlashData(c, output.Message, "fail")
	} else {
		system.AddFlashData(c, "添加文章分类频道成功", "success")
	}

	_ = services.NewUcSystemMasterVisitService().WriteSystemMasterVisitorLogs(c, 1, 5, 0, "添加文章频道")
	system.RedirectGet(c, paths.AdminRoot+paths.AdminCreateArticleChannel)
	return
}

// ListArticleChannel
//
// @Title ListArticleChannel
// @Description: 渲染文章频道信息列表页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-20 01:03:30
// @receiver ctrl
// @param c
func (ctrl *cArticle) ListArticleChannel(c *gin.Context) {
	// 渲染文章频道信息列表页面
	system.Render(c, "admin/article_channel/list.html", pongo2.Context{
		"title": "文章频道息列表",
	})
}

// AjaxListArticleChannel
//
// @Title AjaxListArticleChannel
// @Description: Ajax获取文章分类信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-20 01:03:55
// @receiver ctrl
// @param c
func (ctrl *cArticle) AjaxListArticleChannel(c *gin.Context) {
	start, err := utils.ToInt(c.DefaultQuery("start", "0"))
	if err != nil {
		start = 0
	}

	length, err := utils.ToInt(c.DefaultQuery("length", "10"))
	if err != nil {
		length = 10
	}

	output := services.NewArticleService().ShowArticleChannel(dto.ListArticleChannelInput{
		Order:  "id desc",
		Start:  start,
		Length: length,
	})

	if output.Code == 1 {
		system.EmptyJSON(c)
		return
	}

	data := output.Data.(dto.ListArticleChannelOutput)
	c.JSON(http.StatusOK, gin.H{
		"draw":            c.Query("draw"),
		"data":            data.List,
		"recordsTotal":    data.Total,
		"recordsFiltered": data.Total,
	})
}

// EditArticleChannel
//
// @Title EditArticleChannel
// @Description: 渲染编辑文章频道信息页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-21 21:34:28
// @receiver ctrl
// @param c
func (ctrl *cArticle) EditArticleChannel(c *gin.Context) {
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

	// 查询分类信息
	output := services.NewArticleService().GetArticleChannelByID(uint32ID)
	if output.Code == 1 {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	channel, ok := output.Data.(*model.ArticleChannel)
	if !ok {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	// 渲染文章分类信息编辑页面
	system.Render(c, "admin/article_channel/edit.html", pongo2.Context{
		"title":   "文章分类信息编辑",
		"channel": channel,
		"id":      uint32ID,
	})
}

// HandleAjaxEditArticleChannel
//
// @Title HandleAjaxEditArticleChannel
// @Description: Ajax处理编辑管文章频道信息请求
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-21 21:35:32
// @receiver ctrl
// @param c
func (ctrl *cArticle) HandleAjaxEditArticleChannel(c *gin.Context) {
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

	var form validators.ArticleCategoryForm
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
	sort, errUInt32 := utils.ToUint32(form.Sort)
	if errUInt32 != nil {
		sort = 0
	}

	status, errUInt32 := utils.ToUint32(form.Status)
	if errUInt32 != nil {
		status = 0
	}

	// 更新信息
	output := services.NewArticleService().ChangeArticleChannelByID(uint32ID, &dto.ArticleChannelInput{
		Name:   form.Name,
		Remark: form.Remark,
		Sort:   sort,
		Status: status,
	})
	if output.Code == 1 {
		c.JSON(http.StatusOK, &system.SysResponse{
			Code:    1,
			Message: output.Message,
			Data:    nil,
		})
		return
	}

	_ = services.NewUcSystemMasterVisitService().WriteSystemMasterVisitorLogs(c, 1, 5, 0, "编辑文章频道信息")
	// 更新成功后，可以跳转到文章频道列表页面或显示成功信息
	c.JSON(http.StatusOK, &system.SysResponse{
		Code:    0,
		Message: "Success",
		Data:    nil,
	})
	return
}

// HandleAjaxEditStatusArticleChannel
//
// @Title HandleAjaxEditStatusArticleChannel
// @Description: Ajax处理编辑管文章频道状态信息请求
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-21 16:54:28
// @receiver ctrl
// @param c
func (ctrl *cArticle) HandleAjaxEditStatusArticleChannel(c *gin.Context) {
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

	// 更新信息
	output := services.NewArticleService().ChangeArticleChannelByID(uint32ID, &dto.ArticleChannelInput{
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

	_ = services.NewUcSystemMasterVisitService().WriteSystemMasterVisitorLogs(c, 1, 5, 0, "编辑文章频道信息状态")
	// 更新成功后，可以跳转到文章频道列表页面或显示成功信息
	c.JSON(http.StatusOK, &system.SysResponse{
		Code:    0,
		Message: "Success",
		Data:    nil,
	})
	return
}

// HandleAjaxEditDeleteArticleChannel
//
// @Title HandleAjaxEditDeleteArticleChannel
// @Description: Ajax处理删除管文章频道信息请求
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-21 21:35:51
// @receiver ctrl
// @param c
func (ctrl *cArticle) HandleAjaxEditDeleteArticleChannel(c *gin.Context) {
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

	output := services.NewArticleService().DeleteArticleChannelByID(uint32ID)
	if output.Code == 1 {
		c.JSON(http.StatusOK, &system.SysResponse{
			Code:    1,
			Message: output.Message,
			Data:    nil,
		})
		return
	}

	_ = services.NewUcSystemMasterVisitService().WriteSystemMasterVisitorLogs(c, 1, 5, 0, "删除文章频道信息")
	// 更新成功后，可以跳转到文章频道列表页面或显示成功信息
	c.JSON(http.StatusOK, &system.SysResponse{
		Code:    0,
		Message: "Success",
		Data:    nil,
	})
	return
}

// CreateArticleTag
//
// @Title CreateArticleTag
// @Description: 渲染新建文章tag页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-22 23:58:01
// @receiver ctrl
// @param c
func (ctrl *cArticle) CreateArticleTag(c *gin.Context) {
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

	system.Render(c, "admin/article_tag/create.html", pongo2.Context{
		"title":   "新建文章标签信息",
		"success": success,
		"fail":    fail,
		"err_msg": errMsg,
	})
	return
}

// HandelCreateArticleTag
//
// @Title HandelCreateArticleTag
// @Description: 处理新建文章tag页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-28 23:12:18
// @receiver ctrl
// @param c
func (ctrl *cArticle) HandelCreateArticleTag(c *gin.Context) {
	var form validators.ArticleTagForm
	if err := c.ShouldBind(&form); err != nil {
		// 获取验证错误信息
		errMsg := system.GetValidationErrors(err, form)
		// 将错误信息存储到会话中
		errFlash := system.AddDataToFlash(c, errMsg, "err_msg")
		if errFlash != nil {
			system.RedirectGet(c, ctrl.pageNotFound)
			return
		}

		// 写入表单提交信息
		v := reflect.ValueOf(&form).Elem()
		t := v.Type()
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldName := field.Tag.Get("form")
			fieldValue := v.Field(i).String()
			system.SetOldInput(c, fieldName, fieldValue)
		}

		system.RedirectGet(c, paths.AdminRoot+paths.AdminCreateArticleTag)
		return
	}

	sort, errUInt32 := utils.ToUint32(form.Sort)
	if errUInt32 != nil {
		sort = 0
	}

	status, errUInt32 := utils.ToUint32(form.Status)
	if errUInt32 != nil {
		status = 0
	}

	output := services.NewArticleService().CreateArticleTag(&dto.ArticleTagInput{
		Name:   form.Name,
		Remark: form.Remark,
		Sort:   sort,
		Status: status,
	})

	if output.Code == 1 {
		system.AddFlashData(c, output.Message, "fail")
	} else {
		system.AddFlashData(c, "添加文章标签信息成功", "success")
	}

	_ = services.NewUcSystemMasterVisitService().WriteSystemMasterVisitorLogs(c, 1, 5, 0, "添加文章标签")
	system.RedirectGet(c, paths.AdminRoot+paths.AdminCreateArticleTag)
	return
}

// ListArticleTag
//
// @Title ListArticleTag
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-30 23:57:56
// @receiver ctrl
// @param c
func (ctrl *cArticle) ListArticleTag(c *gin.Context) {
	// 渲染文章标签信息列表页面
	system.Render(c, "admin/article_tag/list.html", pongo2.Context{
		"title": "文章标签息信列表",
	})
}

// AjaxListArticleTag
//
// @Title AjaxListArticleTag
// @Description: Ajax获取文章标签信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-05-01 00:00:20
// @receiver ctrl
// @param c
func (ctrl *cArticle) AjaxListArticleTag(c *gin.Context) {
	start, err := utils.ToInt(c.DefaultQuery("start", "0"))
	if err != nil {
		start = 0
	}

	length, err := utils.ToInt(c.DefaultQuery("length", "10"))
	if err != nil {
		length = 10
	}

	output := services.NewArticleService().ShowArticleTag(dto.ListArticleTagInput{
		Order:  "id desc",
		Start:  start,
		Length: length,
	})

	if output.Code == 1 {
		system.EmptyJSON(c)
		return
	}

	data := output.Data.(dto.ListArticleTagOutput)
	c.JSON(http.StatusOK, gin.H{
		"draw":            c.Query("draw"),
		"data":            data.List,
		"recordsTotal":    data.Total,
		"recordsFiltered": data.Total,
	})
}

// EditArticleTag
//
// @Title EditArticleTag
// @Description: 渲染文章标签信息编辑页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-05-02 00:05:18
// @receiver ctrl
// @param c
func (ctrl *cArticle) EditArticleTag(c *gin.Context) {
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

	// 查询分类信息
	output := services.NewArticleService().GetArticleTagByID(uint32ID)
	if output.Code == 1 {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	tag, ok := output.Data.(*model.ArticleTag)
	if !ok {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	// 渲染文章标签信息编辑页面
	system.Render(c, "admin/article_tag/edit.html", pongo2.Context{
		"title": "文章标签信息编辑",
		"tag":   tag,
		"id":    uint32ID,
	})
}

// HandelAjaxEditArticleTag
//
// @Title HandelAjaxEditArticleTag
// @Description: 处理新建文章标签页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-05-02 00:17:50
// @receiver ctrl
// @param c
func (ctrl *cArticle) HandelAjaxEditArticleTag(c *gin.Context) {
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

	var form validators.ArticleTagForm
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
	sort, errUInt32 := utils.ToUint32(form.Sort)
	if errUInt32 != nil {
		sort = 0
	}

	status, errUInt32 := utils.ToUint32(form.Status)
	if errUInt32 != nil {
		status = 1
	}

	// 更新信息
	output := services.NewArticleService().ChangeArticleTagByID(uint32ID, &dto.ArticleTagInput{
		Name:   form.Name,
		Remark: form.Remark,
		Sort:   sort,
		Status: status,
	})
	if output.Code == 1 {
		c.JSON(http.StatusOK, &system.SysResponse{
			Code:    1,
			Message: output.Message,
			Data:    nil,
		})
		return
	}

	_ = services.NewUcSystemMasterVisitService().WriteSystemMasterVisitorLogs(c, 1, 5, 0, "编辑文章标签信息")
	// 更新成功后，可以跳转到文章标签列表页面或显示成功信息
	c.JSON(http.StatusOK, &system.SysResponse{
		Code:    0,
		Message: "Success",
		Data:    nil,
	})
	return
}

// HandelAjaxEditStatusArticleTag
//
// @Title HandelAjaxEditStatusArticleTag
// @Description: Ajax处理文章标签状态
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-05-11 10:45:55
// @receiver ctrl
// @param c
func (ctrl *cArticle) HandelAjaxEditStatusArticleTag(c *gin.Context) {
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

	var form validators.ArticleTagForm
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
	sort, errUInt32 := utils.ToUint32(form.Sort)
	if errUInt32 != nil {
		sort = 0
	}

	status, errUInt32 := utils.ToUint32(form.Status)
	if errUInt32 != nil {
		status = 0
	}

	// 更新信息
	output := services.NewArticleService().ChangeArticleTagByID(uint32ID, &dto.ArticleTagInput{
		Name:   form.Name,
		Remark: form.Remark,
		Sort:   sort,
		Status: status,
	})
	if output.Code == 1 {
		c.JSON(http.StatusOK, &system.SysResponse{
			Code:    1,
			Message: output.Message,
			Data:    nil,
		})
		return
	}

	_ = services.NewUcSystemMasterVisitService().WriteSystemMasterVisitorLogs(c, 1, 5, 0, "编辑文章标签信息")
	// 更新成功后，可以跳转到文章标签列表页面或显示成功信息
	c.JSON(http.StatusOK, &system.SysResponse{
		Code:    0,
		Message: "Success",
		Data:    nil,
	})
	return

}

// HandelAjaxDeleteArticleTag
//
// @Title HandelAjaxDeleteArticleTag
// @Description: Ajax删除文章标签
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-05-11 10:47:26
// @receiver ctrl
// @param c
func (ctrl *cArticle) HandelAjaxDeleteArticleTag(c *gin.Context) {
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

	output := services.NewArticleService().DeleteArticleTagByID(uint32ID)
	if output.Code == 1 {
		c.JSON(http.StatusOK, &system.SysResponse{
			Code:    1,
			Message: output.Message,
			Data:    nil,
		})
		return
	}

	_ = services.NewUcSystemMasterVisitService().WriteSystemMasterVisitorLogs(c, 1, 5, 0, "删除文章标签信息")
	// 更新成功后，可以跳转到文章频道列表页面或显示成功信息
	c.JSON(http.StatusOK, &system.SysResponse{
		Code:    0,
		Message: "Success",
		Data:    nil,
	})
	return
}

// CreateArticle
//
// @Title CreateArticle
// @Description: 新建文章
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-05-22 18:24:55
// @receiver ctrl
// @param c
func (ctrl *cArticle) CreateArticle(c *gin.Context) {
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

	// 分类信息
	output := services.NewArticleService().ShowArticleCategoryByFID(0)
	if output.Code == 1 {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	categorySelect, ok := output.Data.([]*dto.SelectOptionOutput)
	if !ok {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	// 频道信息
	output = services.NewArticleService().GetArticleChannelAll()
	if output.Code == 1 {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}
	channelSelect, ok := output.Data.([]*dto.SelectOptionOutput)
	if output.Code == 1 {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	// 标签信息
	output = services.NewArticleService().GetArticleTagAll()
	if output.Code == 1 {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}
	tagSelect, ok := output.Data.([]*dto.SelectOptionOutput)
	if output.Code == 1 {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	system.Render(c, "admin/article/create.html", pongo2.Context{
		"title":           "新建文章信息",
		"success":         success,
		"fail":            fail,
		"err_msg":         errMsg,
		"channel_select":  channelSelect,
		"category_select": categorySelect,
		"tag_select":      tagSelect,
	})
}

// HandelCreateArticle
//
// @Title HandelCreateArticle
// @Description: 处理新建文章g页面
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-06-25 11:04:46
// @receiver ctrl
// @param c
func (ctrl *cArticle) HandelCreateArticle(c *gin.Context) {
	var form validators.ArticleForm
	if err := c.ShouldBind(&form); err != nil {
		// 获取验证错误信息
		errMsg := system.GetValidationErrors(err, form)
		// 将错误信息存储到会话中
		err := system.AddDataToFlash(c, errMsg, "err_msg")
		if err != nil {
			system.RedirectGet(c, ctrl.pageNotFound)
			return
		}

		// 写入表单提交信息
		v := reflect.ValueOf(&form).Elem()
		t := v.Type()
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldName := field.Tag.Get("form")
			fieldValue := v.Field(i).String()
			system.SetOldInput(c, fieldName, fieldValue)
		}

		system.RedirectGet(c, paths.AdminRoot+paths.AdminCreateArticle)
		return
	}

	fmt.Println(form.Tags)
}
func (ctrl *cArticle) ListArticle(c *gin.Context)           {}
func (ctrl *cArticle) AjaxListArticle(c *gin.Context)       {}
func (ctrl *cArticle) EditArticle(c *gin.Context)           {}
func (ctrl *cArticle) HandleAjaxEditArticle(c *gin.Context) {}
