package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ice_flame_gin/internal/app/constants"
	"ice_flame_gin/internal/pkg/utils"
	"ice_flame_gin/internal/system"
	"ice_flame_gin/routers/paths"
)

func MasterAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := system.GetSession(c, "ice_flame_master")
		fmt.Println(token)
		if token == nil {
			//system.RedirectGet(c, paths.AdminRoot)
		}
		claims, err := utils.ParseToken(token.(string), constants.MasterSecretKey)
		if err != nil {
			system.RedirectGet(c, paths.AdminRoot)
		}

		c.Set("master_id", claims.Id)
		c.Set("user_info", claims.UserInfo)
		c.Next()
	}
}
