package autoload

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/memcached"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"ice_flame_gin/config"
	"ice_flame_gin/internal/pkg/utils"
	"log"
)

func UseLoader(r *gin.Engine) {
	// 加在session
	session := config.GlobalConfig.Session
	secret := []byte(utils.HashWithSHA256(session.Secret))
	switch session.Type {
	case "cookie":
		// 初始化基于 Cookie 存储的会话存储引擎
		store := cookie.NewStore(secret)
		// 将会话存储引擎添加到 Gin 中间件
		r.Use(sessions.Sessions(session.Cookie.Name, store))
	case "redis":
		// 初始化基于 Redis 存储的会话存储引擎
		store, err := redis.NewStoreWithDB(
			session.Redis.PrefixLength,
			session.Redis.Network,
			session.Redis.Address,
			session.Redis.Password,
			session.Redis.DB,
			secret,
		)
		if err != nil {
			log.Fatalf("Redis配置信息出错：%v", err)
		}
		r.Use(sessions.Sessions(session.Redis.Name, store))
	case "memcached":
		// 创建 memcache 客户端与 Memcached 服务器建立连接
		mc := memcache.New(session.Memcached.Address)
		// 初始化基于 Memcached 存储的会话存储引擎
		store := memcached.NewStore(mc, "", secret)
		r.Use(sessions.Sessions(session.Memcached.Name, store))
	}

}
