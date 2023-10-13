package study

import (
	"fmt"
	"os"
	s "strings"
)

var p = fmt.Println

func TestStrings() {
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}

type point struct {
	x, y int
}

func TestFormat() {
	p := point{1, 2}
	// struct1: {1 2}
	fmt.Printf("struct1: %v\n", p)
	// struct2: {x:1 y:2}，包括结构体字段
	fmt.Printf("struct2: %+v\n", p)
	// struct3: study.point{x:1, y:2}，源码形式
	fmt.Printf("struct3: %#v\n", p)
	// 打印类型用%T
	fmt.Printf("Type: %T Value: %v\n", p, p)

	fmt.Printf("bool: %t\n", true)

	fmt.Printf("int: %d\n", 123)

	fmt.Printf("bin: %b\n", 2)

	fmt.Printf("hex: %x\n", 15)

	fmt.Printf("oct: %o\n", 9)

	fmt.Printf("float1: %f\n", 1.2)

	fmt.Printf("float2: %e\n", 12340000.0)

	fmt.Printf("float3: %E\n", 12340000.0)

	fmt.Printf("char: %c\n", 33)

	fmt.Printf("string: %s\n", "\"string\"")

	fmt.Printf("string: %q\n", "\"string\"")
	// %x以 16 为进制呈现字符串，每个输入字节有两个输出字符。
	fmt.Printf("string: %x\n", "hex this")

	fmt.Printf("pointer: %p\n", &p)

	// 默认情况下，结果将右对齐并用空格填充
	fmt.Printf("width: %6.2f\n", 1.2)

	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
	// 左对齐
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 10.2, 3.45)

	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")
	// 返回格式化后的字符串而不打印
	s := fmt.Sprintf("a %s", "string")
	fmt.Printf("%s\n", s)

	bytes, _ := fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
	fmt.Printf("bytes: %d\n", bytes)
}
