package study

import (
	"fmt"
)

func mayPanic() {
	panic("a problem")
}

func TestRecover() {
	defer func() {
		// recover必须在延迟函数中调用。当封闭函数发生恐慌时，延迟将激活，并且recover其中的调用将捕获恐慌。
		// recover的返回值是调用中引发的错误panic。
		if r := recover(); r != nil {
			fmt.Println("Recovered Error:\n", r)
		}
	}()
	mayPanic()
	fmt.Println("After mayPanic")
}
