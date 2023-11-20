package main

import "lin.com/study/dao"

// main is the entry point of the program.
//
// It prints "霖霖" to the console.
func main() {
	// study.TestHttpServer()
	// study.TestRateLimit()
	// study.TestConcurrentPrime()
	// testStruct := study.TestStruct1{
	// 	Name: "lin",
	// 	TestStruct2: study.TestStruct2{
	// 		Hello: "hello",
	// 	},
	// }
	// fmt.Println("testStruct:", testStruct.Hello)
	// study.TestLogger()

	// user := study.Model{
	// 	Name: "John Doe",
	// 	Age:  25,
	// 	Reference: &study.User{
	// 		Name: "Bob",
	// 		Age:  30,
	// 	},
	// 	Date: utils.LocalDateTimeFormat(time.Now()),
	// }

	// fmt.Println("user:", user)
	// fmt.Printf("user:%v \n", user)
	// fmt.Printf("user:%+v \n", user)
	// fmt.Printf("user:%#v \n", user)

	dao.DBGorm()
}
