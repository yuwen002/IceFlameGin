package autoload

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func LoadTemplates(r *gin.Engine) {
	// 加载模板一级目录
	files1, err := filepath.Glob("./web/templates/**/*.html")
	if err != nil {
		panic("文件加载错误")
	}
	// 加载模板二级机目录
	files2, err := filepath.Glob("./web/templates/**/**/*.html")
	if err != nil {
		panic("文件加载错误")
	}
	// 合并两个列表
	files := append(files1, files2...)

	r.LoadHTMLFiles(files...)
	// 配置静态文件路径
	r.Static("/static", "./web/static")
}
