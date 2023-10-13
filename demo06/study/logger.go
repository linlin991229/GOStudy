package study

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func TestLogger() {
	log.Println("标准日志")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("有文件/行号的日志")

	myLog := log.New(os.Stdout, "my:", log.LstdFlags)
	myLog.Println("from myLog")

	myLog.SetPrefix("myPrefix:")
	myLog.Println("from myLog")

	var buf bytes.Buffer

	bufLog := log.New(&buf, "buf:", log.LstdFlags)
	// 输出到缓冲区
	bufLog.Println("hello")
	fmt.Println("from bufLog:", buf.String())

	// 以JSON格式输出
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	myslog.Info("hello again", "key", "val", "age", 25)
}
