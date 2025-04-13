package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Teja")
	fmt.Println(message)
}
