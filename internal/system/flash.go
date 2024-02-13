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
