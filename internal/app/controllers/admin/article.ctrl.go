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
		"title": "文章分类息列表",
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
	// 更新成功后，可以跳转到用户角色列表页面或显示成功信息
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
	system.RedirectGet(c, paths.AdminRoot+paths.AdminCreateArticleCategory)
	return
}
