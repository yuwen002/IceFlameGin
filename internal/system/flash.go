package system

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
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

// SetStructOldInput
//
// @Title SetStructOldInput
// @Description: struct表单提交函数，写入数据
// @Author liuxingyu
// @Date 2024-02-20 23:57:42
// @param c
// @param s
func SetStructOldInput(c *gin.Context, s interface{}) {
	// 通过反射获取结构体的值
	v := reflect.ValueOf(s)

	// 遍历结构体的字段
	for i := 0; i < v.NumField(); i++ {
		// 获取字段名称对应的表单标签
		fieldName := v.Type().Field(i).Tag.Get("form")

		// 获取字段的值
		fieldValue := v.Field(i).Interface()

		// 将字段值转换为字符串类型
		var valueStr string
		switch fieldValue := fieldValue.(type) {
		case int:
			valueStr = strconv.Itoa(fieldValue)
		case float64:
			valueStr = strconv.FormatFloat(fieldValue, 'f', -1, 64)
		case bool:
			valueStr = strconv.FormatBool(fieldValue)
		case string:
			valueStr = fieldValue
		default:
			// 如果无法转换为字符串，可以根据需要进行适当的处理
			// 这里我们选择跳过该字段
			continue
		}

		// 将字段名和字段值保存到会话中
		SetOldInput(c, fieldName, valueStr)
	}
}
