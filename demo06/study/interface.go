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

type Show interface {
	// 值接收者方法
	show()
}
type ShowPoint interface {
	showPoint()
}

func TestInterface2() {
	user := User{
		Name: "John Doe",
		Age:  25,
	}
	show(user)
	// 指针接收者实现的接口要使用引用
	show2(&user)

	// Duration.pretty() 错误
	d := Duration(10)
	fmt.Printf("d.pretty(): %v\n", d.pretty())
	fmt.Printf("Duration(20).pretty2(): %v\n", Duration(20).pretty2())
	fmt.Println("d.pretty2():", (&d).pretty2())
}

func show(s Show) {
	s.show()
}
func show2(s ShowPoint) {
	s.showPoint()
}

type Duration int

func (d *Duration) pretty() string {
	return fmt.Sprintf("Duration: %d", *d)
}
func (d Duration) pretty2() string {
	return fmt.Sprintf("Duration: %d", d)
}
