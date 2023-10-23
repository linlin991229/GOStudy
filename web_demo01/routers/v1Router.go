package routers

import (
	"github.com/gin-gonic/gin"
	"lin.com/web_demo01/controller"
	"lin.com/web_demo01/middlewares"
)

func V1Router(router *gin.Engine) {
	// 简单的路由组：v1
	v1 := router.Group("/v1")
	// 路由分组中的中间件配置
	// router.Group("/v1", initMiddleware, initMiddlewareTwo)
	// v1.Use(initMiddleware, initMiddlewareTwo)

	{
		v1.GET("/hello", controller.Hello)
		v1.GET("/hello2", controller.Hello2)

		v1.GET("/middleware1", middlewares.InitMiddleware, controller.Hello2)

	}
}
