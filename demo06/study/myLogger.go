package study

import (
	"io"
	"log"
	"os"
)

var (
	Trace   *log.Logger // 所有日志
	Info    *log.Logger // 普通日志
	Warning *log.Logger // 警告日志
	Error   *log.Logger // 错误日志
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}
	// 会忽略所有写入这一变量的数据。当某个等级的日志不重要时，使用Discard变量可以禁用这个等级的日志。
	Trace = log.New(io.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	// io.MultiWriter()可以接收多个io.Writer，返回一个io.Writer。这个返回的io.Writer把所有传入的io.Writer绑定在一起。
	// 当对这个io.Writer进行写入时，所有传入的io.Writer都会被写入
	// 因此这里的日志会被写入到文件和标准输出
	Error = log.New(io.MultiWriter(file, os.Stdout), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func TestMyLogger() {
	Trace.Println("我是trace")
	Info.Println("我是info")
	Warning.Println("我是warning")
	Error.Println("我是error")
}
