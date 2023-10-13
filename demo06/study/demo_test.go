package study

import (
	"fmt"
	"testing"
)

// 文件名必须以_test.go结尾
// 测试函数必须以Test开头
// 参数必须是*testing.T
func TestReverseRunes(t *testing.T) {

	if false {
		fmt.Println("true")
	} else {
		t.Error("测试失败")
	}
}
