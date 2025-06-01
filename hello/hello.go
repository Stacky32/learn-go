package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Configure logging
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Jim", "Sally", "Rosie"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
