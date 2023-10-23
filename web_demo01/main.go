package main

import (
	"github.com/gin-gonic/gin"
	"lin.com/web_demo01/routers"
)

func main() {
	r := gin.Default()
	// 设置受信任的代理
	r.SetTrustedProxies([]string{"127.0.0.1"})
	// 用于配置全局中间件
	// r.Use()

	routers.V1Router(r)
	routers.V2Router(r)
	routers.FileUpRouter(r)
	routers.StudentRouter(r)

	r.Run(":8080")

}
