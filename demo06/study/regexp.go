package study

import (
	"bytes"
	"fmt"
	"regexp"
)

func TestRegexp() {
	match, _ := regexp.MatchString("p([a-z]+)", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))
	// 找到第一个匹配
	fmt.Println(r.FindString("peach punch"))
	fmt.Println("idx: ", r.FindStringIndex("peach punch"))
	// 会查找匹配模式中的子匹配信息，如p([a-z]+)ch会返回([a-z]+)匹配到的字符串
	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// 第二个非负参数用于限制匹配的次数
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println("all: ", r.FindAllStringIndex("peach punch pinch", -1))

	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	fmt.Println(r.FindAllString("peach punch pinch", 2))
	fmt.Println("all: ", r.FindAllStringIndex("peach punch pinch", 2))

	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp: ", r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	// 该regexp包还可用于将字符串子集替换为其他值。
	// 函数用于转换匹配的字符串。
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}

func TestArray() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(arr)
	strs := []string{"aaa", "bbb"}
	fmt.Println(strs)
	// 字符串转换为字节切片
	bytes := []byte("a b")
	fmt.Println(bytes)
}
