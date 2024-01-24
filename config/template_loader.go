package config

import "github.com/gin-gonic/gin"

func LoadTemplates(r *gin.Engine) {
	// 加载模板
	r.LoadHTMLGlob("web/templates/**/*")
	// 配置静态文件路径
	r.Static("/static", "./web/static")
}
