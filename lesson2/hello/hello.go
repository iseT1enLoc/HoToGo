package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.PrintGreetingsMessage("Welcome to your appearance", "Peter")
	fmt.Println(message)
}
