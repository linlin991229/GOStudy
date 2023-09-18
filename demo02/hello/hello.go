package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("greetings: ")
	// message, err := greetings.Hello("")
	// message, err := greetings.Hello("Gladys")
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(message)
	fmt.Println(messages)
}
