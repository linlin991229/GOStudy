package routers

import (
	"github.com/gin-gonic/gin"
	"lin.com/web_demo01/controller"
)

func StudentRouter(router *gin.Engine) {
	stuRouter := router.Group("/students")
	{
		stuRouter.GET("/", controller.StudentController{}.GetStudents)
		stuRouter.GET("/:id", controller.StudentController{}.GetStudent)
		stuRouter.POST("/", controller.StudentController{}.SaveStudent)
		stuRouter.PUT("/", controller.StudentController{}.UpdateStudent)
		stuRouter.DELETE("/:id", controller.StudentController{}.DeleteStudent)
	}
}
