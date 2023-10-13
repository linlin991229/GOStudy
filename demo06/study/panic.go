package study

import "os"

func TestPanic() {
	panic("a problem")
	// 引发panic下面代码不会执行
	_, err := os.Create("/tmp/file")

	if err != nil {
		panic(err)
	}

}
