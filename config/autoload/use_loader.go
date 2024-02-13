package autoload

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func UseLoader(r *gin.Engine) {
	// 初始化基于 Cookie 存储的会话存储引擎
	store := cookie.NewStore([]byte("secret"))
	// 将会话存储引擎添加到 Gin 中间件
	r.Use(sessions.Sessions("mysession", store))
}
