package admin

import (
	"github.com/flosch/pongo2/v6"
	"github.com/gin-gonic/gin"
	"ice_flame_gin/internal/app/dto"
	"ice_flame_gin/internal/app/services"
	"ice_flame_gin/internal/app/validators"
	"ice_flame_gin/internal/system"
	"ice_flame_gin/routers/paths"
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

func (ctrl *cUcSystemMasterVisitor) CreateVisitorCategory(c *gin.Context) {
	system.Render(c, "admin/system_master_visitor/create.html", pongo2.Context{
		"title": "新建访问类型",
	})
}

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
		services.NewUcSystemMasterVisitor().CreateVisitorCategory(dto.SystemMasterVisitorCategoryInput{
			Title: form.Title,
		})
	}
}
