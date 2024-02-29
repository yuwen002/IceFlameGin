package system

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) (string, map[string]string, error) {
	masterID, exists := c.Get("master_id")
	if !exists {
		return "", nil, errors.New("用户ID未在上下文中找到")
	}

	userInfo, exists := c.Get("master_info")
	if !exists {
		return "", nil, errors.New("用户信息未在上下文中找到")
	}

	return masterID.(string), userInfo.(map[string]string), nil
}
