package system

import (
	"github.com/flosch/pongo2/v6"
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
func Render(c *gin.Context, name string, data pongo2.Context) {
	data["ginContext"] = c
	c.HTML(http.StatusOK, name, data)
}

// EmptyJSON
//
// @Title EmptyJSON
// @Description: 返回空JSON
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-14 10:39:23
// @param c
func EmptyJSON(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
