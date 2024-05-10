package main

import (
	"github.com/gin-gonic/gin"
	"ice_flame_gin/config/autoload"
	"ice_flame_gin/routers"
)

func main() {
	r := gin.Default()
	autoload.ValidatorLoader()
	// 注册使用组件
	autoload.UseLoader(r)
	// 注册模版文件
	autoload.LoadTemplates(r)
	// 注册路由
	routers.RegisterRouters(r)

	// 监听端口默认为8000
	r.Run(":8001")
}
