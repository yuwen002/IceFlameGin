package autoload

import (
	"github.com/flosch/pongo2/v6"
	"github.com/gin-gonic/gin"
	"ice_flame_gin/internal/pkg/utils"
)

// LoadTemplates
//
// @Title LoadTemplates
// @Description: 加载模板函数
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-11 01:19:26
// @param r
func LoadTemplates(r *gin.Engine) {
	// 加载模板
	r.HTMLRender = utils.TemplatePath("web/templates")
	// 注册自定义标签
	pongo2.RegisterFilter("old", utils.OldFilter)
	// 配置静态文件路径
	r.Static("/static", "./web/static")
}
