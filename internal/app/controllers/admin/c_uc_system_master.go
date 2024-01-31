package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"ice_flame_gin/internal/app/validators"
	"ice_flame_gin/internal/system"
	"net/http"
)

var UcSystemMaster = cUcSystemMaster{}

type cUcSystemMaster struct{}

func (c *cUcSystemMaster) Login(ctx *gin.Context) {
	system.Render(ctx, "admin/login.html", gin.H{"title": "后台登入"})
}

func (c *cUcSystemMaster) HandleLogin(ctx *gin.Context) {
	var form validators.AdminLoginForm
	if err := ctx.ShouldBind(&form); err != nil {
		// 检查验证错误类型
		if errs, ok := err.(validator.ValidationErrors); ok {
			type ErrorResponse struct {
				Field   string `json:"field"`
				Message string `json:"message"`
			}

			for _, e := range errs {
				// 获取字段名和对应的错误消息
				field := e.Field()
				msg := e.Tag()

				// 返回自定义错误响应
				ctx.JSON(http.StatusBadRequest, ErrorResponse{
					Field:   field,
					Message: msg,
				})
				return
			}
		}
		fmt.Println(err.Error())
		system.Render(ctx, "admin/login.html", gin.H{
			"title": "后台登入",
			"error": err.Error(),
		})
		return
	}
}
