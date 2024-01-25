package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"lin.com/im/common"
	"lin.com/im/models"
	"lin.com/im/services"
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
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param user body models.UserBase true "创建用户"
// @Success 200 {object} common.LinResult
// @Router /user/createUser [post]
func (u UserController) CreateUser(c *gin.Context) {
	user := models.UserBase{}
	c.BindJSON(&user)
	if ok, err := (services.UserService{}.CreateUser(&user)); ok {
		c.JSON(http.StatusCreated, common.LinResult{
			Code: 200,
			Msg:  "创建成功",
			Data: user,
		})
	} else {
		c.JSON(http.StatusCreated, common.LinResult{
			Code: 200,
			Msg:  "创建失败" + err.Error(),
			Data: user,
		})
	}

}

// UpdateUser
// @Summary 更新用户
// @Schemes http https
// @Description 更新用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param user body models.UserBase true "更新用户"
// @Success 200 {object} common.LinResult
// @Router /user/updateUser [put]
func (u UserController) UpdateUser(c *gin.Context) {
	user := models.UserBase{}
	c.ShouldBind(&user)
	if ok, err := (services.UserService{}.UpdateUser(&user)); ok {
		c.JSON(http.StatusOK, common.LinResult{
			Code: 200,
			Msg:  "更新成功",
			Data: user,
		})
	} else {
		c.JSON(http.StatusOK, common.LinResult{
			Code: 200,
			Msg:  "更新失败" + err.Error(),
			Data: user,
		})
	}
}

// GetUser
// @Summary 查询用户
// @Schemes http https
// @Description 查询用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id query string true "查询用户"
// @Success 302 {object} common.LinResult
// @Router /user/getUser/{id} [get]
func (u UserController) GetUser(c *gin.Context) {
	userId := c.Query("id")
	if user, err := (services.UserService{}.GetUser(userId)); err == nil {
		c.JSON(http.StatusFound, common.LinResult{
			Code: 200,
			Msg:  "查询成功",
			Data: user,
		})
	} else {
		c.JSON(http.StatusNotFound, common.LinResult{
			Code: 200,
			Msg:  "查询失败" + err.Error(),
			Data: user,
		})
	}
}

// GetUsers
// @Summary 查询所有用户
// @Schemes http https
// @Description 查询所有用户
// @Tags 用户管理
// @Produce json
// @Success 200 {object} common.LinResult
// @Router /user/getUsers [get]
func (u UserController) GetUsers(c *gin.Context) {
	if users, err := (services.UserService{}.GetUsers()); err == nil {
		c.JSON(http.StatusCreated, common.LinResult{
			Code: 200,
			Msg:  "查询成功",
			Data: users,
		})
	} else {
		c.JSON(http.StatusCreated, common.LinResult{
			Code: 200,
			Msg:  "查询失败" + err.Error(),
			Data: users,
		})
	}
}

// DeleteUser
// @Summary 删除用户
// @Schemes http https
// @Description 删除用户
// @Tags 用户管理
// @Produce json
// @Param id path string true "删除用户"
// @Success 200 {object} common.LinResult
// @Router /user/deleteUser/{id} [delete]
func (u UserController) DeleteUserById(c *gin.Context) {
	userId := c.Param("id")
	if ok, err := (services.UserService{}.DeleteUserById(userId)); ok {
		c.JSON(http.StatusCreated, common.LinResult{
			Code: 200,
			Msg:  "删除成功",
			Data: userId,
		})
	} else {
		c.JSON(http.StatusCreated, common.LinResult{
			Code: 200,
			Msg:  "删除失败" + err.Error(),
			Data: userId,
		})
	}
}
