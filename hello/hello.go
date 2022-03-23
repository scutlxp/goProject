package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	message, err := greetings.Hello("lxp")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	names := []string{"lxp", "dbb", "lyb", "hz"}
	messages, errs := greetings.Hellos(names)

	if errs != nil {
		log.Fatal(errs)
	}

	fmt.Println(messages)
}
