package system

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// GetMasterInfo
//
// @Title GetMasterInfo
// @Description: 获取中间件用户基本信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-01 11:21:13
// @param c
// @return string
// @return map[string]string
// @return error
func GetMasterInfo(c *gin.Context) (uint32, map[string]interface{}, error) {
	masterID, exists := c.Get("master_id")
	if !exists {
		return 0, nil, errors.New("用户ID未在上下文中找到")
	}

	userInfo, exists := c.Get("master_info")
	if !exists {
		return 0, nil, errors.New("用户信息未在上下文中找到")
	}

	return masterID.(uint32), userInfo.(map[string]interface{}), nil
}

// GetMasterID
//
// @Title GetMasterID
// @Description: 获取管理员ID，account id
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-29 11:20:59
// @param c
// @return uint32
// @return error
func GetMasterID(c *gin.Context) (uint32, error) {
	masterID, exists := c.Get("master_id")
	if !exists {
		return 0, errors.New("用户ID未在上下文中找到")
	}

	return masterID.(uint32), nil
}
