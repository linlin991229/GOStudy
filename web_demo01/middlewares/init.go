package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("into")
	fmt.Println("url:", c.Request.URL)
	// 获取请求头信息
	fmt.Println(c.GetHeader("Authorization"))

	c.Set("middlewareData", "middlewareData")
	// 下一个处理器
	c.Next()
	// 会终止下一个处理器执行
	// c.Abort()

	// 使用协程注意！！！！
	cCp := c.Copy()
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("协程执行完毕", cCp.Request.URL.Path)
	}()
	// 处理完毕后，执行
	fmt.Println("out,执行耗时:", time.Now().UnixNano()-start, "ns")
}
