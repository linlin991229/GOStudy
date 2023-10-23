package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 推荐使用结构体进行分离，封装
type V2Controller struct{}

func (v2 V2Controller) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"say": "hello V2",
	})
}

func (v2 V2Controller) Hello2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "world V2",
	})
}
