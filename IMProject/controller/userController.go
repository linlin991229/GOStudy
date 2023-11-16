package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"lin.com/im/common"
	"lin.com/im/models"
)

type UserController struct{}

// @BasePath /api/v1

// swagger测试例子 godoc
// @Summary swagger测试
// @Schemes http https
// @Description swagger测试描述
// @Tags swagger测试标签
// @Accept json
// @Produce json
// @Success 200 {object} common.LinResult
// @Router /user/test [get]
func (u UserController) HandlerUserTest(c *gin.Context) {
	log.Println("HandlerUserTest")
	c.JSON(http.StatusOK, common.LinResult{
		Code: 200,
		Msg:  "success",
		Data: "HandlerUserTest",
	})
}

// CreateUser
// @Summary 创建用户
// @Schemes http https
// @Description 创建用户
// @Tags 首页
// @Accept json
// @Produce json
// @Param user body models.UserBase true "创建用户"
// @Success 200 {object} common.LinResult
// @Router /user/createUser [post]
func (u UserController) CreateUser(c *gin.Context) {
	user := models.UserBase{}

	c.BindJSON(&user)
	c.JSON(http.StatusCreated, common.LinResult{
		Code: 200,
		Msg:  "success",
		Data: user,
	})
}
