package autoload

import (
	"github.com/flosch/pongo2/v6"
	"github.com/gin-gonic/gin"
	"ice_flame_gin/internal/pkg/utils"
)

//func LoadTemplates(r *gin.Engine) {
//	// 注册模板函数
//	r.SetFuncMap(template.FuncMap{
//		"GetOldInput": func(c *gin.Context, key string) string {
//			return system.GetOldInput(c, key)
//		},
//	})
//
//	// 加载模板一级目录
//	files1, err := filepath.Glob("./web/templates/**/*.html")
//	if err != nil {
//		panic("文件加载错误")
//	}
//	// 加载模板二级机目录
//	files2, err := filepath.Glob("./web/templates/**/**/*.html")
//	if err != nil {
//		panic("文件加载错误")
//	}
//	// 合并两个列表
//	files := append(files1, files2...)
//
//	r.LoadHTMLFiles(files...)
//	// 配置静态文件路径
//	r.Static("/static", "./web/static")
//}

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
