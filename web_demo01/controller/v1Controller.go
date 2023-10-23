package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"say": "world",
	})
}

func Hello2(c *gin.Context) {
	fmt.Println("Hello2 处理中。。。")
	if v, ok := c.Get("middlewareData"); ok {
		fmt.Println("获取上一个中间件数据:", v, ok)
		c.JSON(200, gin.H{
			"hello": v,
		})
	} else {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	}
}
