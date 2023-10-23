package routers

import (
	"github.com/gin-gonic/gin"
	"lin.com/web_demo01/controller"
)

func FileUpRouter(router *gin.Engine) {
	// 简单的路由组: v2
	fileUp := router.Group("/file")
	{
		fileUp.POST("/upload", controller.FileUpController{}.Upload)
		fileUp.POST("/multiUpload", controller.FileUpController{}.MultiUpload)
	}
}
