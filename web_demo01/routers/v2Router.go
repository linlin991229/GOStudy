package routers

import (
	"github.com/gin-gonic/gin"
	"lin.com/web_demo01/controller"
	"lin.com/web_demo01/middlewares"
)

func V2Router(router *gin.Engine) {
	// 简单的路由组: v2
	v2 := router.Group("/v2", middlewares.InitMiddleware)
	{
		v2.GET("/hello", controller.V2Controller{}.Hello)
		v2.GET("/hello2", controller.V2Controller{}.Hello2)
	}
}
