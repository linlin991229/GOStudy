package routers

import (
	"github.com/gin-gonic/gin"
	"lin.com/im/controller"
)

func UserRouter(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.GET("/test", controller.UserController{}.HandlerUserTest)
		user.POST("/createUser", controller.UserController{}.CreateUser)
	}
}
