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
	system.Render(c, "admin/single_page/create.html", pongo2.Context{
		"title": "新建单页信息",
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
			system.RedirectGet(c, ctrl.pageNotFound)
			return
		}

		status, errUInt32 := utils.ToUint32(form.Status)
		if errUInt32 != nil {
			system.RedirectGet(c, ctrl.pageNotFound)
			return
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
}

func (ctrl *cSinglePage) ListSinglePage(c *gin.Context) {

}

func (ctrl *cSinglePage) AjaxListSinglePage(c *gin.Context) {

}

func (ctrl *cSinglePage) EditSinglePage(c *gin.Context) {

}

func (ctrl *cSinglePage) HandleEditSinglePage(c *gin.Context) {

}

func (ctrl *cSinglePage) DeleteSinglePage(c *gin.Context) {

}
