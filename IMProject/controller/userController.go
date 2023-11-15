package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"lin.com/im/common"
)

type UserController struct{}

// @BasePath /api/v1

// swagger测试例子 godoc
// @Summary swagger测试
// @Schemes
// @Description swagger测试描述
// @Tags swagger测试标签
// @Accept json
// @Produce json
// @Success 200 struct HandlerUserTest
// @Router /user/test [get]
func (u UserController) HandlerUserTest(c *gin.Context) {
	log.Println("HandlerUserTest")
	c.JSON(http.StatusOK, common.LinResult{
		Code: 200,
		Msg:  "success",
		Data: "HandlerUserTest",
	})
}
