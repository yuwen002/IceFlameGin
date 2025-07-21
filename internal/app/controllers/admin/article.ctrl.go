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
	"strings"
)

// cArticle 文章管理控制器
// @Description: 文章管理相关的控制器，处理文章、分类、频道和标签相关的操作
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-10 00:13:44
type cArticle struct {
	pageNotFound string
}

// Article 文章管理控制器实例
// @Description: 文章管理控制器的全局实例
var Article = &cArticle{
	pageNotFound: paths.AdminRoot + paths.Admin404,
}

// CreateArticleCategory
// @Summary         创建文章分类页面
// @Description     渲染新建文章分类的页面，提供分类创建表单
// @Tags            文章管理
// @Accept          html
// @Produce         html
// @Success         200 {object} pongo2.Context "成功渲染页面"
// @Router          /admin/article-category/create [GET]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2024-04-10 00:17:11
// @receiver        ctrl
// @param           c gin.Context 上下文对象
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
// @Summary         处理文章分类创建请求
// @Description     处理新建文章分类的表单提交，验证数据并保存分类信息
// @Tags            文章管理
// @Accept          json
// @Produce         json
// @Param           form body validators.ArticleCategoryForm true "分类表单数据"
// @Success         200 {object} system.SysResponse "成功响应"
// @Failure         400 {object} system.SysResponse "参数错误"
// @Router          /admin/article-category/create [POST]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2024-04-10 00:18:49
// @receiver        ctrl
// @param           c gin.Context 上下文对象
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
// @Summary         查询二级分类
// @Description     Ajax请求，通过一级分类ID查询对应的二级分类
// @Tags            文章管理
// @Accept          json
// @Produce         json
// @Param           fid formData string true "一级分类ID"
// @Success         200 {object} system.SysResponse "成功响应"
// @Failure         400 {object} system.SysResponse "参数错误"
// @Router          /admin/article-category/list-json [POST]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2024-06-01 00:34:56
// @receiver        ctrl
// @param           c gin.Context 上下文对象
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
// @Summary         文章分类列表
// @Description     渲染文章分类信息列表页面
// @Tags            文章管理
// @Accept          html
// @Produce         html
// @Success         200 {object} pongo2.Context "成功渲染页面"
// @Router          /admin/article-category [GET]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2024-04-14 00:58:17
// @receiver        ctrl
// @param           c gin.Context 上下文对象
func (ctrl *cArticle) ListArticleCategory(c *gin.Context) {
	// 渲染文章分类信息列表页面
	system.Render(c, "admin/article_category/list.html", pongo2.Context{
		"title": "文章分类信息列表",
	})
}

// AjaxListArticleCategory
// @Summary         获取分类列表
// @Description     Ajax请求获取文章分类信息列表
// @Tags            文章管理
// @Accept          json
// @Produce         json
// @Success         200 {object} system.SysResponse "成功响应"
// @Failure         400 {object} system.SysResponse "参数错误"
// @Router          /admin/article-category/list [GET]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2024-04-14 00:57:55
// @receiver        ctrl
// @param           c gin.Context 上下文对象
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
// @Summary         编辑文章分类
// @Description     渲染文章分类信息编辑页面
// @Tags            文章管理
// @Accept          html
// @Produce         html
// @Success         200 {object} pongo2.Context "成功渲染页面"
// @Router          /admin/article-category/edit [GET]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2024-04-14 22:40:14
// @receiver        ctrl
// @param           c gin.Context 上下文对象
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
// @Summary         处理分类编辑
// @Description     Ajax处理文章分类信息编辑请求
// @Tags            文章管理
// @Accept          json
// @Produce         json
// @Param           form body validators.ArticleCategoryForm true "分类表单数据"
// @Success         200 {object} system.SysResponse "编辑成功"
// @Failure         400 {object} system.SysResponse "参数错误"
// @Router          /admin/article-category/edit [POST]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2024-04-14 22:54:49
// @receiver        ctrl
// @param           c gin.Context 上下文对象
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
// @Summary         修改分类状态
// @Description     Ajax处理文章分类状态修改请求
// @Tags            文章管理
// @Accept          json
// @Produce         json
// @Param           id formData string true "分类ID"
// @Param           status formData uint32 true "状态值"
// @Success         200 {object} system.SysResponse "修改成功"
// @Failure         400 {object} system.SysResponse "参数错误"
// @Router          /admin/article-category/status [POST]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2024-04-21 16:53:22
// @receiver        ctrl
// @param           c gin.Context 上下文对象
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
// @Summary         删除文章分类
// @Description     Ajax请求处理，删除指定的文章分类
// @Tags            文章管理
// @Accept          json
// @Produce         json
// @Param           id formData string true "分类ID"
// @Success         200 {object} system.SysResponse "删除成功"
// @Failure         400 {object} system.SysResponse "参数错误"
// @Router          /admin/article-category/delete [POST]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2024-04-21 16:52:13
// @receiver        ctrl
// @param           c gin.Context 上下文对象
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
// @Summary         创建文章频道页面
// @Description     渲染新建文章频道页面
// @Tags            文章管理
// @Accept          html
// @Produce         html
// @Success         200 {object} pongo2.Context "成功渲染页面"
// @Router          /admin/article-channel/create [GET]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2024-04-17 23:34:13
// @receiver        ctrl
// @param           c gin.Context 上下文对象
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
// @Summary         创建文章频道
// @Description     处理新建文章频道的表单提交，验证数据并保存频道信息
// @Tags            文章管理
// @Accept          json
// @Produce         json
// @Param           form body validators.ArticleChannelForm true "频道表单数据"
// @Success         200 {object} system.SysResponse "创建成功"
// @Failure         400 {object} system.SysResponse "参数错误"
// @Router          /admin/article-channel/create [POST]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2024-04-17 23:40:37
// @receiver        ctrl
// @param           c gin.Context 上下文对象
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
// @Summary         频道列表
// @Description     渲染文章频道信息列表页面
// @Tags            文章管理
// @Accept          html
// @Produce         html
// @Success         200 {object} pongo2.Context "成功渲染页面"
// @Router          /admin/article-channel [GET]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2024-04-20 01:03:30
// @receiver        ctrl
// @param           c gin.Context 上下文对象
func (ctrl *cArticle) ListArticleChannel(c *gin.Context) {
	// 渲染文章频道信息列表页面
	system.Render(c, "admin/article_channel/list.html", pongo2.Context{
		"title": "文章频道息列表",
	})
}

// AjaxListArticleChannel
// @Summary         获取频道列表
// @Description     Ajax请求获取文章频道信息列表
// @Tags            文章管理
// @Accept          json
// @Produce         json
// @Success         200 {object} system.SysResponse "成功响应"
// @Failure         400 {object} system.SysResponse "参数错误"
// @Router          /admin/article-channel/list [GET]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2024-04-20 01:03:55
// @receiver        ctrl
// @param           c gin.Context 上下文对象
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
// @Summary         编辑频道
// @Description     渲染文章频道信息编辑页面
// @Tags            文章管理
// @Accept          html
// @Produce         html
// @Success         200 {object} pongo2.Context "成功渲染页面"
// @Router          /admin/article-channel/edit [GET]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2024-04-21 21:34:28
// @receiver        ctrl
// @param           c gin.Context 上下文对象
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
// @Summary         处理频道编辑
// @Description     Ajax处理文章频道信息编辑请求
// @Tags            文章管理
// @Accept          json
// @Produce         json
// @Param           form body validators.ArticleChannelForm true "频道表单数据"
// @Success         200 {object} system.SysResponse "编辑成功"
// @Failure         400 {object} system.SysResponse "参数错误"
// @Router          /admin/article-channel/edit [POST]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2024-04-21 21:35:32
// @receiver        ctrl
// @param           c gin.Context 上下文对象
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
// @Summary         修改频道状态
// @Description     Ajax处理文章频道状态修改请求
// @Tags            文章管理
// @Accept          json
// @Produce         json
// @Param           id formData string true "频道ID"
// @Param           status formData uint32 true "状态值"
// @Success         200 {object} system.SysResponse "修改成功"
// @Failure         400 {object} system.SysResponse "参数错误"
// @Router          /admin/article-channel/status [POST]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2024-04-21 16:54:28
// @receiver        ctrl
// @param           c gin.Context 上下文对象
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
// @Summary         删除频道
// @Description     Ajax请求处理，删除指定的文章频道
// @Tags            文章管理
// @Accept          json
// @Produce         json
// @Param           id formData string true "频道ID"
// @Success         200 {object} system.SysResponse "删除成功"
// @Failure         400 {object} system.SysResponse "参数错误"
// @Router          /admin/article-channel/delete [POST]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2024-04-21 21:35:51
// @receiver        ctrl
// @param           c gin.Context 上下文对象
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
// @Summary         创建文章标签页面
// @Description     渲染新建文章标签页面
// @Tags            文章管理
// @Accept          html
// @Produce         html
// @Success         200 {object} pongo2.Context "成功渲染页面"
// @Router          /admin/article-tag/create [GET]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2024-04-22 23:58:01
// @receiver        ctrl
// @param           c gin.Context 上下文对象
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
// @Summary         创建文章标签
// @Description     处理新建文章标签的表单提交
// @Tags            文章管理
// @Accept          json
// @Produce         json
// @Param           form body validators.ArticleTagForm true "标签表单数据"
// @Success         200 {object} system.SysResponse "创建成功"
// @Failure         400 {object} system.SysResponse "参数错误"
// @Router          /admin/article-tag/create [POST]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2024-04-28 23:12:18
// @receiver        ctrl
// @param           c gin.Context 上下文对象
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
// @Summary         标签列表
// @Description     渲染文章标签信息列表页面
// @Tags            文章管理
// @Accept          html
// @Produce         html
// @Success         200 {object} pongo2.Context "成功渲染页面"
// @Router          /admin/article-tag [GET]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2024-04-30 23:57:56
// @receiver        ctrl
// @param           c gin.Context 上下文对象
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
// @Summary         编辑标签
// @Description     渲染文章标签信息编辑页面
// @Tags            文章管理
// @Accept          html
// @Produce         html
// @Success         200 {object} pongo2.Context "成功渲染页面"
// @Router          /admin/article-tag/edit [GET]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2024-05-02 00:05:18
// @receiver        ctrl
// @param           c gin.Context 上下文对象
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
// @Summary         修改标签状态
// @Description     Ajax处理文章标签状态修改请求
// @Tags            文章管理
// @Accept          json
// @Produce         json
// @Param           id formData string true "标签ID"
// @Param           status formData uint32 true "状态值"
// @Success         200 {object} system.SysResponse "修改成功"
// @Failure         400 {object} system.SysResponse "参数错误"
// @Router          /admin/article-tag/status [POST]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2024-05-11 10:45:55
// @receiver        ctrl
// @param           c gin.Context 上下文对象
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
// @Summary         删除标签
// @Description     Ajax请求处理，删除指定的文章标签
// @Tags            文章管理
// @Accept          json
// @Produce         json
// @Param           id formData string true "标签ID"
// @Success         200 {object} system.SysResponse "删除成功"
// @Failure         400 {object} system.SysResponse "参数错误"
// @Router          /admin/article-tag/delete [POST]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2024-05-11 10:47:26
// @receiver        ctrl
// @param           c gin.Context 上下文对象
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
	return
}

// HandelCreateArticle
// @Summary         创建文章
// @Description     处理新建文章的表单提交
// @Tags            文章管理
// @Accept          json
// @Produce         json
// @Param           form body validators.ArticleForm true "文章表单数据"
// @Success         200 {object} system.SysResponse "创建成功"
// @Failure         400 {object} system.SysResponse "参数错误"
// @Router          /admin/article/create [POST]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2024-06-25 11:04:46
// @receiver        ctrl
// @param           c gin.Context 上下文对象
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

	// 处理 tags 字段，将 []string 转为 "1,2,3" 格式
	tagIDs := ""
	if len(form.Tags) > 0 {
		tagIDs = strings.Join(form.Tags, ",")
	}

	input := dto.ArticleInput{
		CategoryId:  uint32(form.SecondCategory),
		ChannelId:   uint32(form.ChannelID),
		Title:       form.Title,
		Keywords:    form.Keyword,
		Description: form.Description,
		Content:     form.Content,
		Link:        form.Link,
		Author:      form.Author,
		Tags:        tagIDs,
		PubDate:     form.PubDate,
		Thumbnail:   form.Thumbnail,
		Summary:     form.Summary,
		Status:      form.Status,
		Click:       form.Click,
	}

	output := services.NewArticleService().CreateArticle(&input)
	if output.Code == 1 {
		system.AddFlashData(c, output.Message, "fail")
	} else {
		system.AddFlashData(c, "添加文章信息成功", "success")
	}

	system.RedirectGet(c, paths.AdminRoot+paths.AdminCreateArticle)
	return
}

// ListArticle
// @Summary         文章列表
// @Description     渲染文章信息列表页面
// @Tags            文章管理
// @Accept          html
// @Produce         html
// @Success         200 {object} pongo2.Context "成功渲染页面"
// @Router          /admin/article [GET]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2025-06-16 09:20:00
// @receiver        ctrl
// @param           c gin.Context 上下文对象
func (ctrl *cArticle) ListArticle(c *gin.Context) {
	system.Render(c, "admin/article/list.html", pongo2.Context{
		"title": "文章信息列表",
	})
}

// AjaxListArticle
// @Summary         获取文章列表
// @Description     Ajax请求获取文章信息列表
// @Tags            文章管理
// @Accept          json
// @Produce         json
// @Success         200 {object} system.SysResponse "成功响应"
// @Failure         400 {object} system.SysResponse "参数错误"
// @Router          /admin/article/list [GET]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2025-06-16 09:20:00
// @receiver        ctrl
// @param           c gin.Context 上下文对象
func (ctrl *cArticle) AjaxListArticle(c *gin.Context) {
	start, err := utils.ToInt(c.DefaultQuery("start", "0"))
	if err != nil {
		start = 0
	}

	length, err := utils.ToInt(c.DefaultQuery("length", "10"))
	if err != nil {
		length = 10
	}

	output := services.NewArticleService().ShowArticle(dto.ListArticleInput{
		Order:  "id desc",
		Start:  start,
		Length: length,
	})

	if output.Code == 1 {
		system.EmptyJSON(c)
		return
	}

	data := output.Data.(dto.ListArticleOutput)
	c.JSON(http.StatusOK, gin.H{
		"draw":            c.Query("draw"),
		"data":            data.List,
		"recordsTotal":    data.Total,
		"recordsFiltered": data.Total,
	})
}

// EditArticle
// @Summary         编辑文章
// @Description     渲染文章信息编辑页面
// @Tags            文章管理
// @Accept          html
// @Produce         html
// @Param           id query string true "文章ID"
// @Success         200 {object} pongo2.Context "成功渲染页面"
// @Failure         400 {object} system.SysResponse "参数错误"
// @Router          /admin/article/edit [GET]
// @Security        JWT
// @Author          liuxingyu <yuwen002@163.com>
// @Date            2024-06-25 11:04:46
// @receiver        ctrl
// @param           c gin.Context 上下文对象
func (ctrl *cArticle) EditArticle(c *gin.Context) {
	id, err := utils.ToUint32(c.Query("id"))
	if err != nil {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	output := services.NewArticleService().GetArticleByID(id)
	if output.Code == 1 {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	article, ok := output.Data.(*model.Article)
	if !ok {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	// 分类信息
	output = services.NewArticleService().ShowArticleCategoryByFID(0)
	if output.Code == 1 {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	categorySelect, ok := output.Data.([]*dto.SelectOptionOutput)
	if !ok {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	// 一级分类选择信息
	output = services.NewArticleService().GetArticleCategoryByID(article.CategoryID)
	if output.Code == 1 {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}

	categoryCurrent, ok := output.Data.(*model.ArticleCategory)
	if !ok {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}
	// 二级分类选择信息
	output = services.NewArticleService().ShowArticleCategoryByFID(categoryCurrent.Fid)
	if output.Code == 1 {
		system.RedirectGet(c, ctrl.pageNotFound)
		return
	}
	subCategory, ok := output.Data.([]*dto.SelectOptionOutput)
	fmt.Println("formatfff", subCategory)
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

	// 渲染文章标签信息列表页面
	system.Render(c, "admin/article/edit.html", pongo2.Context{
		"title":            "编辑文章信息",
		"article":          article,
		"channel_select":   channelSelect,
		"category_select":  categorySelect,
		"category_current": categoryCurrent,
		"sub_category":     subCategory,
		"tag_select":       tagSelect,
	})
}
func (ctrl *cArticle) HandleAjaxEditArticle(c *gin.Context) {}
