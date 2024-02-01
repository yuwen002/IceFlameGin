package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Render
//
// @Title Render
// @Description: 使用模板渲染 HTML 页面
// @Author liuxingyu
// @Date 2024-01-25 00:13:38
// @param c
// @param name
// @param data
func Render(c *gin.Context, name string, data gin.H) {
	c.HTML(http.StatusOK, name, data)
}