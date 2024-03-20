package middlewares

import (
	"github.com/gin-gonic/gin"
	"ice_flame_gin/internal/app/constants"
	"ice_flame_gin/internal/pkg/utils"
	"ice_flame_gin/internal/system"
	"ice_flame_gin/routers/paths"
)

func MasterAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置响应头，禁用缓存
		c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

		token := system.GetSession(c, "ice_flame_master")
		if token == nil {
			system.RedirectGet(c, paths.AdminRoot)
			return
		}
		claims, err := utils.ParseToken(token.(string), constants.MasterSecretKey)
		if err != nil {
			system.RedirectGet(c, paths.AdminRoot)
			return
		}

		c.Set("master_id", claims.Id)
		c.Set("master_info", claims.UserInfo)
		c.Next()
	}
}
