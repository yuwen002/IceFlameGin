package system

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

// SetSession
//
// @Title SetSession
// @Description: 存储session值
// @Author liuxingyu
// @Date 2024-02-28 00:13:52
// @param c
// @param key
// @param value
func SetSession(c *gin.Context, key string, value interface{}) {
	session := sessions.Default(c)
	session.Set(key, value)
	err := session.Save()
	if err != nil {
		log.Fatalf("session出错：%v", err)
	}
}

// GetSession
//
// @Title GetSession
// @Description: 获取session值
// @Author liuxingyu
// @Date 2024-02-28 00:14:02
// @param c
// @param key
// @return interface{}
func GetSession(c *gin.Context, key string) interface{} {
	session := sessions.Default(c)
	return session.Get(key)
}
