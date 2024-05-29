package greetings

import "fmt"

// Hello returns a greeting for the named person.
func PrintGreetingsMessage(msg string, emp_name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi %v, %v!", emp_name, msg)
	return message
}
