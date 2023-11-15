package main

import (
	"fmt"

	"lin.com/study/study"
)

// main is the entry point of the program.
//
// It prints "霖霖" to the console.
func main() {
	fmt.Println("霖霖")

	// study.TestHttpServer()
	// study.TestRateLimit()
	// study.TestConcurrentPrime()

	study.TestLogger()
}
