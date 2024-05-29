package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	//log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}
	// Request a greeting message.
	messages, err := greetings.PrintGreetingsMessageToManyPeople(names)
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}
	/*for _, value := range messages {
		fmt.Print(value)
	}*/
	// Get a greeting message and print it.

	fmt.Println(messages)
}
