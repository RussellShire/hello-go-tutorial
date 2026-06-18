package main

import (
	"fmt"

	"example/hello-go-tutorial/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
