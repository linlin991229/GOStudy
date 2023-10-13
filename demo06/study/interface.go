package study

import "fmt"

// struct嵌入
type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// 嵌入struct
type container struct {
	base
	str string
}

func TestInterface() {
	// co := container{base: base{num: 1}, str: "hello"}
	co := container{base{1}, "some name"}

	fmt.Printf("co={num:%v,str:%v}\n", co.num, co.str)
	fmt.Println("also num:", co.base.num)
	fmt.Println("describe:", co.describe())

	type describe interface {
		describe() string
	}

	// d是一个接口类型变量，可以存储任意实现了describe接口的具体类型的值
	var d describe = co
	// 接口断言
	if desc, ok := d.(describe); ok {
		fmt.Println("describe:", desc.describe())
	}
}
