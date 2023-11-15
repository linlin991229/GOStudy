package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"lin.com/im/docs"
	"lin.com/im/routers"
)

// @title           IM即时通信
// @version         1.0
// @description     一个go的即时通信项目.
// @termsOfService  API服务条款

// @contact.name   霖霖
// @contact.url    http://www.swagger.io/support
// @contact.email  1589861957@qq.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /lin/im/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	routers.UserRouter(r)

	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run("localhost:8080")
}
