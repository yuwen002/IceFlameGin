package system

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AddFlashData
//
// @Title AddFlashData
// @Description: 写入闪存信息
// @Author liuxingyu
// @Date 2024-02-13 23:25:30
// @param c
// @param value
// @param key
func AddFlashData(c *gin.Context, value interface{}, key string) {
	session := sessions.Default(c)
	session.AddFlash(value, key)
	session.Save()
}

// GetFlashedData
//
// @Title GetFlashedData
// @Description: 读取闪存信息
// @Author liuxingyu
// @Date 2024-02-13 23:25:16
// @param c
// @param key
// @return interface{}
func GetFlashedData(c *gin.Context, key string) interface{} {
	session := sessions.Default(c)
	flashes := session.Flashes(key)
	session.Save()

	if len(flashes) > 0 {
		return flashes[0]
	}

	return nil
}

// AddDataToFlash
//
// @Title AddDataToFlash
// @Description: struct或map写入闪存信息
// @Author liuxingyu
// @Date 2024-02-13 23:40:11
// @param c
// @param data
// @param key
// @return error
func AddDataToFlash(c *gin.Context, data interface{}, key string) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	session := sessions.Default(c)
	session.AddFlash(string(jsonData), key)
	err = session.Save()
	return err
}

// GetDataFromFlash
//
// @Title GetDataFromFlash
// @Description: struct或map读取闪存信息
// @Author liuxingyu
// @Date 2024-02-13 23:40:53
// @param c
// @param key
// @param out
// @return error
func GetDataFromFlash(c *gin.Context, key string, out interface{}) error {
	session := sessions.Default(c)
	flashes := session.Flashes(key)
	if len(flashes) == 0 {
		return nil
	}
	flashData, ok := flashes[0].(string)
	if !ok {
		return nil
	}
	err := json.Unmarshal([]byte(flashData), out)
	if err != nil {
		return err
	}
	session.Save()
	return nil
}

// SetOldInput
//
// @Title SetOldInput
// @Description: 表单提交函数，写入数据
// @Author liuxingyu
// @Date 2024-02-20 01:28:08
// @param c
// @param key
// @param value
func SetOldInput(c *gin.Context, key string, value string) {
	// 获取会话
	session := sessions.Default(c)

	// 设置旧输入值到会话中
	session.Set(key, value)
	_ = session.Save()
}

// GetOldInput
//
// @Title GetOldInput
// @Description: 表单取出函数，取出数据
// @Author liuxingyu
// @Date 2024-02-20 01:29:00
// @param c
// @param key
// @return string
func GetOldInput(c *gin.Context, key string) string {
	// 获取会话
	session := sessions.Default(c)

	// 检查会话中是否存在旧输入数据
	if val := session.Get(key); val != nil {
		// 如果存在，将旧输入值作为字符串返回
		if strVal, ok := val.(string); ok {
			return strVal
		}
	}

	return ""
}
