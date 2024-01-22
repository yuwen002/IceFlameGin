package main

import "ice_flame_gin/routers"

func main() {
	r := routers.RegisterRouters()
	//监听端口默认为8080
	r.Run(":8000")
}
