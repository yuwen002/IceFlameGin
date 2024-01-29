package admin

import (
	"github.com/gin-gonic/gin"
	"ice_flame_gin/internal/system"
)

var UcSystemMaster = cUcSystemMaster{}

type cUcSystemMaster struct{}

func (c *cUcSystemMaster) Login(ctx *gin.Context) {
	system.Render(ctx, "admin/login.html", gin.H{"title": "后台登入"})
}
