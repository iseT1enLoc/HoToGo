package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func PrintGreetingsMessage(msg string, emp_name string) (string, error) {
	if emp_name == "" || msg == "" {
		return "", errors.New("The msg or employee name might be empty")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi %v, %v!.%v\n", emp_name, msg, GiveRandomWishes())
	return message, nil
}

func PrintGreetingsMessageToManyPeople(names []string) (map[string]string, error) {
	//initialize massage map
	map_messages := make(map[string]string)

	for _, name := range names {
		message, err := PrintGreetingsMessage("Nice to meet you", name)
		if err != nil {
			return nil, err
		}
		map_messages[name] = message
	}
	return map_messages, nil
}

func GiveRandomWishes() string {
	formats := []string{
		"Have a good day",
		"Have a productive day",
		"May your work run well",
	}
	return formats[rand.Intn(len(formats))]
}
