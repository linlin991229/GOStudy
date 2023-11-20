package study

import (
	"fmt"
	"time"

	"lin.com/study/utils"
)

type User struct {
	Name string
	Age  int
}

// 值接收者
func (u User) show() {
	fmt.Println("Name:", u.Name, "Age:", u.Age)
}
func (u *User) showPoint() {
	fmt.Println("showPoint Name:", u.Name, "Age:", u.Age)
}

// 指针接收者
func (u *User) changeAge(age int) {
	u.Age = age
}

func TestStruct() {
	user := User{
		Name: "John Doe",
		Age:  25,
	}
	showUser(user)
	user.show()

	// 会被解析为(&user).show()
	user.changeAge(30)
	user.showPoint()
	user.show()

	var user2 *User = &user

	// 会被解析为(*user2).show()
	user2.show()
	user2.changeAge(35)
	user2.show()

}
func showUser(user User) {
	user.Name = "Bob"
	fmt.Println("showUserFunc Name:", user.Name, "Age:", user.Age)
}

type TestStruct2 struct {
	Hello string
}
type TestStruct1 struct {
	Name string
	TestStruct2
}
type Model struct {
	Name string
	Age  int

	Reference *User

	Date utils.LocalDateTimeFormat
}
type MyTime time.Time

func (t MyTime) String() string {
	// return time.Time(t).String()
	// return time.Time(t).Format("2006-01-02 15:04:05")
	return fmt.Sprintf("%s", time.Time(t).Format("2006-01-02 15:04:05"))

}
