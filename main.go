package main

import (
	"ice_flame_gin/config"
	"ice_flame_gin/routers"
)

func main() {
	r := routers.RegisterRouters()
	config.LoadTemplates(r)
	//监听端口默认为8080
	r.Run(":8000")
}
