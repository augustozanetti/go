package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main(){
	log.SetPrefix("grettings: ")
	log.SetFlags(0)

	// message, err := greetings.Hello("Gladys")

	names := []string{"Gladys", "Joe", "Nic"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}