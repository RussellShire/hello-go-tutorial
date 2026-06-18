package main

import (
	"fmt"
	"log"

	"example/hello-go-tutorial/greetings"
)

func main() {
	log.SetPrefix("greetings: ") // Sets prefix to log (overwriting default of date and time)
	log.SetFlags(0)              // Disables default log dates

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
