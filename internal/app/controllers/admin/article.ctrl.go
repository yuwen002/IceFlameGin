package admin

import (
	"github.com/flosch/pongo2/v6"
	"github.com/gin-gonic/gin"
	"ice_flame_gin/internal/app/validators"
	"ice_flame_gin/internal/system"
	"ice_flame_gin/routers/paths"
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

	system.Render(c, "admin/article_category/create.html", pongo2.Context{
		"title":   "新建文章分类信息",
		"success": success,
		"fail":    fail,
		"err_msg": errMsg,
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
}
