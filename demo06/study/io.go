package study

import (
	"bytes"
	"fmt"
	"os"
)

func TestReadWrite() {
	// 创建一个Buffer值，并将一个字符串写入Buffer
	// 使用实现io.Writer的Write方法
	var buf bytes.Buffer
	buf.Write([]byte("Hello"))

	// 使用Fprintf来将一个字符串拼接到Buffer里
	// 将bytes.Buffer的地址作为io.Writer类型值传入
	fmt.Fprintf(&buf, "World!")

	// 将Buffer的内容输出到标准输出设备
	// 将os.File值的地址作为io.Writer类型值传入
	buf.WriteTo(os.Stdout)
}
