package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var UcSystemMaster = cUcSystemMaster{}

type cUcSystemMaster struct{}

func (c *cUcSystemMaster) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
