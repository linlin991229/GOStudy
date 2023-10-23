package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Grade string
}

func (stu Student) getName() string {
	stu.Name = "李四"
	return stu.Name
}
func (stu *Student) getNamePointer() string {
	stu.Name = "李四"
	return stu.Name
}
func main() {
	stu := Student{
		Name:  "张三",
		Age:   18,
		Grade: "三年级",
	}
	// t := reflect.TypeOf(stu)
	// v := reflect.ValueOf(stu)
	// fmt.Println("v:", v)
	// fmt.Println("t:", t)

	fmt.Println("stu:", stu)

	fmt.Println("stu.String():", stu.getName())
	// fmt.Println("stu.StringPointer():", stu.getNamePointer())

	fmt.Println("student:", Student{
		Name:  "张三",
		Age:   18,
		Grade: "三年级",
	}.getName())
	fmt.Printf("stu.String(): %v\n", stu.getName())
	fmt.Println("stu:", stu)
	stu.getNamePointer()
	// Student{"张三", 18, "三年级"}
	fmt.Printf("stu: %v\n", stu)
}
